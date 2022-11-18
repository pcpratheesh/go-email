package models

var (
	EmailProviderSMTP = "SMTP"
)

type EmailConfig struct {
	Provider string
	From     string
}

type SMTPCredentials struct {
	Username string
	Password string
	Host     string
}
