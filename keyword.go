package badsoap

import (
	"encoding/xml"
)

type AmountType struct {
	Amount float64 `xml:"https://adcenter.microsoft.com/v8 ExactMatchBid>Amount,omitempty"`
}

// EditorialStatus: "Active","Disapproved","Inactive"
// Status: "Active","Paused","Deleted","Inactive"
type Keyword struct {
	BroadMatchBid    *AmountType `xml:"https://adcenter.microsoft.com/v8 BroadMatchBid,omitempty"`
	ContentMatchBid  *AmountType `xml:"https://adcenter.microsoft.com/v8 ContentMatchBid,omitempty"`
	EditorialStatus  string      `xml:"https://adcenter.microsoft.com/v8 EditorialStatus,omitempty"`
	ExactMatchBid    *AmountType `xml:"https://adcenter.microsoft.com/v8 ExactMatchBid,omitempty"`
	Id               int64       `xml:"https://adcenter.microsoft.com/v8 Id,omitempty"`
	NegativeKeywords []string    `xml:"https://adcenter.microsoft.com/v8 NegativeKeywords"`
	Param1           string      `xml:"https://adcenter.microsoft.com/v8 Param1"`
	Param2           string      `xml:"https://adcenter.microsoft.com/v8 Param2"`
	Param3           string      `xml:"https://adcenter.microsoft.com/v8 Param3"`
	PhraseMatchBid   *AmountType `xml:"https://adcenter.microsoft.com/v8 PhraseMatchBid,omitempty"`
	Status           string      `xml:"https://adcenter.microsoft.com/v8 Status,omitempty"`
	Text             string      `xml:"https://adcenter.microsoft.com/v8 Text,omitempty"`
}

// Location: "Unknown","Keyword","KeywordParam1","KeywordParam2","KeywordParam3","AdTitleDescription","AdTitle"
//   "AdDescription","DisplayUrl","LandingUrl","SiteDomain","BusinessName","PhoneName","CashbackTextParam","AltText",
//   "Audio","Video","Flash","CAsset","Image","Destination","Asset","Ad","Order","BiddingKeyword","Association","Script"

type EditorialReason struct {
	Location           string   `xml:"https://adcenter.microsoft.com/v8 Location"`
	PublisherCountries []string `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays PublisherCountries>string"`
	ReasonCode         int64    `xml:"https://adcenter.microsoft.com/v8 ReasonCode"`
	Term               string   `xml:"https://adcenter.microsoft.com/v8 Term"`
}

// AppealStatus: Apealable 1

type AppealStatus int64

const (
	Appealable AppealStatus = iota + 1
	AppealPending
	NotAppealable
)

type EditorialReasons struct {
	AdOrKeywordId int64             `xml:"https://adcenter.microsoft.com/v8 AdOrKeywordId"`
	AppealStatus  AppealStatus      `xml:"https://adcenter.microsoft.com/v8 AppealStatus"`
	Reasions      []EditorialReason `xml:"https://adcenter.microsoft.com/v8 EditorialReasons>EditorialReasonCollection"`
}

func unmarshalKeywords(xmlBytes []byte) (keywords []Keyword, err error) {
	opResponse := struct {
		Keywords []Keyword `xml:"Keywords>Keyword"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return keywords, err
	}
	emptyBid := AmountType{}
	for i, _ := range opResponse.Keywords {
		if *opResponse.Keywords[i].BroadMatchBid == emptyBid {
			opResponse.Keywords[i].BroadMatchBid = nil
		}
		if *opResponse.Keywords[i].ContentMatchBid == emptyBid {
			opResponse.Keywords[i].ContentMatchBid = nil
		}
		if *opResponse.Keywords[i].ExactMatchBid == emptyBid {
			opResponse.Keywords[i].ExactMatchBid = nil
		}
		if *opResponse.Keywords[i].PhraseMatchBid == emptyBid {
			opResponse.Keywords[i].PhraseMatchBid = nil
		}
	}
	return opResponse.Keywords, err
}

func unmarshalKeywordIds(xmlBytes []byte) (keywordIds []int64, err error) {
	opResponse := struct {
		KeywordIds []int64 `xml:"KeywordIds>long"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return keywordIds, err
	}
	return opResponse.KeywordIds, err
}

