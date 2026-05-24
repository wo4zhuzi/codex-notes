# 替换 Function Calling Demo 为 Go 版本

## 任务背景

用户要求将当前 Function Calling demo 替换成 Golang 版本，并确认保留命令行和网页两种体验方式。

## 根因定位

迁移前的 Function Calling 实现不只是单个脚本，而是文章、demo README、命令行入口、Web 入口和依赖说明都绑定了旧语言版本：

- `function-calling.md` 中示例代码、运行方式和 Web 链路使用旧版本说明。
- `demos/function-calling-orders/` 中包含旧语言入口、公共逻辑文件和依赖文件。
- Web 页面调用 `/api/chat`，后端由旧 Web 框架提供。

因此需要整体替换实现和文档，而不是只改函数调用片段。

## 执行计划

1. 使用官方 OpenAI Go SDK `github.com/openai/openai-go/v3` 实现 Responses API Function Calling。
2. 使用 Go 标准库 `net/http` 保留网页版本，避免引入额外 Web 框架。
3. 复用现有 `orders.json` 和 `templates/index.html`。
4. 删除旧入口、旧依赖和生成缓存文件。
5. 更新 `function-calling.md` 和 demo README 中的示例、运行命令和关键文件说明。
6. 运行 Go 测试、无 API Key 行为验证、关键词检查和 diff 空白检查。

## 变更内容

- 新增 Go 模块文件 `go.mod`、`go.sum`。
- 新增 `assistant.go`：
  - 定义 `get_order_status` 函数工具 schema。
  - 调用 OpenAI Responses API。
  - 处理模型返回的 `function_call`。
  - 执行本地订单查询并用 `function_call_output` 回传结果。
- 新增 `orders.go` 和 `orders_test.go`：
  - 读取 `orders.json`。
  - 查询存在和不存在的订单。
- 新增 `server.go` 和 `server_test.go`：
  - 使用 `net/http` 提供 `/` 和 `/api/chat`。
  - 验证正常响应、空消息和后端错误。
- 新增 `main.go`：
  - 默认作为命令行 demo 运行。
  - 使用 `-web` 启动网页版本，默认监听 `127.0.0.1:8080`。
- 删除旧入口文件、旧依赖文件和生成缓存目录。
- 更新 `function-calling.md` 和 `demos/function-calling-orders/README.md`，替换为 Go 版本说明。

## 验证结果

- 已执行 `go test ./...`：Go demo 编译通过，订单查询和 HTTP handler 单元测试通过。
- 已执行 `go run . "帮我查一下订单 ORD-1001 到哪了"`：未设置 `OPENAI_API_KEY` 时返回明确错误提示。
- 已执行旧版本关键词扫描：确认 Function Calling 相关文档没有残留旧版本运行说明。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：用于检查遗留占位符和疑似密钥。
- 已执行 `git diff --check`：检查 Markdown 和 Go 文件 diff 空白。

## 后续建议

如果后续要演示真实业务接口，可在 `GetOrderStatus` 后面接入 HTTP/RPC 客户端，并为网络错误、权限不足和订单不存在分别返回结构化错误；当前 demo 继续保持本地 JSON 数据，便于理解 Function Calling 主链路。
