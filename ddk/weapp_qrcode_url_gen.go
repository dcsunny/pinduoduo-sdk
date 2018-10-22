package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多客生成单品推广小程序二维码url
*/
type WeappQrcodeUrlGenParams struct {
	PID              string  `json:"p_id"`                        //推广位ID
	GoodsIdList      []int64 `json:"goods_id_list"`               //商品ID，仅支持单个查询
	CustomParameters *string `json:"custom_parameters,omitempty"` //自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	ZsDuoID          *int64  `json:"zs_duo_id,omitempty"`         //招商多多客ID
}

type WeappQrcodeUrlGenResult struct {
	WeappQrcodeGenerateResponse struct {
		Url string `json:"url"` //单品推广小程序二维码url
	} `json:"weapp_qrcode_generate_response"`
	common.CommonResult
}

func (this *DuoduoKe) WeappQrcodeUrlGen(p *WeappQrcodeUrlGenParams) (*WeappQrcodeUrlGenResult, error) {
	apiType := `pdd.ddk.weapp.qrcode.url.gen`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result WeappQrcodeUrlGenResult
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
