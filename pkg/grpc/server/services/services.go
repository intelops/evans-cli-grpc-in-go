package services

import (
	"github.com/azar-intelops/grpc-for-evans/pkg/grpc/server/daos"
	"github.com/azar-intelops/grpc-for-evans/pkg/grpc/server/models"
)



type OfficeService struct{

}

var officeDao = daos.OfficeDao{}

func (officeService *OfficeService) CreateOffice(office models.Office) error  {
	return officeDao.CreateOffice(office)
}

func (officeService *OfficeService) GetOffice(id string) (models.Office, error)  {
	return officeDao.GetOffice(id)
}

func (officeService *OfficeService) ListOffices() ([]models.Office, error)  {
	return officeDao.ListOffice()
}

func (officeService *OfficeService) UpdateOffice(id string, office models.Office) (error)  {
	return officeDao.UpdateOffice(id, office)
}

func (officeService *OfficeService) DeleteOffice(id string) (error)  {
	return officeDao.DeleteOffice(id)
}