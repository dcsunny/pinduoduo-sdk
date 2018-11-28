package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多进宝商品查询
*/
type GoodsSearchParams struct {
	Keyword      *string             `json:"keyword,omitempty"`       //商品关键词，与opt_id字段选填一个或全部填写
	OptID        *int64              `json:"opt_id,omitempty"`        //商品标签类目ID，使用pdd.goods.opt.get获取
	Page         *int                `json:"page,omitempty"`          //默认值1，商品分页数
	PageSize     *int                `json:"page_size,omitempty"`     //默认100，每页商品数量
	SortType     int                 `json:"sort_type"`               //排序方式:0-综合排序;1-按佣金比率升序;2-按佣金比例降序;3-按价格升序;4-按价格降序;5-按销量升序;6-按销量降序;7-优惠券金额排序升序;8-优惠券金额排序降序;9-券后价升序排序;10-券后价降序排序;11-按照加入多多进宝时间升序;12-按照加入多多进宝时间降序;13-按佣金金额升序排序;14-按佣金金额降序排序;15-店铺描述评分升序;16-店铺描述评分降序;17-店铺物流评分升序;18-店铺物流评分降序;19-店铺服务评分升序;20-店铺服务评分降序;27-描述评分击败同类店铺百分比升序，28-描述评分击败同类店铺百分比降序，29-物流评分击败同类店铺百分比升序，30-物流评分击败同类店铺百分比降序，31-服务评分击败同类店铺百分比升序，32-服务评分击败同类店铺百分比降序
	WithCoupon   bool                `json:"with_coupon"`             //是否只返回优惠券的商品，false返回所有商品，true只返回有优惠券的商品
	RangeList    *[]GoodsSearchRange `json:"range_list,omitempty"`    //范围列表，可选值：[{"range_id":0,"range_from":1,"range_to":1500},{"range_id":1,"range_from":1,"range_to":1500}]
	CatID        *int64              `json:"cat_id,omitempty"`        //商品类目ID，使用pdd.goods.cats.get接口获取
	GoodsIDList  *[]int64            `json:"goods_id_list,omitempty"` //商品ID列表。例如：[123456,123]，当入参带有goods_id_list字段，将不会以opt_id、 cat_id、keyword维度筛选商品
	ZsDuoID      *int64              `json:"zs_duo_id,omitempty"`     //招商多多客ID
	MerchantType *int                `json:"merchant_type,omitempty"` //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店（未传为全部）
}

type GoodsSearchRange struct {
	RangeTo   *int  `json:"range_to,omitempty"`   //如果左区间不限制，range_from传空就行，右区间不限制，range_to传空就行
	RangeFrom *int  `json:"range_from,omitempty"` //如果左区间不限制，range_from传空就行，右区间不限制，range_to传空就行
	RangeID   int64 `json:"range_id"`             //查询维度ID，枚举值如下：0-商品拼团价格区间，1-商品券后价价格区间，2-佣金比例区间，3-优惠券金额区间，4-加入多多进宝时间区间，5-销量区间，6-佣金金额区间，7-店铺描述评分区间，8-店铺物流评分区间，9-店铺服务评分区间
}

type GoodsSearchResult struct {
	GoodsSearchResponse struct {
		GoodsList  []GoodsSearchInfo `json:"goods_list"`  //商品列表
		TotalCount int               `json:"total_count"` //返回商品总数
	} `json:"goods_search_response"`
	common.CommonResult
}
type GoodsSearchInfo struct {
	CreateAt             int64    `json:"create_at"`               //创建时间（unix时间戳）
	GoodsID              int64    `json:"goods_id"`                //商品id
	GoodsName            string   `json:"goods_name"`              //商品名称
	GoodsDesc            string   `json:"goods_desc"`              //商品描述
	GoodsThumbnailUrl    string   `json:"goods_thumbnail_url"`     //商品缩略图
	GoodsImageUrl        string   `json:"goods_image_url"`         //商品主图
	GoodsGalleryUrls     []string `json:"goods_gallery_urls"`      //商品轮播图
	SoldQuantity         int      `json:"sold_quantity"`           //已售卖件数
	MinGroupPrice        int      `json:"min_group_price"`         //最小拼团价（单位为分）
	MinNormalPrice       int      `json:"min_normal_price"`        //最小单买价格（单位为分）
	MallName             string   `json:"mall_name"`               //店铺名字
	MerchantType         int      `json:"merchant_type"`           //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
	CategoryID           int64    `json:"category_id"`             //商品类目ID，使用pdd.goods.cats.get接口获取
	CategoryName         string   `json:"category_name"`           //商品类目名
	OptID                int64    `json:"opt_id"`                  //商品标签ID，使用pdd.goods.opts.get接口获取
	OptName              string   `json:"opt_name"`                //商品标签名
	OptIds               []int64  `json:"opt_ids"`                 //商品标签id
	CatIds               []int64  `json:"cat_ids"`                 //商品类目id
	MallCps              int      `json:"mall_cps"`                //该商品所在店铺是否参与全店推广，0：否，1：是
	HasCoupon            bool     `json:"has_coupon"`              //商品是否有优惠券 true-有，false-没有
	CouponMinOrderAmount int      `json:"coupon_min_order_amount"` //优惠券门槛价格，单位为分
	CouponDiscount       int      `json:"coupon_discount"`         //优惠券面额，单位为分
	CouponTotalQuantity  int      `json:"coupon_total_quantity"`   //优惠券总数量
	CouponRemainQuantity int      `json:"coupon_remain_quantity"`  //优惠券剩余数量
	CouponStartTime      int64    `json:"coupon_start_time"`       //优惠券生效时间，UNIX时间戳
	CouponEndTime        int64    `json:"coupon_end_time"`         //优惠券失效时间，UNIX时间戳
	PromotionRate        int      `json:"promotion_rate"`          //佣金比例，千分比
	GoodsEvalScore       float64  `json:"goods_eval_score"`        //商品评价分
	GoodsEvalCount       int      `json:"goods_eval_count"`        //商品评价数量
	CatID                int64    `json:"cat_id"`                  //描述评分
	AvgDesc              int      `json:"avg_desc"`                //物流评分
	AvgLgst              int      `json:"avg_lgst"`                //服务评分
	AvgServ              int      `json:"avg_serv"`                //描述分击败同类店铺百分比
	DescPct              float64  `json:"desc_pct"`                //物流分击败同类店铺百分比
	LgstPct              float64  `json:"lgst_pct"`                //服务分击败同类店铺百分比
	ServPct              float64  `json:"serv_pct"`
}

func (this *DuoduoKe) GoodsSearch(p *GoodsSearchParams) (*GoodsSearchResult, error) {
	apiType := "pdd.ddk.goods.search"
	params, paramsURL := util.FormatURLParams(p)

	url := this.GetURL(apiType, "", params, paramsURL)

	var result GoodsSearchResult

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
