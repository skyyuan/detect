package utils

import (
	"github.com/astaxie/beego"
	"github.com/franela/goreq"
	"time"
)

const (
	SMSURL = "http://hi.qychbb.com/sms/v1/push"
	APPNAME = "［彩虹宝贝］"
)

type Sms struct {
	SmsType string     `json:"SmsType"`
	SmsSubType string  `json:"SmsSubType"`
	Mobiles []string   `json:"Mobiles"`
	Yzm string         `json:"Yzm"`
	Message string     `json:"Message"`
}

type SmsResponse struct {
	Code string `json:"code"`
	State string `json:"state"`
}

// 只支持基本的注册验证码,不是所以的都支持
func SendSmsMessage(mobile string, message string) error {
	sms := Sms{SmsType: "qyys", SmsSubType: "yzm", Mobiles: []string{mobile}, Yzm: message}
	req := goreq.Request{
		Method: "POST",
		Uri:    SMSURL,
		//ContentType: "text/plain;charset=UTF-8",
		Body:    sms,
		Timeout: 5 * time.Second,
	}
	res, err := req.Do()
	defer closeGoreqResponse(res, err)
	go NewMobileMessageLog(mobile, message)
	return err
}

func SendSmsOtherMessage(mobile string, message string) error {
	sms := Sms{Mobiles: []string{mobile}, Message: message}
	req := goreq.Request{
		Method: "POST",
		Uri:    SMSURL,
		//ContentType: "text/plain;charset=UTF-8",
		Body:    sms,
		Timeout: 5 * time.Second,
	}
	res, err := req.Do()
	defer closeGoreqResponse(res, err)
	go NewMobileMessageLog(mobile, message)
	return err
}

func closeGoreqResponse(resp *goreq.Response, err error) {
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	} else {
		beego.Error(err)
	}
}
