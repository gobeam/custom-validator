package validator

import (
	"gopkg.in/go-playground/validator.v8"
	"testing"
)


func TestLcFirst(t *testing.T) {
	var tests = []struct {
		input string
		expected string
	} {
		{"Roshan","roshan"},
		{"Rana","rana"},
		{"Test","test"},
	}
	for _, test := range tests {
		if output := LcFirst(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}


func TestUcFirst(t *testing.T) {
	var tests = []struct {
		input string
		expected string
	} {
		{"roshan","Roshan"},
		{"rana","Rana"},
		{"test","Test"},
	}
	for _, test := range tests {
		if output := UcFirst(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}


func TestSplit(t *testing.T) {
	var tests = []struct {
		input string
		expected string
	} {
		{"myName","My name"},
		{"testDriven","Test driven"},
		{"myOutput","My output"},
	}
	for _, test := range tests {
		if output := Split(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestValidationErrorToText(t *testing.T) {
	validate := []ExtraValidation{
		{Tag: "test", Message:"%s is passed!"},
	}
	makeExtraValidation(validate)

	error1 := validator.FieldError{}
	error1.Tag = "required"
	error1.Field = "firstName"


	error2 := validator.FieldError{}
	error2.Tag = "test"
	error2.Field = "thisTest"

	var tests = []struct {
		input *validator.FieldError
		expected string
	} {
		{&error1, "First name is required!"},
		{&error2, "This test is passed!"},
	}

	for _, test := range tests {
		if output := ValidationErrorToText(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}

	}

}