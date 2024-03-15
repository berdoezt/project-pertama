package controller

import (
	"net/http"
	"project-pertama/model"
	"project-pertama/repository"
	"project-pertama/util"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *orderController {
	return &orderController{
		orderRepository: orderRepository,
	}
}

func (oc *orderController) GetAll(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "unauthorized"))
		return
	}

	mapClaims, ok := claims.(map[string]any)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "unauthorized"))
		return
	}

	userId, ok := mapClaims["sub"]
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "unauthorized"))
		return
	}

	orders, err := oc.orderRepository.GetAllByUserId(userId.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	orderResponses := make([]model.OrderResponse, 0)
	for _, order := range orders {
		orderResponses = append(orderResponses, model.OrderResponse{
			UUID:        order.UUID,
			TotalAmount: order.TotalAmount,
		})
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, orderResponses, ""))
}

func (oc *orderController) Create(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := util.GetSubFromClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.CreateResponse(false, nil, "unauthorized"))
		return
	}

	var orderRequest model.Order

	err = ctx.ShouldBindJSON(&orderRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	orderRequest.UserUUID = sub.(string)
	orderCreated, err := oc.orderRepository.Create(orderRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, orderCreated, ""))
}

func (oc *orderController) Delete(ctx *gin.Context) {

}
