package database

import "context"

func GetUserByUsername(username string) (*string, error) {
	var passwordHash string
	err := DB.QueryRow(context.Background(),
		"SELECT password_hash FROM users WHERE username = $1", username).
		Scan(&passwordHash)

	if err != nil {
		return nil, err
	}

	return &passwordHash, nil
}
