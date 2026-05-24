# 调整 Function Calling Demo 为无状态回传函数结果

## 任务背景

用户使用 sub2 代理运行 Function Calling demo 时，第二次请求报错：

```text
previous_response_id is only supported on Responses WebSocket v2
```

需要修改 demo，让它兼容不支持 HTTP `previous_response_id` 续接的 OpenAI 兼容供应商代理。

## 根因定位

Go demo 初版在模型返回 `function_call` 后，第二次调用 Responses API 时使用 `PreviousResponseID` 回传 `function_call_output`。sub2 当前代理的 `/v1/responses` HTTP 接口返回 400，说明该兼容层只在 Responses WebSocket v2 支持 `previous_response_id`，不能用于当前 HTTP 请求链路。

## 执行计划

1. 去掉第二次 Responses API 请求里的 `PreviousResponseID`。
2. 第二次请求改为显式传入用户原始问题、模型生成的 `function_call` 和本地执行结果 `function_call_output`。
3. 增加单元测试，断言回传函数结果时不会再设置 `PreviousResponseID`。
4. 更新 README 和主文档，说明 demo 使用无状态回传方式以兼容供应商代理。
5. 执行 Go 测试、关键词检查和 diff 空白检查。

## 变更内容

- 更新 `assistant.go`：
  - 第二次请求不再设置 `PreviousResponseID`。
  - 使用 `ResponseInputItemParamOfMessage`、`ResponseInputItemParamOfFunctionCall` 和 `ResponseInputItemParamOfFunctionCallOutput` 组成无状态输入列表。
- 更新 `assistant_test.go`：
  - 新增 fake Responses client，覆盖完整两次请求链路。
  - 断言第二次请求没有 `PreviousResponseID`。
  - 断言第二次请求包含用户消息、函数调用和函数输出。
- 更新 `demos/function-calling-orders/README.md` 和 `function-calling.md`：
  - 说明 demo 不依赖 `previous_response_id`。
  - 标注该方式对 sub2 等 OpenAI 兼容供应商代理更友好。

## 验证结果

- 已执行 `go test ./...`：确认 Go demo 编译和单元测试通过。
- 已执行 `rg -n "PreviousResponseID|previous_response_id|function_call_output|sub2" demos/function-calling-orders function-calling.md docs/changes/2026-05-24-function-calling-stateless-output.md`：确认代码不再发送 `PreviousResponseID`，文档记录兼容说明。
- 已执行 `git diff --check`：未发现空白错误。

## 后续建议

如果某个供应商对 Responses API 的 function call item 仍不兼容，可以再增加 Chat Completions Function Calling 版本作为降级 demo；当前优先保持 Responses API 主线。
