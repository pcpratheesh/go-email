package providers

import "github.com/pcpratheesh/go-email/models"

type SMTPProvider struct {
	UnimplementedServer
	Credentials models.SMTPCredentials
}

func NewSMTPProvider(cred models.SMTPCredentials) EmailProvider {
	return &SMTPProvider{
		Credentials: cred,
	}
}
