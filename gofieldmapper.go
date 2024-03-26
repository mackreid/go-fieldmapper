package gofieldmapper

import (
	"errors"
	"reflect"
	"strings"
)

const package_tag = "fmap"

func Get(o any) (map[string]any, error) {
	rm := make(map[string]any)
	val := reflect.ValueOf(o)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, errors.New("must be a struct or pointer to a struct")
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldTyp := typ.Field(i)
		tag := fieldTyp.Tag.Get(package_tag)
		if tag == "-" || tag == "" {
			continue
		}
		tagn, tago := splitTag(tag)
		fVal := field.Interface()
		if (fieldTyp.Type.Kind() == reflect.Slice || fieldTyp.Type.Kind() == reflect.Array) && reflect.ValueOf(fVal).Len() == 0 {
			continue
		}
		if fVal == reflect.Zero(fieldTyp.Type).Interface() && strings.Contains(tago, "omitempty") {
			continue
		}
		rm[tagn] = fVal
	}
	return rm, nil
}

func splitTag(s string) (string, string) {
	sp := strings.Split(s, ",")
	return sp[0], strings.Join(sp[1:], ",")
}
