package helpers

//func for validate if data is empty
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}