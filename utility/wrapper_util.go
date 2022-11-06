package utils

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gstructs"
)

const (
	QueryTag  = "q"
	Eq        = "EQ"
	In        = "IN"
	Like      = "LIKE"
	LeftLike  = "LEFT_LIKE"
	RightLike = "RIGHT_LIKE"
)

func GetWrapper(param interface{}, wrapper *gdb.Model, columns map[string]string) (*gdb.Model, error) {
	r, err := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         param,
		RecursiveOption: 1,
	})
	if err != nil {
		return wrapper, err
	}
	for _, field := range r {
		if field.Value.IsZero() {
			continue
		}
		column := columns[field.Field.Name]
		if column == "" {
			continue
		}
		query := field.Tag(QueryTag)
		switch query {
		case Eq:
			wrapper = wrapper.Where(column, field.Value.Interface())
		case Like:
			wrapper = wrapper.WhereLike(column, "%"+field.Value.String()+"%")
		case LeftLike:
			wrapper = wrapper.WhereLike(column, "%"+field.Value.String())
		case RightLike:
			wrapper = wrapper.WhereLike(column, field.Value.String()+"%")
		case In:
			wrapper = wrapper.WhereIn(column, field.Value.Interface())
		default:
			wrapper = wrapper.Where(column, field.Value.Interface())
		case "-":
			break
		}
	}

	return wrapper, nil
}
