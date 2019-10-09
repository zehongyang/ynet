package yxnet

import (
	"sync"
	"ynet/iyxnetface"
)

type ConnManager struct {
	ConnMap sync.Map
}

func NewConnManager() *ConnManager {
	return &ConnManager{}
}
//新增客户端
func (c *ConnManager) AddConnection (connection iyxnetface.IConnection)  {
	c.ConnMap.Store(connection.GetConnId(),connection)
}
//移除客户端
func (c *ConnManager) RemoveConnection(connection iyxnetface.IConnection){
	c.ConnMap.Delete(connection.GetConnId())
}