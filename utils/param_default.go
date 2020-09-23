package utils

func SetDefault(data map[string]interface{}, key string, defaultValue interface{}) {
	_, ok := data[key]
	if !ok {
		data[key] = defaultValue
	}
}
