# 2026-05-20 补充 planning-with-files skill 使用规范

## 任务背景

用户希望整理 `planning-with-files` skill 的使用方式和触发方式，并明确 `task_plan.md`、`findings.md`、`progress.md` 与本仓库 `docs/changes/` 的关系，避免执行进度文件和长期变更记录职责重叠。

## 根因定位

`progress.md` 和 `docs/changes/` 都容易被理解为“进度记录”，但两者生命周期不同：`progress.md` 面向当前任务执行中的状态恢复，`docs/changes/YYYY-MM-DD-<topic>.md` 面向任务完成后的长期事实沉淀。冲突根因不是文件数量，而是缺少明确职责边界。

## 执行计划

1. 阅读 `codex-skills.md`、`AGENTS.md` 和现有工作流文档，确认已有规则。
2. 在 `codex-skills.md` 增加 `planning-with-files` 使用规范、触发方式和三文件分工。
3. 在 `AGENTS.md` 增加文件型计划流程硬性规则，明确 `progress.md` 不替代 `docs/changes/`。
4. 记录本次变更文档并运行仓库约定检查。

## 变更内容

- `codex-skills.md`：新增 `planning-with-files` 使用规范，包含适用场景、显式触发、自然触发、`task_plan.md` / `findings.md` / `progress.md` 分工、`progress.md` 示例，以及与 `docs/changes/` 的边界。
- `AGENTS.md`：新增文件型计划流程规则，明确 `task_plan.md` 记录执行前计划、`findings.md` 记录过程上下文、`progress.md` 仅记录当前任务执行状态，并要求这些文件不得写入敏感信息。
- 本地安装：通过 `skill-installer` 从 `OthmanAdi/planning-with-files` 的 `skills/planning-with-files` 路径安装到 `~/.codex/skills/planning-with-files`。
- `docs/changes/2026-05-20-planning-with-files-skill.md`：记录本次会话背景、根因、计划、变更和验证。

## 验证结果

- `git status --short --untracked-files=all`：显示本次修改 `AGENTS.md`、`codex-skills.md`，并新增 `docs/changes/2026-05-20-planning-with-files-skill.md`。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。

## 后续建议

后续使用 `planning-with-files` 时，统一按上游模板维护 `task_plan.md`、`findings.md` 和 `progress.md`。
