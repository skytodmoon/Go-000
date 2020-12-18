package main

import (
	"Week04/global"
	"Week04/internal/biz"
	"Week04/internal/dao"
	"Week04/internal/endpoint"
	"Week04/internal/service"
	"Week04/pkg/setting"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	var (
		// 服务地址和服务名
		servicePort = flag.Int("service.port", 10086, "service port")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	err := dao.InitMysql("127.0.0.1", "3306", "root", "mysql-password", "user")
	if err != nil {
		log.Fatal(err)
	}
	userService := service.MakeUserServiceImpl(&dao.UserDAOImpl{})

	userEndpoints := &endpoint.UserEndpoints{
		QueryUserEndpoint: endpoint.MakeQueryUserEndpoint(userService),
	}

	r := transport.MakeHTTPHandler(ctx, userEndpoints)

	go func() {
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), r)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	log.Println(error)
}

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func init() {

}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = biz.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
