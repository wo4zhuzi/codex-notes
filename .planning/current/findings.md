# 调查发现

## 仓库状态

- 调查开始时 `git status --short` 无输出，工作区干净。
- 仓库内没有自定义 Superpower Skill 目录。
- `codex-skills.md` 记录了来自 `obra/superpowers` 的 `brainstorming`、`writing-plans`、`systematic-debugging`、`verification-before-completion`。
- 当前实际加载的相关 Skill 位于用户级 `~/.codex/skills/`，尚未确定用户希望修改单个 Skill 还是整套 Superpowers 规则。
- `~/.codex/skills/` 中单独安装了四个来自 Superpowers 的 Skill：`brainstorming`、`writing-plans`、`systematic-debugging`、`verification-before-completion`；没有发现完整 Superpowers 插件目录。
- `codex-skills.md` 当前明确推荐这四个 Skill 作为轻量组合，并说明完整插件会更积极触发规划、TDD、review、worktree 和分支收尾。

## 初步冲突证据

- 当前 `brainstorming` 含硬门槛：任何行为修改都必须先逐项澄清、提出多个方案、取得设计确认、写设计文档、等待再次确认，再进入实施计划。
- 该规则可能与仓库自身 `/plan`、Spec、ReAct 和日期变更文档工作流叠加，形成重复规划与重复确认。
- 用户级 `brainstorming` 的硬门槛覆盖“任何行为修改”，即使任务很小也要求多轮问题、2-3 个方案、设计落盘、提交设计、再次等待用户审阅，再调用 `writing-plans`。这与更高层仓库工作流构成确定性的控制协议叠加。
- `writing-plans` 强制把每一步压缩到 2-5 分钟、展示完整代码、频繁 commit，并要求从不存在于当前会话的 `subagent-driven-development` 或 `executing-plans` 二选一执行。这既与仓库“未经用户确认不得 commit”冲突，也与当前禁止自行委派子 agent 的约束冲突。
- `systematic-debugging` 的“先证据、后根因、再最小修复”与 GPT-5.6 和本仓库 ReAct 兼容；问题主要是它强制引用当前未安装的 `test-driven-development`，应降级为“可用且任务需要时调用”。
- `verification-before-completion` 只要求新鲜验证证据，和本仓库目标一致，原则上无需重写。

## 推荐维护方式

- 仓库已有约定：个人组合 Skill 的唯一维护源放在 `my-skills/my-*`，再用 `my-skills/install-skill.sh` 安装到 `~/.codex/skills/`。
- 直接修改第三方安装目录最快，但后续重新安装或升级会覆盖修改，且改动无法随仓库审查和恢复。
- 创建仓库托管的兼容版个人 Skill 更稳定，但必须同步停用或收窄旧 `brainstorming`、`writing-plans` 的触发，否则冲突协议仍会同时加载。

## 根因结论草案

- 模型升级没有让 Skill 文件格式失效；冲突发生在“决策控制权”。
- 旧 Skill 为能力较弱、容易跳步的模型设计了高密度硬门槛，以强制顺序换稳定性。
- 更强模型能从仓库、风险、工具反馈中动态决定流程，继续叠加无条件门槛会造成重复规划、重复确认、无意义文档与 commit、调用不存在的 Skill，并压制模型根据上下文缩放流程的能力。
- 推荐从“无条件流程脚本”改成“风险分级约束”：保留根因、验证、安全和授权边界；把方案数量、确认次数、TDD、子 agent、commit 和文档粒度改为按风险触发。

## 事实边界

- “GPT-5.6 更聪明”可作为用户观察和设计前提，但不能据此声称特定模型内部机制。
- 应将冲突解释为提示协议与工作流控制层的冲突，而不是模型能力本身与 Markdown Skill 格式不兼容。
