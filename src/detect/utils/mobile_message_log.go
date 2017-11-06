package utils

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

type MobileMessageLog struct {
	Id_         bson.ObjectId `bson:"_id"`
	Mobile      string `bson:"mobile"`
	Message     string `bson:"message"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func NewMobileMessageLog(mobile , message string) error{
	mdb, mSession := GetMgoDbSession()
	defer mSession.Close()
	var log MobileMessageLog
	mobileMessageLogCollection := mdb.C("mobile_message_logs")
	err := mobileMessageLogCollection.Find(bson.M{"mobile": mobile}).One(&log)
	if err == mgo.ErrNotFound {
		currentTime := bson.Now()
		log.Id_ = bson.NewObjectId()
		log.Mobile = mobile
		log.Message = message
		log.CreatedAt = currentTime
		log.UpdatedAt = currentTime
		err := mobileMessageLogCollection.Insert(&log)
		return err
	} else if err != nil {
		beego.Error("Something is wrong" + err.Error())
		return err
	} else {
		currentTime := bson.Now()
		err = mobileMessageLogCollection.Update(bson.M{"_id": log.Id_}, bson.M{"$set": bson.M{"message": message, "updated_at": currentTime}})
		return err
	}
}

