package lib

import (
	"context"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/theovidal/onyxcord"
)

// Router is the structure holding the multiple channels together
type Router struct {
	ID          string `bson:"_id" json:"_id"`
	Token       string `json:"token"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Plugs       []Plug `json:"plugs"`
}

func CreateRouter(bot *onyxcord.Bot, name, description, channel string) (err error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72000).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Carrbridge",
		Subject:   name,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))

	collection := bot.Database.Database(bot.Config.Database.Database).Collection(bot.Config.Assets["routers_collection"])

	var result *mongo.InsertOneResult
	result, err = collection.InsertOne(
		context.Background(),
		Router{
			Token:       tokenString,
			Name:        name,
			Description: description,
			Plugs:       nil,
		},
	)
	if err != nil {
		return
	}

	var router Router
	err = collection.FindOne(context.Background(), bson.D{{"_id", result.InsertedID}}).Decode(&router)
	if err != nil {
		return
	}

	err = CreatePlug(bot, &router, channel)
	return
}

// GetRoutersFromChannel returns all the routers linked to a channel
func GetRoutersFromChannel(bot *onyxcord.Bot, channel string) (routers []Router, err error) {
	collection := bot.Database.Database(bot.Config.Database.Database).Collection(bot.Config.Assets["routers_collection"])
	var cursors *mongo.Cursor
	cursors, err = collection.Find(context.Background(), bson.M{
		"plugs": bson.M{
			"$elemMatch": bson.M{
				"channel": channel,
			},
		},
	})
	if err != nil {
		return
	}

	err = cursors.All(context.Background(), &routers)
	return
}

func GetRouterFromToken(bot *onyxcord.Bot, token string) (router Router, err error) {
	err = bot.Database.Database(bot.Config.Database.Database).Collection(bot.Config.Assets["routers_collection"]).FindOne(
		context.Background(),
		bson.M{"token": token},
	).Decode(&router)

	return
}
