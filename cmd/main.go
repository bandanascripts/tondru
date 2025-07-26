package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bandanascripts/tondru/pkg/core"
	"github.com/bandanascripts/tondru/pkg/server/routes"
	"github.com/bandanascripts/tondru/pkg/service/redis"
	"github.com/gin-gonic/gin"
)

func init() {
	redis.Connect()
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	core.GenAndStoreKey(ctx, "RSA:PRIVATEKEY:", "RSA:PUBLICKEY:", 3600)
	
	var r = gin.Default()
	routes.RegisteredRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server at port 8080 : %v", err)
	}
}
