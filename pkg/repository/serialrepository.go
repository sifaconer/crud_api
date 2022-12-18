package repository

type SerialRepository interface {
	Create()
	Update()
	Delete()
	All()
}
