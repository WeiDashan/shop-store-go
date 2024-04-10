package constants

const (
	APPUSER = "/app-user"
)

var resultMap map[string]bool

func InitFilter() {
	resultMap = make(map[string]bool)
	resultMap[APPUSER+"/test"] = true
}
func NeedFilter(fullPath string) bool {
	return resultMap[fullPath]
}
