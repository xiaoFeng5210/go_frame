package main

import (
	handler "dqq/go/frame/post/handler/fiber"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

func main() {
	Init()
	go ListenTermSignal() //监听信号

	engine := html.New("post/views/html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use("/css", static.New("post/views/css")) //在url是访问目录/js相当于访问文件系统中的views/js目录
	app.Use("/js", static.New("post/views/js"))
	app.Use("/favicon.ico", static.New("post/views/img/dqq.png"))

	app.Get("/login", func(ctx fiber.Ctx) error { return ctx.Status(http.StatusOK).Render("login", nil) })
	app.Get("/regist", func(ctx fiber.Ctx) error { return ctx.Status(http.StatusOK).Render("user_regist", nil) })
	app.Get("/modify_pass", func(ctx fiber.Ctx) error { return ctx.Status(http.StatusOK).Render("update_pass", nil) }, handler.Auth)
	app.Post("/login/submit", handler.Login)
	app.Post("/regist/submit", handler.RegistUser)
	app.Post("/modify_pass/submit", handler.UpdatePassword, handler.Auth)
	app.Get("/user", handler.GetUserInfo)
	app.Get("/logout", handler.Logout)

	group := app.Group("/news")
	group.Get("", handler.NewsList)
	group.Get("/issue", func(ctx fiber.Ctx) error { return ctx.Status(http.StatusOK).Render("news_issue", nil) })
	group.Post("/issue/submit", handler.PostNews, handler.Auth)
	group.Get("/belong", handler.NewsBelong)
	group.Get("/:id", handler.GetNewsById)
	group.Get("/delete/:id", handler.DeleteNews, handler.Auth)
	group.Post("/update", handler.UpdateNews, handler.Auth)

	app.Get("", func(ctx fiber.Ctx) error { return ctx.Status(http.StatusMovedPermanently).Redirect().To("/news") }) //新闻列表页是默认的首页

	if err := app.Listen("127.0.0.1:5678"); err != nil {
		panic(err)
	}
}

// go run ./Post
// 在浏览器里访问 http://localhost:5678/
