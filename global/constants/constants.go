package constants

import "strings"

const (
	LOGIN_USER_TOKEN    = "login{id}"
	LOGIN_CODE_GENERATE = "GeneratingCode{email}"
)

func LoginCodeGenerate(email string) string {
	return strings.Replace(LOGIN_CODE_GENERATE, "{email}", email, -1)
}
func LoginTarget(target string) string {
	return strings.Replace(LOGIN_USER_TOKEN, "{id}", target, -1)
}
