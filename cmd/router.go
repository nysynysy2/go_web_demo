package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nysynysy2/go_web_demo/internal/dept"
	"github.com/nysynysy2/go_web_demo/internal/env"
)

func Run() {
	r := gin.Default()
	dept.Routes(r)
	addr := env.ReadOrFatal("ADDR")
	r.Run(addr)
}
