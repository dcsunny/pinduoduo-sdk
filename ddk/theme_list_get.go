package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
查询多多进宝主题列表
*/

type ThemeListGetParams struct {
	PageSize int `json:"page_size,omitempty"` //返回的一页数据数量
	Page     int `json:"page,omitempty"`      //返回的页码
}

type ThemeListGetResult struct {
	ThemeListGetResponse struct {
		ThemeList []ThemeListGetInfo `json:"theme_list"` //返回的主题列表
		Total     int                `json:"total"`      //返回的元素数量
	} `json:"theme_list_get_response"`
	common.CommonResult
}

type ThemeListGetInfo struct {
	ID       int64  `json:"id"`        //主题ID
	ImageUrl string `json:"image_url"` //主题图片
	Name     string `json:"name"`      //主题名称
	GoodsNum int    `json:"goods_num"` //主题包含的商品数量
}

func (this *DuoduoKe) ThemeListGet(p *ThemeListGetParams) (*ThemeListGetResult, error) {
	apiType := `pdd.ddk.theme.list.get`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result ThemeListGetResult
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
