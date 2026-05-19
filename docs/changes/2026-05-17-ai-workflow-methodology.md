# 2026-05-17 新增 AI 编程工作流总纲

## 任务背景

用户希望把关于 AI 编程工作流和 harness engineering 的理解沉淀为一份 Markdown 文档，并补充 TDD 在流程中的位置。文档定位为方法论总纲，不替代 `AGENTS.md` 和 `codex-core-commands.md`。

## 根因定位

现有仓库已有 `AGENTS.md`、`codex-core-commands.md` 和 `codex-skills.md`，分别承担仓库导航、命令说明和 skills 说明。但缺少一份上层方法论文档，用于解释什么时候选择 vibe coding、plan 模式、spec 模式，以及如何把 TDD、机械验证、迭代自愈和上下文沉淀组合成完整工作流。

## 执行计划

1. 阅读 `README.md`、`AGENTS.md`、`codex-core-commands.md` 和 `codex-skills.md`，确认现有文档分工。
2. 新增根目录 `ai-workflow.md`，作为方法论总纲。
3. 在文档中覆盖三种协作模式、任务分级、TDD 建议、harness engineering 原则、上下文生命周期和项目落地清单。
4. 更新 `README.md` 的目录结构、快速入口和使用建议。
5. 运行本地检查，确认链接和敏感词扫描无异常。

AI 自检结论：计划已覆盖根因定位、修改步骤、验证方式和风险边界；本次新增方法论文档，不修改既有命令规则和配置模板。

## 变更内容

- `ai-workflow.md`：新增 AI 编程工作流方法论总纲。
- `README.md`：新增 `ai-workflow.md` 目录项、快速入口和阅读顺序。
- `docs/changes/2026-05-17-ai-workflow-methodology.md`：记录本次会话背景、根因、计划、变更和验证。

## 验证结果

- `git status --short --untracked-files=all`：显示 `README.md` 已修改，`.gitignore`、`ai-workflow.md` 和 `docs/changes/` 下多份变更记录为未跟踪文件；本次未回滚已有未提交内容。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中记录的检查命令示例，未发现真实密钥或遗留占位符。
- `rg -n "\\[.*\\]\\(([^)]+)\\)" README.md ai-workflow.md docs/changes/2026-05-17-ai-workflow-methodology.md`：确认新增 Markdown 链接集中在 `README.md` 和 `ai-workflow.md`，路径指向现有文档或目录。

## 后续建议

后续如果要把这套方法论复制到其他项目，建议只把摘要写进 `AGENTS.md`，把完整说明保留为独立工作流文档，避免 `AGENTS.md` 变成过长手册。
