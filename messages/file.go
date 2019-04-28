package messages

type Hello struct{}

func NewHello() *Hello {
	return &Hello{}
}

func (Hello) Say() string {
	message := "你好，北京！"
	return message
}
