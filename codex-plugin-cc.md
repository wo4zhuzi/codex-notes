# codex-plugin-cc 在 Claude Code 中使用

## 结论

`codex-plugin-cc` 是 OpenAI 提供的 Claude Code 插件，用于在 Claude Code 会话中调用本机 Codex。它适合把 Claude Code 作为主工作台，让 Claude Code 承载较长上下文、需求梳理和连续实现，再让 Codex 作为独立协作方做审查、挑战方案或后台救援。

这个插件不是把 Codex 模型直接内置到 Claude Code，也不是替代 Claude Code 的执行环境。它依赖本机已安装并登录的 `codex` CLI，并复用当前 Codex 配置、权限和仓库环境。

推荐使用方式：

- Claude Code：主控任务、读取长上下文、整理需求、执行主要修改。
- Codex：通过插件命令做只读审查、对抗式审查、卡住时的后台排查。
- 人：决定是否采纳 Codex 结果，尤其是涉及架构、权限、安全和数据一致性的问题。

## 为什么用插件协作

Claude Code 和 Codex 都能独立完成编程任务，但插件协作的价值不是让两个 agent 同时乱改代码，而是建立一个更稳定的主从分工：

- Claude Code 保持主上下文，避免需求、历史讨论、长日志和当前执行状态分散在多个工具中。
- Codex 从另一个模型和运行时视角审查当前 diff 或方案，降低单一 agent 自证正确的风险。
- 后台任务可以把排查工作拆出去，Claude Code 不必在一个长会话里阻塞等待所有探索结果。

因此，`codex-plugin-cc` 更适合高风险变更、复杂 bugfix、发布前审查和方案不确定的任务，不适合把同一批文件同时交给两个 agent 并行修改。

## 安装与初始化

在 Claude Code 中添加插件市场：

```text
/plugin marketplace add openai/codex-plugin-cc
```

安装插件：

```text
/plugin install codex@openai-codex
```

重新加载插件：

```text
/reload-plugins
```

检查本机 Codex 环境：

```text
/codex:setup
```

如果本机还没有安装 Codex CLI，可以先安装并登录：

```bash
npm install -g @openai/codex
codex login
```

也可以在 Claude Code 中用 shell 命令触发登录：

```text
!codex login
```

初始化完成后，建议在目标仓库中先运行一次 `/codex:setup`，确认插件能找到 `codex`、当前账号已登录，并且 Codex 能访问当前工作区。

## 常用命令

### `/codex:review`

用于让 Codex 对当前改动做只读代码审查。适合提交前检查、PR 前自审、修复完成后的二次确认。

```text
/codex:review --background
```

常见提问方式：

```text
/codex:review --background

请重点检查：
1. 当前 diff 是否引入数据一致性风险。
2. 测试是否覆盖失败分支和边界输入。
3. 是否存在与 AGENTS.md 冲突的流程问题。
```

使用建议：

- 审查前先让 Claude Code 跑完关键测试。
- 审查时明确关注点，不要只说“帮我看看”。
- Codex 返回的问题要回到真实代码、测试和日志中验证，不要直接全盘接受。

### `/codex:adversarial-review`

用于让 Codex 从更挑剔的角度挑战方案。它适合在实现前或高风险改动后使用，重点不是找格式问题，而是检查根因、假设和风险边界。

```text
/codex:adversarial-review --background
```

示例：

```text
/codex:adversarial-review --background

Claude Code 当前判断缓存失效逻辑是根因，并计划只修改订单详情查询路径。
请从反方视角检查：
1. 根因证据是否充分。
2. 是否还有并发、事务、兼容性或降级路径遗漏。
3. 这个修复是否可能掩盖更上游的数据写入问题。
4. 最小验证集应该包含哪些测试或手动检查。
```

使用建议：

- 在方案进入实现前使用，成本最低。
- 对架构调整、权限、安全、支付、数据迁移类任务优先使用。
- 结果应该转化为计划修订或测试补充，而不是让两个工具互相辩论。

### `/codex:rescue`

用于在 Claude Code 卡住时，把排查任务交给 Codex 后台执行。适合测试连续失败、根因不明、日志很多、需要另一条探索路径的情况。

```text
/codex:rescue --background
```

示例：

```text
/codex:rescue --background

当前 bugfix 卡在 go test ./... 失败。
失败命令：
go test ./...

已知现象：
- TestOrderStatusAfterCancel 偶现失败。
- Claude Code 已检查订单状态机，但还没有确认是否与缓存有关。

限制：
- 先只做根因定位和建议，不要修改文件。
- 请输出最可能根因、证据、建议验证命令和最小修复方向。
```

使用建议：

- 交给 Codex 的输入要包含失败命令、错误摘要、已尝试方案和限制条件。
- 默认让 Codex 先只读定位；需要写入时再显式授权。
- Claude Code 读取结果后，应先验证证据，再决定是否修改。

### 后台任务管理

查看后台任务：

```text
/codex:status
```

读取任务结果：

```text
/codex:result
```

取消任务：

```text
/codex:cancel
```

后台任务适合审查和排查，但不应该变成无人值守的自动合并流程。关键结论仍要由 Claude Code 或开发者回到当前 diff 中核对。

## 复杂 Bugfix 协作流程

下面以“订单状态偶现错误”为例，展示 Claude Code 和 Codex 如何通过插件协作。

