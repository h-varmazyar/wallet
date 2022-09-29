package wallets

import (
	"github.com/gin-gonic/gin"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	"github.com/h-varmazyar/wallet/pkg/grpcext"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	walletService walletApi.WalletServiceClient
	logger        *log.Logger
}

func NewHandler(configs *Configs, logger *log.Logger) *Handler {
	walletConn := grpcext.NewConnection(configs.TransactionServiceAddress)
	handler := &Handler{
		logger:        logger,
		walletService: walletApi.NewWalletServiceClient(walletConn),
	}

	return handler
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	walletRoutes := router.Group("/wallets")

	walletRoutes.GET("/phone/:phone_number", h.returnByPhoneNumber)

}

func (h *Handler) returnByPhoneNumber(c *gin.Context) {
	req := new(walletApi.WalletReturnByPhoneReq)
	req.PhoneNumber = c.Param("phone_number")
	if wallet, err := h.walletService.ReturnByPhoneNumber(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, wallet)
	}
}
