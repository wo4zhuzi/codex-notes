# 新增 n8n AI 重构工作流专题

## 任务背景

用户希望新增一篇独立专题，说明在 AI 参与重构项目时，n8n 具体能做什么，并要求不要绑定本仓库自身的知识库规则，而是给出可落地的实战方案。

## 根因定位

前序讨论中对 n8n 的说明偏抽象，缺少一个围绕真实重构项目的端到端方案。需要把 n8n 从“自动化工具”具体落到 GitHub PR、CI、AI agent、通知、审批和 MCP 工具网关这几个工程节点上。

## 执行计划

1. 确认工作区状态，避免覆盖用户已有改动。
2. 新增一篇独立专题，按真实重构场景说明 n8n 的职责、workflow、输入输出、AI prompt 和权限边界。
3. 更新根目录 README 和 `external-tools/README.md`，补充专题入口。
4. 生成本次日期变更记录。
5. 运行文档级检查，确认链接、关键词和敏感信息没有明显问题。

## 变更内容

- 新增 `external-tools/n8n-ai-refactor-workflow.md`：
  - 定义 AI agent、n8n、GitHub / CI 的分工。
  - 给出重构计划质量检查、Diff 偏离计划检测、CI 失败摘要、高风险人工审批、重构日报与上下文恢复五个场景。
  - 给出 n8n 节点链路、JSON 输入输出、AI prompt 模板、MCP 工具接口和第一周落地路线。
  - 补充凭证、审计、模型输出和失败处理边界。
- 更新 `external-tools/README.md`：
  - 在外部工具列表中增加 n8n 专题入口。
- 更新 `README.md`：
  - 在内容说明、目录结构、快速入口和使用建议中增加 n8n 专题。

## 验证结果

已执行：

```bash
git status --short
rg -n "n8n-ai-refactor-workflow|AI 重构项目中 n8n 可以做什么|n8n" README.md external-tools/README.md external-tools/n8n-ai-refactor-workflow.md docs/changes/2026-06-11-n8n-ai-refactor-workflow.md
rg -n "TODO|FIXME|your-api-key|sk-" .
```

结果：

- `git status --short` 显示本次仅修改 `README.md`、`external-tools/README.md`，并新增 `external-tools/n8n-ai-refactor-workflow.md` 和本变更记录。
- n8n 入口检查命中 `README.md`、`external-tools/README.md`、专题文档和本变更记录，确认入口已接上。
- 敏感信息扫描仅命中文档中的检查命令示例、既有 `risk-reviewer` 名称、既有 `project-task-branch` 示例分支名和 Codex 指令示例，未发现真实 API Key、Token、私有代理地址或遗留调试标记。

## 后续建议

- 如果后续要继续细化，可新增 n8n workflow JSON 蓝图或截图说明。
- 如果用户已有具体 GitHub / CI / 飞书环境，可再补一版面向单一平台的配置手册。
