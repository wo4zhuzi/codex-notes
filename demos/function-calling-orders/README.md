# Function Calling 订单查询 Demo

这个 demo 演示如何用 OpenAI Responses API 的 Function Calling 查询本地订单数据。目录中同时包含 Go 命令行版本和 Go 标准库 HTTP 网页版本。

运行链路：

```text
用户问题
-> 模型生成 get_order_status 的 function_call
-> main.go 执行本地函数并读取 orders.json
-> main.go 回传 function_call_output
-> 模型生成最终回答
```

demo 回传函数结果时使用无状态请求：第二次请求会显式带上用户问题、模型生成的 `function_call` 和本地 `function_call_output`，不依赖 `previous_response_id`。这种方式对 sub2 等 OpenAI 兼容供应商代理更友好。

## 准备环境

```bash
cd demos/function-calling-orders
go mod tidy
```

设置 API Key：

```bash
export OPENAI_API_KEY="你的 API Key"
```

如果使用 sub2 或其他 OpenAI 兼容供应商代理，需要同时设置供应商 Key 和接口地址：

```bash
export OPENAI_API_KEY="你的供应商 API Key"
export OPENAI_BASE_URL="https://你的供应商接口地址/v1"
```

如果供应商不支持默认模型，可指定供应商支持的模型名：

```bash
export OPENAI_MODEL="供应商支持的模型名"
```

如果只想让环境变量影响本次 demo，不影响当前终端后续启动的其他程序，可以使用一次性命令：

```bash
OPENAI_API_KEY="你的供应商 API Key" \
OPENAI_BASE_URL="https://你的供应商接口地址/v1" \
OPENAI_MODEL="供应商支持的模型名" \
go run . "帮我查一下订单 ORD-1001 到哪了"
```

`OPENAI_BASE_URL` 是否需要带 `/v1` 以供应商文档为准；OpenAI 兼容代理通常使用 `/v1` 作为 base URL。

不要把真实 API Key 写入仓库文件。

## 运行

### 命令行版本

```bash
go run . "帮我查一下订单 ORD-1001 到哪了"
```

也可以查询其他示例订单：

```bash
go run . "订单 ORD-1002 现在是什么状态？"
go run . "查一下 ORD-1003 的物流信息"
```

### 网页版本

```bash
go run . -web
```

浏览器打开：

```text
http://127.0.0.1:8080
```

在页面输入：

```text
帮我查一下订单 ORD-1001 到哪了
```

## 示例订单

- `ORD-1001`：已发货。
- `ORD-1002`：待支付。
- `ORD-1003`：已签收。

## 关键文件

- `orders.json`：本地订单数据。
- `assistant.go`：Function Calling 公共流程。
- `orders.go`：本地订单读取和查询逻辑。
- `server.go`：网页版本 HTTP 路由和 `/api/chat` 接口。
- `main.go`：命令行入口和网页启动入口。
- `templates/index.html`：网页聊天界面。
- `go.mod`、`go.sum`：Go 模块依赖。

## 验证

```bash
go test ./...
```

## 注意事项

- 这个 demo 不访问真实订单系统。
- 模型只负责生成函数调用参数，不会直接读取本地文件。
- 本地文件读取由 `orders.go` 中的 `GetOrderStatus` 执行。
- demo 不使用 `previous_response_id`，避免部分供应商代理不支持 Responses HTTP 续接。
- Web 页面不会保存聊天记录，刷新后页面消息会清空。
- 如果没有设置 `OPENAI_API_KEY`，demo 会直接报错并提示设置环境变量。
