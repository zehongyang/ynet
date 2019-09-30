package yxnet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	server *Server
)

type Server struct {
	//tcp
	Net string
	//端口
	Port uint32
	//读超时
	ReadTimeOut uint32
	//写超时
	WriteTimeOut uint32
}

func init()  {
	server = &Server{
		Net:"tcp",
		Port:8999,
		ReadTimeOut:60,
		WriteTimeOut:10,
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
func (s *Server) Start() error{

	return nil
}

//服务停止
func (s *Server) Stop() error{
	return nil
}

//启动服务
func (s *Server) Serve() error{
	err := s.Start()
	return err
}
