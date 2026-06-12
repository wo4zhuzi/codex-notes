# last30days Skill 使用笔记

## 任务背景

用户希望了解 `mvanhorn/last30days-skill` 如何安装、如何最佳使用，并将总结记录到仓库中。用户选择新建独立文档，而不是只追加到现有 `codex-skills.md`。

## 根因定位

当前仓库已有 `codex-skills.md` 和 `personal-skills.md`：

- `codex-skills.md` 适合记录通用 skill 机制和安装方法。
- `personal-skills.md` 适合记录个人自建 skill 的治理方式。

`last30days` 是第三方研究型 skill，内容涉及安装、API Key、输出目录、HTML brief、长期监控和使用边界。直接塞进通用文档会导致篇幅膨胀，因此需要独立主题文档，并从 README 与 `codex-skills.md` 建立入口。

## 执行计划

1. 查阅 `mvanhorn/last30days-skill` 的 README、运行时 `SKILL.md` 和 `CONFIGURATION.md`。
2. 新增 `last30days-skill.md`，记录安装、配置、最佳使用方式和生产使用边界。
3. 更新 `README.md` 的内容说明、目录结构、快速入口和阅读建议。
4. 更新 `codex-skills.md`，在第三方 skill 安装部分加入 `last30days` 短入口。
5. 运行文档级验证和敏感信息扫描。

## 变更内容

- 新增 `last30days-skill.md`：
  - 说明 `last30days` 的定位：近 30 天跨社区、社媒、预测市场、GitHub 和 Web 的研究型 skill。
  - 推荐 Codex 使用 `npx skills add mvanhorn/last30days-skill -g -a codex` 安装。
  - 补充 Codex 内置 `skill-installer` 的备用安装命令。
  - 记录零配置来源、可选 API Key、输出目录和 `.env` 路径。
  - 整理会议前人物研究、竞品对比、AI 工具趋势、prompt 调研、HTML brief 和长期监控等使用场景。
  - 补充生产使用边界，避免把 API Key、raw research 文件或客户数据提交进仓库。
- 更新 `README.md`：
  - 增加 `last30days` 主题说明。
  - 在目录结构和快速入口加入 `last30days-skill.md`。
  - 在阅读建议中加入近 30 天社区研究的阅读入口。
- 更新 `codex-skills.md`：
  - 在第三方 skill 安装说明后新增 `last30days` 安装示例。
  - 链接到独立主题文档，避免通用文档承载过多细节。

## 验证结果

- 已执行 `git status --short`：仅显示本次计划内改动，包含 `README.md`、`codex-skills.md`、`last30days-skill.md` 和本文档。
- 已执行 `rg -n "TODO|FIXME|your-api-key|sk-" .`：仅命中文档中的检查命令示例、既有 `project-task-branch` 示例分支名和 `risk-reviewer` 名称，未发现真实 API Key、Token、私有代理地址或遗留调试标记。
- 已执行 `git diff --check -- README.md codex-skills.md last30days-skill.md docs/changes/2026-06-12-last30days-skill.md`：未发现 Markdown 空白错误。
- 已执行 `rg -n "last30days-skill|last30days Skill|last30days Skill 使用笔记|mvanhorn/last30days-skill" README.md codex-skills.md last30days-skill.md docs/changes/2026-06-12-last30days-skill.md`：确认 README 入口、Codex Skills 入口、独立文档和变更记录均已写入。

未实际执行安装命令，避免在用户未明确要求安装前改变本机全局 Codex skills 状态。

## 后续建议

如果后续实际安装该 skill，建议单独记录本机安装结果、`npx skills list -g` 输出摘要和 Codex 重启后的触发验证结果。不要把本机 `.env` 或 raw research 文件提交到仓库。
