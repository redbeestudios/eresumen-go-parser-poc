package mongodb

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	Id              primitive.ObjectID `bson:"_id"`
	Nombre          string             `bson:"nombre,omitempty"`
	Direccion       string             `bson:"direccion,omitempty"`
	Localidad       string             `bson:"localidad,omitempty"`
	CodigoPostal    string             `bson:"codigo_postal,omitempty"`
	Email           string             `bson:"email,omitempty"`
	NumeroCuenta    string             `bson:"numero_cuenta,omitempty"`
	NumeroDocumento string             `bson:"numero_documento,omitempty"`
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
