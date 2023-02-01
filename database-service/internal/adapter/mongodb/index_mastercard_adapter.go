package mongodb

import (
	"context"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	dto "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/mongodb/dto"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	sf "github.com/sa-/slicefunk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var _ out.IndexRepository = (*MastercardMongoIndexAdapter)(nil)

type MastercardMongoIndexAdapter struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func (a *MastercardMongoIndexAdapter) Flush() error {

	return nil
}

func (a *MastercardMongoIndexAdapter) SaveIndex(index *domain.Indice, usuario string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	indexDto := &dto.IndiceResumen{
		Id:            primitive.NewObjectID(),
		Usuario:       usuario,
		FechaCierre:   index.FechaCierre(),
		Marca:         index.Marca(),
		Path:          index.Path(),
		Month:         index.Month(),
		Year:          index.Year(),
		ModeloResumen: index.ModeloResumen(),
	}

	_, err := a.collection().InsertOne(ctx, indexDto)

	return err
}

func (a *MastercardMongoIndexAdapter) GetIndex(nroCuenta string, fechaCierre string) (*domain.Indice, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var indexDto = &dto.IndiceResumen{}

	err := a.collection().FindOne(ctx, bson.M{"numero_cuenta": nroCuenta, "fecha_cierra": fechaCierre}).Decode(indexDto)
	if err != nil {
		if err.Error() == "cannot Decode to nil value" {
			log.Printf("Could not find index %s error: %s", fechaCierre, err)
		}
		return nil, err
	}

	index := indiceDtoToDomain(indexDto)

	return index, nil
}

func indiceDtoToDomain(indexDto *dto.IndiceResumen) *domain.Indice {
	index := domain.NewIndice()

	index.SetPath(indexDto.Path)
	index.SetYear(indexDto.Year)
	index.SetMonth(indexDto.Month)
	index.SetMarca(indexDto.Marca)
	index.SetFechaCierre(indexDto.FechaCierre)
	index.SetModeloResumen(indexDto.ModeloResumen)
	return index
}

func (a *MastercardMongoIndexAdapter) GetIndices(nroCuenta string) ([]*domain.Indice, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var indicesDto = []*dto.IndiceResumen{}

	cursor, err := a.collection().Find(ctx, bson.M{"numero_cuenta": nroCuenta})
	if err != nil {
		log.Printf("Could not find index %s error: %s", nroCuenta, err)
		return nil, err
	}

	err = cursor.Decode(indicesDto)
	if err != nil {
		if err.Error() == "cannot Decode to nil value" {
			log.Printf("Could not find index %s error: %s", nroCuenta, err)
		}
		return nil, err
	}

	indices := sf.Map[*dto.IndiceResumen, *domain.Indice](indicesDto, func(item *dto.IndiceResumen) *domain.Indice {
		return indiceDtoToDomain(item)
	})

	return indices, nil
}

func (a *MastercardMongoIndexAdapter) collection() *mongo.Collection {
	return a.client.Database(a.databaseName).Collection(a.collectionName)
}

func NewIndexMongoRepository(config *pkg.DatabaseServiceConfig) *MastercardMongoIndexAdapter {

	mongodbConfig := config.Mongodb

	opts := options.Client()

	opts.ApplyURI(mongodbConfig.Uri)
	opts.SetTimeout(time.Duration(mongodbConfig.SessionTimeoutMs) * time.Millisecond)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbConfig.Uri))

	if err != nil {
		panic(err)
	}
	return &MastercardMongoIndexAdapter{
		client:         client,
		databaseName:   mongodbConfig.DatabaseName,
		collectionName: mongodbConfig.IndexCollectionName,
	}
}
