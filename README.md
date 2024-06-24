# My Go Project 
## using Go - Gin framework

#### Mockery
- **Used** for generating mocks for testing
- **Installation**: 
  - Following the instructions [here](https://vektra.github.io/mockery/v2.32/installation/)
    ```bash
    go install github.com/vektra/mockery/v2@v2.32.0
    ```
- **Usage:**
  - Add `//go:generate mockery --name=<InterfaceName>` above the interface declaration
  - Run `go generate ./...` in the *root directory* of the project


#### Using make in cli 
    -- make run-api : for run api 
    -- make run-consumer: for run consumer
    -- make swagger : for generate swagger documents
    -- make build : for build docker compose
    -- make run : to run service
    -- make stop: to stop service
### Generate jwt token env
    -- get JWT_SECERT -> cli run : node -e "console.log(require('crypto').randomBytes(32).toString('hex'))" 

### Docker compose 
    -- Create folder mongodb_data for mount database
    -- Include Rabbitmq,Mongodb in docker-compose

### Database setup
    -- In env have ENCRYPT_KEY , using is for encrypt uri string of database (key size = 32)    
    -- mkdir ./mongo_data && sudo chmod 755 -R mongo_data (this for mount data from docker-compose)

### Swagger 
    -- access to http://localhost:8088/swagger/index.html to see swagger
    -- or access to domain https://t.hoangdeptrai.online/musics/swagger/index.html



### API LIST 
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/music-tracks      --> github.com/quachhoang2002/Music-Library/internal/music/delivery/http.Handler.List-fm (4 handlers)
[GIN-debug] GET    /api/v1/music-tracks/:id  --> github.com/quachhoang2002/Music-Library/internal/music/delivery/http.Handler.Detail-fm (4 handlers)
[GIN-debug] POST   /api/v1/music-tracks      --> github.com/quachhoang2002/Music-Library/internal/music/delivery/http.Handler.Create-fm (4 handlers)
[GIN-debug] PUT    /api/v1/music-tracks/:id  --> github.com/quachhoang2002/Music-Library/internal/music/delivery/http.Handler.Update-fm (4 handlers)
[GIN-debug] DELETE /api/v1/music-tracks/:id  --> github.com/quachhoang2002/Music-Library/internal/music/delivery/http.Handler.Delete-fm (4 handlers)
[GIN-debug] GET    /api/v1/music-tracks/:id/file --> github.com/quachhoang2002/Music-Library/internal/music/delivery/http.Handler.GetFile-fm (4 handlers) // get file

[GIN-debug] GET    /api/v1/playlists/:user_id --> github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http.Handler.List-fm (4 handlers)
[GIN-debug] POST   /api/v1/playlists/:user_id --> github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http.Handler.Create-fm (4 handlers)
[GIN-debug] PUT    /api/v1/playlists/:user_id/:id --> github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http.Handler.Update-fm (4 handlers)
[GIN-debug] DELETE /api/v1/playlists/:user_id/:id --> github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http.Handler.Delete-fm (4 handlers)
[GIN-debug] GET    /api/v1/playlists/:user_id/:id --> github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http.Handler.Detail-fm (4 handlers)
[GIN-debug] POST   /api/v1/playlists/:user_id/:id/tracks/:track_id --> github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http.Handler.AddTrack-fm (4 handlers)
[GIN-debug] DELETE /api/v1/playlists/:user_id/:id/tracks/:track_id --> github.com/quachhoang2002/Music-Library/internal/playlist/delivery/http.Handler.RemoveTrack-fm (4 handlers)

