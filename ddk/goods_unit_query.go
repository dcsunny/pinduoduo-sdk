package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
查询商品的推广计划
*/
type GoodsUnitQueryParams struct {
	GoodsID int64  `json:"goods_id"`            //商品id
	ZsDuoID *int64 `json:"zs_duo_id,omitempty"` //招商多多客ID
}

type GoodsUnitQueryResult struct {
	DdkGoodsUnitQueryResponse struct {
		CouponEndTime   int64  `json:"coupon_end_time"`   //优惠券失效时间，UNIX时间戳
		CouponStartTime int64  `json:"coupon_start_time"` //优惠券生效时间，UNIX时间戳
		CouponID        string `json:"coupon_id"`         //优惠券id
		Discount        int    `json:"discount"`          //优惠券面额，单位：厘
		InitQuantity    int64  `json:"init_quantity"`     //优惠券总数量
		Rate            int    `json:"rate"`              //商品的佣金比例，单位：千分位，比如100，表示10%
		RemainQuantity  int    `json:"remain_quantity"`   //优惠券剩余数量
		UnitType        int    `json:"unit_type"`         //商品的推广计划类型，1-通用推广，2-专属推广，3-招商推广，4-全店推广
	} `json:"ddk_goods_unit_query_response"`
	common.CommonResult
}

func (this *DuoduoKe) GoodsUnitQuery(p *GoodsUnitQueryParams) (*GoodsUnitQueryResult, error) {
	apiType := "pdd.ddk.goods.unit.query"
	params, paramsURL := util.FormatURLParams(p)

	url := this.GetURL(apiType, "", params, paramsURL)

	var result GoodsUnitQueryResult

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
