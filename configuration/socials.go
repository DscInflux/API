package configuration

import (
	"go.dscinflux.xyz/types"
)

func getSocials() types.Social {
	return types.Social{
		{
			ID:    "1",
			Name:  "Github",
			URL:   "https://github.com/{username}",
			Color: "#000000",
			Icon: types.Icon{
				Label: "fa-github",
				Value: "fab",
				URL:   "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
			},
			Enabled: true,
		},
		{
			ID:    "2",
			Name:  "Twitter/X",
			URL:   "https://x.com/{username}",
			Color: "#1DA1F2",
			Icon: types.Icon{
				Label: "fa-twitter",
				Value: "fab",
				URL:   "https://cdn.topiclist.xyz/images/svg/x.com.svg",
			},
			Enabled: true,
		},
		{
			ID:    "3",
			Name:  "Facebook",
			URL:   "https://facebook.com/{username}",
			Color: "#3b5998",
			Icon: types.Icon{
				Label: "fa-facebook",
				Value: "fab",
				URL:   "https://www.facebook.com/images/fb_icon_325x325.png",
			},
			Enabled: true,
		},
		{
			ID:    "4",
			Name:  "Instagram",
			URL:   "https://instagram.com/{username}",
			Color: "#E1306C",
			Icon: types.Icon{
				Label: "fa-instagram",
				Value: "fab",
				URL:   "https://www.instagram.com/static/images/ico/favicon-192.png/68d99ba29cc8.png",
			},
			Enabled: true,
		},
		{
			ID:    "5",
			Name:  "LinkedIn",
			URL:   "https://linkedin.com/in/{username}",
			Color: "#0077B5",
			Icon: types.Icon{
				Label: "fa-linkedin",
				Value: "fab",
				URL:   "https://static-exp1.licdn.com/sc/h/7fx9nkd7mx8avdpqm5hqcbi97",
			},
			Enabled: true,
		},
		{
			ID:    "6",
			Name:  "StackOverflow",
			URL:   "https://stackoverflow.com/users/{username}",
			Color: "#f48024",
			Icon: types.Icon{
				Label: "fa-stack-overflow",
				Value: "fab",
				URL:   "https://cdn.sstatic.net/Sites/stackoverflow/img/favicon.ico?v=4f32ecc8f43d",
			},
			Enabled: true,
		},
		{
			ID:    "7",
			Name:  "Reddit",
			URL:   "https://reddit.com/user/{username}",
			Color: "#FF4500",
			Icon: types.Icon{
				Label: "fa-reddit",
				Value: "fab",
				URL:   "https://www.redditstatic.com/desktop2x/img/favicon/favicon-32x32.png",
			},
			Enabled: true,
		},
		{
			ID:    "8",
			Name:  "YouTube",
			URL:   "https://youtube.com/channel/{username}",
			Color: "#FF0000",
			Icon: types.Icon{
				Label: "fa-youtube",
				Value: "fab",
				URL:   "https://s.ytimg.com/yts/img/favicon-vfl8qSV2F.ico",
			},
			Enabled: true,
		},
		{
			ID:    "9",
			Name:  "Steam",
			URL:   "https://steamcommunity.com/{username}",
			Color: "#FF0000",
			Icon: types.Icon{
				Label: "fa-steam",
				Value: "fab",
				URL:   "https://w7.pngwing.com/pngs/49/992/png-transparent-steam-logo-computer-icons-steam-tuners-steam-miscellaneous-rim-area-thumbnail.png",
			},
			Enabled: true,
		},
		{
			ID:    "10",
			Name:  "Twitch",
			URL:   "https://www.twitch.tv/{username}",
			Color: "#FF0000",
			Icon: types.Icon{
				Label: "fa-twitch",
				Value: "fab",
				URL:   "https://cdn.topiclist.xyz/logos/TwitchBlack.png",
			},
			Enabled: true,
		},
		{
			ID:    "11",
			Name:  "MyAnimeList",
			URL:   "https://myanimelist.net/profile/{username}",
			Color: "#FF0000",
			Icon: types.Icon{
				Label: "fa-myanimelist",
				Value: "fab",
				URL:   "https://cdn.topiclist.xyz/logos/myanimelist.png",
			},
			Enabled: true,
		},
	}
}

func GetSocials() types.Social {
	return getSocials()
}
