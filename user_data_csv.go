package main

import (
	"encoding/csv"
	"os"
)

var UserMemCache = map[string]User{}

func initUserFromCSV() error {
	f, err := os.Open(userFilepath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		if len(record) != 3 {
			continue
		}

		user := User{
			Username: record[0],
			Email:    record[1],
			Password: record[2],
		}

		_, ok := UserMemCache[user.Username]
		if !ok {
			UserMemCache[user.Username] = user
		}
	}

	return nil
}

func saveUserDataToCSV(users map[string]User, filepath string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, user := range users {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return err
		}

		header := []string{"Name", "Email", "Password"}
		if err := w.Write(header); err != nil {
			return err
		}

		if err := w.Write([]string{user.Username, user.Email, hashedPassword}); err != nil {
			return err
		}
	}

	return nil
}
