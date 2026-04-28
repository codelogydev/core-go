package mailer

import "github.com/wneessen/go-mail"

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var cfg Config

func Init(c Config) {
	cfg = c
}

func Send(to, subject, body string) error {
	m := mail.NewMsg()

	if err := m.From(cfg.From); err != nil {
		return err
	}
	if err := m.To(to); err != nil {
		return err
	}

	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, body)

	c, err := mail.NewClient(cfg.Host,
		mail.WithPort(cfg.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Username),
		mail.WithPassword(cfg.Password),
	)
	if err != nil {
		return err
	}

	return c.DialAndSend(m)
}
