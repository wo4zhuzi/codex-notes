# 新增 Next AI Draw.io 外部工具文档

## 任务背景

用户希望在外部工具目录中整理 `DayuanJiang/next-ai-draw-io` 的安装和使用方式，方便后续将它与 CodeGraph、Understand Anything、`my-codebase-intel` 搭配，用于生成 draw.io 架构图和流程图。

## 根因定位

仓库当前 `external-tools/` 只记录了 CodeGraph，缺少 AI 出图工具的安装、MCP 接入、提示词模板和使用边界说明。`next-ai-draw-io` 的定位不同于 CodeGraph：它负责 draw.io 图生成、编辑、预览和导出，不负责代码事实理解，因此需要单独文档说明职责边界。

## 执行计划

1. 查阅官方 README、MCP Server README 和安全索引，确认安装命令、MCP 配置、工具能力、端口、私有部署配置和版本注意事项。
2. 新增 `external-tools/next-ai-draw-io.md`，整理安装、基础使用、MCP 接入、与 CodeGraph / Understand Anything 配合方式和常见问题。
3. 更新 `external-tools/README.md`，加入工具索引。
4. 更新根目录 `README.md`，补充目录结构和使用建议。
5. 执行静态检查和敏感信息扫描。

## 变更内容

- 新增 `external-tools/next-ai-draw-io.md`：
  - 记录在线 demo、桌面应用、本地源码运行、Docker 运行和 MCP 接入方式。
  - 记录 `codex mcp add drawio -- npx @next-ai-drawio/mcp-server@latest` 和通用 MCP JSON 配置。
  - 说明 `PORT`、`DRAWIO_BASE_URL`、`start_session`、`create_new_diagram`、`edit_diagram`、`get_diagram`、`export_diagram` 的用途。
  - 增加 MCP server 旧版本安全提示，建议使用 `@latest` 或 `0.1.19` 及以上版本。
  - 增加与 CodeGraph、Understand Anything、`my-codebase-intel` 的搭配方式和提示词。
- 更新 `external-tools/README.md`：
  - 在工具表中加入 Next AI Draw.io。
- 更新 `README.md`：
  - 在内容说明、目录结构和使用建议中加入 Next AI Draw.io。

## 验证结果

- 已确认 `external-tools/next-ai-draw-io.md` 文件存在。
- 已确认 README、外部工具目录和变更记录中包含 `next-ai-draw-io` / `drawio` 相关入口。
- 已执行敏感信息扫描，未发现真实密钥形态或遗留标记。
- 已执行 `command -v markdownlint`：本机未安装 `markdownlint`，因此未执行 Markdown lint。
- 未执行 `npx @next-ai-drawio/mcp-server@latest`、`codex mcp add ...`、`claude mcp add ...` 或 Docker 启动命令；本次只整理文档。

## 后续建议

- 如需真正接入 Codex，先由用户确认是否允许修改本机 MCP 配置，再执行 `codex mcp add drawio -- npx @next-ai-drawio/mcp-server@latest`。
- 如果在安全敏感环境使用，优先配置私有 draw.io 实例，并在仓库文档中只写占位域名。
