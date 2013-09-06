package mongodb

import (
	// "fmt"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
)

func Insert(url, db, collectionName string, docs ...interface{}) {
	mongo, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	collection := mongo.DB(db).C(collectionName)
	err = collection.Insert(docs...)
	if err != nil {
		panic(err)
	}
}

// func mmain() {
// 	// 连接数据库
// 	mongo, err := mgo.Dial("127.0.0.1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mongo.Close()

// 	// 获取数据库,获取集合
// 	collection := mongo.DB("test").C("mail")

// 	// 存储数据
// 	// m1 := Mail{bson.NewObjectId(), "user1", "user1@dotcoo.com"}
// 	// m2 := Mail{bson.NewObjectId(), "user1", "user2@dotcoo.com"}
// 	// m3 := Mail{bson.NewObjectId(), "user3", "user3@dotcoo.com"}
// 	// m4 := Mail{bson.NewObjectId(), "user3", "user4@dotcoo.com"}
// 	m1 := Mail{"user1", "user1@dotcoo.com"}
// 	m2 := Mail{"user1", "user2@dotcoo.com"}
// 	m3 := Mail{"user3", "user3@dotcoo.com"}
// 	m4 := Mail{"user3", "user4@dotcoo.com"}
// 	err = collection.Insert(&m1, &m2, &m3, &m4)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 读取数据
// 	ms := []Mail{}
// 	err = collection.Find(&bson.M{"name": "user3"}).All(&ms)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 显示数据
// 	for i, m := range ms {
// 		// fmt.Printf("%s, %d, %s\n", m.Id.Hex(), i, m.Email)
// 		fmt.Printf(" %d, %s\n", i, m.Email)
// 	}
// }
