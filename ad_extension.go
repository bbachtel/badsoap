package badsoap

import (
	"encoding/xml"

//	"log"
)

// EditorialStatus: "Active", "ActiveLimited", "Disapproved", "Inactive"
type CampaignAdExtension struct {
	AdExtension     AdExtension
	CampaignId      int64
	EditorialStatus string
}

type AddressType struct {
	CityName       string `xml:"https://adcenter.microsoft.com/v8 CityName"`
	CountryCode    string `xml:"https://adcenter.microsoft.com/v8 CountryCode"`
	PostalCode     string `xml:"https://adcenter.microsoft.com/v8 PostalCode"`
	ProvinceCode   string `xml:"https://adcenter.microsoft.com/v8 ProvinceCode"`
	ProvinceName   string `xml:"https://adcenter.microsoft.com/v8 ProvinceName"`
	StreetAddress  string `xml:"https://adcenter.microsoft.com/v8 StreetAddress"`
	StreetAddress2 string `xml:"https://adcenter.microsoft.com/v8 StreetAddress2"`
}

type GeoPointType struct {
	Latitude  int64 `xml:"https://adcenter.microsoft.com/v8 LatitudeInMicroDegrees"`
	Longitude int64 `xml:"https://adcenter.microsoft.com/v8 LatitudeInMicroDegrees"`
}

type SiteLink struct {
	DestinationUrl string `xml:"https://adcenter.microsoft.com/v8 DestinationUrl"`
	DisplayText    string `xml:"https://adcenter.microsoft.com/v8 DisplayText"`
}

// Location: "AddressLine1","AddressLine2","BusinessImage","Country","LocationExtensionBusinessName","MapIcon","SiteLinkDestinationUrl","SiteLinkDisplayText"
type AdExtensionEditorialReason struct {
	Location           string
	PublisherCountries []string
	ReasonCode         int64
	Term               string
}

// Operand: "BingAds_Grouping","BingAds_Label","Brand","Condition","ProductType","SKU","Id"
// Attribute:
//   "Condition" -> "New","Used","Refurbished","Remanufactured","Collectable","Open Box"
type ProductCondition struct {
	Attribute string `xml:"https://adcenter.microsoft.com/v8 Attribute"`
	Operand   string `xml:"https://adcenter.microsoft.com/v8 Operand"`
}

// Status: "Active", "Deleted"
// BusinessGeoCodeStatus: "Pending", "Complete", "Invalid", "Failed"
type AdExtension struct {
	AdExtensionType string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	Id      int64  `xml:"https://adcenter.microsoft.com/v8 Id,omitempty"`
	Status  string `xml:"https://adcenter.microsoft.com/v8 Status,omitempty"`
	Type    string `xml:"https://adcenter.microsoft.com/v8 Type,omitempty"`
	Version int64  `xml:"https://adcenter.microsoft.com/v8 Version,omitempty"`

	// CallAdExtension
	CountryCode           string `xml:"https://adcenter.microsoft.com/v8 CountryCode,omitempty"`
	IsCallOnly            bool   `xml:"https://adcenter.microsoft.com/v8 IsCallOnly,omitempty"`
	IsCallTrackingEnabled bool   `xml:"https://adcenter.microsoft.com/v8 IsCallTrackingEnabled,omitempty"`

	// SiteLinksAdExtension
	SiteLinks []SiteLink `xml:"https://adcenter.microsoft.com/v8 SiteLinks>SiteLink,omitempty"`

	// ProductAdExtension
	Name             string             `xml:"https://adcenter.microsoft.com/v8 Name,omitempty"`
	ProductSelection []ProductCondition `xml:"https://adcenter.microsoft.com/v8 ProductSelection,omitempty"`
	StoreId          int64              `xml:"https://adcenter.microsoft.com/v8 StoreId,omitempty"`
	StoreName        string             `xml:"https://adcenter.microsoft.com/v8 StoreName,omitempty"`

	// LocationAdExtension
	Address       *AddressType  `xml:"https://adcenter.microsoft.com/v8 Address,omitempty"`
	CompanyName   string        `xml:"https://adcenter.microsoft.com/v8 CompanyName,omitempty"`
	GeoCodeStatus string        `xml:"https://adcenter.microsoft.com/v8 GeoCodeStatus,omitempty"`
	GeoPoint      *GeoPointType `xml:"https://adcenter.microsoft.com/v8 GeoPoint,omitempty"`
	IconMediaId   int64         `xml:"https://adcenter.microsoft.com/v8 IconMediaId,omitempty"`
	ImageMediaId  int64         `xml:"https://adcenter.microsoft.com/v8 ImageMediaId,omitempty"`

	PhoneNumber                   string `xml:"https://adcenter.microsoft.com/v8 PhoneNumber,omitempty"`
	RequireTollFreeTrackingNumber bool   `xml:"https://adcenter.microsoft.com/v8 RequireTollFreeTrackingNumber,omitempty"`
}