### 1. Claude Code 主持上下文

先让 Claude Code 读取需求、报错、日志、测试和相关代码，形成初步判断：

```text
请先定位订单状态偶现错误的根因。

要求：
1. 读取 AGENTS.md 和相关测试。
2. 先运行或说明需要运行的失败命令。
3. 修改前给出根因证据、最小修复计划和验证命令。
4. 不要先大范围重构。
```

Claude Code 的优势在这里体现为：持续承载较长上下文，把历史说明、日志片段、当前 diff、计划和验证结果串起来，作为主线推进任务。

### 2. Codex 挑战根因和方案

当 Claude Code 形成初步方案后，不急着实现或提交，先让 Codex 做对抗式审查：

```text
/codex:adversarial-review --background

当前方案：
Claude Code 判断订单取消后的缓存失效时机不一致，计划在 CancelOrder 成功提交后统一刷新订单详情缓存。

请只读审查这个方案：
1. 根因证据是否足够支持该修复。
2. 是否存在事务提交前刷新缓存导致脏读的风险。
3. 是否还有列表页、详情页、通知任务等其他读取路径会绕过该修复。
4. 应补哪些测试来证明修复有效。
```

这一步的目标是让 Codex 找反例。推荐把 Codex 输出归类为三类：

- 必须处理：能对应到代码、测试或日志证据的问题。
- 需要验证：合理但缺少证据的风险。
- 暂不处理：超出本次 bugfix 范围的重构建议。

### 3. Claude Code 执行最小修复

Claude Code 根据审查结果修订计划，再实现最小修复。执行时仍保持单一写入方，避免 Claude Code 和 Codex 同时修改同一批文件。

推荐提示词：

```text
请基于刚才 Codex 的审查结果修订计划并执行。

要求：
1. 只采纳有代码或测试证据支持的问题。
2. 先补失败测试，再实现修复。
3. 不做无关重构。
4. 修复后运行指定测试和 go test ./...。
5. 更新 docs/changes/YYYY-MM-DD-<topic>.md。
```

### 4. 卡住时交给 Codex 救援

如果修复后测试仍失败，可以把当前失败状态交给 Codex 后台排查：

```text
/codex:rescue --background

当前修复后仍失败。

失败命令：
go test ./...

失败摘要：
TestOrderStatusAfterCancel 期望 canceled，实际 paid。

已尝试：
1. 在 CancelOrder 提交后刷新详情缓存。
2. 补充详情页读取测试。

请只读定位：
1. 最可能根因。
2. 需要查看的关键代码路径。
3. 最小复现或验证命令。
4. 不要修改文件。
```

Claude Code 读取 `/codex:result` 后，应把 Codex 的结论当作审查材料，而不是自动执行脚本。真正进入代码修改前，仍要确认根因证据。

### 5. Codex 做最终只读审查

修复完成并通过测试后，再让 Codex 做最终审查：

```text
/codex:review --background

请审查当前 diff。
重点检查：
1. 是否只解决订单状态偶现错误，没有引入无关重构。
2. 新增测试是否能在修复前失败、修复后通过。
3. 是否遗漏缓存、事务、并发或兼容性风险。
4. docs/changes 记录是否能说明根因、改动和验证结果。
```

最终是否提交仍由开发者决定。插件审查不能替代本地测试、人工判断和 Git 授权流程。

## 推荐协作边界

推荐方案：

```text
Claude Code 主执行
-> Codex 插件只读审查或后台救援
-> Claude Code 汇总采纳
-> 本地验证
-> 人确认提交
```

不推荐：

- 两个 agent 同时修改同一批核心文件。
- 在根因不清楚时直接让 Codex `rescue` 写代码。
- 把 Codex 审查结果当成无需验证的事实。
- 只依赖插件审查，不运行项目测试。

如果任务涉及权限、安全、支付、数据迁移或兼容性，应先写 Spec 或计划，再用 `/codex:adversarial-review` 挑战方案，最后由 Claude Code 或开发者收敛到一个可验证的实现路径。

## 排查清单

如果 `/codex:*` 命令不可用，按顺序检查：

1. Claude Code 是否已经添加插件市场并安装 `codex@openai-codex`。
2. 是否执行过 `/reload-plugins` 或重启 Claude Code。
3. 本机是否能直接运行 `codex --help`。
4. 本机 Codex 是否已经 `codex login`。
5. 当前仓库是否允许 Codex 读取，并且没有被权限策略阻断。
6. `/codex:setup` 输出是否提示缺少 Codex app server、登录状态或配置。

如果 Codex 后台任务没有结果，先用 `/codex:status` 确认任务是否仍在运行，再用 `/codex:result` 读取结果；不需要的任务用 `/codex:cancel` 取消。

## 与现有协作文档的关系

本文只讲 `codex-plugin-cc` 在 Claude Code 中的插件化使用。跨工具规则文件、`AGENTS.md` 与 `CLAUDE.md` 的分工、文件型计划交接和冲突处理，见 [Codex 与 Claude Code 协作](./codex-claude-collaboration.md)。

## 来源与更新时间

本文更新时间：2026-06-03。

参考资料：

- [openai/codex-plugin-cc](https://github.com/openai/codex-plugin-cc)
- [Codex CLI 使用笔记](./codex-cli.md)
- [Codex 与 Claude Code 协作](./codex-claude-collaboration.md)
