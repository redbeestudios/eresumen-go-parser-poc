package main

import (
	"fmt"
	cmd2 "github.com/redbeestudios/go-parser-poc/commons/cmd"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"github.com/redbeestudios/go-parser-poc/database-service/cmd"
	"sync"
)

func main() {

	env, err := pkg.NewEnv("dev")
	if err != nil {
		panic(fmt.Sprintf("error creating environment: %s", err.Error()))
	}

	config := cmd2.InitConfig(env)
	deps := cmd.InitDependencies(config.DatabaseService)
	router := cmd.InitRoutes(deps)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd.StartServer(config.DatabaseService, router)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd.StartConsumer(config, deps)
	}()

	wg.Wait()
}
