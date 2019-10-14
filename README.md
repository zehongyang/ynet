# Ynet
Ynet 是一款轻量级的TCP框架
### 快速开始
基于Ynet框架开发业务最多需要3步即可。
1. 创建server
2. 注册业务回调函数
3. 服务启动
```go
server := yxnet.NewServer()
//客户端连接回调
server.SetOnConnect(func(conn iyxnetface.IConnection) {

})

//客户端消息发送回调
server.SetOnMessage(func(conn iyxnetface.IConnection, msg []byte) {
    //客户端和uid为1进行绑定
    conn.BindUid(1)
    //客户端解绑已有的uid
    conn.UnBindUid()
    //客户端加入某个组
    conn.JoinGroup(1)
    //客户端解绑所属组
    conn.LeaveGroup()
    //给当前客户端发送消息
    conn.Write([]byte("hello"))
    //向uid为1的客户端发送消息
    conn.GetTcpServer().SendToUid(1,[]byte("hello"))
    //向组id为1的所有客户端发送数据
    conn.GetTcpServer().SendToGroup(1,[]byte("hello"))
    //向所有在线客户端发送数据
    conn.GetTcpServer().SendToAll([]byte("hello"))
})

//客户端关闭回调
server.SetOnClose(func(conn iyxnetface.IConnection) {
	
})
server.Serve()
```
### Ynet 配置
```json
{
  "Net":"tcp",
  "Host":"127.0.0.1",
  "Port":8888,
  "ReadTimeOut":60,
  "ServerName": "ynet"
}
```

