package bridge

import "fmt"

type Mailman interface {
	SendMail(mail Mail) string
}

type CommonMailman struct {
	sender Sender
}

func (m *CommonMailman) SendMail(mail Mail) string {
	msg := fmt.Sprintf("[ ] %s", mail.Serialize())
	return m.sender.Send(msg)
}

func NewCommonMailman(sender Sender) *CommonMailman {
	return &CommonMailman{
		sender: sender,
	}
}

type SpecialMailman struct {
	sender Sender
}

func (m *SpecialMailman) SendMail(mail Mail) string {
	msg := fmt.Sprintf("[*] %s", mail.Serialize())
	return m.sender.Send(msg)
}

func NewSpecialMailman(sender Sender) *SpecialMailman {
	return &SpecialMailman{
		sender: sender,
	}
}
