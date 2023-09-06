package configuration

import (
	"go.topiclist.xyz/types"
)

func getConfig() types.Config {
	return types.Config{
		ApiVersion: 1,
		Database: types.Database{
			Url: "mongodb+srv://Admin:RanveerSoni11@topic.q8qcpfz.mongodb.net/discordinflux",
		},
		Web: types.Web{
			Port:           "8080",
			ImageUploadKey: "d6b358da96fd1e63b3eadc17fbae7bdb",
			ReturnUrl:      "https://beta.discordinflux.xyz",
		},
		Client: types.Client{
			Id:       "925740376948609034",
			Secret:   "vZZ5V5kki93ECcaB9RegynGA8Itua4vg",
			Token:    "OTI1NzQwMzc2OTQ4NjA5MDM0.GKYieQ.WUhkXZZQ2cAIDkxaG0DhmnpwnByqxjIRycNEpc",
			Callback: "https://beta.discordinflux.xyz/v1/auth/callback",
		},
		Collection: "entities",
		APIUrl:     "https://beta.discordinflux.xyz/v1",
	}
}

func GetConfig() types.Config {
	return getConfig()
}
