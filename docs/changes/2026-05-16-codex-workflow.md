# 2026-05-16 Codex 工作流调整

## 任务背景

用户希望把 Codex 协作流程调整为：每次执行前先规划，由 AI 检查计划，执行后把本次会话改动写入日期开头的文档，方便后续 AI 和开发者理解历史上下文。

## 根因定位

当前 `codex-core-commands.md` 的推荐工作流允许小改动跳过 `/plan`，并说明计划默认不会自动写入仓库文件。该规则无法保证每次执行前都有计划、自检和长期上下文沉淀。

## 执行计划

1. 阅读仓库结构、`README.md` 和 `codex-core-commands.md`，确认核心指令位置。
2. 将核心工作流升级为“规划 -> AI 检查计划 -> 执行 -> 生成变更文档 -> diff -> review”。
3. 补充 `docs/changes/YYYY-MM-DD-<topic>.md` 的命名和内容规范。
4. 同步更新 `README.md` 的目录结构、入口和使用建议。
5. 运行本地检查，确认没有占位符、疑似密钥或无关改动。

AI 自检结论：计划已覆盖根因定位、修改文件、验证方式和风险边界；本次只修改核心指令、入口文档和新增会话变更记录，不涉及配置模板和工具链。

## 变更内容

- `codex-core-commands.md`：更新推荐工作流，新增会话变更文档规范和 AI 检查计划规则，并调整 `/plan` 为默认执行前步骤。
- `README.md`：补充 `docs/changes/` 目录说明、快速入口和使用建议。
- `AGENTS.md`：新增本仓库 AI 协作工作流硬性规则，要求产生文件改动时自动生成或更新日期变更文档。
- `docs/changes/2026-05-16-codex-workflow.md`：记录本次会话背景、根因、计划、变更和验证。

追加更新：

- 将“生成日期变更文档”明确为本仓库项目级自动执行规则：只要会话产生仓库文件改动，AI 必须在结束前自动创建或更新 `docs/changes/YYYY-MM-DD-<topic>.md`。
- 在 `codex-core-commands.md` 中补充新项目的全局启用方式：可写入 Codex 全局指令或 `/init` 使用的 `AGENTS.md` 模板，让新项目默认继承。

再次追加：

- 在 `codex-core-commands.md` 中新增 `/collab` 说明，用于记录从 Plan Mode 切回默认协作模式的入口和注意事项。
- 安装 `writing-plans`、`systematic-debugging`、`verification-before-completion` 三个 Superpowers skills；`brainstorming` 此前已安装。
- 在 `codex-skills.md` 中记录程序员推荐安装组合，以及切换为完整 Superpowers 插件的安装方式和取舍。
- 调整 `.gitignore`，继续忽略其他 `docs` 内容，但允许 `docs/changes/` 下的日期变更文档被 Git 识别。

## 验证结果

- `git status --short`：显示 `README.md` 已修改，`AGENTS.md`、`codex-core-commands.md` 和 `docs/` 为未跟踪文件；本次未回滚已有未提交内容。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中记录的检查命令本身，未发现真实密钥或遗留占位符。

追加验证：

- `git status --short`：显示 `README.md` 已修改，`AGENTS.md`、`codex-core-commands.md` 和 `docs/` 为未跟踪文件，符合本次文档变更范围。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中 `AGENTS.md` 和本变更记录中的检查命令示例，未发现真实密钥或遗留占位符。

再次验证：

- `find ~/.codex/skills -maxdepth 2 -name SKILL.md`：确认 `brainstorming`、`writing-plans`、`systematic-debugging`、`verification-before-completion` 已安装。
- `git status --short --untracked-files=all`：确认 `docs/changes/2026-05-16-codex-workflow.md` 已可作为未跟踪文件显示，不再被 `docs` 忽略规则屏蔽。

## 后续建议

后续每次会话产生仓库改动时，都按 `docs/changes/YYYY-MM-DD-<topic>.md` 新增记录，并在记录中保留最终采用的计划摘要和验证结果。
