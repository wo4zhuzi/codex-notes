# 2026-05-26 Agent Evaluation Reference 模板

## 任务背景

用户指出 `agent-evaluation` 多场景 skill 方案中只记录了 references 目录结构和创建提示词，但没有把各个 reference Markdown 文件实际写出来，希望新增文件夹保存这些场景标准，并在 `codex-skills.md` 和 `agent-evaluation.md` 中补充引用。

## 根因定位

此前 `codex-skills.md` 已说明“一个 `SKILL.md` 总入口 + 多个 `references/*.md`”的结构，但只停留在设计层面。缺少可复用的仓库模板会导致后续创建全局 skill 时仍需重新整理每个场景的边界、必须验证项和停止线。

## 执行计划

- 新增 `docs/agent-evaluation-references/` 目录。
- 新增 8 个场景 reference 模板：文档、bugfix、新功能、重构、配置、工具集成、code review、高风险任务。
- 在 `agent-evaluation.md` 中补充这些模板的入口链接。
- 在 `codex-skills.md` 中说明这些模板可复制到全局 skill 的 `references/` 目录。
- 运行仓库约定的本地检查命令。

## 变更内容

- 新增 `docs/agent-evaluation-references/docs.md`。
- 新增 `docs/agent-evaluation-references/bugfix.md`。
- 新增 `docs/agent-evaluation-references/feature.md`。
- 新增 `docs/agent-evaluation-references/refactor.md`。
- 新增 `docs/agent-evaluation-references/config.md`。
- 新增 `docs/agent-evaluation-references/tool-integration.md`。
- 新增 `docs/agent-evaluation-references/code-review.md`。
- 新增 `docs/agent-evaluation-references/high-risk.md`。
- 更新 `agent-evaluation.md`，补充 reference 模板入口。
- 更新 `codex-skills.md`，说明创建全局 skill 时可复制或参考这些模板。

## 验证结果

- 已执行 `git status --short`：确认本次新增 `docs/agent-evaluation-references/` 场景模板，并修改 `agent-evaluation.md`、`codex-skills.md`。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `find docs/agent-evaluation-references -maxdepth 1 -type f | sort`：确认 8 个 reference 模板文件已创建。
- 已执行 `rg -n "agent-evaluation-references|references/" agent-evaluation.md codex-skills.md`：确认 `agent-evaluation.md` 和 `codex-skills.md` 已补充 reference 模板入口与全局 skill 复制说明。

## 后续建议

- 后续真正创建全局 `agent-evaluation` skill 时，可直接把 `docs/agent-evaluation-references/*.md` 复制到 `~/.codex/skills/agent-evaluation/references/`。
- 如果场景继续增多，应优先新增 reference 文件，而不是把所有细节塞回 `SKILL.md`。
