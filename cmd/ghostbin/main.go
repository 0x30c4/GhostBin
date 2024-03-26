package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0x30c4/ghostbin/internal/env"
	"github.com/0x30c4/ghostbin/internal/logger"
	"github.com/0x30c4/ghostbin/internal/middleware"
	"github.com/0x30c4/ghostbin/internal/redis"
	"github.com/0x30c4/ghostbin/internal/routes"
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
