package main

import (
	"os"

	"github.com/0x30c4/ghostbin/internal/repositories"
	"github.com/0x30c4/ghostbin/internal/services"
	"github.com/0x30c4/ghostbin/pkg/log"
	"github.com/0x30c4/ghostbin/pkg/redis"
	"github.com/0x30c4/ghostbin/pkg/webappsrv"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/exp/slog"
)

var (
  uploadDir       = os.Getenv("PASTE_UPLOAD_DIR")
  maxFileSize     = os.Getenv("MAX_FILE_SIZE")

	domain          = os.Getenv("DOMAIN")
	host            = os.Getenv("HOST")
	port            = os.Getenv("PORT")

	logFilePath     = os.Getenv("LOG_FILE_PATH")
	logAsTextStr    = os.Getenv("LOG_AS_TEXT")
)

func main() {
  var logAsText bool = true
  if logAsTextStr == "false" {
    logAsText = false
  }

  logger := log.NewLogger(logFilePath, logAsText)
	logger.Info("starting service",
		slog.String("release_mode", os.Getenv("RELEASE_MODE")),
	)

  redisConn, err := redis.NewConnectionFromEnvVar()

  if err != nil {
		logger.Error("failed to connect to redis database", slog.String("err", err.Error()))
		return
	}
	defer func() {
    redisConn.Close()
    logger.Info("Redis connection is closing")
  }()

	logger.Info("Redis connection established",
		slog.String("address", os.Getenv("REDIS_ADDRESS")),
	)

  repo := repositories.NewPasteRepository(redisConn)
  pasteSrv := services.NewPasteService(repo)
  fileSrv := services.NewFileService(repo, uploadDir, maxFileSize)

  logger.Info("Repositories & Services are initialized")

  webappsrv.RunServer(pasteSrv, fileSrv, logger)

  logger.Info("Web Server is running",
    slog.String("domain", domain),
    slog.String("host", host),
    slog.String("port", port),
  )
}
