package mongo

import "date_users_app/models"

type User struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Email      string  `bson:"email"`
	LastName   string `bson:"lastname"`
	Country    string `bson:"country"`
	City       string `bson:"city"`
	Gender     string `bson:"gender"`
	Birth_date string `bson:"birth_date"`
}


type UserRepository struct {
	db *mongo.Collection
}

func toMongoUser(u *models.User) *User {
	return &User{
		Email: u.Email,
	}
}

func toModel(u *User) *models.User {
	return &models.User{
		Id:         u.Id.Hex(),
		Email:      u.Email,
		LastName:   u.LastName,
		Country:    u.Country,
		City:       u.City,
		Gender:     u.Gender,
		Birth_date: u.Birth_date,
	}
}
