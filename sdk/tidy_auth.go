package sdk

import (
	"crypto/md5"
	"encoding/base32"
	"fmt"
	"time"

	"github.com/FishPiOffical/golang-sdk/types"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// PostApiGetKey 获取登录key
func (s *FishPiSDK) PostApiGetKey() error {
	conf := s.configProvider.Get()

	if conf.Username == "" || (conf.Password == "" && conf.PasswordMd5 == "") {
		return fmt.Errorf("username and password are required")
	}

	request := &types.PostApiGetKeyRequest{
		NameOrEmail:  conf.Username,
		UserPassword: conf.PasswordMd5,
	}
	if conf.PasswordMd5 == "" {
		sum := md5.Sum([]byte(conf.Password))
		request.UserPassword = fmt.Sprintf("%x", sum)
	}
	if conf.MfaCode != "" {
		request.MfaCode = conf.MfaCode
	}
	if conf.Totp != "" {
		secret := base32.StdEncoding.EncodeToString([]byte(conf.Totp))
		passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
			Period:    30,
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA512,
		})
		if err != nil {
			return fmt.Errorf("failed to generate TOTP: %w", err)
		}
		request.MfaCode = passcode
	}

	res, err := s.client.R().
		SetBodyJsonMarshal(request).
		Post("/api/getKey")
	if err != nil {
		return fmt.Errorf("failed to get api key: %w", err)
	}

	resp := new(types.PostApiGetKeyResponse)
	if err = res.Unmarshal(resp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if resp.Code != 0 {
		return fmt.Errorf("get api key failed: %s", resp.Msg)
	}

	// 更新配置中的 ApiKey
	conf.ApiKey = resp.Key
	if err = s.configProvider.Update(conf); err != nil {
		return fmt.Errorf("failed to save api key: %w", err)
	}

	// 更新 HTTP 客户端配置
	s.client.SetCommonQueryParam("apiKey", conf.ApiKey)

	return nil
}
