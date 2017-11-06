package utils

type QYError struct {
	Code int "json: code"
	//Message string "json: message"
}

var QYErrorMap = map[int]string{
	//system
	10001: "签名错误",
	10002: "请求已经过期",
	10003: "签名参数没有",
	10004: "没有权限访问",
	//user
	20000: "缺少参数",
	20100: "用户名或密码不对",
	20101: "没有这个手机号码的用户",
	20102: "用户已经注册，请直接登录",
	20103: "验证码不对",
	20104: "已经具备家长身份",
	20105: "该手机号码的用户不存在",
	20106: "手机号码格式不对",
	20007: "您的版本太低，无法登录",
	20108: "两次输入的密码不一致，请重新输入",
	20109: "原密码不对，请重新输入",
}

func(qyEroor *QYError)GetCode() string{
	return QYErrorMap[qyEroor.Code]
}


func(qyEroor QYError)Error() string{
	return QYErrorMap[qyEroor.Code]
}