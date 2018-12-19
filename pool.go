package http

// Abstract definition of goroutine pool.
// GoPool is used for http Server:
// type Server struct {
//		...         // some field of server
//
//		Pool GoPool // goroutine pool of server
// }
//
// If the pool of http server is nil,create a new goroutine
// to handle a new request as usual,or submit the task to
// goroutine pool:
// func (srv *Server) Serve(l net.Listener) error {
// 		... //some code
// 		for {
// 			rw, e := l.Accept()
//			... //some code
//			c, err := srv.newConn(rw)
//			... //some code
//			if srv.Pool == nil {
//				go c.serve()
//			} else {
//				srv.Pool.SubmitTask(func() { c.serve() })
//			}
//		}
// }
//
// You can implement GoPool with synchronous or asynchronous.
type GoPool interface {
	SubmitTask(task func()) error
}
