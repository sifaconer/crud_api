package repository

type FabricanteRepository interface {
	Create()
	Update()
	Delete()
	All()
}
