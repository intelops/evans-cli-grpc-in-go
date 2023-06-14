package daos

import (
	"fmt"

	"github.com/azar-intelops/grpc-for-evans/pkg/grpc/server/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var offices []models.Office

type OfficeDao struct {}

func (OfficeDao *OfficeDao) CreateOffice(office models.Office) error {
	offices = append(offices, office)
	return nil
}

func (OfficeDao *OfficeDao) GetOffice(id string) (models.Office, error) {
	if id == "" {
		return models.Office{}, status.Error(codes.InvalidArgument, "id can't be empty")
	}

	for _, office := range offices {
		if office.Id == id {
			return office, nil
		}
	}

	return models.Office{}, status.Error(codes.Internal, "something went wrong")
}

func (OfficeDao *OfficeDao) DeleteOffice(id string) error {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id can't be empty")
	} 
	
	for idx, office := range offices {
		if office.Id == id {
			offices = append(offices[:idx], offices[idx+1:]...)
			return nil
		}
	} 
	return status.Error(codes.Internal, "something went wrong")
}

func (OfficeDao *OfficeDao) UpdateOffice(id string, office models.Office) error {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id can't be empty")
	}

	for idx, office := range offices {
		if office.Id == id{
			offices[idx] = office
			return nil
		}
	}
	return status.Error(codes.Internal, "something went wrong")
}

func (OfficeDao *OfficeDao) ListOffice() ([]models.Office,error) {
	if len(offices) < 1 {
		return []models.Office{}, fmt.Errorf("data is empty %v", offices )
	}

	return offices, nil
}