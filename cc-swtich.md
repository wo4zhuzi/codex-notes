# CC-Switch

## 文档

- [CC-Switch 中文 README](https://github.com/farion1231/cc-switch/blob/main/README_ZH.md)

## 安装

### macOS 用户

#### 方式一：通过 Homebrew 安装（推荐）

```shell
brew tap farion1231/ccswitch
brew install --cask cc-switch
```

更新：

```shell
brew upgrade --cask cc-switch
```

#### 方式二：手动下载安装

从 Releases 页面下载以下安装包之一：

- `CC-Switch-v{版本号}-macOS.dmg`（推荐）
- `CC-Switch-v{版本号}-macOS.zip`

注意：CC Switch macOS 版本已通过 Apple 代码签名和公证，可直接安装打开。

## 实现原理

CC-Switch 通过修改配置文件实现 Codex、Claude Code 的服务商切换。

## 配置文件参考

以下配置可作为 CC-Switch 的配置参考：

```toml
# =========================
# 核心输出控制（工程规范层）
# =========================

developer_instructions = """
请使用中文回答。

回答要求：
1. 优先提供工程可落地的实现方案（可直接用于生产环境）。
2. 涉及代码必须保证完整性，避免仅给思路或片段代码。
3. 修改代码遵循最小变更原则，避免无意义重构或结构重排。
4. 必须先定位问题根因，再给出解决方案，不允许直接猜测结论。
5. 多方案必须提供推荐方案，并说明取舍依据。
6. 输出结构必须清晰：结论优先，其次关键依据，最后细节说明。
7. 避免冗余解释与概念扩展，保持工程导向表达。

文档语言要求：
1. 生成或修改仓库文档、配置说明、README、AGENTS.md 时，默认使用简体中文；除非用户明确要求英文，否则不要生成英文文档。
2. 执行 /init 生成 AGENTS.md 时，标题、章节名、正文和示例说明均使用简体中文。
"""

# 用途：
# 控制模型输出风格与工程约束能力

# 当前策略：
# - 强工程约束
# - 强可执行性
# - 偏后端工程实践

# 适用场景：
# - 后端服务开发
# - API / RPC / 中间件
# - 高并发系统设计
# - 数据处理 / 系统调试

# 可调整场景：
# - 学习 / 概念理解：可简化指令
# - PoC / demo：降低约束


# =========================
# 模型选择（审查模型）
# =========================

review_model = "gpt-5.5"

# 作用：
# 用于代码审查 / 二次校验 / 逻辑检查阶段

# 当前选择理由：
# - 逻辑一致性较强
# - 更适合做工程质量检查
# - 稳定性优于速度

# 可调整策略：
# - 高吞吐场景：使用更快模型
# - 高风险逻辑审查：使用更强模型


# =========================
# 推理强度控制（核心性能开关）
# =========================

model_reasoning_effort = "high"

# 可选等级说明：

# low：
# - 快速响应模式
# - 适用于简单任务
# - 如：脚本 / 文档 / 格式转换

# high：
# - 默认工程开发模式
# - 适用于大多数工程开发场景

# exhigh：
# - 深度分析模式
# - 用于复杂系统问题定位与设计

# 当前策略：
# 默认使用 high（token 充足时优先保证推理质量）

# 适用场景切换建议：

# 使用 exhigh：
# - 复杂问题排查（多模块联动）
# - 性能瓶颈分析
# - 架构设计评审
# - 难以复现的线上问题

# 使用 low：
# - CLI 工具类任务
# - 简单代码生成
# - 批量文本处理


# =========================
# 响应存储控制
# =========================

disable_response_storage = true

# 作用：
# 控制是否保留模型输出记录

# true（推荐）：
# - 避免上下文污染
# - 保持多项目隔离
# - 提升安全性

# false：
# - 允许历史追溯
# - 可能带来上下文干扰

# 建议使用场景：
# - 多系统并行开发：推荐保持 true
# - 单系统长期演进：可考虑 false


# =========================
# 输出风格控制
# =========================

personality = "pragmatic"

# 可选风格：

# pragmatic：
# - 工程导向（推荐）
# - 强执行力输出

# creative：
# - 偏设计与发散
# - 用于方案讨论

# academic：
# - 偏理论解释
# - 用于学习与文档

# chatty：
# - 对话风格
# - 不适用于工程开发

# 当前选择：
# pragmatic（后端工程最优风格）


# =========================
# 版本映射规则
# =========================

[notice.model_migrations]
"gpt-5.2" = "gpt-5.2-codex"

# 作用：
# - 旧模型名称自动映射
# - 保证兼容性
# - 避免配置失效

# 一般无需调整


# =========================
# 功能开关（运行时行为）
# =========================

[features]

fast_mode = false

# 作用：
# 控制响应速度优先级

# false（推荐）：
# - 稳定优先
# - 更适合调试与生产问题分析

# true：
# - 响应更快
# - 适合轻量任务或交互式操作

# 建议：
# 默认关闭 fast_mode


responses_websockets_v2 = false

# 作用：
# 底层通信协议版本控制

# 当前：
# 保持关闭（稳定优先）

# 一般无需修改


memories = false

# 作用：
# 是否允许跨会话记忆

# false（推荐）：
# - 保证不同项目隔离
# - 避免上下文污染
# - 提高可预测性

# true：
# - 适合单系统长期演进项目

# 不建议在多项目环境开启


multi_agent = false

# 作用：
# 是否启用多模型协作分析

# false（推荐默认）：
# - 输出直接
# - 更高执行效率
# - 更适合工程开发

# true：
# - 多视角分析
# - 更适合方案设计与评审

# 使用建议：
# - 开发 / debug：false
# - 架构设计 / 技术选型：true


# =========================
# 记忆子系统（高级功能）
# =========================

[memories]

consolidation_model = "gpt-5.4"
extract_model = "gpt-5.4"

# 作用：
# - extract_model：提取关键信息
# - consolidation_model：压缩与整理记忆

# 当前选择：
# 使用较强模型保证信息准确性

# 一般无需调整
# 除非成本优化或模型升级


max_raw_memories_for_consolidation = 512

# 作用：
# 单次记忆整理最大输入量

# 当前设置：
# 平衡稳定性与性能

# 调整建议：
# - 大规模系统：可调大
# - 轻量使用：可调小


max_unused_days = 30

# 作用：
# 记忆过期时间（未使用）

# 当前：
# 30 天（标准工程周期）

# 调整场景：
# - 高频系统：缩短
# - 长期项目：延长


max_rollout_age_days = 45

# 作用：
# 记忆生命周期上限

# 当前：
# 偏保守设置（安全优先）

# 可调：
# - 研究项目：可延长
# - 临时项目：可缩短
```

## 多种配置参考

多套配置已拆分到 [cc-switch-configs](./cc-switch-configs/) 目录，便于单独维护和复制使用。

注意：这些文件是整份配置模板，不要在模板顶层添加 `profile = "..."`。Codex 的 `--profile` 会从 `config.toml` 中查找已定义的 `[profiles.<name>]`，如果顶层 `profile` 指向未定义名称，会出现 `Error loading configuration: config profile <name> not found`。按文件模板切换时，应让 CC-Switch 覆盖整份配置，或使用 Codex 的 `--profile-v2 <name>` 加载独立配置文件。

| 配置文件 | 场景 | 推荐用途 |
| --- | --- | --- |
| [dev-main.toml](./cc-switch-configs/dev-main.toml) | 日常开发配置 | 常规需求开发、Bug 修复、接口联调 |
| [dev-arch.toml](./cc-switch-configs/dev-arch.toml) | 架构设计与复杂问题分析配置 | 架构评审、技术选型、高风险变更评估 |
| [dev-cheap.toml](./cc-switch-configs/dev-cheap.toml) | 低成本快速处理配置 | 文档整理、格式转换、简单脚本 |
| [dev-debug.toml](./cc-switch-configs/dev-debug.toml) | 问题排查与调试配置 | 线上问题定位、日志排查、复杂 Bug 修复 |
| [dev-review.toml](./cc-switch-configs/dev-review.toml) | 代码审查配置 | PR 审查、重构评估、测试覆盖检查 |
| [dev-subagent-review.toml](./cc-switch-configs/dev-subagent-review.toml) | Subagent 分支审查配置 | 只读并行审查 PR / 分支风险 |
| [dev-subagent-bugfix.toml](./cc-switch-configs/dev-subagent-bugfix.toml) | Subagent 复杂 Bugfix 配置 | 并行定位根因，主 agent 串行修复 |
| [dev-subagent-project.toml](./cc-switch-configs/dev-subagent-project.toml) | Subagent 新项目开发配置 | 先 Spec/Plan，再做架构、测试、安全专项审查 |

推荐默认使用 `dev-main.toml`。当任务复杂度明显升高时切换到 `dev-arch.toml` 或 `dev-debug.toml`；当任务低风险且追求响应速度时使用 `dev-cheap.toml`。

## Subagent 场景配置

Subagent 场景建议单独配置，而不是直接改动日常开发配置。原因是 subagent 会放大并发工具调用、token 消耗和上下文汇总成本，适合在审查、复杂排查和新项目设计阶段显式启用。

推荐新增三类 CC-Switch 配置：

- `dev-subagent-review.toml`：用于分支审查和 PR 审查。Subagent 默认只读，负责并行取证和专项初判；主 agent 负责去重、排序和最终审查结论。
- `dev-subagent-bugfix.toml`：用于复杂 Bugfix。Subagent 只负责日志、错误、调用链、测试失败和影响面分析；修复动作由主 agent 串行执行。
- `dev-subagent-project.toml`：用于新项目或新模块。先进入 Spec/Plan，再让 subagent 分别审查架构边界、测试策略、安全风险和文档/API 依据。

具体 custom subagent 模板放在：

```text
cc-switch-configs/subagents/
```

这个目录是仓库内的模板源，方便复制和维护；Codex 不会自动读取这里的文件。实际启用时，需要复制到全局或项目级 agents 目录。

三类配置的共同核心开关：

```toml
[features]
fast_mode = false
responses_websockets_v2 = false
memories = false
multi_agent = true

[agents]
max_threads = 6
max_depth = 1
```

配置取舍：

- `multi_agent = true`：启用 Codex 多 agent 能力。
- `max_threads = 4` 到 `6`：限制并发线程，避免审查结果过多导致主 agent 难以收口。
- `max_depth = 1`：只允许主 agent 派生一层 subagent，避免递归分派失控。
- `memories = false`：保持多项目隔离，避免跨项目上下文污染。

### Subagent 模板安装

全局安装适合通用审查角色，所有项目都能复用：

```bash
mkdir -p ~/.codex/agents
cp cc-switch-configs/subagents/*.toml ~/.codex/agents/
```

项目级安装适合绑定某个项目的业务规则、权限边界或领域术语：

```bash
mkdir -p .codex/agents
cp cc-switch-configs/subagents/*.toml .codex/agents/
```

推荐先使用全局安装；如果某个项目需要更严格约束，再把模板复制到项目级 `.codex/agents/` 并单独调整。

修改 `cc-switch-configs/subagents/` 下的模板后，需要重新复制到 `~/.codex/agents/` 或 `.codex/agents/`。如果当前 Codex 会话已经启动，建议重启会话，确保新 agent 配置被加载。

### 主 agent 与 subagent 分工

主 agent 是审查负责人和最终执行者：

- 明确任务目标、审查范围和验证方式。
- 决定派发哪些 subagent。
- 限定 subagent 只读或专项范围。
- 汇总 subagent 结果，去重、处理冲突结论并按严重级别排序。
- 在需要修改代码时串行执行修复，并运行最终验证。

Subagent 是专项审查员或取证员：

- `pr-explorer`：分析 diff、入口点、调用链和影响范围。
- `docs-researcher`：核对官方文档、API 行为、模型和工具版本说明。
- `risk-reviewer`：检查 correctness、数据一致性、并发和错误处理风险。
- `security-reviewer`：检查权限、密钥、注入、敏感日志和供应链配置。
- `compat-reviewer`：检查接口兼容、配置兼容、迁移和回滚风险。
- `test-impact-reviewer`：判断测试覆盖、回归路径和缺失用例。

### Subagent 推理强度建议

主 agent 建议使用：

```toml
model_reasoning_effort = "exhigh"
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

### 推荐口令

分支审查：

```text
请使用只读 subagent 审查当前分支相对 main 的改动，不要修改文件。让 pr-explorer 分析影响范围，risk-reviewer 检查行为风险，security-reviewer 检查安全问题，test-impact-reviewer 检查测试缺口，docs-researcher 核对外部 API 或框架行为。最后由主 agent 汇总 findings，按严重程度排序。
```

复杂 Bugfix：

```text
请使用只读 subagent 先定位这个问题的根因，不要修改文件。分别分析错误日志、调用链、相关测试失败和影响面。等所有 subagent 返回后，由主 agent 汇总根因判断、最小修复方案和验证计划，再开始修复。
```

新项目开发：

```text
请先进入 Spec/Plan，不要直接实现。使用 subagent 分别审查架构边界、测试策略、安全风险和外部文档依据。主 agent 汇总后给出目标、非目标、实施步骤、验证方式和风险边界。
```
