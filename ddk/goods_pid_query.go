package ddk

import "github.com/dcsunny/pinduoduo-sdk/util"

/**
查询已经生成的推广位信息
*/
type GoodsPidQueryParams struct {
	Page     *int `json:"page,omitempty"`      //返回的页数
	PageSize *int `json:"page_size,omitempty"` //返回的每页推广位数量
}

type GoodsPidQueryResult struct {
	PIdQueryResponse struct {
		PIdList    []GoodsPidQueryInfo `json:"p_id_list"`   //多多进宝推广位对象列表
		TotalCount int                 `json:"total_count"` //返回推广位总数
	} `json:"p_id_query_response"`
	util.CommonResult
}

type GoodsPidQueryInfo struct {
	Name       string `json:"name"`        //推广位名称
	PID        string `json:"p_id"`        //推广位ID
	CreateTime int64  `json:"create_time"` //推广位生成时间
}

func (this *DuoduoKe) GoodsPidQuery(p *GoodsPidQueryParams) (*GoodsPidQueryResult, error) {
	apiType := `pdd.ddk.goods.pid.query`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result GoodsPidQueryResult
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
