package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanhoang/music-library/pkg/mongo"
	"github.com/xuanhoang/music-library/pkg/rabbitmq"
	"github.com/xuanhoang/music-library/pkg/redis"

	pkgLog "github.com/xuanhoang/music-library/pkg/log"
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
	microservice Microservice
	telegram     TeleCredentials
}
type Config struct {
	Port         int
	JWTSecretKey string
	Database     mongo.Database
	Mode         string
	AMQPConn     *rabbitmq.Connection
	Redis        redis.Client
	Microservice Microservice
	Telegram     TeleCredentials
}
type Microservice struct {
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
		microservice: cfg.Microservice,
		telegram:     cfg.Telegram,
	}
}
