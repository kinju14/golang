package main

func main() {
	age := 18

	if age < 18 {
		println("You are a minor.")
	} else if age >= 18 && age < 65 {
		println("You are an adult.")
	} else {
		println("You are a senior citizen.")
	}

	role := "admin"
	hasPermission := false

	if role == "admin" && hasPermission {
		println("You have access to the admin panel.")
	} else {
		println("You do not have access to the admin panel.")
	}
}
