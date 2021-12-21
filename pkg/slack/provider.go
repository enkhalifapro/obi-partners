package slack

type Provider struct {
}

func NewProvider(webhookURL string) *Provider {
	return &Provider{}
}

func (p *Provider) Notify(msg string) error {
	return nil
}
