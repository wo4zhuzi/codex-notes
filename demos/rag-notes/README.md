# RAG 笔记助手 Demo

这个 demo 演示最小 RAG 链路：先检索本地 Markdown 知识库，再把检索到的内容交给模型回答。目录中同时包含 Go 命令行版本和 Go 标准库 HTTP 网页版本。

运行链路：

```text
用户问题
-> retriever.go 检索 knowledge/*.md
-> assistant.go 把检索结果拼进 prompt
-> OpenAI Responses API 生成答案
-> 命令行或网页返回答案和来源
```

这个 demo 不使用 embedding 和向量数据库，目的是先看懂 RAG 在项目中的位置。

## RAG 执行过程详解

当前 demo 的 RAG 链路由 Go 程序和模型共同完成，但两者职责不同：

```text
Go 程序：加载知识库 -> 检索相关文档 -> 组装增强 prompt -> 调用模型
模型：读取增强 prompt -> 基于用户问题和检索上下文生成答案
```

这里的“检索”不是模型自己完成的，也不是把整个 `knowledge/` 目录直接发送给模型。程序启动时，`retriever.go` 会读取 `knowledge/*.md`，把每篇 Markdown 文档保存到内存中的 `Retriever.docs`。用户提问后，Go 程序会从问题中提取关键词，在已加载文档的文件名、标题和正文中做匹配与评分，然后按分数取前 `topK` 篇文档作为检索结果。

拿到检索结果后，`assistant.go` 会进入 prompt augmentation 阶段：把“用户问题”和“检索命中的文档内容”拼成一个新的 prompt，再通过 OpenAI Responses API 发给模型。模型只接收这份增强后的 prompt，负责基于上下文总结、对比和组织答案，并按要求输出引用来源。

因此，这个 demo 的专业分工可以描述为：

```text
Knowledge Loading：Go 启动时读取本地 Markdown 知识库。
Retrieval：Go 根据用户问题做关键词召回和相关性排序。
Prompt Augmentation：Go 把用户问题和 topK 检索结果拼接成模型输入。
Generation：模型基于增强 prompt 生成面向用户的自然语言答案。
```

## 准备环境

```bash
cd demos/rag-notes
go mod tidy
```

设置 API Key：

```bash
export OPENAI_API_KEY="你的 API Key"
```

如果使用 OpenAI 兼容供应商代理，需要同时设置供应商 Key 和接口地址：

```bash
export OPENAI_API_KEY="你的供应商 API Key"
export OPENAI_BASE_URL="https://你的供应商接口地址/v1"
```

如果供应商不支持默认模型，可指定供应商支持的模型名：

```bash
export OPENAI_MODEL="供应商支持的模型名"
```

不要把真实 API Key 写入仓库文件。

## 运行

### 命令行版本

```bash
go run . "RAG 是什么？"
go run . "RAG 和 Function Calling 有什么区别？"
go run . -k 5 "MCP 和 RAG 是什么关系？"
```

### 网页版本

```bash
go run . -web
```

浏览器打开：

```text
http://127.0.0.1:8081
```

## 关键文件

- `knowledge/*.md`：本地 Markdown 知识库。
- `retriever.go`：关键词检索和评分逻辑。
- `assistant.go`：RAG prompt 组装和 OpenAI 调用。
- `server.go`：网页版本 HTTP 路由和 `/api/chat` 接口。
- `main.go`：命令行入口和网页启动入口。
- `templates/index.html`：网页聊天界面。

## 这个 demo 属于哪种 RAG

它属于最小 RAG：

```text
本地文档
-> 关键词检索
-> 上下文注入
-> 模型回答
```

它不是生产级 Hybrid RAG。生产系统通常还会加入：

- embedding 和向量数据库，支持语义检索。
- BM25 或全文检索，支持精确命中。
- RRF 融合，合并多路检索结果。
- reranker，重排候选片段。
- 权限过滤、版本过滤和引用来源。
- 检索质量评估和用户反馈闭环。

## 验证

```bash
go test ./...
```
