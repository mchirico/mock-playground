package imp

import "github.com/gin-gonic/gin"

//go:generate mockgen -destination ../../mocks/mock_car.go -package=Mocks github.com/mchirico/mock-playground/pkg/car/imp CAR

type CAR interface {
	Start(*gin.Context, string, CAR) bool
	EngineCheck(*gin.Context, CAR) bool
}
