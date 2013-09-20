package badsoap

import (
	"encoding/xml"
)

type Campaign struct {
	BudgetType                string  `xml:"https://adcenter.microsoft.com/v8 BudgetType"`
	ConversionTrackingEnabled bool    `xml:"https://adcenter.microsoft.com/v8 ConversionTrackingEnabled"`
	DailyBudget               float64 `xml:"https://adcenter.microsoft.com/v8 DailyBudget"`
	DaylightSaving            bool    `xml:"https://adcenter.microsoft.com/v8 DaylightSaving"`
	Description               string  `xml:"https://adcenter.microsoft.com/v8 Description"`
	Id                        int64   `xml:"https://adcenter.microsoft.com/v8 Id,omitempty"`
	MonthlyBudget             float64 `xml:"https://adcenter.microsoft.com/v8 MonthlyBudget"`
	Name                      string  `xml:"https://adcenter.microsoft.com/v8 Name"`
	Status                    string  `xml:"https://adcenter.microsoft.com/v8 Status,omitempty"`
	TimeZone                  string  `xml:"https://adcenter.microsoft.com/v8 TimeZone"`
}

// Unmarshal Return values
func unmarshalCampaigns(xmlBytes []byte) (campaigns []Campaign, err error) {
	opResponse := struct {
		Campaigns []Campaign `xml:"Campaigns>Campaign"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return campaigns, err
	}
	return opResponse.Campaigns, err
}

func unmarshalCampaignIds(xmlBytes []byte) (campaignIds []int64, err error) {
	opResponse := struct {
		CampaignIds []int64 `xml:"CampaignIds>long"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return campaignIds, err
	}
	return opResponse.CampaignIds, err
}

func (a *Auth) AddCampaigns(accountId int64, campaigns []Campaign) (campaignIds []int64, err error) {
	cs := []Campaign{}
	for _, c := range campaigns {
		c.Id = 0
		cs = append(cs, c)
	}
	addCampaignsRequest := struct {
		XMLName   xml.Name   `xml:"https://adcenter.microsoft.com/v8 AddCampaignsRequest"`
		AccountId int64      `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		Campaigns []Campaign `xml:"https://adcenter.microsoft.com/v8 Campaigns>Campaign"`
	}{XMLNS, accountId, cs}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("AddCamapigns"), SoapRequestBody{addCampaignsRequest}}, " ", " ")
	if err != nil {
		return campaignIds, err
	}
	soapRespBody, err := a.Request("AddCampaigns", reqBody)
	if err != nil {
		return campaignIds, err
	}
	return unmarshalCampaignIds(soapRespBody)
}

func (a *Auth) DeleteCampaigns(accountId int64, campaignIds []int64) error {
	deleteCampaignsRequest := struct {
		XMLName     xml.Name `xml:"https://adcenter.microsoft.com/v8 DeleteCampaignsRequest"`
		AccountId   int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		CampaignIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays CampaignIds>long"`
	}{XMLNS, accountId, campaignIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("DeleteCamapigns"), SoapRequestBody{deleteCampaignsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("DeleteCampaigns", reqBody)
	return err
}

func (a *Auth) GetCampaignsByAccountId(accountId int64) (campaigns []Campaign, err error) {
	getCampaignsByAccountIdRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 GetCampaignsByAccountIdRequest"`
		AccountId int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
	}{XMLNS, accountId}
	reqBody, err := xml.Marshal(SoapRequestEnvelope{ENVNS, a.authHeader("GetCampaignsByAccountId"), SoapRequestBody{getCampaignsByAccountIdRequest}})
	if err != nil {
		return campaigns, err
	}
	respBody, err := a.Request("GetCampaignsByAccountId", reqBody)
	if err != nil {
		return campaigns, err
	}
	return unmarshalCampaigns(respBody)
}

func (a *Auth) GetCampaignsByIds(accountId int64, campaignIds []int64) (campaigns []Campaign, err error) {
	getCampaignsByIdsRequest := struct {
		XMLName     xml.Name `xml:"https://adcenter.microsoft.com/v8 GetCampaignsByIdsRequest"`
		AccountId   int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		CampaignIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays CampaignIds>long"`
	}{XMLNS, accountId, campaignIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetCampaignsByIds"), SoapRequestBody{getCampaignsByIdsRequest}}, " ", " ")
	if err != nil {
		return campaigns, err
	}
	respBody, err := a.Request("GetCampaignsByIds", reqBody)
	if err != nil {
		return campaigns, err
	}
	return unmarshalCampaigns(respBody)
}

func (a *Auth) PauseCampaigns(accountId int64, campaignIds []int64) error {
	pauseCampaignsRequest := struct {
		XMLName     xml.Name `xml:"https://adcenter.microsoft.com/v8 PauseCampaignsRequest"`
		AccountId   int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		CampaignIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays CampaignIds>long"`
	}{XMLNS, accountId, campaignIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("PauseCampaigns"), SoapRequestBody{pauseCampaignsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("PauseCampaigns", reqBody)
	return err
}

func (a *Auth) ResumeCampaigns(accountId int64, campaignIds []int64) error {
	resumeCampaignsRequest := struct {
		XMLName     xml.Name `xml:"https://adcenter.microsoft.com/v8 ResumeCampaignsRequest"`
		AccountId   int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		CampaignIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays CampaignIds>long"`
	}{XMLNS, accountId, campaignIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("ResumeCampaigns"), SoapRequestBody{resumeCampaignsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("ResumeCampaigns", reqBody)
	return err
}

func (a *Auth) UpdateCampaigns(accountId int64, campaigns []Campaign) error {
	cs := []Campaign{}
	for _, c := range campaigns {
		c.Status = ""
		cs = append(cs, c)
	}
	updateCampaignsRequest := struct {
		XMLName   xml.Name   `xml:"https://adcenter.microsoft.com/v8 UpdateCampaignsRequest"`
		AccountId int64      `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		Campaigns []Campaign `xml:"https://adcenter.microsoft.com/v8 Campaigns>Campaign"`
	}{XMLNS, accountId, cs}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("UpdateCampaigns"), SoapRequestBody{updateCampaignsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("UpdateCampaigns", reqBody)
	return err
}
