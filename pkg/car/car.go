package car

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mchirico/mock-playground/pkg/car/imp"
)

type CARSTRUCT struct {
}

func (car CARSTRUCT) Start(context *gin.Context, key string, carInterface imp.CAR) bool {
	if carInterface.EngineCheck(context, carInterface) {
		if key == "carkey" {
			fmt.Println("Let's start the car!")
			return true
		} else {
			fmt.Println("You don't have the right key!")
			return false
		}
	} else {
		return false
	}
}
func (car CARSTRUCT) EngineCheck(context *gin.Context, carInterface imp.CAR) bool {
	return true
}
