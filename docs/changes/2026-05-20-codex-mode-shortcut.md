# 2026-05-20 更新 Codex 模式切换说明

## 任务背景

用户希望删除当前文档中 `/collab` 命令介绍，并改为记录 Codex 切换模式快捷键是 `Shift + Tab`。

## 根因定位

`codex-core-commands.md` 仍把 `/collab` 写成协作模式切换入口，并在推荐工作流中把 `/collab` 与 `/plan` 并列。当前实际使用习惯是通过 `Shift + Tab` 切换模式，继续保留 `/collab` 命令介绍会误导后续执行。

## 执行计划

1. 检查 `git status --short`，确认没有用户未提交改动。
2. 搜索 `/collab` 和模式切换相关说明，定位当前文档入口。
3. 最小修改 `codex-core-commands.md`：删除 `/collab` 命令入口，改为 `Shift + Tab` 快捷键说明。
4. 新增本次日期变更文档，并执行仓库约定检查。

## 变更内容

- `codex-core-commands.md`：推荐工作流移除 `/collab`；原 `/collab：切换协作模式` 章节改为 `Shift + Tab：切换协作模式`，示例和使用规则同步改为快捷键表述。
- `docs/changes/2026-05-20-codex-mode-shortcut.md`：记录本次任务背景、根因定位、执行计划、变更内容和验证结果。

## 验证结果

- `rg -n "/collab" codex-core-commands.md README.md AGENTS.md ai-workflow.md ai-project-checklist.md spec-workflow.md tdd-workflow.md codex-skills.md git-workflow.md`：无命中，当前说明文档不再介绍 `/collab` 命令入口。
- `rg -n "Shift \+ Tab|shift \+ tab" codex-core-commands.md docs/changes/2026-05-20-codex-mode-shortcut.md`：命中 `codex-core-commands.md` 的模式切换章节和本变更记录。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- `git status --short`：显示本次修改 `codex-core-commands.md`，并新增 `docs/changes/2026-05-20-codex-mode-shortcut.md`。

## 后续建议

历史变更文档中保留过去曾新增 `/collab` 说明的记录，不回写历史事实；后续如 Codex 客户端快捷键变化，应优先更新 `codex-core-commands.md` 的当前说明。
