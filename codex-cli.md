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
