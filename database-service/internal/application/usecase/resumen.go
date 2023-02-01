package usecase

import (
	"context"
	"fmt"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	"log"
)

type SaveResumen struct {
	usuarioRepository out.UsuarioRepository
	indexRepository   out.IndexRepository
	resumenRepository out.ResumenRepository
}

func (c *SaveResumen) FinishSaving() error {
	c.usuarioRepository.Flush()
	c.indexRepository.Flush()
	return nil
}

func (c *SaveResumen) SaveResumen(ctx context.Context, usuario *domain.Usuario, index *domain.Indice, resumen *domain.Resumen) error {

	path, err := c.resumenRepository.SaveResumen(resumen, fmt.Sprintf("%s-%s", usuario.NumeroCuenta(), index.FechaCierre()))
	if err != nil {
		log.Printf("Error saving file : %s %s", usuario.NumeroCuenta(), err.Error())
		return nil
	}

	index.SetPath(path)

	err = c.indexRepository.SaveIndex(index, usuario.NumeroCuenta())
	if err != nil {
		log.Printf("Error saving in database: %s %s", usuario.NumeroCuenta(), err.Error())
		return nil
	}

	err = c.usuarioRepository.SaveUsuario(usuario)
	if err != nil {
		log.Printf("Error saving in database: %s %s", usuario.NumeroCuenta(), err.Error())
		return nil
	}

	return nil
}

func NewSaveResumen(resumenRepository out.ResumenRepository, indexRepository out.IndexRepository, usuarioRepository out.UsuarioRepository) *SaveResumen {
	return &SaveResumen{
		resumenRepository: resumenRepository,
		indexRepository:   indexRepository,
		usuarioRepository: usuarioRepository,
	}
}
