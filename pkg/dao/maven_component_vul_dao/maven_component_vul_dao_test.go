package maven_component_vul_dao

import (
	"github.com/golang-infrastructure/go-pointer"
	"github.com/scagogogo/sca-base-module-vuls/pkg/models"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpsertMavenComponentVul(t *testing.T) {
	err := UpsertMavenComponentVul(&models.MavenComponentVul{
		GroupId:    "test-group-id",
		ArtifactId: "test-artifact-id",
		Version:    "test-version",
		VulId:      "test-vul-id",
		CreateTime: pointer.Now(),
		UpdateTime: pointer.Now(),
	})
	assert.Nil(t, err)
}

func TestFindMavenComponentVul(t *testing.T) {

	_ = UpsertMavenComponentVul(&models.MavenComponentVul{
		GroupId:    "test-group-id",
		ArtifactId: "test-artifact-id",
		Version:    "test-version",
		VulId:      "test-vul-id",
		CreateTime: pointer.Now(),
		UpdateTime: pointer.Now(),
	})

	vul, err := FindMavenComponentVul("test-group-id", "test-artifact-id", "test-version", "test-vul-id")
	assert.Nil(t, err)
	assert.NotNil(t, vul)
}
