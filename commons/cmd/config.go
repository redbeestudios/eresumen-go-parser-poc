package cmdcommons

import (
	"encoding/json"
	"fmt"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"os"
)

func InitConfig(env pkg.Env) *pkg.Config {
	config := &pkg.Config{}

	// TODO: este codigo podria moverse a pkg bajo el nombre de ReadJsonFile,
	// para en el futuro formar parte de una lib de Redbee
	jsonConfig, err := os.Open(fmt.Sprintf("../env_%s.json", env.String()))
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %s", err.Error()))
	}
	defer jsonConfig.Close()

	jsonParser := json.NewDecoder(jsonConfig)
	if err = jsonParser.Decode(config); err != nil {
		panic(fmt.Sprintf("Error parsing config file: %s", err.Error()))
	}

	return config
}
