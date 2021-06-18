package bridge

import "fmt"

type Mail struct {
	from    string
	to      string
	content string
}

func NewMail(from, to, content string) *Mail {
	return &Mail{
		from:    from,
		to:      to,
		content: content,
	}
}

func (m *Mail) Serialize() string {
	return fmt.Sprintf("%s -> %s: %s",
		m.from, m.to, m.content)
}
