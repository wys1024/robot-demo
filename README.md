# Robot-Go 聊天机器人

一个基于Ollama的智能聊天机器人项目，使用Go语言开发，提供简洁的Web界面进行交互。

## 功能特点

- 基于Ollama的智能对话功能
- 实时流式响应
- 简洁美观的Web界面
- 支持自定义端口配置

## 技术栈

- 后端：Go + Kratos框架
- AI模型：Ollama (qwen2:7b)
- 前端：原生JavaScript + SSE

## 环境要求

- Go 1.22+
- Ollama（需要预先安装并运行）
- qwen2:7b 模型（需要在Ollama中提前下载）

## 安装步骤

1. 克隆项目
```bash
git clone git@github.com:wys1024/robot-demo.git
cd robot-demo
```

2. 安装依赖
```bash
go mod tidy
```


3. 运行项目
```bash
kratos run
```

默认情况下，服务器将在 http://localhost:9000 启动

## 配置说明

可以通过环境变量配置以下参数：

- `PORT`: 服务器端口号（默认：9000）

## API接口

### 聊天接口

- 路径：`POST /api/chat`
- 请求格式：
```json
{
    "question": "你的问题"
}
```
- 响应格式：Server-Sent Events (SSE)

## 使用说明

1. 确保Ollama服务正在运行，并已下载qwen2:7b模型(也可以是其他模型)
2. 启动服务后，访问 http://localhost:9000
3. 在输入框中输入问题，点击发送按钮或按回车键发送
4. 等待AI助手的实时回复

## 注意事项

- 请确保Ollama服务正常运行
- 首次使用需要下载模型，可能需要一些时间
- 建议在本地网络环境良好的情况下使用

---