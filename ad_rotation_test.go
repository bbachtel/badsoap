package badsoap

import (
	"testing"
)

func TestGetAdRotation(t *testing.T) {

	a.Testing = t
	campaignId, adGroupId := testCampaign("Ad Rotation")
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	err := a.SetAdRotationToAdGroups(campaignId, map[int64]AdRotation{
		adGroupId: {"RotateAdsEvenly"},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	adGroupsAdRotation, err := a.GetAdRotationByAdGroupIds(campaignId, []int64{adGroupId})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adGroupsAdRotation == nil {
		t.Fatalf("failed to get ad groups ad rotation")
	}

}
