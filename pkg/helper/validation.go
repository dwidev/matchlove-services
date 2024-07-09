package helper

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validation(v *validator.Validate, s interface{}) []string {
	results := make([]string, 0)

	filterError := func(e validator.FieldError) {
		fieldPath := strings.Split(e.StructNamespace(), ".")
		typ := reflect.TypeOf(s).Elem()

		var field reflect.StructField
		var jsonPath []string

		for _, part := range fieldPath {
			var found bool
			field, found = typ.FieldByName(part)
			if !found {
				continue
			}

			jsonPath = append(jsonPath, field.Tag.Get("json"))

			typ = field.Type
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}
		}
		res := strings.Join(jsonPath, " -> ")
		validate := field.Tag.Get("validate")

		switch e.Tag() {
		case "required":
			msg := fmt.Sprintf("%s is required", res)
			results = append(results, msg)
			return
		case "max":
			maxVal := getValueFromTag(validate, "max=")
			msg := fmt.Sprintf("%s is max %v", res, maxVal)
			results = append(results, msg)
			return
		case "min":
			minVal := getValueFromTag(validate, "min=")
			msg := fmt.Sprintf("%s is min %v", res, minVal)
			results = append(results, msg)
			return
		}

	}

	if err := v.Struct(s); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, ev := range ve { // ev is error validation
				filterError(ev)
			}
		}
	}

	return results
}

func getValueFromTag(tag string, prefix string) string {
	tagValues := strings.Split(tag, ",")

	for _, value := range tagValues {
		if strings.HasPrefix(value, prefix) {
			maxValueStr := strings.TrimPrefix(value, prefix)
			return maxValueStr
		}
	}

	return ""
}
