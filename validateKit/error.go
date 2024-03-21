package validateKit

/*func TransErrors(v *validator.Validate, trans locales.Translator, err error) string {
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return (err.(validator.ValidationErrors)).Translate(v, trans)
		}
	}
	return ""
}*/

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
