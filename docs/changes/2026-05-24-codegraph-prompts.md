# 补充 CodeGraph 提示词模板

## 任务背景

用户希望把 CodeGraph MCP 的实际使用提示词记录到 CodeGraph 使用文档中，尤其是进入项目后需要改代码时，如何让 Codex 先用 CodeGraph 分析项目结构、调用链和影响面，从而更清晰、更省 token。

## 根因定位

`external-tools/codegraph.md` 已记录 CodeGraph 的安装、建索引、MCP 接入、常用命令和日常流程，但提示词部分只有一句“先用 codegraph 查一下这个函数的调用链和影响面，再给我修改方案”。这不足以覆盖项目理解、改代码、bug 排查、重构、代码审查和省 token 这几类高频场景。

## 执行计划

1. 在 `external-tools/codegraph.md` 的“什么时候让 AI 使用”之后新增“提示词模板”章节。
2. 按项目理解、改代码前、bug 排查、重构前、代码审查和省 token 用法整理可复制提示词。
3. 更新“推荐日常流程”中的 Codex 示例提示词，让它强调先查定义、调用方、被调用方和影响面。
4. 执行 Git 状态、关键词、敏感占位符和 Markdown 空白检查。

## 变更内容

- 新增“提示词模板”章节，说明进入项目根目录并启动 Codex 后，可以要求 Codex 先通过 CodeGraph MCP 查询结构，再读取少量必要源码。
- 补充项目理解提示词，用于分析核心模块、入口点、热点函数和主要依赖关系。
- 补充改代码前提示词，用于先查目标符号定义、调用方、被调用方和影响面，再给最小修改方案。
- 补充 bug 排查、重构前、代码审查和省 token 用法提示词。
- 在日常流程中替换原有简短示例，并引导读者查看“提示词模板”。
- 补充边界说明：CodeGraph 更适合有函数、类、模块调用关系的代码项目，纯文档或配置仓库收益可能有限。

## 验证结果

- 已执行 `git status --short`：执行前工作区干净，执行后仅包含 `external-tools/codegraph.md` 和本变更记录文件。
- 已执行 `rg -n "提示词模板|项目理解|改代码前|bug 排查|staged diff|不要直接大范围读文件" external-tools/codegraph.md`：确认核心模板已写入。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `git diff --check -- external-tools/codegraph.md docs/changes/2026-05-24-codegraph-prompts.md`：未发现 Markdown 空白错误。

## 后续建议

后续如果增加 Serena、Context7、ast-grep 等外部工具文档，也可以按相同方式补充“提示词模板”，把工具适用场景和提示词分开记录。
