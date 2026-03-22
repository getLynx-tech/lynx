package middleware

import (
	"github.com/getLynx-tech/lynx/internal/conf"
	"github.com/gin-gonic/gin"
)

type BasicAuthMiddleware struct {
	cfg *conf.Config
}

func NewBasicAuthMiddleware(cfg *conf.Config) *BasicAuthMiddleware {
	return &BasicAuthMiddleware{cfg: cfg}
}

func (b *BasicAuthMiddleware) WithBasicAuth() gin.HandlerFunc {
	user := b.cfg.BasicAuthUser
	password := b.cfg.BasicAuthPassword

	return gin.BasicAuth(gin.Accounts{
		user: password,
	})
}
