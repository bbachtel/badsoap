package badsoap

import (
	"testing"
)

func TestAddCampaigns(t *testing.T) {

	a.Testing = t
	ids_created, err := a.AddCampaigns(a.AccountId, []Campaign{{
		Name:           "Campaign Name",
		Description:    "Campaign Description",
		BudgetType:     "DailyBudgetAccelerated",
		DailyBudget:    20,
		MonthlyBudget:  240,
		TimeZone:       "OsakaSapporoTokyo",
		DaylightSaving: true,
		Status:         "Paused",
	}})
	if err != nil {
		t.Fatalf("%#v", err)
	}
	if ids_created == nil {
		t.Fatalf("failed to create campaign")
	}

	campaigns_got, err := a.GetCampaignsByIds(a.AccountId, ids_created)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if campaigns_got == nil {
		t.Fatalf("failed to get campaigns by ids")
	}

	err = a.ResumeCampaigns(a.AccountId, ids_created)
	if err != nil {
		t.Fatalf(err.Error())
	}
	campaigns_got, err = a.GetCampaignsByIds(a.AccountId, ids_created)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if campaigns_got[0].Status != "Active" {
		t.Fatalf("expected camapign status to be \"Active\" but was \"%s\"", campaigns_got[0].Status)
	}

	err = a.PauseCampaigns(a.AccountId, ids_created)
	if err != nil {
		t.Fatalf(err.Error())
	}
	campaigns_got, err = a.GetCampaignsByIds(a.AccountId, ids_created)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if campaigns_got[0].Status != "Paused" {
		t.Fatalf("expected camapign status to be \"Paused\" but was \"%s\"", campaigns_got[0].Status)
	}

	campaigns_got[0].Name = "test"
	a.UpdateCampaigns(a.AccountId, campaigns_got)
	if err != nil {
		t.Fatalf(err.Error())
	}
	campaigns_got, err = a.GetCampaignsByIds(a.AccountId, ids_created)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if campaigns_got[0].Name != "test" {
		t.Fatalf("failed to update campaign name")
	}

	err = a.DeleteCampaigns(a.AccountId, ids_created)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
