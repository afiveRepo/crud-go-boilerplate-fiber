package services

import (
	"crud-go-boilerplate-fiber/app/models/entities"
	"crud-go-boilerplate-fiber/app/models/requests"
	"crud-go-boilerplate-fiber/app/models/responses"
	"crud-go-boilerplate-fiber/app/repository"
	"errors"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(req requests.InputProduct) (*responses.ProductResponse, error)
	Update(req requests.UpdateProduct) (*responses.ProductResponse, error)
	FindAll() (*[]responses.ProductResponse, error)
	FindByID(uid string) (*responses.ProductResponse, error)
}
type productService struct {
	base    repository.BaseRepository
	product repository.ProductRepository
}

func NewProductService(base repository.BaseRepository, p repository.ProductRepository) ProductService {
	return &productService{
		base:    base,
		product: p,
	}
}
func (s *productService) Create(req requests.InputProduct) (*responses.ProductResponse, error) {
	uid := uuid.New().String()
	newProduct := entities.Product{
		ProductUID:         uid,
		ProductName:        req.ProductName,
		ProductStok:        req.ProductStok,
		PoductPrice:        req.PoductPrice,
		ProductInformation: req.ProductInformation,
	}
	tx := s.base.GetBegin()
	product, err := s.product.Create(tx, newProduct)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	res := responses.ProductResponse{
		ProductID:          uint64(product.ID),
		ProductName:        product.ProductName,
		PoductPrice:        product.PoductPrice,
		ProductStok:        product.ProductStok,
		ProductInformation: product.ProductInformation,
	}
	return &res, nil
}
func (s *productService) Update(req requests.UpdateProduct) (*responses.ProductResponse, error) {
	product, err := s.product.FindByUID(req.ProductUID)
	if err != nil {
		errs := errors.New("product not found")
		return nil, errs
	}
	product.ProductName = req.ProductName
	product.ProductUID = req.ProductUID
	product.PoductPrice = req.PoductPrice
	product.ProductStok = product.ProductStok
	product.ProductInformation = product.ProductInformation
	tx := s.base.GetBegin()
	NewProduct, err := s.product.Update(tx, *product)
	if err != nil {
		tx.Rollback()
		errs := errors.New("update fail")
		return nil, errs
	}
	tx.Commit()
	res := responses.ProductResponse{
		ProductID:          uint64(NewProduct.ID),
		ProductName:        NewProduct.ProductName,
		PoductPrice:        NewProduct.PoductPrice,
		ProductStok:        NewProduct.ProductStok,
		ProductInformation: NewProduct.ProductInformation,
	}
	return &res, nil
}
func (s *productService) FindAll() (*[]responses.ProductResponse, error) {
	var res []responses.ProductResponse
	product, err := s.product.FindAll()
	if err != nil {
		errs := errors.New("product not found")
		return nil, errs
	}
	for _, v := range *product {
		row := responses.ProductResponse{
			ProductID:          uint64(v.ID),
			ProductName:        v.ProductName,
			PoductPrice:        v.PoductPrice,
			ProductStok:        v.ProductStok,
			ProductInformation: v.ProductInformation,
		}
		res = append(res, row)
	}

	return &res, nil
}
func (s *productService) FindByID(uid string) (*responses.ProductResponse, error) {
	product, err := s.product.FindByUID(uid)
	if err != nil {
		errs := errors.New("product not found")
		return nil, errs
	}
	res := responses.ProductResponse{
		ProductID:          uint64(product.ID),
		ProductName:        product.ProductName,
		PoductPrice:        product.PoductPrice,
		ProductStok:        product.ProductStok,
		ProductInformation: product.ProductInformation,
	}
	return &res, nil
}
