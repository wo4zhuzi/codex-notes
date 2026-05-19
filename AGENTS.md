# Repository Guidelines

## 项目结构与模块组织

本仓库是 Codex / Claude Code 使用笔记与配置模板集合，当前无应用源码、构建产物或自动化测试目录。

- `README.md`：仓库入口、内容索引和使用建议。
- `codex-cli.md`、`codex-core-commands.md`、`codex-skills.md`、`cc-swtich.md`：按主题拆分的使用笔记。
- `docs/changes/`：按日期记录每次 AI 会话产生的仓库改动，文件名格式为 `YYYY-MM-DD-<topic>.md`。
- `cc-switch-configs/`：CC-Switch 场景化 TOML 配置模板，例如 `dev-main.toml`、`dev-review.toml`。

新增主题笔记优先放在根目录 Markdown 文档中；会话变更记录放入 `docs/changes/`；仅配置模板放入 `cc-switch-configs/`。`cc-swtich.md` 文件名暂保持不变，避免破坏已有链接。

## AI 协作工作流

本仓库执行会产生文件改动的任务时，默认遵循：

```text
/init -> /status -> /plan -> AI 检查计划 -> 执行任务 -> 生成日期变更文档 -> /diff -> /review
```

硬性规则：

- 每次执行前先给出计划，并自检计划是否包含根因定位、修改步骤、验证方式和风险边界。
- 高风险任务、新项目、跨模块重构、数据迁移、权限/支付/安全相关改动，必须先进入 Spec 模式；Spec 文档位置为 `docs/specs/YYYY-MM-DD-<topic>.md`。
- 用户确认完整 Spec 前，不得开始实现；执行完成后仍必须生成 `docs/changes/YYYY-MM-DD-<topic>.md`。
- 涉及核心业务规则、bugfix、重构、权限、安全、数据一致性时，优先使用 TDD；操作流程见 `tdd-workflow.md`。
- 有测试框架时，应先写失败测试，再实现，再验证通过；没有测试框架时，必须记录等价的手动验证步骤。
- 如果本次会话产生任何仓库文件改动，必须在结束前自动创建或更新 `docs/changes/YYYY-MM-DD-<topic>.md`。
- 不需要等待用户额外要求“生成日期变更文档”，也不要只在最终回复中用文字描述替代文档。
- 变更文档应记录任务背景、根因定位、执行计划、变更内容、验证结果和直接相关的后续建议。
- 变更文档不得写入真实 API Key、个人账户 Token、私有代理地址或内部服务 URL。

## 构建、测试与本地检查命令

本项目没有包管理器脚本或编译步骤。提交前建议执行以下本地检查：

```bash
git status --short
rg -n "TODO|FIXME|your-api-key|sk-" .
```

- `git status --short`：确认只包含本次意图修改。
- `rg ...`：检查遗留占位符、调试标记和疑似密钥。

如已安装 Markdown 工具，可额外运行：

```bash
markdownlint README.md *.md
```

## 编码风格与命名约定

- 文档默认使用简体中文，保留工具名、命令名和配置键的英文原文。
- Markdown 使用清晰层级：单个 `#` 标题，正文按 `##` 分节。
- 命令、路径、文件名和配置键使用反引号，例如 `cc-switch-configs/dev-main.toml`。
- TOML 配置保持 2 空格或原文件既有缩进风格，不做无关重排。
- 新增配置模板建议使用 `dev-<scenario>.toml` 命名，场景名使用小写短横线。

## 测试指南

当前仓库没有自动化测试框架。修改文档时需手动验证：

- Markdown 链接路径是否存在，例如 `./cc-switch-configs`。
- 示例命令是否仍与文档上下文一致。
- 配置文件是否为合法 TOML；可用支持 TOML 的编辑器或 CLI 工具检查。

涉及安装步骤、第三方工具参数或模型名称时，应标注来源或更新时间，避免过期说明误导使用者。

## 提交与 Pull Request 规范

Git 历史以短中文提交信息为主，例如 `备注`、`参数调整`、`初始化中文模板`。继续沿用简短、动宾结构的中文描述，示例：

```text
更新 codex skills 排查说明
调整 dev-review 配置参数
```

PR 应包含：变更摘要、影响文件、是否涉及配置默认值变化。若修改安装或使用流程，请附上本地验证命令或手动检查结果。

## Git 与提交权限

- AI 执行修改前必须先查看 `git status --short`，识别用户已有未提交改动。
- AI 不得回滚、覆盖或混入用户已有未提交改动。
- AI 可以自动准备 commit，但不能自动执行 commit。
- 准备 commit 时，AI 必须先汇报待提交文件、变更摘要、验证结果和建议 commit message。
- 只有用户明确回复“确认提交”“yes”“commit”后，AI 才能执行 `git add` 和 `git commit`。
- AI 默认不得执行 `git push`；push 必须单独获得用户明确授权。
- 完整 Git 工作流见 `git-workflow.md`。

## 安全与配置注意事项

不要提交真实 API Key、个人账户 Token、私有代理地址或内部服务 URL。配置示例应使用占位符，并明确提示使用者按自己的供应商、模型和成本策略调整。
