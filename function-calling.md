# Function Calling 使用笔记

更新时间：2026-05-24。

本文记录 Function Calling 的运行原理、函数定义方式和一个最小订单查询 demo。本文以 OpenAI Responses API 为主，不讨论 Function Calling 是否会被其他工具形态替代。

参考来源：

- [OpenAI Function Calling](https://platform.openai.com/docs/guides/function-calling)
- [OpenAI Help Center: Function Calling in the OpenAI API](https://help.openai.com/en/articles/8555517-function-calling-in-the-openai-api)
- [OpenAI Go SDK](https://github.com/openai/openai-go)

## 结论

Function Calling 适合让模型把自然语言转换成结构化函数调用参数，再由应用程序执行真实函数、接口或数据库查询。

关键点：

- 模型不会真的执行函数。
- API 请求中只传函数说明和 JSON Schema，不传真实函数代码。
- 模型返回要调用的函数名和参数。
- 应用程序负责执行函数，并把执行结果回传给模型。
- 模型基于函数执行结果生成最终自然语言回答。

## 运行原理

完整链路如下：

```text
用户问题
-> 应用把函数 schema 和用户问题发给模型
-> 模型判断是否需要调用函数
-> 模型返回 function_call：函数名 + JSON 参数 + call_id
-> 应用解析参数并执行本地函数或业务接口
-> 应用把结果作为 function_call_output 回传给模型
-> 模型根据函数结果生成最终回答
```

以查订单为例：

```text
用户：帮我查一下订单 ORD-1001 到哪了
模型：需要调用 get_order_status({"order_id": "ORD-1001"})
应用：读取本地订单数据并返回订单状态
模型：根据订单状态回复用户
```

Function Calling 的核心价值不是让模型直接访问系统，而是让模型“决定调用什么、生成什么参数”，真实执行仍由应用控制。

## 如何定义函数

在 Responses API 中，函数通过 `tools` 定义。最小结构如下：

```go
tools := []responses.ToolUnionParam{{
	OfFunction: &responses.FunctionToolParam{
		Name:        "get_order_status",
		Description: openai.String("根据订单号查询订单状态、物流信息和预计送达时间。"),
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"order_id": map[string]string{
					"type":        "string",
					"description": "订单号，例如 ORD-1001。",
				},
			},
			"required":             []string{"order_id"},
			"additionalProperties": false,
		},
		Strict: openai.Bool(true),
	},
}}
```

字段含义：

| 字段 | 作用 |
| --- | --- |
| `type` | 固定为 `function`，表示这是一个函数工具。 |
| `name` | 函数名，模型会返回这个名字，应用据此分发到真实函数。 |
| `description` | 告诉模型什么时候应该调用这个函数。 |
| `parameters` | JSON Schema，用来约束模型生成的参数。 |
| `required` | 必填参数列表。 |
| `additionalProperties` | 是否允许 schema 外的额外字段，生产环境建议设为 `False`。 |
| `strict` | 要求模型生成的参数严格匹配 schema。 |

函数定义要写清楚“什么时候调用”和“参数是什么”，不要把真实密钥、数据库连接串或内部服务 URL 写进 schema。

## 应用如何执行函数调用

模型返回的响应里可能包含一个或多个 `function_call`。应用需要遍历响应项，找到函数调用并执行：

```go
for _, item := range response.Output {
	if item.Type != "function_call" {
		continue
	}

	toolCall := item.AsFunctionCall()
	result := callFunction(toolCall.Name, toolCall.Arguments)

	finalResponse, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Model:              openai.ResponsesModel(model),
		PreviousResponseID: openai.String(response.ID),
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: []responses.ResponseInputItemUnionParam{{
				OfFunctionCallOutput: &responses.ResponseInputItemFunctionCallOutputParam{
					CallID: toolCall.CallID,
					Output: responses.ResponseInputItemFunctionCallOutputOutputUnionParam{
						OfString: openai.String(result),
					},
				},
			}},
		},
		Tools: tools,
	})
}
```

这里的 `call_id` 很重要。它用于告诉模型“这份输出对应刚才哪一次函数调用”。

## Codex 如何参与

Function Calling 的运行时不是 Codex CLI。典型分工是：

```text
Codex：帮你写 demo、解释代码、运行检查、排查错误
main.go：作为应用调用 OpenAI API
OpenAI 模型：生成 function_call
本地函数：执行订单查询
```

也就是说，Codex 不会替你的业务系统自动执行 Function Calling；你需要在自己的应用代码里实现“发送 schema、处理 function_call、执行函数、回传 function_call_output”这条链路。

## 订单查询 Demo

示例目录：

```text
demos/function-calling-orders/
├── README.md
├── assistant.go
├── go.mod
├── go.sum
├── main.go
├── orders.go
├── orders_test.go
├── orders.json
├── server.go
├── server_test.go
└── templates/
    └── index.html
```

运行方式：

```bash
cd demos/function-calling-orders
export OPENAI_API_KEY="你的 API Key"
go run . "帮我查一下订单 ORD-1001 到哪了"
```

如果希望用网页形式体验：

```bash
go run . -web
```

然后在浏览器打开：

```text
http://127.0.0.1:8080
```

这个 demo 会：

1. 把 `get_order_status` 的 schema 发给模型。
2. 让模型从用户问题中提取订单号。
3. 应用侧执行 `get_order_status(order_id)`。
4. 从本地 `orders.json` 查询订单。
5. 把查询结果回传给模型。
6. 打印最终自然语言回答。

Web 版链路如下：

```text
浏览器输入问题
-> Go HTTP /api/chat
-> OpenAI Function Calling
-> get_order_status 读取 orders.json
-> Go HTTP 返回最终回答
-> 浏览器展示回答
```

## 适用场景

适合使用 Function Calling：

- 查订单、查物流、查账户余额、查审批状态。
- 创建工单、创建日程、发送通知等受控业务动作。
- 根据用户自然语言生成结构化查询参数。
- 需要模型在多个业务函数之间选择一个来调用。
- 需要强约束参数格式，避免模型自由编造字段。

不适合使用 Function Calling：

- 只需要普通问答、总结、翻译或文案生成。
- 没有外部数据或业务动作需要执行。
- 高风险写操作没有权限校验、审计和二次确认。
- 想让多个 AI 客户端复用同一外部工具服务，这类场景更适合 MCP。

## 生产建议

- 函数 schema 尽量小而明确，一个函数只做一类动作。
- 参数必须在应用侧再次校验，不能只相信模型输出。
- 写操作、删除操作、支付和权限变更必须增加确认、鉴权和审计。
- 函数返回结果建议使用 JSON 字符串，便于模型稳定理解。
- 对函数调用失败、订单不存在、权限不足等情况返回结构化错误。
- 不要把真实 API Key、个人 Token、私有地址或内部服务 URL 写入文档和 demo。
