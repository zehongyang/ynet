package iyxnetface

type IConnection interface {
	Start()
	Stop()
	GetConnId() uint64
	Write([]byte)
	GetProperty() map[string]interface{}
	GetTcpServer() IServer
	BindUid(int)
	UnBindUid()
	JoinGroup(int)
	LeaveGroup()
}
