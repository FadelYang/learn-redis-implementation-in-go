package tools

var (
	VALIDATION_ERROR = "VALIDATION_ERROR"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Code    string              `json:"code"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		Code:    "VALIDATION_ERROR",
		Message: "Invalid input",
		Errors:  make(map[string][]string),
	}
}

func (v *ValidationError) Add(field, msg string) {
	v.Errors[field] = append(v.Errors[field], msg)
}

func (v *ValidationError) Error() string {
	return v.Code
}
