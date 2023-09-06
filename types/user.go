package types

type User struct {
	AccentColor      int         `json:"accent_color" bson:"accent_color"`
	Avatar           string      `json:"avatar"`
	AvatarDecoration interface{} `json:"avatar_decoration" bson:"avatar_decoration"`
	Banner           interface{} `json:"banner"`
	BannerColor      string      `json:"banner_color" bson:"banner_color"`
	Flags            int         `json:"flags"`
	Locale           string      `json:"locale"`
	MfaEnabled       bool        `json:"mfa_enabled"`
	PremiumType      int         `json:"premium_type" bson:"premium_type"`
	PublicFlags      int         `json:"public_flags" bson:"public_flags"`
	Token            string      `json:"token"`
	AccessToken      string      `json:"access_token" bson:"access_token"`
	IsBanned         bool        `json:"is_banned" bson:"is_banned"`
	IsAdmin          bool        `json:"is_admin" bson:"is_admin"`
	AppID            interface{} `json:"appId" bson:"appId"`
	Entity           interface{} `json:"entity" bson:"entity"`
	ID               string      `json:"id" description:"The users ID"`
	Username         string      `json:"username"`
	DisplayName      string          `json:"display_name"`
}
