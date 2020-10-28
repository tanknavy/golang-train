package message

//消息类型
const (
	LoginMsgType    = "LoginMsg"
	loginResMsgType = "LoginResMsg"
)

type Message struct {
	Type string `json:"type"` //消息类型，使用常量, 使用tag,大写为了export, 传递的/序列化时小写
	Data string `json:"data"` //消息的内容
}

//login message
type LoginMsg struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

//response message
type LoginResMsg struct {
	Code  int    `json:"code"`  //返回状态码， 500:未注册，200:登录成功
	Error string `json:"error"` //返回错误信息
}
