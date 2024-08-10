package main

import "fmt"

// User represents a user with an email and password.
type User struct { // Use CamelCase for type names
	Email    string
	Password string
}

// UserRepo handles operations related to user data storage and retrieval.
type UserRepo struct { // Use CamelCase for type names
	DB []User // Use CamelCase for type names
}

// Register adds a new user to the repository if the user data is valid.
func (r *UserRepo) Register(u User) error { // Use pointer receiver for methods that modify the receiver
	if u.Email == "" || u.Password == "" {
		return fmt.Errorf("register failed: email or password is empty") // Return error instead of printing
	}

	r.DB = append(r.DB, u)
	return nil
}

// Login checks if the provided user credentials match any user in the repository.
// Returns an authentication token if successful.
func (r *UserRepo) Login(u User) (string, error) { // Return error to handle failure cases
	if u.Email == "" || u.Password == "" {
		return "", fmt.Errorf("login failed: email or password is empty") // Return error instead of printing
	}

	for _, us := range r.DB {
		if us.Email == u.Email && us.Password == u.Password {
			return "auth token", nil // Return token with nil error on success
		}
	}

	return "", fmt.Errorf("login failed: invalid credentials") // Return error for invalid credentials
}

func main() {
	repo := UserRepo{}
	user := User{Email: "example@example.com", Password: "password123"}

	if err := repo.Register(user); err != nil {
		fmt.Println(err)
	}

	token, err := repo.Login(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Login successful, token:", token)
	}
}
