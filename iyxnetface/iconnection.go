package iyxnetface

type IConnection interface {
	Start()
	Stop()
	GetConnId() uint64
}
