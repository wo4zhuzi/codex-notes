# Codex Subagent 使用笔记

Subagent 用于把一个复杂任务拆给多个专项 agent 并行取证、审查或验证。它适合 PR 审查、复杂 bug 根因定位、安全/兼容性检查、测试影响分析和外部文档核对。

## Subagent 是什么

主 agent 是任务负责人，负责明确目标、分派范围、汇总结论和最终执行修改。Subagent 是专项审查员或取证员，负责在限定范围内只读分析并输出证据、风险、置信度和建议。

常见分工：

- `pr-explorer`：分析 diff、入口点、调用链和影响范围。
- `docs-researcher`：核对官方文档、API 行为、模型和工具版本说明。
- `risk-reviewer`：检查 correctness、数据一致性、并发和错误处理风险。
- `security-reviewer`：检查权限、密钥、注入、敏感日志和供应链配置。
- `compat-reviewer`：检查接口兼容、配置兼容、迁移和回滚风险。
- `test-impact-reviewer`：判断测试覆盖、回归路径和缺失用例。

## 仓库内模板

本仓库的 custom subagent 模板源放在：

```text
cc-switch-configs/subagents/
```

这些文件方便复制和维护，但 Codex 不会自动读取仓库内的模板源。要让自定义 agent 真正按 TOML 生效，需要把模板复制到 Codex 的 agents 目录，并重启会话。

默认推理强度分工：

```text
pr-explorer              -> medium
docs-researcher          -> medium
test-impact-reviewer     -> high
risk-reviewer            -> high
security-reviewer        -> high
compat-reviewer          -> high
```

## 如何注册 agent

全局注册适合通用审查角色，所有项目都能复用：

```bash
mkdir -p ~/.codex/agents
cp cc-switch-configs/subagents/*.toml ~/.codex/agents/
```

项目级注册适合绑定某个项目的业务规则、权限边界或领域术语：

```bash
mkdir -p .codex/agents
cp cc-switch-configs/subagents/*.toml .codex/agents/
```

推荐先使用全局注册；如果某个项目需要更严格约束，再把模板复制到项目级 `.codex/agents/` 并单独调整。

修改 `cc-switch-configs/subagents/` 下的模板后，需要重新复制到 `~/.codex/agents/` 或 `.codex/agents/`。如果当前 Codex 会话已经启动，建议重启会话，确保新 agent 配置被加载。

## 如何确认注册是否生效

重启后，需要确认当前 Codex 暴露的 `spawn_agent` 可选类型里是否包含这些自定义 agent：

```text
pr-explorer
test-impact-reviewer
risk-reviewer
security-reviewer
compat-reviewer
docs-researcher
```

如果当前工具只暴露：

```text
default
explorer
worker
```

说明这些 TOML 没有被本次会话注册成可调用 `agent_type`。此时即使提示词里写“让 pr-explorer 分析影响范围”，也只是让通用 subagent 扮演这个角色，不会自动读取 `pr-explorer.toml`。

## 如何调用 agent

如果自定义 agent 已注册，主 agent 应直接创建对应类型：

```text
agent_type = "pr-explorer"
agent_type = "test-impact-reviewer"
agent_type = "risk-reviewer"
```

这种情况下，才应按对应 TOML 中的 `model_reasoning_effort` 和 `sandbox_mode` 生效。

如果自定义 agent 未注册，只能降级使用通用类型：

```text
agent_type = "explorer"
```

再在提示词中限定“你是 pr-explorer / risk-reviewer，只读审查并输出证据”。这种 fallback 方式不会自动使用 `cc-switch-configs/subagents/*.toml`，且未显式覆盖时通常会继承主 agent 的模型和推理强度。比如主 agent 是 `xhigh` 时，通用 `explorer` subagent 也可能继承 `xhigh`。

一次有效启用应同时满足：

- `~/.codex/agents/` 或 `.codex/agents/` 中存在对应 TOML。
- 当前 Codex 会话已重启。
- `spawn_agent` 可选类型中能看到自定义 agent 名称。
- 创建时使用的是 `agent_type = "pr-explorer"` 这类自定义名称，而不是通用 `explorer`。
- Subagent 推理强度符合 TOML：取证类 `medium`，风险审查类 `high`，主 agent `xhigh`。

## 推理强度与只读约束

主 agent 建议使用：

```toml
model_reasoning_effort = "xhigh"
```

原因是主 agent 需要做最终裁决：合并重复 findings、处理互相矛盾的结论、判断严重级别和决定是否进入修复。

取证类 subagent 建议使用：

```toml
model_reasoning_effort = "medium"
```

适用于 `pr-explorer` 和 `docs-researcher`。这类 agent 主要负责枚举、查证和归纳，`medium` 通常足够，且能控制并发成本。

高风险审查类 subagent 建议使用：

```toml
model_reasoning_effort = "high"
```

适用于 `risk-reviewer`、`security-reviewer`、`compat-reviewer` 和 `test-impact-reviewer`。这类 agent 需要判断风险是否成立，不只是收集证据，因此应提高推理强度。

不建议把所有 subagent 默认设为 `high`。这样会增加延迟、费用和长篇分析噪声，主 agent 汇总时也更容易被低价值细节干扰。

审查类 subagent 默认应使用只读约束：

```toml
sandbox_mode = "read-only"
```

如果需要修改文件，应由主 agent 在汇总结论后串行执行，避免多个 agent 并行写入造成冲突或遗漏验证。

## 推荐口令

分支审查：

```text
请使用只读 subagent 审查当前分支相对 main 的改动，不要修改文件。优先调用已注册的 pr-explorer 分析影响范围，risk-reviewer 检查行为风险，security-reviewer 检查安全问题，test-impact-reviewer 检查测试缺口，docs-researcher 核对外部 API 或框架行为。如果当前运行时没有暴露这些自定义 agent_type，请明确说明并使用通用只读 explorer fallback，不要假装 TOML 配置已生效。最后由主 agent 汇总 findings，按严重程度排序。
```

复杂 Bugfix：

```text
请使用只读 subagent 先定位这个问题的根因，不要修改文件。优先调用已注册的 pr-explorer 分析调用链和影响面，test-impact-reviewer 分析相关测试失败和验证缺口，risk-reviewer 判断行为风险。如果当前运行时没有暴露这些自定义 agent_type，请明确说明并使用通用只读 explorer fallback，不要假装 TOML 配置已生效。等所有 subagent 返回后，由主 agent 汇总根因判断、最小修复方案和验证计划，再开始修复。
```

新项目开发：

```text
请先进入 Spec/Plan，不要直接实现。使用 subagent 分别审查架构边界、测试策略、安全风险和外部文档依据。主 agent 汇总后给出目标、非目标、实施步骤、验证方式和风险边界。
```
