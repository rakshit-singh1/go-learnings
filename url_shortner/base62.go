package main

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func encodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	result := ""
	base := int64(len(base62Chars))

	for num > 0 {
		result = string(base62Chars[num%base]) + result
		num = num / base
	}

	return result
}
