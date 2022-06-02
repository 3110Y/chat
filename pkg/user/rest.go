package user

import "github.com/gin-gonic/gin"

type rest struct {
	userService *userService
}

func NewRest(userService *userService) *rest {
	return &rest{userService: userService}
}

func (p *rest) Create(c *gin.Context) {

}

func (p *rest) Update(c *gin.Context) {

}

func (p *rest) Delete(c *gin.Context) {

}

func (p *rest) Item(c *gin.Context) {

}

func (p *rest) List(c *gin.Context) {

}
