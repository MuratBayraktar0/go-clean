package infrastructure

import (
	"context"
	"golang-clean/domain"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	collection *mongo.Collection
}

func NewMongodb(ctx context.Context, uri, database, collection string) *Mongodb {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(database)
	cl := db.Collection(collection)

	return &Mongodb{
		collection: cl,
	}
}

func (u *Mongodb) Get(ctx context.Context, id string) (*domain.UserEntity, error) {
	var result domain.UserEntity
	filter := bson.M{"_id": id}
	err := u.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *Mongodb) GetAll(ctx context.Context) (*[]domain.UserEntity, error) {
	var results []domain.UserEntity
	filter := bson.M{}
	cursor, err := u.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user domain.UserEntity
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (u *Mongodb) Create(ctx context.Context, user domain.UserEntity) error {
	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *Mongodb) Update(ctx context.Context, user domain.UserEntity) error {
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email, "updated_at": user.UpdatedAt}}
	_, err := u.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
