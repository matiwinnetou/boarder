package core

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"os"
	"time"
)

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func init() {
	c, err := redis.DialTimeout("tcp", "pi:6379", time.Duration(10) * time.Second)
	errHndlr(err)
	defer c.Close()

	//* Strings
	r := c.Cmd("set", "mykey0", "myval0")
	errHndlr(r.Err)

	s, err := c.Cmd("get", "mykey0").Str()
	errHndlr(err)
	fmt.Println("mykey0:", s)
}
