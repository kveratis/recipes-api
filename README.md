# recipes-api

Gin Recipes Api

## Generate Swagger

```cli
swag init
```

[Swagger Docs](http://localhost:8080/swagger/index.html)

## Databases

```cli
docker-compose up -d
```

### MongoDB

For initial setup, go into Mongo Compass and create a database called **demo** with a collection called **recipes**. Then import the recipes.json file into that collection.

mongodb://root:password@localhost:27017/