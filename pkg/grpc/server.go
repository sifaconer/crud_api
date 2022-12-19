package grpc

import (
	"net"

	"github.com/sifaconer/crud_api/pkg/domain/usecase"
	"github.com/sifaconer/crud_api/pkg/grpc/implements"
	"github.com/sifaconer/crud_api/pkg/grpc/proto"
	"google.golang.org/grpc"
)

type serverGRPC struct {
	Port           string
	MedidorUseCase usecase.MedidorUseCase
}

func NewServerGRPC(port string,
	useCase usecase.MedidorUseCase) *serverGRPC {
	return &serverGRPC{
		Port:           port,
		MedidorUseCase: useCase,
	}
}

func (s *serverGRPC) RunServer() error {
	lis, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		return err
	}

	register := grpc.NewServer()

	service := &implements.MedidorImpl{
		Usecase: s.MedidorUseCase,
	}
	proto.RegisterMedidorServicesServer(register, service)

	return register.Serve(lis)
}
