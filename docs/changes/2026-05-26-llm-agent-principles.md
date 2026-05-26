# 新增 LLM 与 Agent 原理地图

## 任务背景

用户指出当前文档更偏 AI 编程实践和工具使用，希望补齐原理层知识，例如 LLM 原理、Agent 多轮交互如何实现、工具调用循环如何工作。

## 根因定位

仓库已有 [AI 编程工作流](../../ai-workflow.md)、[Agent 上下文管理](../../context-management.md)、[Function Calling 使用笔记](../../function-calling.md)、[RAG 使用笔记](../../rag.md) 和 [Codex MCP 使用笔记](../../mcp.md)，但这些文档主要解释实践流程和工程链路，缺少一个统一入口说明：

- LLM 如何基于上下文生成回答。
- 多轮对话如何由客户端或 API 状态维护。
- Agent runtime 如何组织工具调用、观察结果和继续执行。
- 上下文、记忆、RAG、Function Calling、MCP 的边界。
- Agent 失败原因和可靠性验证的原理基础。

## 执行计划

1. 新增 `llm-agent-principles.md`，作为 LLM 与 Agent 原理层入口。
2. 更新 `README.md` 的内容说明、目录结构、快速入口和使用建议。
3. 保持已有实践文档职责不变，避免重复展开 Function Calling、RAG、MCP 的细节。
4. 运行文档任务相关验证命令，检查链接、敏感信息和工作区状态。

## 变更内容

- 新增 `llm-agent-principles.md`：
  - 说明 LLM 的 token、上下文窗口、消息角色和生成机制。
  - 解释推理、采样和常见模型参数。
  - 说明多轮对话的无状态和有状态实现方式。
  - 解释 Agent 工具调用循环和 ReAct 的 runtime 视角。
  - 对比 Function Calling、MCP、RAG、Shell / 本地工具的分工。
  - 区分当前上下文、历史消息、摘要压缩、长期记忆和 RAG。
  - 总结 Agent 常见失败来源和可靠性要求。
  - 补充 Codex 的状态模型：模型本身不永久记忆，Function Calling demo 采用无状态请求，Codex CLI 作为 Agent 客户端维护会话、工具结果、计划和恢复状态。
  - 说明 `codex resume`、`codex fork` 和 `/compact` 在多轮 Agent 状态管理中的位置。
  - 补充 API 层有状态/无状态判断：客户端本地有状态不等于模型服务端有状态，是否有状态取决于请求是否依赖 `previous_response_id`、`conversation_id` 或等价服务端状态。
  - 补充 Function Calling 应用建议：通用应用优先采用无状态请求，以提升兼容性、可复现性和调试便利性。
- 更新 `README.md`：
  - 在内容说明中加入 LLM 与 Agent 原理。
  - 在目录结构和快速入口中加入 `llm-agent-principles.md`。
  - 将原理地图放到使用建议的第一步。

## 验证结果

已执行：

```bash
git status --short
rg -n "TODO|FIXME|your-api-key|sk-" .
test -f llm-agent-principles.md
rg -n "llm-agent-principles|LLM 与 Agent 原理地图" README.md llm-agent-principles.md
rg -n "Codex.*有状态|无状态 API|有状态 API|resume|compact|previous_response_id" llm-agent-principles.md
rg -n "API 层|previous_response_id|conversation_id|登录账号|服务商|Function Calling 应用|无状态请求" llm-agent-principles.md
rg -n "previous_response_id|无状态请求" function-calling.md demos/function-calling-orders/README.md
codex resume --help
codex fork --help
```

结果：

- `git status --short` 显示本次改动包含 `README.md`、`llm-agent-principles.md` 和本变更记录。
- `test -f llm-agent-principles.md` 通过，确认新增文档存在。
- README 中已加入 `llm-agent-principles.md` 目录项、快速入口和使用建议。
- `codex resume --help` 显示 `resume` 可恢复历史交互会话，支持 `SESSION_ID` 和 `--last`。
- `codex fork --help` 显示 `fork` 可基于历史交互会话创建新路线，支持 `SESSION_ID` 和 `--last`。
- `llm-agent-principles.md` 已补充 Codex CLI 有状态、模型层不永久记忆、Function Calling demo 无状态请求三者的分层说明。
- `llm-agent-principles.md` 已补充 API 层判断规则，明确登录账号不等于 API 有状态，本机 session 不等于服务端 conversation。
- `function-calling.md` 和 `demos/function-calling-orders/README.md` 仍保持 demo 使用无状态请求的说明。
- 敏感信息扫描只命中文档中的检查命令示例和既有示例分支名 `project-task-branch`，未发现真实密钥、Token、私有代理地址或内部服务 URL。

## 后续建议

后续如果需要深入，可按主题拆分独立文档，例如 `llm-basics.md`、`agent-runtime.md`、`context-memory-rag.md`。当前阶段先保留为原理地图，避免过早拆成多篇长文档。
