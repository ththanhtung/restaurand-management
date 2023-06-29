package helpers

import (
	"context"
	"log"
	"mongotest/database"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignDetail struct {
	Email  string
	UserID string
	jwt.StandardClaims
}

var jwtSecret string = os.Getenv("JWT_KEY")

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func GenerateToken(email, userID string) (token string, err error) {
	claims := &SignDetail{
		Email:  email,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))

	if err != nil {
		panic(err)
	}

	return token, err
}

func UpdateToken(userID, signedToken string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var updatedObj primitive.D

	updatedObj = append(updatedObj, bson.E{"token", signedToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updatedObj = append(updatedObj, bson.E{"updated_at", updatedAt})

	upsert := true
	filter := bson.M{"userid": userID}
	options := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{"$set", updatedObj}}, &options)

	defer cancel()

	if err != nil {
		log.Panic(err)
	}
}

func VerifyToken(signedToken string) (claims *SignDetail, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, &SignDetail{}, func(t *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil })

	claims, ok := token.Claims.(*SignDetail)

	if err != nil {
		log.Panic(err)
		return
	}

	if !ok {
		msg = "token is invalid"
		return nil, msg
	}

	if claims.ExpiresAt < time.Now().Unix(){
		msg = "token is expired"
		return nil, msg
	}

	return claims, msg
}
