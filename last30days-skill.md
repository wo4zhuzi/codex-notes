# last30days Skill 使用笔记

## 结论

`last30days` 是一个近 30 天研究型 Codex / Claude / Agent Skills skill，用于围绕人物、公司、产品、技术、舆情或竞品，从 Reddit、X、YouTube、TikTok、Hacker News、Polymarket、GitHub 和 Web 等来源汇总近期讨论，并按社区参与度和事实信号生成简报。

它和 `brainstorming`、`writing-plans`、`systematic-debugging` 这类工程化 skills 的属性不同。工程化 skills 主要服务项目执行、代码质量和交付闭环；`last30days` 更像外部信息入口和研究引擎，服务最新信息获取、趋势追踪、社区反馈分析和个人知识库上游输入。

Codex 中推荐使用官方 Agent Skills CLI 安装：

```bash
npx skills add mvanhorn/last30days-skill -g -a codex
```

安装后重启 Codex，再用自然指令触发：

```text
使用 last30days 调研 OpenAI Codex 最近 30 天的社区反馈
```

或在支持 slash skill 的宿主中直接使用：

```text
/last30days OpenAI Codex
```

## 根因定位

普通 Web 搜索更偏向网页索引和编辑内容，容易遗漏社区讨论、短视频评论、预测市场、GitHub 近期活动和社媒高参与内容。`last30days` 的价值在于把这些分散来源统一交给 agent 检索、评分和综合，适合回答“最近大家真实在讨论什么”。

因此它适合单独拆分成主题文档：它不是“帮我把当前工程任务做完”的助手，而是“帮我持续打开外部信息渠道”的工具。后续如果结合自动化工作流，可以把一次性搜索升级成持续信息雷达，再沉淀到 Obsidian、Notion、RAG 知识库或团队协作系统中。

该 skill 的运行时规范以仓库内 `skills/last30days/SKILL.md` 为准。README 说明当前主流程，但当 README、配置文档和 `SKILL.md` 出现差异时，优先以 `SKILL.md` 为准。

## 和其他 Skills 的分工

| 类型 | 代表工具 | 核心职责 |
| --- | --- | --- |
| 工程化 skills | `brainstorming`、`writing-plans`、`systematic-debugging`、`verification-before-completion` | 服务需求澄清、计划拆解、bug 根因定位、验证和交付闭环。 |
| 代码库理解 skills | `my-codebase-intel`、Understand Anything、CodeGraph 相关流程 | 服务已有代码库理解、调用链分析、影响面评估和改动审查。 |
| 信息渠道 skill | `last30days` | 服务外部信息发现、趋势追踪、社区反馈、竞品情报和人物动态。 |
| 知识库与自动化 | RAG、Obsidian、Notion、n8n、Slack / 飞书推送 | 服务信息归档、周期性摘要、团队分发和后续检索增强。 |

推荐把 `last30days` 放在信息输入层，而不是工程执行层：

```text
外部社区 / 平台 / 市场信号
-> last30days 调研和保存 raw 结果
-> 摘要、分类、去重和标签化
-> Obsidian / Notion / RAG 知识库
-> Codex / Claude 后续分析和工程决策上下文
```

## 安装方式

### 推荐：Agent Skills CLI

Codex、Cursor、Gemini CLI、Copilot 等 Agent Skills hosts 推荐用 `npx skills` 安装。

```bash
npx skills add mvanhorn/last30days-skill -g -a codex
```

参数说明：

| 参数 | 含义 |
| --- | --- |
| `mvanhorn/last30days-skill` | GitHub 仓库。 |
| `-g` | 全局安装到当前用户，适合所有项目复用。 |
| `-a codex` | 明确安装到 Codex 宿主，避免自动识别错误。 |

常用维护命令：

```bash
npx skills update last30days -g
npx skills list -g
npx skills remove last30days -g
```

如果同一台机器还要给其他宿主安装，可以重复指定 `-a`：

```bash
npx skills add mvanhorn/last30days-skill -g -a codex -a cursor
```

### 备用：Codex 内置 skill-installer

如果希望沿用 Codex 内置 `skill-installer`，安装路径必须指向 skill 目录，而不是 `SKILL.md` 文件：

