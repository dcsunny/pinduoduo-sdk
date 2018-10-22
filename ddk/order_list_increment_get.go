package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
增量查询推广订单信息（根据最后更新时间）
*/
type OrderListIncrementGetParams struct {
	StartUpdateTime int64 `json:"start_update_time"`   //最近90天内多多进宝商品订单更新时间--查询时间开始。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	EndUpdateTime   int64 `json:"end_update_time"`     //最近90天内多多进宝商品订单更新时间--查询时间结束。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	PageSize        *int  `json:"page_size,omitempty"` //返回的每页结果订单数，默认为100，范围为10到100，建议使用40~50，可以提高成功率，减少超时数量。
	Page            *int  `json:"page,omitempty"`      //第几页，从1到10000，默认1
}

type OrderListIncrementGetResult struct {
	OrderListGetResponse struct {
		OrderList []OrderListIncrementGetInfo `json:"order_list"` //多多进宝推广位对象列表
	} `json:"order_list_get_response"`
	TotalCount      int `json:"total_count"`       //请求到的结果数
	DuoCouponAmount int `json:"duo_coupon_amount"` //使用多多进宝券的面额（单位为分）
	common.CommonResult
}

type OrderListIncrementGetInfo struct {
	OrderReceiveTime      int64  `json:"order_receive_time"`       //订单确认签收时间
	CustomParameters      string `json:"custom_parameters"`        //自定义参数，标志订单来源于哪个自定义参数
	Type                  int    `json:"type"`                     //订单来源：0-单品（领券页）推广，1-红包活动推广，2-领券页底部推荐，54-大转盘主页的商品，55-抽中免单的商品，56-大转盘抽中红包的商品，13-大转盘拉新锁佣的
	OrderVerifyTime       int64  `json:"order_verify_time"`        //审核时间
	OrderPayTime          int64  `json:"order_pay_time"`           //支付时间
	OrderGroupSuccessTime int64  `json:"order_group_success_time"` //成团时间
	OrderModifyAt         int64  `json:"order_modify_at"`          //最后更新时间
	OrderStatusDesc       string `json:"order_status_desc"`        //订单状态描述
	PID                   string `json:"p_id"`                     //推广位ID
	OrderStatus           int    `json:"order_status"`             //订单状态： -1 未支付; 0-已支付；1-已成团；2-确认收货；3-审核成功；4-审核失败（不可提现）；5-已经结算；8-非多多进宝商品（无佣金订单）;10-已处罚
	PromotionAmount       int    `json:"promotion_amount"`         //佣金金额，单位为分
	PromotionRate         int    `json:"promotion_rate"`           //佣金比例，千分比
	OrderCreateTime       int64  `json:"order_create_time"`        //订单生成时间，UNIX时间戳
	OrderAmount           int    `json:"order_amount"`             //实际支付金额，单位为分
	GoodsPrice            int    `json:"goods_price"`              //订单中sku的单件价格，单位为分
	GoodsQuantity         int    `json:"goods_quantity"`           //购买商品的数量
	GoodsThumbnailUrl     string `json:"goods_thumbnail_url"`      //商品缩略图
	GoodsName             string `json:"goods_name"`               //商品标题
	GoodsID               int64  `json:"goods_id"`                 //商品ID
	OrderSn               string `json:"order_sn"`                 //推广订单编号
}

func (this *DuoduoKe) OrderListIncrementGet(p *OrderListIncrementGetParams) (*OrderListIncrementGetResult, error) {
	apiType := `pdd.ddk.order.list.increment.get`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result OrderListIncrementGetResult
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
