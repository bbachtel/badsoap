package badsoap

import (
	"testing"
)

func TestNegativeSites(t *testing.T) {

	a.Testing = t
	campaignId, adGroupId := testCampaign("Negative Sites")
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	err := a.SetNegativeSitesToCampaigns(a.AccountId, map[int64][]string{
		campaignId: {"http://example.com", "http://other.com"},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	campaignIdsNegativeSites, err := a.GetNegativeSitesByCampaignIds(a.AccountId, []int64{campaignId})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if campaignIdsNegativeSites == nil {
		t.Fatalf("Failed to add negative sites to campaigns")
	}

	err = a.SetNegativeSitesToAdGroups(campaignId, map[int64][]string{
		adGroupId: {"http://example.com", "http://other.com"},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	adGroupIdsNegativeSites, err := a.GetNegativeSitesByAdGroupIds(campaignId, []int64{adGroupId})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adGroupIdsNegativeSites == nil {
		t.Fatalf("Failed to add negative sites to adGroups")
	}

}
