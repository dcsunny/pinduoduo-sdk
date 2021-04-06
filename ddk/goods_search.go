package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

/**
多多进宝商品查询
*/
type GoodsSearchParams struct {
	Keyword          *string             `json:"keyword,omitempty"`           //商品关键词，与opt_id字段选填一个或全部填写
	OptID            *int64              `json:"opt_id,omitempty"`            //商品标签类目ID，使用pdd.goods.opt.get获取
	Page             *int                `json:"page,omitempty"`              //默认值1，商品分页数
	PageSize         *int                `json:"page_size,omitempty"`         //默认100，每页商品数量
	SortType         *int                `json:"sort_type,omitempty"`         //排序方式:0-综合排序;1-按佣金比率升序;2-按佣金比例降序;3-按价格升序;4-按价格降序;5-按销量升序;6-按销量降序;7-优惠券金额排序升序;8-优惠券金额排序降序;9-券后价升序排序;10-券后价降序排序;11-按照加入多多进宝时间升序;12-按照加入多多进宝时间降序;13-按佣金金额升序排序;14-按佣金金额降序排序;15-店铺描述评分升序;16-店铺描述评分降序;17-店铺物流评分升序;18-店铺物流评分降序;19-店铺服务评分升序;20-店铺服务评分降序;27-描述评分击败同类店铺百分比升序，28-描述评分击败同类店铺百分比降序，29-物流评分击败同类店铺百分比升序，30-物流评分击败同类店铺百分比降序，31-服务评分击败同类店铺百分比升序，32-服务评分击败同类店铺百分比降序
	WithCoupon       *bool               `json:"with_coupon,omitempty"`       //是否只返回优惠券的商品，false返回所有商品，true只返回有优惠券的商品
	RangeList        *[]GoodsSearchRange `json:"range_list,omitempty"`        //范围列表，可选值：[{"range_id":0,"range_from":1,"range_to":1500},{"range_id":1,"range_from":1,"range_to":1500}]
	CatID            *int64              `json:"cat_id,omitempty"`            //商品类目ID，使用pdd.goods.cats.get接口获取
	GoodsIDList      *[]int64            `json:"goods_id_list,omitempty"`     //商品ID列表。例如：[123456,123]，当入参带有goods_id_list字段，将不会以opt_id、 cat_id、keyword维度筛选商品
	ZsDuoID          *int64              `json:"zs_duo_id,omitempty"`         //招商多多客ID
	MerchantType     *int                `json:"merchant_type,omitempty"`     //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店（未传为全部）
	ActivityTags     *[]int              `json:"activity_tags,omitempty"`     //商品活动标记数组，例：[4,7]，4-秒杀，7-百亿补贴，24-品牌高佣，20-行业精选，21-金牌商家，10044-潜力爆品，10475-爆品上新，其他的值请忽略
	CustomParameters *string             `json:"custom_parameters,omitempty"` //自定义参数，为链接打上自定义标签；自定义参数最长限制64个字节；格式为： {"uid":"11111","sid":"22222"} ，其中 uid 用户唯一标识，可自行加密后传入，每个用户仅且对应一个标识，必填； sid 上下文信息标识，例如sessionId等，非必填。该json字符串中也可以加入其他自定义的key
	IsBrandGoods     *bool               `json:"is_brand_goods,omitempty"`    //是否为品牌商品
	ListId           *string             `json:"list_id,omitempty"`           //翻页时建议填写前页返回的list_id值
	MerchantTypeList *int                `json:"merchantTypeList,omitempty"`  //店铺类型数组
	Pid              *string             `json:"pid,omitempty"`               //推广位id
	BlockCats        *[]int64            `json:"blockCats,omitempty"`         //自定义屏蔽一级/二级/三级类目ID，自定义数量不超过20个;使用pdd.goods.cats.get接口获取cat_id
	BlockCatPackages *[]int              `json:"blockCatPackages,omitempty"`  //屏蔽商品类目包：1-拼多多小程序屏蔽类目;2-虚拟类目;3-医疗器械;4-处方药;5-非处方药
}

