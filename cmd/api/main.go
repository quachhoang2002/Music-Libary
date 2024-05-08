package main

import (
	"github.com/xuanhoang/music-library/config"
	"github.com/xuanhoang/music-library/internal/appconfig/mongo"
	"github.com/xuanhoang/music-library/internal/appconfig/redis"
	"github.com/xuanhoang/music-library/internal/httpserver"
	pkgCrt "github.com/xuanhoang/music-library/pkg/encrypter"
	pkgLog "github.com/xuanhoang/music-library/pkg/log"
	"github.com/xuanhoang/music-library/pkg/rabbitmq"
)

// @title Mucsic Library API
// @description This is the API documentation for the Music Library.
// @description Error codes: 0000 - 1000("Music Track Error")
// @description `0001`(Wrong pagination query)
// @description `0002`(Invalid body)
// @description `0003`(Invalid form data)
// @description `0004`(Invalid params query)
// @description `0005`(Invalid validation)
// @description `0006`(Unauthorized)
// @description `0007`(Music track not found)

// @description Error codes: 1001 - 2000("User Error")

// @version 1
// @host localhost:8088
// @BasePath /
// @schemes http
func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	crp := pkgCrt.NewEncrypter(cfg.Encrypter.Key)

	client, err := mongo.Connect(cfg.Mongo, crp)
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(client)

	database := client.Database(cfg.Mongo.Database)

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	amqpConn, err := rabbitmq.Dial(cfg.RabbitMQConfig.URL, true)
	if err != nil {
		panic(err)
	}
	defer amqpConn.Close()

	redisClient, err := redis.Connect(cfg.RedisConfig)
	if err != nil {
		panic(err)
	}
	defer redisClient.Disconnect()

	srv := httpserver.New(l, httpserver.Config{
		Port:         cfg.HTTPServer.Port,
		Database:     database,
		JWTSecretKey: cfg.JWT.SecretKey,
		Mode:         cfg.HTTPServer.Mode,
		AMQPConn:     amqpConn,
		Redis:        redisClient,
		Microservice: httpserver.Microservice{},
		Telegram: httpserver.TeleCredentials{
			BotKey: cfg.Telegram.BotKey,
			ChatIDs: httpserver.ChatIDs{
				ReportBug: cfg.Telegram.ChatIDs.ReportBug,
			},
		},
	})

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
