# 新增 CodeGraph 外部工具文档

## 任务背景

用户希望后续持续记录 CodeGraph、Serena、Context7、ast-grep 等类似外部工具的安装、配置和 MCP 接入方式，因此需要先明确仓库目录规范，并新增 CodeGraph 的完整落地方案。用户明确要求只写文档，不由 AI 代为安装。

## 根因定位

仓库此前只有根目录主题文档和 `cc-switch-configs/` 配置模板目录，没有专门存放外部工具使用笔记的位置。如果继续把所有工具文档放在根目录，后续工具增多后会降低 README 和根目录可读性。

CodeGraph 这类工具既有 CLI 安装、索引文件、gitignore、MCP 接入，又有 AI 使用边界，适合放入独立的外部工具目录，并在 `AGENTS.md` 中固化后续放置规则。

## 执行计划

1. 查看 `git status --short`，确认工作区没有用户已有未提交改动。
2. 阅读 `AGENTS.md`、`README.md`、`.gitignore` 和 `mcp.md`，确认现有结构和写法。
3. 查阅 `@optave/codegraph` 公开说明，核验安装、建索引、索引文件和 MCP server 命令。
4. 新增 `external-tools/README.md`，作为外部工具索引。
5. 新增 `external-tools/codegraph.md`，写 CodeGraph 安装、建索引、gitignore、Codex MCP 接入、日常流程和排查清单。
6. 更新 `AGENTS.md`，补充 `external-tools/` 目录职责和新增文档放置规则。
7. 更新 `README.md`，补充外部工具目录结构、快速入口和阅读建议。
8. 更新 `.gitignore`，忽略 CodeGraph 本地索引目录。
9. 更新 `mcp.md`，补充 CodeGraph 类工具的“两段式”接入模式。
10. 新增本变更文档，并执行本地检查。

AI 自检结论：计划包含根因定位、修改步骤、验证方式和风险边界；本次只修改文档和 `.gitignore`，不安装 npm 包、不执行 `codegraph build`、不启动 MCP server、不提交 commit。

## 变更内容

- 更新 `AGENTS.md`：
  - 新增 `external-tools/` 目录职责。
  - 明确通用主题文档、外部工具文档、变更记录和配置模板的放置规则。
- 更新 `README.md`：
  - 增加外部工具文档说明。
  - 增加 `external-tools/` 目录结构和快速入口。
  - 在使用建议中补充外部工具阅读顺序。
- 更新 `.gitignore`：
  - 忽略 `.codegraph/` 和 `.code-graph/`。
- 更新 `mcp.md`：
  - 新增 CodeGraph 类工具接入模式，说明“CLI / watch / CI 构建索引，Codex 通过 MCP 查询索引”。
- 新增 `external-tools/README.md`：
  - 定义外部工具目录定位、当前工具索引和新增工具文档模板。
- 新增 `external-tools/codegraph.md`：
  - 记录 `@optave/codegraph` 安装、建索引、MCP 接入、AI 使用边界、日常流程和排查清单。

## 验证结果

- 已执行 `git status --short`：显示本次修改 `.gitignore`、`AGENTS.md`、`README.md`、`mcp.md`，新增 `external-tools/` 和本变更文档；未发现其他无关改动。
- 已执行 `test -f external-tools/README.md && test -f external-tools/codegraph.md`：命令退出码为 0，确认新目录文件存在。
- 已执行 `rg -n "external-tools|CodeGraph|codegraph" README.md AGENTS.md mcp.md external-tools`：确认 README、AGENTS、MCP 文档和外部工具目录均包含预期入口、链接和 CodeGraph 内容。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `command -v markdownlint`：本机未安装 `markdownlint`，因此未执行 Markdown lint。
- 未执行 `npm install -g @optave/codegraph`、`codegraph build` 或 `codex mcp add codegraph ...`，符合用户要求。

## 后续建议

- 后续新增 Serena、Context7、ast-grep 等工具时，直接在 `external-tools/` 下新增独立 Markdown 文件，并更新 `external-tools/README.md`。
- 如果某个工具需要真实 Token、私有服务地址或团队内部配置，只写占位符和配置位置，不写真实值。
