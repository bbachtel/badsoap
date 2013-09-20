package badsoap

import (
	"testing"
)

func TestAdExtensions(t *testing.T) {

	a.Testing = t
	campaignId, _ := testCampaign("AdExtensions")
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	adExtensionIdentities, err := a.AddAdExtensions(a.AccountId, []AdExtension{
		// NewCallAdExtension(),
		// NewLocationAdExtension(AddressType),
		// NewProductAdExtension(),
		NewSiteLinksAdExtension([]SiteLink{
			{"http://classdo.com", "classdo.com"},
			{"http://test.classdo.com", "test.classdo.com"},
		}),
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adExtensionIdentities == nil {
		t.Fatalf("failed to add ad extension")
	}

	adExtensionIds := []int64{}
	for _, adExtensionIdentity := range adExtensionIdentities {
		adExtensionIds = append(adExtensionIds, adExtensionIdentity.Id)
	}
	err = a.SetAdExtensionsToCampaigns(a.AccountId, map[int64][]int64{campaignId: adExtensionIds})
	if err != nil {
		t.Fatalf(err.Error())
	}

	campaignIdsAdExtensions, err := a.GetAdExtensionsByCampaignIds(a.AccountId, []int64{campaignId}, "SiteLinksAdExtension")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if campaignIdsAdExtensions == nil {
		t.Fatalf("failed to add ad extension to campaign")
	}

	adExtensions, err := a.GetAdExtensionsByIds(a.AccountId, adExtensionIds, "SiteLinksAdExtension")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adExtensions == nil {
		t.Fatalf("failed to get ad extensions by ids")
	}

	campaignIdToAdExtensionIds := map[int64][]int64{campaignId: adExtensionIds}
	adExtensionIdsEditorialReasons, err := a.GetAdExtensionsEditorialReasonsByCampaignIds(a.AccountId, campaignIdToAdExtensionIds, "SiteLinksAdExtension")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adExtensionIdsEditorialReasons == nil {
		t.Fatalf("failed to get ad extension editorial reasons")
	}

	adExtensionIds, err = a.GetAdExtensionIdsByAccountId(a.AccountId, "SiteLinksAdExtension", "All")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adExtensionIds == nil {
		t.Fatalf("failed to get ad extension ids by account id")
	}

	adExtensions[0].SiteLinks[0].DestinationUrl = "http://example.jp"
	adExtensions[0].SiteLinks[0].DisplayText = "example.jp"
	err = a.UpdateAdExtensions(a.AccountId, adExtensions)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.DeleteAdExtensionsFromCampaigns(a.AccountId, campaignIdToAdExtensionIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.DeleteAdExtensions(a.AccountId, adExtensionIds)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
