# 个人自建 Skills 主题文档

## 任务背景

用户希望将“个人创建 skill”作为一个独立主题沉淀到仓库中，尤其是把多个 skills、MCP 和外部工具融合成固定工作流时，能和第三方 skill 明确区分。

## 根因定位

当前仓库已有 `codex-skills.md` 记录 Codex Skills 的安装、触发和 Understand Anything 接入，也有 `codex-plugins.md` 说明 plugin 与 skill 的区别。但缺少一个集中说明个人自建 skill 的治理文档，包括：

- 如何用命名区分自建 skill 和三方 skill。
- 如何选择用户级或项目级存放路径。
- 如何把多个 skills、MCP 工具和验证步骤组合成个人工作流。
- 如何避免把组合型 workflow 写成三方 skill 内容复制。

因此需要新增独立主题文档，并在现有入口文档中建立链接。

## 执行计划

1. 新增 `personal-skills.md`，记录个人自建 skill 的定位、命名、路径、元数据、组合型设计原则和示例模板。
2. 更新 `README.md`，加入内容说明、目录结构、快速入口和阅读顺序。
3. 更新 `codex-skills.md`，在 skill 存放位置说明中补充个人自建 skill 的短规则，并链接到新文档。
4. 运行文档检查和敏感信息扫描。
5. 根据用户反馈新增 `my-skills/` 模板库，将具体个人 skill 内容从说明文档迁移到模板目录。

## 变更内容

- 新增 `personal-skills.md`：
  - 明确个人自建 skill 是个人工作流资产，不是三方工具说明或 plugin 分发包。
  - 固定 `my-` 前缀规范，例如 `my-codebase-intel`、`my-debug-flow`、`my-review-gate`。
  - 区分用户级 `~/.codex/skills/my-*` 和项目级 `.agents/skills/my-*`。
  - 给出中文 `SKILL.md` frontmatter 和 `agents/openai.yaml` 的个人风格示例。
  - 补充组合型 skill 的事实优先级，并链接到 `my-skills/my-codebase-intel/` 模板。
  - 补充 `my-codebase-intel` 使用示例，覆盖项目看板、代码分析、项目介绍、模块介绍和改动审查。
  - 将使用示例调整为短指令，具体执行流程沉淀到 `my-skills/my-codebase-intel/SKILL.md`。
- 新增 `my-skills/my-codebase-intel/`：
  - 新增 `SKILL.md`，定义项目看板、项目介绍、模块介绍、代码分析、bug 排查、重构评估和改动审查等任务类型。
  - 新增 `agents/openai.yaml`，提供 `个人 · Codebase Intel` 的 UI 元数据和默认短 prompt。
- 更新 `README.md`：
  - 增加个人自建 Skills 的内容说明、`my-skills/` 目录项和快速入口。
  - 在使用建议中加入阅读 `personal-skills.md` 的时机。
- 更新 `codex-skills.md`：
  - 新增“个人自建 Skills”小节。
  - 说明 `my-` 前缀、典型路径和不要写入 `.system` 目录。
  - 说明 `my-skills/` 是仓库模板库，真实启用仍需复制或链接到 Codex 识别目录。

## 验证结果

- 已执行 `git status --short`：仅显示本次计划内文档改动，包含 `README.md`、`codex-skills.md`、`personal-skills.md` 和本文档。
- 已执行 `rg -n "个人自建 Skills|personal-skills|my-codebase-intel|my-" README.md codex-skills.md personal-skills.md docs/changes/2026-06-06-personal-skills.md`：确认 README 入口、skills 文档入口、新主题文档和变更记录均已写入。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有 `project-task-branch` 示例分支名和 `risk-reviewer` 名称，未发现真实密钥或遗留调试标记。
- 已执行 `git diff --check -- README.md codex-skills.md personal-skills.md docs/changes/2026-06-06-personal-skills.md`：未发现 Markdown 空白错误。
- 根据用户反馈，已将 `personal-skills.md` 中的 `my-codebase-intel` 示例 skill 内容改为中文，保留 Understand Anything、CodeGraph、diff review 等工具名和技术关键词。
- 根据用户反馈，已补充可直接复制的使用示例，包括生成项目理解看板、分析代码并给出修改方案、生成项目介绍、生成模块介绍和审查当前改动。
- 根据用户反馈，已新增 `my-skills/my-codebase-intel/` 模板目录，并将长流程收敛到真实 `SKILL.md` 模板中，`personal-skills.md` 只保留短指令示例和模板链接。
- 已执行 `find my-skills -maxdepth 4 -type f | sort`：确认新增 `my-skills/my-codebase-intel/SKILL.md` 和 `my-skills/my-codebase-intel/agents/openai.yaml`。
- 已执行 `python3 ~/.codex/skills/.system/skill-creator/scripts/quick_validate.py my-skills/my-codebase-intel`：因当前 Python 环境缺少 `yaml` 模块失败；已人工核对 `SKILL.md` frontmatter 仅包含 `name` 和 `description`，`name: my-codebase-intel` 符合小写短横线规则。
- 已执行 `rg -n "my-skills|my-codebase-intel|项目看板|项目介绍|模块介绍|改动审查|个人 · Codebase Intel" ...`：确认 README、skills 文档、个人 skill 文档、模板目录和变更记录均已同步。
- 已执行 `git diff --check -- README.md codex-skills.md personal-skills.md my-skills docs/changes/2026-06-06-personal-skills.md`：未发现 Markdown 空白错误。

## 后续建议

后续创建 `my-codebase-intel` 等用户级 skill 时，可直接以 `my-skills/<skill-name>/` 中的模板为基准，再根据实际工具可用性微调。
