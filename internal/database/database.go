package database

import (
	"chi_pos_counter_service_api/internal/server/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	Health() map[string]string
	GetIngredientById(id string) (model.Ingredient, error)
	GetIngredients() ([]model.Ingredient, error)
	GetStoreById(id string) (model.Store, error)
	GetEmployeeById(id string) (model.Employee, error)
}

type service struct {
	db *mongo.Client
}

var (
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	//database = os.Getenv("DB_DATABASE")
)

func New() Service {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)))

	if err != nil {
		log.Fatal(err)

	}
	return &service{
		db: client,
	}
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetIngredients() ([]model.Ingredient, error) {
	var ingredients []model.Ingredient
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	collection := s.db.Database("test").Collection("ingredients")
	cursor, err := collection.Find(ctx, bson.M{})
	fmt.Println(cursor)
	if err != nil {
		log.Fatalf(fmt.Sprintf("error getting ingredients: %v", err))
	}
	for cursor.Next(ctx) {
		var ingredient model.Ingredient
		cursor.Decode(&ingredient)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func (s *service) GetIngredientById(id string) (model.Ingredient, error) {
	var ingredient model.Ingredient
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	collection := s.db.Database("test").Collection("ingredients")
	fmt.Printf(id)
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("error converting id to object id: %v", err)
		return ingredient, err
	}
	fmt.Println(_id)

	filter := bson.M{
		"_id": _id,
	}
	fmt.Println(collection.FindOne(ctx, filter))
	collection.FindOne(ctx, filter).Decode(&ingredient)

	return ingredient, nil
}

func (s *service) GetStoreById(id string) (model.Store, error) {
	var store model.Store
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := s.db.Database("test").Collection("stores")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("error converting id to object id: %v", err)
		return store, err
	}
	filter := bson.M{
		"_id": _id,
	}
	collection.FindOne(ctx, filter).Decode(&store)
	return store, nil
}

func (s *service) GetEmployeeById(id string) (model.Employee, error) {
	var employee model.Employee
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := s.db.Database("test").Collection("employees")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("error converting id to object id: %v", err)
		return employee, err
	}
	filter := bson.M{
		"_id": _id,
	}
	collection.FindOne(ctx, filter).Decode(&employee)
	return employee, nil
}
