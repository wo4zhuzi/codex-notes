# 2026-05-17 AI 工作流工程化补强

## 任务背景

用户希望在不修改 `.gitignore`、不处理 `AGENTS.md` 提交策略的前提下，先补充 AI 项目落地清单、AI 权限边界，以及 Git / commit / PR 操作边界。

## 根因定位

当前知识库已经覆盖 vibe、plan、spec、TDD、skills、resume、compact 和变更记录，但更偏方法论说明。新项目实际落地时，还缺少一份可复制的检查清单，以及对 AI 能做什么、必须确认什么、Git 操作边界是什么的明确说明。

## 执行计划

AI 自检结论：计划覆盖根因定位、修改步骤、验证方式和风险边界。本次只补文档，不修改 `.gitignore`，不调整 `AGENTS.md` 提交策略，不执行 commit。

计划：

1. 新增 `ai-project-checklist.md`，作为新项目接入 AI 协作的落地清单。
2. 更新 `ai-workflow.md`，补充 AI 权限边界，并链接项目落地清单。
3. 更新 `codex-core-commands.md`，补充 Git / commit / PR 操作边界。
4. 更新 `README.md`，加入新文档入口和阅读顺序。
5. 执行链接、敏感信息和工作区状态检查。

## 变更内容

- `ai-project-checklist.md`：新增 AI 项目落地清单，覆盖启动前检查、协作规则、文档结构、验证命令、Git 工作流、权限边界、发布和回滚。
- `ai-workflow.md`：新增“AI 权限边界”章节，并在项目落地清单中链接 `ai-project-checklist.md`。
- `codex-core-commands.md`：新增“Git / commit / PR 操作边界”章节，说明 `git status` 检查、commit 授权、PR 描述和提交前顺序。
- `README.md`：加入 `ai-project-checklist.md` 的目录结构、快速入口和阅读建议。

## 验证结果

已执行：

```bash
rg -n "\[.*\]\(([^)]+)\)" README.md ai-workflow.md codex-core-commands.md ai-project-checklist.md docs/changes/2026-05-17-ai-workflow-hardening.md
rg -n "TODO|FIXME|your-api-key|sk-" .
git status --short --untracked-files=all
```

结果：

- Markdown 链接扫描命中 `README.md` 和 `ai-workflow.md` 中的仓库内链接，新增 `ai-project-checklist.md` 入口已出现。
- 敏感信息扫描仅命中文档中记录的检查命令示例，未发现真实密钥。
- `git status --short --untracked-files=all` 显示本次修改 `README.md`、`ai-workflow.md`、`codex-core-commands.md`，新增 `ai-project-checklist.md` 和本变更记录；仓库中仍存在此前未跟踪的 `.gitignore` 与历史变更记录，本次未修改 `.gitignore`。

## 后续建议

- 后续如果希望 `docs/specs/` 也进入 Git 跟踪，再单独评估 `.gitignore`。
- 如果未来把本仓库从个人知识库升级为团队模板，再决定是否提交项目级 `AGENTS.md`。
