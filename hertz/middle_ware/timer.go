package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// 业务Handler以及Server端中间件都要满足此函数签名
func TimeMW(ctx context.Context, c *app.RequestContext) {
	begin := time.Now()
	c.Next(ctx)
	// c.Abort()   //终止后续的Handler。或者使用AbortWithMsg、AbortWithStatusJSON
	hlog.Infof("interface %s use time %d ms", c.Path(), time.Since(begin).Milliseconds())
}
