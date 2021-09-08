package application

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"goauthentik.io/api"
	"goauthentik.io/internal/config"
	"gopkg.in/boj/redistore.v1"
)

func GetStore(p api.ProxyOutpostConfig) sessions.Store {
	var store sessions.Store
	if config.G.Redis.Host != "" {
		rs, err := redistore.NewRediStoreWithDB(10, "tcp", fmt.Sprintf("%s:%d", config.G.Redis.Host, config.G.Redis.Port), config.G.Redis.Password, strconv.Itoa(config.G.Redis.OutpostSessionDB))
		if err != nil {
			panic(err)
		}
		rs.Options.Domain = *p.CookieDomain
		rs.Options.SameSite = http.SameSiteNoneMode
		store = rs
	} else {
		cs := sessions.NewCookieStore([]byte(*p.CookieSecret))
		cs.Options.Domain = *p.CookieDomain
		cs.Options.SameSite = http.SameSiteNoneMode
		store = cs
	}
	return store
}
