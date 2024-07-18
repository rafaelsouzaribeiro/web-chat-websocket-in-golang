package server

type Server struct {
	host    string
	port    int
	pattern string
}

func NewServer(host, pattern string, port int) *Server {
	return &Server{
		host:    host,
		port:    port,
		pattern: pattern,
	}
}
