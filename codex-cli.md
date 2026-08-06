# Codex CLI

## 安装

### 前置条件：Node.js 环境

#### 安装 nvm

```shell
wget -qO- https://raw.githubusercontent.com/creationix/nvm/v0.34.0/install.sh | bash
```

配置 `NVM_DIR`：

```shell
export NVM_DIR="${XDG_CONFIG_HOME/:-$HOME/.}nvm"
```

加载 `nvm`：

```shell
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
```

注意：`nvm` 安装完成后，建议重新打开一个终端窗口再执行后续命令。

安装 Node.js：

```shell
nvm install v22
```

#### 安装最新版本

```shell
npm install -g @openai/codex@latest
```

#### 验证是否安装成功

```shell
codex --version
```

## 登录 / 初始化

```shell
codex
```

## `/init` 生成仓库说明

进入某个项目仓库后，可以在 Codex 会话中执行：

```text
/init
```

`/init` 会扫描当前仓库，并生成或更新仓库级说明文件 `AGENTS.md`。该文件用于记录项目结构、常用命令、测试方式、代码规范和协作约束，后续 Codex 在该仓库中工作时会读取其中的说明。

推荐使用方式：

1. 先进入目标仓库根目录。
2. 执行 `codex` 启动 Codex 会话。
3. 在会话中输入 `/init`。
4. 检查生成的 `AGENTS.md`，补充项目特有的构建、测试、发布和编码规范。
5. 将 `AGENTS.md` 提交到仓库，保证团队和后续会话使用同一份上下文。

## 跨项目参考代码

当项目 B 是当前开发项目，但实现时需要参考项目 A 的代码或设计，可以将项目 B 设为主工作目录，并把项目 A 加入额外目录：

```shell
codex \
  -C /data/b_project \
  -s workspace-write \
  --add-dir /data/a_project
```

参数含义：

- `-C /data/b_project`：将项目 B 设置为 Codex 的工作根目录。
- `-s workspace-write`：允许 Codex 在工作区内写入文件，避免依赖全局配置中的默认沙箱模式。
- `--add-dir /data/a_project`：在 `workspace-write` 模式下，将项目 A 加入主工作区之外的额外可写目录。

`-s workspace-write` 和 `--add-dir` 的职责不同：前者启用工作区写入模式，后者扩展该模式允许写入的目录范围。如果当前默认沙箱是 `read-only`，只设置 `--add-dir` 不能获得预期的写入能力，因此推荐在命令中显式指定二者。

### 仅将参考项目作为只读上下文

`--add-dir` 授予的是额外目录的写入权限，不是只读权限。如果项目 A 只用于参考，应在启动提示中明确修改边界：

```shell
codex \
  -C /data/b_project \
  -s workspace-write \
  --add-dir /data/a_project \
  "项目 B 是唯一修改目标；项目 A 仅用于只读参考，禁止修改其中任何文件。实现前先分析项目 A 的相关设计。"
```

上述提示属于任务约束，不能替代操作系统权限控制。生产环境或高风险代码库需要强制只读时，推荐将项目 A 以只读方式挂载，或者提供一个只读快照，再通过 `--add-dir` 加入该路径。不要为了读取参考项目而使用 `--sandbox danger-full-access` 或 `--dangerously-bypass-approvals-and-sandbox`。

### 使用边界

- 项目 B 是当前工作根目录，修改、测试和 Git 检查默认应围绕项目 B 执行。
- 不要依赖 Codex 自动采用项目 A 的 `AGENTS.md`。如需遵循项目 A 的架构或编码约定，应在提示中要求先读取对应文件，并明确这些规则仅用于参考。
- 可以重复使用 `--add-dir` 加入多个参考目录。
- 路径包含空格时，应使用引号包裹完整路径。
- 开始修改前可要求 Codex 分别执行 `git status --short`，确认项目 A 没有产生变更，项目 B 只包含预期改动。

资料来源（核对时间：2026-08-06）：

- [Codex Developer commands：Flag combinations and safety tips](https://learn.chatgpt.com/docs/developer-commands#flag-combinations-and-safety-tips)
- 本机 `codex-cli 0.146.0` 的 `codex --help` 输出

## 配置文件

### 全局配置文件

```text
~/.codex/config.toml
```

### 查看当前状态

```text
/status
```

在状态信息中查看 `Model provider`，确认当前模型供应商是否已切换。