type GoodsSearchRange struct {
	RangeTo   *int  `json:"range_to,omitempty"`   //如果左区间不限制，range_from传空就行，右区间不限制，range_to传空就行
	RangeFrom *int  `json:"range_from,omitempty"` //如果左区间不限制，range_from传空就行，右区间不限制，range_to传空就行
	RangeID   int64 `json:"range_id"`             //查询维度ID，枚举值如下：0-商品拼团价格区间，1-商品券后价价格区间，2-佣金比例区间，3-优惠券金额区间，4-加入多多进宝时间区间，5-销量区间，6-佣金金额区间，7-店铺描述评分区间，8-店铺物流评分区间，9-店铺服务评分区间
}

type GoodsSearchResult struct {
	GoodsSearchResponse struct {
		GoodsList  []GoodsSearchInfo `json:"goods_list"`  //商品列表
		ListId     string            `json:"list_id"`     //翻页时必填前页返回的list_id值
		SearchID   string            `json:"search_id"`   //搜索id，建议生成推广链接时候填写，提高收益
		TotalCount int               `json:"total_count"` //返回商品总数
	} `json:"goods_search_response"`
	common.CommonResult
}
type GoodsSearchInfo struct {
	ActivityTags                []int    `json:"activity_tags,omitempty"`         //商品活动标记数组，例：[4,7]，4-秒杀，7-百亿补贴，24-品牌高佣，20-行业精选，21-金牌商家，10044-潜力爆品，10475-爆品上新，其他的值请忽略
	HasMallCoupon               bool     `json:"has_mall_coupon"`                 //是否有店铺券
	MallCouponID                int64    `json:"mall_coupon_id"`                  //店铺券id
	MallCouponDiscountPct       int      `json:"mall_coupon_discount_pct"`        //店铺券折扣
	MallCouponMinOrderAmount    int      `json:"mall_coupon_min_order_amount"`    //最小使用金额
	MallCouponMaxDiscountAmount int      `json:"mall_coupon_max_discount_amount"` //最大使用金额
	MallCouponTotalQuantity     int64    `json:"mall_coupon_total_quantity"`      //店铺券总量
	MallCouponRemainQuantity    int64    `json:"mall_coupon_remain_quantity"`     //店铺券余量
	MallCouponStartTime         int64    `json:"mall_coupon_start_time"`          //店铺券开始使用时间
	MallCouponEndTime           int64    `json:"mall_coupon_end_time"`            //店铺券结束使用时间
	CreateAt                    int64    `json:"create_at"`                       //创建时间（unix时间戳）
	GoodsID                     int64    `json:"goods_id"`                        //商品id
	GoodsName                   string   `json:"goods_name"`                      //商品名称
	GoodsDesc                   string   `json:"goods_desc"`                      //商品描述
	GoodsThumbnailUrl           string   `json:"goods_thumbnail_url"`             //商品缩略图
	GoodsImageUrl               string   `json:"goods_image_url"`                 //商品主图
	GoodsGalleryUrls            []string `json:"goods_gallery_urls"`              //商品轮播图
	MinGroupPrice               int      `json:"min_group_price"`                 //最小拼团价（单位为分）
	MinNormalPrice              int      `json:"min_normal_price"`                //最小单买价格（单位为分）
	MallName                    string   `json:"mall_name"`                       //店铺名字
	MerchantType                int      `json:"merchant_type"`                   //店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店
	CategoryID                  int64    `json:"category_id"`                     //商品类目ID，使用pdd.goods.cats.get接口获取
	CategoryName                string   `json:"category_name"`                   //商品类目名
	OptID                       int64    `json:"opt_id"`                          //商品标签ID，使用pdd.goods.opts.get接口获取
	OptName                     string   `json:"opt_name"`                        //商品标签名
	OptIds                      []int64  `json:"opt_ids"`                         //商品标签id
	CatIds                      []int64  `json:"cat_ids"`                         //商品类目id
	MallCps                     int      `json:"mall_cps"`                        //该商品所在店铺是否参与全店推广，0：否，1：是
	HasCoupon                   bool     `json:"has_coupon"`                      //商品是否有优惠券 true-有，false-没有
	CouponMinOrderAmount        int      `json:"coupon_min_order_amount"`         //优惠券门槛价格，单位为分
	CouponDiscount              int      `json:"coupon_discount"`                 //优惠券面额，单位为分
	CouponTotalQuantity         int      `json:"coupon_total_quantity"`           //优惠券总数量
	CouponRemainQuantity        int      `json:"coupon_remain_quantity"`          //优惠券剩余数量
	CouponStartTime             int64    `json:"coupon_start_time"`               //优惠券生效时间，UNIX时间戳
	CouponEndTime               int64    `json:"coupon_end_time"`                 //优惠券失效时间，UNIX时间戳
	PromotionRate               int      `json:"promotion_rate"`                  //佣金比例，千分比
	GoodsEvalCount              int      `json:"goods_eval_count"`                //商品评价数量
	CatID                       int64    `json:"cat_id"`                          //商品类目id
	ActivityType                int      `json:"activity_type"`                   //活动类型，0-无活动;1-秒杀;3-限量折扣;12-限时折扣;13-大促活动;14-名品折扣;15-品牌清仓;16-食品超市;17-一元幸运团;18-爱逛街;19-时尚穿搭;20-男人帮;21-9块9;22-竞价活动;23-榜单活动;24-幸运半价购;25-定金预售;26-幸运人气购;27-特色主题活动;28-断码清仓;29-一元话费;30-电器城;31-每日好店;32-品牌卡;101-大促搜索池;102-大促品类分会场;
	SalesTip                    string   `json:"sales_tip"`                       //	已售卖件数
	DescTxt                     string   `json:"desc_txt"`                        //	描述分
	ServTxt                     string   `json:"serv_txt"`                        //服务分
	LgstTxt                     string   `json:"lgst_txt"`                        //	物流分
	ServiceTags                 []int    `json:"service_tags"`                    //服务标签: 1-全场包邮,2-七天退换,3-退货包运费,4-送货入户并安装,5-送货入户,6-电子发票,7-诚信发货,8-缺重包赔,9-坏果包赔,10-果重保证,11-闪电退款,12-24小时发货,13-48小时发货,14-免税费,15-假一罚十,16-贴心服务,17-顺丰包邮,18-只换不修,19-全国联保,20-分期付款,21-纸质发票,22-上门安装,23-爱心助农,24-极速退款,25-品质保障,26-缺重包退,27-当日发货,1000001-正品发票,1000002-送货入户并安装,2000001-价格保护
	CltCpnBatchSn               string   `json:"clt_cpn_batch_sn"`                //店铺收藏券id
	CltCpnStartTime             int64    `json:"clt_cpn_start_time"`              //店铺收藏券起始时间
	CltCpnEndTime               int64    `json:"clt_cpn_end_time"`                //店铺收藏券截止时间
	CltCpnQuantity              int64    `json:"clt_cpn_quantity"`                //店铺收藏券总量
	CltCpnRemainQuantity        int64    `json:"clt_cpn_remain_quantity"`         //店铺收藏券剩余量
	CltCpnDiscount              int64    `json:"clt_cpn_discount"`                //店铺收藏券面额，单位为分
	CltCpnMinAmt                int64    `json:"clt_cpn_min_amt"`                 //店铺收藏券使用门槛价格，单位为分
	ZsDuoID                     int64    `json:"zs_duo_id"`                       //招商多多客ID
	PredictPromotionRate        int64    `json:"predict_promotion_rate"`          //比价行为预判定佣金，需要用户备案
	SearchID                    string   `json:"search_id"`                       //搜索id，建议生成推广链接时候填写，提高收益
	PlanType                    int      `json:"plan_type"`                       //推广计划类型 3:定向 4:招商
	GoodsSign                   string   `json:"goods_sign"`
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
