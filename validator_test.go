package validator

import (
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestLcFirst(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"Roshan", "roshan"},
		{"Rana", "rana"},
		{"Test", "test"},
	}
	for _, test := range tests {
		if output := LcFirst(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestUcFirst(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"roshan", "Roshan"},
		{"rana", "Rana"},
		{"test", "Test"},
	}
	for _, test := range tests {
		if output := UcFirst(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"myName", "My name"},
		{"testDriven", "Test driven"},
		{"myOutput", "My output"},
	}
	for _, test := range tests {
		if output := Split(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

type FirstNameValidationTestStruct struct {
	FirstName string `validate:"required" json:"firstName"`
}

type ThisTestVlaidationStruct struct {
	ThisTest string `validate:"required" json:"thisTest"`
}

func TestValidationErrorToText(t *testing.T) {
	extraValidator := []ExtraValidation{
		{Tag: "required", Message: "%s is passed!"},
	}
	MakeExtraValidation(extraValidator)

	validates := validator.New()
	err := validates.Struct(FirstNameValidationTestStruct{})
	error1 := err.(validator.ValidationErrors)

	err2 := validates.Struct(ThisTestVlaidationStruct{})
	error2 := err2.(validator.ValidationErrors)

	var tests = []struct {
		input    validator.ValidationErrors
		expected string
	}{
		{error1, "First name is passed!"},
		{error2, "This test is passed!"},
	}

	for _, test := range tests {
		if output := ValidationErrorToText(test.input[0]); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}

	}

}
