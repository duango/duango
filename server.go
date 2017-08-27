package main

const (
	MINECRAFT_VERSION = "1.11.2"
	VERSION = "0.0.1"
)

type Server struct {
	IPAddress string
	Port uint16
	MaxConn uint32

	// status
	running bool



}

func NewServer() *Server {
	server := &Server{}

	return server
}

func (s *Server) ListenAndServer() {

}