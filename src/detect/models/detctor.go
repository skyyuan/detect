package models

import "gopkg.in/mgo.v2/bson"

type Detecor struct {
	Id_         bson.ObjectId `bson:"_id"`
	DeviceId  string        `bson:"device_id"`
	CameraId  string        `bson:"camera_id"`
	PrevDetectorId    bson.ObjectId        `bson:"prev_detector_id"`
	NextDetectorId    bson.ObjectId        `bson:"next_detector_id"`
	Location   string        `bson:"location"`
	CommonModel `bson:",inline"`
}
