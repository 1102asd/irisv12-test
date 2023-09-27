package controllers

func ConvertStringMapToStringInterfaceMap(inputMap map[string]string) map[string]interface{} {
	resultMap := make(map[string]interface{})

	for key, value := range inputMap {
		resultMap[key] = value
	}

	return resultMap
}
