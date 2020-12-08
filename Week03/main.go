package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"github.com/go-kratos/kratos/pkg/sync/errgroup"
	"golang.org/x/sync/errgroup"
)

// errgroup
// https://github.com/golang/sync/blob/09787c993a3a/errgroup/errgroup.go
// https://github.com/go-kratos/kratos/blob/76da31effb5ece597cff22e970816a5ddd7a7659/pkg/sync/errgroup/errgroup.go
//https://zhuanlan.zhihu.com/p/64983626
// 主动关闭服务器

var server *http.Server

func main() {

	//g := errgroup.WithCancel(context.Background())
	//g := errgroup.WithContext(context.Background())
	//g, cancel := errgroup.WithContext(context.Background())
	g := new(errgroup.Group)

	g.Go(func() error {
		defer fmt.Println("finish errgroup1")
		return serveSIG()
	})
	var doneErr error

	g.Go(func() error {
		defer fmt.Println("finish errgroup2")
		serveApp()
		// select {
		// case <-ctx.Done():
		// 	doneErr = ctx.Err()
		// }
		return doneErr
	})
	g.Wait()
	log.Println("cancel all tasks")
	if doneErr != context.Canceled {
		log.Fatal("error should be Canceled")
	}

}

func serveApp() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	server = &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}
	log.Println("Starting v3 httpserver")
	if err := server.ListenAndServe(); err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")
}

func serveSIG() error {
	// 一个通知退出的chan
	quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	// 接收退出信号
	sign := <-quit
	log.Println("Accept interupt sign:", sign)

	if err := server.Close(); err != nil {
		log.Fatal("Close server:", err)
	}
	return fmt.Errorf("boom")
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

//关闭http
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the server")) // 没有输出
	log.Println("bye bye ,shutdown the server")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("shutdown the server err:", err)
	}
}
