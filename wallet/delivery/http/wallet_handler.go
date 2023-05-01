package http

import (
	"mini-wallet/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	walletHandler struct {
		accountUseCase domain.AccountUseCase
		walletUsecase  domain.WalletUseCase
	}
)

func NewWalletHandler(r *gin.Engine, us domain.AccountUseCase, wus domain.WalletUseCase) {
	handler := &walletHandler{us, wus}
	r.SetTrustedProxies(nil)

	rg := r.Group("/api/v1")
	rg.POST("/init", handler.initAccount)

	walletRG := rg.Group("/wallet")
	walletRG.POST("", handler.enableWallet)
	walletRG.GET("", handler.getWallet)
	walletRG.PATCH("", handler.disableWallet)
	walletRG.POST("/deposits", handler.depositWallet)
	walletRG.PATCH("/withdrawals", handler.withdrawWallet)
	walletRG.GET("/transactions", handler.getTransaction)

}

func (h *walletHandler) getTransaction(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Token ")

	transactions, err := h.walletUsecase.GetWalletTransactions(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when get wallet transactions",
		})
		return
	}
	c.JSON(http.StatusCreated, response{
		Status: "success",
		Data: map[string]interface{}{
			"transations": transactions,
		},
	})

}

func (h *walletHandler) withdrawWallet(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Token ")

	amountStr := c.PostForm("amount")
	refID := c.PostForm("reference_id")
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when withdraw wallet",
		})
		return
	}

	wallet, err := h.walletUsecase.WithdrawMoneyWallet(token, refID, amount)
	if err != nil {
		if err == domain.ErrWalletMustEnabled || err == domain.ErrWalletNotFound || err == domain.ErrRefIDTransactionAlreadyExists || err == domain.ErrWalletInsufficantBalance {
			c.JSON(http.StatusBadRequest, response{
				Status:  "error",
				Message: domain.ErrWalletAlreadyEnabled.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when withdraw wallet",
		})
		return
	}

	c.JSON(http.StatusCreated, response{
		Status: "success",
		Data: map[string]interface{}{
			"deposit": depositResponseDTO(wallet, refID, amount),
		},
	})
}

func (h *walletHandler) depositWallet(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Token ")

	amountStr := c.PostForm("amount")
	refID := c.PostForm("reference_id")
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when deposit wallet",
		})
		return
	}

	wallet, err := h.walletUsecase.DepositMoneyWallet(token, refID, amount)
	if err != nil {
		if err == domain.ErrWalletMustEnabled || err == domain.ErrWalletNotFound || err == domain.ErrRefIDTransactionAlreadyExists {
			c.JSON(http.StatusBadRequest, response{
				Status:  "error",
				Message: domain.ErrWalletAlreadyEnabled.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when deposit wallet",
		})
		return
	}

	c.JSON(http.StatusCreated, response{
		Status: "success",
		Data: map[string]interface{}{
			"deposit": depositResponseDTO(wallet, refID, amount),
		},
	})
}

func (h *walletHandler) enableWallet(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Token ")

	wallet, err := h.walletUsecase.EnableWallet(token)
	if err != nil {
		if err == domain.ErrWalletAlreadyEnabled {
			c.JSON(http.StatusBadRequest, response{
				Status:  "error",
				Message: domain.ErrWalletAlreadyEnabled.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when enabling wallet",
		})
		return
	}

	c.JSON(http.StatusCreated, response{
		Status: "success",
		Data: map[string]interface{}{
			"wallet": walletResponseDTO(wallet),
		},
	})

}

func (h *walletHandler) disableWallet(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Token ")

	wallet, err := h.walletUsecase.DisableWallet(token)
	if err != nil {
		if err == domain.ErrWalletAlreadyDisabled {
			c.JSON(http.StatusBadRequest, response{
				Status:  "error",
				Message: domain.ErrWalletAlreadyDisabled.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when enabling wallet",
		})
		return
	}

	c.JSON(http.StatusCreated, response{
		Status: "success",
		Data: map[string]interface{}{
			"wallet": walletResponseDTO(wallet),
		},
	})

}

func (h *walletHandler) getWallet(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	token := strings.TrimPrefix(authHeader, "Token ")

	wallet, err := h.walletUsecase.GetWalletBalance(token)
	if err != nil {
		if err == domain.ErrWalletMustEnabled {
			c.JSON(http.StatusBadRequest, response{
				Status:  "error",
				Message: domain.ErrWalletMustEnabled.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, response{
			Status:  "error",
			Message: "Error when getting wallet",
		})
		return
	}

	c.JSON(http.StatusCreated, response{
		Status: "success",
		Data: map[string]interface{}{
			"wallet": walletResponseDTO(wallet),
		},
	})
}

func (h *walletHandler) initAccount(c *gin.Context) {
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
