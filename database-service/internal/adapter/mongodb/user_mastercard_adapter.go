package mongodb

import (
	"context"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	dto "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/mongodb/dto"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var _ out.UsuarioRepository = (*MastercardMongoUserAdapter)(nil)

type MastercardMongoUserAdapter struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
	users          []*dto.Usuario
}

func (m *MastercardMongoUserAdapter) Flush() error {
	return nil
}

func (m *MastercardMongoUserAdapter) SaveUsuario(usuario *domain.Usuario) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	userDto, err := m.getUsuario(usuario.NumeroCuenta())
	if err != nil {
		userDto = &dto.Usuario{
			Id:              primitive.NewObjectID(),
			Nombre:          usuario.Nombre(),
			NumeroCuenta:    usuario.NumeroCuenta(),
			Direccion:       usuario.Direccion(),
			Localidad:       usuario.Localidad(),
			CodigoPostal:    usuario.CodigoPostal(),
			Email:           usuario.Email(),
			NumeroDocumento: usuario.NumeroDocumento(),
		}

		_, err = m.collection().InsertOne(ctx, userDto)

		return err
	}

	return err
}

func (m *MastercardMongoUserAdapter) GetUsuario(nroCuenta string) (*domain.Usuario, error) {

	usuarioDto, err := m.getUsuario(nroCuenta)
	if err != nil {
		log.Printf("Could not find usuario %s error: %s", nroCuenta, err)
		return nil, err
	}

	usuario := usuarioDtoToDomain(usuarioDto)

	return usuario, nil
}

func (m *MastercardMongoUserAdapter) getUsuario(nroCuenta string) (*dto.Usuario, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var usuarioDto = &dto.Usuario{}

	err := m.collection().FindOne(ctx, bson.M{"numero_cuenta": nroCuenta}).Decode(usuarioDto)
	if err != nil {
		if err.Error() == "cannot Decode to nil value" {
			log.Printf("Could not find usuario %s error: %s", nroCuenta, err)
		}
		return nil, err
	}

	return usuarioDto, nil
}

func usuarioDtoToDomain(indexDto *dto.Usuario) *domain.Usuario {
	usuario := domain.NewUsuario()

	usuario.SetEmail(indexDto.Email)
	usuario.SetNombre(indexDto.Nombre)
	usuario.SetLocalidad(indexDto.Localidad)
	usuario.SetNumeroCuenta(indexDto.NumeroCuenta)
	usuario.SetNumeroDocumento(indexDto.NumeroDocumento)
	usuario.SetCodigoPostal(indexDto.CodigoPostal)
	usuario.SetDireccion(indexDto.Direccion)
	return usuario
}

func (a *MastercardMongoUserAdapter) collection() *mongo.Collection {
	return a.client.Database(a.databaseName).Collection(a.collectionName)
}

func NewUserRepository(config *pkg.DatabaseServiceConfig) *MastercardMongoUserAdapter {

	mongodbConfig := config.Mongodb

	opts := options.Client()

	opts.ApplyURI(mongodbConfig.Uri)
	opts.SetTimeout(time.Duration(mongodbConfig.SessionTimeoutMs) * time.Millisecond)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbConfig.Uri))

	if err != nil {
		panic(err)
	}
	return &MastercardMongoUserAdapter{
		client:         client,
		databaseName:   mongodbConfig.DatabaseName,
		collectionName: mongodbConfig.UserCollectionName,
	}
}
