# 2026-06-01 补充项目级 Skills 说明

## 任务背景

用户在其他项目中添加了项目级 skill，例如 `.agents/skills/redesign-existing-projects/SKILL.md`，同时生成了 `skills-lock.json`，希望理解这些文件的作用，以及是否必须在 `AGENTS.md` 中写明才能被 Codex 发现。

## 根因定位

仓库已有 `codex-skills.md`，但此前主要说明全局 skills 和系统内置 skills，例如 `~/.codex/skills/{skill-name}/SKILL.md` 和 `~/.codex/skills/.system/`。缺少项目级 `.agents/skills/`、`skills-lock.json`、`AGENTS.md` 与 skill 发现关系的说明，容易让用户误以为 `AGENTS.md` 是 skill 注册表。

实际应区分：

- `SKILL.md`：skill 的主说明文件，定义触发场景和执行流程。
- `.agents/skills/`：当前仓库的项目级 skill 目录。
- `skills-lock.json`：项目级 skills 的锁文件或清单。
- `AGENTS.md`：项目规则和协作约束，不负责注册 skill，但可说明何时优先使用某个 skill。

## 执行计划

AI 自检结论：计划覆盖根因定位、修改步骤、验证方式和风险边界。本次仅修改文档，不创建实际项目级 skill，不修改配置模板或 demo。

计划：

1. 在 `codex-skills.md` 的“存放位置”后新增“项目级 Skills”说明。
2. 在显式触发、判断是否启用、常见误判、推荐实践和排查流程中补充项目级 skill 相关内容。
3. 新增本日期变更记录。
4. 执行关键词检查、敏感信息检查和 git 状态检查。

## 变更内容

- 更新 `codex-skills.md`：
  - 新增项目级 skill 常见目录结构。
  - 补充系统 skill、全局 skill、项目级 skill 的区别。
  - 说明 `skills-lock.json` 的锁文件职责和提交前安全检查。
  - 明确 `AGENTS.md` 不负责注册 skill，但适合记录项目级 skill 的使用场景。
  - 增加 `redesign-existing-projects` 显式触发示例。
  - 补充项目级 skill 的本地检查和排查命令。

## 验证结果

已执行：

```bash
git status --short
rg -n "项目级 Skills|\\.agents/skills|skills-lock.json|redesign-existing-projects|AGENTS.md.*注册" codex-skills.md docs/changes/2026-06-01-project-level-skills.md
```

结果：

- `git status --short` 显示本次改动包含 `codex-skills.md` 和新增本变更记录。
- 关键词检查通过，确认项目级 skill、`.agents/skills/`、`skills-lock.json`、`redesign-existing-projects` 示例和 `AGENTS.md` 关系说明已写入。
- 占位符和严格密钥形态检查未发现问题。

## 后续建议

如果后续在真实项目中长期使用项目级 skills，可再补充一份项目级 skill 创建模板，明确 `SKILL.md` frontmatter、description 写法、references 分层和 `AGENTS.md` 短规则示例。
