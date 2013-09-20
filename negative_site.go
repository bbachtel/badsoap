package badsoap

import (
	"encoding/xml"
)

func (a *Auth) GetNegativeSitesByAdGroupIds(campaignId int64, adGroupIds []int64) (adGroupIdsNegativeSites map[int64][]string, err error) {
	getNegativeSitesByAdGroupIdsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetNegativeSitesByAdGroupIdsRequest"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroupIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdGroupIds>long"`
	}{XMLNS, campaignId, adGroupIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetNegativeSitesByAdGroupIds"), SoapRequestBody{getNegativeSitesByAdGroupIdsRequest}}, " ", " ")
	if err != nil {
		return adGroupIdsNegativeSites, err
	}
	soapRespBody, err := a.Request("GetNegativeSitesByAdGroupIds", reqBody)
	if err != nil {
		return adGroupIdsNegativeSites, err
	}
	type unmarshalAdGroupNegativeSites struct {
		AdGroupId     int64    `xml:"AdGroupId"`
		NegativeSites []string `xml:"NegativeSites>string"`
	}
	opResponse := struct {
		AdGroupNegativeSites []unmarshalAdGroupNegativeSites `xml:"AdGroupNegativeSites>AdGroupNegativeSites"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return adGroupIdsNegativeSites, err
	}
	adGroupIdsNegativeSites = map[int64][]string{}
	for _, ans := range opResponse.AdGroupNegativeSites {
		adGroupIdsNegativeSites[ans.AdGroupId] = ans.NegativeSites
	}
	return adGroupIdsNegativeSites, err
}

func (a *Auth) SetNegativeSitesToAdGroups(campaignId int64, adGroupIdsNegativeSites map[int64][]string) error {
	type negativeSite struct {
		XMLName       xml.Name `xml:"https://adcenter.microsoft.com/v8 AdGroupNegativeSites"`
		AdGroupId     int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		NegativeSites []string `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays NegativeSites>string"`
	}
	allNegativeSites := []negativeSite{}
	for adGroupId, negativeSites := range adGroupIdsNegativeSites {
		allNegativeSites = append(allNegativeSites, negativeSite{XMLNS, adGroupId, negativeSites})
	}
	setNegativeSitesToAdGroupsRequest := struct {
		XMLName       xml.Name       `xml:"https://adcenter.microsoft.com/v8 SetNegativeSitesToAdGroupsRequest"`
		CampaignId    int64          `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		NegativeSites []negativeSite `xml:"https://adcenter.microsoft.com/v8 AdGroupNegativeSites>AdGroupNegativeSites"`
	}{XMLNS, campaignId, allNegativeSites}
	requestXml := SoapRequestEnvelope{ENVNS, a.authHeader("SetNegativeSitesToAdGroups"), SoapRequestBody{setNegativeSitesToAdGroupsRequest}}
	reqBody, err := xml.MarshalIndent(requestXml, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetNegativeSitesToAdGroups", reqBody)
	return err
}

func (a *Auth) GetNegativeSitesByCampaignIds(accountId int64, campaignIds []int64) (campaignIdsNegativeSites map[int64][]string, err error) {
	getNegativeSitesByCampaignIdsRequest := struct {
		XMLName     xml.Name `xml:"https://adcenter.microsoft.com/v8 GetNegativeSitesByCampaignIdsRequest"`
		AccountId   int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		CampaignIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays CampaignIds>long"`
	}{XMLNS, accountId, campaignIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetNegativeSitesByCampaignIds"), SoapRequestBody{getNegativeSitesByCampaignIdsRequest}}, " ", " ")
	if err != nil {
		return campaignIdsNegativeSites, err
	}
	soapRespBody, err := a.Request("GetNegativeSitesByCampaignIds", reqBody)
	if err != nil {
		return campaignIdsNegativeSites, err
	}
	type unmarshalCampaignNegativeSites struct {
		CampaignId    int64    `xml:"CampaignId"`
		NegativeSites []string `xml:"NegativeSites>string"`
	}
	opResponse := struct {
		CampaignNegativeSites []unmarshalCampaignNegativeSites `xml:"CampaignNegativeSites>CampaignNegativeSites"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return campaignIdsNegativeSites, err
	}
	campaignIdsNegativeSites = map[int64][]string{}
	for _, a := range opResponse.CampaignNegativeSites {
		campaignIdsNegativeSites[a.CampaignId] = a.NegativeSites
	}
	return campaignIdsNegativeSites, err
}

func (a *Auth) SetNegativeSitesToCampaigns(accountId int64, campaignNegativeSites map[int64][]string) error {
	type negativeSite struct {
		XMLName       xml.Name `xml:"CampaignNegativeSites"`
		CampaignId    int64    `xml:"CampaignId"`
		NegativeSites []string `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays NegativeSites>string"`
	}
	allNegativeSites := []negativeSite{}
	for campaignId, negativeSites := range campaignNegativeSites {
		allNegativeSites = append(allNegativeSites, negativeSite{XMLNS, campaignId, negativeSites})
	}
	setNegativeSitesToCampaignsRequest := struct {
		XMLName   xml.Name       `xml:"https://adcenter.microsoft.com/v8 SetNegativeSitesToCampaignsRequest"`
		AccountId int64          `xml:"AccountId"`
		Campaigns []negativeSite `xml:"CampaignNegativeSites>CampaignNegativeSites"`
	}{XMLNS, accountId, allNegativeSites}
	requestXml := SoapRequestEnvelope{ENVNS, a.authHeader("SetNegativeSitesToCampaigns"), SoapRequestBody{setNegativeSitesToCampaignsRequest}}
	reqBody, err := xml.MarshalIndent(requestXml, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetNegativeSitesToCampaigns", reqBody)
	return err
}
