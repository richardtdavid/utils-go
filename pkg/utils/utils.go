package utils

func SafeMapNestKeysConv(input map[string]interface{}, key ...string) map[string]interface{} {
	for _, k := range key {
		input = SafeMapKeyConv(input, k)
	}
	return input
}

func SafeMapKeyConv(input map[string]interface{}, key string) map[string]interface{} {
	outInterface := input[key]
	if outInterface == nil {
		return nil
	}
	result, ok := outInterface.(map[string]interface{})
	if !ok {
		return nil
	}
	return result
}

func SafeListKeyConv(input map[string]interface{}, key string) []interface{} {
	outInterface := input[key]
	if outInterface == nil {
		return nil
	}
	result, ok := outInterface.([]interface{})
	if !ok {
		return nil
	}
	return result
}

func SafeStringKeyConv(input map[string]interface{}, key string) string {
	if inputVal, ok := input[key]; ok {
		if inputString, ok := inputVal.(string); ok {
			return inputString
		}
	}
	return ""
}

func SafeIntKeyConv(input map[string]interface{}, key string) int {
	if inputVal, ok := input[key]; ok {
		if inputFloat, ok := inputVal.(float64); ok {
			return int(inputFloat)
		}
	}
	return 0
}
