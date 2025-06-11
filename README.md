# GoghNight

一个类似 Raphael.app 的定价页面项目。

## 项目结构

```
.
├── frontend/          # React 前端项目
│   ├── src/
│   │   ├── components/  # React 组件
│   │   ├── pages/      # 页面组件
│   │   └── styles/     # 样式文件
│   └── package.json
│
└── backend/           # Golang 后端项目
    ├── cmd/           # 主程序入口
    ├── internal/      # 内部包
    └── pkg/           # 公共包
```

## 开发指南

### 前端开发

```bash
cd frontend
npm install
npm run dev
```

### 后端开发

```bash
cd backend
go mod tidy
go run cmd/main.go
```

## 技术栈

- 前端：React + Next.js + TypeScript + Tailwind CSS
- 后端：Golang + Gin 