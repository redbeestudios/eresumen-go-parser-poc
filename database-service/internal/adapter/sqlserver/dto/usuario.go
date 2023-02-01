package sqlserver

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre          string `gorm:"column:nombre"`
	Direccion       string `gorm:"column:direccion"`
	Localidad       string `gorm:"column:localidad"`
	CodigoPostal    string `gorm:"column:codigo_postal"`
	Email           string `gorm:"column:email"`
	NumeroCuenta    string `gorm:"column:numero_cuenta"`
	NumeroDocumento string `gorm:"column:numero_documento"`
}

func UsuarioFromDomain(u *domain.Usuario) *Usuario {

	return &Usuario{
		Nombre:          u.Nombre(),
		Direccion:       u.Direccion(),
		Localidad:       u.Localidad(),
		CodigoPostal:    u.CodigoPostal(),
		Email:           u.Email(),
		NumeroCuenta:    u.NumeroCuenta(),
		NumeroDocumento: u.NumeroDocumento(),
	}
}

func (Usuario) TableName() string {
	return "user"

}