```bash
python3 ~/.codex/skills/.system/skill-installer/scripts/install-skill-from-github.py \
  --repo mvanhorn/last30days-skill \
  --path skills/last30days
```

验证：

```bash
test -f ~/.codex/skills/last30days/SKILL.md && echo "installed"
```

安装或更新后建议重启 Codex，让新 skill 被当前会话加载。

### Claude Code 安装

Claude Code 官方推荐走 plugin marketplace：

```text
/plugin marketplace add mvanhorn/last30days-skill
/plugin install last30days
```

如果 Claude Code 中同时用 marketplace plugin 和 `npx skills` 安装，同名 `/last30days` 可能出现重复入口。推荐一台机器只保留一种安装方式。

## 基础配置

### 零配置可用来源

以下来源默认或接近默认可用：

| 来源 | 依赖 |
| --- | --- |
| Reddit public JSON | 无需 API Key。 |
| Hacker News | 无需 API Key。 |
| Polymarket | 无需 API Key。 |
| GitHub | 安装并登录 `gh` CLI 时效果更好。 |
| YouTube | 安装 `yt-dlp` 后可抓取公开视频和字幕。 |

安装 `yt-dlp`：

```bash
brew install yt-dlp
```

### 可选 API Key

全局配置文件默认位置：

```text
~/.config/last30days/.env
```

项目级配置文件位置：

```text
.claude/last30days.env
```

项目级配置优先于全局配置。配置文件中只写自己的真实密钥，本仓库文档只保留占位符。

示例：

```bash
# 输出目录
LAST30DAYS_MEMORY_DIR=~/Documents/Last30Days

# Web 搜索后端，Brave 是常用低成本选项
BRAVE_API_KEY=<your-brave-key>

# TikTok / Instagram / Threads / Pinterest / YouTube comments 等来源
SCRAPECREATORS_API_KEY=<your-scrapecreators-key>
INCLUDE_SOURCES=tiktok,instagram,threads

# X / Twitter，选择一种方式即可
XAI_API_KEY=<your-xai-key>
# FROM_BROWSER=firefox

# Bluesky
BSKY_HANDLE=<your-handle>.bsky.social
BSKY_APP_PASSWORD=<your-app-password>

# Perplexity Sonar / Deep Research
OPENROUTER_API_KEY=<your-openrouter-key>
```

POSIX 系统建议限制权限：

```bash
chmod 600 ~/.config/last30days/.env
```

### 输出目录

默认输出目录：

```text
~/Documents/Last30Days/
```

每次运行会保存原始研究文件：

```text
<topic-slug>-raw.md
```

可以用环境变量或单次参数调整：

```bash
export LAST30DAYS_MEMORY_DIR="$HOME/Documents/Last30Days"
```

```text
/last30days OpenAI Codex --save-dir ~/Documents/Last30Days --save-suffix=codex
```

## 最佳使用方法

### 适合场景

| 场景 | 示例 |
| --- | --- |
| 会议前人物研究 | `/last30days Peter Steinberger` |
| 公司或产品近况 | `/last30days OpenAI Codex` |
| 竞品对比 | `/last30days Cursor vs Codex vs Claude Code` |
| AI 工具趋势 | `/last30days AI coding agents` |
| 社区需求调研 | `/last30days what users want in React` |
| prompt 技法调研 | `/last30days Nano Banana Pro prompting` |
| 近期舆情和新闻 | `/last30days Universal Epic Universe` |

### 知识库与自动化工作流

如果只是手动运行一次，`last30days` 是研究工具；如果结合自动化工作流，它可以变成个人信息渠道的上游采集层。

推荐链路：

```text
定期主题监控
-> last30days 拉取近 30 天社区和市场信号
-> 保存 raw Markdown 或 SQLite
-> 自动摘要成日报 / 周报
-> 推送到 Obsidian / Notion / Slack / 飞书
-> 进入 RAG 知识库
-> 作为后续 Agent 分析上下文
```

适合持续跟踪的主题：

