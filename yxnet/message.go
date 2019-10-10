package yxnet

import "ynet/iyxnetface"

type Message struct {
	Msg []byte
	Conn iyxnetface.IConnection
}
//创建消息
func NewMessage(msg []byte,conn iyxnetface.IConnection) *Message {
	return &Message{
		Msg:msg,
		Conn:conn,
	}
}

//获取客户端
func (m *Message) GetConnection () iyxnetface.IConnection {
	return m.Conn
}
//获取消息
func (m *Message) GetMessage () []byte {
	return m.Msg
}
