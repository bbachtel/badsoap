package badsoap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	//	"log"
	"net/http"
	"strings"
	"testing"

/*
	"crypto/md5"
	"os"
	"io"
*/
)

var (
	CM_PRODUCTION_SERVICE_URI = "https://adcenterapi.microsoft.com/api/advertiser/v8/CampaignManagement/CampaignManagementService.svc?wsdl"
	CM_SANDBOX_SERVICE_URI    = "https://api.sandbox.bingads.microsoft.com/Api/Advertiser/v8/CampaignManagement/CampaignManagementService.svc?wsdl"
	ENVNS                     = xml.Name{"http://schemas.xmlsoap.org/soap/envelope/", "e"}
	XMLSchemaNS               = xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "s"}
	XMLNS                     = xml.Name{"https://adcenter.microsoft.com/v8", "ms"}
	CM_API_NS                 = "https://adcenter.microsoft.com/v8"
	MS_ARRAY_NS               = "http://schemas.microsoft.com/2003/10/Serialization/Arrays"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type Auth struct {
	AccountId      int64
	CustomerId     int64
	DeveloperToken string
	UserName       string
	Password       string
	ServiceUrl     string
	Client         Client
	Testing        *testing.T
}

type ArrayOfLong struct {
	Long []int64 `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays long"`
}

type HeaderAction struct {
	MustUnderstand string ``
}

type AuthHeader struct {
	Action              string `xml:"https://adcenter.microsoft.com/v8 Action"`
	ApplicationToken    string `xml:"https://adcenter.microsoft.com/v8 ApplicationToken"`
	AuthenticationToken string `xml:"https://adcenter.microsoft.com/v8 AuthenticationToken"`
	CustomerAccountId   int64  `xml:"https://adcenter.microsoft.com/v8 CustomerAccountId"`
	CustomerId          int64  `xml:"https://adcenter.microsoft.com/v8 CustomerId"`
	DeveloperToken      string `xml:"https://adcenter.microsoft.com/v8 DeveloperToken"`
	UserName            string `xml:"https://adcenter.microsoft.com/v8 UserName"`
	Password            string `xml:"https://adcenter.microsoft.com/v8 Password"`
}

type AdApiError struct {
	Code      int64  `xml:"AdApiError>Code"`
	Details   string `xml:"AdApiError>Details"`
	ErrorCode string `xml:"AdApiError>ErrorCode"`
	Message   string `xml:"AdApiError>Message"`
}

type BatchError struct {
	Code      int64  `xml:"BatchError>Code"`
	Details   string `xml:"BatchError>Details"`
	ErrorCode string `xml:"BatchError>ErrorCode"`
	Index     int64  `xml:"BatchError>Index"`
	Message   string `xml:"BatchError>Message"`
}

type EditorialError struct {
	Appealable       bool   `xml:"EditorialError>Appealable"`
	Code             int64  `xml:"EditorialError>Code"`
	DisapprovedText  string `xml:"EditorialError>DisapprovedText"`
	ErrorCode        string `xml:"EditorialError>ErrorCode"`
	Index            int64  `xml:"EditorialError>Index"`
	Message          string `xml:"EditorialError>Message"`
	PublisherCountry string `xml:"EditorialError>PublisherCountry"`
}

type GoalError struct {
	BatchErrors []BatchError `xml:"GoalError>BatchErrors"`
	Index       int64        `xml:"GoalError>Index"`
	StepErrors  []BatchError `xml:"GoalError>StepErrors"`
}

type OperationError struct {
	Code      int64  `xml:"OperationError>Code"`
	Details   string `xml:"OperationError>Details"`
	ErrorCode string `xml:"OperationError>ErrorCode"`
	Message   string `xml:"OperationError>Message"`
}

type Fault struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      struct {
		XMLName xml.Name   `xml:"detail"`
		Errors  ErrorsType `xml:",any"`
	}
}

type ErrorsType struct {
	TrackingId      string           `xml:"TrackingId"`
	AdApiErrors     []AdApiError     `xml:"Errors"`
	BatchErrors     []BatchError     `xml:"BatchErrors"`
	EditorialErrors []EditorialError `xml:"EditorialErrors"`
	GoalErrors      []GoalError      `xml:"GoalErrors"`
	OperationErrors []OperationError `xml:"OperationErrors"`
}

