package main

import (
	"fmt"
	"github.com/franksong/logger"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

type HelloArgs struct {
	Who  string
	Name string
	Age  uint32
}

type HelloReply struct {
	Message string
	Namer   string
	Ager    uint32
}

type HelloService struct{}

func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	reply.Message = "Hello, " + args.Who + "!"
	reply.Namer = args.Name
	reply.Ager = args.Age
	logger.Error("Say:, ", args)
	return nil
}

func main() {
	logger.Debug("test debug!!!")
	logger.Info("test info!!!")
	logger.Warning("test warning!!!")
	logger.Error("test error!!!")
	// logger.Fatal("test fatal!!!")

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json-rpc")
	pServer := &HelloService{}
	err := s.RegisterService(pServer, "Hello")
	fmt.Println("RegisterService: ")
	logger.Debug("RegisterService")
	if err != nil {
		fmt.Println("RegisterService: ", err)
	}
	http.Handle("/rpc", s)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("listenandserve error: ", err)
	}
}
