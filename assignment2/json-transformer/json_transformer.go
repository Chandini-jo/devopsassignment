package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sanitizeValue(value string) string {
	return strings.TrimSpace(value)
}

func transformString(value string) interface{} {
	sanitizedValue := sanitizeValue(value)
	rfc3339Pattern := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`)

	if rfc3339Pattern.MatchString(sanitizedValue) {
		t, err := time.Parse(time.RFC3339, sanitizedValue)
		if err == nil {
			return t.Unix()
		}
	}
	if sanitizedValue == "" {
		return nil
	}
	return sanitizedValue
}

func transformNumber(value string) interface{} {
	sanitizedValue := sanitizeValue(value)
	if sanitizedValue == "" {
		return nil
	}
	if num, err := strconv.ParseFloat(sanitizedValue, 64); err == nil {
		return num
	}
	return nil
}

func transformBoolean(value string) interface{} {
	sanitizedValue := sanitizeValue(value)
	if sanitizedValue == "" {
		return nil
	}
	if sanitizedValue == "1" || sanitizedValue == "t" || sanitizedValue == "T" ||
		sanitizedValue == "TRUE" || sanitizedValue == "true" || sanitizedValue == "True" {
		return true
	} else if sanitizedValue == "0" || sanitizedValue == "f" || sanitizedValue == "F" ||
		sanitizedValue == "FALSE" || sanitizedValue == "false" || sanitizedValue == "False" {
		return false
	}
	return nil
}

func transformNull(value string) interface{} {
	sanitizedValue := sanitizeValue(value)
	if sanitizedValue == "" {
		return nil
	}
	if sanitizedValue == "1" || sanitizedValue == "t" || sanitizedValue == "T" ||
		sanitizedValue == "TRUE" || sanitizedValue == "true" || sanitizedValue == "True" {
		return nil
	} else if sanitizedValue == "0" || sanitizedValue == "f" || sanitizedValue == "F" ||
		sanitizedValue == "FALSE" || sanitizedValue == "false" || sanitizedValue == "False" {
		return false
	}
	return nil
}

func transformList(value interface{}) interface{} {
	list, ok := value.([]interface{})
	if !ok {
		return nil
	}
	result := []interface{}{}
	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			for key, val := range m {
				switch key {
				case "S":
					transformedValue := transformString(val.(string))
					if transformedValue != nil {
						result = append(result, transformedValue)
					}
				case "N":
					transformedValue := transformNumber(val.(string))
					if transformedValue != nil {
						result = append(result, transformedValue)
					}
				case "BOOL":
					transformedValue := transformBoolean(val.(string))
					if transformedValue != nil {
						result = append(result, transformedValue)
					}
				case "NULL":
					transformedValue := transformNull(val.(string))
					if transformedValue != nil {
						result = append(result, transformedValue)
					}
				}
			}
		}
	}
	if len(result) > 0 {
		return result
	}
	return nil
}

func transformMap(value interface{}) interface{} {
	m, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	result := map[string]interface{}{}
	for key, val := range m {
		switch key {
		case "S":
			transformedValue := transformString(val.(string))
			if transformedValue != nil {
				result[key] = transformedValue
			}
		case "N":
			transformedValue := transformNumber(val.(string))
			if transformedValue != nil {
				result[key] = transformedValue
			}
		case "BOOL":
			transformedValue := transformBoolean(val.(string))
			if transformedValue != nil {
				result[key] = transformedValue
			}
		case "NULL":
			transformedValue := transformNull(val.(string))
			if transformedValue != nil {
				result[key] = transformedValue
			}
		case "L":
			transformedValue := transformList(val)
			if transformedValue != nil {
				result[key] = transformedValue
			}
		}
	}
	if len(result) > 0 {
		return result
	}
	return nil
}

func jsonTransformer(inputJSON map[string]interface{}) []interface{} {
	result := []interface{}{}
	for key, item := range inputJSON {
		sanitizedKey := sanitizeValue(key)
		if sanitizedKey != "" {
			if m, ok := item.(map[string]interface{}); ok {
				for dataType, val := range m {
					switch dataType {
					case "S":
						transformedValue := transformString(val.(string))
						if transformedValue != nil {
							result = append(result, map[string]interface{}{sanitizedKey: transformedValue})
						}
					case "N":
						transformedValue := transformNumber(val.(string))
						if transformedValue != nil {
							result = append(result, map[string]interface{}{sanitizedKey: transformedValue})
						}
					case "BOOL":
						transformedValue := transformBoolean(val.(string))
						if transformedValue != nil {
							result = append(result, map[string]interface{}{sanitizedKey: transformedValue})
						}
					case "NULL":
						transformedValue := transformNull(val.(string))
						if transformedValue != nil {
							result = append(result, map[string]interface{}{sanitizedKey: transformedValue})
						}
					case "L":
						transformedValue := transformList(val)
						if transformedValue != nil {
							result = append(result, map[string]interface{}{sanitizedKey: transformedValue})
						}
					case "M":
						transformedValue := transformMap(val)
						if transformedValue != nil {
							result = append(result, map[string]interface{}{sanitizedKey: transformedValue})
						}
					}
				}
			}
		}
	}
	return result
}

func main() {
	// Specify the file path
	filePath := "input.json"

	// Loading the contents of the JSON file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parsing the JSON content into a map
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(content, &jsonMap); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Transforming the JSON input
	transformedJSON := jsonTransformer(jsonMap)

	// Print the transformed JSON
	transformedJSON
