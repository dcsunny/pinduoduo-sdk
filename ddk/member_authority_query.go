package ddk

import (
	"github.com/dcsunny/pinduoduo-sdk/common"
	"github.com/dcsunny/pinduoduo-sdk/util"
)

type MemberAuthorityQueryParams struct {
	Pid              *string `json:"pid,omitempty"`
	CustomParameters *string `json:"custom_parameters,omitempty"`
}

type MemberAuthorityQueryResult struct {
	AuthorityQueryResponse struct {
		Bind int64 `json:"bind"`
	} `json:"authority_query_response"`
	common.CommonResult
}

func (this *DuoduoKe) MemberAuthorityQuery(p *MemberAuthorityQueryParams) (*MemberAuthorityQueryResult, error) {
	apiType := "pdd.ddk.member.authority.query"
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result MemberAuthorityQueryResult

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
