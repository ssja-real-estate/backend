package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}
type conn struct {
	session *mgo.Session
}

func NewConnection() Connection {
	var c conn
	var err error
	c.session, err = mgo.Dial(getUrl())
	if err != nil {
		log.Println(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	c.session.Close()
}

func (c *conn) DB() *mgo.Database {
	// return c.session.DB(os.Getenv("DATABASE_NAME"))
	return c.session.DB("Amlak")
}
func getUrl() string {
	// port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	// if err != nil {
	// 	log.Println("error in load os.getenv", err.Error())
	// 	port = 2707
	// }
	port := 27017
	return fmt.Sprintf("mongodb://localhost:%d", port)
	// return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", os.Getenv("DATABASE_USER"),

	// 	os.Getenv("DATABASE_PASS"),
	// 	os.Getenv("DATABASE_HOST"),
	// 	port,
	// 	os.Getenv("DATABASE_NAME"),
	// )

}
