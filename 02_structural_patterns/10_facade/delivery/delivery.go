package delivery

func SendExpress(receiver, item string) bool {
	if receiver == "Void" {
		return false
	}
	return true
}
