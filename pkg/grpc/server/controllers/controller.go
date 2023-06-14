package controllers

import (
	"context"

	gen "github.com/azar-intelops/grpc-for-evans/gen/api/v1"
	"github.com/azar-intelops/grpc-for-evans/pkg/grpc/server/models"
	"github.com/azar-intelops/grpc-for-evans/pkg/grpc/server/services"
	"github.com/google/uuid"
)

var officeService = services.OfficeService{}

type OfficeServer struct {
	gen.UnimplementedOfficeServiceServer
}

func NewOfficeServer() *OfficeServer {
	return &OfficeServer{}
}

// ping
func (s *OfficeServer) Ping(ctx context.Context, req *gen.PingRequest) (*gen.PingResponse, error) {
	return &gen.PingResponse{
		Message: "Server is healthy and Working!",
	}, nil
}

// create
func (s *OfficeServer) CreateOffice(ctx context.Context, req *gen.CreateRequest) (*gen.CreateResponse, error) {
	id := uuid.New().String()
	request := req.GetOffice()

	office := models.Office{
		Id:             id,
		OfficeName:     request.GetOfficeName(),
		Address:        request.GetAddress(),
		EmployeesCount: request.GetEmployeesCount(),
		TurnOver:       request.GetTurnOver(),
	}

	if err := officeService.CreateOffice(office); err != nil {
		return nil, err
	}

	return &gen.CreateResponse{
		Message: "Office Created Successfully!",
	}, nil
}

// Read office by id
func (s *OfficeServer) GetOffice(ctx context.Context, req *gen.GetRequest) (*gen.GetResponse, error) {
	id := req.GetId()

	office, err := officeService.GetOffice(id)
	if err != nil {
		return nil, err
	}

	return &gen.GetResponse{
		Office: &gen.Office{
			Id:             office.Id,
			OfficeName:     office.OfficeName,
			Address:        office.Address,
			EmployeesCount: office.EmployeesCount,
			TurnOver:       office.TurnOver,
		},
	}, nil
}

// Delete by ID
func (s *OfficeServer) DeleteOffice(ctx context.Context, req *gen.DeleteRequest) (*gen.DeleteResponse, error) {
	id := req.Id

	err := officeService.DeleteOffice(id)

	if err != nil {
		return nil, err
	}

	return &gen.DeleteResponse{
		Message: "Office Deleted Successfully!",
	}, nil
}

// Update Office
func (s *OfficeServer) UpdateOffice(ctx context.Context, req *gen.UpdateRequest) (*gen.UpdateResponse, error) {
	request := req.GetOffice()
	response := models.Office{
		Id: request.Id,
		OfficeName: request.GetOfficeName(),
		Address: request.GetAddress(),
		EmployeesCount: request.GetEmployeesCount(),
		TurnOver: request.GetTurnOver(),
	}

	err := officeService.UpdateOffice(request.GetId(), response)

	if err != nil {
		return nil, err
	}

	return &gen.UpdateResponse{
		Message: "Office Updated Successfully!",
	},nil
}

// List all Offices
func (s *OfficeServer) ListOffice(ctx context.Context, req *gen.ListRequest) (*gen.ListResponse, error) {
	offices, err := officeService.ListOffices()

	if err != nil {
		return nil, err
	}

	var officeList []*gen.Office

	for _, office := range offices{
		officeList = append(officeList, &gen.Office{
			Id:             office.Id,
			OfficeName:     office.OfficeName,
			Address:        office.Address,
			EmployeesCount: office.EmployeesCount,
			TurnOver:       office.TurnOver,
		})
	}

	return &gen.ListResponse{
		Office: officeList,
	}, nil
}
