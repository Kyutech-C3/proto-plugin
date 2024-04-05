package utils

func TypeConvert(goType string) string {
	switch goType {
	case "string":
		return "string"
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "byte", "rune":
		return "number"
	case "float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	default:
		return "any"
	}
}
