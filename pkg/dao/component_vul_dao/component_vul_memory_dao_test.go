package component_vul_dao

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponentVulMemoryDao(t *testing.T) {
	dao, err := NewComponentVulMemoryDaoFromJsonLine(context.Background(), []byte{})
	assert.Nil(t, err)
	assert.NotNil(t, dao)
	ComponentVulDaoTest(t, dao)
}
