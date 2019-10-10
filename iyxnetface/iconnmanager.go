package iyxnetface

import "sync"

type IConnManager interface {
	AddConnection(IConnection)
	RemoveConnection(IConnection)
	GetConnMap() sync.Map
}
