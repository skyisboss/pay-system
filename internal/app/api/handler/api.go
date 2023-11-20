package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/balance"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

func (h *Handler) ApplyAddress(ctx *gin.Context) {
	ioc := h.Ioc
	context := context.Background()
	var req struct {
		// AppID string `json:"appid" binding:"required" validate:"alphanum"`
		// Nonce string `json:"nonce" binding:"required" validate:"min=6,max=42,alphanum"`
		// Sign  string `json:"sign" binding:"required" validate:"len=32,alphanum"`
		Chain string `json:"chain" binding:"required" validate:"oneof=eth tron bsc"`
	}
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "参数不完整", "data": gin.H{}})
		ctx.Abort()
		return
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": err.Error(), "data": gin.H{}})
		ctx.Abort()
		return
	}

	// productID := ctx.GetInt64("product_id")
	// if productID == 0 {
	// 	ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "无法获取appid", "data": gin.H{}})
	// 	ctx.Abort()
	// 	return
	// }

	// 获取币种配置信息
	cfgs, err := ioc.BlockchainService().GetAll(context)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("BlockchainService GetAll")
		return
	}
	var chainIDs []uint64
	for _, cfg := range cfgs {
		if cfg.Status != 1 {
			continue
		}
		if req.Chain == cfg.Chain && req.Chain == cfg.Types {
			chainIDs = append(chainIDs, cfg.ID)
		}
	}

	// chainIDs := util.MapSlice(cfgs, func(e *ent.Blockchain) uint64 {
	// 	return e.ID
	// })

	// 分配一个新地址
	// address, err := ioc.AddressService().GetNewAddress(context, chainID)
	// if err != nil {
	// 	ioc.Logger().Error().Err(err).Msg("查询address")
	// 	return
	// }

	// 开启事物，分配balance
	tx, err := ioc.DBClient().Tx(ctx)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("new transactional client")
		return
	}
	// 创建product
	appid := util.GetUUID()
	appName := util.GetUUID()
	appSecret := util.GetUUID()
	product, err := tx.Product.Create().
		SetAppID(appid).
		SetAppName(appName).
		SetAppSecret(appSecret).
		SetAppStatus(1).
		SetWithdrawStatus(1).
		SetWebHook("").
		SetCreatedAt(time.Now()).
		Save(context)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("Product.Create")
		tx.Rollback()
		return
	}

	// 创建地址
	for _, id := range chainIDs {
		chainID := id
		address, err := ioc.AddressService().GetNewAddress(context, chainID)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("查询address")
			return
		}

		// 更新 UseTo
		_, err = tx.Addres.UpdateOneID(address.ID).SetUseTo(int64(chainID)).Save(context)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("Update")
			tx.Rollback()
			return
		}

	}

	// 创建余额
	rows := []*ent.BalanceCreate{}
	for _, id := range chainIDs {
		chainID := id
		row := tx.Balance.Create().
			SetChainID(chainID).
			SetProductID(uint64(product.ID))

		rows = append(rows, row)
	}
	_, err = tx.Balance.CreateBulk(rows...).Save(context)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("CreateBulk")
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("Commit")
		tx.Rollback()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"err": 0, "msg": "ok", "data": gin.H{}})
}

