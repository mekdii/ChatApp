package UserRepo

import (
	"context"
	"errors"

	"github.com/samuael/Project/Weg/internal/pkg/User"
	"github.com/samuael/Project/Weg/internal/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo struct
type UserRepo struct {
	DB *mongo.Database
}

// NewUserRepo function returning client repo pointer
func NewUserRepo(db *mongo.Database) User.UserRepo {
	return &UserRepo{
		DB: db,
	}
}

// UserEmailExist check the existence of the user
func (urepo *UserRepo) UserEmailExist(email string) error {
	filter := bson.D{{"email", email}}
	return urepo.DB.Collection("user").FindOne(context.TODO(), filter).Err()
}

// RegisterUser user registration
func (urepo *UserRepo) RegisterUser(user *entity.User) (*entity.User, error) {
	_, er := urepo.DB.Collection(entity.USER).InsertOne(context.TODO(), user)
	if er != nil {
		return user, er
	}
	return user, nil
}

// GetUserByEmail repository method returning user if exist
func (urepo *UserRepo) GetUserByEmail(Email string) (*entity.User, error) {
	filter := bson.D{{"email", Email}}
	user := &entity.User{}
	era := urepo.DB.Collection("user").FindOne(context.TODO(), filter).Decode(user)
	return user, era
}

// GetUserByEmailAndID function to Get User By Email and ID
func (urepo *UserRepo) GetUserByEmailAndID(Email, ID string) (*entity.User, error) {
	filter := bson.D{{"email", Email},}
	user := &entity.User{}
	era := urepo.DB.Collection("user").FindOne(context.TODO(), filter).Decode(user)
	return user, era
}

// SaveUser functio fro fetching a user
func (urepo *UserRepo) SaveUser(usr *entity.User) (*entity.User, error) {

	oid  , er := primitive.ObjectIDFromHex(usr.ID)
	if er != nil {
		return nil  , er
	}
	_, er = urepo.DB.Collection(entity.USER).UpdateOne(context.TODO(), bson.D{{"_id", oid}},
		bson.D{
			{"$set", bson.D{{"username", usr.Username},
			 {"email", usr.Email}, 
			 {"password", usr.Password}, 
			 {"imageurl", usr.Imageurl}, 
			 {"bio", usr.Bio}, 
			 {"mygroups"  , usr.MyGroups},
			 {"last_updated", usr.LastUpdated}, 
			 {"last_seen", usr.LastSeen}}},
		})
	if er != nil {
		return usr, er
	}
	return usr, nil
}


// GetUserByID (id string ) (*entity.User , error )
func (urepo *UserRepo) GetUserByID(id string ) (*entity.User , error ){
	user := &entity.User{}
	oid  , er := primitive.ObjectIDFromHex(id)
	if er != nil {
		return nil  , er 
	}
	er = urepo.DB.Collection(entity.USER).FindOne( context.TODO() , bson.D{{"_id", oid }} ).Decode(user)
	return user  , er 
}


// UserWithIDExist ( friendID string ) error
func (urepo *UserRepo) UserWithIDExist(friendID string ) error {
	user := &entity.User{}
	oid  , er := primitive.ObjectIDFromHex(friendID)
	if er != nil {
		return  er 
	}
	er = urepo.DB.Collection(entity.USER).FindOne( context.TODO() , bson.D{{"_id", oid }}).Decode(user)
	if user.ID =="" || user.Email=="" || user.Password==""{
		return errors.New("No User Found!")
	}
	return nil 
}


// IsGroupMember  repository returning whether the use is a member or not 
// return value error if the user is not a member of that group
// otherwise it returns nil
func (urepo *UserRepo) IsGroupMember(userid , groupid string  ) error {
	user := &entity.User{}
	oid  , er := primitive.ObjectIDFromHex(userid)
	if er != nil {
		return er
	}
	er = urepo.DB.Collection(entity.USER).FindOne(context.TODO()  , bson.D{{"_id"  , oid}}).Decode(user)
	if er != nil {
		return er
	}

	for _  , val := range user.MyGroups{
		if val == groupid {
			return nil 
		}
	}
	return er 
}