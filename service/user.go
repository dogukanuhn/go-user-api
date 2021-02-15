package service

import (
	"context"
	"user-basic/common"
	"user-basic/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
  
)



type UserService struct {
}



func Login(email string) (*models.User,error) {


	filter := bson.D{primitive.E{Key: "email", Value: email}}
	user := &models.User{}
	//Get MongoDB connection using connectionhelper.
	client, err := common.GetMongoClient()
	if err != nil {
		return user, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(common.DB).Collection(common.USERS)
	//Perform InsertOne operation & validate against the error.
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err
	}

	rand.Seed(time.Now().UnixNano())

    min := 100000
    max := 1000000

    randomCode := rand.Intn(max - min + 1) + min

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	rdb.Set(context.Background(), email+":"+guid.New().String(), randomCode, 0)

	return user, nil
}


func Register(user *models.User) (bool, error) {
	
	
	//Get MongoDB connection using connectionhelper.
	client, err := common.GetMongoClient()
	if err != nil {
		return false, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(common.DB).Collection(common.USERS)
	//Perform InsertOne operation & validate against the error.
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return false, err
	}

	return true, err
}

// func Login(ctx echo.Context) string {
// 	// var userMail = ctx.Param("email")
// 	// fmt.Println(userMAil)
// 	//  user := User{"10","test@mail.com"}
// 	 return token
// }

// func  GetUser(ctx echo.Context) User  {

// 	sort.Slice(users, func(i, j int) bool {
//         return crowd[i].Name <= crowd[j].Name
//     })

// }
