# Go 双分支深度审查 Skill

## 任务背景

用户希望创建一个面向 Go 开发的只读审查 Skill。用户提供原始 ref 和开发 ref 后，Skill 不只检查 diff 语法，还需要结合旧实现、直接关联代码和历史提交，重点判断改动后的逻辑是否自洽，以及事务和数据库锁相关实现是否合理。

用户进一步要求仓库中的 Skill 模板能够跨电脑迁移：当前机器需要完成用户级安装，换电脑后也应能从本仓库重复安装和安全更新。

## 根因定位

普通 diff review 只能展示文本变化，不能稳定回答业务不变量、调用方兼容、历史修复保留、事务边界、`tx` 传递、SQL 锁范围和死锁风险等问题。

设计过程中评估过 CodeGraph。由于 CodeGraph 索引对应建立索引时的工作树，可靠比较两个 ref 需要维护两个隔离索引，索引一致性和执行成本高于本任务收益。因此最终选择 Git、`rg` 和 Go 工具链，从真实 diff 出发定向扩展相关代码。

原有个人 Skill 模板只能通过手动复制安装，缺少可验证的覆盖保护和换机恢复流程。为避免直接覆盖本机修改，本次增加通用安装器：首次安装直接复制，重复安装默认拒绝，只有显式使用 `--replace` 才会在备份旧版本后替换。

仓库还有一个既有配置问题：`.gitignore` 使用 `docs/*` 忽略整个文档目录，但未放行仓库规范指定的 `docs/specs/`，导致新 Spec 不会出现在 Git 状态中。

## 执行计划

1. 检查仓库结构、已有个人 Skill 模板、代码审查规范和近期提交。
2. 使用 `$brainstorming` 逐项确认权限、输入、merge-base 语义、验证深度、核心审查模型、输出和报告确认门。
3. 写入完整 Spec，并最小修改 `.gitignore` 放行 `docs/specs/`。
4. 使用 `writing-plans` 生成可执行实现计划并完成自检。
5. 使用 `skill-creator` 初始化 Skill 骨架。
6. 使用 TDD 实现 Git 上下文收集脚本和通用安装器。
7. 编写逻辑自洽、事务锁和 finding 分级 reference，再实现主 `SKILL.md`。
8. 更新 README、个人 Skill 说明和 Codex Skills 模板入口。
9. 执行脚本、Skill 元数据、安装器和仓库级验证。
10. 使用已验证安装器完成用户级安装，不自动 commit 或 push。

## 变更内容

- 新增 `docs/specs/2026-07-10-go-change-review-skill.md`：记录目标、输入、权限边界、审查流程、事务锁专项、输出和测试设计。
- 新增 `docs/specs/2026-07-10-go-change-review-skill-implementation-plan.md`：记录 TDD 实现步骤、完整文件内容、验证命令和权限门。
- 新增 `my-skills/my-go-change-review/`：
  - `SKILL.md` 编排双 ref 只读审查、定向历史追踪、分层验证和报告确认门。
  - `references/logic-consistency.md` 检查业务不变量、成功与失败路径、状态机、兼容性和历史修复。
  - `references/transaction-locking.md` 检查事务边界、`tx` 逃逸、锁范围、索引、死锁、隔离级别、重试和外部副作用。
  - `references/severity-guide.md` 定义 `P0` 至 `P3`、证据状态和 finding 输出模板。
  - `scripts/collect-review-context.sh` 固定 base/head SHA 和 merge-base，输出提交、变更文件和 Go module 清单。
  - `tests/test-collect-review-context.sh` 使用临时 Git fixture 验证基线前进、重命名、多 module、脏工作树、浅克隆和异常 ref。
  - `agents/openai.yaml` 提供中文 UI 元数据和显式包含 `$my-go-change-review` 的默认提示词。
- 新增 `my-skills/install-skill.sh`：安装指定个人 Skill，默认拒绝覆盖，`--replace` 时备份旧版本并执行 staged 替换。
- 新增 `my-skills/tests/test-install-skill.sh`：在临时 `CODEX_HOME` 验证首次安装、重复安装拒绝、替换备份、非法名称和缺失模板。
- 更新 `README.md`、`personal-skills.md` 和 `codex-skills.md`：增加新模板入口、安装命令、更新方式、备份位置和换机恢复流程。
- 更新 `.gitignore`：放行 `docs/specs/`，使仓库规范要求的 Spec 可以正常纳入版本控制。
- 使用仓库安装器把通过验证的模板安装到用户级 `${CODEX_HOME:-$HOME/.codex}/skills/my-go-change-review`。仓库模板仍是唯一维护源。

## 验证结果

- Git 上下文脚本 Red 阶段按预期失败，错误为缺少 `collect-review-context.sh`；实现后测试输出 `PASS: collect-review-context.sh`。
- 安装器 Red 阶段按预期失败，错误为缺少 `install-skill.sh`；实现后测试输出 `PASS: install-skill.sh`。
- 两个实现脚本和两个测试脚本均通过 `bash -n`。
- `skill-creator/scripts/quick_validate.py my-skills/my-go-change-review` 输出 `Skill is valid!`。
- `agents/openai.yaml` 已通过 PyYAML 解析，短描述长度合法，默认提示包含 `$my-go-change-review`。
- Skill 目录未发现初始化 `TODO`、`TBD`、`FIXME` 或示例占位内容。
- 用户级安装完成后，安装目录与仓库模板执行 `diff -qr` 无差异；已安装副本再次通过官方 Skill 校验。
- 安装器依赖测试全部使用临时 `CODEX_HOME`，没有在测试阶段改动真实用户目录。
- `git diff --check` 无输出，两个新增文档目录和 Skill 目录的行尾空白扫描无命中。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`。结果仅命中仓库既有检查命令、示例分支名、`risk-reviewer` 名称，以及本次计划和变更记录中的检查说明，未发现真实 API Key、Token、私有代理地址或遗留实现占位符。
- 真实凭据形态扫描无命中；README 中新增的 Skill 链接目标存在。
- 最终 `git status --short --untracked-files=all` 只包含本次 `.gitignore`、入口文档、Spec、Skill、安装器、测试和变更记录。

## 后续建议

- 重启 Codex，让新安装的 Skill 被运行时发现。
- 在真实 Go 项目中用一组包含事务或数据库锁改动的 base/head ref 做首次前向验证，并根据误报或漏报继续收紧 reference。
- 当前仓库变更尚未 commit；只有用户明确授权后才准备并执行提交，默认不 push。
