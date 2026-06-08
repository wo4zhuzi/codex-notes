# Codex Agent 运行机制文档

## 任务背景

用户整理了 Codex 任务上下文发现逻辑，希望进一步完善。讨论后确认该主题不只是上下文管理，还涉及 Codex Agent runtime 的运行机制，因此应从 `context-management.md` 中单独拆出一篇文章。

## 根因定位

原整理同时覆盖了 `AGENTS.md` 加载、skills 触发、memories、工具探索、文件读取范围、任务恢复和提示词写法。若直接放入上下文管理文档，会混淆两类问题：

- 上下文管理：人类如何准备、收敛和沉淀上下文。
- Agent 运行机制：Codex 如何加载规则、发现上下文、调用工具和继续执行。

因此需要新增独立文档承载运行机制说明，并在现有入口文档中建立引用。

## 执行计划

1. 新增 `codex-agent-runtime.md`，系统说明 Codex 的启动上下文、文件发现、工具探索和任务状态沉淀。
2. 更新 `README.md`，把新文章加入内容说明、目录结构、快速入口和阅读顺序。
3. 更新 `context-management.md`，说明它与新文章的边界。
4. 更新 `llm-agent-principles.md`，把新文章作为 Codex 具体运行机制的延伸阅读。
5. 运行仓库检查命令，确认没有明显占位符、密钥或非预期改动。

## 变更内容

- 新增 `codex-agent-runtime.md`：
  - 说明 Codex 不会默认读取完整仓库。
  - 区分自动加载、按需读取和工具探索。
  - 解释“扫描多”和“扫描少”的提示词写法。
  - 补充普通规划文档为什么有时会被发现。
  - 提供 `AGENTS.md` 任务恢复入口模板和可复制 prompt。
- 更新 `README.md`：
  - 增加 `Codex Agent 运行机制` 条目。
  - 更新目录结构和快速入口。
  - 调整推荐阅读顺序。
- 更新 `context-management.md` 和 `llm-agent-principles.md`：
  - 增加到新文章的交叉引用，明确文档边界。

## 验证结果

已执行：

```bash
git status --short
rg -n "TODO|FIXME|your-api-key|sk-" .
```

结果：

- `git status --short` 显示本次仅涉及 `README.md`、`context-management.md`、`llm-agent-principles.md`、新增 `codex-agent-runtime.md` 和本变更记录。
- 敏感信息扫描仅命中文档中的检查命令示例、既有示例分支名 `project-task-branch` 和 `risk-reviewer` 名称，未发现真实 API Key、Token、私有代理地址或遗留调试标记。

## 后续建议

- 如后续有固定学习路线图或长期计划文件，可在 `AGENTS.md` 增加“任务恢复入口”，减少新会话依赖文件名猜测。
- 如果继续扩展该主题，可补一篇“如何写让 Agent 扫描多/扫描少的 prompt 模板库”。
