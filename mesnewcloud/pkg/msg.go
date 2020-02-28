package pkg

var Msgflags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	NOTENANTID:     "租户ID不能为空",
}

func GetMsg(code int) string {
	msg, ok := Msgflags[code]
	if ok {
		return msg
	}
	return Msgflags[ERROR]
}
