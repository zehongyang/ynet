package yxnet

import (
	"net"
	"time"
	"ynet/iyxnetface"
)


type Connection struct {
	ConnId uint64
	TcpServer iyxnetface.IServer
	Conn *net.TCPConn
	ExitChan chan bool
	MsgChan chan []byte
	property map[string]interface{}
}

//创建连接
func NewConnection(connId uint64,server iyxnetface.IServer,conn *net.TCPConn) *Connection {
	return &Connection{
		ConnId:connId,
		TcpServer:server,
		Conn:conn,
		ExitChan:make(chan bool),
		MsgChan:make(chan []byte),
		property:make(map[string]interface{}),
	}
}
//客户端开始
func (c *Connection) Start()  {
	go c.startRead()
	go c.startWrite()
}
//客户端读
func (c *Connection) startRead ()  {
	//设置读超时
	c.Conn.SetReadDeadline(time.Now().Add(time.Duration(c.TcpServer.GetReadTimeOut()) * time.Second))
	defer c.Stop()
	var buf = make([]byte,4096)
	for  {
		n, err := c.Conn.Read(buf)
		c.Conn.SetReadDeadline(time.Now().Add(time.Duration(c.TcpServer.GetReadTimeOut()) * time.Second))
		if err != nil {
			c.ExitChan <- true
			return
		}
		var msg = NewMessage(buf[:n],c)
		c.TcpServer.AssignMessage(msg)
	}
}
//客户端写
func (c *Connection) startWrite()  {
	//设置写超时
	c.Conn.SetWriteDeadline(time.Now().Add(time.Duration(c.TcpServer.GetWriteTimeOut()) * time.Second))
	for  {
		select {
		case msg, ok := <-c.MsgChan:
			if !ok {
				return
			}
			if _, err := c.Conn.Write(msg);err != nil{
				return
			}
			c.Conn.SetWriteDeadline(time.Now().Add(time.Duration(c.TcpServer.GetWriteTimeOut()) * time.Second))
		case <-c.ExitChan:
			return
		}
	}
}

//获取客户端id
func (c *Connection) GetConnId () uint64 {
	return c.ConnId
}

//客户端结束
func (c *Connection) Stop()  {
	//移除该客户端
	c.TcpServer.GetConnManager().RemoveConnection(c)
	//关闭客户端
	c.Conn.Close()
	//关闭管道
	close(c.MsgChan)
	close(c.ExitChan)
	//客户端关闭后处理
	c.TcpServer.CallOnClose(c)
}

//将数据写入客户端管道
func (c *Connection) Write (msg []byte) {
	c.MsgChan <- msg
}

//绑定uid
func (c *Connection) BindUid (uid int)  {
	c.property["uid"] = uid
}

//解绑uid
func (c *Connection) UnBindUid ()  {
	delete(c.property,"uid")
}

//绑定组id
func (c *Connection) JoinGroup (groupId int)  {
	c.property["groupId"] = groupId
}
//解绑组id
func (c *Connection) LeaveGroup () {
	delete(c.property,"groupId")
}
//获取客户端属性
func (c *Connection) GetProperty () map[string]interface{} {
	return c.property
}