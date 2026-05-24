# 补充 ReAct Execution 工作流与 Skill 推荐

## 任务背景

用户认可将 ReAct Execution 作为执行阶段的小步反馈循环，并希望把它补充到工作流文档中，同时把 `react-execution` 作为不同任务场景下的推荐 skill。用户还明确要求在 `AGENTS.md` 中加入已确认的 ReAct 执行约束，保持长期硬规则简短清晰。

## 根因定位

仓库已有 `ai-workflow.md` 说明 vibe、Plan、Spec、TDD 和 Harness Engineering，也有 `codex-skills.md` 说明现有 skills 的安装与触发方式。但此前缺少 ReAct Execution 的定位：

- `ai-workflow.md` 没有说明 ReAct 与 Spec、Plan、Harness 的层级关系。
- `codex-skills.md` 没有把 `react-execution` 放进分场景推荐。
- `AGENTS.md` 没有长期硬规则要求复杂执行任务按 ReAct 小步循环推进。

## 执行计划

1. 在 `AGENTS.md` 新增 `ReAct 执行约束` 小节，只写已确认的 7 条硬规则。
2. 在 `ai-workflow.md` 新增 `ReAct Execution 的位置` 小节，说明它是执行阶段机制，不替代 Spec 或 Plan。
3. 更新 `ai-workflow.md` 的任务分级模型和默认流程，加入 `react-execution` 与 Harness 验证的组合。
4. 更新 `codex-skills.md`，把 `react-execution` 写入当前可用 skills 表和分场景推荐。
5. 执行关键词检查、安全占位检查和 Markdown diff 空白检查。

## 变更内容

- `AGENTS.md`：
  - 新增 `ReAct 执行约束`。
  - 保留用户确认的 7 条规则，包括每轮只做一个明确动作、修改前基于观察说明判断、修改后验证、连续 3 轮不收敛时停止汇报。
- `ai-workflow.md`：
  - 新增 `ReAct Execution 的位置`。
  - 保留英文工作流 `Goal -> Observe -> Decide -> Act -> Verify -> Adjust -> Final`。
  - 更新任务分级模型，增加推荐 skill 列。
  - 更新推荐默认流程和 Spec 后执行路径。
- `codex-skills.md`：
  - 当前可用 skills 表新增 `react-execution`。
  - 新增 `分场景推荐`，说明不同任务应显式点名的 skill。

## 验证结果

- 已执行 `rg -n "ReAct 执行约束|复杂任务、bugfix、接口兼容|Goal -> Observe -> Decide -> Act -> Verify -> Adjust -> Final|react-execution" AGENTS.md ai-workflow.md codex-skills.md`：确认 ReAct 约束、工作流和 skill 推荐已写入。
- 已执行 `rg -n "brainstorming|planning-with-files|writing-plans|systematic-debugging|react-execution|verification-before-completion" codex-skills.md`：确认分场景推荐覆盖预期 skills。
- 已执行 `rg -n "TODO|FIXME|sk-|真实 API Key" AGENTS.md ai-workflow.md codex-skills.md docs/changes/2026-05-25-react-execution-workflow.md`：仅命中既有安全提示和检查命令示例，未发现真实密钥或遗留占位符。
- 已执行 `git diff --check`：检查 Markdown diff 空白。

## 后续建议

后续可单独创建 `~/.codex/skills/react-execution/SKILL.md`，把完整执行循环沉淀为实际可触发的 Codex skill；当前变更先完成仓库工作流和推荐项说明。
