package utils

func ContainsKey(dict map[string]interface{}, key string) bool {
	if _, ok := dict[key]; ok {
		return true
	} else {
		return false
	}
}
