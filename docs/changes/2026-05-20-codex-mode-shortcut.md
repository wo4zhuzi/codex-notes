# 2026-05-20 更新 Codex 模式切换说明

## 任务背景

用户希望删除当前文档中 `/collab` 命令介绍，并改为记录 Codex 切换模式快捷键是 `Shift + Tab`。后续又补充：即使配置中推理强度是 `high`，切换到 `/plan` / Plan Mode 后也可能显示为 `medium`，需要在 Plan Mode 内通过 `/model` 永久覆盖该模式下的推理设置。

## 根因定位

`codex-core-commands.md` 仍把 `/collab` 写成协作模式切换入口，并在推荐工作流中把 `/collab` 与 `/plan` 并列。当前实际使用习惯是通过 `Shift + Tab` 切换模式，继续保留 `/collab` 命令介绍会误导后续执行。

同时，原模式切换章节只说明如何进入或退出 Plan Mode，没有说明 Plan Mode 内的模型和推理强度可能独立于全局配置。这样会导致用户误以为配置文件中的 `high` 会自动继承到 Plan Mode，但实际进入 Plan Mode 后仍可能需要用 `/model` 单独永久覆盖。

## 执行计划

1. 检查 `git status --short`，确认没有用户未提交改动。
2. 搜索 `/collab` 和模式切换相关说明，定位当前文档入口。
3. 最小修改 `codex-core-commands.md`：删除 `/collab` 命令入口，改为 `Shift + Tab` 快捷键说明。
4. 在模式切换章节补充 Plan Mode 下通过 `/model` 永久覆盖模型和推理强度的规则。
5. 新增或更新本次日期变更文档，并执行仓库约定检查。

## 变更内容

- `codex-core-commands.md`：推荐工作流移除 `/collab`；原 `/collab：切换协作模式` 章节改为 `Shift + Tab：切换协作模式`，示例和使用规则同步改为快捷键表述。
- `codex-core-commands.md`：在模式切换使用规则中补充 Plan Mode 的模型和推理强度可能独立于全局配置；如全局为 `high` 但 Plan Mode 显示 `medium`，需进入 Plan Mode 后通过 `/model` 重新选择并永久覆盖该模式设置。
- `docs/changes/2026-05-20-codex-mode-shortcut.md`：记录本次任务背景、根因定位、执行计划、变更内容和验证结果。

## 验证结果

- `rg -n "/collab" codex-core-commands.md README.md AGENTS.md ai-workflow.md ai-project-checklist.md spec-workflow.md tdd-workflow.md codex-skills.md git-workflow.md`：无命中，当前说明文档不再介绍 `/collab` 命令入口。
- `rg -n "Shift \+ Tab|shift \+ tab" codex-core-commands.md docs/changes/2026-05-20-codex-mode-shortcut.md`：命中 `codex-core-commands.md` 的模式切换章节和本变更记录。
- `rg -n "/model|推理|medium|high|Shift \+ Tab|/plan" codex-core-commands.md docs/changes/2026-05-20-codex-mode-shortcut.md`：命中模式切换快捷键、Plan Mode 推理强度覆盖说明和本变更记录。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- `git status --short`：显示本次修改 `codex-core-commands.md`，并新增 `docs/changes/2026-05-20-codex-mode-shortcut.md`。

## 后续建议

历史变更文档中保留过去曾新增 `/collab` 说明的记录，不回写历史事实；后续如 Codex 客户端快捷键变化，应优先更新 `codex-core-commands.md` 的当前说明。
