package goemail

import (
	"github.com/pcpratheesh/go-email/models"
	"github.com/pcpratheesh/go-email/providers"
)

// New provider interface init
func New(cfg models.EmailConfig, credentials interface{}) (providers.EmailProvider, error) {
	return providers.ChooseProvider(cfg, credentials)
}
