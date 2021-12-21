package main

import (
	"github.com/enkhalifapro/obi-partners/internal/api"
	"github.com/enkhalifapro/obi-partners/internal/partners"
	"github.com/enkhalifapro/obi-partners/pkg/mongo"
	"github.com/enkhalifapro/obi-partners/pkg/slack"
	"log"
	"net/http"
)

type Config struct {
	DBURI      string
	DBName     string
	WebHookURL string
}

func main() {
	// config parsing
	cfg := &Config{
		DBURI: "",
	}
	// dependencies setup

	connector, err := mongo.NewConnector(cfg.DBURI, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	slackProvider := slack.NewProvider(cfg.WebHookURL)
	partnerMgr := partners.NewManager(connector, slackProvider)
	h := api.Handler(partnerMgr)

	http.ListenAndServe("localhost:8090", h)
}
