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
