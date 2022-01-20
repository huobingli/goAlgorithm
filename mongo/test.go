package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func testmongo() {
	var (
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://47.114.171.118:32334").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	// 2.选择数据库 test
	db = client.Database("test1")
	fmt.Print("111222")
	fmt.Print(db)
	//3.选择表 my_collection
	collection = db.Collection("testcollection")
	collection = collection
}

type mgo struct {
	uri        string //数据库网络地址
	database   string //要连接的数据库
	collection string //要连接的集合
}

type Trainer struct {
	Name string
	Age  int
	City string
}

func (m *mgo) Connect1() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //养成良好的习惯，在调用WithTimeout之后defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	if err != nil {
		print("123321")
		log.Print(err)
	}
	//collection := client.Database(m.database).Collection(m.collection)

	err1 := client.Ping(context.TODO(), nil)

	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	trainers := []interface{}{ash, misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//return collectionz`
}

func main() {

	var mgo = &mgo{
		// "mongodb+srv://huobingli:george199579@cluster0.qytuu.mongodb.net/test",
		"mongodb+srv://admin:123456@47.114.171.118:32334/test",
		"MainSite",
		"UsersM12",
	}

	mgo.Connect1()
}
