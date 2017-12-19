package json

import (
	"encoding/json"
    "io"
    "log"
	"strings"
	"reflect"
)


func parse(jsonStream string) (interface{}, string, error) {

    dec := json.NewDecoder(strings.NewReader(jsonStream))

	var result interface{}
	var str string
	err := dec.Decode(&result)
	if err == io.EOF {
		log.Fatal(err)
	} else if err != nil {
		log.Fatal(err)
	}
	switch t := result.(type) {
	case string:
		str = t
	default:
		result = t
	}
	return result, str, err
}

func parseOne(jsonStream string) (interface{}, error) {
	
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	var result interface{}
	err := dec.Decode(&result)
	if err == io.EOF {
		log.Fatal(err)
	} else if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func parseObject(jsonStream string, jsonKey string) (interface{}, error) {
	
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	var result interface{}
	err := dec.Decode(&result)
	if err == io.EOF {
		log.Fatal(err)
	} else if err != nil {
		log.Fatal(err)
	}

	v := reflect.ValueOf(result)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			if (key.Interface().(string) == jsonKey) {
				result = strct.Interface()
			}
		}
	}
	
	return result, err
}

