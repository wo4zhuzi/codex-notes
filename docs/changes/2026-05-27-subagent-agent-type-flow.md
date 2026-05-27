# 2026-05-27 Subagent 文档拆分与调用流程说明

## 任务背景

用户确认：即使仓库内存在 `cc-switch-configs/subagents/*.toml`，当前 Codex 会话实际创建 subagent 时仍可能只暴露 `default`、`explorer`、`worker` 等通用类型，导致自定义 TOML 中的 `model_reasoning_effort` 未自动生效。

随后用户认为把完整 subagent 注册和调用说明放在 `cc-swtich.md` 中不合适，希望单独新增 `subagent.md`，把相关内容拎出来并详细介绍如何注册 agent。

## 根因定位

`cc-swtich.md` 原本已说明模板需要复制到 `~/.codex/agents/` 或 `.codex/agents/`，但推荐口令没有明确区分两种情况：

1. 自定义 agent 已被运行时注册成可选 `agent_type`。
2. 运行时未注册自定义 agent，只能用通用 `explorer` 按提示词扮演角色。

第二种情况下，subagent 不会自动读取 `cc-switch-configs/subagents/*.toml`，推理强度也可能继承主 agent。

同时，`cc-swtich.md` 的职责应集中在 CC-Switch 安装、配置切换和场景模板说明；完整 subagent 注册、调用与 fallback 说明更适合独立文档承载。

## 执行计划

1. 新增根目录 `subagent.md`，集中说明 subagent 概念、模板源、注册方式、运行时确认、调用方式、推理强度和推荐口令。
2. 精简 `cc-swtich.md` 的 Subagent 场景配置，只保留 CC-Switch 场景模板入口和跳转链接。
3. 更新 `README.md`，补充 `subagent.md` 到内容说明、目录结构、快速入口和使用建议。
4. 更新本变更记录，记录拆分背景、变更内容和验证结果。

## 变更内容

- `subagent.md`：新增 Codex Subagent 使用笔记，详细说明如何注册 agent、如何确认自定义 `agent_type` 是否生效，以及未注册时的通用 `explorer` fallback 行为。
- `cc-swtich.md`：保留 `dev-subagent-*.toml` 场景配置说明，并把详细注册和调用说明迁移到 `subagent.md`。
- `README.md`：新增 `subagent.md` 索引和阅读建议。

## 验证结果

- `git status --short`：确认本次修改 `README.md`、`cc-swtich.md`，并新增 `subagent.md` 和本变更记录。
- `git diff --check`：通过，未发现空白错误。
- `python3 -B -c 'import pathlib,tomllib; ...'`：成功解析 `cc-switch-configs/` 下 14 个 TOML 文件。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有示例分支名 `project-task-branch` 和 `risk-reviewer` 名称，未发现真实密钥或遗留调试标记。
- `command -v markdownlint`：本机未安装，未运行可选 Markdown lint。

## 后续建议

后续如果 Codex 运行时支持列出已注册自定义 agent，应把对应确认命令补进 `subagent.md`，替代当前基于工具 schema 的人工确认说明。
