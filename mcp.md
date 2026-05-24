# Codex MCP 使用笔记

更新时间：2026-05-24。

本文记录 Codex 中 MCP（Model Context Protocol）的使用方式、基本原理和服务端设计思路。第一版以 Codex 使用为主，不覆盖 Claude Code 的 MCP 配置差异。

## 结论

MCP 适合把外部系统、文档、数据库、浏览器、内部工具或本地脚本，以标准化工具能力接入 Codex。Codex 负责作为 MCP client 管理和调用 MCP server；MCP server 负责暴露工具、资源或提示模板。

推荐分工：

- 命令速查放在 [Codex 核心指令](./codex-core-commands.md)。
- MCP 原理、配置、认证、设计和排查放在本文。
- 真实密钥、Token、私有代理地址和内部服务 URL 不写入仓库文档或配置示例。

## MCP 是什么

MCP 是一套让 AI 客户端连接外部能力的协议。它解决的核心问题是：不要为每个 AI 工具单独写一套接入方式，而是让外部能力通过统一协议暴露给客户端。

基本角色：

| 角色 | 职责 |
| --- | --- |
| MCP client | 发起连接、发现能力、把可用工具交给模型使用。Codex 在这里通常是 client。 |
| MCP server | 暴露外部能力，例如查询文档、访问数据库、调用内部 API 或执行受控脚本。 |
| Tool | 可被模型请求调用的动作，适合有副作用或需要参数的操作。 |
| Resource | 可读取的上下文数据，适合文档、文件、配置、知识库条目等内容。 |
| Prompt | 服务端提供的提示模板，适合沉淀固定任务流程。 |

常见调用链路：

```text
用户请求 -> Codex 判断需要外部能力 -> MCP client 发现并调用 tool/resource/prompt
       -> MCP server 执行受控操作 -> 返回结构化结果 -> Codex 汇总并继续任务
```

