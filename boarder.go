package main

import (
	"fmt"
	"time"
	"log"
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/matiwinnetou/boarder/chess"
"github.com/matiwinnetou/boarder/core"
)

func main() {
	t1 := chess.NewChessTable(1)
	t1.Player1 = &core.Player{Name: "Jan"}
	t1.Player2 = &core.Player{Name: "Stefan"}

	fmt.Printf("%d\n", t1.TableNo)
	fmt.Printf("%d\n", t1.PlayerCount())

	cfg := client.Config{
		Endpoints:               []string{"http://pi:2379"},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err = kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	//log.Printf("Get is done. Metadata is %q\n", resp)
	log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
}
