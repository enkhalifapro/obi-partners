package partners

import (
	"github.com/enkhalifapro/obi-partners/internal/types"
	"github.com/enkhalifapro/obi-partners/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	connector := &mocks.Connector{}
	connector.On("CreateOne", "collName1", &types.Partner{}).Return("", errors.New("invalid coll name"))
	connector.On("CreateOne", "collName", &types.Partner{}).Return("", nil)

	slackProvider := &mocks.SlackProvider{}
	slackProvider.On("Notify","partner added").Return(nil)

	// Act
	mgr := NewManager(connector, slackProvider)
	err := mgr.Add(&types.Partner{})

	// Assert
	assert.Nil(t, err)

}
