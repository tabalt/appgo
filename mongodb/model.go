package mongodb

import (
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"fmt"
)
/*
type Model interface {
	Add() 
    Edit()
    Find()
    Exist()
    List()
    Del()
}
*/

func NewMongoModel(name string) *MongoModel {
	return &MongoModel{Mongo: GetMongoDbFromConfig(), Name : name, Error: ""}
}

type MongoModel struct {
	Mongo *mgo.Database
	Name string
	Error string
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
    	result = true;
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

//find
func (this *MongoModel) Find(query interface{}) ([]map[string]interface{}, error){
    itemList := make([]map[string]interface{}, 100)

    err := this.Mongo.C(this.Name).Find(query).All(&itemList)
    return itemList, err
}

//get all
func (this *MongoModel) All() ([]map[string]interface{}, error){
    return this.Find(nil)
}

//check if exists
func (this *MongoModel) Exist(dataSet interface{}) (result bool) {
    result = false
    itemList,_ := this.Find(dataSet)
    if len(itemList) > 0 {
        result = true
    }
    return
}