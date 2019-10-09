package iyxnetface

type IServer interface {
	Start()
	Stop() error
	Serve()
	SetOnConnect(func(IConnection))
	SetOnMessage(func(IConnection,[]byte))
	SetOnClose(func(IConnection))
	CallOnConnect(IConnection)
	CallOnMessage(IConnection,[]byte)
	CallOnClose(IConnection)
}
