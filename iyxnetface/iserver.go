package iyxnetface

type IServer interface {
	Start() error
	Stop() error
	Serve() error
}
