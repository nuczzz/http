package main

import (
	"fmt"
	"github.com/nuczzz/gopool"
	"github.com/nuczzz/http"
	"net"
	"sync"
)

type NoPoolServer struct{}

func (nps *NoPoolServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "no pool server response: hello world")
}

type PoolServer struct {
	Pool gopool.Pool
}

var poolCount int
var lock sync.Mutex

func (ps *PoolServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	poolCount++
	fmt.Fprintf(resp, "pool server response: count: %v, total[%v], working[%v], free[%v]",
		poolCount, ps.Pool.GetTotalGoroutineNum(), ps.Pool.GetWorkingGoroutineNum(), ps.Pool.GetFreeGoroutineNum())
}

func main() {
	ln1, _ := net.Listen("tcp", "127.0.0.1:8080")
	ln2, _ := net.Listen("tcp", "127.0.0.1:8081")

	server1 := http.Server{
		Handler: &NoPoolServer{},
	}
	poolServer := &PoolServer{Pool: gopool.NewPoolWithDefault()}
	server2 := http.Server{
		Handler: poolServer,
		Pool:    poolServer.Pool,
	}

	fmt.Println("Listen and Serve: ", ln1.Addr())
	go server1.Serve(ln1)
	fmt.Println("Listen and Serve: ", ln2.Addr())
	go server2.Serve(ln2)

	select {}
}
