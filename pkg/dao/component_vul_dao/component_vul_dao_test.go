package component_vul_dao

import (
	"context"
	"github.com/golang-infrastructure/go-pointer"
	"github.com/scagogogo/sca-base-module-vuls/pkg/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testName    = "test-name"
	testVersion = "test-version"
	testVulId   = "test-vul-id"
)

func buildComponentVul() *domain.ComponentVul {
	return &domain.ComponentVul{
		Name:       testName,
		Version:    testVersion,
		VulId:      testVulId,
		CreateTime: pointer.Now(),
		ChangeTime: pointer.Now(),
		UpdateTime: pointer.Now(),
	}
}

func ensureClear(t *testing.T, dao ComponentVulDao) {
	_, _ = dao.DeleteByComponentNameAndVersion(context.Background(), testName, testVersion)
}

func ensureExists(t *testing.T, dao ComponentVulDao) {
	_ = dao.Create(context.Background(), buildComponentVul())
}

func ComponentVulDaoTest(t *testing.T, dao ComponentVulDao) {
	CreateTest(t, dao)
	DeleteByComponentNameTest(t, dao)
	DeleteByComponentNameAndVersionTest(t, dao)
	DeleteByVulIdTest(t, dao)
	Find(t, dao)
	FindByComponentNameTest(t, dao)
	LoadAllTest(t, dao)
	UpdateTest(t, dao)
	UpsertTest(t, dao)
}

func CreateTest(t *testing.T, dao ComponentVulDao) {

	ensureClear(t, dao)

	err := dao.Create(context.Background(), buildComponentVul())
	assert.Nil(t, err)

	cv, err := dao.Find(context.Background(), testName, testVersion)
	assert.Nil(t, err)
	assert.NotNil(t, cv)
	assert.Equal(t, testVulId, cv[0].VulId)
}

func DeleteByComponentNameTest(t *testing.T, dao ComponentVulDao) {

	ensureExists(t, dao)

	c, err := dao.DeleteByComponentName(context.Background(), testName)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, c)

}

func DeleteByComponentNameAndVersionTest(t *testing.T, dao ComponentVulDao) {
	ensureExists(t, dao)

	c, err := dao.DeleteByComponentNameAndVersion(context.Background(), testName, testVersion)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, c)
}

func DeleteByVulIdTest(t *testing.T, dao ComponentVulDao) {
	ensureExists(t, dao)

	c, err := dao.DeleteByVulId(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, c)
}

func Find(t *testing.T, dao ComponentVulDao) {
	ensureExists(t, dao)

	c, err := dao.DeleteByVulId(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, c)
}

func FindByComponentNameTest(t *testing.T, dao ComponentVulDao) {
	ensureExists(t, dao)

	cv, err := dao.FindByComponentName(context.Background(), testName)
	assert.Nil(t, err)
	assert.NotNil(t, cv)
	assert.Equal(t, testVulId, cv[0].VulId)

}

func LoadAllTest(t *testing.T, dao ComponentVulDao) {
	ensureExists(t, dao)

	all, err := dao.LoadAll(context.Background())
	assert.Nil(t, err)
	assert.True(t, len(all) > 0)
}

func UpdateTest(t *testing.T, dao ComponentVulDao) {
	ensureExists(t, dao)

	err := dao.Update(context.Background(), buildComponentVul())
	assert.Nil(t, err)

	ensureClear(t, dao)

	err = dao.Update(context.Background(), buildComponentVul())
	assert.Nil(t, err)
}

func UpsertTest(t *testing.T, dao ComponentVulDao) {
	ensureExists(t, dao)

	err := dao.Upsert(context.Background(), buildComponentVul())
	assert.Nil(t, err)

	err = dao.Upsert(context.Background(), buildComponentVul())
	assert.Nil(t, err)

	ensureClear(t, dao)

	err = dao.Upsert(context.Background(), buildComponentVul())
	assert.Nil(t, err)
}
