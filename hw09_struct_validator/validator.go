package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrNotAStruct         = errors.New("not struct is passed to validation")
	ErrNoRule             = errors.New("no such validation rule")
	ErrValidationLength   = errors.New("invalid: length")
	ErrValidationMinimum  = errors.New("invalid: minimum")
	ErrValidationMaximum  = errors.New("invalid: maximum")
	ErrValidationRegexp   = errors.New("invalid: regexp")
	ErrValidationContains = errors.New("invalid: not contains")
)

type RuleType int

const (
	Rule RuleType = iota
	Min
	Max
	Len
	In
	Regexp
)

type ValidationRule struct {
	Rule  RuleType
	Value interface{}
}

type ValidationRules []ValidationRule

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	if !isStruct(v) {
		return ErrNotAStruct
	}

	r := reflect.ValueOf(v)
	rtype := r.Type()

	var validationErrs ValidationErrors

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		fieldValue := r.Field(i)

		validationRules, err := getValidationRules(field)
		if err != nil {
			return err
		}
		if len(validationRules) == 0 {
			continue
		}

		vErr := validate(fieldValue, validationRules)
		if vErr != nil {
			validationErrs = append(
				validationErrs,
				ValidationError{Err: vErr, Field: field.Name},
			)
		}
	}

	return validationErrs
}

func getValidationRules(field reflect.StructField) (ValidationRules, error) {
	structTag := field.Tag
	rawValidationRules := structTag.Get("validate")

	if len(rawValidationRules) == 0 {
		return ValidationRules{}, nil
	}

	return makeRules(rawValidationRules)
}

func makeRules(rawData string) (ValidationRules, error) {
	result := ValidationRules{}
	rules := strings.Split(rawData, "|")
	for _, r := range rules {
		splitRuleAndValue := strings.Split(r, ":")
		if len(splitRuleAndValue) != 2 {
			continue
		}

		rule, err := getValidationRule(splitRuleAndValue[0])
		if err != nil {
			return ValidationRules{}, err
		}

		value, err := prepareValidationRule(rule, splitRuleAndValue[1])
		if err != nil {
			return ValidationRules{}, err
		}

		result = append(result, ValidationRule{
			Rule:  rule,
			Value: value,
		})
	}

	return result, nil
}

func validate(value reflect.Value, rules ValidationRules) error {
	switch value.Type().Kind() {
	case reflect.Int:
		return validateInt(value.Int(), rules)
	case reflect.String:
		return validateString(value.String(), rules)
	case reflect.Slice:
		return validateSlice(value, rules)
	default:
		return nil
	}
}

func getValidationRule(s string) (RuleType, error) {
	switch s {
	case "min":
		return Min, nil
	case "max":
		return Max, nil
	case "len":
		return Len, nil
	case "in":
		return In, nil
	case "regexp":
		return Regexp, nil
	}
	return 0, ErrNoRule
}

func prepareValidationRule(r RuleType, s string) (interface{}, error) {
	switch r {
	case Min:
		fallthrough
	case Max:
		fallthrough
	case Len:
		return strconv.Atoi(s)
	case In:
		return strings.Split(s, ","), nil
	case Regexp:
		return regexp.Compile(s)
	default:
		return nil, ErrNoRule
	}
}

func validateInt(v int64, rules ValidationRules) error {
	for _, r := range rules {
		switch r.Rule {
		case Min:
			if v < int64(r.Value.(int)) {
				return ErrValidationMinimum
			}
		case Max:
			if v > int64(r.Value.(int)) {
				return ErrValidationMaximum
			}
		case In:
			if !validateContains(v, r.Value) {
				return ErrValidationContains
			}
		default:
			continue
		}

	}
	return nil
}

func validateString(v string, rules ValidationRules) error {
	for _, r := range rules {
		switch r.Rule {
		case Len:
			if len(v) != r.Value.(int) {
				return ErrValidationLength
			}
		case In:
			if !validateContains(v, r.Value) {
				return ErrValidationContains
			}
		case Regexp:
			if !validateByRegexp(v, r.Value.(*regexp.Regexp)) {
				return ErrValidationRegexp
			}
		default:
			continue
		}
	}
	return nil
}

func validateSlice(v reflect.Value, r ValidationRules) error {
	for i := 0; i < v.Len(); i++ {
		err := validate(v.Index(i), r)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateContains(v interface{}, slice interface{}) bool {
	s := reflect.ValueOf(slice)
	strValue := fmt.Sprintf("%v", v)
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).String() == strValue {
			return true
		}
	}
	return false
}

func validateByRegexp(v string, regexp *regexp.Regexp) bool {
	return regexp.FindStringIndex(v) != nil
}

func isStruct(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Struct
}
