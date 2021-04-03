module github.com/theovidal/carrbridge

go 1.13

require (
	github.com/bwmarrin/discordgo v0.23.3-0.20210327033043-f637c37ba2f0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/theovidal/onyxcord v0.0.0
	go.mongodb.org/mongo-driver v1.4.3
)

replace github.com/theovidal/onyxcord => ../onyxcord
