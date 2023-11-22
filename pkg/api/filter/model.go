package filter

import "fmt"

const (
	DataTypeStr  = "string"
	DataTypeInt  = "int"
	DataTypeDate = "date"
	DataTypeBool = "bool"

	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "lte"
	OperatorGreaterThan   = "gt"
	OperatorGreaterThanEq = "gte"
	OperatorBetween       = "between"
	OperatorLike          = "like"
)

type options struct {
	limit  int
	fields []Field
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

type Options interface {
	GetLimit() int
	AddField(name, operator, value, dtype string) error
	Fields() []Field
}

func NewOptions(limit int) Options {
	return &options{limit: limit}
}

func (o *options) GetLimit() int {
	return o.limit
}

func (o *options) AddField(name, operator, value, dtype string) error {
	if err := validateOperator(operator); err != nil {
		return err
	}
	o.fields = append(o.fields, Field{
		Name:     name,
		Operator: operator,
		Value:    value,
		Type:     dtype,
	})

	return nil
}

func (o *options) Fields() []Field {
	return o.fields
}

func validateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	case OperatorNotEq:
	case OperatorLowerThan:
	case OperatorLowerThanEq:
	case OperatorGreaterThan:
	case OperatorGreaterThanEq:
	default:
		return fmt.Errorf("bad operator")
	}

	return nil
}
