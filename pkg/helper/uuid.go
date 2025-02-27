package helper

import "regexp"

func IsValidUUID(u string) bool {
	re := regexp.MustCompile(`^[a-fA-F0-9]{8}\-[a-fA-F0-9]{4}\-4[a-fA-F0-9]{3}\-[89abAB][a-fA-F0-9]{3}\-[a-fA-F0-9]{12}$`)
	return re.MatchString(u)
}
