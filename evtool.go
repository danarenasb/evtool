package evtool

import (
	"net"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsEmailValid takes the email as a string and checks for 3 thingsL 1)length (min 3 max 254), 2) that it macthes REGEX "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$") and 3) uses net package to check DNX MX records for the domain.
func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

// NormalizeGmail removes all periods and anything after the + sign on gmail addresses, as these are considered ways to create aliases for one gmail address
func NormalizeGmail(e string) string {
	emailWithoutDots := strings.ReplaceAll(e, ".", "")
	if strings.Contains(emailWithoutDots, "+") {
		normalizedEmail := strings.Split(emailWithoutDots, "+")
		return normalizedEmail[0]
	}
	return emailWithoutDots
}
