package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	DisplayName      string      `json:"display_name"`
}

type Partner struct {
	Title      string `json:"title"`
	Logo       string `json:"logo"`
	Banner     string `json:"banner"`
	OwnerName  string `json:"ownername"`
	OwnerID    string `json:"ownerid"`
	Desc       string `json:"desc"`
	Link1      string `json:"link1"`
	Link1Title string `json:"link1title"`
	Link1Icon  string `json:"link1icon"`
	Link2      string `json:"link2"`
	Link2Title string `json:"link2title"`
	Link2Icon  string `json:"link2icon"`
}

type Entity struct {
	ID      string `json:"id"`
	Discord struct {
		ID          string `json:"id" description:"The users ID"`
		Username    string `json:"username" description:"The users username"`
		DisplayName string `json:"display_name"`
	} `json:"discord"`
	URL        string      `json:"url"`
	Banner     string      `json:"banner"`
	Avatar     string      `json:"avatar"`
	About      interface{} `json:"about"`
	Occupation interface{} `json:"occupation"`
	Staff      string      `json:"staff"`
	Birthday   interface{} `json:"birthday"`
	Location   interface{} `json:"location"`
	Gender     interface{} `json:"gender"`
	Pronouns   interface{} `json:"pronouns"`
	Language   string      `json:"language"`
	Website    interface{} `json:"website"`
	Like       int         `json:"like"`
	Email      interface{} `json:"email"`
	Views      []string    `json:"views"`
	Premium    bool        `json:"isPremium" bson:"isPremium"`
	Verified   bool        `json:"isVerified" bson:"isVerified"`
	Privacy    struct {
		IsShow            bool `json:"isShow"`
		IsEmailPrivate    bool `json:"isEmailPrivate"`
		IsBirthdayPrivate bool `json:"isBirthdayPrivate"`
		IsLocationPrivate bool `json:"isLocationPrivate"`
		IsGenderPrivate   bool `json:"isGenderPrivate"`
		IsPronounsPrivate bool `json:"isPronounsPrivate"`
	} `json:"privacy"`
	Roles   []string `json:"roles"`
	Likes   []string `json:"likes"`
	Skills  []string `json:"skills"`
	Socials []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
		Icon struct {
			Label string `json:"label"`
			Value string `json:"value"`
			URL   string `json:"url"`
		} `json:"icon"`
	} `json:"socials"`
	CreatedAt    primitive.DateTime `json:"createdAt"`
	UpdatedAt    primitive.DateTime `json:"updatedAt"`
	DeletedAt    interface{}        `json:"deletedAt"`
	IsLiked      bool               `json:"isLiked"`
	IsSelf       bool               `json:"isSelf"`
	IsTeamMember bool               `json:"isTeamMember"`
}

type Sort []struct {
	Default bool   `json:"default"`
	Label   string `json:"label"`
	Value   string `json:"value"`
	Icon    Icon   `json:"icon"`
}

type Social []struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Color   string `json:"color"`
	Icon    Icon   `json:"icon"`
	Enabled bool   `json:"enabled"`
}

type Icon struct {
	Label string `json:"label"`
	Value string `json:"value"`
	URL   string `json:"url"`
}

type Pagination struct {
	Page     string `json:"page"`
	Limit    string `json:"limit"`
	Sort     string `json:"sort"`
	Roles    string `json:"roles"`
	Skills   string `json:"skills"`
	Query    string `json:"query"`
	Language string `json:"language"`
}

type ImgbbImage struct {
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	Mime      string `json:"mime"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
}

type ImgbbResponse struct {
	Data    ImgbbResponseData `json:"data"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
}

type ImgbbResponseData struct {
	ID string `json:"id"`

	DisplayURL string `json:"display_url"`
	DeleteURL  string `json:"delete_url"`

	Expiration string `json:"expiration"`

	Height string `json:"height"`
	Width  string `json:"width"`

	Image  ImgbbImage `json:"image"`
	Thumb  ImgbbImage `json:"thumb"`
	Medium ImgbbImage `json:"medium"`
}

type Config struct {
	ApiVersion int `json:"apiVersion"`
	Database   `json:"database"`
	Web        `json:"web"`
	Client     `json:"client"`
	Collection string `json:"collection"`
	APIUrl     string `json:"apiUrl"`
}

type Database struct {
	Url string `json:"url"`
}

type Web struct {
	Port           string `json:"port"`
	ImageUploadKey string `json:"imageUploadKey"`
	ReturnUrl      string `json:"returnUrl"`
}

type Client struct {
	Id       string `json:"id"`
	Secret   string `json:"secret"`
	Token    string `json:"token"`
	Callback string `json:"callback"`
}

type Roles struct {
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Skills struct {
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
