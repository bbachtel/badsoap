package badsoap

import (
	"encoding/xml"
)

func (a *Auth) GetNegativeKeywordsByCampaignIds(accountId int64, campaignIds []int64) (campaignIdsNegativeKeywords map[int64][]string, err error) {
	getNegativeKeywordsByCampaignIdsRequest := struct {
		XMLName     xml.Name `xml:"https://adcenter.microsoft.com/v8 GetNegativeKeywordsByCampaignIdsRequest"`
		AccountId   int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		CampaignIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays CampaignIds>long"`
	}{XMLNS, accountId, campaignIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetNegativeKeywordsByCampaignIds"), SoapRequestBody{getNegativeKeywordsByCampaignIdsRequest}}, " ", " ")
	if err != nil {
		return campaignIdsNegativeKeywords, err
	}
	soapRespBody, err := a.Request("GetNegativeKeywordsByCampaignIds", reqBody)
	if err != nil {
		return campaignIdsNegativeKeywords, err
	}
	type unmarshalCampaignNegativeKeywords struct {
		CampaignId       int64    `xml:"CampaignId"`
		NegativeKeywords []string `xml:"NegativeKeywords>string"`
	}
	opResponse := struct {
		CampaignNegativeKeywords []unmarshalCampaignNegativeKeywords `xml:"CampaignNegativeKeywords>CampaignNegativeKeywords"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return campaignIdsNegativeKeywords, err
	}
	campaignIdsNegativeKeywords = map[int64][]string{}
	for _, ans := range opResponse.CampaignNegativeKeywords {
		campaignIdsNegativeKeywords[ans.CampaignId] = ans.NegativeKeywords
	}
	return campaignIdsNegativeKeywords, err
}

func (a *Auth) SetNegativeKeywordsToCampaigns(accountId int64, campaignIdsNegativeKeywords map[int64][]string) error {
	type negativeKeyword struct {
		XMLName          xml.Name `xml:"CampaignNegativeKeywords"`
		CampaignId       int64
		NegativeKeywords []string `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays NegativeKeywords>string"`
	}
	allNegativeKeywords := []negativeKeyword{}
	for campaignId, negativeKeywords := range campaignIdsNegativeKeywords {
		allNegativeKeywords = append(allNegativeKeywords, negativeKeyword{XMLNS, campaignId, negativeKeywords})
	}
	setNegativeKeywordsToCampaignsRequest := struct {
		XMLName   xml.Name          `xml:"https://adcenter.microsoft.com/v8 SetNegativeKeywordsToCampaignsRequest"`
		AccountId int64             `xml:"AccountId"`
		Campaigns []negativeKeyword `xml:"CampaignNegativeKeywords>CampaignNegativeKeywords"`
	}{XMLNS, accountId, allNegativeKeywords}
	requestXml := SoapRequestEnvelope{ENVNS, a.authHeader("SetNegativeKeywordsToCampaigns"), SoapRequestBody{setNegativeKeywordsToCampaignsRequest}}
	reqBody, err := xml.MarshalIndent(requestXml, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetNegativeKeywordsToCampaigns", reqBody)
	return err
}

func (a *Auth) GetNegativeKeywordsByAdGroupIds(campaignId int64, adGroupIds []int64) (adGroupIdsNegativeKeywords map[int64][]string, err error) {
	getNegativeKeywordsByAdGroupIdsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetNegativeKeywordsByAdGroupIdsRequest"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroupIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdGroupIds>long"`
	}{XMLNS, campaignId, adGroupIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetNegativeKeywordsByAdGroupIds"), SoapRequestBody{getNegativeKeywordsByAdGroupIdsRequest}}, " ", " ")
	if err != nil {
		return adGroupIdsNegativeKeywords, err
	}
	soapRespBody, err := a.Request("GetNegativeKeywordsByAdGroupIds", reqBody)
	if err != nil {
		return adGroupIdsNegativeKeywords, err
	}
	type unmarshalAdGroupNegativeKeywords struct {
		AdGroupId        int64    `xml:"AdGroupId"`
		NegativeKeywords []string `xml:"NegativeKeywords>string"`
	}
	opResponse := struct {
		AdGroupNegativeKeywords []unmarshalAdGroupNegativeKeywords `xml:"AdGroupNegativeKeywords>AdGroupNegativeKeywords"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return adGroupIdsNegativeKeywords, err
	}
	adGroupIdsNegativeKeywords = map[int64][]string{}
	for _, ans := range opResponse.AdGroupNegativeKeywords {
		adGroupIdsNegativeKeywords[ans.AdGroupId] = ans.NegativeKeywords
	}
	return adGroupIdsNegativeKeywords, err
}

func (a *Auth) SetNegativeKeywordsToAdGroups(campaignId int64, adGroupIdsNegativeKeywords map[int64][]string) error {
	type negativeKeyword struct {
		XMLName          xml.Name `xml:"AdGroupNegativeKeywords"`
		AdGroupId        int64
		NegativeKeywords []string `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays NegativeKeywords>string"`
	}
	allNegativeKeywords := []negativeKeyword{}
	for adGroupId, negativeKeywords := range adGroupIdsNegativeKeywords {
		allNegativeKeywords = append(allNegativeKeywords, negativeKeyword{XMLNS, adGroupId, negativeKeywords})
	}
	setNegativeKeywordsToAdGroupsRequest := struct {
		XMLName    xml.Name          `xml:"https://adcenter.microsoft.com/v8 SetNegativeKeywordsToAdGroupsRequest"`
		CampaignId int64             `xml:"CampaignId"`
		AdGroups   []negativeKeyword `xml:"AdGroupNegativeKeywords>AdGroupNegativeKeywords"`
	}{XMLNS, campaignId, allNegativeKeywords}
	requestXml := SoapRequestEnvelope{ENVNS, a.authHeader("SetNegativeKeywordsToAdGroups"), SoapRequestBody{setNegativeKeywordsToAdGroupsRequest}}
	reqBody, err := xml.MarshalIndent(requestXml, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetNegativeKeywordsToAdGroups", reqBody)
	return err
}
