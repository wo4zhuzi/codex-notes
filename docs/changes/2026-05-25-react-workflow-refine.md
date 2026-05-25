# 细化 ReAct 与 Plan Execute 的关系

## 任务背景

用户反馈昨天补充的 ReAct 内容仍然偏抽象，并且在 `ai-workflow.md` 中像是突然插在三种协作模式后面，容易让人误以为 ReAct 是和 vibe、Plan、Spec 并列的第四种协作模式。

用户希望文档明确说明：ReAct 是 `Thought -> Action -> Observation -> Adjust` 的执行反馈循环；`Plan -> Execute` 是提前规划好再做；当 plan 链路较长时，可以在每个执行步骤中融合 ReAct。

后续用户继续补充：需要强调 ReAct 必须基于 Observation 反馈，在动态环境中容错率更高，但更依赖 LLM 的步进式推理能力；`Plan -> Execute` 更节省 token，并且对复杂任务有更好的宏观掌控力。

## 根因定位

`ai-workflow.md` 原有结构先介绍“三种协作模式”，再在任务分级后单独放置 `ReAct Execution 的位置`。虽然内容已经说明 ReAct 不替代 Spec 或 Plan，但层级表达仍不够清晰：

- 协作模式和执行机制没有明确分层。
- `Plan -> Execute` 与 ReAct 的区别没有展开。
- ReAct 对 Observation 的依赖，以及两种机制在动态容错、token 成本和宏观掌控力上的取舍没有写清楚。
- 长链路 plan 中如何嵌入 ReAct 没有说明。
- 推荐默认流程中 `ReAct Execution` 像一个独立阶段，而不是执行阶段内部的小步闭环。

## 执行计划

1. 重写 `ai-workflow.md` 中 `ReAct Execution 的位置` 小节。
2. 将小节标题调整为 `执行机制：Plan / Execute 与 ReAct`。
3. 补充 Vibe / Plan / Spec、ReAct、Harness 的层级关系。
4. 补充 `Thought -> Action -> Observation -> Adjust` 的含义。
5. 补充 `Plan -> Execute` 与 ReAct 的区别，以及长链路 plan 中嵌入 ReAct 的方式。
6. 补充 ReAct 与 `Plan -> Execute` 的工程取舍。
7. 调整推荐默认流程中的 ReAct 表述。
8. 执行本仓库约定的文档检查命令。

## 变更内容

- `ai-workflow.md`：
  - 将 `ReAct Execution 的位置` 改为 `执行机制：Plan / Execute 与 ReAct`。
  - 明确 ReAct 不是第四种协作模式，而是执行过程中的小步反馈循环。
  - 补充三层关系：Vibe / Plan / Spec 决定边界和授权，ReAct 决定执行反馈，Harness 提供外部验证反馈。
  - 补充 `Thought -> Action -> Observation -> Adjust` 的解释。
  - 强调 ReAct 的关键是每次 Action 后读取真实 Observation，再决定下一步，避免退化成“边想边猜”。
  - 增加 ReAct 与 `Plan -> Execute` 的取舍表：ReAct 动态容错更强但更依赖 LLM 步进式推理能力并消耗更多 token；`Plan -> Execute` 更节省 token，并且对复杂任务有更好的宏观掌控力。
  - 补充 `Plan -> Execute` 与 ReAct 的差异，以及在长链路 plan 中逐步骤嵌入 ReAct 的推荐做法。
  - 将默认流程中的 `ReAct Execution` 改为 `执行阶段用 ReAct 小步闭环`。

## 验证结果

- 已执行 `git status --short`：确认本次只修改 `ai-workflow.md`，并新增 `docs/changes/2026-05-25-react-workflow-refine.md`。
- 已执行 `rg -n "ReAct|Plan / Execute|Thought -> Action -> Observation|执行机制" ai-workflow.md`：确认新标题、核心循环、Plan / Execute 区分和默认流程表述均已写入。
- 已执行 `rg -n "Observation|动态环境|步进式推理|节省 token|宏观掌控力|边想边猜" ai-workflow.md docs/changes/2026-05-25-react-workflow-refine.md`：确认新增取舍说明已写入正文和变更记录。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `git diff --check`：未发现 Markdown diff 空白问题。

## 后续建议

如后续继续细化 `react-execution` skill，可把本次文档中的执行循环提炼为独立 `SKILL.md`，但本次不扩大范围。