type AdExtensionIdentity struct {
	Id      int64 `xml:"Id"`
	Version int64 `xml:"Version"`
}

func NewCallAdExtension(countryCode string, isCallOnly, isCallTrackingEnabled bool, phoneNumber string, requireTollFreeTrackingNumber bool) AdExtension {
	return AdExtension{
		AdExtensionType:               "CallAdExtension",
		Type:                          "CallAdExtension",
		CountryCode:                   countryCode,
		IsCallOnly:                    isCallOnly,
		IsCallTrackingEnabled:         isCallTrackingEnabled,
		PhoneNumber:                   phoneNumber,
		RequireTollFreeTrackingNumber: requireTollFreeTrackingNumber,
	}
}

func NewLocationAdExtension(address AddressType, companyName, geoCodeStatus string, geoPoint GeoPointType, iconMediaId, imageMediaId int64, phoneNumber string) AdExtension {
	return AdExtension{
		AdExtensionType: "LocationAdExtension",
		Type:            "LocationAdExtension",
		Address:         &address,
		CompanyName:     companyName,
		GeoCodeStatus:   geoCodeStatus,
		GeoPoint:        &geoPoint,
		IconMediaId:     iconMediaId,
		ImageMediaId:    imageMediaId,
		PhoneNumber:     phoneNumber,
	}
}

func NewProductAdExtension(name string, productSelection []ProductCondition, storeId int64, storeName string) AdExtension {
	return AdExtension{
		AdExtensionType:  "ProductAdExtension",
		Type:             "ProductAdExtension",
		Name:             name,
		ProductSelection: productSelection,
		StoreId:          storeId,
		StoreName:        storeName,
	}
}

func NewSiteLinksAdExtension(siteLinks []SiteLink) AdExtension {
	return AdExtension{
		AdExtensionType: "SiteLinksAdExtension",
		Type:            "SiteLinksAdExtension",
		SiteLinks:       siteLinks,
	}
}

type adExtIdToCmpId struct {
	AdExtensionId int64 `xml:"https://adcenter.microsoft.com/v8 AdExtensionId"`
	CampaignId    int64 `xml:"https://adcenter.microsoft.com/v8 CampaignId"`
}

func marshalAdExtIdToCmpIds(campaignIdToAdExtensionIds map[int64][]int64) []adExtIdToCmpId {
	adExtIdToCmpIds := []adExtIdToCmpId{}
	for campaignId, adExtensionIds := range campaignIdToAdExtensionIds {
		for _, adExtensionId := range adExtensionIds {
			adExtIdToCmpIds = append(adExtIdToCmpIds, adExtIdToCmpId{adExtensionId, campaignId})
		}
	}
	return adExtIdToCmpIds
}

func (a *Auth) AddAdExtensions(accountId int64, adExtensions []AdExtension) (adExtensionIdentities []AdExtensionIdentity, err error) {
	addAdExtensionsRequest := struct {
		XMLName      xml.Name      `xml:"https://adcenter.microsoft.com/v8 AddAdExtensionsRequest"`
		AccountId    int64         `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtensions []AdExtension `xml:"https://adcenter.microsoft.com/v8 AdExtensions>AdExtension2"`
	}{XMLNS, accountId, adExtensions}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("AddAdExtensions"), SoapRequestBody{addAdExtensionsRequest}}, " ", " ")
	if err != nil {
		return adExtensionIdentities, err
	}
	soapRespBody, err := a.Request("AddAdExtensions", reqBody)
	if err != nil {
		return adExtensionIdentities, err
	}
	opResponse := struct {
		AdExtensionIdentities []AdExtensionIdentity `xml:"AdExtensionIdentities>AdExtensionIdentity"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return adExtensionIdentities, err
	}
	return opResponse.AdExtensionIdentities, err
}

