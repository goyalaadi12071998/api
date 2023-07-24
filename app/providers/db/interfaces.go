package db

type IRepo interface {
	Create(model any) error

	FindOne(model any, filter map[string]interface{}) error
}
