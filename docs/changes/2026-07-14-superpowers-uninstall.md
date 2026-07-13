# 卸载 Superpowers Skills

## 任务背景

本机原先单独安装了 `brainstorming`、`writing-plans`、`systematic-debugging` 和 `verification-before-completion` 四个来自 `obra/superpowers` 的 Skill。模型升级且仓库工作流逐步完善后，这些 Skill 的强制流程与现有协作约束出现重复和冲突。用户最终决定直接卸载，不保留备份。

## 根因定位

问题不是 Codex 无法读取 Skill，也没有证据表明 GPT-5.6 与 Superpowers 在文件格式或加载机制上存在官方定义的不兼容。实际冲突发生在工作流控制层：

- `brainstorming` 对简单改动也强制多轮澄清、多个方案、设计文档、提交和再次确认，与仓库已有 Plan / Spec 流程重复。
- `writing-plans` 强制频繁 commit，并要求调用本机不一定存在的执行或子 agent Skill，与仓库提交授权边界冲突。
- 仓库已经通过 `AGENTS.md`、Plan / Spec、ReAct、测试和变更记录约束根因定位、执行与验证，再叠加相同职责会产生重复规划和过量流程。
- 更强模型可以根据任务风险、仓库约束和实时工具反馈调整流程粒度；无条件硬门槛会限制这种动态判断能力。

因此，本次将问题定义为“旧 Skill 的固定流程与当前模型及仓库工作流不匹配”，不将其表述为 OpenAI 官方发布的兼容性结论。

## 执行计划

1. 确认本机实际安装的 Superpowers Skill 目录。
2. 按用户授权直接删除四个目录，不创建备份。
3. 更新 `codex-skills.md` 的当前安装状态、卸载原因和未来重装方式。
4. 检查卸载结果、文档内容、敏感信息和 Git 差异。

## 变更内容

- 删除用户级 `~/.codex/skills/brainstorming/`。
- 删除用户级 `~/.codex/skills/writing-plans/`。
- 删除用户级 `~/.codex/skills/systematic-debugging/`。
- 删除用户级 `~/.codex/skills/verification-before-completion/`。
- 更新 `codex-skills.md`，不再把上述 Skill 列为当前安装项或默认推荐组合。
- 保留重新安装命令，便于后续按需恢复。

## 验证结果

- 删除命令退出码为 0。
- 重新列出 `~/.codex/skills/` 后，四个目标目录均不存在。
- 其他个人 Skill、`planning-with-files-zh` 和系统 Skill 未被删除。
- `find ~/.codex/skills ...` 精确检查退出码为 0 且无输出，确认四个目标目录未残留。
- `git diff --check` 退出码为 0，未发现空白错误。
- 仓库级 `rg -n "TODO|FIXME|your-api-key|sk-" .` 只命中既有检查命令和普通示例文本；对本次文件执行精确正则后无匹配，未发现占位符或真实密钥形态。
- 本机没有安装 `markdownlint`；已人工检查本次 Markdown diff，未新增失效的本地链接。

## 后续建议

- 重启 Codex 或开启新会话，使当前会话的 Skill 列表刷新。
- 后续如重新安装，先检查上游 Skill 的提交、TDD、文档和子 agent 约束是否与目标仓库兼容。
