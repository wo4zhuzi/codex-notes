# 新增 Function Calling 原理文章与订单查询 Demo

## 任务背景

用户希望写一篇关于 Function Calling 的小文章，并配套一个可运行 demo。文章重点不是讨论 Function Calling 是否淘汰，而是说明 Function Calling 的运行原理、如何定义函数，以及应用程序如何调用 OpenAI API、接收模型返回的函数调用并执行本地函数。后续用户希望 demo 支持网页形式，因此在命令行版本之外补充 Go Web 版本。

## 根因定位

仓库此前没有 Function Calling 专题文档，也没有可运行 demo。现有 `mcp.md` 只说明 MCP 原理；`codex-skills.md` 说明 Skills；但 Function Calling 与 MCP、Skill 的层级不同，需要单独用文章和 demo 解释：

- 函数 schema 如何定义。
- 模型返回的 `function_call` 是什么。
- 应用侧如何解析参数并执行真实函数。
- 应用侧如何用 `function_call_output` 把结果回传给模型。
- Codex CLI 在这里是辅助编写和调试 demo，不是 Function Calling runtime。

## 执行计划

1. 新增 `function-calling.md`，记录 Function Calling 原理、函数定义方式、应用执行链路和生产建议。
2. 新增 `demos/function-calling-orders/`，提供本地订单查询 demo。
3. 抽出 Function Calling 公共逻辑，复用到命令行版本和 Web 版本。
4. 新增 Go HTTP 页面和 `/api/chat` 接口。
5. 更新 `README.md`，加入 Function Calling 文章和 demo 目录入口。
6. 执行 Go 测试、关键词检查、敏感占位符检查和 Markdown 空白检查。

## 变更内容

- 新增 `function-calling.md`：
  - 说明 Function Calling 的完整链路。
  - 说明 `tools` 中函数 schema 的关键字段。
  - 说明模型不会执行函数，真实执行发生在应用侧。
  - 说明 Codex CLI 只负责辅助编写、运行和调试 demo。
- 新增 `demos/function-calling-orders/assistant.go`：
  - 使用 OpenAI Responses API。
  - 定义 `get_order_status` 函数工具。
  - 读取本地 `orders.json` 查询订单。
  - 处理 `function_call` 并回传 `function_call_output`。
- 新增 `demos/function-calling-orders/main.go`：
  - 提供命令行入口和 Web 启动入口。
- 新增 `demos/function-calling-orders/server.go` 和 `templates/index.html`：
  - 提供 Go HTTP Web 页面和 `/api/chat` 接口。
  - 页面输入自然语言问题，后端执行 Function Calling 并返回最终回答。
- 新增 `demos/function-calling-orders/orders.go`、`orders.json`、`go.mod`、`go.sum` 和 `README.md`。
- 更新根目录 `README.md`，补充 Function Calling 入口和 demo 目录。

## 验证结果

- 已执行 `go test ./...`：确认 Go demo 编译通过，订单查询、函数参数解析和 HTTP handler 测试通过。
- 已执行 `go run . "帮我查一下订单 ORD-1001 到哪了"`：未设置 `OPENAI_API_KEY` 时返回明确错误提示。
- 已执行 `rg -n "Function Calling|function_call|get_order_status|ORD-1001|OPENAI_API_KEY|Go HTTP|/api/chat" function-calling.md demos/function-calling-orders README.md`：确认文章、命令行 demo 和 Web demo 入口完整。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `git diff --check`：未发现空白错误。

## 后续建议

如果后续要把 demo 扩展成真实接口服务，可在 `GetOrderStatus` 后接入 HTTP/RPC 客户端，演示 Function Calling 调用业务 API；当前版本保持最小可运行，便于理解原理。
