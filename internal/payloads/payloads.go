package payloads

func GetDefaultPayloads() []string {
	return []string{
		"null",
		"http://evil.com",
		"https://evil.com",
		"http://localhost",
		"http://127.0.0.1",
	}
}
