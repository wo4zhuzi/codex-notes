# 2026-06-03 codex-plugin-cc 使用文章

## 任务背景

用户希望新增一篇文章，介绍 `openai/codex-plugin-cc` 插件在 Claude Code 中的使用方式。文章重点不是泛泛说明 Codex 与 Claude Code 可以协作，而是以插件形式说明：Claude Code 作为主工作台承载较长上下文，Codex 通过插件命令参与审查、对抗式审查和复杂 bugfix 救援。

## 根因定位

仓库已有 `codex-claude-collaboration.md`，主要覆盖 `AGENTS.md`、`CLAUDE.md`、文件型计划和跨工具接续规则，但缺少 `codex-plugin-cc` 的插件安装、命令使用、后台任务管理和实际协作案例。

因此本次不直接扩写原协作文档，避免把规则总览和插件实操混在一起；选择新增独立文章，并在 `README.md` 增加入口。

## 执行计划

1. 新增 `codex-plugin-cc.md`，介绍插件定位、安装流程、常用命令和复杂 Bugfix 协作流程。
2. 更新 `README.md` 的内容说明、目录结构、快速入口和使用建议。
3. 新增本日期变更记录，方便后续 AI 或开发者理解本次文档来源和边界。
4. 运行文档入口、核心命令和敏感占位符检查。

## 变更内容

- 新增 `codex-plugin-cc.md`：
  - 说明 `codex-plugin-cc` 是 Claude Code 插件，依赖本机 `codex` CLI 和登录状态。
  - 记录 `/plugin marketplace add openai/codex-plugin-cc`、`/plugin install codex@openai-codex`、`/reload-plugins` 和 `/codex:setup` 等安装初始化步骤。
  - 记录 `/codex:review`、`/codex:adversarial-review`、`/codex:rescue`、`/codex:status`、`/codex:result`、`/codex:cancel` 的适用场景。
  - 以复杂 Bugfix 为主案例，说明 Claude Code 主持上下文、Codex 挑战方案和后台救援的协作边界。
  - 补充排查清单和与 `codex-claude-collaboration.md` 的分工。
- 更新 `README.md`：
  - 内容说明中增加 `codex-plugin-cc` 插件化协作流程。
  - 目录结构和快速入口中加入 `codex-plugin-cc.md`。
  - 使用建议中增加 Claude Code 调用 Codex 审查或救援的阅读路径。

## 验证结果

```bash
git status --short
rg -n "codex-plugin-cc|/codex:review|/codex:rescue" README.md codex-plugin-cc.md
rg -n "TODO|FIXME|your-api-key|sk-" .
```

- `git status --short`：仅显示本次文档改动，包含 `README.md`、`codex-plugin-cc.md` 和本文档。
- `rg -n "codex-plugin-cc|/codex:review|/codex:rescue" README.md codex-plugin-cc.md`：确认 README 入口和新文章核心命令均存在。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有示例分支名 `project-task-branch` 和 `risk-reviewer` 名称，未发现真实密钥或遗留调试标记。

本次只修改 Markdown 文档，不涉及 `demos/` 下 Go 示例，因此不需要运行 `go test ./...`。

## 后续建议

如果后续 `openai/codex-plugin-cc` 增加新命令或调整安装方式，应优先更新 `codex-plugin-cc.md` 的安装与命令章节，并同步更新本文档的来源说明。
