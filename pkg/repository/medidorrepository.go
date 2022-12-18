package repository

type MedidorRepository interface {
	Create()
	Update()
	Delete()
	All()       // todos o uno
	Installed() // condicionar para energía true o false
	ByID()
}
