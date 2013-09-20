package badsoap

import (
	"encoding/xml"
)

type PublisherCountry struct {
	Country   string `xml:"https://adcenter.microsoft.com/v8 Country"`
	IsOptedIn bool   `xml:"https://adcenter.microsoft.com/v8 IsOptedIn"`
}

// AdDistribution: "Search", "Content"
// BiddingModel: "Keyword", "SitePlacement"
// Network: "OwnedAndOperatedAndSyndicatedSearch",
//          "OwnedAndOperatedOnly", "SyndicatedSearchOnly"
// PricingModel: "Cpc","Cpm"
// Status: "Draft","Active","Paused","Deleted"
type AdGroup struct {
	AdDistribution     string             `xml:"https://adcenter.microsoft.com/v8 AdDistribution"`
	BiddingModel       string             `xml:"https://adcenter.microsoft.com/v8 BiddingModel,omitempty"`
	BroadMatchBid      float64            `xml:"https://adcenter.microsoft.com/v8 BroadMatchBid>Amount,omitempty"`
	ContentMatchBid    float64            `xml:"https://adcenter.microsoft.com/v8 ContentMatchBid>Amount,omitempty"`
	EndDate            *Date              `xml:"https://adcenter.microsoft.com/v8 EndDate,omitempty"`
	ExactMatchBid      float64            `xml:"https://adcenter.microsoft.com/v8 ExactMatchBid>Amount,omitempty"`
	Id                 int64              `xml:"https://adcenter.microsoft.com/v8 Id,omitempty"`
	Language           string             `xml:"https://adcenter.microsoft.com/v8 Language,omitempty"`
	Name               string             `xml:"https://adcenter.microsoft.com/v8 Name"`
	Network            string             `xml:"https://adcenter.microsoft.com/v8 Network"`
	PhraseMatchBid     float64            `xml:"https://adcenter.microsoft.com/v8 PhraseMatchBid>Amount,omitempty"`
	PricingModel       string             `xml:"https://adcenter.microsoft.com/v8 PricingModel"`
	PublisherCountries []PublisherCountry `xml:"https://adcenter.microsoft.com/v8 PublisherCountries"`
	StartDate          *Date              `xml:"https://adcenter.microsoft.com/v8 StartDate"`
	Status             string             `xml:"https://adcenter.microsoft.com/v8 Status,omitempty"`
}

func unmarshalAdGroups(xmlBytes []byte) (adGroups []AdGroup, err error) {
	opResponse := struct {
		AdGroups []AdGroup `xml:"AdGroups>AdGroup"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return adGroups, err
	}
	// nil out empty end date
	emptyDate := Date{}
	for i, _ := range opResponse.AdGroups {
		if *opResponse.AdGroups[i].EndDate == emptyDate {
			opResponse.AdGroups[i].EndDate = nil
		}
	}
	return opResponse.AdGroups, err
}

func unmarshalAdGroupIds(xmlBytes []byte) (adGroupIds []int64, err error) {
	opResponse := struct {
		AdGroupIds []int64 `xml:"AdGroupIds>long"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return adGroupIds, err
	}
	return opResponse.AdGroupIds, err
}

func (a *Auth) AddAdGroups(campaignId int64, adGroups []AdGroup) (adGroupIds []int64, err error) {
	ags := []AdGroup{}
	for _, ag := range adGroups {
		ag.Id = 0
		ags = append(ags, ag)
	}
	addAdGroupsRequest := struct {
		XMLName    xml.Name  `xml:"https://adcenter.microsoft.com/v8 AddAdGroupsRequest"`
		CampaignId int64     `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroups   []AdGroup `xml:"https://adcenter.microsoft.com/v8 AdGroups>AdGroup"`
	}{XMLNS, campaignId, ags}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("AddAdGroups"), SoapRequestBody{addAdGroupsRequest}}, " ", " ")
	if err != nil {
		return adGroupIds, err
	}
	soapRespBody, err := a.Request("AddAdGroups", reqBody)
	if err != nil {
		return adGroupIds, err
	}
	return unmarshalAdGroupIds(soapRespBody)
}

func (a *Auth) DeleteAdGroups(campaignId int64, adGroupIds []int64) error {
	deleteAdGroupsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 DeleteAdGroupsRequest"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroupIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdGroupIds>long"`
	}{XMLNS, campaignId, adGroupIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("DeleteAdGroups"), SoapRequestBody{deleteAdGroupsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("DeleteAdGroups", reqBody)
	return err
}

func (a *Auth) GetAdGroupsByCampaignId(campaignId int64) (adGroups []AdGroup, err error) {
	getAdGroupsByCampaignIdRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdGroupsByCampaignIdRequest"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
	}{XMLNS, campaignId}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdGroupsByCampaignId"), SoapRequestBody{getAdGroupsByCampaignIdRequest}}, " ", " ")
	if err != nil {
		return adGroups, err
	}
	respBody, err := a.Request("GetAdGroupsByCampaignId", reqBody)
	if err != nil {
		return adGroups, err
	}
	return unmarshalAdGroups(respBody)
}

func (a *Auth) GetAdGroupsByIds(campaignId int64, adGroupIds []int64) (adGroups []AdGroup, err error) {
	getAdGroupsByIdsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdGroupsByIdsRequest"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroupIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdGroupIds>long"`
	}{XMLNS, campaignId, adGroupIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdGroupsByIds"), SoapRequestBody{getAdGroupsByIdsRequest}}, " ", " ")
	if err != nil {
		return adGroups, err
	}
	respBody, err := a.Request("GetAdGroupsByIds", reqBody)
	if err != nil {
		return adGroups, err
	}
	return unmarshalAdGroups(respBody)
}

func (a *Auth) PauseAdGroups(campaignId int64, adGroupIds []int64) error {
	pauseAdGroupsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 PauseAdGroupsRequest"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroupIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdGroupIds>long"`
	}{XMLNS, campaignId, adGroupIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("PauseAdGroups"), SoapRequestBody{pauseAdGroupsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("PauseAdGroups", reqBody)
	return err
}

func (a *Auth) ResumeAdGroups(campaignId int64, adGroupIds []int64) error {
	resumeAdGroupsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 ResumeAdGroupsRequest"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroupIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdGroupIds>long"`
	}{XMLNS, campaignId, adGroupIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("ResumeAdGroups"), SoapRequestBody{resumeAdGroupsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("ResumeAdGroups", reqBody)
	return err
}

func (a *Auth) SubmitAdGroupForApproval(adGroupId int64) error {
	submitAdGroupForApprovalRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 SubmitAdGroupForApprovalRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
	}{XMLNS, adGroupId}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("SubmitAdGroupForApproval"), SoapRequestBody{submitAdGroupForApprovalRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SubmitAdGroupForApproval", reqBody)
	return err
}

func (a *Auth) UpdateAdGroups(campaignId int64, adGroups []AdGroup) error {
	ags := []AdGroup{}
	for _, ag := range adGroups {
		ag.BiddingModel = ""
		ag.Language = ""
		ag.Status = ""
		ags = append(ags, ag)
	}
	updateAdGroupsRequest := struct {
		XMLName    xml.Name  `xml:"https://adcenter.microsoft.com/v8 UpdateAdGroupsRequest"`
		CampaignId int64     `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
		AdGroups   []AdGroup `xml:"https://adcenter.microsoft.com/v8 AdGroups>AdGroup"`
	}{XMLNS, campaignId, ags}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("UpdateAdGroups"), SoapRequestBody{updateAdGroupsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("UpdateAdGroups", reqBody)
	return err
}
