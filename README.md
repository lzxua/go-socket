## 一个简单的socket聊天室

### 启动方式
```go
cd socket
go run socket.go
```

### 连接聊天
#### 打开本地终端
```go
nc 127.0.0.1 8080
```
#### 再开启一个新终端重复上述指令
然后在随意一个终端下，输入聊天内容即可


#### 目前只有广播聊天，后续会慢慢的加入功能

#### 1. 加入私聊，以及退出功能


#### 2. 更换map集合为线程安全的sync.Map