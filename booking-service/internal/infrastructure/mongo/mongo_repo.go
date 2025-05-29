package mongo

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gemdivk/LUMERA-SPA/booking-service/internal/domain"
)

type bookingRepo struct {
	db *mongo.Database
}

func NewMongoRepo(db *mongo.Database) *bookingRepo {
	return &bookingRepo{db: db}
}

func (r *bookingRepo) Create(b *domain.Booking) (*domain.Booking, error) {
	b.ID = uuid.New().String()
	b.Status = "active"
	_, err := r.db.Collection("bookings").InsertOne(context.Background(), b)
	return b, err
}

func (r *bookingRepo) Cancel(id string) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"status": "cancelled"}}
	_, err := r.db.Collection("bookings").UpdateOne(context.Background(), filter, update)
	return err
}

func (r *bookingRepo) Reschedule(id, newDate, newTime string) (*domain.Booking, error) {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"date": newDate, "starttime": newTime, "status": "rescheduled"}}
	_, err := r.db.Collection("bookings").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	var updated domain.Booking
	err = r.db.Collection("bookings").FindOne(context.Background(), filter).Decode(&updated)
	return &updated, err
}

func (r *bookingRepo) ListByClient(clientID string) ([]*domain.Booking, error) {
	filter := bson.M{"clientid": clientID}
	cursor, err := r.db.Collection("bookings").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var result []*domain.Booking
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *bookingRepo) GetAll() ([]*domain.Booking, error) {
	cursor, err := r.db.Collection("bookings").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var bookings []*domain.Booking
	if err := cursor.All(context.Background(), &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}
