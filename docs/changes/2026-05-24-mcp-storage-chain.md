# 补充 MCP 存储与调用链路

## 任务背景

用户已经安装 CodeGraph MCP，希望在文档中补充安装后的存储路径、存储格式，以及 Codex 使用 MCP 查询 CodeGraph 的原理和链路。用户进一步明确不要写个人本机绝对路径，应说明 MCP 会注册到 `~/.codex/config.toml` 以及 TOML 配置格式。随后用户指出这类注册格式和调用链路属于通用 MCP 原理，应主要体现在 `mcp.md`，而不是 CodeGraph 专题文档。

## 根因定位

`mcp.md` 已记录 MCP 基本角色、Codex MCP 命令、stdio / Streamable HTTP、认证和排查，但缺少以下通用落地信息：

- `codex mcp add` 写入 `~/.codex/config.toml` 后大致是什么 TOML 结构。
- Codex 作为 MCP client 到 MCP server 再到外部工具、索引、文档、数据库或远程服务的完整调用链路。

`external-tools/codegraph.md` 则应只保留 CodeGraph 专属信息，例如 `.codegraph/graph.db` 的存储格式、索引内容和 `codegraph mcp -d` 参数。

## 执行计划

1. 在 `mcp.md` 中补充“配置落盘位置与格式”，说明 `~/.codex/config.toml` 和 `[mcp_servers.<name>]`。
2. 在 `mcp.md` 中补充通用“MCP 调用链路”，说明 stdio、Streamable HTTP、JSON-RPC 和 MCP server 后端能力的关系。
3. 收窄 `external-tools/codegraph.md`，移除通用 MCP 注册和链路展开，只保留 CodeGraph 专属存储格式与 `codegraph mcp -d` 示例。
4. 执行 Git 状态、关键词和敏感占位符检查，确认文档改动范围符合预期。

## 变更内容

- 在 `mcp.md` 中补充 `~/.codex/config.toml` 的 MCP 配置落盘位置。
- 在 `mcp.md` 中补充 `[mcp_servers.local-docs]` 和 `[mcp_servers.codegraph]` 的 stdio TOML 示例。
- 在 `mcp.md` 中补充 Codex 读取配置、启动或连接 MCP server、通过 JSON-RPC 通信、由 MCP server 访问后端能力并返回结构化结果的通用链路。
- 在 `external-tools/codegraph.md` 中保留 CodeGraph CLI、npm 全局包和 `.codegraph/graph.db` 三类 CodeGraph 相关位置说明。
- 在 `external-tools/codegraph.md` 中说明 `graph.db` 是 SQLite 3 数据库，并补充 `graph.db-wal`、`graph.db-shm`、`changes.journal` 等同目录文件。
- 在 `external-tools/codegraph.md` 中保留显式 `-d` 指定数据库示例，并链接回 `mcp.md` 查看通用 MCP 配置和链路说明。

## 验证结果

- 已执行 `git status --short`：执行前工作区干净，执行后包含 `mcp.md`、`external-tools/codegraph.md` 和本变更记录文件。
- 已执行 `rg -n "config.toml|mcp_servers|graph.db|SQLite|stdio|JSON-RPC|调用链路" mcp.md external-tools/codegraph.md`：确认核心说明已写入。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。

## 后续建议

如果后续 Codex 调整 MCP 配置字段，应以 `codex mcp get <name> --json` 和 `~/.codex/config.toml` 实际内容为准同步更新 `mcp.md`；如果 CodeGraph 新版本调整 `codegraph mcp` 参数，再同步更新 `external-tools/codegraph.md`。
