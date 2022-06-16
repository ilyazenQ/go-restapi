package actions

// TODO Сходить в дб
func GetAllStatus() map[int]string {
	return map[int]string{
		1: "NEW",
		2: "SUCCESS",
		3: "FAIL",
		4: "ERROR",
		5: "CLOSED",
	}
}
