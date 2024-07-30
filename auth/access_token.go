package auth

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
  授权 使用code换access_token
*/

type ReportOrPromotionLinkParams struct {
	Code string `json:"code"` //授权code，grantType==authorization_code 时需要
}

type SpecialParam struct {
	Key   string `json:"$key"`
	Value string `json:"$value"`
}

type ReportOrPromotionLinkResult struct {
	PopAuthTokenCreateResponse struct {
		AccessToken           string   `json:"access_token"`             //access_token
		ExpiresAt             int      `json:"expires_at"`               //access_token过期时间点
		ExpiresIn             int      `json:"expires_in"`               //access_token过期时间段，10（表示10秒后过期）
		OwnerId               string   `json:"owner_id"`                 //商家店铺id
		OwnerName             string   `json:"owner_name"`               //商家账号名称
		R1ExpiresAt           int      `json:"r1_expires_at"`            //r1级别API或字段的访问过期时间点
		R1ExpiresIn           int      `json:"r1_expires_in"`            //r1级别API或字段的访问过期时间； 10（表示10秒后过期）
		R2ExpiresAt           int      `json:"r2_expires_at"`            //r2级别API或字段的访问过期时间点
		R2ExpiresIn           int      `json:"r2_expires_in"`            //r2级别API或字段的访问过期时间；10（表示10秒后过期）
		RefreshToken          string   `json:"refresh_token"`            //refresh token，可用来刷新access_token
		RefreshTokenExpiresAt int      `json:"refresh_token_expires_at"` //Refresh token过期时间点
		RefreshTokenExpiresIn int      `json:"refresh_token_expires_in"` //refresh_token过期时间段，10表示10秒后过期
		Scope                 []string `json:"scope"`                    //接口列表
		W1ExpiresAt           int      `json:"w1_expires_at"`            //w1级别API或字段的访问过期时间点
		W1ExpiresIn           int      `json:"w1_expires_in"`            //w1级别API或字段的访问过期时间； 10（表示10秒后过期）
		W2ExpiresAt           int      `json:"w2_expires_at"`            //w2级别API或字段的访问过期时间点
		W2ExpiresIn           int      `json:"w2_expires_in"`            //w2级别API或字段的访问过期时间；10（表示10秒后过期）
	} `json:"pop_auth_token_create_response"`

	common.CommonResult
}

func (this *Auth) ReportOrPromotionLink(p *ReportOrPromotionLinkParams) (*ReportOrPromotionLinkResult, error) {
	apiType := `pdd.pop.auth.token.create`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ReportOrPromotionLinkResult
	err := util.HttpPOST(url, nil, &result)
	if err != nil {
		return nil, err
	}
	err = common.CheckErrCode(result.CommonResult)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
