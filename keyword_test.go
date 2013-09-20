package badsoap

import (
	"testing"
)

func TestKeywords(t *testing.T) {

	a.Testing = t
	campaignId, adGroupId := testCampaign("Keywords")
	defer a.DeleteCampaigns(a.AccountId, []int64{campaignId})

	keywordIds, err := a.AddKeywords(adGroupId, []Keyword{
		{
			Text:             "testing",
			ExactMatchBid:    &AmountType{0.5},
			NegativeKeywords: []string{},
			Status:           "Paused",
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	if keywordIds == nil {
		t.Fatalf("failed to add keywords")
	}

	destinationUrls, err := a.GetDestinationUrlByKeywordIds(adGroupId, keywordIds)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if destinationUrls == nil {
		t.Fatalf("failed to get destination urls from keyword ids")
	}

	editorialReasons, err := a.GetKeywordEditorialReasonsByIds(a.AccountId, keywordIds)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if editorialReasons == nil {
		t.Fatalf("failed to get editorial resions from keyword ids")
	}

	keywords, err := a.GetKeywordsByAdGroupId(adGroupId)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if keywords == nil {
		t.Fatalf("failed to get keywords from ad group ids")
	}

	keywords, err = a.GetKeywordsByEditorialStatus(adGroupId, "Active")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if keywords == nil {
		t.Fatalf("failed to get keywords by editorial reasion")
	}

	keywords, err = a.GetKeywordsByIds(adGroupId, keywordIds)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if keywords == nil {
		t.Fatalf("failed to get keywords by their ids")
	}

	err = a.ResumeKeywords(adGroupId, keywordIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.PauseKeywords(adGroupId, keywordIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.SetDestinationUrlToKeywords(adGroupId, map[int64]string{
		keywordIds[0]: "http://example.com/",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	keywords[0].ExactMatchBid = &AmountType{0.9}
	err = a.UpdateKeywords(adGroupId, keywords)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = a.DeleteKeywords(adGroupId, keywordIds)
	if err != nil {
		t.Fatalf(err.Error())
	}

}
