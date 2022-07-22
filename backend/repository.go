package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/greetings/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func (repository *Repository) CreateBlog(blog models.Blog) error {
	collection := repository.client.Database("blog").Collection("blog")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, blog)

	if err != nil {
		return err
	}

	return nil

}

func NewRepository() *Repository {
	uri := "mongodb+srv://Cluster:cluster@cluster0.hnmuy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	/* uri := "mongodb+srv://Cluster:cluster@cluster0.hnmuy.mongodb.net/test?authSource=admin&replicaSet=atlas-jbznje-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true" */
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func NewTestRepository() *Repository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func (repository *Repository) GetBlogs() ([]models.Blog, error) {
	collection := repository.client.Database("blog").Collection("blog")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	blogs := []models.Blog{}
	for cur.Next(ctx) {
		var blog models.Blog
		err := cur.Decode(&blog)
		if err != nil {
			log.Fatal(err)
		}
		//go'da ekleme append ile yapılır
		blogs = append(blogs, blog)
	}

	return blogs, nil

}

func (repository *Repository) PostBlogs(blog models.Blog) error {
	collection := repository.client.Database("blog").Collection("blog")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, blog)

	if err != nil {
		return err
	}

	return nil

}

func (repository *Repository) DeleteBlogs(blogId string) error {
	collection := repository.client.Database("blog").Collection("blog")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": blogId}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}

func (repository *Repository) UpdateBlogs(blog models.BlogDTO, ID string) (models.Blog, error) {
	collection := repository.client.Database("blog").Collection("blog")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	updateBlog := bson.M{"id": ID,
		"title":       blog.Title,
		"description": blog.Description,
	}

	_, err := collection.ReplaceOne(ctx, bson.M{"id": ID}, updateBlog)

	if err != nil {
		return models.Blog{}, err
	}
	updatedBlog, err := repository.GetBlog(ID)

	if err != nil {
		return models.Blog{}, err
	}

	return updatedBlog, nil

}

func (repository *Repository) GetBlog(ID string) (models.Blog, error) {
	collection := repository.client.Database("blog").Collection("blog")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	blog := models.Blog{}
	if err := collection.FindOne(ctx, bson.M{}).Decode(&blog); err != nil {
		log.Fatal(err)
	}

	fmt.Println("dsafdfsafd", blog)

	return blog, nil

}

func GetCleanTestRepository() *Repository {

	repository := NewRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	blogDB := repository.client.Database("blog")
	blogDB.Drop(ctx)

	return repository
}
