package main

import (
	// 必须要空载注册 provider !!!
	_ "bookstore/internal/store"
	"bookstore/server"
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem")
	if err != nil {
		panic(err)
	}

	serv := server.NewBookStoreServer(":8080", s)
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}

	errChan, err := serv.ListenAndServe() // 运行http服务
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}

	log.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select { // 监视来自errChan以及c的事件
	case err = <-errChan:
		log.Println("web server run failed:", err)
	case <-c:
		log.Println("bookstore program is exiting...")
		_, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		//err = serv.Shutdown(ctx)
	}
	if err != nil {
		log.Println("bookstore program exit error:", err)
	}
	log.Println("bookstore program exit ok")
}
