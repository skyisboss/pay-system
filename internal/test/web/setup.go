package web

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/cmd"
	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/skyisboss/pay-system/internal/config"
)

type RespType struct {
	Rows    gin.H  `json:"rows"`
	Data    gin.H  `json:"data"`
	Err     int    `json:"err"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

func Setup() (context.Context, *config.Config, *boot.Boot, *zerolog.Logger) {
	var (
		ctx    = context.Background()
		cfg    = cmd.RegisterConfig()
		boot   = boot.New(ctx, cfg)
		logger = boot.Logger()
	)

	return ctx, cfg, boot, logger
}

func UseRequest(r http.Handler, method, path string, playload io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, playload)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