func (a *Auth) DeleteAdExtensions(accountId int64, adExtensionIds []int64) error {
	deleteAdExtensionsRequest := struct {
		XMLName        xml.Name `xml:"https://adcenter.microsoft.com/v8 DeleteAdExtensionsRequest"`
		AccountId      int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtensionIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdExtensionIds>long"`
	}{XMLNS, accountId, adExtensionIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("DeleteAdExtensions"), SoapRequestBody{deleteAdExtensionsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("DeleteAdExtensions", reqBody)
	return err
}

// adExtensionIdToCampaignIdAssociations AdExtensionIdToCampaignIdAssociation[]
func (a *Auth) DeleteAdExtensionsFromCampaigns(accountId int64, campaignIdToAdExtensionIds map[int64][]int64) error {
	deleteAdExtensionsFromCampaignsRequest := struct {
		XMLName         xml.Name         `xml:"https://adcenter.microsoft.com/v8 DeleteAdExtensionsFromCampaignsRequest"`
		AccountId       int64            `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtIdToCmpIds []adExtIdToCmpId `xml:"https://adcenter.microsoft.com/v8 AdExtensionIdToCampaignIdAssociations>AdExtensionIdToCampaignIdAssociation"`
	}{XMLNS, accountId, marshalAdExtIdToCmpIds(campaignIdToAdExtensionIds)}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("DeleteAdExtensionsFromCampaigns"), SoapRequestBody{deleteAdExtensionsFromCampaignsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("DeleteAdExtensionsFromCampaigns", reqBody)
	return err
}

// adExtensionType: "SiteLinksAdExtension","CallAdExtension","LocationAdExtension","ProductsAdExtension"
// associationFilter: "All","Associated"
func (a *Auth) GetAdExtensionIdsByAccountId(accountId int64, adExtensionType, associationFilter string) (adExtensionIds []int64, err error) {
	getAdExtensionIdsByAccountIdRequest := struct {
		XMLName           xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdExtensionIdsByAccountIdRequest"`
		AccountId         int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtensionType   string   `xml:"https://adcenter.microsoft.com/v8 AdExtensionType"`
		AssociationFilter string   `xml:"https://adcenter.microsoft.com/v8 AssociationFilter"`
	}{XMLNS, accountId, adExtensionType, associationFilter}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdExtensionIdsByAccountId"), SoapRequestBody{getAdExtensionIdsByAccountIdRequest}}, " ", " ")
	if err != nil {
		return adExtensionIds, err
	}
	soapRespBody, err := a.Request("GetAdExtensionIdsByAccountId", reqBody)
	if err != nil {
		return adExtensionIds, err
	}
	opResponse := struct {
		AdExtensionIds []int64 `xml:"AdExtensionIds>long"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return adExtensionIds, err
	}
	return opResponse.AdExtensionIds, err
}

// adExtensionType: "SiteLinksAdExtension","CallAdExtension","LocationAdExtension","ProductsAdExtension"
func (a *Auth) GetAdExtensionsByCampaignIds(accountId int64, campaignIds []int64, adExtensionType string) (campaignAdExtensions []CampaignAdExtension, err error) {
	getAdExtensionsByCampaignIdsRequest := struct {
		XMLName         xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdExtensionsByCampaignIdsRequest"`
		AccountId       int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		CampaignIds     []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays CampaignIds>long"`
		AdExtensionType string   `xml:"https://adcenter.microsoft.com/v8 AdExtensionType"`
	}{XMLNS, accountId, campaignIds, adExtensionType}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdExtensionsByCampaignIds"), SoapRequestBody{getAdExtensionsByCampaignIdsRequest}}, " ", " ")
	if err != nil {
		return campaignAdExtensions, err
	}
	soapRespBody, err := a.Request("GetAdExtensionsByCampaignIds", reqBody)
	if err != nil {
		return campaignAdExtensions, err
	}
	opResponse := struct {
		CampaignAdExtensions []CampaignAdExtension `xml:"CampaignAdExtensionCollection>CampaignAdExtensionCollection>CampaignAdExtensions>CampaignAdExtension"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return campaignAdExtensions, err
	}
	return opResponse.CampaignAdExtensions, err
}

// adExtensionType: "SiteLinksAdExtension","CallAdExtension","LocationAdExtension","ProductsAdExtension"
func (a *Auth) GetAdExtensionsByIds(accountId int64, adExtensionIds []int64, adExtensionType string) (adExtensions []AdExtension, err error) {
	getAdExtensionsByIdsRequest := struct {
		XMLName         xml.Name `xml:"https://adcenter.microsoft.com/v8 GetAdExtensionsByIdsRequest"`
		AccountId       int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtensionIds  []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays AdExtensionIds>long"`
		AdExtensionType string   `xml:"https://adcenter.microsoft.com/v8 AdExtensionType"`
	}{XMLNS, accountId, adExtensionIds, adExtensionType}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdExtensionsByIds"), SoapRequestBody{getAdExtensionsByIdsRequest}}, " ", " ")
	if err != nil {
		return adExtensions, err
	}
	soapRespBody, err := a.Request("GetAdExtensionsByIds", reqBody)
	if err != nil {
		return adExtensions, err
	}
	opResponse := struct {
		AdExtensions []AdExtension `xml:"AdExtensions>AdExtension2"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return adExtensions, err
	}
	return opResponse.AdExtensions, err
}

// adExtensionType: "SiteLinksAdExtension","CallAdExtension","LocationAdExtension","ProductsAdExtension"
func (a *Auth) GetAdExtensionsEditorialReasonsByCampaignIds(accountId int64, campaignIdToAdExtensionIds map[int64][]int64, adExtensionType string) (adExtensionIdsEditorialReasons map[int64][]AdExtensionEditorialReason, err error) {
	getAdExtensionsEditorialReasonsByCampaignIdsRequest := struct {
		XMLName         xml.Name         `xml:"https://adcenter.microsoft.com/v8 GetAdExtensionsEditorialReasonsByCampaignIdsRequest"`
		AccountId       int64            `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtIdToCmpIds []adExtIdToCmpId `xml:"https://adcenter.microsoft.com/v8 AdExtensionIdToCampaignIdAssociations>AdExtensionIdToCampaignIdAssociation"`
		AdExtensionType string           `xml:"https://adcenter.microsoft.com/v8 AdExtensionType"`
	}{XMLNS, accountId, marshalAdExtIdToCmpIds(campaignIdToAdExtensionIds), adExtensionType}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetAdExtensionsEditorialReasonsByCampaignIds"), SoapRequestBody{getAdExtensionsEditorialReasonsByCampaignIdsRequest}}, " ", " ")
	if err != nil {
		return adExtensionIdsEditorialReasons, err
	}
	soapRespBody, err := a.Request("GetAdExtensionsEditorialReasonsByCampaignIds", reqBody)
	if err != nil {
		return adExtensionIdsEditorialReasons, err
	}
	type adExtEditorialReasions struct {
		AdExtensionId    int64                        `xml:"AdExtensionId"`
		EditorialReasons []AdExtensionEditorialReason `xml:"Reasons>AdExtensionEditorialReason"`
	}
	opResponse := struct {
		AdExtIdsEditorialReasons []adExtEditorialReasions `xml:"EditorialReasons>AdExtensionEditorialReasonCollection"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return adExtensionIdsEditorialReasons, err
	}
	adExtensionIdsEditorialReasons = map[int64][]AdExtensionEditorialReason{}
	for _, er := range opResponse.AdExtIdsEditorialReasons {
		_, ok := adExtensionIdsEditorialReasons[er.AdExtensionId]
		if ok {
			for _, aeer := range er.EditorialReasons {
				aers := adExtensionIdsEditorialReasons[er.AdExtensionId]
				adExtensionIdsEditorialReasons[er.AdExtensionId] = append(aers, aeer)
			}
		} else {
			adExtensionIdsEditorialReasons[er.AdExtensionId] = er.EditorialReasons
		}
	}
	return adExtensionIdsEditorialReasons, nil
}

func (a *Auth) SetAdExtensionsToCampaigns(accountId int64, campaignIdToAdExtensionIds map[int64][]int64) error {
	setAdExtensionsToCampaignsRequest := struct {
		XMLName         xml.Name         `xml:"https://adcenter.microsoft.com/v8 SetAdExtensionsToCampaignsRequest"`
		AccountId       int64            `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtIdToCmpIds []adExtIdToCmpId `xml:"https://adcenter.microsoft.com/v8 AdExtensionIdToCampaignIdAssociations>AdExtensionIdToCampaignIdAssociation"`
	}{XMLNS, accountId, marshalAdExtIdToCmpIds(campaignIdToAdExtensionIds)}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("SetAdExtensionsToCampaigns"), SoapRequestBody{setAdExtensionsToCampaignsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetAdExtensionsToCampaigns", reqBody)
	return err
}

func (a *Auth) UpdateAdExtensions(accountId int64, adExtensions []AdExtension) error {
	for a, _ := range adExtensions {
		adExtensions[a].Version = 0
	}
	updateAdExtensionsRequest := struct {
		XMLName      xml.Name      `xml:"https://adcenter.microsoft.com/v8 UpdateAdExtensionsRequest"`
		AccountId    int64         `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		AdExtensions []AdExtension `xml:"https://adcenter.microsoft.com/v8 AdExtensions>AdExtension2"`
	}{XMLNS, accountId, adExtensions}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("UpdateAdExtensions"), SoapRequestBody{updateAdExtensionsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("UpdateAdExtensions", reqBody)
	return err
}
