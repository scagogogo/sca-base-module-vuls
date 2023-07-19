package vul_dao

import (
	"context"
	"github.com/golang-infrastructure/go-pointer"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

var (
	testVulId    = "vul_id_test"
	testVulCode1 = "vul_code_test_001"
	testVulCode2 = "vul_code_test_002"
	testVulCode3 = "vul_code_test_003"
)

func buildTestVul() *models.Vul {
	return &models.Vul{
		VulId: testVulId,
		CVSS3: "test-cvss3",
		Title: map[language.Tag]string{
			language.Chinese: "中文标题",
			language.English: "英文标题",
		},
		Description: map[language.Tag]string{
			language.Chinese: "中文描述",
			language.English: "英文描述",
		},
		PublishedTime: pointer.Now(),
		CreateTime:    pointer.Now(),
		UpdateTime:    pointer.Now(),
		ChangeTime:    pointer.Now(),
	}
}

func buildTestVulCode(code string) *models.VulCode {
	return &models.VulCode{
		VulId:      testVulId,
		Code:       code,
		CodeType:   models.CodeTypeGHSA,
		CreateTime: pointer.Now(),
		UpdateTime: pointer.Now(),
		ChangeTime: pointer.Now(),
	}
}

// 清空之前的测试数据
func ensureClear(t *testing.T, dao VulDao) {
	err := dao.Delete(context.Background(), testVulId)
	assert.Nil(t, err)
}

func ensureExists(t *testing.T, dao VulDao) {
	err := dao.Upsert(context.Background(), buildTestVul(), []*models.VulCode{buildTestVulCode(testVulCode1)})
	assert.Nil(t, err)
}

func VulDaoTest(t *testing.T, dao VulDao) {
	CreateTest(t, dao)
	CreateCodesTest(t, dao)
	DeleteTest(t, dao)
	DeleteCodeTest(t, dao)
	DeleteCodeByVulIdsTest(t, dao)
	FindTest(t, dao)
	FindCodesTest(t, dao)
	FindManyTest(t, dao)
	LoadAllTest(t, dao)
	LoadAllCodesTest(t, dao)
	ReplaceCodesTest(t, dao)
	UpdateTest(t, dao)
	UpsertTest(t, dao)
	UpsertCodesTes(t, dao)
}

func CreateTest(t *testing.T, dao VulDao) {

	ensureClear(t, dao)

	err := dao.Create(context.Background(), buildTestVul(), []*models.VulCode{
		buildTestVulCode(testVulCode1),
		buildTestVulCode(testVulCode2),
	})
	assert.Nil(t, err)
}

func CreateCodesTest(t *testing.T, dao VulDao) {

	ensureClear(t, dao)

	err := dao.CreateCodes(context.Background(), testVulId, []*models.VulCode{
		buildTestVulCode(testVulCode3),
	})
	assert.Nil(t, err)
}

func DeleteTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	err := dao.Delete(context.Background(), testVulId)
	assert.Nil(t, err)
}

func DeleteCodeTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	err := dao.DeleteCode(context.Background(), testVulCode1)
	assert.Nil(t, err)

}

func DeleteCodeByVulIdsTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	c, err := dao.DeleteCodeByVulId(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.True(t, c != 0)
}

func FindTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	vul, err := dao.Find(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.NotNil(t, vul)

}

func FindCodesTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	codes, err := dao.FindCodes(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.True(t, len(codes) != 0)

}

func FindManyTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	many, err := dao.FindMany(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.True(t, len(many) > 0)
}

func LoadAllTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	all, err := dao.LoadAll(context.Background())
	assert.Nil(t, err)
	assert.True(t, len(all) > 0)

}

func LoadAllCodesTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	codes, err := dao.LoadAllCodes(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, codes)
}

func ReplaceCodesTest(t *testing.T, dao VulDao) {

	ensureClear(t, dao)

	err := dao.ReplaceCodes(context.Background(), testVulId, []*models.VulCode{
		buildTestVulCode(testVulCode1),
		buildTestVulCode(testVulCode2),
		buildTestVulCode(testVulCode3),
	})
	assert.Nil(t, err)

	codes, err := dao.FindCodes(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.True(t, len(codes) >= 3)

	err = dao.ReplaceCodes(context.Background(), testVulId, []*models.VulCode{
		buildTestVulCode(testVulCode2),
		buildTestVulCode(testVulCode3),
	})
	assert.Nil(t, err)

	codes, err = dao.FindCodes(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.True(t, len(codes) >= 2)

}

func UpdateTest(t *testing.T, dao VulDao) {

	ensureClear(t, dao)

	err := dao.Update(context.Background(), buildTestVul(), []*models.VulCode{
		buildTestVulCode(testVulCode1),
		buildTestVulCode(testVulCode2),
		buildTestVulCode(testVulCode3),
	})
	assert.Nil(t, err)

}

func UpsertTest(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	err := dao.Upsert(context.Background(), buildTestVul(), []*models.VulCode{
		buildTestVulCode(testVulCode1),
		buildTestVulCode(testVulCode2),
	})
	assert.Nil(t, err)

}

func UpsertCodesTes(t *testing.T, dao VulDao) {

	ensureExists(t, dao)

	err := dao.UpsertCodes(context.Background(), testVulId, []*models.VulCode{
		buildTestVulCode(testVulCode1),
		buildTestVulCode(testVulCode2),
		buildTestVulCode(testVulCode3),
	})
	assert.Nil(t, err)

	codes, err := dao.FindCodes(context.Background(), testVulId)
	assert.Nil(t, err)
	assert.True(t, len(codes) >= 3)

}
