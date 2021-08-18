package mongo

import (
	"context"
	"date_users_app/models"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"

	"github.com/labstack/echo/v4"
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

func (r UserRepository) CreateUserDB(ctx context.Context, user *models.User) error {
	model := toMongoUser(user)
	res, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	user.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r UserRepository) GetAllUsersDB(ctx context.Context, skip, limit int) ([]*models.User, error) {
	rez := []models.User{}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	skip, err := strconv.Atoi(c.QueryParam("skip"))

	limit := int64(limit)
	skip := int64(skip)

	cursor, err := r.Find(context.Background(), bson.M{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bson.D{{"email", 1}},
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		err = cursor.Close(context.Background())
		if err != nil {
			log.Error(err)
		}
	}()

	// parse all
	for cursor.Next(context.TODO()) {
		var episode models.User
		if err = cursor.Decode(&episode); err != nil {
			log.Error(err)
		}

		result = append(result, episode)
	}

	// ---------------Find many

	return &result, nil
}

}

func (r UserRepository) UpdateUserDB(ctx context.Context, user *models.User, id string) error {

}

func toMongoUser(u *models.User) *User {
	return &User{

		Email:     u.Email,
		LastName:  u.LastName,
		Country:   u.Country,
		City:      u.City,
		Gender:    u.Gender,
		BirthDate: u.BirthDate,
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
