# 2026-06-06 Codex Plugins 与 Understand Anything 使用方案

## 任务背景

用户希望新增 Codex plugin 说明，说明 plugin 和 skill 的区别，并规划 CodeGraph 和 Understand Anything 的最佳搭配使用方案。后续用户明确要求：`codex-plugins.md` 只介绍 plugin，Understand Anything 的 Codex 安装应放在 skill 相关文档中。

## 根因定位

仓库已有 `codex-skills.md` 说明 Codex Skills，也有 `codex-plugin-cc.md` 专门介绍 `openai/codex-plugin-cc`。但当前缺少一个独立文档回答：

- plugin 作为安装和分发单元的定位。
- plugin 与 skill 的职责边界。
- 第三方 plugin 的安装流程和安全边界。
- Understand Anything 与 CodeGraph 如何分工搭配。
- Understand Anything 在 Codex 侧应按 skills 安装方式记录，而不是混入 plugin 文档。

根因修正：Codex 有自己的 plugin 机制，可以通过 Codex app / CLI 插件目录和 `codex plugin marketplace ...` 管理；Understand Anything 支持 Codex，但 Codex 侧通过 `install.sh codex` 将相关能力接入 skills 目录，不应放在 `codex-plugins.md` 的主体中。

因此本次最终分工为：`codex-plugins.md` 只讲 Codex plugin；`codex-skills.md` 记录 Understand Anything 的 Codex 安装方式；`external-tools/codegraph.md` 只保留与 Understand Anything 的搭配入口。

## 执行计划

1. 查看 `git status --short`，确认工作区没有用户已有未提交改动。
2. 阅读 `README.md`、`codex-skills.md`、`codex-plugin-cc.md` 和 `external-tools/codegraph.md`，确认现有文档分工。
3. 新增 `codex-plugins.md`，说明 Codex plugin 定位、plugin 与 skill 的区别、Codex plugin 安装和 marketplace 管理方式。
4. 更新 `codex-skills.md`，补充 Understand Anything 的 Codex 安装、更新、卸载、验证和排查说明。
5. 更新 `README.md`，增加新文档入口和阅读顺序。
6. 更新 `external-tools/codegraph.md`，补充与 Understand Anything 的搭配入口并指向 skills 文档。
7. 新增本日期变更记录。
8. 运行文档入口和敏感占位符检查。

## 变更内容

- 新增 `codex-plugins.md`：
  - 说明 Codex plugin 是 Codex 的可安装分发单元，skill 是任务触发后的工作流程说明。
  - 对比 Codex plugin 与 Codex skill 在定位、内容、使用方式、适合场景和生命周期上的区别。
  - 记录 Codex CLI 中 `/plugins` 和 `codex plugin marketplace ...` 的基本用法。
  - 移除 Understand Anything 安装和 CodeGraph 搭配内容，只保留第三方工具通过 skills 接入时应记录到 skills 文档的边界说明。
- 更新 `codex-skills.md`：
  - 新增 Understand Anything 安装说明。
  - 记录 `install.sh codex`、更新、卸载、验证和排查命令。
  - 说明安装脚本会把相关 skills 链接到 `~/.agents/skills`，重启 Codex 后生效。
- 更新 `README.md`：
  - 内容说明、目录结构、快速入口和使用建议中加入 `codex-plugins.md`，并把 Understand Anything 接入口径移到 skills。
- 更新 `external-tools/codegraph.md`：
  - 更新时间调整为 2026-06-06。
  - 新增“与 Understand Anything 搭配”章节，说明全局图谱和精确调用链查询的分工，并链接到 `codex-skills.md`。

## 验证结果

```bash
git status --short
rg -n "Understand Anything|install.sh codex|~/.agents/skills|--uninstall codex|--update" codex-skills.md
rg -n "codex-plugins|Codex Plugins|Understand Anything|Codex Skills" README.md external-tools/codegraph.md docs/changes/2026-06-06-codex-plugins-understand-anything.md
rg -n "TODO|FIXME|your-api-key|sk-" .
```

- `git status --short`：仅显示本次计划内文档改动，包含 `README.md`、`codex-skills.md`、`external-tools/codegraph.md`、`codex-plugins.md` 和本文档。
- `rg -n "Understand Anything|install.sh codex|~/.agents/skills|--uninstall codex|--update" codex-skills.md`：确认 Understand Anything 的 Codex skills 安装、更新、卸载和验证说明均已写入。
- `rg -n "codex-plugins|Codex Plugins|Understand Anything|Codex Skills" ...`：确认 README 入口、CodeGraph 搭配入口和变更记录口径一致。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有 `project-task-branch` 示例分支名和 `risk-reviewer` 名称，未发现真实密钥或遗留调试标记。

本次只修改 Markdown 文档，不涉及 `demos/` 下 Go 示例，因此不需要运行 `go test ./...`。

## 后续建议

后续如果 Understand Anything 提供 Codex 兼容的 `.codex-plugin/plugin.json` 或 Codex marketplace，可再在 `codex-plugins.md` 中补充；在此之前，Codex 侧安装说明保留在 `codex-skills.md`。
