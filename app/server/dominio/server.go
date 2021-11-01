package server

type servidorPuerto interface {
	new() interface{}
}

type Server struct {
	puerto      uint16
	tiempoFuera uint16
}

func (s *Server) new() *Server {
	server := &Server{}

	return server
}
