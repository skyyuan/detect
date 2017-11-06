package utils

import (
	"reflect"
	"time"
	"math/rand"
	"gopkg.in/mgo.v2/bson"
)

func Contains(slice interface{}, val interface{}) bool {
	sv := reflect.ValueOf(slice)

	for i := 0; i < sv.Len(); i++ {
		if sv.Index(i).Interface() == val {
			return true
		}
	}
	return false
}

func ContainsAny(slice interface{}, val interface{}) bool {
	vv  := reflect.ValueOf(val)
	for i := 0; i < vv.Len(); i++ {
		if Contains(slice, vv.Index(i).Interface()){
			return true
		}
	}
	return false
}


func RemoveSliceBsonId(slice []bson.ObjectId, bid bson.ObjectId) (results []bson.ObjectId) {
	for i := 0; i < len(slice); i++ {
		if  slice[i] != bid{
			results = append(results, slice[i])
		}
	}
	return
}

func RemoveSliceBsonIds(slice []bson.ObjectId, bids []bson.ObjectId) (results []bson.ObjectId) {
	for i := 0; i < len(slice); i++ {
		if  !Contains(bids, slice[i]){
			results = append(results, slice[i])
		}
	}
	return
}


func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func WeekDay(day time.Weekday) string {
	weeDays := []string {
		"周一",
		"周二",
		"周三",
		"周四",
		"周五",
		"周六",
		"周日",
	}
	return weeDays[day]
}
