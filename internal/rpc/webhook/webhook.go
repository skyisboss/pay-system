package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

// 发送回调
func HttpRequest(url, playload string) (res string, err error) {
	if err := validateURL(url); err != nil {
		return "", err
	}
	resp, body, errs := gorequest.New().
		Post(url).
		Timeout(time.Second*20).
		Send(playload).
		Set("User-Agent", "PaymentBot/1.11").
		End()

	// 回调地址无法访问
	if errs != nil {
		return body, errs[0]
	}

	// 回调地址状态非200
	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("req status error: %d", resp.StatusCode)
	}

	// 准备更新数据
	result := gin.H{}
	toJson := json.Unmarshal([]byte(body), &result)

	// 通知地址响应数据不是json
	if toJson != nil {
		return body, fmt.Errorf("err: [%T] %s", toJson, toJson.Error())
	}

	// {
	// 	err: string,
	// 	msg: string,
	// 	ret: any
	// }
	// 通知地址响应数据json里没有包含规定字段
	_, ok := result["err"]
	if !ok {
		return body, errors.New("field error not in result")
	}

	// 处理成功状态
	return body, nil
}

func validateURL(u string) error {
	parsed, err := url.ParseRequestURI(u)
	if err != nil {
		return err
	}

	if parsed.Hostname() == "" {
		return errors.New("invalid hostname")
	}

	return nil
}