func (a *Auth) AddKeywords(adGroupId int64, keywords []Keyword) (keywordIds []int64, err error) {
	ks := []Keyword{}
	for _, k := range keywords {
		k.EditorialStatus = ""
		k.Id = 0
		ks = append(ks, k)
	}
	addKeywordsRequest := struct {
		XMLName   xml.Name  `xml:"https://adcenter.microsoft.com/v8 AddKeywordsRequest"`
		AdGroupId int64     `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		Keywords  []Keyword `xml:"https://adcenter.microsoft.com/v8 Keywords>Keyword"`
	}{XMLNS, adGroupId, ks}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("AddKeywords"), SoapRequestBody{addKeywordsRequest}}, " ", " ")
	if err != nil {
		return keywordIds, err
	}
	soapRespBody, err := a.Request("AddKeywords", reqBody)
	if err != nil {
		return keywordIds, err
	}
	return unmarshalKeywordIds(soapRespBody)
}

func (a *Auth) DeleteKeywords(adGroupId int64, keywordIds []int64) error {
	deleteKeywordsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 DeleteKeywordsRequest"`
		AdGroupId  int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		KeywordIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays KeywordIds>long"`
	}{XMLNS, adGroupId, keywordIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("DeleteKeywords"), SoapRequestBody{deleteKeywordsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("DeleteKeywords", reqBody)
	return err
}

func (a *Auth) GetDestinationUrlByKeywordIds(adGroupId int64, keywordIds []int64) (destinationUrls []string, err error) {
	getDestinationUrlByKeywordIdsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetDestinationUrlByKeywordIdsRequest"`
		AdGroupId  int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		KeywordIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays KeywordIds>long"`
	}{XMLNS, adGroupId, keywordIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetDestinationUrlByKeywordIds"), SoapRequestBody{getDestinationUrlByKeywordIdsRequest}}, " ", " ")
	if err != nil {
		return destinationUrls, err
	}
	soapRespBody, err := a.Request("GetDestinationUrlByKeywordIds", reqBody)
	if err != nil {
		return destinationUrls, err
	}
	opResponse := struct {
		DestinationUrls []string `xml:"DestinationUrls>string"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return destinationUrls, err
	}
	return opResponse.DestinationUrls, err
}

func (a *Auth) GetKeywordEditorialReasonsByIds(accountId int64, keywordIds []int64) (editorialReasons []EditorialReasons, err error) {
	getKeywordEditorialReasonsByIdsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetKeywordEditorialReasonsByIdsRequest"`
		KeywordIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays KeywordIds>long"`
		AccountId  int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
	}{XMLNS, keywordIds, accountId}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetKeywordEditorialReasonsByIds"), SoapRequestBody{getKeywordEditorialReasonsByIdsRequest}}, " ", " ")
	if err != nil {
		return editorialReasons, err
	}
	soapRespBody, err := a.Request("GetKeywordEditorialReasonsByIds", reqBody)
	if err != nil {
		return editorialReasons, err
	}
	opResponse := struct {
		EditorialReasonsCollections []EditorialReasons `xml:"EditorialReasons>EditorialReasonCollection"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return editorialReasons, err
	}
	return opResponse.EditorialReasonsCollections, err
}

func (a *Auth) GetKeywordsByAdGroupId(adGroupId int64) (keywords []Keyword, err error) {
	getKeywordsByAdGroupIdRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 GetKeywordsByAdGroupIdRequest"`
		AdGroupId int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
	}{XMLNS, adGroupId}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetKeywordsByAdGroupId"), SoapRequestBody{getKeywordsByAdGroupIdRequest}}, " ", " ")
	if err != nil {
		return keywords, err
	}
	soapRespBody, err := a.Request("GetKeywordsByAdGroupId", reqBody)
	if err != nil {
		return keywords, err
	}
	return unmarshalKeywords(soapRespBody)
}

func (a *Auth) GetKeywordsByEditorialStatus(adGroupId int64, editorialStatus string) (keywords []Keyword, err error) {
	getKeywordsByEditorialStatusRequest := struct {
		XMLName         xml.Name `xml:"https://adcenter.microsoft.com/v8 GetKeywordsByEditorialStatusRequest"`
		AdGroupId       int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		EditorialStatus string   `xml:"https://adcenter.microsoft.com/v8 EditorialStatus"`
	}{XMLNS, adGroupId, editorialStatus}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetKeywordsByEditorialStatus"), SoapRequestBody{getKeywordsByEditorialStatusRequest}}, " ", " ")
	if err != nil {
		return keywords, err
	}
	soapRespBody, err := a.Request("GetKeywordsByEditorialStatus", reqBody)
	if err != nil {
		return keywords, err
	}
	return unmarshalKeywords(soapRespBody)
}

func (a *Auth) GetKeywordsByIds(adGroupId int64, keywordIds []int64) (keywords []Keyword, err error) {
	getKeywordsByIdsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 GetKeywordsByIdsRequest"`
		AdGroupId  int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		KeywordIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays KeywordIds>long"`
	}{XMLNS, adGroupId, keywordIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetKeywordsByIds"), SoapRequestBody{getKeywordsByIdsRequest}}, " ", " ")
	if err != nil {
		return keywords, err
	}
	soapRespBody, err := a.Request("GetKeywordsByIds", reqBody)
	if err != nil {
		return keywords, err
	}
	return unmarshalKeywords(soapRespBody)
}

func (a *Auth) PauseKeywords(adGroupId int64, keywordIds []int64) error {
	pauseKeywordsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 PauseKeywordsRequest"`
		AdGroupId  int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		KeywordIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays KeywordIds>long"`
	}{XMLNS, adGroupId, keywordIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("PauseKeywords"), SoapRequestBody{pauseKeywordsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("PauseKeywords", reqBody)
	return err
}

func (a *Auth) ResumeKeywords(adGroupId int64, keywordIds []int64) error {
	resumeKeywordsRequest := struct {
		XMLName    xml.Name `xml:"https://adcenter.microsoft.com/v8 ResumeKeywordsRequest"`
		AdGroupId  int64    `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		KeywordIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays KeywordIds>long"`
	}{XMLNS, adGroupId, keywordIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("ResumeKeywords"), SoapRequestBody{resumeKeywordsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("ResumeKeywords", reqBody)
	return err
}

func (a *Auth) SetDestinationUrlToKeywords(adGroupId int64, keywordIdsDestinationUrl map[int64]string) error {
	type kwDestinationUrl struct {
		XMLName        xml.Name `xml:"KeywordDestinationUrl"`
		DestinationUrl string   `xml:"https://adcenter.microsoft.com/v8 DestinationUrl"`
		KeywordId      int64    `xml:"https://adcenter.microsoft.com/v8 KeywordId"`
	}
	allDestinationUrls := []kwDestinationUrl{}
	for keywordId, destinationUrl := range keywordIdsDestinationUrl {
		allDestinationUrls = append(allDestinationUrls, kwDestinationUrl{XMLNS, destinationUrl, keywordId})
	}
	setDestinationUrlToKeywordsRequest := struct {
		XMLName           xml.Name           `xml:"https://adcenter.microsoft.com/v8 SetDestinationUrlToKeywordsRequest"`
		AdGroupId         int64              `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		KwDestinationUrls []kwDestinationUrl `xml:"https://adcenter.microsoft.com/v8 KeywordDestinationUrls>KeywordDestinationUrl"`
	}{XMLNS, adGroupId, allDestinationUrls}
	requestXml := SoapRequestEnvelope{ENVNS, a.authHeader("SetDestinationUrlToKeywords"), SoapRequestBody{setDestinationUrlToKeywordsRequest}}
	reqBody, err := xml.MarshalIndent(requestXml, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetDestinationUrlToKeywords", reqBody)
	return err
}

func (a *Auth) UpdateKeywords(adGroupId int64, keywords []Keyword) error {
	ks := []Keyword{}
	for _, k := range keywords {
		k.EditorialStatus = ""
		k.Status = ""
		k.Text = ""
		ks = append(ks, k)
	}
	updateKeywordsRequest := struct {
		XMLName   xml.Name  `xml:"https://adcenter.microsoft.com/v8 UpdateKeywordsRequest"`
		AdGroupId int64     `xml:"https://adcenter.microsoft.com/v8 AdGroupId"`
		Keywords  []Keyword `xml:"https://adcenter.microsoft.com/v8 Keywords>Keyword"`
	}{XMLNS, adGroupId, ks}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("UpdateKeywords"), SoapRequestBody{updateKeywordsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("UpdateKeywords", reqBody)
	return err
}
