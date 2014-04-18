package mongodb

import (
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"fmt"
)

func NewMongoModel(name string) *MongoModel {
	return &MongoModel{Mongo: GetMongoDbFromConfig(), Name : name, Error: ""}
}

type MongoModel struct {
	Mongo *mgo.Database
	Name string
	Error string
    Condition interface{}
}

//add 
func (this *MongoModel) Add(dataSet interface{}) (result bool){
	fmt.Println(dataSet)
	err := this.Mongo.C(this.Name).Insert(dataSet)
    if err != nil {
    	//TODO deal exception
    	//jLogger.Fatal(err.Error())
    	result = false
    } else {
    	result = true
    	//TODO　return id
    }
    return
}

//unique add 
func (this *MongoModel) UniqueAdd(dataSet interface{}) (result bool){
    if this.Exist(dataSet) {
    	this.Error = "";
        result = false
    } else {
    	result = this.Add(dataSet)
    }
    return
}

//condition
func (this *MongoModel) Condition(condition interface{}) (*MongoModel){
    this.Condition = condition
    return this
}

//select
func (this *MongoModel) Select(infoList *interface{}) (result bool){
    err := this.Mongo.C(this.Name).Find(this.Condition).All(&infoList)
    if err != nil {
        result = false
    } else {
        result = true
    }
    return 
}

//find
func (this *MongoModel) Find(info *interface{}) (result bool){
    err := this.Mongo.C(this.Name).Find(this.Condition).One(&info)
    if err != nil {
        result = false
    } else {
        result = true
    }
    return 
}

/*
//find
func (this *MongoModel) Find(query interface{}) ([]map[string]interface{}, error){
    itemList := make([]map[string]interface{}, 100)

    err := this.Mongo.C(this.Name).Find(query).All(&itemList)
    return itemList, err
}
*/

//check if exists
func (this *MongoModel) Exist(dataSet interface{}) (result bool) {
    result = false
    itemList,_ := this.GetList(dataSet)
    if len(itemList) > 0 {
        result = true
    }
    return
}