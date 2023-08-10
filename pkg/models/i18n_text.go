package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"reflect"
)

// I18nText 支持国际化的文本
type I18nText map[language.Tag]string

var _ sql.Scanner = &I18nText{}
var _ driver.Valuer = &I18nText{}

func NewI18nText() I18nText {
	return make(map[language.Tag]string)
}

func (x I18nText) MergeAll(i18nText I18nText) I18nText {
	for language, text := range i18nText {
		x = x.Merge(language, text)
	}
	return x
}

func (x I18nText) SetAll(i18nText I18nText) I18nText {
	for language, text := range i18nText {
		x = x.Set(language, text)
	}
	return x
}

// Merge 仅当给定的language不存在的时候才设置，如果已经存在的话则忽略
func (x I18nText) Merge(language language.Tag, text string) I18nText {
	if x == nil {
		x = NewI18nText()
	}
	if _, exists := x[language]; !exists {
		x[language] = text
	}
	return x
}

// Set 不管给定的language存不存在都会更新为value，已经存在的话会被覆盖掉
func (x I18nText) Set(language language.Tag, text string) I18nText {
	if x == nil {
		x = NewI18nText()
	}
	x[language] = text
	return x
}

func (x I18nText) Append(language language.Tag, text string) I18nText {
	x[language] = text
	return x
}

func (x I18nText) Value() (driver.Value, error) {
	if len(x) == 0 {
		return nil, nil
	}
	return json.Marshal(x)
}

func (x *I18nText) Scan(src any) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("can not scan from %s to %s", reflect.TypeOf(src).Name(), reflect.TypeOf(x).Name())
	}
	return json.Unmarshal(bytes, &x)
}
