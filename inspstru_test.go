package inspstru_test

import (
	"fmt"

	"github.com/crgimenes/inspstru"
)

func ExamplePrintElements() {
	type Inner struct {
		InnerValue string
	}
	type Outer struct {
		OuterValue string
		Inner      Inner
	}

	obj := Outer{
		OuterValue: "outer",
		Inner: Inner{
			InnerValue: "inner",
		},
	}

	inspstru.PrintElements(obj, false)
	// Output:
	// .Inner.InnerValue (string) = inner
	// .OuterValue (string) = outer
}

func ExampleBuildTemplate() {
	type Inner struct {
		InnerValue string
	}
	type Outer struct {
		OuterValue string
		Inner      Inner
	}

	obj := Outer{
		OuterValue: "outer",
		Inner: Inner{
			InnerValue: "inner",
		},
	}

	templateStr := inspstru.BuildTemplate(obj, "")

	fmt.Print(templateStr)
	// Output:
	// {{ .OuterValue }}
	// {{ .Inner.InnerValue }}
}
