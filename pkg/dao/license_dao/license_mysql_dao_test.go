package license_dao

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLicense(t *testing.T) {
	//license, err := FindLicense(context.Background(), "AGPL-3.0-only")
	//assert.Nil(t, err)
	//assert.NotNil(t, license)
}

func TestLoadAllLicenses(t *testing.T) {
	licenses, err := LoadAllLicenses(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, licenses)
}

func TestSaveLicense(t *testing.T) {

}
