package badsoap

import (
	"testing"
)

func TestGoals(t *testing.T) {

	a.Testing = t
	campaignId, _ := testCampaign("Goals")
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	err := a.SetAnalyticsType(a.AccountId, "Enabled")
	if err != nil {
		t.Fatalf(err.Error())
	}

	goals := []Goal{
		{
			Name:      "signup",
			CostModel: "None",
			RevenueModel: RevenueModel{
				Type: "None",
			},
			DaysApplicableForConversion: "Seven",
			Steps: []Step{
				{
					Position: 1,
					Name:     "Step1",
					Type:     "Lead",
				},
				{
					Position: 2,
					Name:     "Step2",
					Type:     "Browse",
				},
				{
					Position: 3,
					Name:     "Step3",
					Type:     "Conversion",
				},
			},
		},
	}

	goalResults, err := a.AddGoals(a.AccountId, goals)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if goalResults == nil {
		t.Fatalf("failed to add goal")
	}

	goals, err = a.GetGoals(a.AccountId)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if goals == nil {
		t.Fatalf("failed to get goals from account id")
	}

	goals[0].DaysApplicableForConversion = "Fifteen"
	goals[0].Steps[2].Name = "new conversion name"
	updatedGoalResults, err := a.UpdateGoals(a.AccountId, goals)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if updatedGoalResults == nil {
		t.Fatalf("failed to update goals")
	}

	err = a.SetAnalyticsType(a.AccountId, "Disabled")
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.DeleteGoals(a.AccountId, []int64{goalResults[0].GoalId})
	if err != nil {
		t.Fatalf(err.Error())
	}

}
