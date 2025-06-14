package val

import (
	"fmt"

	"github.com/DebdipWritesCode/MUN_Scoresheet/backend/pb"
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)

	if n < minLength || n > maxLength {
		return fmt.Errorf("value must be between %d and %d characters long, got %d", minLength, maxLength, n)
	}
	return nil
}

func ValidateScore(score float64) error {
	if score < -100 || score > 100 {
		return fmt.Errorf("score must be between -100 and 100, got %f", score)
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
				if char != ' ' {
					return fmt.Errorf("name must contain only alphabetic characters, found '%c'", char)
				}
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

func ValidateRule(rule string, rules []string) error {
	for _, r := range rules {
		if r == rule {
			return nil
		}
	}

	return fmt.Errorf("rule '%s' is not valid, must be one of %v", rule, rules)
}

func ValidateWeight(fieldName string, value *float64) error {
	if value == nil {
		return nil
	}
	if *value < 0 || *value > 1 {
		return fmt.Errorf("%s must be between 0 and 1", fieldName)
	}
	return nil
}

func ValidateSpecialConditionRule(ruleType string, isSpecialParameter *bool) error {
	if ruleType == "special" {
		if isSpecialParameter == nil || !*isSpecialParameter {
			return fmt.Errorf("is_special_parameter must be true when rule_type is 'special'")
		}
	} else {
		if isSpecialParameter != nil && *isSpecialParameter {
			return fmt.Errorf("is_special_parameter must be false or unset unless rule_type is 'special'")
		}
	}
	return nil
}

func ValidateDelegateInput(delegate *pb.DelegateInput) error {
	if err := ValidateName(delegate.GetDelegateName()); err != nil {
		return fmt.Errorf("invalid delegate name: %w", err)
	}
	for _, parameter := range delegate.GetParameters() {
		if err := ValidateParameterInput(parameter); err != nil {
			return fmt.Errorf("invalid parameter input: %w", err)
		}
	}
	return nil
}

func ValidateParameterInput(parameter *pb.ParameterInput) error {
	if err := ValidateName(parameter.GetParameterName()); err != nil {
		return fmt.Errorf("invalid parameter name: %w", err)
	}

	if err := ValidateScore(parameter.GetReceived()); err != nil {
		return fmt.Errorf("invalid received score: %w", err)
	}

	if err := ValidateScore(parameter.GetHighest()); err != nil {
		return fmt.Errorf("invalid highest score: %w", err)
	}

	return nil
}
