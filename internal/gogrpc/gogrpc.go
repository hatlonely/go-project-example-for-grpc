package gogrpc

import (
	"context"
	apigogrpc "github.com/hatlonely/go-project-example-for-grpc/api/gogrpc"
	"github.com/sirupsen/logrus"
)

var InfoLog *logrus.Logger
var WarnLog *logrus.Logger
var AccessLog *logrus.Logger

func init() {
	InfoLog = logrus.New()
	WarnLog = logrus.New()
	AccessLog = logrus.New()
}

func NewService() *Service {
	return &Service{}
}

type Service struct{}

func (s *Service) Do(ctx context.Context, request *apigogrpc.Request) (*apigogrpc.Response, error) {
	response := &apigogrpc.Response{
		Message: request.Message,
	}

	AccessLog.WithFields(logrus.Fields{
		"request":  request,
		"response": response,
	}).Info()

	return response, nil
}
