package utils

func StrPtrToStr(pointer *string) string {
	if pointer == nil {
		return ""
	}
	return *pointer
}
func Int64PtrToInt(pointer *int64) int64 {
	if pointer == nil {
		return 0
	}
	return *pointer
}
