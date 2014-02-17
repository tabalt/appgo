package mongodb

import (
	"labix.org/v2/mgo"
	"os"
	. "github.com/tabalt/appgo"
	"fmt"
)

// init mongo
func GetMongoDb(host string, port string, user string, password string, database string) *mgo.Database {
    addr := host + ":" + port
    session, err := mgo.Dial(addr)
    if err != nil {
        //TODO deal exception
        //jLogger.Fatal(err.Error())
        fmt.Println(err.Error())
        os.Exit(1)
    }
    mongoDb := session.DB(database)
    mongoDb.Login(user, password)
    //defer session.Close()
    return mongoDb
}

func GetMongoDbFromConfig() *mgo.Database {

	host, _ := AppConfig.String("mongo_host")
	port, _ := AppConfig.String("mongo_port")
    user, _ := AppConfig.String("mongo_user")
    password, _ := AppConfig.String("mongo_password")
    database, _ := AppConfig.String("mongo_database")
    
    return GetMongoDb(host, port, user, password, database);
}

