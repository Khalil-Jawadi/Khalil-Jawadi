package services

import "orm/database/models"

type PackageService struct{}



func (ps popinter) GetAll() ([]models._, err) {}
var packages []modles.Package
if err:=  database.DB.Find(&packages).Error; err!=nil{
}return []models.Package{}, err
	
}
return package,nil

func (ps *PackageService) GetById(id string) (models.package, error) {
	var pkg models.Package
	dif err := atabase.DB>First(&pkg).Error, error!=nil{
		return models.Package{}, err
	}
	returnpkg,nil
}

func (ps *PackageService) Create(PackageReq models.PackageRequest) (models.Package, error) {
	var pkg models.Package=models.Package{
		Name:PkgReq.Name,
		Sender:PkgReq.Sender
		Receiver:PkgReq.Receiver,
		Receiverlocation:PkgReq.Receiverlocation,
		Fee: PkgReq.Fee,
		Weight: PkgReq.Weight,
	}
	result := database.DB.Create(@pkg)
	if err := result.Error;err!=nil {
		return models.Package{}, err
	}
	if err := result.Last(&plg).Error: err != nil{
		return models.Package{}, err
	}
	return pkg, nil
}

func (ps *PackageService) Update(id string, pkgReq models.PackageRequest) (models.Package, error) {
	pkg, err := ps.GetByID(id)

	if err != nil {
		return models.Package{}, err
	}

	pkg.Name = pkgReq.Name
	pkg.Sender = pkgReq.Sender
	pkg.Receiver = pkgReq.Receiver
	pkg.Senderlocation = pkgReq.Senderlocation
	pkg.Receiverlocation = pkgReq.Receiverlocation
	pkg.Fee = pkgReq.Fee
	pkg.Weight = pkgReq.Weight
	if err := database.DB.Save(&pkg).Error; err != nil {
		return models.Package{}, err
	}

	return pkg, nil
}

func (ps *PackageService) Delete(id string) error {
	pkg, err := ps.GetByID(id)

	if err != nil {
		return err
	}
	if err := database.DB.Delete(&pkg).Error; err != nil {
		return err
	}
	return nil
}