MCP 底层使用 JSON-RPC 消息。官方 MCP 文档当前定义的标准传输方式包括 `stdio` 和 `Streamable HTTP`，其中 `stdio` 适合本地进程，`Streamable HTTP` 适合远程或常驻服务。参考：[MCP Transports](https://modelcontextprotocol.io/docs/concepts/transports)。

## Codex MCP 命令

Codex CLI 当前提供两类 MCP 相关命令：

| 命令 | 作用 |
| --- | --- |
| `codex mcp` | 管理外部 MCP server 配置。 |
| `codex mcp-server` | 将 Codex 作为 stdio MCP server 启动。 |

查看已配置 server：

```bash
codex mcp list
codex mcp list --json
```

查看单个 server：

```bash
codex mcp get <name>
codex mcp get <name> --json
```

移除 server：

```bash
codex mcp remove <name>
```

### 配置落盘位置与格式

`codex mcp add ...` 会把 MCP server 注册信息写入本机 Codex 配置文件：

```text
~/.codex/config.toml
```

配置通常以 `[mcp_servers.<name>]` 作为表名。以 stdio MCP server 为例：

```toml
[mcp_servers.local-docs]
command = "node"
args = ["/path/to/mcp-server.js"]
```

如果注册 CodeGraph：

```bash
codex mcp add codegraph -- codegraph mcp
```

对应配置通常类似：

```toml
[mcp_servers.codegraph]
command = "codegraph"
args = ["mcp"]
```

日常推荐通过 `codex mcp add/get/remove` 管理配置，不直接手写 `~/.codex/config.toml`。如果确实需要排查，可以用：

```bash
codex mcp get <name>
codex mcp get <name> --json
```

手写或分享配置时，不要写入真实 Token、私有代理地址、内部服务 URL 或个人机器上的绝对路径。需要固定路径时，优先使用 `/path/to/...` 这类占位符示例。

### 添加 stdio MCP server

`stdio` 模式下，Codex 会把 MCP server 作为子进程启动，通过标准输入输出交换 JSON-RPC 消息。它适合本地脚本、本地 CLI、轻量开发工具和只在当前机器使用的集成。

命令格式：

```bash
codex mcp add <name> -- <command> [args...]
```

示例：

```bash
codex mcp add local-docs -- node /path/to/mcp-server.js
```

如果 server 启动需要环境变量：

```bash
codex mcp add local-docs \
  --env DOCS_ROOT=/path/to/docs \
  -- node /path/to/mcp-server.js
```

注意事项：

- server 的 `stdout` 只能输出合法 MCP 消息，日志应写到 `stderr`。
- 不要把真实 Token 直接写进命令示例或仓库文档。
- 本地 server 默认应只访问必要目录，不要把整个用户主目录暴露给工具。

### 添加 Streamable HTTP MCP server

`Streamable HTTP` 模式下，MCP server 是独立 HTTP 服务。它适合远程服务、团队共享服务、需要 OAuth 或 bearer token 的场景。

命令格式：

```bash
codex mcp add <name> --url <url>
```

示例：

```bash
codex mcp add docs-server --url https://example.com/mcp
```

如果服务需要 bearer token，建议让 Codex 从环境变量读取：

```bash
codex mcp add docs-server \
  --url https://example.com/mcp \
  --bearer-token-env-var DOCS_MCP_TOKEN
```

使用规则：

- token 存在本机环境变量或安全凭据系统中，不写入仓库。
- HTTP MCP server 应启用认证、TLS 和最小权限。
- 本地 HTTP server 应绑定 `127.0.0.1`，避免暴露到局域网或公网。

### OAuth 登录与退出

部分 HTTP MCP server 支持 OAuth。添加 server 后，可以通过 Codex 发起登录：

```bash
codex mcp login <name>
```

指定 OAuth scopes：

```bash
codex mcp login <name> --scopes scope-a,scope-b
```

退出登录：

```bash
codex mcp logout <name>
```

OAuth 适合用户级授权，例如文档平台、任务系统、代码托管平台等。服务端应只请求任务所需 scope，不要默认申请过宽权限。

## MCP 调用链路

不同 MCP server 的业务能力不同，但 Codex 使用它们的基本链路相同：

```text
用户请求
-> Codex 判断需要外部能力
-> Codex 读取 ~/.codex/config.toml 中的 mcp_servers.<name>
-> Codex 作为 MCP client 启动或连接 MCP server
-> Codex 与 MCP server 通过 JSON-RPC 交换消息
-> MCP server 访问它背后的工具、索引、文档、数据库或远程服务
-> MCP server 返回结构化结果
-> Codex 基于结果继续分析、计划或修改代码
```

两类传输的差异主要在连接方式：

| 传输方式 | 链路特点 | 适用场景 |
| --- | --- | --- |
| `stdio` | Codex 启动本地子进程，通过 stdin/stdout 交换 JSON-RPC 消息。 | 本地 CLI、本地索引工具、本地脚本。 |
| `Streamable HTTP` | Codex 连接一个 HTTP MCP endpoint，通过 HTTP 传输 JSON-RPC 消息。 | 远程服务、团队共享服务、需要 OAuth 或 bearer token 的服务。 |

因此，CodeGraph、文档检索、数据库只读查询等 MCP 的接入原理一致：Codex 不直接理解这些系统的内部存储格式，而是通过 MCP server 暴露的 tool/resource/prompt 获取受控结果。

## 什么时候适合使用 MCP

适合接入 MCP：

- Codex 需要稳定访问外部系统，例如文档库、issue 系统、数据库只读查询或内部工具。
- 同一能力会在多个项目或多次会话中复用。
- 能力边界清晰，可以用少量结构化参数描述。
- 需要把工具能力和模型对话解耦，避免每次都复制大量上下文。

不建议接入 MCP：

- 一次性任务，用 shell 命令或直接读取文件即可完成。
- 工具会执行高风险写操作，但没有权限隔离、审计或确认机制。
- server 输出不稳定，容易把日志、调试文本或大段无关内容混入结果。
- 需要提交真实密钥、私有地址或个人 Token 才能复现。

## 设计 MCP server 的思路

设计 MCP server 时，优先把它当成“给 AI 使用的受控 API”，而不是把现有内部接口原样暴露出去。

### 工具边界

一个 tool 只做一件明确的事：

- 好：`search_docs(query, limit)`、`get_issue(id)`、`run_readonly_sql(query)`。
- 差：`do_anything(input)`、`execute_shell(command)`、`admin_api(payload)`。

工具参数应尽量结构化，避免让模型拼接复杂命令或自由格式 JSON。返回值应短、稳定、可解释，必要时包含下一步建议或错误原因。

### 权限边界

默认按最小权限设计：

- 读操作和写操作分开。
- 高风险写操作需要显式确认或额外授权。
- 按项目、目录、租户或账号隔离可访问范围。
- 不把原始密钥、内部 URL、数据库连接串返回给模型。

### 失败处理

MCP server 应把失败分成可操作的错误：

- 参数错误：说明哪个参数缺失或格式不对。
- 权限错误：说明缺少哪类授权，但不泄露敏感策略。
- 外部系统错误：说明依赖服务不可用或超时。
- 空结果：明确返回未找到，而不是抛出模糊异常。

### 上下文控制

工具返回内容越多，越容易稀释模型上下文。建议：

- 默认限制结果数量，例如 `limit=10`。
- 长文档先返回摘要和定位信息，再按需读取详情。
- 搜索结果返回标题、路径、更新时间和短摘要。
- 对二进制、大文件、日志流等内容提供分页或过滤参数。

## 安全注意事项

MCP 扩展了 Codex 可访问的外部能力，因此安全边界要前置设计。

最低要求：

- 不提交真实 API Key、Token、私有代理地址或内部服务 URL。
- 远程 MCP server 必须启用认证，生产环境必须使用 TLS。
- 本地 HTTP server 优先绑定 `127.0.0.1`。
- Streamable HTTP server 应校验 `Origin`，防止 DNS rebinding 风险。
- 工具描述要准确，避免名称相似的高风险工具伪装成低风险工具。
- 对写操作、删除操作、外部通知和资金相关操作设置二次确认。

## 排查清单

### `codex mcp` 命令不可用

先确认 Codex 版本和帮助信息：

```bash
codex --version
codex --help
```

如果 `codex --help` 中没有 `mcp`，说明当前 Codex CLI 版本不支持该命令或安装路径不是预期版本。先升级或确认实际执行的是哪个 `codex`。

### server 没有出现在列表中

检查配置：

```bash
codex mcp list
codex mcp list --json
```

如果未出现，重新执行 `codex mcp add ...`，并确认 `<name>` 没有写错。

### stdio server 启动失败

重点检查：

- `<command>` 是否存在并可执行。
- 工作目录和参数是否正确。
- 依赖是否已安装。
- server 是否把日志写到了 `stdout`，导致 MCP 消息被污染。
- 需要的环境变量是否通过 `--env KEY=VALUE` 传入。

### HTTP server 连接失败

重点检查：

- `--url` 是否指向 MCP endpoint。
- 服务是否正在运行。
- 网络、代理、TLS 证书是否正常。
- bearer token 环境变量是否存在。
- server 是否要求 OAuth 登录。

### 工具不可见或不可用

重点检查：

- server 是否成功启动并完成初始化。
- server 是否正确声明 tools/resources/prompts。
- 当前账号或 token 是否有权限访问该能力。
- 工具描述是否过于模糊，导致模型不知道何时使用。

## 推荐落地顺序

1. 先用 `codex mcp list` 查看当前已有配置。
2. 从只读、低风险 MCP server 开始接入。
3. 使用 `codex mcp get <name>` 确认配置是否符合预期。
4. 在一次真实任务中验证工具是否能稳定返回结果。
5. 再逐步接入需要认证、远程服务或写操作的 server。

## CodeGraph 类工具接入模式

CodeGraph 这类代码索引工具适合采用“两段式”接入：

```text
CLI / watch / CI 构建索引 -> MCP server 暴露查询能力 -> Codex 通过 MCP 查询索引
```

索引构建是确定性的工程动作，建议由用户命令、watch 进程、CI 或 git hook 控制；Codex 主要通过 MCP 消费已有索引，用于查询符号、调用链、影响面和模块关系。

具体 CodeGraph 安装、建索引、gitignore 和 Codex MCP 接入流程见 [CodeGraph 使用笔记](./external-tools/codegraph.md)。
