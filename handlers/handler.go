package handlers

import (
	"fmt"

	"github.com/SudhanAruna/dist-go-recipes-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipesHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewRecipesHandler(ctx context.Context, collection *mango.Collection) *RecipesHandler {
	return &RecipesHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *RecipesHandler) ListRecipesHandler(ctx *gin.Context) {
	cursor, err := handler.collection.Find(
		handler.ctx,
		bson.M{},
	)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	defer cursor.Close(handler.ctx)

	recipes := make([]models.Recipe, 0)
	for cursor.Next(handler.ctx) {
		var recipe models.Recipe
		cursor.Decode(&recipe)
		recipes = append(recipes, recipe)
	}
	ctx.JSON(http.StatusOK, recipes)
}
