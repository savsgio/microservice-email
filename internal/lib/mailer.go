package lib

import "gopkg.in/gomail.v2"

type Email struct {
	To          []string
	Subject     string
	Body        string
	ContentType string `json:"content_type"`
}

func (m *Email) Send() error {
	gm := gomail.NewMessage()
	gm.SetHeader("From", Conf.Smtp.User)
	gm.SetHeader("To", m.To...)
	gm.SetHeader("Subject", m.Subject)
	gm.SetBody(m.ContentType, m.Body)

	d := gomail.NewDialer(Conf.Smtp.Host, Conf.Smtp.Port, Conf.Smtp.User, Conf.Smtp.Password)

	err := d.DialAndSend(gm)

	return err
}
