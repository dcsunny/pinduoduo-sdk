package goods

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
获得拼多多商品标签列表（非商品类目cat，当前仅开放给多多客使用）
*/

type OptGetParams struct {
	ParentOptID int64 `json:"parent_opt_id"` //值=0时为顶点opt_id,通过树顶级节点获取opt树
}

type OptGetResult struct {
	GoodsOptGetResponse struct {
		GoodsOptList []OptGetInfo `json:"goods_opt_list"` //opt列表
	} `json:"goods_opt_get_response"`
	common.CommonResult
}

type OptGetInfo struct {
	Level       int    `json:"level"`         //层级，1-一级，2-二级，3-三级，4-四级
	ParentOptID int64  `json:"parent_opt_id"` //id所属父ID，其中，parent_id=0时为顶级节点
	OptName     string `json:"opt_name"`      //商品标签名
	OptID       int64  `json:"opt_id"`        //商品标签ID
}

func (this *Goods) OptGet(p *OptGetParams) (*OptGetResult, error) {
	apiType := `pdd.goods.opt.get`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result OptGetResult
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
