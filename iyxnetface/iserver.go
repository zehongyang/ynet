package iyxnetface

type IServer interface {
	Start()
	Serve()
	SetOnConnect(func(IConnection))
	SetOnMessage(func(IConnection,[]byte))
	SetOnClose(func(IConnection))
	CallOnConnect(IConnection)
	CallOnMessage(IConnection,[]byte)
	CallOnClose(IConnection)
	AssignMessage(IMessage)
	GetConnManager() IConnManager
	SendToUid(int,[]byte)
	SendToGroup(int,[]byte)
	SendToAll([]byte)
	GetReadTimeOut() int
	GetWriteTimeOut() int
}
