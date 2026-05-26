# 创建 agent-evaluation 个人 Skill

## 任务背景

用户要求使用 `skill-creator` 创建个人 Codex skill：`agent-evaluation`，用于在 AI Agent 完成任务前，根据任务阶段、项目风险和改动类型选择合适的评估标准，判断当前结果是可交付、部分完成，还是必须停止交给人。

## 根因定位

本地 `~/.codex/skills/agent-evaluation/` 目录原本不存在，缺少可触发的 `SKILL.md` 和按场景拆分的评估 references。仓库内已有 `docs/agent-evaluation-references/` 参考资料，可作为生成该个人 skill 的内容依据。

## 执行计划

1. 检查 `git status --short`，确认仓库当前改动状态。
2. 读取 `skill-creator` 使用约束和 `docs/agent-evaluation-references/` 现有参考文件。
3. 在 `~/.codex/skills/agent-evaluation/` 下创建 `SKILL.md` 和 8 个 `references/*.md` 文件。
4. 校验文件结构、frontmatter、触发描述和敏感信息扫描结果。
5. 按仓库规则生成本日期变更文档。

## 变更内容

新增个人 skill 文件：

- `~/.codex/skills/agent-evaluation/SKILL.md`
- `~/.codex/skills/agent-evaluation/references/docs.md`
- `~/.codex/skills/agent-evaluation/references/bugfix.md`
- `~/.codex/skills/agent-evaluation/references/feature.md`
- `~/.codex/skills/agent-evaluation/references/refactor.md`
- `~/.codex/skills/agent-evaluation/references/config.md`
- `~/.codex/skills/agent-evaluation/references/tool-integration.md`
- `~/.codex/skills/agent-evaluation/references/code-review.md`
- `~/.codex/skills/agent-evaluation/references/high-risk.md`

新增仓库变更记录：

- `docs/changes/2026-05-26-agent-evaluation-skill.md`

## 验证结果

已执行：

```bash
find /Users/kongdebo/.codex/skills/agent-evaluation -type f | sort
sed -n '1,120p' /Users/kongdebo/.codex/skills/agent-evaluation/SKILL.md
rg -n "name: agent-evaluation|验收|完成前检查|评估 Agent 输出|判断是否可交付|检查验证证据|区分阶段性评估和交付评估" /Users/kongdebo/.codex/skills/agent-evaluation/SKILL.md
rg -n "README|INSTALLATION|QUICK_REFERENCE|CHANGELOG" /Users/kongdebo/.codex/skills/agent-evaluation
rg -n "TODO|FIXME|your-api-key|sk-" /Users/kongdebo/.codex/skills/agent-evaluation
```

结果：

- 目标目录包含用户要求的 1 个 `SKILL.md` 和 8 个 references 文件。
- `SKILL.md` frontmatter 的 `name` 为 `agent-evaluation`，`description` 包含用户要求的触发场景。
- 未创建 README、安装说明或其他额外无关文件。
- 敏感信息扫描只命中文档中的示例扫描命令文本，不包含真实密钥、Token、私有代理地址或内部服务 URL。

## 后续建议

后续使用该 skill 时，可用“完成前检查”“评估是否可交付”“检查验证证据”等提示语触发，并根据任务类型按需读取对应 reference 文件。
