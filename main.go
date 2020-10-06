package main

import (
	"log"

	D "github.com/NeoJRotary/describe-go"
	"github.com/NeoJRotary/describe-go/dhttp"
	"github.com/zenportinc/GCB-bridge/app"
	"github.com/zenportinc/GCB-bridge/gcloud"
	"github.com/zenportinc/GCB-bridge/webhook"
)

func init() {
	app.Init()
	gcloud.Init()
}

func main() {
	addr := D.GetENV("LISTEN_ON", "0.0.0.0:8080")
	server := dhttp.Server().ListenOn(addr)
	webhook.InitWebhookRoute(server)

	log.Println("Server listen on 8080")
	err := server.Start()
	log.Fatal(err)
}
