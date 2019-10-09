package iyxnetface

type IConnManager interface {
	AddConnection(IConnection)
	RemoveConnection(IConnection)
}
