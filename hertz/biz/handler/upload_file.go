package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// 业务Handler以及Server端中间件都要满足此函数签名
func UploadFile(c context.Context, ctx *app.RequestContext) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Printf("get file error %v\n", err)
		ctx.String(http.StatusInternalServerError, "upload file failed")
	} else {
		if err = ctx.SaveUploadedFile(file, "./data/"+file.Filename); err == nil { //把用户上传的文件存到data目录下
			ctx.String(http.StatusOK, file.Filename)
		} else {
			fmt.Printf("save file to %s failed: %v\n", "./data/"+file.Filename, err)
		}
	}
}

func UploadMultiFile(c context.Context, ctx *app.RequestContext) {
	form, err := ctx.MultipartForm() //MultipartForm中包含多个文件
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	} else {
		//从MultipartForm中获取上传的文件
		files := form.File["files"]
		for _, file := range files {
			ctx.SaveUploadedFile(file, "./data/"+file.Filename) //把用户上传的文件存到data目录下
		}
		ctx.String(http.StatusOK, "upload "+strconv.Itoa(len(files))+" files")
	}
}
