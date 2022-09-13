package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	_ "recipes-api/docs"
	"time"
)

type Recipe struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Tags         []string           `json:"tags" bson:"tags"`
	Ingredients  []string           `json:"ingredients" bson:"ingredients"`
	Instructions []string           `json:"instructions" bson:"instructions"`
	PublishedAt  time.Time          `json:"publishedAt" bson:"publishedAt"`
}

type Message struct {
	Description string `json:"message,omitempty"`
	Error       string `json:"error,omitempty"`
}

var ctx context.Context
var collection *mongo.Collection
var err error
var client *mongo.Client

func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to MongoDB")
	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
}

// ListRecipesHandler godoc
// @Summary get list of recipes
// @Produce json
// @Success	200 {array} Recipe "Successful operation"
// @Failure 500 {object} Message "Server Error"
// @Router	/recipes [get]
func ListRecipesHandler(c *gin.Context) {
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		message := Message{Error: err.Error()}
		c.JSON(http.StatusInternalServerError, message)
		return
	}
	defer cur.Close(ctx)

	recipes := make([]Recipe, 0)
	for cur.Next(ctx) {
		var recipe Recipe
		cur.Decode(&recipe)
		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, recipes)
}

// GetOneRecipeHandler godoc
// @Summary get one recipe
// @Produce json
// @Param id path string true "ID of the recipe"
// @Success 200 {object} Recipe "Successful operation"
// @Failure 404 {object} Message "Invalid recipe ID"
// @Failure 500 {object} Message "Server Error"
// @Router /recipes/{id} [get]
func GetOneRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		message := Message{Error: "Recipe not found"}
		c.JSON(http.StatusNotFound, message)
		return
	}

	cur := collection.FindOne(ctx, bson.M{
		"_id": objectId,
	})
	var recipe Recipe
	err = cur.Decode(&recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// NewRecipeHandler godoc
// @Summary	Create a new recipe
// @Accept json
// @Produce json
// @Param recipe body Recipe true "Recipe to add"
// @Success 200 {object} Recipe "Successful operation"
// @Failure 400 {object} Message "Invalid input"
// @Failure 500 {object} Message "Server Error"
// @Router /recipes [post]
func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		message := Message{Error: err.Error()}
		c.JSON(http.StatusBadRequest, message)
		return
	}
	recipe.ID = primitive.NewObjectID()
	recipe.PublishedAt = time.Now()
	_, err = collection.InsertOne(ctx, recipe)
	if err != nil {
		fmt.Println(err)
		message := Message{Error: "Error while inserting a new recipe"}
		c.JSON(http.StatusInternalServerError, message)
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// UpdateRecipeHandler godoc
// @Summary update an existing recipe
// @Accept json
// @Produce json
// @Param id path string true "ID of the recipe"
// @Param recipe body Recipe true "Updated recipe"
// @Success 200 {object} Message "Successful operation"
// @Failure 400 {object} Message "Invalid input"
// @Failure 404 {object} Message "Invalid recipe ID"
// @Failure 500 {object} Message "Server Error"
// @Router /recipes/{id} [put]
func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		message := Message{Error: "Recipe not found"}
		c.JSON(http.StatusNotFound, message)
		return
	}

	recipe := Recipe{ID: objectId}
	if err := c.ShouldBindJSON(&recipe); err != nil {
		message := Message{Error: err.Error()}
		c.JSON(http.StatusBadRequest, message)
		return
	}

	_, err = collection.UpdateOne(ctx, bson.M{
		"_id": objectId,
	}, bson.D{{"$set", bson.D{
		{"name", recipe.Name},
		{"instructions", recipe.Instructions},
		{"ingredients", recipe.Ingredients},
		{"tags", recipe.Tags},
	}}})

	if err != nil {
		fmt.Println(err)
		message := Message{Error: err.Error()}
		c.JSON(http.StatusInternalServerError, message)
		return
	}

	message := Message{Description: "Recipe has been updated"}
	c.JSON(http.StatusOK, message)
}

// DeleteRecipeHandler godoc
// @Summary delete an existing recipe
// @Produce json
// @Param id path string true "ID of the recipe"
// @Success 200 {object} Message "Successful operation"
// @Failure 404 {object} Message "Invalid recipe ID"
// @Failure 500 {object} Message "Server Error"
// @Router /recipes/{id} [delete]
func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		message := Message{Error: "Recipe not found"}
		c.JSON(http.StatusNotFound, message)
		return
	}

	_, err = collection.DeleteOne(ctx, bson.M{
		"_id": objectId,
	})
	if err != nil {
		message := Message{Error: err.Error()}
		c.JSON(http.StatusInternalServerError, message)
		return
	}

	message := Message{Description: "Recipe has been deleted"}
	c.JSON(http.StatusOK, message)
}

// SearchRecipesHandler godoc
// @Summary search for recipe by tag
// @Produce json
// @Param tag query string true "Tag of the recipes"
// @Success 200 {array} Recipe "Successful operation"
// @Router /recipes/search [get]
//func SearchRecipesHandler(c *gin.Context) {
//	tag := c.Query("tag")
//	listOfRecipes := make([]Recipe, 0)
//
//	for i := 0; i < len(recipes); i++ {
//		found := false
//		for _, t := range recipes[i].Tags {
//			if strings.EqualFold(t, tag) {
//				found = true
//			}
//		}
//		if found {
//			listOfRecipes = append(listOfRecipes, recipes[i])
//		}
//	}
//
//	c.JSON(http.StatusOK, listOfRecipes)
//}

// @title Recipes API
// @version 1.0.0
// @description This is a sample recipes api.
// @contact.name Daniel Petersen
// @schemes http
// @host localhost:8080
// @BasePath /
// @accept json
// @produce json
func main() {
	router := gin.Default()
	router.GET("/recipes", ListRecipesHandler)
	router.GET("/recipes/:id", GetOneRecipeHandler)
	router.POST("/recipes", NewRecipeHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	//router.GET("/recipes/search", SearchRecipesHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run()
	if err != nil {
		return
	}
}
