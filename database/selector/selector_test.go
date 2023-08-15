package selector

import (
	"ZuzinhoBot/env"
	"github.com/jackc/pgx"
	"testing"
)

func TestSelect(t *testing.T) {
	var err error

	env.ConnConfig, err = pgx.ParseConnectionString("postgres://tgbotdb_d4vs_user:alpbF7ZdCwZhePS0zqFe1EZJvZsbLanT@dpg-cjch7tk5kgrc73bp2jvg-a.frankfurt-postgres.render.com/tgbotdb_d4vs")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		value, err := Select("phrases")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(value)
	}
}
