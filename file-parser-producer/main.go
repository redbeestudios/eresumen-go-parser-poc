package main

import (
	"fmt"
	cmdcommons "github.com/redbeestudios/go-parser-poc/commons/cmd"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/cmd"
)

func main() {
	env, err := pkg.NewEnv("dev")
	if err != nil {
		panic(fmt.Sprintf("error creating environment: %s", err.Error()))
	}

	config := cmdcommons.InitConfig(env).FileParserProducer
	deps := cmd.InitDependencies(config)
	router := cmd.InitRoutes(deps)

	cmd.StartServer(config, router)
}
