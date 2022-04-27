//go:generate mockgen -destination=../../../mocks/adapter/validatoradapter/adapter.go github.com/adityarev/go-be-starter-2/pkg/core/adapter/validatoradapter Adapter

package validatoradapter

type Adapter interface {
	Struct(s interface{}) error
	Var(field Field) error
}

type Field struct {
	Name  string
	Value interface{}
	Tag   tag
}
