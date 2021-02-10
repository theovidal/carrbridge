package lib

import (
	"context"
	"github.com/theovidal/onyxcord"
	"go.mongodb.org/mongo-driver/bson"
)

// Plug is part of a Router and is more specifically an ending of this router
type Plug struct {
	Channel string `json:"channel"`
	Webhook string `json:"webhook"`
	Token   string `json:"token"`
}

func CreatePlug(bot *onyxcord.Bot, router *Router, channel string) (err error) {
	webhook, err := bot.Client.WebhookCreate(channel, "Carrbridge Plug", "")
	if err != nil {
		return
	}

	collection := bot.Database.Database(bot.Config.Database.Database).
		Collection(bot.Config.Assets["routers_collection"])

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{
			"token": router.Token,
		},
		bson.M{
			"$addToSet": bson.M{
				"plugs": bson.M{
					"channel": channel,
					"webhook": webhook.ID,
					"token":   webhook.Token,
				},
			},
		},
	)
	return
}
