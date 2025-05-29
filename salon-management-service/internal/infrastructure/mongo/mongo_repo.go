package mongo

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gemdivk/LUMERA-SPA/salon-management-service/internal/domain/entity"
)

type mongoRepo struct {
	db *mongo.Database
}

func NewMongoRepo(db *mongo.Database) *mongoRepo {
	return &mongoRepo{db: db}
}

func (r *mongoRepo) AddSalon(s *entity.Salon) (*entity.Salon, error) {
	s.ID = uuid.New().String()
	_, err := r.db.Collection("salons").InsertOne(context.Background(), s)
	return s, err
}

func (r *mongoRepo) AddProcedure(p *entity.Procedure) (*entity.Procedure, error) {
	p.ID = uuid.New().String()
	_, err := r.db.Collection("procedures").InsertOne(context.Background(), p)
	return p, err
}

func (r *mongoRepo) AddSpecialist(sp *entity.Specialist) (*entity.Specialist, error) {
	sp.ID = uuid.New().String()
	_, err := r.db.Collection("specialists").InsertOne(context.Background(), sp)
	return sp, err
}
