package inspstru

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Element struct {
	Path  string
	Type  string
	Value string
}

const (
	ANSI_RESET  = "\x1b[0m"
	ANSI_GREEN  = "\x1b[32m"
	ANSI_YELLOW = "\x1b[33m"
	ANSI_CYAN   = "\x1b[36m"
	ANSI_WHITE  = "\x1b[37m"
)

func CollectElements(obj interface{}, prefix string) []Element {
	var elements []Element

	adjustPrefix := func(base, field string) string {
		if base == "" {
			return "." + field
		}
		return base + "." + field
	}

	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValue := v.Field(i).Interface()
			newPrefix := adjustPrefix(prefix, fieldName)

			subElements := CollectElements(fieldValue, newPrefix)
			elements = append(elements, subElements...)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key).Interface()
			keyStr := fmt.Sprintf("%v", key.Interface())
			newPrefix := adjustPrefix(prefix, keyStr)

			subElements := CollectElements(val, newPrefix)
			elements = append(elements, subElements...)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			newPrefix := adjustPrefix(prefix, strconv.Itoa(i))
			subElements := CollectElements(v.Index(i).Interface(), newPrefix)
			elements = append(elements, subElements...)
		}
	default:
		elements = append(elements, Element{
			Path:  prefix,
			Type:  t.Name(),
			Value: fmt.Sprintf("%v", obj),
		})
	}

	return elements
}

func BuildTemplate(obj interface{}, prefix string) string {
	var result strings.Builder

	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValue := v.Field(i).Interface()
			newPrefix := prefix + "." + fieldName
			subTemplate := BuildTemplate(fieldValue, newPrefix)
			result.WriteString(subTemplate)
		}
	case reflect.Map:
		result.WriteString(fmt.Sprintf("{{ range $key, $value := %s }}\n", prefix))
		result.WriteString("\tKey: {{$key}}\n")
		result.WriteString("\tValue: {{ $value }}\n")
		result.WriteString("{{ end }}\n")

	case reflect.Slice, reflect.Array:
		result.WriteString(fmt.Sprintf("{{ range %s }}\n", prefix))
		result.WriteString("\t{{ . }}\n")
		result.WriteString("{{ end }}\n")

	default:
		result.WriteString(fmt.Sprintf("{{ %s }}\n", prefix))
	}

	return result.String()
}

func PrintElements(obj interface{}) {
	r := CollectElements(obj, "")

	sort.Slice(r, func(i, j int) bool {
		return r[i].Path < r[j].Path
	})

	for _, e := range r {
		fmt.Printf("%s%s%s (%s) = %s%s%s\n",
			ANSI_CYAN,
			e.Path,
			ANSI_GREEN,
			e.Type,
			ANSI_YELLOW,
			e.Value,
			ANSI_RESET,
		)
	}
}
