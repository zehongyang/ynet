package yxnet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	version float64
	//tcp
	Net string
	//端口
	Port uint32
	//读超时
	ReadTimeOut int
	//host
	Host string
	//客户端连接处理
	OnConnect func(connection iyxnetface.IConnection)
	//客户端消息处理
	OnMessage func(connection iyxnetface.IConnection,msg []byte)
	//客户端关闭处理
	OnClose func(connection iyxnetface.IConnection)
	//客户端管理
	ConnManager iyxnetface.IConnManager
	//任务队列
	TaskQueue map[int]chan iyxnetface.IMessage
}

func init()  {
	server = &Server{
		Net:"tcp",
		Port:8999,
		Host:"0.0.0.0",
		ReadTimeOut:60,
		ServerName:"YNET",
		version:1.0,
		ConnManager:NewConnManager(),
		TaskQueue:make(map[int]chan iyxnetface.IMessage),
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
	fmt.Printf("Server %s Version %.1f listening on Port %d\n",s.ServerName,s.version,s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.Net, fmt.Sprintf("%s:%d", s.Host, s.Port))
		if err != nil {
			log.Fatal("resolve tcp addr err:",err)
		}
		ln, err := net.ListenTCP(s.Net, addr)
		if err != nil {
			log.Fatal("listen tcp err:",err)
		}
		s.InitTask()
		var connId uint64
		//监听连接
		for  {
			conn, err := ln.AcceptTCP()
			if err != nil {
				continue
			}
			connection := NewConnection(connId, s, conn)
			s.ConnManager.AddConnection(connection)
			go func() {
				s.CallOnConnect(connection)
				connection.Start()
			}()
			connId++
		}
	}()
	select {}
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
//处理消息
func (s *Server) DoMessageHandler (msgChan chan iyxnetface.IMessage)  {
	for  {
		select {
		case msg,ok := <-msgChan:
			if !ok {
				return
			}
			s.CallOnMessage(msg.GetConnection(),msg.GetMessage())
		}
	}
}
//将消息分发到管道
func (s *Server) AssignMessage (message iyxnetface.IMessage)  {
	index := int(message.GetConnection().GetConnId() % uint64(len(s.TaskQueue)))
	s.TaskQueue[index] <- message
}

//初始10个协程来分别处理对应管道的数据
func (s *Server) InitTask ()  {
	for i := 0; i < 10 ; i++  {
		s.TaskQueue[i] = make(chan iyxnetface.IMessage)
		go s.DoMessageHandler(s.TaskQueue[i])
	}
}
//获取客户端管理
func (s *Server) GetConnManager () iyxnetface.IConnManager  {
	return s.ConnManager
}

//将数据发送给指定uid
func (s *Server) SendToUid (uid int,msg []byte){
	s.ConnManager.GetConnMap().Range(func(key, value interface{}) bool {
		if conn,ok := value.(iyxnetface.IConnection);ok{
			if cuid,ok := conn.GetProperty()["uid"];ok {
				if cuid == uid {
					conn.Write(msg)
				}
			}
		}
		return true
	})
}

//将数据发送到指定组
func (s *Server) SendToGroup (groupId int,msg []byte) {
	s.ConnManager.GetConnMap().Range(func(key, value interface{}) bool {
		if conn,ok := value.(iyxnetface.IConnection);ok{
			if cgroupId,ok := conn.GetProperty()["groupId"];ok {
				if cgroupId == groupId {
					conn.Write(msg)
				}
			}
		}
		return true
	})
}
//将数据发送给全部客户端
func (s *Server) SendToAll (msg []byte){
	s.ConnManager.GetConnMap().Range(func(key, value interface{}) bool {
		if conn,ok := value.(iyxnetface.IConnection);ok{
			conn.Write(msg)
		}
		return true
	})
}
//获取读超时
func (s *Server) GetReadTimeOut () int  {
	return s.ReadTimeOut
}
