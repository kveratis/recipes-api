# Initial Setup
```cli
mkdir recipes-api
cd recipes-api
go mod init recipes-api
```
## Add Gin

```cli
go get -u github.com/gin-gonic/gin
```

## Add XID for unique IDs

```cli
go get github.com/rs/xid
go mod tidy
```

## Add Swag for OpenApi/Swagger Docs

```cli
go install github.com/swaggo/swag/cmd/swag@latest
```

## Add MongoDB Drivers

```cli
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
```