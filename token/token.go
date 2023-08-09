package token

import (
	"ZuzinhoBot/errlog"
	"encoding/json"
	"io"
	"os"
)

type config struct {
	BotTgToken string `json:"tg_bot_token"`
}

func Token() string {
	return "6336763734:AAGXqgxjMWnksl1_hef5Drpc3UeJM1OloXA"
	f, err := os.Open("config.json")
	errlog.LogOnErr(err)

	jsonStr, err := io.ReadAll(f)
	errlog.LogOnErr(err)

	conf := config{}

	err = json.Unmarshal(jsonStr, &conf)
	errlog.LogOnErr(err)

	return conf.BotTgToken
}
