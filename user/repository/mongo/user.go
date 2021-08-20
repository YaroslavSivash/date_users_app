package mongo

import (
	"context"
	"date_users_app/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"

	//"github.com/labstack/gommon/log"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"strconv"
	//"github.com/labstack/echo/v4"
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

func (r UserRepository) CreateUserDB(c echo.Context, user *models.User) error {
	_, err := r.db.InsertOne(context.Background(), user)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, err)
}

func (r UserRepository) GetAllUsersDB(c echo.Context, skip, limit int) ([]*models.User, error) {
	//sliceAllUsers := []models.User{}

	return nil, nil
}

func (r UserRepository) UpdateUserDB(c echo.Context, user *models.User) error {
	//filter := bson.D{{"_id", user.Id}}
	return nil
	//user := r.db.InsertOne(context.Background(), filter)
}

func toMongoUser(u *models.User) *User {
	uid, _ := primitive.ObjectIDFromHex(u.Id)
	return &User{
		Id:        uid,
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
