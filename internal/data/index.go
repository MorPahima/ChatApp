package data

type MongoRepository interface {
	Set(body interface{}) error
	GetAll() ([]interface{}, error)
}