# 2026-05-26 Agent 评估与验证

## 任务背景

用户希望将 Agent Evaluation 融入现有 AI 协作工作流，并明确不同阶段、项目风险和改动类型下如何确定验收边界。

## 根因定位

当前仓库已有 AI 工作流、上下文管理、Spec、TDD、Git 和核心命令说明，但缺少一个集中说明“任务完成前如何判断是否可交付”的验收闸门。现有验证更多分散在 Harness Engineering、TDD 和 Git 流程中，未明确区分开发阶段的阶段性评估和最终收尾前的交付评估。

## 执行计划

- 新增 `agent-evaluation.md`，定义阶段性评估、交付评估和按场景选择评估强度的方法。
- 更新 `README.md`，加入 Agent 评估与验证入口和学习顺序。
- 更新 `ai-workflow.md`，将 Agent Evaluation 融入推荐默认流程和 Spec 执行路径。
- 更新 `codex-core-commands.md`，在 `/diff` 和 `/review` 前补充 Agent Evaluation 验收小节和可复制 prompt。
- 运行仓库约定的本地检查命令。

## 变更内容

- 新增 `agent-evaluation.md`：
  - 定义 Agent Evaluation 的目标和核心原则。
  - 区分阶段性评估与交付评估。
  - 按阶段、项目风险和改动类型定义评估强度。
  - 提供文档、bugfix、重构、配置、RAG/MCP/Function Calling、Code Review 的验收清单。
  - 提供阶段性评估和交付评估 prompt。
- 更新 `README.md`：
  - 内容说明、目录结构、快速入口和使用建议加入 Agent 评估与验证。
- 更新 `ai-workflow.md`：
  - 在机械验证原则中加入 Agent Evaluation 的交付判断定位。
  - 推荐默认流程加入阶段性评估和交付评估。
  - Spec + Plan、Spec + TDD 路径加入评估步骤。
- 更新 `codex-core-commands.md`：
  - 提交前顺序加入 Agent Evaluation 验收。
  - 新增 Agent Evaluation 任务验收小节。

## 验证结果

- 已执行 `git status --short`：确认本次改动包含 `README.md`、`ai-workflow.md`、`codex-core-commands.md`、`agent-evaluation.md` 和本变更记录。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `test -f agent-evaluation.md`：确认 README 新增链接目标文件存在。
- 已执行 `rg -n "agent-evaluation|Agent Evaluation|Agent 评估|阶段性评估|交付评估" README.md ai-workflow.md codex-core-commands.md`：确认关键入口和流程引用已写入。

## 后续建议

- 后续如需进一步提高执行稳定性，可把 Agent Evaluation 的最终验收 prompt 纳入常用 `AGENTS.md` 模板。
