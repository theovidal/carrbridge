package lib

import (
	"context"

	"github.com/theovidal/onyxcord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Router is the structure holding the multiple channels together
type Router struct {
	Token       string `json:"token"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Plugs       []Plug `json:"plugs"`
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

// Plug is part of a Router and is more specifically an ending of this router
type Plug struct {
	Channel string `json:"channel"`
	Webhook string `json:"webhook"`
	Token   string `json:"token"`
}
