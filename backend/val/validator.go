package val

import "fmt"

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)

	if n < minLength || n > maxLength {
		return fmt.Errorf("value must be between %d and %d characters long, got %d", minLength, maxLength, n)
	}
	return nil
}

func ValidateEmail(email string) error {
	if len(email) < 5 || len(email) > 254 {
		return fmt.Errorf("email must be between 5 and 254 characters long, got %d", len(email))
	}

	atIndex := -1
	for i, char := range email {
		if char == '@' {
			if atIndex != -1 {
				return fmt.Errorf("email must contain only one '@' character, found multiple")
			}
			atIndex = i
		} else if char == '.' && atIndex == -1 {
			return fmt.Errorf("email must contain '@' before '.'")
		}
	}

	if atIndex == -1 || atIndex == 0 || atIndex == len(email)-1 {
		return fmt.Errorf("email must contain a valid local part and domain")
	}

	return nil
}

func ValidateName(name string) error {
	if err := ValidateString(name, 1, 30); err != nil {
		return err
	}

	for _, char := range name {
		if char < 'A' || char > 'Z' {
			if char < 'a' || char > 'z' {
				return fmt.Errorf("name must contain only alphabetic characters, found '%c'", char)
			}
		}
	}
	return nil
}

func ValidatePassword(password string) error {
	if err := ValidateString(password, 8, 100); err != nil {
		return err
	}
	return nil
}

func ValidateID(id int32) error {
	if id <= 0 {
		return fmt.Errorf("ID must be a positive integer, got %d", id)
	}
	return nil
}

func ValidateNote(note string) error {
	if err := ValidateString(note, 0, 5); err != nil {
		return err
	}
	return nil
}
