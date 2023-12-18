package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ErrorResponse struct {
	Error          bool
	FailedField    string
	Tag            string
	DefaultMessage string
	CustomMessage  string
	Value          interface{}
}

type ValidationMessage struct {
	Error   bool
	Message string
	Errors  map[string]string
}

func (v *ValidationMessage) NewValidationErrorResponse() *ValidationErrorResponse {
	return &ValidationErrorResponse{
		Success: false,
		Error:   v.Message,
		Errors:  v.Errors,
	}
}

type ValidationErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Failed validation"`
	Errors  map[string]string
}

type Validator struct {
	validator *validator.Validate
}

// New validator for model fields.
func NewValidator() *Validator {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Rename struct fields.
	validate.RegisterTagNameFunc(func(fl reflect.StructField) string {
		name := strings.SplitN(fl.Tag.Get("json"), ",", 2)
		return name[0]
	})

	validator := &Validator{
		validator: validate,
	}

	return validator
}

// Validate struct
func (v *Validator) Validate(data interface{}) []ErrorResponse {
	errs := v.validator.Struct(data)
	if errs != nil {
		return validationErrorResponse(errs)
	}

	return nil
}

// ValidatorErrors func for show validation errors for each invalid fields.
func (v *Validator) errors(err []ErrorResponse) ValidationMessage {
	msg := ValidationMessage{
		Error:   true,
		Message: "Failed validation",
	}

	msg.Errors = make(map[string]string)

	// Make error message for each invalid field.
	for _, err := range err {
		msg.Errors[err.FailedField] = err.CustomMessage
	}

	return msg
}

// Return http Json response
func (v *Validator) JsonResponse(ctx *fiber.Ctx, err []ErrorResponse) error {
	validation := v.errors(err)
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(validation.NewValidationErrorResponse())
}

// Build validation error message
func validationErrorResponse(errs error) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	for _, err := range errs.(validator.ValidationErrors) {
		// In this case data object is actually holding the User struct
		var elem ErrorResponse
		// Get name of the field's struct.
		structName := strings.Split(err.Namespace(), ".")[0]

		elem.FailedField = err.Field() // Export struct field name
		elem.Tag = err.Tag()           // Export struct tag
		elem.Value = err.Value()       // Export field value
		elem.DefaultMessage = err.Error()
		elem.CustomMessage = fmt.Sprintf(
			"%s failed '%s' tag check (value '%s' is not valid for %s struct)",
			cases.Title(language.English).String(strings.Replace(err.Field(), "_", " ", -1)),
			err.Tag(),
			err.Value(),
			structName,
		)
		elem.Error = true

		validationErrors = append(validationErrors, elem)
	}
	return validationErrors
}
