package httpserver

import (
	"context"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Import this to execute the init function in docs.go which setups the Swagger docs.
	_ "github.com/quachhoang2002/Music-Library/docs"
	"github.com/quachhoang2002/Music-Library/pkg/jwt"

	musicHTTP "github.com/quachhoang2002/Music-Library/internal/music/delivery/http"
	musicRepo "github.com/quachhoang2002/Music-Library/internal/music/repository/mongo"
	musicUC "github.com/quachhoang2002/Music-Library/internal/music/usecase"

	playlistHTTP "github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http"
	playlistRepo "github.com/quachhoang2002/Music-Library/internal/playlist/repository/mongo"
	playlistUC "github.com/quachhoang2002/Music-Library/internal/playlist/usecase"

	"github.com/quachhoang2002/Music-Library/internal/middleware"
)

func (srv HTTPServer) mapHandlers() error {
	srv.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jwtManager, err := jwt.NewJWTMaker(srv.jwtSecretKey)
	if err != nil {
		srv.l.Fatal(context.Background(), err)
		return err
	}

	// Middleware
	mw := middleware.New(srv.l, jwtManager)

	// Repositories
	musicRepo := musicRepo.New(srv.l, srv.database)
	playlistRepo := playlistRepo.New(srv.l, srv.database)

	// UseCases
	musicUC := musicUC.New(srv.l, musicRepo)
	playlistUC := playlistUC.New(srv.l, playlistRepo, musicUC)

	// Handlers
	musicH := musicHTTP.New(srv.l, musicUC)
	playlistH := playlistHTTP.New(srv.l, playlistUC)

	// External routes
	externalAPI := srv.gin.Group("/api/v1")
	musicHTTP.MapMusicTrackRoutes(externalAPI.Group("/music-tracks"), musicH, mw)
	playlistHTTP.MapMusicPlaylistRoutes(externalAPI.Group("/playlists"), playlistH, mw)

	// Internal routes
	// internalAPI := srv.gin.Group("/internal/api/v1")

	return nil
}
