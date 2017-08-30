package handlers

import (
	"log"
	"os"

	"github.com/docker/docker/client"
	"github.com/turkenh/play-with-ansible/config"
	"github.com/turkenh/play-with-ansible/docker"
	"github.com/turkenh/play-with-ansible/pwd"
	"github.com/turkenh/play-with-ansible/storage"
)

var core pwd.PWDApi
var Broadcast pwd.BroadcastApi

func Bootstrap() {
	c, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	d := docker.NewDocker(c)

	Broadcast, err = pwd.NewBroadcast(WS, WSError)
	if err != nil {
		log.Fatal(err)
	}

	t := pwd.NewScheduler(Broadcast, d)

	s, err := storage.NewFileStorage(config.SessionsFile)

	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error decoding sessions from disk ", err)
	}
	core = pwd.NewPWD(d, t, Broadcast, s)

}
