package helper

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validation(v *validator.Validate, s interface{}) []string {
	results := make([]string, 0)

	filterError := func(e validator.FieldError) {
		field, _ := reflect.TypeOf(s).Elem().FieldByName(e.Field())
		jsonTag := field.Tag.Get("json")
		queryTag := field.Tag.Get("query")
		validate := field.Tag.Get("validate")

		var res string
		if len(jsonTag) > 0 {
			res = jsonTag
		} else {
			res = queryTag
		}

		switch e.Tag() {
		case "required":
			msg := fmt.Sprintf("%s is required", res)
			results = append(results, msg)
			return
		case "max":
			max := getValueFromTag(validate, "max=")
			msg := fmt.Sprintf("%s is max %v", res, max)
			results = append(results, msg)
			return
		case "min":
			min := getValueFromTag(validate, "min=")
			msg := fmt.Sprintf("%s is min %v", res, min)
			results = append(results, msg)
			return
		}

	}

	if err := v.Struct(s); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				filterError(e)
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
