package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	hosts    = "mongodb:27017"
	database = "Amlak"
	username = "admin"
	password = "admin"
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
	// info := &mgo.DialInfo{
	// 	Addrs:    []string{hosts},
	// 	Timeout:  60 * time.Second,
	// 	Database: database,
	// 	Username: username,
	// 	Password: password,
	// }
	// c.session, err = mgo.DialWithInfo(info)
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
	// port := 27017
	// return fmt.Sprintf("mongodb://mongodb:%d", port)
	return "mongodb:27017"
	// return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", "amanc",

	// 	"Amanc1101!",
	// 	"127.0.0.1",
	// 	port,
	// 	"Amlak",
	// )

}
