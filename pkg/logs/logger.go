package logs

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	ExcludedConfidential = "This has been excluded due to confidentialty"
)

type Logger struct {
	ObjectPrint ObjectPrinter
}

type ObjectPrinter interface {
	PrintField(name string, val interface{})

	Build() string
}

type StringObjectPrinter struct {
	builder *strings.Builder
}

func NewStringObjectPrinter() StringObjectPrinter {
	return StringObjectPrinter{
		builder: &strings.Builder{},
	}
}

func (stringObjectPrinter StringObjectPrinter) PrintField(name string, val interface{}) {
	fmt.Println(fmt.Sprintf("%s: %s,", name, val))
	_, err := stringObjectPrinter.builder.WriteString(fmt.Sprintf("%s: %s,", name, val))
	fmt.Println(err)
}

func (stringObjectPrinter StringObjectPrinter) Build() string {
	defer stringObjectPrinter.builder.Reset()
	str := stringObjectPrinter.builder.String()
	fmt.Println(stringObjectPrinter.builder.String())
	return str
}

func (logger Logger) Print(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Struct {
		//we can do our processing
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			_, ok := field.Tag.Lookup("confidential")
			if !ok {
				//we can log this
				logger.ObjectPrint.PrintField(field.Name, v.Field(i).Interface())
			} else {
				// we will inform the reader that this has been explicitly excluded
				logger.ObjectPrint.PrintField(field.Name, ExcludedConfidential)
			}
		}
		return logger.ObjectPrint.Build()
	}

	return fmt.Sprint(obj)
}
