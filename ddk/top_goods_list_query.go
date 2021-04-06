package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

type TopGoodsListQueryParams struct {
	Pid      string `json:"p_id,omitempty"`
	SortType int    `json:"sort_type"`
	Offset   int    `json:"offset,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type TopGoodsListQueryResult struct {
	TopGoodsListGetResponse struct {
		List     []TopGoodsListQueryInfo `json:"list"`
		ListID   string                  `json:"list_id"`
		SearchID string                  `json:"search_id"`
		Total    int                     `json:"total"`
	} `json:"top_goods_list_get_response"`
	common.CommonResult
}

type TopGoodsListQueryInfo struct {
	GoodsID              int64    `json:"goods_id"`
	GoodsName            string   `json:"goods_name"`
	GoodsDesc            string   `json:"goods_desc"`
	GoodsThumbnailUrl    string   `json:"goods_thumbnail_url"`
	GoodsImageUrl        string   `json:"goods_image_url"`
	GoodsGalleryUrls     []string `json:"goods_gallery_urls"`
	MinGroupPrice        int      `json:"min_group_price"`
	MinNormalPrice       int      `json:"min_normal_price"`
	MallName             string   `json:"mall_name"`
	MerchantType         int      `json:"merchant_type"`
	CategoryId           int      `json:"category_id"`
	CategoryName         string   `json:"category_name"`
	OptId                int      `json:"opt_id"`
	OptName              string   `json:"opt_name"`
	OptIds               []int    `json:"opt_ids"`
	CatIds               []int    `json:"cat_ids"`
	MallCps              int      `json:"mall_cps"`
	HasCoupon            bool     `json:"has_coupon"`
	CouponMinOrderAmount int      `json:"coupon_min_order_amount"`
	CouponDiscount       int      `json:"coupon_discount"`
	CouponTotalQuantity  int      `json:"coupon_total_quantity"`
	CouponRemainQuantity int      `json:"coupon_remain_quantity"`
	CouponStartTime      int64    `json:"coupon_start_time"`
	CouponEndTime        int64    `json:"coupon_end_time"`
	PromotionRate        int      `json:"promotion_rate"`
	GoodsEvalCount       int      `json:"goods_eval_count"`
	SalesTip             string   `json:"sales_tip"`
	DescTxt              string   `json:"desc_txt"`
	ServTxt              string   `json:"serv_txt"`
	LgstTxt              string   `json:"lgst_txt"`
	SearchID             string   `json:"search_id"`
	GoodsSign            string   `json:"goods_sign"`
}

func (this *DuoduoKe) TopGoodsListQuery(p *TopGoodsListQueryParams) (*TopGoodsListQueryResult, error) {
	apiType := `pdd.ddk.top.goods.list.query`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result TopGoodsListQueryResult
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
