package main

import (
	"bytes"
	"dqq/go/math/biz/model/student_service"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"google.golang.org/protobuf/proto"
)

var (
	host = "http://localhost:5678/"
)

func PrintResp(resp *http.Response, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("StatusCode", resp.StatusCode)
	}
	bs, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bs))
}

func processResponse(resp *http.Response) {
	defer resp.Body.Close()
	fmt.Println("响应头：")
	for k, v := range resp.Header {
		fmt.Printf("%s=%s\n", k, v[0])
	}
	fmt.Print("响应体：")
	io.Copy(os.Stdout, resp.Body)
	os.Stdout.WriteString("\n")
	if resp.StatusCode != http.StatusOK {
		slog.Error("异常状态码", "http response code", resp.StatusCode)
	}
	os.Stdout.WriteString("\n")
}

func Get(path string) {
	fmt.Println("GET " + path)
	resp, err := http.Get(host + path)
	if err != nil {
		slog.Error("http get failed", "error", err)
		return
	}
	processResponse(resp)
}

func Add() {
	fmt.Print("Add ")
	resp, err := http.Get(host + "/add?left=3&right=5")
	PrintResp(resp, err)
}

func Sub() {
	fmt.Print("Sub ")
	resp, err := http.PostForm(host+"/sub", url.Values{"left": {"5"}, "right": {"8"}})
	PrintResp(resp, err)
}

func Query() {
	fmt.Print("Query ")
	resp, err := http.PostForm(host+"/student?name=大乔乔&addr=北京", url.Values{"score": {"100"}})
	PrintResp(resp, err)
}

func Restful() {
	fmt.Print("Restful ")
	resp, err := http.PostForm(host+"/student/大乔乔/北京/海淀", url.Values{"score": {"100"}})
	PrintResp(resp, err)
}

func Header() {
	fmt.Print("Header ")
	request, _ := http.NewRequest(http.MethodPost, host+"/student/header", strings.NewReader("score=100"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("name", "大乔乔")
	request.Header.Add("addr", "北京")
	client := http.Client{}
	resp, err := client.Do(request)
	PrintResp(resp, err)
}

func Cookie() {
	fmt.Print("Cookie ")
	request, _ := http.NewRequest(http.MethodPost, host+"/student/cookie", strings.NewReader("score=100"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(&http.Cookie{Name: "name", Value: url.QueryEscape("大乔乔")})
	request.AddCookie(&http.Cookie{Name: "addr", Value: url.QueryEscape("北京")})
	client := http.Client{}
	resp, err := client.Do(request)
	PrintResp(resp, err)
}

func PostForm() {
	fmt.Print("PostForm ")
	resp, err := http.PostForm(host+"/student/form", url.Values{"name": {"大乔乔"}, "addr": {"北京"}, "score": {"100"}})
	PrintResp(resp, err)
}

func PostJson() {
	fmt.Print("PostJson ")
	ago := time.Now().Add(-24 * time.Hour)
	future := time.Now().Add(24 * time.Hour)
	stu := &student_service.Student{
		Name:       "大乔乔",
		Address:    "北京",
		Score:      29,
		Enrollment: ago.Unix(),
		Graduation: future.Unix(),
	}
	bs, _ := sonic.Marshal(stu)
	resp, err := http.Post(host+"/student/json", "application/json", bytes.NewReader(bs))
	PrintResp(resp, err)
}

func PostPb() {
	fmt.Print("PostPb ")
	ago := time.Now().Add(-24 * time.Hour)
	future := time.Now().Add(24 * time.Hour)
	stu := student_service.Student{
		Name:       "大乔乔",
		Address:    "北京",
		Score:      29,
		Enrollment: ago.Unix(),
		Graduation: future.Unix(),
	}
	bs, err := proto.Marshal(&stu)
	if err != nil {
		log.Fatal("proto.Marshal error", err)
	}
	resp, err := http.Post(host+"/student/pb", "application/x-protobuf", bytes.NewReader(bs))
	PrintResp(resp, err)
}

func GetUserPB(path string) {
	fmt.Println("GET " + path)
	resp, err := http.Get(host + path)
	if err != nil {
		slog.Error("http get failed", "error", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("响应头：")
	for k, v := range resp.Header {
		fmt.Printf("%s=%s\n", k, v[0])
	}
	if resp.StatusCode != http.StatusOK {
		slog.Error("异常状态码", "http response code", resp.StatusCode)
		io.Copy(os.Stdout, resp.Body)
	} else {
		bs, _ := io.ReadAll(resp.Body)
		var stu student_service.Student
		if err := proto.Unmarshal(bs, &stu); err == nil {
			fmt.Printf("pb反序列化成功, name=%s, addr=%s\n", stu.Name, stu.Address)
		}
	}
	os.Stdout.WriteString("\n\n")
}

func DownloadFile(path string, file string) {
	resp, err := http.Get(host + path + file)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("StatusCode", resp.StatusCode)
	}

	fout, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()
	io.Copy(fout, resp.Body)
}

func main() {
	Add()
	Sub()
	Query()
	Restful()
	PostForm()
	Header()
	Cookie()
	PostJson()
	PostPb()
	fmt.Println()
	fmt.Println()

	Get("/user/text")
	Get("/user/json")
	Get("/user/xml")
	GetUserPB("/user/pb")
	Get("/user/cookie")
	Get("/user/html")
	Get("/user/old_page")
	DownloadFile("/file/", "通天.mp4")
}
