# 2026-05-31 AI Agent 完整项目实战复盘

## 任务背景

用户近期使用 AI Agent 完成了一个完整项目，希望把项目背景、遇到的问题、根因判断和规避方案整理成后续可复用的个人实战手册。

项目背景包括：

- 后端使用 Golang + go-zero 微服务框架。
- 前端使用 Next.js + Ant Design X。
- 数据库使用三套 MySQL，不同微服务使用不同数据库，同时使用 Redis。
- 数据库、后端、前端使用 Docker / Dockerfile / docker-compose 一键部署。
- 业务包括注册、登录、AI Function Calling 和看板。
- 开发过程使用 Spec 模式。

## 根因定位

本次问题不是单纯的模型代码能力不足，而是完整项目执行时缺少可执行的产品验收细则、失败停止规则和最终审查机制。

具体表现：

- Docker、docker-compose、脚手架、基础接口和单元测试这类规则明确任务成功率高。
- 前期规划不足会导致后续接口、前端、测试和文档联动返工，放大 token 消耗。
- `npm install` 等依赖命令卡住时，如果没有重试上限，Agent 容易空转。
- 基础功能能跑，但注册跳转、重复注册提示、错误文案映射等产品细节容易遗漏。

## 执行计划

AI 自检结论：计划包含根因定位、修改步骤、验证方式和风险边界。本次仅新增知识库文档和 README 入口，不修改代码、配置模板或 demo。

计划：

1. 新增独立实践复盘文档，完整记录项目背景、现象、根因、推荐工作流、token 控制、验收清单和可复用 Agent 约束。
2. 更新 README 内容说明、目录结构、快速入口和阅读顺序。
3. 新增本日期变更记录。
4. 执行文档链接、关键词、敏感信息和 git 状态检查。

## 变更内容

- 新增 `ai-agent-project-practice.md`：
  - 开头明确记录本次 Go-zero 微服务、Next.js、MySQL、Redis、Docker Compose 和 Spec 模式项目背景。
  - 总结 Docker 和基础工程能力强、Spec 执行顺畅、规划不足导致返工、命令卡住导致 token 空转、产品细节不足等现象。
  - 给出 Spec、垂直切片、验收清单、失败预算、最终审查的推荐工作流。
  - 补充 token 控制策略、功能验收、接口验收、工程部署验收和可复用 Agent 约束。
- 更新 `README.md`：
  - 在内容说明中加入 AI Agent 完整项目实战复盘。
  - 在目录结构和快速入口中加入 `ai-agent-project-practice.md`。
  - 在使用建议中加入阅读该复盘文档的顺序说明。

## 验证结果

已执行：

```bash
git status --short
test -f ai-agent-project-practice.md
rg -n "AI Agent 完整项目实战复盘|ai-agent-project-practice|垂直切片|失败预算|token" README.md ai-agent-project-practice.md docs/changes/2026-05-31-ai-agent-project-practice.md
```

结果：

- `git status --short` 显示本次改动包含 `README.md`、新增 `ai-agent-project-practice.md` 和新增本变更记录。
- `test -f ai-agent-project-practice.md` 通过，确认新增文档存在。
- 关键词检查通过，确认 README 入口、复盘文档和变更记录均包含本次主题。
- 占位符和严格密钥形态检查未发现问题。

## 后续建议

后续如果继续完善，可把本文中的可复用 Agent 约束拆成项目 `AGENTS.md` 模板，或补充一份完整项目 Spec 模板，专门覆盖注册登录、错误码、前端文案和 Docker Compose 验收。