func (f *ErrorsType) Error() string {
	errors := []string{}
	for _, e := range f.AdApiErrors {
		errors = append(errors, fmt.Sprintf("%s", e.Message))
	}
	for _, e := range f.BatchErrors {
		errors = append(errors, fmt.Sprintf("%s", e.Message))
	}
	for _, e := range f.EditorialErrors {
		errors = append(errors, fmt.Sprintf("%s", e.Message))
	}
	for _, e := range f.GoalErrors {
		for _, be := range e.BatchErrors {
			errors = append(errors, fmt.Sprintf("%s", be.Message))
		}
		for _, be := range e.StepErrors {
			errors = append(errors, fmt.Sprintf("%s", be.Message))
		}
	}
	for _, e := range f.OperationErrors {
		errors = append(errors, fmt.Sprintf("%s", e.Message))
	}
	return strings.Join(errors, "\n")
}

type SoapRequestEnvelope struct {
	XMLName xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  AuthHeader      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    SoapRequestBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

func (a *Auth) authHeader(action string) AuthHeader {
	return AuthHeader{action, "", "", a.AccountId, a.CustomerId, a.DeveloperToken, a.UserName, a.Password}
}

type SoapRequestBody struct {
	Body interface{}
}

type SoapResponseBody struct {
	OperationResponse []byte `xml:",innerxml"`
}

type SoapResponseEnvelope struct {
	XMLName xml.Name         `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  TrackingId       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    SoapResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type TrackingId struct {
	Nil        bool   `xml:"http://www.w3.org/2001/XMLSchema-instance nil,attr"`
	TrackingId string `xml:"https://adcenter.microsoft.com/v8 TrackingId"`
}

type Date struct {
	Day   int64 `xml:"https://adcenter.microsoft.com/v8 Day"`
	Month int64 `xml:"https://adcenter.microsoft.com/v8 Month"`
	Year  int64 `xml:"https://adcenter.microsoft.com/v8 Year"`
}

func (c *Auth) Request(action string, reqBody []byte) (respBody []byte, err error) {
	req, err := http.NewRequest("POST", c.ServiceUrl, bytes.NewReader(reqBody))
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("Accept", "multipart/*")
	req.Header.Add("Content-type", "text/xml; charset=\"UTF-8\"")
	contentLength := fmt.Sprintf("%d", len(reqBody))
	req.Header.Add("Content-length", contentLength)
	req.Header.Add("SOAPAction", action)
	if c.Testing != nil {
		c.Testing.Logf("request ->\n%s\n%#v\n", string(reqBody), req.URL.String())
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBody, err
	}

	/*
		h := md5.New()
		io.WriteString(h,string(reqBody))
		dir := fmt.Sprintf("./data/%s/%x",action,h.Sum(nil))
		os.MkdirAll(dir,0744)
		fo,err := os.Create(fmt.Sprintf("%s/Request.xml",dir))
		fo.Write(reqBody)
		fo.Close()
		fo,err = os.Create(fmt.Sprintf("%s/Response.xml",dir))
		fo.Write(respBody)
		fo.Close()
	*/
	if c.Testing != nil {
		c.Testing.Logf("respBody ->\n%s\n%s\n", string(respBody), resp.Status)
	}
	response := SoapResponseEnvelope{}
	err = xml.Unmarshal([]byte(respBody), &response)
	if err != nil {
		return respBody, err
	}
	if resp.StatusCode == 400 || resp.StatusCode == 401 || resp.StatusCode == 403 || resp.StatusCode == 405 || resp.StatusCode == 500 {
		//log.Printf("request ->\n%s\n%#v\n", string(reqBody), req.URL.String())
		//log.Printf("%s ->\n%s\n",action,response.Body.OperationResponse)
		fault := Fault{}
		err = xml.Unmarshal(response.Body.OperationResponse, &fault)
		if err != nil {
			return respBody, err
		}
		return respBody, &fault.Detail.Errors //errors
	}
	return response.Body.OperationResponse, err
}