- AI 编程工具和模型生态，例如 Codex、Claude Code、Cursor、OpenClaw。
- 行业趋势和竞品动态，例如某个 SaaS 赛道、开源项目或客户行业。
- 关键人物和团队动态，例如创始人、投资人、开源维护者。
- 用户真实反馈，例如某类工具的抱怨、付费意愿、迁移原因和使用场景。
- prompt 技法和工作流实践，例如图像生成、代码 agent、自动化编排。

这类工作流的关键不是“多查几个网页”，而是稳定地把外部信号转成可复用知识资产。

### 推荐工作流

1. 先用明确主题触发一次研究：

```text
/last30days OpenAI Codex
```

2. 等 skill 输出综合简报和原始结果路径。
3. 在同一会话继续追问，不要立刻换新主题：

```text
从结果里提炼 Codex 用户最常抱怨的 5 个问题
```

```text
基于这些发现，帮我写一份内部分享提纲
```

```text
把 Cursor、Codex、Claude Code 的差异整理成决策表
```

4. 需要可分享文件时，要求导出 HTML：

```text
/last30days OpenAI Codex --emit=html
```

或：

```text
使用 last30days 调研 OpenAI Codex，并生成适合发到 Slack 的 HTML brief
```

5. 需要长期观察时，使用 `--store` 保存到 SQLite，再配合 watchlist / briefing：

```text
/last30days "AI coding agents" --store
```

### 提问质量建议

推荐写法：

```text
/last30days OpenAI Codex
/last30days Cursor vs Codex
/last30days what users dislike about AI coding agents
/last30days Claude Code subagents best practices
```

不推荐写法：

```text
/last30days 最近有什么新鲜事
/last30days 帮我查一下
```

原因是该 skill 会先解析人物、产品、社区、账号、仓库和关键词。主题越具体，预研究和来源分配越稳定。

## 生产使用建议

- 把 `last30days` 当作“近期社区研究入口”，不要当作唯一事实来源。
- 涉及法律、医疗、金融、生产决策时，应把输出作为线索，再回到一手来源核验。
- 保存的 raw 文件可能包含公开帖文、链接和摘要，不建议提交进仓库。
- 不要把 `~/.config/last30days/.env`、`.claude/last30days.env` 或 raw research 文件提交到 Git。
- 多客户场景建议按客户设置 `LAST30DAYS_MEMORY_DIR` 或 `--save-suffix`，避免研究文件互相覆盖。
- 需要 X / TikTok / Instagram 等付费或账号来源时，先确认预算和 API Key 权限。
- 如果结果缺少预期来源，优先运行诊断命令：

```bash
python3 scripts/last30days.py --diagnose
```

诊断命令需要在已安装的 `last30days` skill 目录内运行，例如：

```bash
cd ~/.codex/skills/last30days
python3 scripts/last30days.py --diagnose
```

## 常见问题

### 是否必须配置 API Key？

不必须。Reddit public JSON、Hacker News、Polymarket 等可零配置使用。但 X、TikTok、Instagram、Threads、Pinterest、Perplexity、部分 Web 搜索后端需要额外凭证或 CLI。

### Codex 中是否必须使用 `/last30days`？

不一定。是否支持 slash 形式取决于宿主。Codex 中更稳妥的写法是直接点名：

```text
使用 last30days skill 调研 Cursor vs Codex 最近 30 天社区反馈
```

### 为什么建议全局安装？

`last30days` 是跨项目研究工具，通常不属于某个仓库的业务依赖。全局安装后可以在所有项目复用，也避免把第三方 skill 文件提交到业务仓库。

### 什么时候不适合使用？

- 只需要查询一个确定事实，例如官方版本号或某个文档页面。
- 用户明确要求只使用官方来源。
- 主题涉及私有数据、内部客户信息或不能外发的敏感上下文。
- 输出需要严格可审计引用时，应额外保存和核对原始链接。

## 参考来源

- GitHub 仓库：<https://github.com/mvanhorn/last30days-skill>
- 运行时规范：<https://github.com/mvanhorn/last30days-skill/blob/main/skills/last30days/SKILL.md>
- 配置说明：<https://github.com/mvanhorn/last30days-skill/blob/main/CONFIGURATION.md>
