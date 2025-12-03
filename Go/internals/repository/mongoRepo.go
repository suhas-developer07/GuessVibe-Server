package repository

import (
	"context"
	"time"

	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/User_model"
	"github.com/suhas-developer07/GuessVibe-Server/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// MongoRepo struct definition
type MongoRepo struct {
	db  *mongo.Database
	ctx context.Context
}

func (r *MongoRepo) RegisterUser(user models.User) (int64, error) {
	collection := r.db.Collection("users")
	Hashedpassword, err := utils.Hashedpassword(user.Password)
	if err != nil {
		return 0, err
	}

	objectID := primitive.NewObjectID()
	user.ID = objectID.Hex()
	user.UserID = user.ID
	user.Password = Hashedpassword

	result, err := collection.InsertOne(r.ctx, user)
	if err != nil {
		return 0, err
	}
	insertedID := result.InsertedID.(int64)
	return insertedID, nil
}
func (r *MongoRepo) LoginUser(Email, password string) (string, error) {
	collection := r.db.Collection("users")
	var user models.User
	err := collection.FindOne(r.ctx, map[string]interface{}{"email": Email}).Decode(&user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		return "", err
	}
	collection.UpdateOne(r.ctx, bson.M{"email": Email}, bson.M{"$set": bson.M{"token": token}})
	return token, nil
}
func (r *MongoRepo) LogoutUser(UserID, token string) error {
	collection := r.db.Collection("users")
	updateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateData := bson.M{
		"$set": bson.M{
			"token":     token,
			"updatedat": updateAt,
		},
	}
	_, err := collection.UpdateOne(r.ctx, bson.M{"userid": UserID}, updateData)
	if err != nil {
		return err
	}
	return nil
}
