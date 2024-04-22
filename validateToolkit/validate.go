package validateToolkit

import (
	"context"
)

func Struct(s interface{}) error {
	return validate.Struct(s)
}

func StructCtx(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}

func ValidateMap(data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return validate.ValidateMap(data, rules)
}
func ValidateMapCtx(ctx context.Context, data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return validate.ValidateMapCtx(ctx, data, rules)
}

func ValidateVar(field interface{}, tag string) error {
	return validate.Var(field, tag)
}

func ValidateVarCtx(ctx context.Context, field interface{}, tag string) error {
	return validate.VarCtx(ctx, field, tag)
}

func ValidateVarWithValue(field interface{}, other interface{}, tag string) error {
	return validate.VarWithValue(field, other, tag)
}

func ValidateVarWithValueCtx(ctx context.Context, field interface{}, other interface{}, tag string) error {
	return validate.VarWithValueCtx(ctx, field, other, tag)
}
