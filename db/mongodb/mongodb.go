package mongodb

import (
	// "fmt"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
)

func Insert(url, db, collection string, docs ...interface{}) {
	mongo, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	coll := mongo.DB(db).C(collection)
	err = coll.Insert(docs...)
	if err != nil {
		panic(err)
	}
}
func GetCollection(url, db, collection string) (*mgo.Collection, error) {
	mongo, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	defer mongo.Close()
	coll := mongo.DB(db).C(collectionName)
	return coll, nil
}
func Find(url, db, collction string, query interface{}) (*mgo.Query, error) {
	coll, err := GetCollection(url, db, collction)
	if err != nil {
		return nil, err
	}
	return coll.Find(), nil
}
