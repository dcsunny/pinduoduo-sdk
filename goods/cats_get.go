package goods

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

type CatsGetParams struct {
	ParentCatID int64 `json:"parent_cat_id"`
}

type CatsGetResult struct {
	GoodsCatsGetResponse struct {
		GoodsCatsList []CatsGetInfo `json:"goods_cats_list"`
	} `json:"goods_cats_get_response"`
	common.CommonResult
}

type CatsGetInfo struct {
	Level       int    `json:"level"`
	CatID       int64  `json:"cat_id"`
	ParentCatID int64  `json:"parent_cat_id"`
	CatName     string `json:"cat_name"`
}

func (this *Goods) CatsGet(p *CatsGetParams) (*CatsGetResult, error) {
	apiType := `pdd.goods.cats.get`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result CatsGetResult
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
