package service

import (
	"context"
	"user-basic/common"
	"user-basic/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
    "time"
	"github.com/go-redis/redis/v8"
	"github.com/beevik/guid"
	"strconv"
	"encoding/json"
	"fmt"
)



type UserService struct {
}

type RedisData struct{
	Code string
	AccessGuid string
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

	data := &RedisData{Code:strconv.Itoa(randomCode),AccessGuid:guid.New().String()}

	jsonData, err := json.Marshal(data)
	
	fmt.Println(string(jsonData))

	rdb.Set(context.Background(), email, string(jsonData), 5*time.Minute)

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

func Authenticate(auth *models.Authenticate) (string, error){

	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })


	val, err := rdb.Get(context.Background(), auth.Email).Result()
    if err != nil {
        return "", err
    }

	var redisData *RedisData
	json.Unmarshal([]byte(val), &redisData)
	
	result1 := redisData.Code == auth.Code
	result2 := redisData.AccessGuid == auth.AccessGuid

	if !result1 && !result2 {
		return "", err
	} 


	
	
	return common.Authenticate(auth.Email)
}

