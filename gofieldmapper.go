package gofieldmapper

import (
	"errors"
	"reflect"
	"slices"
	"strings"
)

const package_tag = "fmap"

type FieldRule func(value any, tag string, opts string) (k string, v any, ok bool)

func WithOmit() FieldRule {
	return func(value any, tag string, opts string) (k string, v any, ok bool) {
		o := GetOptKeyValue(opts, "omitempty")
		if o == "" {
			return tag, value, true
		}
		if reflect.ValueOf(value).IsZero() {
			return "", "", false
		}
		return tag, value, true
	}
}

func WithMask(mask []string) FieldRule {
	return func(value any, tag string, opts string) (k string, v any, ok bool) {
		o := GetOptKeyValue(opts, "mask")
		if !slices.Contains(mask, o) {
			return "", "", false
		}
		return tag, value, true
	}
}

func Make(obj any, rules ...FieldRule) (map[string]any, error) {
	rm := make(map[string]any)
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Pointer {
		objValue = objValue.Elem()
	}
	if objValue.Kind() != reflect.Struct {
		return nil, errors.New("must be a struct or pointer to a struct")
	}
	objTyp := objValue.Type()

fields:
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldTyp := objTyp.Field(i)
		tag := fieldTyp.Tag.Get(package_tag)
		if tag == "-" || tag == "" {
			continue
		}
		n, o := parseTag(tag)
		key := n
		val := field.Interface()
		for _, rule := range rules {
			k, v, ok := rule(val, n, o)
			if !ok {
				continue fields
			}
			key = k
			val = v
		}
		rm[key] = val
	}
	return rm, nil
}

func GetOptKeyValue(tag string, key string) string {
	parts := strings.Split(tag, ",")
	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		if kv[0] == key {
			if len(kv) > 1 {
				return kv[1]
			}
			return kv[0]
		}
	}
	return ""
}

func parseTag(tag string) (string, string) {
	pt := strings.Split(tag, ",")
	return pt[0], strings.Join(pt[1:], ",")
}
