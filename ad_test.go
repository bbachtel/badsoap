package badsoap

import (
	"testing"
)

func TestUnmarshalAds(t *testing.T) {
	xmlStr := `<GetAdsByIdsResponse xmlns="https://adcenter.microsoft.com/v8"><Ads xmlns:i="http://www.w3.org/2001/XMLSchema-instance"><Ad i:type="TextAd"><EditorialStatus>Active</EditorialStatus><Id>2400239690</Id><Status>Paused</Status><Type>Text</Type><DestinationUrl>http://example.com</DestinationUrl><DisplayUrl>example.com</DisplayUrl><Text>fadfd fdfda fdafd</Text><Title>test</Title></Ad></Ads></GetAdsByIdsResponse>`
	_, err := unmarshalAds([]byte(xmlStr))
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUnmarshalAdIds(t *testing.T) {
	xmlStr := `<AddAdsResponse xmlns="https://adcenter.microsoft.com/v8"><AdIds xmlns:a="http://schemas.microsoft.com/2003/10/Serialization/Arrays" xmlns:i="http://www.w3.org/2001/XMLSchema-instance"><a:long>2400239732</a:long></AdIds></AddAdsResponse>`
	_, err := unmarshalAdIds([]byte(xmlStr))
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAds(t *testing.T) {
	a.Testing = t
	campaignId, adGroupId := testCampaign("Ad")
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	adIds, err := a.AddAds(adGroupId, []Ad{
		NewTextAd("test1", "fadfd fdfda fdafd", "https://example.com", "example.com", "Paused"),
		NewMobileAd("test2", "fdsaf gafdsfs", "Business Name", "https://example.com", "example.com", "62653657645", "Paused"),
		//NewProductAd("promotional text","Paused"),
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if adIds == nil {
		t.Fatalf("failed to add ads")
	}

	ads, err := a.GetAdsByAdGroupId(adGroupId)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if ads == nil {
		t.Fatalf("failed to get ads by ad group id")
	}

	ads, err = a.GetAdsByEditorialStatus(adGroupId, "Active")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if ads == nil {
		t.Fatalf("failed to get ads by editorial status")
	}

	ads, err = a.GetAdsByIds(adGroupId, adIds)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if ads == nil {
		t.Fatalf("failed to get ads by ad ids")
	}

	err = a.ResumeAds(adGroupId, adIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.PauseAds(adGroupId, adIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

	ads[0].Text = "testing testing one two three"
	err = a.UpdateAds(adGroupId, ads)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.DeleteAds(adGroupId, adIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

}
