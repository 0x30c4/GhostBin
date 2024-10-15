package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0x30c4/ghostbin/pkg/env"
	"github.com/0x30c4/ghostbin/pkg/logger"
	"github.com/0x30c4/ghostbin/pkg/middleware"
	"github.com/0x30c4/ghostbin/pkg/redis"
	"github.com/0x30c4/ghostbin/pkg/routes"
)

func main() {

  env.EnvInit()
	router := routes.Router()

  host := env.HOST
  port := env.PORT

  address := fmt.Sprintf("%s:%s", host, port)

  fmt.Printf("Server listening on http://%s\n", address)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  err := redis.InitRedis()

  if err != nil {
    log.Fatal(err)
  }

  stack := middleware.CreateStack(
    logger.LogRequest,
  )

  server := http.Server{
    Addr: address,
    Handler: stack(router),
  }

  err = server.ListenAndServe()

  if err != nil {
    panic(err)
  }
}
