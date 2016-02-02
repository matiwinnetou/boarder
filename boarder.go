package main

import "fmt"

import (
	"log"
	"time"
	"github.com/matiwinnetou/boarder/utils"
	"golang.org/x/net/context"
	"github.com/coreos/etcd/client'
	"github.com/coreos/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}

	fmt.Printf(utils.Reverse("abc"))
}
