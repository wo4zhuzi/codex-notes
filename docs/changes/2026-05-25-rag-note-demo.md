# 2026-05-25 RAG 笔记与 Demo

## 任务背景

用户已经理解 embedding 是底层能力，但对“RAG 在项目中怎么应用”仍然模糊，希望新增一篇 RAG 笔记，并配套一个可运行 demo，通过调用模型理解完整链路。

## 根因定位

现有仓库已有 Function Calling、MCP、CodeGraph 等笔记和 demo，但缺少 RAG 的独立说明与最小闭环示例。仅解释概念无法直观看到 `检索 -> 上下文注入 -> 模型回答` 这条链路。

## 执行计划

- 新增 `rag.md`，解释 RAG 的项目位置、与 embedding / Function Calling / MCP / Agent 的关系，以及从最小 RAG 到 Hybrid RAG 的升级路线。
- 新增 `demos/rag-notes/`，提供命令行和网页两种体验。
- 使用本地 Markdown 小知识库和关键词检索实现最小 RAG，不引入 embedding、向量库和 reranker。
- 更新 `README.md` 的内容说明、目录结构、快速入口和使用建议。
- 运行 Go 测试和仓库级敏感信息检查。

## 变更内容

- 新增 `rag.md`。
- 新增 `demos/rag-notes/`：
  - `knowledge/*.md`：本地知识库。
  - `retriever.go`：关键词检索与评分。
  - `assistant.go`：RAG prompt 组装和 OpenAI Responses API 调用。
  - `main.go`：命令行与 Web 启动入口。
  - `server.go`：HTTP 路由和 `/api/chat`。
  - `templates/index.html`：网页聊天界面。
  - `*_test.go`：检索、Assistant 和 HTTP 层测试。
  - `README.md`：运行说明。
- 更新 `README.md`，增加 RAG 笔记和 demo 入口。

## 验证结果

```bash
cd demos/rag-notes && go test ./...
```

结果：通过，输出 `ok rag-notes`。

```bash
git status --short
```

结果：确认仅包含本次 RAG 笔记、demo、README 和变更文档相关改动。

```bash
rg -n "TODO|FIXME|your-api-key|sk-" .
```

结果：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。

## 后续建议

- 如果后续文档规模增加，可在当前 demo 基础上增加 embedding 和向量检索。
- 如果召回准确率不足，可升级为 BM25 / 全文搜索 + embedding + RRF 融合。
- 如果用于团队或企业知识库，需要补充权限过滤、版本过滤、引用追踪和评估集。
