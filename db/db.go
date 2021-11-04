package db

import (
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
	// info := &mgo.DialInfo{
	// 	Addrs:    []string{"mongodb:27017"},
	// 	Database: "Amlak",
	// 	Username: "amanc",
	// 	Password: "Amanc1101!",
	// }
	// c.session, err = mgo.DialWithInfo(info)
	c.session, err = mgo.Dial("mongodb:27017")
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
	return "mongodb://amanc:Amanc1101!@mongodb:27017/Amlak"
	// return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", "amanc",

	// 	"Amanc1101!",
	// 	"127.0.0.1",
	// 	port,
	// 	"Amlak",
	// )

}
