package functions

import "fmt"

func ShowFunctions() {
	var username string = buildUsername("Abbul","Rodriguez")
	var secretKey, accessKey= buildCredentials()
	var group, _= buildProfile()
	fmt.Println("username: ", username)
	fmt.Println("secretKey: ", secretKey)
	fmt.Println("accessKey: ", accessKey)
	fmt.Println("group: ", group)
}

func buildUsername(name string, lastname string) string  {
	return name + "." + lastname
}

func buildCredentials() (secretKey string, accessKey string)  {
	return "ImSecretKey","ImAccessKey"
}

func buildProfile() (group string, role string)  {
	return "admin-group","admin-role"
}