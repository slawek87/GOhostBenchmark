package settings

import "gopkg.in/mgo.v2"

func MongoDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return session
}
