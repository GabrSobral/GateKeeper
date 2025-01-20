package errors

import (
	"fmt"
	"reflect"
)

type InvalidRequestBody struct {
	Schema interface{} `json:"schema"`
}

type InvalidRequestBodyResponse struct {
	Title   string                    `json:"title"`
	Message string                    `json:"message"`
	Fields  []InvalidRequestBodyField `json:"fields"`
}

type InvalidRequestBodyField struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (s *InvalidRequestBody) Error() string {
	return "Invalid Request Body"
}

func (s *InvalidRequestBodyResponse) Error() string {
	return s.Message
}

// Function to parse the schema and return InvalidRequestBodyResponse
func NewInvalidBodyResponse(schema *InvalidRequestBody) (*InvalidRequestBodyResponse, error) {
	s := &InvalidRequestBodyResponse{Title: "Invalid Request Body", Message: "Invalid Request Body"}

	t := reflect.TypeOf(schema.Schema)

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("schema is not a struct")
	}

	// Iterate through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type.String()

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name // Use field name if json tag is not specified
		}

		s.Fields = append(s.Fields, InvalidRequestBodyField{
			Field: jsonTag,
			Type:  fieldType,
		})
	}

	return s, nil
}
