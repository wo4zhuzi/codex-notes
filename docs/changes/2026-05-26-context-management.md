# 2026-05-26 Agent 上下文管理

## 任务背景

用户询问上下文工程是否已被 Harness Engineering 取代，并希望将相关内容融入当前 AI Agent 学习笔记项目。

## 根因定位

当前仓库已经在 `ai-workflow.md`、`codex-core-commands.md` 和 README 中分散记录了 `/mention`、`/compact`、`resume`、`docs/changes/` 等上下文相关机制，但缺少一篇面向学习路线的集中说明，无法清晰解释上下文管理与 Harness Engineering 的分工。

## 执行计划

- 新增 `context-management.md`，将上下文工程落地为“Agent 上下文管理”。
- 更新 `README.md`，加入新文档入口并调整学习顺序。
- 在 `ai-workflow.md` 中补充上下文管理与 Harness Engineering 的边界说明。
- 在 `codex-core-commands.md` 中补充 `/mention` 与上下文收敛的关系。
- 运行仓库约定的本地检查命令。

## 变更内容

- 新增 `context-management.md`：
  - 说明上下文管理目标。
  - 梳理上下文选择、结构、生命周期、污染和常见场景。
  - 明确上下文管理与 Harness Engineering 的分工。
- 更新 `README.md`：
  - 内容说明加入 Agent 上下文管理。
  - 目录结构和快速入口加入 `context-management.md`。
  - 使用建议中将上下文管理放在 AI 编程工作流之后。
- 更新 `ai-workflow.md`：
  - 在 Harness Engineering 原则下补充边界说明和交叉引用。
- 更新 `codex-core-commands.md`：
  - 在 `/mention` 章节补充“上下文收敛”说明和交叉引用。

## 验证结果

- 已执行 `git status --short`：确认本次改动包含 `README.md`、`ai-workflow.md`、`codex-core-commands.md`、`context-management.md` 和本变更记录。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已手动检查 README 新增链接：`./context-management.md` 指向真实文件。

## 后续建议

- 后续如果继续扩展 AI Agent 学习路线，可优先补充 Agent 评估、安全与权限治理、生产化可观测性等主题。
