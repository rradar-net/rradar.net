package proto

func OptionalString(value *string) string {
	if value != nil {
		return *value
	}

	return ""
}
