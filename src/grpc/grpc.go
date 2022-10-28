package grpc

import (
	"editory/editory_sms_service/config"
	"editory/editory_sms_service/genproto/sms_service"
	"editory/editory_sms_service/grpc/service"
	"editory/editory_sms_service/pkg/logger"
	"editory/editory_sms_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	sms_service.RegisterSmsServiceServer(grpcServer, service.NewSendService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
