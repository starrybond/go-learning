# 📚 Blog-Gin 个人博客后端

> 一个简洁、RESTful 风格的个人博客后端，使用 **Go + Gin + GORM + JWT + Zap** 实现。  
> 克隆即可运行，快速二次开发！

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

---

## ✨ 主要功能
- RESTful API 设计
- JWT 登录认证
- 自动数据库迁移（GORM）
- 密码 bcrypt 加密
- Zap 结构化日志
- 支持 Swagger 文档（已预留）
- 优雅关机示例
- 易扩展的仓库模式

---

## 🚀 快速开始

### 1. 克隆项目
```bash
git clone https://github.com/YOUR_NAME/blog-gin.git
cd blog-gin
```

### 2. 配置数据库
编辑 `main.go` 中的 DSN：
```go
dsn := "root:1234@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
```
> 支持 **MySQL** 或 **SQLite**（改 driver 即可）。

### 3. 安装依赖
```bash
go mod download
```

### 4. 启动服务
```bash
go run -mod=mod .
```
服务地址：http://localhost:8080

---

## 📖 接口一览

| 方法 | 路径 | 说明 | 需登录 |
|------|------|------|--------|
| POST | `/api/register` | 用户注册 | × |
| POST | `/api/login` | 用户登录 | × |
| GET | `/api/posts` | 文章列表 | × |
| GET | `/api/posts/:id` | 文章详情 | × |
| GET | `/api/posts/:id/comments` | 评论列表 | × |
| POST | `/api/posts` | 发表文章 | ✓ |
| PUT | `/api/posts/:id` | 更新文章 | ✓ |
| DELETE | `/api/posts/:id` | 删除文章 | ✓ |
| POST | `/api/posts/:id/comments` | 发表评论 | ✓ |

> 登录后请将 JWT 放入 Header：`Authorization: <token>`

---

## 🧪 Apipost 测试用例

### 1. 注册
```http
POST http://localhost:8080/api/register
Content-Type: application/json

{
  "username": "tom",
  "password": "123456",
  "email": "tom@x.com"
}
```
**预期返回**：
```json
{"message": "ok"}
```

### 2. 登录
```http
POST http://localhost:8080/api/login
Content-Type: application/json

{
  "username": "tom",
  "password": "123456"
}
```
**预期返回**：
```json
{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
```

### 3. 发表文章（带 JWT）
```http
POST http://localhost:8080/api/posts
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json

{
  "title": "Hello Gin",
  "content": "This is my first post."
}
```
**预期返回**：
```json
{"id": 1}
```

### 4. 获取文章列表（公开）
```http
GET http://localhost:8080/api/posts
```
**预期返回**：
```json
{
  "data": [
    {
      "ID": 1,
      "Title": "Hello Gin",
      "Content": "This is my first post.",
      "UserID": 1,
      "User": { "ID": 1, "Username": "tom", ... }
    }
  ]
}
```

### 5. 发表评论（带 JWT）
```http
POST http://localhost:8080/api/posts/1/comments
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json

{
  "content": "Nice post!"
}
```
**预期返回**：
```json
{"id": 1}
```

---

## 📁 项目结构
```
blog/
├── controller/     # 控制器（接收请求、返回响应）
├── middleware/     # 中间件（JWT、日志、跨域等）
├── model/          # 数据模型与数据库迁移
├── utils/          # 工具包（JWT、日志、配置）
├── docs/           # Swagger 文档（可扩展）
├── main.go         # 项目入口
├── router.go       # 路由注册
├── go.mod          # 模块依赖
├── go.sum          # 依赖校验
└── README.md       # 项目说明
```

---

## 🧪 运行测试
```bash
go test ./... -v
```

---

## 🔧 扩展开发
1. 新增字段 → 修改 `model/entity.go` → 重启（自动迁移）
2. 新增接口 → 添加 `controller` → 注册到 `router.go`
3. 业务复杂 → 加入 `service` 层，控制器只负责输入输出校验

---

## 📄 开源协议
MIT License © [Starry](https://github.com/YOUR_NAME)

---

> 如果本项目对你有帮助，欢迎 ⭐ Star 和提 Issue / PR ！
```
