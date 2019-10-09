package yxnet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-admin/pkg/log"
	"net"
	"ynet/iyxnetface"
)

var (
	server *Server
)

type Server struct {
	//服务名称
	ServerName string
	//服务版本
	Version float64
	//tcp
	Net string
	//端口
	Port uint32
	//读超时
	ReadTimeOut uint32
	//写超时
	WriteTimeOut uint32
	//host
	Host string
	//客户端连接处理
	OnConnect func(connection iyxnetface.IConnection)
	//客户端消息处理
	OnMessage func(connection iyxnetface.IConnection,msg []byte)
	//客户端关闭处理
	OnClose func(connection iyxnetface.IConnection)
}

func init()  {
	server = &Server{
		Net:"tcp",
		Port:8999,
		Host:"0.0.0.0",
		ReadTimeOut:60,
		WriteTimeOut:10,
		ServerName:"YNET",
		Version:1.0,
	}
	loadConfig()
}


func NewServer() *Server {
	return server
}

//加载配置
func loadConfig()  {
	bytes, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, server)
	if err != nil {
		fmt.Println("json unmarshal err:",err)
		return
	}
}

//服务开始
func (s *Server) Start(){
	fmt.Printf("Server %s Version %f listening on Port %d\n",s.ServerName,s.Version,s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.Net, fmt.Sprintf("%s:%d", s.Host, s.Port))
		if err != nil {
			log.Fatal("resolve tcp addr err:",err)
		}
		ln, err := net.ListenTCP(s.Net, addr)
		if err != nil {
			log.Fatal("listen tcp err:",err)
		}
		var connId uint64
		//监听连接
		for  {
			conn, err := ln.AcceptTCP()
			if err != nil {
				continue
			}
			connId++
			connection := NewConnection(connId, s, conn)
			go func() {
				s.CallOnConnect(connection)
				connection.Start()
			}()
		}
	}()
	select {}
}

//服务停止
func (s *Server) Stop() error{
	return nil
}

//启动服务
func (s *Server) Serve(){
	s.Start()
}

//注册客户端连接
func (s *Server) SetOnConnect (handler func(connection iyxnetface.IConnection))  {
	s.OnConnect = handler
}
//注册客户端消息处理
func (s *Server) SetOnMessage (handler func(connection iyxnetface.IConnection,msg []byte))  {
	s.OnMessage = handler
}
//注册客户端断开
func (s *Server) SetOnClose(handler func(connection iyxnetface.IConnection)) {
	s.OnClose = handler
}

//调用客户端连接
func (s *Server) CallOnConnect (connection iyxnetface.IConnection)  {
	if s.OnConnect != nil {
		s.OnConnect(connection)
	}
}
//调用客户端消息处理
func (s *Server) CallOnMessage (connection iyxnetface.IConnection,msg []byte)  {
	if s.OnMessage != nil {
		s.OnMessage(connection,msg)
	}
}
//调用客户端关闭
func (s *Server) CallOnClose (connection iyxnetface.IConnection)  {
	if s.OnClose != nil {
		s.OnClose(connection)
	}
}
