package to

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"reflect"
	"strings"
	"unicode"
)

//This function creates a graphql query from provided query name, operation, input, arguments and the type of the response.
func Query(name string, operation models.GraphqlOperation, input interface{}, args interface{}, response interface{}) string {
	if args != nil {
		return fmt.Sprintf("%s{response:%s(%s){%s}}",
			operation,
			name,
			Arguments(args),
			GraphqlResponseSelector(response))
	} else if input != nil {
		inputNameSlice := strings.Split(reflect.TypeOf(input).String(), ".")
		nameInput := inputNameSlice[len(inputNameSlice)-1]
		return fmt.Sprintf("%s($input:%s){response:%s(input:$input){%s}}",
			operation,
			nameInput,
			name,
			GraphqlResponseSelector(response))
	} else {
		return fmt.Sprintf("%s{response:%s{%s}}",
			operation,
			name,
			GraphqlResponseSelector(response))
	}
}

//This function creates arguments string from args struct for graphql query
func Arguments(args interface{}) string {
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
		_, _ = fmt.Fprintf(&stringBuilder, "%s:%s,", getFieldName(valueOf.Type().Field(i).Name), val)
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

//This function creates graphql responses from provided struct as value
func GraphqlResponseSelector(value interface{}) string {
	value = populateData(value, 0)
	var stringBuilder strings.Builder
	valueOf := reflect.ValueOf(value)
	if valueOf.Kind() == reflect.Slice {
		graphqlResponseSelector(reflect.ValueOf(toArrayInterface(value)[0]), &stringBuilder)
	} else {
		graphqlResponseSelector(reflect.ValueOf(value), &stringBuilder)
	}
	return stringBuilder.String()
}

//This function creates graphql responses from provided struct as value
func graphqlResponseSelector(valueOf reflect.Value, stringBuilder *strings.Builder) {
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		name := valueOf.Type().Field(i).Name
		if field.Kind() == reflect.Struct {
			_, _ = fmt.Fprintln(stringBuilder, fmt.Sprintf("%s {", getFieldName(name)))
			graphqlResponseSelector(field, stringBuilder)
			_, _ = fmt.Fprintln(stringBuilder, "}")
			continue
		} else if field.Kind() == reflect.Slice {
			var sliceField = toArrayInterface(field.Interface())
			if len(sliceField) == 0 {
				continue
			}
			if reflect.TypeOf(sliceField[0]).Kind() == reflect.String {
				_, _ = fmt.Fprint(stringBuilder, getFieldName(name))
				continue
			}
			_, _ = fmt.Fprint(stringBuilder, fmt.Sprintf("%s {", getFieldName(name)))
			_, _ = fmt.Fprintln(stringBuilder, "")
			graphqlResponseSelector(reflect.ValueOf(sliceField[0]), stringBuilder)
			_, _ = fmt.Fprintln(stringBuilder, "}")
			continue
		}
		_, _ = fmt.Fprint(stringBuilder, getFieldName(name))
		_, _ = fmt.Fprintln(stringBuilder, "")
	}
}
//This function populates fake data to provided struct as value
func populateData(value interface{}, counter int) interface{} {
	if counter == 10 {
		panic("Couldn't populate response selector.")
	}
	firstValue := value
	err := faker.FakeData(&value)
	if err != nil {
		fmt.Println(err)
	}
	if reflect.DeepEqual(firstValue, value) {
		return populateData(value, counter+1)
	}

	return value
}

//This function returns camelcase name of provided field
func getFieldName(field string) string {
	for i, v := range field {
		return string(unicode.ToLower(v)) + field[i+1:]
	}
	return ""
}

//This function returns and converts to array interface from provided input interface
func toArrayInterface(input interface{}) []interface{} {
	object := reflect.ValueOf(input)
	var items []interface{}
	for i := 0; i < object.Len(); i++ {
		items = append(items, object.Index(i).Interface())
	}
	return items
}
