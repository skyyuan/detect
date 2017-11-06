package controllers

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"detect/utils"
	"fmt"
	"time"
)

type DetectorController struct {
	beego.Controller
}


// @router /detect [get]
func (this *DetectorController) Detect() {
	detect_id := this.GetString("did")
	watch_id := this.GetString("wid")
	signal ,_ := this.GetInt("sig")
	created_at ,_ := this.GetInt("cre")
	go DetectRecord(detect_id, watch_id, signal, created_at)
	this.Ctx.Output.Body([]byte("1"))
}


func DetectRecord(detect_id, watch_id string, signal, created_at int){
	conn := utils.RedisPool.Get()
	defer conn.Close()
	cutime := time.Now().Format("2006/01/02")
	redisKey := fmt.Sprintf("%s_%s:%s",cutime, watch_id, detect_id)
	conn.Do("ZADD", redisKey, signal, created_at)
	// 保存12天的数据
	conn.Do("EXPIRE", redisKey, 12 * 60 * 60 * 24)
	redisExpirekey := fmt.Sprintf("expire:%s",redisKey)
	exists, _ := redis.Bool(conn.Do("EXISTS", redisExpirekey))
	if exists {
	} else {
		conn.Do("SETEX", redisExpirekey, 300, "1")
		c := time.Tick(300 * time.Second)
		<-c
		utils.ProducerPublish("test",redisKey)
		// 不知道为什么下面用法有问题
		// utils.ProducerPublishDelay("test",redisKey, 300 * time.Second)
	}
}