# 补充 Function Calling Demo 供应商代理配置说明

## 任务背景

用户使用 sub2 代理的供应商 API 和 Key 运行 Function Calling demo，需要文档说明如何配置供应商 Key、OpenAI 兼容接口地址和模型名。

## 根因定位

Go demo 使用官方 OpenAI Go SDK 的 `openai.NewClient()`，SDK 会自动读取 `OPENAI_API_KEY` 和 `OPENAI_BASE_URL`。此前 demo README 只说明了 `OPENAI_API_KEY`，没有说明供应商代理场景需要额外配置 base URL，也没有提醒一次性命令可以避免影响其他程序。

## 执行计划

1. 在 demo README 的“准备环境”中补充供应商代理配置。
2. 使用占位符展示 `OPENAI_API_KEY`、`OPENAI_BASE_URL` 和 `OPENAI_MODEL`，不写真实 Key 或私有代理地址。
3. 增加一次性命令示例，说明如何避免污染当前终端环境。
4. 在 `function-calling.md` 的运行方式处补充一句，指向 demo README 查看供应商代理配置。
5. 执行关键词检查和 Markdown diff 空白检查。

## 变更内容

- 更新 `demos/function-calling-orders/README.md`：
  - 增加 sub2 或其他 OpenAI 兼容供应商代理配置说明。
  - 说明 `OPENAI_BASE_URL` 通常需要使用供应商提供的 `/v1` base URL。
  - 说明供应商模型不兼容默认模型时可设置 `OPENAI_MODEL`。
  - 增加一次性命令示例，避免影响后续终端程序。
- 更新 `function-calling.md`：
  - 在运行方式下补充供应商代理配置入口说明。

## 验证结果

- 已执行 `rg -n "OPENAI_API_KEY|OPENAI_BASE_URL|OPENAI_MODEL|供应商|sub2" demos/function-calling-orders/README.md function-calling.md`：确认供应商代理配置说明已写入。
- 已执行 `rg -n "真实 API Key|sk-|TODO|FIXME" demos/function-calling-orders/README.md function-calling.md`：仅命中安全提示语，未发现真实 Key 或遗留占位标记。
- 已执行 `git diff --check`：未发现 Markdown 空白错误。

## 后续建议

如果后续发现某个供应商不完全兼容 Responses API，需要在 README 中单独标注该供应商的限制或推荐模型名；当前说明按 OpenAI 兼容代理处理。
