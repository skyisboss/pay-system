// Code generated by ent, DO NOT EDIT.

package blockchain

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the blockchain type in the database.
	Label = "blockchain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldChain holds the string denoting the chain field in the database.
	FieldChain = "chain"
	// FieldTypes holds the string denoting the types field in the database.
	FieldTypes = "types"
	// FieldSymbol holds the string denoting the symbol field in the database.
	FieldSymbol = "symbol"
	// FieldDecimals holds the string denoting the decimals field in the database.
	FieldDecimals = "decimals"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTokenAddress holds the string denoting the token_address field in the database.
	FieldTokenAddress = "token_address"
	// FieldTokenAbi holds the string denoting the token_abi field in the database.
	FieldTokenAbi = "token_abi"
	// FieldColdAddress holds the string denoting the cold_address field in the database.
	FieldColdAddress = "cold_address"
	// FieldHotAddress holds the string denoting the hot_address field in the database.
	FieldHotAddress = "hot_address"
	// FieldScanBlockNum holds the string denoting the scan_block_num field in the database.
	FieldScanBlockNum = "scan_block_num"
	// FieldMinFreeNum holds the string denoting the min_free_num field in the database.
	FieldMinFreeNum = "min_free_num"
	// FieldMinConfirmNum holds the string denoting the min_confirm_num field in the database.
	FieldMinConfirmNum = "min_confirm_num"
	// FieldWithdrawFee holds the string denoting the withdraw_fee field in the database.
	FieldWithdrawFee = "withdraw_fee"
	// FieldWithdrawFeeType holds the string denoting the withdraw_fee_type field in the database.
	FieldWithdrawFeeType = "withdraw_fee_type"
	// FieldMinDeposit holds the string denoting the min_deposit field in the database.
	FieldMinDeposit = "min_deposit"
	// FieldMinWithdraw holds the string denoting the min_withdraw field in the database.
	FieldMinWithdraw = "min_withdraw"
	// FieldMinCollect holds the string denoting the min_collect field in the database.
	FieldMinCollect = "min_collect"
	// FieldGasPrice holds the string denoting the gas_price field in the database.
	FieldGasPrice = "gas_price"
	// Table holds the table name of the blockchain in the database.
	Table = "blockchain"
)

// Columns holds all SQL columns for blockchain fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldChain,
	FieldTypes,
	FieldSymbol,
	FieldDecimals,
	FieldStatus,
	FieldTokenAddress,
	FieldTokenAbi,
	FieldColdAddress,
	FieldHotAddress,
	FieldScanBlockNum,
	FieldMinFreeNum,
	FieldMinConfirmNum,
	FieldWithdrawFee,
	FieldWithdrawFeeType,
	FieldMinDeposit,
	FieldMinWithdraw,
	FieldMinCollect,
	FieldGasPrice,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Blockchain queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByChain orders the results by the chain field.
func ByChain(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChain, opts...).ToFunc()
}

// ByTypes orders the results by the types field.
func ByTypes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTypes, opts...).ToFunc()
}

// BySymbol orders the results by the symbol field.
func BySymbol(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSymbol, opts...).ToFunc()
}

// ByDecimals orders the results by the decimals field.
func ByDecimals(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDecimals, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTokenAddress orders the results by the token_address field.
func ByTokenAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenAddress, opts...).ToFunc()
}

// ByTokenAbi orders the results by the token_abi field.
func ByTokenAbi(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenAbi, opts...).ToFunc()
}

// ByColdAddress orders the results by the cold_address field.
func ByColdAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColdAddress, opts...).ToFunc()
}

// ByHotAddress orders the results by the hot_address field.
func ByHotAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHotAddress, opts...).ToFunc()
}

// ByScanBlockNum orders the results by the scan_block_num field.
func ByScanBlockNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScanBlockNum, opts...).ToFunc()
}

// ByMinFreeNum orders the results by the min_free_num field.
func ByMinFreeNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinFreeNum, opts...).ToFunc()
}

// ByMinConfirmNum orders the results by the min_confirm_num field.
func ByMinConfirmNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinConfirmNum, opts...).ToFunc()
}

// ByWithdrawFee orders the results by the withdraw_fee field.
func ByWithdrawFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWithdrawFee, opts...).ToFunc()
}

// ByWithdrawFeeType orders the results by the withdraw_fee_type field.
func ByWithdrawFeeType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWithdrawFeeType, opts...).ToFunc()
}

// ByMinDeposit orders the results by the min_deposit field.
func ByMinDeposit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinDeposit, opts...).ToFunc()
}

// ByMinWithdraw orders the results by the min_withdraw field.
func ByMinWithdraw(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinWithdraw, opts...).ToFunc()
}

// ByMinCollect orders the results by the min_collect field.
func ByMinCollect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinCollect, opts...).ToFunc()
}
