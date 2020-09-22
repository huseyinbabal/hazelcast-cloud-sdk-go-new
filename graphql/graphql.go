package graphql

import (
	"encoding/json"
	"fmt"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"reflect"
	"strings"
	"unicode"
)

//This function creates a graphql query from provided query name, operation, input, arguments and the type of the response.
func Query(name string, operation models.GraphqlOperation, input interface{}, args interface{}, response interface{}) string {
	if args != nil {
		return fmt.Sprintf("%s{response:%s(%s)%s}",
			operation,
			name,
			argumentsBuilder(args),
			responseBuilder(response))
	} else if input != nil {
		inputNameSlice := strings.Split(reflect.TypeOf(input).String(), ".")
		nameInput := inputNameSlice[len(inputNameSlice)-1]
		return fmt.Sprintf("%s($input:%s){response:%s(input:$input)%s}",
			operation,
			nameInput,
			name,
			responseBuilder(response))
	} else {
		return fmt.Sprintf("%s{response:%s%s}",
			operation,
			name,
			responseBuilder(response))
	}
}

//This function creates response string from interface for graphql query
func responseBuilder(value interface{}) string {
	var stringBuilder strings.Builder
	buildResponse(reflect.TypeOf(value), &stringBuilder)
	return stringBuilder.String()
}

//This function creates response recursively and writes to writer
func buildResponse(t reflect.Type, stringBuilder *strings.Builder) {
	kind := t.Kind()
	if kind == reflect.Slice {
		buildResponse(t.Elem(), stringBuilder)
	} else if kind == reflect.Struct {
		fmt.Fprint(stringBuilder,"{")
		for i := 0; i < t.NumField(); i++ {
			if i != 0 {
				fmt.Fprint(stringBuilder, ",")
			}
			f := t.Field(i)
			fmt.Fprint(stringBuilder, camelCase(f.Name))
			buildResponse(f.Type, stringBuilder)
		}
		fmt.Fprint(stringBuilder,"}")
	}
}

//This function creates arguments for graphql query
func argumentsBuilder(args interface{}) string {
	valueOf := reflect.ValueOf(args)
	var stringBuilder strings.Builder
	for i := 0; i < valueOf.NumField(); i++ {
		val := ""
		if k := valueOf.Field(i).Kind(); k == reflect.String {
			val = fmt.Sprintf("\"%s\"", valueOf.Field(i).Interface())
		} else if k == reflect.Int {
			val = fmt.Sprintf("%d", valueOf.Field(i).Interface())
		} else {
			val = fmt.Sprintf("%s", valueOf.Field(i).Interface())
		}
		_, _ = fmt.Fprintf(&stringBuilder, "%s:%s,", camelCase(valueOf.Type().Field(i).Name), val)
	}
	return strings.TrimSuffix(stringBuilder.String(), ",")
}

//This function creates variables map from args struct for graphql query
func Variables(input interface{}) map[string]interface{} {
	if input == nil {
		return map[string]interface{}{}
	}
	var requestInterface map[string]interface{}
	marshalledRequest, marshallErr := json.Marshal(input)
	if marshallErr != nil {
		panic(marshallErr)
	}
	unmarshallErr := json.Unmarshal(marshalledRequest, &requestInterface)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}
	return requestInterface
}

//This function returns camelcase name of provided field
func camelCase(field string) string {
	for i, v := range field {
		return string(unicode.ToLower(v)) + field[i+1:]
	}
	return ""
}