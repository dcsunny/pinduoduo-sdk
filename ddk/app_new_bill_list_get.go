package ddk

import "github.com/dcsunny/pinduoduo-sdk/util"

/**
获取并查看多多客拉新的奖励账单
*/
type AppNewBillListGetParams struct {
	PID       *string `json:"pid,omitempty"`
	StartTime int64   `json:"start_time"`          //最后更新时间--查询时间开始。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	EndTime   int64   `json:"end_time"`            //最后更新时间--查询时间开始。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	PageSize  *int    `json:"page_size,omitempty"` //返回的每页结果列表数，默认为100，范围为10到100，建议使用40~50，可以提高成功率，减少超时数量。
	Page      *int    `json:"page,omitempty"`      //第几页，从1到10000，默认1，注：使用最后更新时间范围增量同步时，必须采用倒序的分页方式（从最后一页往回取）才能避免漏的问题
}

type AppNewBillListGetResult struct {
	AppNewBillListResponse struct {
		List []AppNewBillListGetInfo `json:"list"` //账单列表对象
	} `json:"app_new_bill_list_response"`
	TotalCount int `json:"total_count"` //账单总数
	util.CommonResult
}

type AppNewBillListGetInfo struct {
	PID              string `json:"pid"`               //推广位ID
	OrderSn          string `json:"order_sn"`          //订单号
	GoodsThumbUrl    string `json:"goods_thumb_url"`   //商品缩略图
	GoodsID          int64  `json:"goods_id"`          //商品id
	GoodsName        string `json:"goods_name"`        //商品名
	Money            int    `json:"money"`             //奖励金额
	PayTime          int64  `json:"pay_time"`          //结算时间
	UpdatedAt        int64  `json:"updated_at"`        //更新时间
	VerifyTime       int64  `json:"verify_time"`       //审核时间
	Status           int    `json:"status"`            //状态：0-未审核、1-审核通过、2-审核失败、3-已结算
	CustomParameters string `json:"custom_parameters"` //自定义参数
}

func (this *DuoduoKe) AppNewBillListGet(p *AppNewBillListGetParams) (*AppNewBillListGetResult, error) {
	apiType := `pdd.ddk.app.new.bill.list.get`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result AppNewBillListGetResult
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
