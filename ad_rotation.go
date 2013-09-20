package badsoap

import (
	"encoding/xml"
)

// Type: "OptimizeForClicks", "RotateAdsEvenly"
type AdRotation struct {
	Type string `xml:"https://adcenter.microsoft.com/v8 Type"`
	//startDate Date   `xml:"https://adcenter.microsoft.com/v8 StartDate"`
	//endDate   Date   `xml:"https://adcenter.microsoft.com/v8 EndDate"`
}

func (a *Auth) GetAdRotationByAdGroupIds(campaignId int64, adGroupIds []int64) (adGroupsAdRotations map[int64]AdRotation, err error) {
	getAdRotationByAdGroupIdsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdRotationByAdGroupIdsRequest"`
		AdGroupIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdGroupIds>long"`
		CampaignId int64    `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
	}{XMLNS, adGroupIds, campaignId}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdRotationByAdGroupIds"), SoapRequestBody{getAdRotationByAdGroupIdsRequest}}, " ", " ")
	soapRespBody, err := a.Request("GetAdRotationByAdGroupIds", reqBody)
	opResponse := struct {
		AdRotations []AdRotation `xml:"AdRotationByAdGroupIds>AdRotation"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return adGroupsAdRotations, err
	}
	adGroupsAdRotations = map[int64]AdRotation{}
	for i, adRotation := range opResponse.AdRotations {
		adGroupsAdRotations[adGroupIds[i]] = adRotation
	}
	return adGroupsAdRotations, err
}

func (a *Auth) SetAdRotationToAdGroups(campaignId int64, adGroupsAdRotations map[int64]AdRotation) error {
	type adGroupAdRotation struct {
		XMLName   xml.Name   `xml:"https://adcenter.microsoft.com/v8 AdGroupAdRotation"`
		AdGroupId int64      `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		AdRot     AdRotation `xml:"https://adcenter.microsoft.com/v8 AdRotation"`
	}
	allAdGroupAdRotations := []adGroupAdRotation{}
	for adGroupId, adRotation := range adGroupsAdRotations {
		allAdGroupAdRotations = append(allAdGroupAdRotations, adGroupAdRotation{XMLNS, adGroupId, adRotation})
	}
	setAdRotationToAdGroupsRequest := struct {
		XMLName            xml.Name            `xml:"https://adcenter.microsoft.com/v8 SetAdRotationToAdGroupsRequest"`
		AdGroupAdRotations []adGroupAdRotation `xml:"https://adcenter.microsoft.com/v8 AdGroupAdRotations>AdGroupAdRotation"`
		CampaignId         int64               `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
	}{
		xml.Name{"https://adcenter.microsoft.com/v8", "SetAdRotationToAdGroupsRequest"},
		allAdGroupAdRotations,
		campaignId,
	}
	requestXml := SoapRequestEnvelope{ENVNS, a.authHeader("SetAdRotationToAdGroups"), SoapRequestBody{setAdRotationToAdGroupsRequest}}
	reqBody, err := xml.MarshalIndent(requestXml, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetAdRotationToAdGroups", reqBody)
	return err
}
