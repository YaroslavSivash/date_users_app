package mongo

import (
	"context"
	"date_users_app/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"

	//"github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r UserRepository) GetAllUsersDB(c echo.Context, perPage, page int) ([]*models.User, error) {
	//sliceAllUsers := []models.User{}
	findOptions := options.Find()
	filter := bson.M{}
	limit := int64(perPage)

	skip := int64(page)
	findOptions.SetSkip((int64(skip) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, _ := r.db.Find(context.Background(), filter, findOptions)
	defer cursor.Close(context.Background())

	out := []*models.User{}

	for cursor.Next(context.Background()) {
		var user *models.User
		cursor.Decode(&user)
		out = append(out, user)
	}

	return out, nil
}

func (r UserRepository) UpdateUserDB(c echo.Context, user *models.User) error {
	filter := bson.D{{"_id", user.Id}}

	userFind := models.User{}
	err := r.db.FindOne(context.Background(), filter).Decode(&userFind)
	if err != nil {
		log.Error(err)
		return err
	}
	if strings.TrimSpace(user.Email) != "" {
		userFind.Email = user.Email
	}
	if strings.TrimSpace(user.LastName) != "" {
		userFind.LastName = user.LastName
	}
	if strings.TrimSpace(user.Country) != "" {
		userFind.Country = user.Country
	}
	if strings.TrimSpace(user.City) != "" {
		userFind.City = user.City
	}
	if strings.TrimSpace(user.Gender) != "" {
		userFind.Gender = user.Gender
	}
	if strings.TrimSpace(user.BirthDate) != "" {
		userFind.BirthDate = user.BirthDate
	}

	UpdateUser := bson.M{"$set": bson.M{
		"email":     userFind.Email,
		"lastname":  userFind.LastName,
		"country":   userFind.Country,
		"city":      userFind.City,
		"gender":    userFind.Gender,
		"birthdate": userFind.BirthDate,
	},
	}
	_, err = r.db.UpdateOne(context.Background(), filter, UpdateUser)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, err)
}
