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

func (r *mongoRepo) GetAllSpecialists() ([]*entity.Specialist, error) {
	cursor, err := r.db.Collection("specialists").Find(context.Background(), struct{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var specialists []*entity.Specialist
	for cursor.Next(context.Background()) {
		var sp entity.Specialist
		if err := cursor.Decode(&sp); err != nil {
			return nil, err
		}
		specialists = append(specialists, &sp)
	}
	return specialists, nil
}

func (r *mongoRepo) GetAllProcedures() ([]*entity.Procedure, error) {
	cursor, err := r.db.Collection("procedures").Find(context.Background(), struct{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var procedures []*entity.Procedure
	for cursor.Next(context.Background()) {
		var p entity.Procedure
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		procedures = append(procedures, &p)
	}
	return procedures, nil
}

func (r *mongoRepo) AssignProcedureToSpecialist(specialistID, procedureID string) error {
	filter := map[string]interface{}{"id": specialistID}
	update := map[string]interface{}{
		"$addToSet": map[string]interface{}{
			"procedure_ids": procedureID,
		},
	}
	_, err := r.db.Collection("specialists").UpdateOne(context.Background(), filter, update)
	return err
}
