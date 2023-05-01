package main

import (
	"fmt"
	accRepo "mini-wallet/account/repository/postgre"
	accUC "mini-wallet/account/usecase"
	"mini-wallet/config"
	"mini-wallet/wallet/delivery/http"
	wRepo "mini-wallet/wallet/repository/postgre"
	wUC "mini-wallet/wallet/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := config.GetConnection()

	r := gin.New()

	accountRepo := accRepo.NewAccountRepository(conn.PostgreCon)
	walletRepo := wRepo.NewWalletRepository(conn.PostgreCon)

	accountUsecase := accUC.NewAccountUseCase(accountRepo)
	walletUseCase := wUC.NewWalletUseCase(walletRepo, accountUsecase)
	http.NewWalletHandler(r, accountUsecase, walletUseCase)

	r.Run(fmt.Sprintf(":%v", config.Conf.ServicePort))
}
