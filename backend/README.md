# GoghNight Backend

## 环境配置

### 1. 创建 .env 文件

在 `backend` 目录下创建 `.env` 文件，内容如下：

```env
# Supabase配置
SUPABASE_URL=your_supabase_project_url
SUPABASE_ANON_KEY=your_supabase_anon_key

# 服务器配置
PORT=8080
```

### 2. 获取Supabase配置信息

1. 登录 [Supabase](https://supabase.com)
2. 进入你的项目
3. 点击左侧菜单的 **Settings** → **API**
4. 复制以下信息：
   - **Project URL**: 填入 `SUPABASE_URL`
   - **anon public**: 填入 `SUPABASE_ANON_KEY`

### 3. 数据库表结构

确保你的Supabase项目中已创建 `messages` 表：

```sql
CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

## API端点

- `GET /api/message` - 获取最新消息
- `GET /api/messages` - 获取所有消息
- `POST /api/messages` - 创建新消息

## 启动服务

```bash
cd backend
go run main.go
```

## 注意事项

- 确保 `.env` 文件在 `backend` 目录下
- 不要将 `.env` 文件提交到版本控制系统
- 如果数据库连接失败，检查Supabase配置是否正确 


## 启动

- 前端：启动命令
```bash
npm run dev
```
http://localhost:3000/

- 后端：启动命令
```bash
go run main.go
```
http://localhost:8080/api/message