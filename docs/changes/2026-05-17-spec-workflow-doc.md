# 2026-05-17 新增 Spec 模式实践文档

## 任务背景

用户希望将 Spec 模式生成流程整理进当前知识库，方便后续新项目参考。目标是保留 `AGENTS.md` 的导航地图定位，同时提供一份可复制、可执行的 Spec 实践手册。

## 根因定位

现有 `ai-workflow.md` 已说明 Spec 模式的理念，但还缺少可直接复用的实践步骤、文档模板和夜间委托口令。如果把这些完整内容写入 `AGENTS.md`，会让 `AGENTS.md` 变成长手册，违背“地图而非手册”的原则。

## 执行计划

1. 新增根目录 `spec-workflow.md`，作为 Spec 模式实践手册。
2. 在 `ai-workflow.md` 的 Spec 模式段落增加实践文档入口。
3. 更新 `README.md` 的目录结构、快速入口和使用建议。
4. 在 `AGENTS.md` 只补充 Spec 触发规则和文档位置，不展开完整流程。
5. 运行本地检查，确认链接和敏感词扫描无异常。

AI 自检结论：计划已覆盖根因定位、修改步骤、验证方式和风险边界；本次仅新增和更新知识库文档，不涉及配置模板或代码。

## 变更内容

- `spec-workflow.md`：新增 Spec 模式实践手册，包含适用场景、生成流程、文档模板、夜间委托和第二天验收口令。
- `ai-workflow.md`：补充 `spec-workflow.md` 入口。
- `README.md`：补充 `spec-workflow.md` 目录项、快速入口和阅读顺序。
- `AGENTS.md`：新增极短 Spec 模式硬规则，保持导航地图定位。
- `docs/changes/2026-05-17-spec-workflow-doc.md`：记录本次会话背景、根因、计划、变更和验证。

## 验证结果

- `git status --short --untracked-files=all`：显示 `README.md`、`ai-workflow.md` 已修改，`.gitignore`、`spec-workflow.md` 和多份 `docs/changes/` 记录为未跟踪文件；本次未回滚已有未提交内容。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中记录的检查命令示例，未发现真实密钥或遗留占位符。
- `rg -n "\\[.*\\]\\(([^)]+)\\)" README.md ai-workflow.md spec-workflow.md docs/changes/2026-05-17-spec-workflow-doc.md`：确认新增链接指向 `ai-workflow.md`、`spec-workflow.md`、`codex-core-commands.md`、`codex-skills.md` 等仓库内文档。

## 后续建议

后续新项目可先复制 `spec-workflow.md` 中的边界、方案、进度和夜间执行授权口令；实际项目 Spec 应存放到 `docs/specs/YYYY-MM-DD-<topic>.md`，执行事实仍写入 `docs/changes/`。
