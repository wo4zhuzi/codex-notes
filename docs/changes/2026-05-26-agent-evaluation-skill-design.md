# 2026-05-26 Agent Evaluation 多场景 Skill 设计

## 任务背景

用户希望将 `agent-evaluation` 多场景 skill 的作用、边界和创建方式记录到 skills 相关章节中，便于后续按“一个 skill + 多个 reference”的方式创建全局 skill。

## 根因定位

当前仓库已有 `agent-evaluation.md` 说明评估方法，也有 `codex-skills.md` 说明 skills 安装与使用，但缺少把 Agent Evaluation 封装为 Codex skill 的具体结构说明。尤其是多场景评估不适合拆成多个独立 skill，需要记录“`SKILL.md` 总入口 + `references/` 场景标准”的组织方式。

## 执行计划

- 在 `codex-skills.md` 的“分场景推荐”后新增 `agent-evaluation` 多场景 skill 设计说明。
- 记录 skill 作用、职责边界、推荐目录结构、场景边界表、创建提示词和使用示例。
- 新增本次变更记录。
- 运行仓库约定的本地检查命令。

## 变更内容

- 更新 `codex-skills.md`：
  - 新增 `agent-evaluation` 多场景 skill 设计章节。
  - 明确它不替代测试、`/diff` 或 `/review`。
  - 记录推荐目录结构和场景 reference 文件。
  - 增加文档、bugfix、新功能、重构、配置、工具集成、code review、高风险生产变更的边界表。
  - 增加使用 `skill-creator` 创建该 skill 的完整提示词。
  - 增加阶段性评估、交付评估和指定场景评估的使用示例。

## 验证结果

- 已执行 `git status --short`：确认本次改动包含 `codex-skills.md` 和本变更记录。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例，以及 `git-workflow.md` 中的示例分支名 `project-task-branch`，未发现真实密钥或遗留调试标记。
- 已执行 `rg -n "agent-evaluation|references|skill-creator|阶段性评估|交付评估" codex-skills.md`：确认新增章节、目录结构、创建提示词和使用示例已写入。
- 已手动检查 `codex-skills.md` 的当前已安装 skills 表：未将 `agent-evaluation` 标记为已安装。

## 后续建议

- 后续真正创建全局 skill 时，优先使用 `skill-creator`，不要手写零散结构。
- 如创建完成，可再更新 `codex-skills.md` 的当前已安装 skills 表，避免提前标记为已安装。
