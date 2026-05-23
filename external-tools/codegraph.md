# CodeGraph 使用笔记

更新时间：2026-05-23。

本文以 `@optave/codegraph` 为准，记录 CodeGraph 的安装、建索引、gitignore、Codex MCP 接入和日常使用流程。本文只写落地流程，不要求 AI 自动安装，也不要求 AI 自动执行索引构建。

参考来源：

- [MCP Server Space: Codegraph](https://mcpserver.space/mcp/codegraph/)
- [Skillget: optave/ops-codegraph-tool](https://skillget.dev/listings/optave-ops-codegraph-tool?lang=en)

## 结论

CodeGraph 的推荐落地方式是：

```text
用户用 CLI 建立索引 -> CodeGraph 在本地生成 .codegraph/graph.db
-> Codex 通过 CodeGraph MCP server 查询索引
-> AI 用查询结果理解符号、调用链、影响面和模块结构
```

不要让 AI 直接读 `graph.db`。数据库文件由 CodeGraph 管理，Codex 通过 MCP 工具查询即可。

## 解决什么问题

CodeGraph 用于给代码库建立函数级依赖图，帮助开发者和 AI 更快回答这类问题：

- 某个函数、类或符号在哪里定义。
- 某个函数有哪些调用方和被调用方。
- 修改某个函数会影响哪些文件或模块。
- staged diff 的影响面是什么。
- 哪些模块依赖关系复杂、存在循环依赖或边界风险。
- AI 在改代码前是否已经理解结构上下文。

它不是替代测试、类型检查或代码审查的工具。它更像一个结构化代码导航层。

## 安装方式

### 前置条件

需要本机已有 Node.js 和 npm。建议先确认：

```bash
node --version
npm --version
npx --version
```

如果没有 Node.js，可先参考仓库中的 [Codex CLI 使用笔记](../codex-cli.md) 安装 Node.js 环境。

### 推荐安装：全局安装

```bash
npm install -g @optave/codegraph
```

验证：

```bash
codegraph --help
codegraph --version
```

### 免全局安装：npx

如果不想全局安装，可以使用 `npx`：

```bash
npx @optave/codegraph --help
```

后续命令也可以写成：

```bash
npx @optave/codegraph build
npx @optave/codegraph mcp
```

两种方式取舍：

| 方式 | 优点 | 缺点 |
| --- | --- | --- |
| 全局安装 | 命令短，适合日常使用。 | 需要管理全局 npm 包版本。 |
| `npx` | 不污染全局命令，适合临时验证。 | 每次命令更长，首次执行可能更慢。 |

## 建立索引

进入目标项目根目录：

```bash
cd /path/to/your-project
```

执行索引构建：

```bash
codegraph build
```

如果使用 `npx`：

```bash
npx @optave/codegraph build
```

预期结果：

```text
.codegraph/graph.db
```

这个文件是 CodeGraph 的本地索引数据库，通常不需要手动读取，也不应提交到 Git。

建索引后可先用 CLI 做一次基础验证：

```bash
codegraph map
```

或查询某个已知符号：

```bash
codegraph query <symbol-name>
```

如果命令失败，先不要接 MCP，先把本地 CLI 跑通。

## 本地文件与 gitignore

CodeGraph 索引属于本地生成产物，建议加入 `.gitignore`：

```gitignore
.codegraph/
.code-graph/
```

原因：

- 索引可由 `codegraph build` 重新生成。
- 数据库可能包含绝对路径、符号信息、代码片段或本机缓存信息。
- 不同机器、分支和 CodeGraph 版本生成结果可能不同。
- 提交索引容易造成无意义 diff。

本仓库已在根目录 `.gitignore` 中加入以上规则。

## 与 Codex / MCP 集成

### 推荐路径

先确认本地索引已经存在：

```bash
test -f .codegraph/graph.db && echo "codegraph index exists"
```

然后把 CodeGraph MCP server 加入 Codex：

```bash
codex mcp add codegraph -- codegraph mcp
```

如果使用 `npx`：

```bash
codex mcp add codegraph -- npx @optave/codegraph mcp
```

查看是否已添加：

```bash
codex mcp list
codex mcp get codegraph
```

后续启动 Codex 后，AI 可以通过 `codegraph` MCP server 查询代码图谱。

### 工作目录边界

CodeGraph MCP server 需要知道查询哪个项目的索引。最稳妥的方式是：

1. 在目标项目根目录执行 `codegraph build`。
2. 在同一个项目目录中启动或使用 Codex。
3. 确认 `.codegraph/graph.db` 位于目标项目根目录下。

如果以后多个项目同时使用 CodeGraph，优先为每个项目明确索引路径或使用工具支持的 registry 机制。不要让 AI 通过猜测路径来直接读取数据库。

## 什么时候建索引

建议在以下时机执行 `codegraph build`：

- 第一次在项目中使用 CodeGraph。
- 切换分支后，尤其是切到别人分支或大 feature 分支。
- `git pull` 后源码、模块结构、接口或依赖有较大变化。
- 目录移动、模块拆分、函数重命名、公共类型调整之后。
- 复杂任务开始前，例如架构梳理、影响面分析、跨模块 bug 排查或代码审查。

轻量文档改动、README 改动、小配置改动通常不需要重建索引。

## 什么时候让 AI 使用

适合让 AI 通过 MCP 查询 CodeGraph：

- 修改函数前，先查调用方和被调用方。
- 排查 bug 时，查入口函数到目标函数的调用链。
- 做重构前，查影响范围和跨模块依赖。
- 代码审查前，查 staged diff 的结构影响。
- 理解陌生项目时，查核心模块、热点函数和依赖关系。

不建议让 AI 做这些事：

- 每次对话都自动重建索引。
- 直接读取或修改 `.codegraph/graph.db`。
- 在没有用户授权时安装 npm 包。
- 把索引文件加入 Git。
- 用 CodeGraph 结果替代测试、类型检查和人工审查。

推荐原则：

```text
索引构建由用户、watch、CI 或显式命令控制；
AI 主要通过 MCP 查询已有索引。
```

## 常用命令

```bash
# 查看帮助
codegraph --help

# 建立或更新索引
codegraph build

# 启动 MCP server
codegraph mcp

# 查看项目结构热点
codegraph map

# 查询符号
codegraph query <symbol-name>

# 查看函数上下文
codegraph context <symbol-name> -T

# 查看函数影响面
codegraph fn-impact <symbol-name> -T

# 查看 staged diff 影响面
codegraph diff-impact --staged -T
```

如果命令参数和当前安装版本不一致，以本机 `codegraph --help` 输出为准。

## 推荐日常流程

### 首次接入

```bash
cd /path/to/your-project
npm install -g @optave/codegraph
codegraph build
test -f .codegraph/graph.db && echo "codegraph index exists"
codex mcp add codegraph -- codegraph mcp
codex mcp list
```

### 每天开始复杂任务前

```bash
cd /path/to/your-project
git status --short
codegraph build
codex
```

进入 Codex 后，可以直接要求：

```text
先用 codegraph 查一下这个函数的调用链和影响面，再给我修改方案。
```

### 修改前

```text
让 AI 查询：
- 目标符号定义位置
- 调用方
- 被调用方
- 影响面
```

### 修改后

```bash
codegraph build
codegraph diff-impact --staged -T
```

然后再运行项目自己的测试、类型检查或 lint。

## 排查清单

### `codegraph` 命令不存在

检查：

```bash
npm list -g @optave/codegraph --depth=0
which codegraph
```

如果没有全局安装，使用：

```bash
npx @optave/codegraph --help
```

### 没有生成 `.codegraph/graph.db`

检查：

```bash
pwd
codegraph build
find . -maxdepth 3 -path "*/.codegraph/*" -o -path "*/.code-graph/*"
```

确认是在项目根目录执行，而不是在父目录、子目录或错误 worktree 中执行。

### Codex 查不到 CodeGraph 工具

检查：

```bash
codex mcp list
codex mcp get codegraph
```

如果没有配置，重新添加：

```bash
codex mcp add codegraph -- codegraph mcp
```

如果使用 `npx`：

```bash
codex mcp add codegraph -- npx @optave/codegraph mcp
```

### AI 查询结果像是旧代码

通常是索引过期。重新构建：

```bash
codegraph build
```

如果刚切换分支或刚完成大规模重构，必须重新 build 后再让 AI 查询。

### 索引文件出现在 Git 变更里

检查：

```bash
git status --short
```

如果看到 `.codegraph/` 或 `.code-graph/`，确认 `.gitignore` 已包含：

```gitignore
.codegraph/
.code-graph/
```

如果索引文件已经被 Git 跟踪，需要人工决定是否从版本控制中移除；AI 不应自动执行破坏性清理。
