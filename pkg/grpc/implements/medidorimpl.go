package implements

import (
	"context"

	"github.com/google/uuid"
	"github.com/sifaconer/crud_api/pkg/domain/entity"
	"github.com/sifaconer/crud_api/pkg/domain/usecase"
	"github.com/sifaconer/crud_api/pkg/grpc/proto"
	"github.com/sifaconer/crud_api/utils"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type MedidorImpl struct {
	proto.UnimplementedMedidorServicesServer
	Usecase usecase.MedidorUseCase
}

func (s *MedidorImpl) Create(ctx context.Context, model *proto.Medidor) (*proto.ResponseMedidor, error) {

	result, err := s.Usecase.Create(entity.Medidor{
		Brand:            model.Brand,
		Address:          model.Address,
		InstallationDate: utils.FormatFromUNIX(model.UnixInstallationDate),
		RetirementDate:   utils.FormatPointerFromUNIX(model.UnixRetirementDate),
		Serial:           model.Serial,
		Lines:            model.Lines,
		IsActive:         model.IsActive,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	response := &proto.ResponseMedidor{
		Id:                   result.ID.String(),
		Brand:                result.Brand,
		Address:              result.Address,
		UnixInstallationDate: utils.FormatToUNIX(result.InstallationDate),
		UnixRetirementDate:   utils.FormatPointerToUNIX(result.RetirementDate),
		Serial:               result.Serial,
		Lines:                result.Lines,
		IsActive:             result.IsActive,
		UnixCreatedAt:        utils.FormatToUNIX(result.CreatedAt),
	}
	return response, nil
}

func (s *MedidorImpl) Update(ctx context.Context, model *proto.Medidor) (*proto.ResponseMedidor, error) {
	id, err := uuid.Parse(model.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	result, err := s.Usecase.Update(entity.Medidor{
		ID:               id,
		Brand:            model.Brand,
		Address:          model.Address,
		InstallationDate: utils.FormatFromUNIX(model.UnixInstallationDate),
		RetirementDate:   utils.FormatPointerFromUNIX(model.UnixRetirementDate),
		Serial:           model.Serial,
		Lines:            model.Lines,
		IsActive:         model.IsActive,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	response := &proto.ResponseMedidor{
		Id:                   result.ID.String(),
		Brand:                result.Brand,
		Address:              result.Address,
		UnixInstallationDate: utils.FormatToUNIX(result.InstallationDate),
		UnixRetirementDate:   utils.FormatPointerToUNIX(result.RetirementDate),
		Serial:               result.Serial,
		Lines:                result.Lines,
		IsActive:             result.IsActive,
		UnixCreatedAt:        utils.FormatToUNIX(result.CreatedAt),
	}
	return response, nil
}

func (s *MedidorImpl) Delete(ctx context.Context, id *proto.UUID) (*proto.Empty, error) {
	uid, err := uuid.Parse(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = s.Usecase.Delete(uid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	response := &proto.Empty{}
	return response, nil
}

func (s *MedidorImpl) All(ctx context.Context, model *proto.Empty) (*proto.ArrayMedidor, error) {
	result, err := s.Usecase.All()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	response := &proto.ArrayMedidor{}

	for _, v := range result {
		var medidor proto.ResponseMedidor

		medidor.Id = v.ID.String()
		medidor.Brand = v.Brand
		medidor.Address = v.Address
		medidor.UnixInstallationDate = utils.FormatToUNIX(v.InstallationDate)
		medidor.UnixRetirementDate = utils.FormatPointerToUNIX(v.RetirementDate)
		medidor.Serial = v.Serial
		medidor.Lines = v.Lines
		medidor.IsActive = v.IsActive
		medidor.UnixCreatedAt = utils.FormatToUNIX(v.CreatedAt)

		response.Medidor = append(response.Medidor, &medidor)
	}

	return response, nil
}

func (s *MedidorImpl) ByID(ctx context.Context, id *proto.UUID) (*proto.ResponseMedidor, error) {
	uid, err := uuid.Parse(id.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	result, err := s.Usecase.ByID(uid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	response := &proto.ResponseMedidor{
		Id:                   result.ID.String(),
		Brand:                result.Brand,
		Address:              result.Address,
		UnixInstallationDate: utils.FormatToUNIX(result.InstallationDate),
		UnixRetirementDate:   utils.FormatPointerToUNIX(result.RetirementDate),
		Serial:               result.Serial,
		Lines:                result.Lines,
		IsActive:             result.IsActive,
		UnixCreatedAt:        utils.FormatToUNIX(result.CreatedAt),
	}
	return response, nil
}

func (s *MedidorImpl) RecentInstallation(ctx context.Context, serial *proto.Serial) (*proto.ResponseMedidor, error) {

	result, err := s.Usecase.RecentInstallation(serial.Serial)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	response := &proto.ResponseMedidor{
		Id:                   result.ID.String(),
		Brand:                result.Brand,
		Address:              result.Address,
		UnixInstallationDate: utils.FormatToUNIX(result.InstallationDate),
		UnixRetirementDate:   utils.FormatPointerToUNIX(result.RetirementDate),
		Serial:               result.Serial,
		Lines:                result.Lines,
		IsActive:             result.IsActive,
		UnixCreatedAt:        utils.FormatToUNIX(result.CreatedAt),
	}
	return response, nil
}

func (s *MedidorImpl) Inactive(ctx context.Context, model *proto.Empty) (*proto.ArrayMedidor, error) {

	result, err := s.Usecase.Inactive()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	response := &proto.ArrayMedidor{}

	for _, v := range result {
		var medidor proto.ResponseMedidor

		medidor.Id = v.ID.String()
		medidor.Brand = v.Brand
		medidor.Address = v.Address
		medidor.UnixInstallationDate = utils.FormatToUNIX(v.InstallationDate)
		medidor.UnixRetirementDate = utils.FormatPointerToUNIX(v.RetirementDate)
		medidor.Serial = v.Serial
		medidor.Lines = v.Lines
		medidor.IsActive = v.IsActive
		medidor.UnixCreatedAt = utils.FormatToUNIX(v.CreatedAt)

		response.Medidor = append(response.Medidor, &medidor)
	}

	return response, nil
}
