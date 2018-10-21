package ddk

import "github.com/dcsunny/pinduoduo-sdk/util"

/**
查询订单详情
*/
type OrderDetailGetParams struct {
	OrderSn string `json:"order_sn"`
}

type OrderDetailGetResult struct {
	OrderDetailResponse OrderDetailGetInfo `json:"order_detail_response"`
	util.CommonResult
}

type OrderDetailGetInfo struct {
	OrderSn               string `json:"order_sn"`                 //订单编号
	GoodsID               int64  `json:"goods_id"`                 //商品id
	GoodsName             string `json:"goods_name"`               //商品名称
	GoodsThumbnailUrl     string `json:"goods_thumbnail_url"`      //商品缩略图
	GoodsQuantity         int    `json:"goods_quantity"`           //商品数量
	GoodsPrice            int    `json:"goods_price"`              //商品价格（分）
	OrderAmount           int    `json:"order_amount"`             //订单价格（分）
	PromotionRate         int    `json:"promotion_rate"`           //佣金比例 千分比
	PromotionAmount       int    `json:"promotion_amount"`         //佣金（分）
	BatchNo               int    `json:"batch_no"`                 //结算批次号
	OrderStatus           int    `json:"order_status"`             //订单状态
	OrderStatusDesc       string `json:"order_status_desc"`        //订单状态描述（ -1 未支付; 0-已支付；1-已成团；2-确认收货；3-审核成功；4-审核失败（不可提现）；5-已经结算；8-非多多进宝商品（无佣金订单）;10-已处罚）
	OrderCreateTime       int64  `json:"order_create_time"`        //订单创建时间（UNIX时间戳）
	OrderPayTime          int64  `json:"order_pay_time"`           //订单支付时间（UNIX时间戳）
	OrderGroupSuccessTime int64  `json:"order_group_success_time"` //订单成团时间（UNIX时间戳）
	OrderReceiveTime      int64  `json:"order_receive_time"`       //订单确认收货时间（UNIX时间戳）
	OrderVerifyTime       int64  `json:"order_verify_time"`        //订单审核时间（UNIX时间戳）
	OrderSettleTime       int64  `json:"order_settle_time"`        //订单结算时间（UNIX时间戳）
	OrderModifyAt         int64  `json:"order_modify_at"`          //订单最后更新时间（UNIX时间戳）
	MatchChannel          int    `json:"match_channel"`            //订单来源 ：0 ：关联，5 ：直接下单页RPC请求
	Type                  int    `json:"type"`                     //订单类型：0：领券页面， 1： 红包页， 2：领券页， 3： 题页
	GroupID               int64  `json:"group_id"`                 //成团编号
	AuthDuoID             int64  `json:"auth_duo_id"`              //多多客工具id
	ZsDuoID               int64  `json:"zs_duo_id"`                //招商多多客id
	CustomParameters      string `json:"custom_parameters"`        //自定义参数
	CpsSign               string `json:"cps_sign"`                 //CPS_Sign
	UrlLastGenerateTime   int64  `json:"url_last_generate_time"`   //链接最后一次生产时间
	PointTime             int64  `json:"point_time"`               //打点时间
	ReturnStatus          int    `json:"return_status"`            //售后状态：0：无，1：售后中，2：售后完成
	PID                   string `json:"p_id"`                     //推广位id
}

func (this *DuoduoKe) OrderDetailGet(p *OrderDetailGetParams) (*OrderDetailGetResult, error) {
	apiType := `pdd.ddk.order.detail.get`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result OrderDetailGetResult
	err := util.HttpPOST(url, nil, &result)
	if err != nil {
		return nil, err
	}

	err = util.CheckErrCode(result.CommonResult)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