// 申请提币
func (h *Handler) ApplyWithdraw(ctx *gin.Context) {
	ioc := h.Ioc
	context := context.Background()
	var req struct {
		AppID   string `json:"appid" binding:"required" validate:"alphanum"`
		Nonce   string `json:"nonce" binding:"required" validate:"min=6,max=42,alphanum"`
		Chain   string `json:"chain" binding:"required" validate:"oneof=eth tron bsc"`
		Symbol  string `json:"symbol" binding:"required" validate:"oneof=eth trx bnb usdt"`
		Address string `json:"address" binding:"required"`
		Amount  string `json:"amount" binding:"required"`
		Sign    string `json:"sign" binding:"required" validate:"len=32,alphanum"`
	}
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "参数不完整", "data": gin.H{}})
		ctx.Abort()
		return
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": err.Error(), "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 判断nonce是否已经存在
	ok, err := ioc.SessionService().GetSession(context, req.AppID, req.Nonce)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": err.Error(), "data": gin.H{}})
		ctx.Abort()
		return
	}
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "nonce重复", "data": gin.H{}})
		ctx.Abort()
		return
	}
	_, err = ioc.SessionService().Create(context, req.AppID, req.Nonce, ctx.ClientIP())
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("SessionService Create")
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "内部错误", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 获取产品信息
	productRow, err := ioc.ProductService().GetByAppID(context, req.AppID)
	if ent.IsNotFound(err) || productRow == nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "商户不存在", "data": gin.H{}})
		ctx.Abort()
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "内部错误", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 验证签名
	checkSign := util.WechatSign(productRow.AppSecret, gin.H{
		"appid":   req.AppID,
		"nonce":   req.Nonce,
		"chain":   req.Chain,
		"symbol":  req.Symbol,
		"address": req.Address,
		"amount":  req.Amount,
	})
	if checkSign != req.Sign {
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "无效签名", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 是否允许提款
	if productRow.WithdrawStatus != 1 {
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "商户提款限制", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 根据chain和symbol获取配置信息
	cfg, err := ioc.BlockchainService().GetByChainAndSymbol(context, req.Chain, req.Symbol)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("GetByChainAndSymbol")
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "获取币种错误", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 根据配置获取Provider 验证address地址是否有效
	Provider := ioc.WalletService().GetProvider(wallet.Blockchain(cfg.Chain))
	if ok := Provider.ValidateAddress(req.Address); !ok {
		ioc.Logger().Error().Err(err).Msg("ValidateAddress")
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "提款地址错误", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 提款金额转换 因为此金额可能存在小数点，所以需要先转成 decimal.Decimal，再转成对应币种的最小单位
	amountObj, err := decimal.NewFromString("0.02")
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("NewFromString")
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "金额错误", "data": gin.H{}})
		ctx.Abort()
		return
	}
	// 判断余额是否大于0
	if amountObj.Cmp(decimal.Zero) <= 0 {
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "金额错误", "data": gin.H{}})
		ctx.Abort()
		return
	}
	amountRaw := amountObj.Mul(decimal.New(1, int32(cfg.Decimals)))

	// 余额是否足够
	balanceRows, err := ioc.BalanceService().ListByProductID(context, productRow.ID)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("ListByProductID")
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "余额错误", "data": gin.H{}})
		ctx.Abort()
		return
	}
	balanceMap := util.KeyFunc(balanceRows, func(e *ent.Balance) uint64 {
		return e.ChainID
	})
	balanceRow := balanceMap[cfg.ID]
	letBalance := balanceRow.BalanceAmount.Sub(amountRaw)
	if letBalance.Cmp(decimal.Zero) < 0 {
		ctx.JSON(http.StatusOK, gin.H{"err": 400, "msg": "余额不足", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 添加提款信息
	// 开启事物
	tx, err := ioc.DBClient().Tx(ctx)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("DBClient")
		ioc.Logger().Error().Err(err).Msg("new transactional client")
		return
	}
	// 1、更新余额信息
	_, err = tx.Balance.Update().
		Where(balance.IDEQ(balanceRow.ID)).
		Where(balance.VersionEQ(balanceRow.Version)).
		SetBalanceAmount(letBalance).
		SetBalanceFreeze(amountRaw).
		SetCountWithdraw(balanceRow.CountWithdraw + 1).
		SetVersion(balanceRow.Version + 1).
		Save(context)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("Update")
		tx.Rollback()
		return
	}
	// 2、添加提款数据
	serialId := util.GetUUID()
	_, err = tx.Withdraw.Create().
		SetProductID(int64(productRow.ID)).
		SetChainID(cfg.ID).
		SetSerialID(serialId).
		SetToAddress(req.Address).
		SetAmountRaw(amountRaw).
		SetAmountStr(req.Amount).
		Save(context)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("Create")
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("Commit")
		tx.Rollback()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"err": 0, "msg": "ok", "data": gin.H{
		"serial_id": serialId,
	}})
}

// 创建账号
func (h *Handler) CreateAccount(ctx *gin.Context) {
	ioc := h.Ioc
	context := context.Background()
	var req struct {
		Chain string `json:"chain" binding:"required" validate:"oneof=eth tron bsc"`
	}
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "参数不完整", "data": gin.H{}})
		ctx.Abort()
		return
	}

	if !util.InArray([]string{"eth", "tron", "bsc"}, req.Chain) {
		ctx.JSON(http.StatusOK, gin.H{"err": 500, "msg": "参数错误", "data": gin.H{}})
		ctx.Abort()
		return
	}

	// 获取币种配置信息
	cfgs, err := ioc.BlockchainService().GetAll(context)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("BlockchainService GetAll")
		return
	}

	// chainIDs := util.MapSlice(cfgs, func(e *ent.Blockchain) uint64 {
	// 	return e.ID
	// })

	// 开启事物，分配balance
	tx, err := ioc.DBClient().Tx(ctx)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("new transactional client")
		return
	}

	// 创建product
	appid := util.GetUUID()
	appName := util.GetUUID()
	appSecret := util.GetUUID()
	product, err := tx.Product.Create().
		SetAppID(appid).
		SetAppName(appName).
		SetAppSecret(appSecret).
		SetAppStatus(1).
		SetWithdrawStatus(1).
		SetWebHook("").
		SetCreatedAt(time.Now()).
		Save(context)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("Product.Create")
		tx.Rollback()
		return
	}

	// 创建地址
	for _, cfg := range cfgs {
		chainID := cfg.ID
		if cfg.TokenAbi == "" && cfg.TokenAddress == "" {
			address, err := ioc.AddressService().GetNewAddress(context, chainID)
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("查询address")
				return
			}
			// 更新 UseTo
			_, err = tx.Addres.UpdateOneID(address.ID).SetUseTo(int64(product.ID)).Save(context)
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("Update")
				tx.Rollback()
				return
			}
		}

		// 新增余额
		_, err = tx.Balance.Create().
			SetChainID(chainID).
			SetProductID(uint64(product.ID)).
			SetBalanceAmount(decimal.Zero).
			SetBalanceFreeze(decimal.Zero).
			SetTotalDeposit(decimal.Zero).
			SetTotalWithdraw(decimal.Zero).
			Save(context)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("Balance.Create")
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("Commit")
		tx.Rollback()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"err": 0, "msg": "ok", "data": gin.H{}})
}
