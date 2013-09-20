package badsoap

import (
	"encoding/xml"
)

// AdType: "MobileAd","TextAd","ProductAd"
// EditorialStatus: "Active","Disapproved","Inactive"
// Status: "Active","Deleted","Inactive","Paused"
// Type: "Image","Mobile","Product","RichSearch","Text"
type Ad struct {
	XMLName         xml.Name `xml:"https://adcenter.microsoft.com/v8 Ad"`
	AdType          string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	EditorialStatus string   `xml:"https://adcenter.microsoft.com/v8 EditorialStatus,omitempty"`
	Id              int64    `xml:"https://adcenter.microsoft.com/v8 Id,omitempty"`
	Status          string   `xml:"https://adcenter.microsoft.com/v8 Status,omitempty"`
	Type            string   `xml:"https://adcenter.microsoft.com/v8 Type"`
	BusinessName    string   `xml:"https://adcenter.microsoft.com/v8 BusinessName,omitempty"`
	DestinationUrl  string   `xml:"https://adcenter.microsoft.com/v8 DestinationUrl,omitempty"`
	DisplayUrl      string   `xml:"https://adcenter.microsoft.com/v8 DisplayUrl,omitempty"`
	PhoneNumber     string   `xml:"https://adcenter.microsoft.com/v8 PhoneNumber,omitempty"`
	Text            string   `xml:"https://adcenter.microsoft.com/v8 Text,omitempty"`
	Title           string   `xml:"https://adcenter.microsoft.com/v8 Title,omitempty"`
	PromotionalText string   `xml:"https://adcenter.microsoft.com/v8 PromotionalText,omitempty"`
}

func NewTextAd(title, text, destinationUrl, displayUrl, status string) Ad {
	return Ad{
		AdType:         "TextAd",
		Type:           "Text",
		Title:          title,
		Text:           text,
		DestinationUrl: destinationUrl,
		DisplayUrl:     displayUrl,
		Status:         status,
	}
}

func NewMobileAd(title, text, businessName, destinationUrl, displayUrl, phoneNumber, status string) Ad {
	return Ad{
		AdType:         "MobileAd",
		Type:           "Mobile",
		Title:          title,
		Text:           text,
		BusinessName:   businessName,
		DestinationUrl: destinationUrl,
		DisplayUrl:     displayUrl,
		PhoneNumber:    phoneNumber,
		Status:         status,
	}
}

func NewProductAd(promotionalText, status string) Ad {
	return Ad{
		AdType:          "ProductAd",
		Type:            "Product",
		PromotionalText: promotionalText,
		Status:          status,
	}
}

func unmarshalAds(xmlBytes []byte) (ads []Ad, err error) {
	opResponse := struct {
		Ads []Ad `xml:"Ads>Ad"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return ads, err
	}
	return opResponse.Ads, err
}

func unmarshalAdIds(xmlBytes []byte) (adIds []int64, err error) {
	opResponse := struct {
		AdIds []int64 `xml:"AdIds>long"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return adIds, err
	}
	return opResponse.AdIds, err
}

func (a *Auth) AddAds(adGroupId int64, ads []Ad) (adIds []int64, err error) {
	as := []Ad{}
	for _, a := range ads {
		a.EditorialStatus = ""
		a.Id = 0
		as = append(as, a)
	}
	addAdsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 AddAdsRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		Ads       []Ad     `xml:"https://adcenter.microsoft.com/v8 Ads>Ad"`
	}{XMLNS, adGroupId, as}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("AddAds"), SoapRequestBody{addAdsRequest}}, " ", " ")
	if err != nil {
		return adIds, err
	}
	soapRespBody, err := a.Request("AddAds", reqBody)
	if err != nil {
		return adIds, err
	}
	return unmarshalAdIds(soapRespBody)
}

func (a *Auth) DeleteAds(adGroupId int64, adIds []int64) error {
	deleteAdsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 DeleteAdsRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		AdIds     []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdIds>long"`
	}{XMLNS, adGroupId, adIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("DeleteAds"), SoapRequestBody{deleteAdsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("DeleteAds", reqBody)
	return err
}

func (a *Auth) GetAdsByAdGroupId(adGroupId int64) (ads []Ad, err error) {
	getAdsByAdGroupIdRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdsByAdGroupIdRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
	}{XMLNS, adGroupId}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdsByAdGroupId"), SoapRequestBody{getAdsByAdGroupIdRequest}}, " ", " ")
	if err != nil {
		return ads, err
	}
	soapRespBody, err := a.Request("GetAdsByAdGroupId", reqBody)
	if err != nil {
		return ads, err
	}
	return unmarshalAds(soapRespBody)
}

func (a *Auth) GetAdsByEditorialStatus(adGroupId int64, editorialStatus string) (ads []Ad, err error) {
	getAdsByEditorialStatusRequest := struct {
		XMLName         xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdsByEditorialStatusRequest"`
		AdGroupId       int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		EditorialStatus string   `xml:"https://adcenter.microsoft.com/v8 EditorialStatus"`
	}{XMLNS, adGroupId, editorialStatus}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdsByEditorialStatus"), SoapRequestBody{getAdsByEditorialStatusRequest}}, " ", " ")
	if err != nil {
		return ads, err
	}
	soapRespBody, err := a.Request("GetAdsByEditorialStatus", reqBody)
	if err != nil {
		return ads, err
	}
	return unmarshalAds(soapRespBody)
}

func (a *Auth) GetAdsByIds(adGroupId int64, adIds []int64) (ads []Ad, err error) {
	getAdsByIdsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdsByIdsRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		AdIds     []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdIds>long"`
	}{XMLNS, adGroupId, adIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdsByIds"), SoapRequestBody{getAdsByIdsRequest}}, " ", " ")
	if err != nil {
		return ads, err
	}
	soapRespBody, err := a.Request("GetAdsByIds", reqBody)
	if err != nil {
		return ads, err
	}
	return unmarshalAds(soapRespBody)
}

func (a *Auth) PauseAds(adGroupId int64, adIds []int64) error {
	pauseAdsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 PauseAdsRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		AdIds     []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdIds>long"`
	}{XMLNS, adGroupId, adIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("PauseAds"), SoapRequestBody{pauseAdsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("PauseAds", reqBody)
	return err
}

func (a *Auth) ResumeAds(adGroupId int64, adIds []int64) error {
	resumeAdsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 ResumeAdsRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		AdIds     []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdIds>long"`
	}{XMLNS, adGroupId, adIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("ResumeAds"), SoapRequestBody{resumeAdsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("ResumeAds", reqBody)
	return err
}

func (a *Auth) UpdateAds(adGroupId int64, ads []Ad) error {
	as := []Ad{}
	for _, a := range ads {
		a.EditorialStatus = ""
		a.Status = ""
		as = append(as, a)
	}
	updateAdsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 UpdateAdsRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		Ads       []Ad     `xml:"https://adcenter.microsoft.com/v8 Ads>Ad"`
	}{XMLNS, adGroupId, as}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("UpdateAds"), SoapRequestBody{updateAdsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("UpdateAds", reqBody)
	return err
}
