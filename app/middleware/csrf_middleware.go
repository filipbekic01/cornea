package middleware

import (
	"github.com/kataras/iris/v12"
)

// GeneralMiddleware .
func CSRFMiddleware(ctx iris.Context) {

	// TODO: Implement CSRF token for MVC requests.
	ctx.Next()
}
