package common

func StringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func PointerToString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
