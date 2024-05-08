# My Go Project

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
    -- mkdir ./mongo_data && sudo chmod 755 -R mongo_data

### 
    -- access to http://localhost:8088/swagger/index.html to see swagger