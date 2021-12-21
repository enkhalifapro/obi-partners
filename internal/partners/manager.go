package partners

import "github.com/enkhalifapro/obi-partners/internal/types"

// Manager ...
type Manager struct {
	connector     Connector
	slackProvider SlackProvider
}

type Connector interface {
	CreateOne(collectionName string, document interface{}) (string, error)
	FindOne(collectionName string, filter interface{}, res interface{}) error
}

type SlackProvider interface {
	Notify(msg string) error
}

func NewManager(connector Connector, slackProvider SlackProvider) *Manager {
	return &Manager{
		connector:     connector,
		slackProvider: slackProvider,
	}
}

// Add ...
func (m *Manager) Add(partner *types.Partner) error {
	_, err := m.connector.CreateOne("collName", partner)
	if err != nil {
		return err
	}
	return m.slackProvider.Notify("partner added")
}

func (m *Manager) List() ([]types.Partner, error) {
	return nil, nil
}

func (m *Manager) Find(id string) (*types.Partner, error) {
	return nil, nil
}

func (m *Manager) Query(partner *types.Partner) ([]types.Partner, error) {
	return nil, nil
}

func (m *Manager) Remove(id string) error {
	return nil
}
