# Codex 协作工作流

Codex 适合参与代码阅读、方案设计、实现、测试和代码审查。

文档型仓库执行会产生文件改动的任务时，推荐流程：

```text
/init -> /status -> /plan -> AI 检查计划 -> 执行任务 -> 生成日期变更文档 -> /diff -> /review
```

关键约束：

- 修改前先看 `git status --short`。
- 修改前先说明计划和风险边界。
- 不回滚用户已有未提交改动。
- 修改后运行对应验证命令。
- 如果会产生仓库文件改动，结束前创建 `docs/changes/YYYY-MM-DD-<topic>.md`。

这些规则适合和 RAG、Function Calling、MCP 等学习 demo 一起沉淀到笔记仓库中。
