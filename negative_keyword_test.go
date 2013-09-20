package badsoap

import (
	"testing"
)

func TestNegativeKeywords(t *testing.T) {

	a.Testing = t
	campaignId, adGroupId := testCampaign("Negative Keywords")
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	err := a.SetNegativeKeywordsToCampaigns(a.AccountId, map[int64][]string{campaignId: {"test1", "test2"}})
	if err != nil {
		t.Fatalf(err.Error())
	}

	campaignIdsNegativeKeywords, err := a.GetNegativeKeywordsByCampaignIds(a.AccountId, []int64{campaignId})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if campaignIdsNegativeKeywords == nil {
		t.Fatalf("failed to get campaign negative keywords")
	}

	err = a.SetNegativeKeywordsToAdGroups(campaignId, map[int64][]string{adGroupId: {"test1", "test2"}})
	if err != nil {
		t.Fatalf(err.Error())
	}

	adGroupIdsNegativeKeywords, err := a.GetNegativeKeywordsByAdGroupIds(campaignId, []int64{adGroupId})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adGroupIdsNegativeKeywords == nil {
		t.Fatalf("failed to get ad group negative keywords")
	}

}
