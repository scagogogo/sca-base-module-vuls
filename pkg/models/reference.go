package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/golang-infrastructure/go-pointer"
	osv_schema "github.com/scagogogo/osv-schema"
	"golang.org/x/text/language"
	"reflect"
)

// ---------------------------------------------------------------------------------------------------------------------

type References []*Reference

var _ sql.Scanner = &References{}
var _ driver.Valuer = &References{}

// FilterByLanguage 根据引用的语言来过滤，一般用于国际化
func (x References) FilterByLanguage(language language.Tag) References {

	newReferences := make([]*Reference, 0)
	for _, r := range x {
		if r.Language == nil {
			continue
		}
		if pointer.FromPointer(r.Language) == language {
			newReferences = append(newReferences, r)
		}
	}
	return newReferences
}

// NewReferencesFromOsv 从OSV的引用中创建sca的支持i18n的引用
func NewReferencesFromOsv(language language.Tag, osvReferences osv_schema.References) References {
	scaReferences := make(References, len(osvReferences))
	for i, osvReference := range osvReferences {
		scaReferences[i] = NewReferenceFromOsv(language, osvReference)
	}
	return scaReferences
}

// FilterByType 根据引用文章的类型来过滤
func (x References) FilterByType(referenceTypes ...osv_schema.ReferenceType) References {

	if len(referenceTypes) == 0 {
		return nil
	}

	referenceTypeSet := make(map[osv_schema.ReferenceType]struct{}, 0)
	for _, r := range referenceTypes {
		referenceTypeSet[r] = struct{}{}
	}

	slice := make([]*Reference, 0)
	for _, r := range x {
		if _, exists := referenceTypeSet[r.Reference.Type]; exists {
			slice = append(slice, r)
		}
	}
	return slice
}

// ContainsUrl 判断引用中是否包含给定的URL
func (x References) ContainsUrl(url string) bool {
	if len(x) == 0 {
		return false
	}
	for _, r := range x {
		if r.URL == url {
			return true
		}
	}
	return false
}

func (x *References) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("scan error")
	}
	if len(bytes) == 0 {
		return nil
	}
	return json.Unmarshal(bytes, &x)
}

func (x References) Value() (driver.Value, error) {
	if len(x) == 0 {
		return nil, nil
	}
	marshal, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	if len(marshal) == 0 {
		return nil, nil
	}
	return string(marshal), nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Reference 漏洞的引用文章，在OSV的Reference的基础上扩展了语言字段，以便能够在国际化的时候根据不同的语言返回不同的引用文章
type Reference struct {
	*osv_schema.Reference `mapstructure:",squash" json:",squash" yaml:",squash" db:",squash" bson:",squash" gorm:"column:type"`
	Language              *language.Tag `mapstructure:"language" json:"language" yaml:"language" db:"language" bson:"language" gorm:"column:language"`
}

var _ sql.Scanner = &Reference{}
var _ driver.Valuer = &Reference{}

// NewReferenceFromOsv 从OSV的引用中创建sca的支持i18n的引用
func NewReferenceFromOsv(language language.Tag, reference *osv_schema.Reference) *Reference {
	return &Reference{
		Reference: reference,
		Language:  pointer.ToPointer(language),
	}
}

func (x *Reference) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *Reference) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not scan from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}

// ---------------------------------------------------------------------------------------------------------------------
