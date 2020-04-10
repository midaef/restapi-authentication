package server

func isAuthorized(responseToken string) (bool, string) {
	conn := NewConnection()
	defer conn.Close()
	var premium string
	var token string
	conn.QueryRow("select token,premium from users where token=$1", responseToken).Scan(&token, &premium)
	if token == responseToken {
		return true, premium
	}
	return false, premium
}
