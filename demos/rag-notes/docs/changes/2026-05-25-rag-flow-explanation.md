# RAG 流程说明补充

## 任务背景

用户希望在 README 中用更专业但可理解的语言说明 `rag-notes` 的 RAG 执行过程，并在 Go 实现代码中标注对应阶段，帮助理解“Go 检索”和“模型生成”的职责边界。

## 根因定位

原 README 只给出高层链路：

```text
用户问题 -> retriever.go 检索 -> assistant.go 拼 prompt -> 模型生成答案
```

这能说明流程顺序，但没有明确解释：

- 检索由 Go 程序完成，不是模型自己读取文档。
- 程序不会把整个 `knowledge/` 目录直接发送给模型。
- 模型收到的是“用户问题 + topK 检索结果”组成的增强 prompt。
- 模型职责是基于上下文生成自然语言答案，而不是执行检索。

## 执行计划

- 在 `README.md` 增加“RAG 执行过程详解”章节。
- 在 `retriever.go` 标注 Knowledge Loading、Retrieval、关键词拆分和评分规则。
- 在 `assistant.go` 标注先检索再生成、Prompt Augmentation 和模型输入边界。
- 保持现有运行逻辑、接口、评分规则和测试不变。

## 变更内容

- `README.md`
  - 新增 Go 程序与模型的职责分工说明。
  - 新增 Knowledge Loading、Retrieval、Prompt Augmentation、Generation 四阶段说明。
  - 明确当前 demo 不使用 embedding、向量数据库，也不会把全部知识库直接喂给模型。

- `retriever.go`
  - 为 `NewRetriever` 添加知识库加载阶段注释。
  - 为 `Search` 添加检索阶段注释。
  - 为 `scoreDocument` 和 `queryTerms` 添加评分与关键词拆分说明。

- `assistant.go`
  - 标注 `Run` 中先检索再调用模型的执行边界。
  - 标注模型接收的是增强 prompt。
  - 为 `buildPrompt` 添加 Prompt Augmentation 阶段注释。

## 验证结果

已执行 Go 测试、工作区状态检查和敏感信息/遗留标记检查。结果见本次会话最终回复。

## 后续建议

如果后续要演示生产级 RAG，可以在现有最小实现之外新增独立示例，逐步加入 embedding、向量库、BM25、reranker 和检索质量评估，避免和当前教学 demo 混在一起。
