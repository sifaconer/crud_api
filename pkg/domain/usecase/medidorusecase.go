package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sifaconer/crud_api/pkg/database/redisdb"
	"github.com/sifaconer/crud_api/pkg/domain/entity"
	"github.com/sifaconer/crud_api/pkg/repository"
)

type MedidorUseCase interface {
	Create(model entity.Medidor) (entity.Medidor, error)
	Update(model entity.Medidor) (entity.Medidor, error)
	Delete(id uuid.UUID) error
	All() ([]entity.Medidor, error)
	ByID(id uuid.UUID) (entity.Medidor, error)
	RecentInstallation(serial string) (entity.Medidor, error)
	Inactive() ([]entity.Medidor, error)
}

//implements MedidorUseCase
type medidorUseCaseImpl struct {
	Repo   repository.MedidorRepository
	Stream *redisdb.StreamRedis
}

func NewMedidorUseCaseImpl(repo repository.MedidorRepository,
	stream *redisdb.StreamRedis) *medidorUseCaseImpl {
	return &medidorUseCaseImpl{
		Repo:   repo,
		Stream: stream,
	}
}

func (m *medidorUseCaseImpl) Create(model entity.Medidor) (entity.Medidor, error) {
	resp, err := m.Repo.RecentInstallation(model.Serial)
	if err != nil {
		return resp, err
	}

	if resp != (entity.Medidor{}) &&
		(resp.Serial+"-"+resp.Brand) == (model.Serial+"-"+model.Brand) {
		return resp, errors.New("duplicate serial-marca")
	}

	respArr, err := m.Repo.All()
	if err != nil {
		return resp, err
	}

	for _, v := range respArr {
		if v.Address == model.Address &&
			v.InstallationDate == model.InstallationDate {
			return resp, errors.New("there can only be one medidor in this predio")
		}
	}

	resp, err = m.Repo.Create(model)
	if err != nil {
		return resp, err
	}
	// stream redis
	m.Stream.PublishEvent(resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *medidorUseCaseImpl) Update(model entity.Medidor) (entity.Medidor, error) {
	resp, err := m.Repo.Update(model)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (m *medidorUseCaseImpl) Delete(id uuid.UUID) error {
	medidor, err := m.Repo.ByID(id)
	if err != nil {
		return err
	}

	if medidor.RetirementDate == nil {
		return errors.New("cannot be removed, currently in use")
	}

	err = m.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (m *medidorUseCaseImpl) All() ([]entity.Medidor, error) {
	resp, err := m.Repo.All()
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (m *medidorUseCaseImpl) ByID(id uuid.UUID) (entity.Medidor, error) {
	resp, err := m.Repo.ByID(id)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (m *medidorUseCaseImpl) RecentInstallation(serial string) (entity.Medidor, error) {
	resp, err := m.Repo.RecentInstallation(serial)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (m *medidorUseCaseImpl) Inactive() ([]entity.Medidor, error) {
	resp, err := m.Repo.Inactive()
	if err != nil {
		return resp, err
	}
	return resp, nil
}
