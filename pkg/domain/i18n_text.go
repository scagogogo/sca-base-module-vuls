package domain

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

func (x I18nText) Value() (driver.Value, error) {
	if x == nil {
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
