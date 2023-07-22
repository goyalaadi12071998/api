package main

import (
	"context"
	"interview/app/boot"
	"interview/app/common"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := common.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	if err := boot.Init(ctx, *env); err != nil {
		log.Fatal(err)
	}
}
