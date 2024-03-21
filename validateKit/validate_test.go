package validateKit

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestOneOf(t *testing.T) {
	type T struct {
		Name   string `validate:"required"`
		Gender string `validate:"oneof=Male Female"`
	}
	simplifyError(Struct(&T{Gender: "Male"}))
	simplifyError(Struct(&T{Gender: "Male1"}))
}

func simplifyError(err error) {
	fmt.Println("-------------------- start --------------------")
	defer func() {
		fmt.Println("--------------------  end  --------------------")
	}()
	if err == nil {
		return
	}
	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		fmt.Printf("InvalidValidationError: %v\n", err)
		return
	}
	var v validator.ValidationErrors
	if errors.As(err, &v) {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("FieldError: %v\n", err)
			fmt.Printf("FieldError(Translation): %s\n", transError(err))
		}
		return
	}
	fmt.Printf("Error: %v\n", err)
}

func transError(err validator.FieldError) string {
	return err.Translate(ZhTrans())
}
