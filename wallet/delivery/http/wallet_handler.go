package http

import (
	"mini-wallet/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	walletHandler struct {
		accountUseCase domain.AccountUseCase
	}
)

func NewWalletHandler(r *gin.Engine, us domain.AccountUseCase) {
	handler := &walletHandler{us}
	r.SetTrustedProxies(nil)

	rg := r.Group("/api/v1")
	rg.POST("/init", handler.InitAccount)

}

func (h *walletHandler) InitAccount(c *gin.Context) {
	xid := c.PostForm("customer_xid")
	token, err := h.accountUseCase.CreateAccount(xid)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when creating account",
		})
		return
	}

	c.JSON(http.StatusCreated, response{
		Status: "success",
		Data: map[string]string{
			"token": token,
		},
	})
}
