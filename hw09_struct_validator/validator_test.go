package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			App{
				Version: "333",
			},
			ValidationErrors{
				ValidationError{Err: ErrValidationLength, Field: "Version"},
			},
		},
		{
			App{
				Version: "12345",
			},
			nil,
		},
		{
			User{
				ID:     "100500",
				Name:   "Testov",
				Age:    62,
				Email:  "totaly.wrong",
				Role:   "admin",
				Phones: []string{"12345678910"},
				meta:   nil,
			},
			ValidationErrors{
				ValidationError{Err: ErrValidationLength, Field: "ID"},
				ValidationError{Err: ErrValidationMaximum, Field: "Age"},
				ValidationError{Err: ErrValidationRegexp, Field: "Email"},
			},
		},
		{
			Token{
				Header:    nil,
				Payload:   nil,
				Signature: nil,
			},
			nil,
		},
		{
			Response{
				Code: 410,
			},
			ValidationErrors{
				ValidationError{
					Err:   ErrValidationContains,
					Field: "Code",
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)

			expectedErr, _ := tt.expectedErr.(ValidationErrors)
			actualErr, _ := err.(ValidationErrors)

			require.Len(t, actualErr, len(expectedErr))

			for i, _ := range expectedErr {
				require.Equal(t, expectedErr[i].Field, actualErr[i].Field)
				require.Equal(t, expectedErr[i].Err, actualErr[i].Err)
			}
		})
	}
}
