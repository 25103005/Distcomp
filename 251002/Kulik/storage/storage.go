package storage

type IStorage[T any] interface {
	Create(entity T) (int64, error)
	Get(id int64) (T, error)
	GetAll() ([]T, []int64, error)
	Update(id int64, entity T) error
	Delete(id int64) error
}
