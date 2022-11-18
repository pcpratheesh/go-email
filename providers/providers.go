package providers

import (
	"fmt"

	"github.com/pcpratheesh/go-email/models"
)

type EmailProvider interface {
	WithTemplate(string) EmailProvider
	RenderTemplate(string) error
	Send(string, string) error
}

type UnimplementedServer struct{}

func (ims *UnimplementedServer) RenderTemplate(string) error {
	return fmt.Errorf("RenderTemplate not implemented")
}

func (ims *UnimplementedServer) Send(string, string) error {
	return fmt.Errorf("Send not implemented")
}

func (ims *UnimplementedServer) WithTemplate(string) EmailProvider {
	return nil
}

// choose provider
func ChooseProvider(cfg models.EmailConfig, credentials interface{}) (EmailProvider, error) {
	switch cfg.Provider {
	case models.EmailProviderSMTP:
		return NewSMTPProvider(credentials.(models.SMTPCredentials)), nil
	}

	return nil, fmt.Errorf("unable to identify the provider")
}
