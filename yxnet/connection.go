package yxnet

import (
	"net"
	"ynet/iyxnetface"
)


type Connection struct {
	ConnId uint64
	TcpServer iyxnetface.IServer
	Conn *net.TCPConn
	ExitChan chan bool
	MsgChan chan []byte
}

//创建连接
func NewConnection(connId uint64,server iyxnetface.IServer,conn *net.TCPConn) *Connection {
	return &Connection{
		ConnId:connId,
		TcpServer:server,
		Conn:conn,
		ExitChan:make(chan bool),
		MsgChan:make(chan []byte),
	}
}
//客户端开始
func (c *Connection) Start()  {
	go c.startRead()
	go c.startWrite()
}
//客户端读
func (c *Connection) startRead ()  {
	defer c.Stop()
	var buf = make([]byte,4096)
	for  {
		n, err := c.Conn.Read(buf)
		if err != nil {
			c.ExitChan <- true
			return
		}
		go c.TcpServer.CallOnMessage(c,buf[:n])
	}
}
//客户端写
func (c *Connection) startWrite()  {
	for  {
		select {
		case msg, ok := <-c.MsgChan:
			if !ok {
				return
			}
			if _, err := c.Conn.Write(msg);err != nil{
				return
			}
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
	//关闭客户端
	c.Conn.Close()
	//关闭管道
	close(c.MsgChan)
	close(c.ExitChan)

}