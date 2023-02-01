package cmd

import (
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	rest "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/controller/mastercard"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/files"
	kafka "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/kafka"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/mongodb"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/sqlserver"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/usecase"
	"os"
)

type Dependencies struct {
	Controller *rest.Controller
	Consumer   *kafka.Consumer
}

func InitDependencies(config *pkg.DatabaseServiceConfig) *Dependencies {

	var useMongo = os.Getenv("use-mongo") == "true"

	var indexRepository out.IndexRepository
	var userRepository out.UsuarioRepository

	if useMongo {
		indexRepository = mongodb.NewIndexMongoRepository(config)
		userRepository = mongodb.NewUserRepository(config)
	} else {

		indexRepository = sqlserver.NewIndexSqlRepository(config)
		userRepository = sqlserver.NewUserRepository(config)
	}

	var fileRepository out.ResumenRepository

	if useMongo {
		fileRepository = files.NewFilesRepository("/home/rodrigo/Parser-Output/Mongo")
	} else {
		fileRepository = files.NewFilesRepository("/home/rodrigo/Parser-Output/SQL")
	}

	saveResumen := usecase.NewSaveResumen(fileRepository, indexRepository, userRepository)

	getResumen := usecase.NewGetResumen(fileRepository, indexRepository)

	controller := rest.NewMasterCardController(saveResumen, getResumen)

	consumer := kafka.NewMasterCardConsumer(saveResumen)

	return &Dependencies{
		Controller: controller,
		Consumer:   consumer,
	}

}
