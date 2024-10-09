package utils

import (
    "regexp"
)

func IsValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9-]+\.com$`)
    return re.MatchString(email)
}
