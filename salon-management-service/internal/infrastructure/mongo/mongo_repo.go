package mongo

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *mongoRepo) UpdateSalon(s *entity.Salon) error {
	_, err := r.db.Collection("salons").UpdateOne(context.Background(),
		bson.M{"id": s.ID},
		bson.M{"$set": bson.M{"name": s.Name, "location": s.Location}},
	)
	return err
}

func (r *mongoRepo) DeleteSalon(id string) error {
	_, err := r.db.Collection("salons").DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

func (r *mongoRepo) GetAllSalons() ([]*entity.Salon, error) {
	cursor, err := r.db.Collection("salons").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.Salon
	for cursor.Next(context.Background()) {
		var s entity.Salon
		if err := cursor.Decode(&s); err != nil {
			return nil, err
		}
		results = append(results, &s)
	}
	return results, nil
}

func (r *mongoRepo) AddProcedure(p *entity.Procedure) (*entity.Procedure, error) {
	p.ID = uuid.New().String()
	_, err := r.db.Collection("procedures").InsertOne(context.Background(), p)
	return p, err
}

func (r *mongoRepo) UpdateProcedure(p *entity.Procedure) error {
	_, err := r.db.Collection("procedures").UpdateOne(context.Background(),
		bson.M{"id": p.ID},
		bson.M{"$set": bson.M{
			"name":        p.Name,
			"duration":    p.Duration,
			"description": p.Description,
			"salonid":     p.SalonID,
		}},
	)
	return err
}

func (r *mongoRepo) DeleteProcedure(id string) error {
	_, err := r.db.Collection("procedures").DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

func (r *mongoRepo) GetAllProcedures() ([]*entity.Procedure, error) {
	cursor, err := r.db.Collection("procedures").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.Procedure
	for cursor.Next(context.Background()) {
		var p entity.Procedure
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		results = append(results, &p)
	}
	return results, nil
}

func (r *mongoRepo) AddSpecialist(sp *entity.Specialist) (*entity.Specialist, error) {
	sp.ID = uuid.New().String()
	_, err := r.db.Collection("specialists").InsertOne(context.Background(), sp)
	return sp, err
}

func (r *mongoRepo) UpdateSpecialist(sp *entity.Specialist) error {
	_, err := r.db.Collection("specialists").UpdateOne(context.Background(),
		bson.M{"id": sp.ID},
		bson.M{"$set": bson.M{
			"name":    sp.Name,
			"bio":     sp.Bio,
			"salonid": sp.SalonID,
		}},
	)
	return err
}

func (r *mongoRepo) DeleteSpecialist(id string) error {
	_, err := r.db.Collection("specialists").DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

func (r *mongoRepo) GetAllSpecialists() ([]*entity.Specialist, error) {
	cursor, err := r.db.Collection("specialists").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.Specialist
	for cursor.Next(context.Background()) {
		var sp entity.Specialist
		if err := cursor.Decode(&sp); err != nil {
			return nil, err
		}
		results = append(results, &sp)
	}
	return results, nil
}

func (r *mongoRepo) AssignProcedureToSpecialist(specialistID, procedureID string) error {
	_, err := r.db.Collection("specialists").UpdateOne(context.Background(),
		bson.M{"id": specialistID},
		bson.M{"$addToSet": bson.M{"procedure_ids": procedureID}},
	)
	return err
}

func (r *mongoRepo) RemoveProcedureFromSpecialist(specialistID, procedureID string) error {
	_, err := r.db.Collection("specialists").UpdateOne(context.Background(),
		bson.M{"id": specialistID},
		bson.M{"$pull": bson.M{"procedure_ids": procedureID}},
	)
	return err
}

func (r *mongoRepo) GetScheduleOverride(procedureID string, date string) (*entity.ProcedureScheduleOverride, error) {
	filter := bson.M{"procedure_id": procedureID, "date": date}

	var result entity.ProcedureScheduleOverride
	err := r.db.Collection("schedule_overrides").FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &result, err
}

func (r *mongoRepo) GetWeeklySchedule(procedureID string, weekday int32) (*entity.WeeklyProcedureSchedule, error) {
	filter := bson.M{"procedure_id": procedureID, "weekday": weekday}

	var result entity.WeeklyProcedureSchedule
	err := r.db.Collection("weekly_schedules").FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &result, err
}
