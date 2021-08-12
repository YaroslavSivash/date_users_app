package mongo

import (
	"context"
	"date_users_app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	LastName  string             `bson:"lastname"`
	Country   string             `bson:"country"`
	City      string             `bson:"city"`
	Gender    string             `bson:"gender"`
	BirthDate string             `bson:"birth_date"`
}

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collection),
	}
}

func (u UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (u UserRepository) GetAllUsers(ctx context.Context, skip, limit int) ([]*models.User, error) {

	panic("implement me")
}

//func (r BookmarkRepository) GetBookmarks(ctx context.Context, user *models.User) ([]*models.Bookmark, error) {
//	uid, _ := primitive.ObjectIDFromHex(user.ID)
//	cur, err := r.db.Find(ctx, bson.M{
//		"userId": uid,
//	})
//	defer cur.Close(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	out := make([]*Bookmark, 0)
//
//	for cur.Next(ctx) {
//		user := new(Bookmark)
//		err := cur.Decode(user)
//		if err != nil {
//			return nil, err
//		}
//
//		out = append(out, user)
//	}
//	if err := cur.Err(); err != nil {
//		return nil, err
//	}
//
//	return toBookmarks(out), nil
//}
func (u UserRepository) UpdateUser(ctx context.Context, user *models.User, id string) error {
	panic("implement me")
}

func toMongoUser(u *models.User) *User {
	return &User{
		Email: u.Email,
	}
}

func toModel(u *User) *models.User {
	return &models.User{
		Id:        u.Id.Hex(),
		Email:     u.Email,
		LastName:  u.LastName,
		Country:   u.Country,
		City:      u.City,
		Gender:    u.Gender,
		BirthDate: u.BirthDate,
	}
}
