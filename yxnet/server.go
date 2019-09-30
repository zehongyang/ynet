package yxnet


var (
	server *Server
)

type Server struct {
	//tcp
	Net string
	//端口
	Port uint32
}

func init()  {
	server = &Server{
		Net:"tcp",
		Port:8999,
	}
}


func NewServer() *Server {
	return server
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
