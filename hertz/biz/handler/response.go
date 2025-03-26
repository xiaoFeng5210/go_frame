package handler

import (
	"context"
	"dqq/go/math/biz/model/student_service"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var (
	stu = &student_service.Student{Name: "大乔乔", Address: "北京"}
)

func Text(c context.Context, ctx *app.RequestContext) {
	// ctx.Status(200)
	// ctx.SetContentType("text/plain")
	// ctx.SetBodyString(fmt.Sprintf("%s live in %s", stu.Name, stu.Address))
	ctx.String(200, "%s live in %s", stu.Name, stu.Address) //包含了上面3行代码
}

func Json(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(200, stu)
}

func Xml(c context.Context, ctx *app.RequestContext) {
	ctx.XML(200, stu)
}

func ProtoBuf(c context.Context, ctx *app.RequestContext) {
	ctx.ProtoBuf(200, stu)
}

func Header(c context.Context, ctx *app.RequestContext) {
	//设置普通的响应头
	ctx.Header("token", "123456")

	//设置Cookie
	name := "language"
	value := "go"
	maxAge := 86400 * 7       //cookie的有效时间，时间单位秒。如果不设置过期时间，默认情况下关闭浏览器后cookie被删除
	path := "/"               //cookie存放目录
	domain := "www.baidu.com" //cookie从属的域名,不区分协议和端口。如果不指定domain则默认为本host(如b.a.com)，如果指定的domain是一级域名(如a.com)，则二级域名(b.a.com)下也可以访问
	secure := false           //是否只能通过https访问
	httpOnly := true          //是否允许别人通过js获取(或修改)该cookie，设为false防止XSS攻击
	//SetCookie只能执行一次,第二次SetCookie无效
	sameSite := protocol.CookieSameSiteDefaultMode
	ctx.SetCookie(name, value, maxAge, path, domain, sameSite, secure, httpOnly) //对应的response header key是"Set-Cookie"
}

func Html(c context.Context, ctx *app.RequestContext) {
	ctx.HTML(http.StatusOK, "template.html", utils.H{"title": "用户信息", "name": "zcy", "addr": "bj"})
}

func Redirect(c context.Context, ctx *app.RequestContext) {
	ctx.Redirect(http.StatusMovedPermanently, []byte("/user/html"))
}

func File(c context.Context, ctx *app.RequestContext) {
	FileName := ctx.Param("file")    //从uri路径里取得参数
	ctx.File("D:/video/" + FileName) //将指定文件写入到 Body Stream
}
