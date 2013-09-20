package badsoap

import (
	"testing"
)

func TestAdGroups(t *testing.T) {

	a.Testing = t
	cids, err := a.AddCampaigns(a.AccountId, []Campaign{
		{
			Name:                      "AdGroups Name",
			Description:               "AdGroups Description",
			BudgetType:                "DailyBudgetAccelerated",
			DailyBudget:               20,
			MonthlyBudget:             240,
			ConversionTrackingEnabled: true,
			DaylightSaving:            true,
			TimeZone:                  "OsakaSapporoTokyo",
			Status:                    "Paused",
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	campaignId := cids[0]
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	adGroupIds, err := a.AddAdGroups(campaignId, []AdGroup{
		{
			Name:           "AdGroup Name",
			AdDistribution: "Search",
			BiddingModel:   "Keyword",
			Status:         "Draft",
			Network:        "OwnedAndOperatedAndSyndicatedSearch",
			StartDate:      &Date{1, 11, 2013},
			PricingModel:   "Cpc",
			Language:       "English",
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adGroupIds == nil {
		t.Fatalf("failed to create ad group")
	}

	adGroups, err := a.GetAdGroupsByCampaignId(campaignId)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adGroups == nil {
		t.Fatalf("failed to get ad group by campaign id")
	}

	adGroups, err = a.GetAdGroupsByIds(campaignId, adGroupIds)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adGroups == nil {
		t.Fatalf("failed to get ad group by ids")
	}

	err = a.PauseAdGroups(campaignId, adGroupIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.ResumeAdGroups(campaignId, adGroupIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

	adGroups[0].Name = "Test"
	err = a.UpdateAdGroups(campaignId, adGroups)
	if err != nil {
		t.Fatalf("%s\n%#v\n", err.Error(), adGroups)
	}

	err = a.DeleteAdGroups(campaignId, adGroupIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

}
