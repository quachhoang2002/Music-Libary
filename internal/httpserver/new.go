package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/quachhoang2002/Music-Library/pkg/mongo"
	"github.com/quachhoang2002/Music-Library/pkg/rabbitmq"
	"github.com/quachhoang2002/Music-Library/pkg/redis"

	pkgLog "github.com/quachhoang2002/Music-Library/pkg/log"
)

const productionMode = "production"

var ginMode = gin.DebugMode

type HTTPServer struct {
	gin          *gin.Engine
	l            pkgLog.Logger
	port         int
	database     mongo.Database
	jwtSecretKey string
	mode         string
	amqpConn     *rabbitmq.Connection
	redis        redis.Client
	telegram     TeleCredentials
}
type Config struct {
	Port         int
	JWTSecretKey string
	Database     mongo.Database
	Mode         string
	AMQPConn     *rabbitmq.Connection
	Redis        redis.Client
	Telegram     TeleCredentials
}

type TeleCredentials struct {
	BotKey string
	ChatIDs
}

type ChatIDs struct {
	ReportBug int64
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	if cfg.Mode == productionMode {
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)

	return &HTTPServer{
		l:            l,
		gin:          gin.Default(),
		port:         cfg.Port,
		database:     cfg.Database,
		jwtSecretKey: cfg.JWTSecretKey,
		mode:         cfg.Mode,
		amqpConn:     cfg.AMQPConn,
		redis:        cfg.Redis,
		telegram:     cfg.Telegram,
	}
}
