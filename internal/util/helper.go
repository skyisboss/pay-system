package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func Println(a ...any) {
	fmt.Println("🟥🟩🟨")
	fmt.Println("")
	fmt.Println(a...)
	fmt.Println("")
	fmt.Println("🟥🟩🟨")
}

func Ptr[T any](v T) *T {
	return &v
}

// GetUUID 获取唯一字符串
func GetUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// IsStringInSlice 字符串是否在数组中
func IsStringInSlice[T comparable](arr []T, str T) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// IsIntInSlice 数字是否在数组中
func IsIntInSlice[T comparable](arr []T, str T) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// 是否在数组中
func InArray[T comparable](arr []T, find T) bool {
	for _, v := range arr {
		if v == find {
			return true
		}
	}
	return false
}

// AddressBytesToStr 地址转化为字符串
func AddressBytesToStr(addressBytes common.Address) string {
	return strings.ToLower(addressBytes.Hex())
}

// 字符串翻转
func StrReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 转换金额 wei to eth
func WeiToEth(wei *big.Int, decimals int32) (string, error) {
	balance, err := decimal.NewFromString(wei.String())
	if err != nil {
		return "0", err
	}
	balanceStr := balance.Mul(decimal.New(1, -decimals)).StringFixed(decimals)
	return balanceStr, nil
}

func EthToWei(eth *big.Int, decimals int32) (string, error) {
	balance, err := decimal.NewFromString(eth.String())
	if err != nil {
		return "0", err
	}
	balanceStr := balance.Mul(decimal.New(1, decimals)).StringFixed(decimals)
	return balanceStr, nil
}

// big.Int to decimal.Decimal
func BigintToDecimal(val *big.Int) (decimal.Decimal, error) {
	ret, err := decimal.NewFromString(val.String())
	if err != nil {
		return decimal.Decimal{}, err
	}
	return ret, nil
}

// str to bigint
func StrToBigInit(str string, decimals int32) (*big.Int, error) {
	amountReal, err := decimal.NewFromString(str)
	if err != nil {
		return nil, err
	}
	amount := amountReal.Mul(decimal.New(1, decimals)).StringFixed(0)
	b := new(big.Int)
	_, ok := b.SetString(amount, 10) // 10进制
	if !ok {
		return nil, errors.New("error str to bigint")
	}
	return b, nil
}

// 美化打印
func ToJson(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	fmt.Println(out.String())
}

// 字节转换为地址
func HashToAddrssStringLower(bytes common.Hash) string {
	return strings.ToLower(HashToAddrss(bytes).Hex())
}
func HashToAddrss(bytes common.Hash) common.Address {
	var b common.Address
	b.SetBytes(bytes[:])
	return b
}

// 获取正在运行的函数名
func GetFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
