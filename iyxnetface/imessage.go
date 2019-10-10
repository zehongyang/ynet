package iyxnetface

type IMessage interface {
	GetConnection() IConnection
	GetMessage() []byte
}
