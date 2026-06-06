# Next AI Draw.io 使用笔记

更新时间：2026-06-06。

本文以 `DayuanJiang/next-ai-draw-io` 官方 README 和 `packages/mcp-server/README.md` 为准，整理安装、基础使用、MCP 接入和与 CodeGraph / Understand Anything 的搭配方式。

参考来源：

- [DayuanJiang/next-ai-draw-io](https://github.com/DayuanJiang/next-ai-draw-io)
- [Next AI Draw.io MCP Server README](https://github.com/DayuanJiang/next-ai-draw-io/blob/main/packages/mcp-server/README.md)
- [Snyk: @next-ai-drawio/mcp-server](https://security.snyk.io/package/npm/%40next-ai-drawio%2Fmcp-server)

## 结论

`next-ai-draw-io` 推荐作为 AI 工作流里的出图工具使用：

```text
CodeGraph / Understand Anything / 源码验证
-> 整理架构事实、模块依赖、流程步骤
-> next-ai-draw-io MCP 生成或编辑 draw.io 图
-> 导出 .drawio 文件
```

不要把 `next-ai-draw-io` 当成代码理解工具。它更适合把已经整理好的架构、流程、依赖关系渲染成 draw.io 图。

## 解决什么问题

`next-ai-draw-io` 是一个基于 Next.js 的 AI draw.io 图表工具，可以通过自然语言生成、修改和增强 draw.io 图。适合：

- 生成项目架构图。
- 生成业务流程图、认证流程图、调用流程图。
- 生成云架构图，例如 AWS、GCP、Azure。
- 根据已有文字说明、PDF 或文本内容生成图。
- 通过 MCP 让 Claude、Cursor、VS Code、Codex 等 AI 客户端调用出图能力。

它不是测试、静态分析或代码图谱工具。涉及代码事实时，应先用源码、测试、CodeGraph 和 Understand Anything 验证。

## 安装方式

### 前置条件

本地 MCP 接入需要 Node.js、npm 和 npx：

```bash
node --version
npm --version
npx --version
```

如果缺少 Node.js，可先参考仓库中的 [Codex CLI 使用笔记](../codex-cli.md) 配置 Node.js 环境。

### 在线使用

如果只是临时试用，可以访问官方 demo：

```text
https://next-ai-drawio.jiang.jp/
```

demo 支持在浏览器设置中配置自己的模型供应商和 API Key。不要在团队共享机器或不可信环境中填入真实密钥。

### 桌面应用

官方 Releases 页面提供 Windows、macOS、Linux 桌面应用。适合只想手动出图、不需要接 MCP 的场景。

```text
https://github.com/DayuanJiang/next-ai-draw-io/releases
```

### 本地源码运行

如果要本地运行完整 Web 应用：

```bash
git clone https://github.com/DayuanJiang/next-ai-draw-io
cd next-ai-draw-io
npm install
cp env.example .env.local
npm run dev
```

然后打开：

```text
http://localhost:6002
```

`.env.local` 中的模型供应商、模型名和 API Key 按自己的供应商配置。不要把真实 `.env.local` 提交到 Git。

### Docker 运行

官方 README 提供 Docker 方式。适合只想本地跑 Web 应用、不想安装前端依赖的场景：

```bash
docker run -d -p 3000:3000 \
  -e AI_PROVIDER=openai \
  -e AI_MODEL=gpt-4o \
  -e OPENAI_API_KEY=your_api_key \
  ghcr.io/dayuanjiang/next-ai-draw-io:latest
```

也可以使用环境变量文件：

```bash
cp env.example .env
docker run -d -p 3000:3000 --env-file .env ghcr.io/dayuanjiang/next-ai-draw-io:latest
```

这里的 `your_api_key` 只是占位符，实际使用时换成自己的供应商配置，并确保 `.env` 不进入 Git。

## 基础使用

直接在 Web 应用或 MCP 客户端里用自然语言描述要画的图：

```text
Create a flowchart showing user authentication with login, MFA, and session management.
```

中文提示也可以：

```text
生成一个用户登录流程图，包含账号密码登录、MFA 校验、会话创建、失败重试和异常退出。
```

项目架构图建议先让 AI 整理事实，再交给 `next-ai-draw-io` 出图：

```text
请根据以下架构摘要生成 draw.io 架构图：

入口层：Web UI、API Server
核心层：Auth Service、Order Service、Payment Service
数据层：PostgreSQL、Redis、Object Storage
外部系统：Payment Gateway、Email Provider

要求：
- 使用三层布局。
- 服务之间用有向箭头表示调用关系。
- 外部系统放在右侧。
- 数据存储放在底部。
```

## 与 Codex / MCP 集成

### 推荐路径

`next-ai-draw-io` 的 MCP server 是自包含服务，会通过 stdio 接收 MCP 工具调用，并启动内嵌 HTTP server 在浏览器中实时预览图。

推荐先用 `npx` 方式接入，不做全局安装：

```bash
codex mcp add drawio -- npx @next-ai-drawio/mcp-server@latest
```

安全提示：`@next-ai-drawio/mcp-server` 旧版本曾存在资源限制相关漏洞。建议使用 `@latest`，或明确固定到 `0.1.19` 及以上版本。

Claude Code CLI 可用：

```bash
claude mcp add drawio -- npx @next-ai-drawio/mcp-server@latest
```

通用 MCP 客户端配置：

```json
{
  "mcpServers": {
    "drawio": {
      "command": "npx",
      "args": ["@next-ai-drawio/mcp-server@latest"]
    }
  }
}
```

配置后重启对应 MCP 客户端，再让 AI 创建图。官方 MCP README 说明图会在浏览器中实时出现。

### 是否需要 API Key

通过 Codex / Claude 的 MCP 使用 `next-ai-draw-io` 时，不需要给 `next-ai-draw-io` 单独配置模型 API Key：

```text
Codex / Claude 负责 LLM 推理
next-ai-draw-io MCP 负责浏览器预览、创建图、编辑图、读取 XML 和导出 .drawio
```

因此 MCP 配置中通常不需要 `OPENAI_API_KEY`、`AI_PROVIDER` 或 `AI_MODEL`。这些变量只在你直接运行 `next-ai-draw-io` Web 应用、Docker 应用，并使用它自己的 AI 聊天能力时才需要。

`DRAWIO_BASE_URL` 不是模型 Key。它只用于指定 draw.io embed 服务地址，例如改成私有 draw.io 部署。

### MCP 工具能力

官方 MCP server 提供的主要工具：

| 工具 | 用途 |
| --- | --- |
| `start_session` | 打开浏览器实时预览会话。 |
| `create_new_diagram` | 根据 draw.io XML 创建新图。 |
| `edit_diagram` | 按 ID 操作编辑已有图。 |
| `get_diagram` | 获取当前图的 XML。 |
| `export_diagram` | 导出 `.drawio` 文件。 |

AI 使用时应先启动 session，再创建或编辑图。遇到 `No active session` 时，先调用 `start_session`。

连接后的典型流程：

```text
1. AI 调用 start_session。
2. 浏览器打开实时预览页，URL 通常带 ?mcp= 参数。
3. AI 调用 create_new_diagram 或 edit_diagram。
4. 用户在浏览器中看到图实时更新。
5. 需要保存时，AI 调用 get_diagram 或 export_diagram。
```

### 端口和私有 draw.io

默认内嵌 HTTP server 端口是 `6002`。如果端口占用，官方说明 server 会尝试下一个可用端口，直到 `6020`。也可以显式设置：

```json
{
  "mcpServers": {
    "drawio": {
      "command": "npx",
      "args": ["@next-ai-drawio/mcp-server@latest"],
      "env": {
        "PORT": "6003"
      }
    }
  }
}
```

默认 draw.io embed 地址是：

```text
https://embed.diagrams.net
```

如果安全敏感环境需要私有部署，可以配置：

```json
{
  "mcpServers": {
    "drawio": {
      "command": "npx",
      "args": ["@next-ai-drawio/mcp-server@latest"],
      "env": {
        "DRAWIO_BASE_URL": "https://drawio.example.com"
      }
    }
  }
}
```

也可以本地启动官方 draw.io Docker 镜像：

```bash
docker run -d -p 8080:8080 jgraph/drawio
```

然后把 `DRAWIO_BASE_URL` 设置为自己的 draw.io 地址。

## 与 CodeGraph / Understand Anything 搭配

推荐分工：

| 工具 | 职责 |
| --- | --- |
| Understand Anything | 获取项目层级、业务域、关键模块和整体知识图谱。 |
| CodeGraph | 查询符号定义、调用链、依赖关系、复杂度热点和 diff impact。 |
| 源码 / 测试 / 运行结果 | 校验关键事实，避免图画错。 |
| next-ai-draw-io | 把已确认的架构或流程渲染为 draw.io 图，并导出 `.drawio`。 |

推荐提示词：

```text
使用 my-codebase-intel 把这个项目整理成 draw.io 生成架构图。
```

```text
使用 my-codebase-intel 把这个项目整理成 draw.io 生成流程图。
```

skill 内部应先结合 Understand Anything、CodeGraph 和必要源码整理事实，再调用 `start_session` 打开 draw.io MCP 实时预览，最后使用 `next-ai-draw-io` MCP 出图。MCP 模式下不需要给 `next-ai-draw-io` 单独配置模型 API Key。

## 什么时候让 AI 使用

适合让 AI 使用：

- 将代码库架构摘要转成 draw.io 架构图。
- 将模块调用链转成流程图。
- 将 CodeGraph 的依赖关系转成模块依赖图。
- 将需求文档、排查流程、部署拓扑转成可编辑图。
- 在已有图基础上做增量修改，并导出 `.drawio` 文件。

不建议让 AI 使用：

- 直接根据猜测画项目图，不先验证源码或图谱。
- 把生成图当成架构事实来源。
- 在没有用户授权时安装 npm 包或修改 MCP 配置。
- 在文档中写入真实 API Key、个人 Token、私有代理地址或内部服务 URL。

## 常见问题

### 端口被占用

默认端口是 `6002`。如果占用，server 会尝试后续端口。也可以在 MCP 配置中设置 `PORT`。

### 出现 `No active session`

先让 AI 调用 `start_session`，打开浏览器预览会话，再创建或编辑图。

### 浏览器没有实时更新

检查浏览器地址是否带有 `?mcp=` 参数。这个参数用于把浏览器预览和 MCP 会话关联起来。

### 私有环境不能访问 `embed.diagrams.net`

部署自己的 draw.io 实例，并设置 `DRAWIO_BASE_URL`。如果使用内网地址，不要把真实内部 URL 写入仓库文档。

### 图和代码事实不一致

回到事实来源重新检查：源码、测试、运行结果、CodeGraph 查询和 Understand Anything 图谱。`next-ai-draw-io` 只负责画图，不负责判断代码结构是否真实。
