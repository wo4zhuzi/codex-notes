# 2026-05-27 新增 CC-Switch Subagent 场景配置

## 任务背景

用户希望把 Codex subagent 的使用方式落到 CC-Switch 配置中，覆盖三个场景：

- 分支审查落地方案。
- Bugfix 修复。
- 新项目开发。

本次目标是新增可复制的场景化 TOML 配置和 custom subagent 模板，并在 `cc-swtich.md` 中说明主 agent 与 subagent 的分工、推理强度取舍、安装方式和推荐口令。

## 根因定位

现有 `cc-switch-configs/` 已按日常开发、架构分析、低成本、调试、代码审查拆分配置，但主要维度是模型、推理强度和输出风格。

缺口在于：

- 没有专门用于 subagent 并行审查的配置模板。
- 复杂 Bugfix 场景没有明确“subagent 只读定位，主 agent 串行修复”的边界。
- 新项目开发场景没有明确“先 Spec/Plan，再由 subagent 做专项审查”的配置入口。
- 既有文档没有说明取证类 subagent 与高风险审查类 subagent 的推理强度差异。
- 仓库中没有集中保存 custom subagent TOML 模板的目录，导致“如何添加 subagent”只能靠口头说明，无法直接复制使用。

## 执行计划

1. 新增三份 CC-Switch TOML 配置：
   - `dev-subagent-review.toml`
   - `dev-subagent-bugfix.toml`
   - `dev-subagent-project.toml`
2. 新增 `cc-switch-configs/subagents/` 目录，保存 custom subagent 模板。
3. 更新 `cc-swtich.md` 的配置索引。
4. 补充 subagent 场景配置说明、模板安装方式、主/子 agent 分工、推理强度建议和推荐口令。
5. 执行 TOML 解析、文本搜索和安全占位符检查。

## 变更内容

- 新增 `cc-switch-configs/dev-subagent-review.toml`：
  - 开启 `multi_agent = true`。
  - 主 agent 使用 `model_reasoning_effort = "exhigh"`。
  - 约束 subagent 只读并行审查 PR / 分支风险。
- 新增 `cc-switch-configs/dev-subagent-bugfix.toml`：
  - 开启 `multi_agent = true`。
  - 设置 `max_threads = 4`，降低复杂排查场景的并发噪声。
  - 约束 subagent 只做根因定位和影响面分析，修复由主 agent 串行完成。
- 新增 `cc-switch-configs/dev-subagent-project.toml`：
  - 开启 `multi_agent = true`。
  - 要求先进入 Spec/Plan，再让 subagent 做架构、测试、安全和文档/API 审查。
- 更新 `cc-swtich.md`：
  - 配置表新增三类 subagent 场景模板。
  - 新增“Subagent 场景配置”章节。
  - 说明 `cc-switch-configs/subagents/` 是模板源，实际启用需复制到 `~/.codex/agents/` 或 `.codex/agents/`。
  - 补充全局安装和项目级安装命令。
  - 明确主 agent 负责最终裁决，subagent 负责专项取证和初判。
  - 说明取证类 subagent 推荐 `medium`，高风险审查类 subagent 推荐 `high`。
  - 补充分支审查、复杂 Bugfix、新项目开发三类推荐口令。
- 新增 `cc-switch-configs/subagents/`：
  - `pr-explorer.toml`：取证类 subagent，使用 `medium`。
  - `docs-researcher.toml`：取证类 subagent，使用 `medium`。
  - `risk-reviewer.toml`：高风险审查类 subagent，使用 `high`。
  - `security-reviewer.toml`：高风险审查类 subagent，使用 `high`。
  - `compat-reviewer.toml`：高风险审查类 subagent，使用 `high`。
  - `test-impact-reviewer.toml`：高风险审查类 subagent，使用 `high`。
  - 所有模板都设置 `sandbox_mode = "read-only"`。

## 验证结果

已执行：

```bash
git status --short
rg -n "dev-subagent|multi_agent|model_reasoning_effort|sandbox_mode" cc-swtich.md cc-switch-configs
rg -n "TODO|FIXME|your-api-key|sk-" .
find cc-switch-configs/subagents -maxdepth 1 -type f -name "*.toml" -print
python3 -c 'import pathlib,tomllib; [tomllib.load(open(p,"rb")) for p in pathlib.Path("cc-switch-configs").glob("**/*.toml")]; print("ok")'
git diff --check
```

结果：

- `git status --short`：仅显示本次修改 `cc-swtich.md`，以及新增三份 `dev-subagent-*` 配置、六份 `subagents/` 模板和本变更文档。
- `rg -n "dev-subagent|multi_agent|model_reasoning_effort|sandbox_mode" cc-swtich.md cc-switch-configs`：确认新增配置、`multi_agent`、`sandbox_mode` 和推理强度说明已写入。
- `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有示例分支名 `project-task-branch`，以及本次新增的 `risk-reviewer` 字符串；未发现真实密钥或遗留调试标记。
- `find cc-switch-configs/subagents -maxdepth 1 -type f -name "*.toml" -print`：确认 6 份 custom subagent 模板存在。
- `python3 ... tomllib.load ...`：输出 `ok`，`cc-switch-configs/` 下所有 TOML 配置语法解析通过。
- `git diff --check`：通过，未发现空白错误。

## 后续建议

- 后续如需实际启用全局 custom agent，可执行 `mkdir -p ~/.codex/agents && cp cc-switch-configs/subagents/*.toml ~/.codex/agents/`。
- 如果某个项目有更严格安全边界，优先使用项目级 `.codex/agents/`，并在项目级模板中继续保持 `sandbox_mode = "read-only"`。
