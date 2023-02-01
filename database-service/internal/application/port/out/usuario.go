package out

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
)

type UsuarioRepository interface {
	SaveUsuario(usuario *domain.Usuario) error
	GetUsuario(nroCuenta string) (*domain.Usuario, error)
	Flush() error
}
