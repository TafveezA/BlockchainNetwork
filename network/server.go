package network

import "fmt"

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh chan RPC
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%v", rpc)

		case <-s.quitCh:
			break free
		}
		fmt.Println("Server Shutdown")
	}
}

func (S *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport){
			for rpc := range tr.Consume(){
				s.rpcCh <- rpc
			}
		}

	}
}

//handle
