package badsoap

import (
	"bytes"
	"crypto/md5"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	auth_file string
	a         = Auth{
		DeveloperToken: "BBD37VB98",
		ServiceUrl:     CM_SANDBOX_SERVICE_URI,
		Client:         &http.Client{},
	}
)

func initDummyServer() {
	server := simpleSoapServer()
	a.ServiceUrl = server.URL
}

func init() {
	flag.StringVar(&auth_file, "auth_file", "sandbox.xml", "xml file containing sandbox credentials")
	xmlFile, _ := os.Open(auth_file)

	xmlFile, err := os.Open(auth_file)
	if err != nil {
		initDummyServer()
		log.Printf("Error opening file:", err)
		return
	}

	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(b, &a)
	if err != nil {

		initDummyServer()

		log.Printf("Error parsing xml file:", err)
		return
	}

}

func testCampaign(name string) (campaignId, adGroupId int64) {
	cids, err := a.AddCampaigns(a.AccountId, []Campaign{
		{
			Name:                      name + " Campaign",
			Description:               name + " Description",
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
		log.Fatalf(err.Error())
	}
	campaignId = cids[0]
	adGroupIds, err := a.AddAdGroups(campaignId, []AdGroup{
		{
			Name:           name + " AdGroup",
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
		log.Fatalf(err.Error())
	}
	_, err = a.GetAdGroupsByIds(campaignId, adGroupIds)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return campaignId, adGroupIds[0]
}

type testMsg struct {
	Headers [][]string
	Body    string
}

type testReqResp struct {
	Request  testMsg
	Response testMsg
}

func testRequest(t *testing.T, wantedRequest testMsg, gotRequest *http.Request) {
	// compare headers
	gotHeaders := gotRequest.Header
	for _, keyval := range wantedRequest.Headers {
		key := keyval[0]
		wanted := keyval[1]
		got := gotHeaders.Get(key)
		if wanted != got {
			t.Errorf("expected key %s to be % but was %s", key, wanted, got)
		}
	}
	// compare total requst
	gotRequestBuf := new(bytes.Buffer)
	gotRequest.Write(gotRequestBuf)
	if strings.HasSuffix(gotRequestBuf.String(), wantedRequest.Body) {
		t.Errorf("expected\n%s\ngot\n%s", wantedRequest, gotRequestBuf.String())
	}
}

func simpleSoapServer() *httptest.Server {
	//cnt := 0
	handler := func(testResponseWriter http.ResponseWriter, gotRequest *http.Request) {
		reqBody, _ := ioutil.ReadAll(gotRequest.Body)
		h := md5.New()
		io.WriteString(h, string(reqBody))
		dir := fmt.Sprintf("./data/%s/%x", gotRequest.Header.Get("SOAPAction"), h.Sum(nil))
		fi, err := os.Open(fmt.Sprintf("%s/Response.xml", dir))
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer fi.Close()

		b, err := ioutil.ReadAll(fi)
		if err != nil {
			log.Fatalf(err.Error())
		}

		// set response
		testResponseWriter.Header().Set("test", "key")
		log.Printf(string(b))
		io.WriteString(testResponseWriter, string(b))
		//cnt += 1
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	return server
}

func soapServer(t *testing.T, tests []testReqResp) *httptest.Server {
	cnt := 0
	handler := func(testResponseWriter http.ResponseWriter, gotRequest *http.Request) {
		testRequest(t, tests[cnt].Request, gotRequest)
		// set response
		testResponse := tests[cnt].Response
		for _, keyvalue := range testResponse.Headers {
			testResponseWriter.Header().Set(keyvalue[0], keyvalue[1])
		}
		io.WriteString(testResponseWriter, testResponse.Body)
		cnt += 1
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	return server
}

type recordingTransport struct {
	client *http.Client
}

func (t *recordingTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	//	request := new(bytes.Buffer)
	//	req.Write(request)
	//	log.Printf("%s\n", request.String())
	if t.client != nil {
		return t.client.Do(req)
	}
	return nil, errors.New("dummy impl")
}
