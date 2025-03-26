package main

import (
	"context"
	"dqq/go/math/biz/model/math_service"
	student_service2 "dqq/go/math/biz/model/student_service"
	"dqq/go/math/hertz_client/math"
	"dqq/go/math/hertz_client/student_service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var (
	host          = "http://localhost:5678"
	mathClient    math.Client
	studentClient student_service.Client
)

func InitLog() {
	f, err := os.OpenFile("./log/client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	hlog.SetOutput(f) //重定向到文件
	hlog.SetLevel(hlog.LevelInfo)
}

func InitClient() {
	var err error
	mathClient, err = math.NewMathClient(host,
		math.WithHertzClientOption(
			client.WithClientReadTimeout(time.Second), // 设置所有接口的默认请求超时
		),
		math.WithHeader(http.Header{"user-agent": {"go-client"}}), //指定若干个公共的request header
		math.WithHertzClientMiddleware(TimeMW),                    //指定若干个公共的中间件
	)
	if err != nil {
		log.Fatal(err)
	}

	studentClient, err = student_service.NewStudentServiceClient(host)
	if err != nil {
		log.Fatal(err)
	}
}

func TimeMW(next client.Endpoint) client.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		begin := time.Now()
		err = next(ctx, req, resp)
		log.Printf("interface %s use time %d ms", req.Path(), time.Since(begin).Milliseconds())
		return
	}
}

func Add() {
	ctx := context.Background()
	req := &math_service.AddRequest{Left: 3, Right: 5}
	resp, rawResp, err := mathClient.Add(ctx, req,
		config.WithReadTimeout(100*time.Millisecond), //设置本次调用的请求超时（优先级高于client级别）
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sum=%d\n", resp.Sum)

	fmt.Printf("response body: %s\n", string(rawResp.Body()))
	fmt.Println("response header:")
	for _, ele := range rawResp.Header.GetHeaders() {
		fmt.Printf("%s=%s\n", ele.GetKey(), ele.GetValue())
	}
	fmt.Println(strings.Repeat("-", 50))
}

func Sub() {
	ctx := context.Background()
	req := &math_service.SubRequest{Left: 5, Right: 8}
	resp, _, err := mathClient.Sub(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("diff=%d\n", resp.Diff)
	fmt.Println(strings.Repeat("-", 50))
}

func Query() {
	ctx := context.Background()
	ago := time.Now().Add(-24 * time.Hour)
	future := time.Now().Add(24 * time.Hour)
	req := &student_service2.Student{Name: "大乔乔", Score: 100, Enrollment: ago.Unix(), Graduation: future.Unix(), Address: "北京"}
	resp, _, err := studentClient.Query(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.DetailInfo)
	fmt.Println(strings.Repeat("-", 50))
}

func main() {
	InitLog()
	InitClient()

	Add()
	Sub()
	Query()
}
