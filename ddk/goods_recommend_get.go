package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
运营频道商品查询
*/
type GoodsRecommendResult struct {
	GoodsBasicDetailResponse struct {
		List []GoodsRecommendInfo `json:"list"`
	} `json:"goods_basic_detail_response"`
	common.CommonResult
}
type GoodsRecommendParams struct {
	Offset      *int `json:"offset,omitempty"`       //从多少位置开始请求；默认值 ： 0
	Limit       *int `json:"limit,omitempty"`        //请求数量；默认值 ： 400
	ChannelType *int `json:"channel_type,omitempty"` //频道类型；0, "1.9包邮", 1, "今日爆款", 2, "品牌清仓", 3, "默认商城", 非必填 ,默认是1
}

type GoodsRecommendInfo struct {
	QrCodeImageUrl       string  `json:"qr_code_image_url"`       //二维码主图
	LockEdit             string  `json:"lock_edit"`               //编辑锁定
	Broker               string  `json:"broker"`                  //代理人
	Rank                 string  `json:"rank"`                    //顺序
	SaleNumToday         int     `json:"sale_num_today"`          //今日成交量
	SaleNum24            int     `json:"sale_num24"`              //成交量
	ServPct              float64 `json:"serv_pct"`                //服务评分击败同类店铺百分比
	LgstPct              float64 `json:"lgst_pct"`                //物流评分击败同类店铺百分比
	DescPct              float64 `json:"desc_pct"`                //描述评分击败同类店铺百分比
	AvgServ              int     `json:"avg_serv"`                //服务评分
	AvgLgst              int     `json:"avg_lgst"`                //物流评分
	AvgDesc              int     `json:"avg_desc"`                //描述评分
	ShareDesc            string  `json:"share_desc"`              //分享描述
	CatID                int64   `json:"cat_id"`                  //商品类目id
	GoodsEvalCount       int     `json:"goods_eval_count"`        //商品评价数量
	GoodsEvalScore       float64 `json:"goods_eval_score"`        //商品评价分
	MarketFee            int     `json:"market_fee"`              //市场服务费
	GoodsRate            int     `json:"goods_rate"`              //商品等级
	CouponPrice          int     `json:"coupon_price"`            //优惠券金额 分
	PromotionRate        int     `json:"promotion_rate"`          //佣金比例,千分比
	CouponEndTime        int64   `json:"coupon_end_time"`         //优惠券失效时间,UNIX时间戳
	CouponStartTime      int64   `json:"coupon_start_time"`       //优惠券生效时间,UNIX时间戳
	CouponRemainQuantity int     `json:"coupon_remain_quantity"`  //优惠券剩余数量
	CouponTotalQuantity  int     `json:"coupon_total_quantity"`   //优惠券总数量
	CouponDiscount       int     `json:"coupon_discount"`         //优惠券面额,单位为分
	CouponMinOrderAmount int     `json:"coupon_min_order_amount"` //优惠券门槛价格,单位为分
	CouponID             int64   `json:"coupon_id"`               //优惠券id
	HasCoupon            bool    `json:"has_coupon"`              //商品是否带券,true-带券,false-不带券
	GoodsType            int     `json:"goods_type"`              //商品类型
	CatIds               []int64 `json:"cat_ids"`                 //商品一~四级类目ID列表
	OptIds               []int64 `json:"opt_ids"`                 //商品一~四级标签类目ID列表
	OptName              string  `json:"opt_name"`                //商品标签名
	OptID                int64   `json:"opt_id"`                  //商品标签类目ID,使用pdd.goods.opt.get获取
	CategoryName         string  `json:"category_name"`           //分类名称
	CategoryID           int64   `json:"category_id"`             //类目id
	MerchantType         int     `json:"merchant_type"`           //商家类型
	MallName             string  `json:"mall_name"`               //店铺名称
	MallID               int64   `json:"mall_id"`                 //商家id
	MinNormalPrice       int     `json:"min_normal_price"`        //最小单买价格，单位分
	MinGroupPrice        int     `json:"min_group_price"`         //最小成团价格，单位分
	GoodsFactPrice       int     `json:"goods_fact_price"`        //商品实际价格
	GoodsMarkPrice       int     `json:"goods_mark_price"`        //商品标准价格
	SoldQuantity         int     `json:"sold_quantity"`           //销售量
	GoodsGalleryUrls     string  `json:"goods_gallery_urls"`      //商品详情图列表
	GoodsImageURL        string  `json:"goods_image_url"`         //商品主图
	GoodsThumbnailUrl    string  `json:"goods_thumbnail_url"`     //商品缩略图
	GoodsDesc            string  `json:"goods_desc"`              //商品描述
	GoodsName            string  `json:"goods_name"`              //商品名称
	GoodsID              int64   `json:"goods_id"`                //商品id
	CreateAt             int64     `json:"create_at"`               //创建时间
}

func (this *DuoduoKe) GetGoodsRecommend(p *GoodsRecommendParams) (*GoodsRecommendResult, error) {

	apiType := "pdd.ddk.goods.recommend.get"

	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result GoodsRecommendResult
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
