package files

import (
	"encoding/json"
	"fmt"
	files "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/files/dto"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	"log"
	"os"
)

var _ out.ResumenRepository = (*MastercardAdapter)(nil)

type MastercardAdapter struct {
	basePath string
}

func (a *MastercardAdapter) GetResumen(codVinculaciion string) (*domain.Resumen, error) {
	b, err := os.ReadFile(codVinculaciion)
	if err != nil {
		return nil, err
	}

	var resumenDto = files.Resumen{}

	err = json.Unmarshal(b, &resumenDto)

	if err != nil {
		return nil, err
	}

	return resumenDto.ToDomain(), nil

}

func NewFilesRepository(path string) *MastercardAdapter {

	return &MastercardAdapter{basePath: path}
}

func (a *MastercardAdapter) SaveResumen(resumen *domain.Resumen, codigoVinculacion string) (string, error) {
	resumenDto := files.DtoFromDomain(resumen)
	path := fmt.Sprintf("%s/%s.json", a.basePath, codigoVinculacion)
	f, err := os.Create(path)
	defer f.Close()
	fileContent, err := json.MarshalIndent(resumenDto, "", " ")
	if err != nil {
		log.Printf("Error creando archivo para resumen %s", resumen.Clave().NroCuenta())
		return "", err
	}
	_, err = f.Write(fileContent)
	if err != nil {
		log.Printf("Error creando archivo para resumen %s", resumen.Clave().NroCuenta())
		return "", err
	}

	return path, nil

}
