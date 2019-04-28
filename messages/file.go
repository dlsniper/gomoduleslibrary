package messages

type Hello struct{}

func NewHello() *Hello {
	return &Hello{}
}

func (Hello) Say(lang string) string {
	message := "你好，北京！"
	if lang == "ro" {
		message = "Salut Beijing!"
	}
	return message
}
