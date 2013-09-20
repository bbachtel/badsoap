package badsoap

import (
	"encoding/xml"
	"strconv"
)

// analyticsType: "Enable","Disable"
func (a *Auth) SetAnalyticsType(accountId int64, analyticsType string) error {
	type accountAnalyticsType struct {
		AccountId int64  `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts AccountId"`
		Type      string `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Type"`
	}
	setAnalyticsTypeRequest := struct {
		XMLName             xml.Name             `xml:"https://adcenter.microsoft.com/v8 SetAnalyticsTypeRequest"`
		AcountAnalyticsType accountAnalyticsType `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts AccountAnalyticsTypes>AccountAnalyticsType"`
	}{XMLNS, accountAnalyticsType{accountId, analyticsType}}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("SetAnalyticsType"), SoapRequestBody{setAnalyticsTypeRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("SetAnalyticsType", reqBody)
	return err
}

// StepType: "Lead","Browse","Prospect","Conversion"
type Step struct {
	Id       int64  `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Id,omitempty"`
	Name     string `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Name"`
	Position int64  `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts PositionNumber,omitempty"`
	Script   string `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Script,omitempty"`
	Type     string `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Type,omitempty"`
}

// Type: "Constant","Variable","None"
type RevenueModel struct {
	ConstantRevenueValue float64 `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts ConstantRevenueValue,omitempty"`
	Type                 string  `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Type"`
}

// CostModel: "None","NonAdvertising","Taxed","Shipped"
// DaysApplicableForConversion: "Fiften","FortyFive","Seven","Thirty"
// RevenueModel: "Constant","Variable","None"
type Goal struct {
	CostModel                   string       `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts CostModel"`
	DaysApplicableForConversion string       `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts DaysApplicableForConversion"`
	Id                          int64        `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Id,omitempty"`
	Name                        string       `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Name"`
	RevenueModel                RevenueModel `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts RevenueModel"`
	Steps                       []Step       `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Steps>Step"`
	YEventId                    int64        `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts YEventId,omitempty"`
}

type GoalResult struct {
	XMLName xml.Name `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts GoalResult"`
	GoalId  int64    `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts GoalId"`
	StepIds []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays StepIds>long"`
}

func unmarshalGoalResults(xmlBytes []byte) (grs []GoalResult, err error) {
	opResponse := struct {
		GoalResults []GoalResult `xml:"GoalResults>GoalResult"`
	}{}
	err = xml.Unmarshal(xmlBytes, &opResponse)
	if err != nil {
		return grs, err
	}
	return opResponse.GoalResults, err
}

func (a *Auth) AddGoals(accountId int64, goals []Goal) (goalResults []GoalResult, err error) {
	addGoalsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 AddGoalsRequest"`
		AccountId int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		Goals     []Goal   `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Goals>Goal"`
	}{XMLNS, accountId, goals}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("AddGoals"), SoapRequestBody{addGoalsRequest}}, " ", " ")
	if err != nil {
		return goalResults, err
	}
	soapRespBody, err := a.Request("AddGoals", reqBody)
	if err != nil {
		return goalResults, err
	}
	return unmarshalGoalResults(soapRespBody)
}

func (a *Auth) DeleteGoals(accountId int64, goalIds []int64) error {
	deleteGoalsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 DeleteGoalsRequest"`
		AccountId int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		GoalIds   []int64  `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays GoalIds>long"`
	}{XMLNS, accountId, goalIds}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("DeleteGoals"), SoapRequestBody{deleteGoalsRequest}}, " ", " ")
	if err != nil {
		return err
	}
	_, err = a.Request("DeleteGoals", reqBody)
	return err
}

func (a *Auth) GetGoals(accountId int64) (goals []Goal, err error) {
	getGoalsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 GetGoalsRequest"`
		AccountId int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
	}{XMLNS, accountId}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("GetGoals"), SoapRequestBody{getGoalsRequest}}, " ", " ")
	if err != nil {
		return goals, err
	}
	soapRespBody, err := a.Request("GetGoals", reqBody)
	if err != nil {
		return goals, err
	}
	type unmarshalRevenueModel struct {
		ConstantRevenueValue string `xml:"ConstantRevenueValue"`
		Type                 string `xml:"Type"`
	}
	type unmarshalGoal struct {
		CostModel                   string                `xml:"CostModel"`
		DaysApplicableForConversion string                `xml:"DaysApplicableForConversion"`
		Id                          int64                 `xml:"Id,omitempty"`
		Name                        string                `xml:"Name"`
		RevenueModel                unmarshalRevenueModel `xml:"RevenueModel"`
		Steps                       []Step                `xml:"Steps>Step"`
		YEventId                    string                `xml:"YEventId"`
	}
	opResponse := struct {
		Goals []unmarshalGoal `xml:"Goals>Goal"`
	}{}
	err = xml.Unmarshal(soapRespBody, &opResponse)
	if err != nil {
		return goals, err
	}
	// hack to get around encoding/xml failing on empty int64 and float64 values without
	// exposing string encoded parameters in exported struct
	goals = []Goal{}
	for _, g := range opResponse.Goals {
		goal := Goal{
			Id:                          g.Id,
			Name:                        g.Name,
			CostModel:                   g.CostModel,
			DaysApplicableForConversion: g.DaysApplicableForConversion,
			RevenueModel: RevenueModel{
				Type: g.RevenueModel.Type,
			},
		}
		if g.RevenueModel.ConstantRevenueValue != "" {
			crv, err := strconv.ParseFloat(g.RevenueModel.ConstantRevenueValue, 64)
			if err != nil {
				goal.RevenueModel.ConstantRevenueValue = crv
			}
		}
		if g.YEventId != "" {
			yei, err := strconv.ParseInt(g.YEventId, 10, 64)
			if err != nil {
				goal.YEventId = yei
			}
		}
		steps := []Step{}
		for _, s := range g.Steps {
			steps = append(steps, Step{s.Id, s.Name, s.Position, s.Script, s.Type})
		}
		goal.Steps = steps
		goals = append(goals, goal)
	}
	return goals, err
}

func (a *Auth) UpdateGoals(accountId int64, goals []Goal) (goalResults []GoalResult, err error) {
	for a, g := range goals {
		for b, _ := range g.Steps {
			if goals[a].Steps[b].Id != 0 {
				goals[a].Steps[b].Type = ""
				goals[a].Steps[b].Position = 0
			}
			goals[a].Steps[b].Script = ""
		}
	}
	updateGoalsRequest := struct {
		XMLName   xml.Name `xml:"https://adcenter.microsoft.com/v8 UpdateGoalsRequest"`
		AccountId int64    `xml:"https://adcenter.microsoft.com/v8 AccountId"`
		Goals     []Goal   `xml:"http://schemas.datacontract.org/2004/07/Microsoft.AdCenter.Advertiser.CampaignManagement.Api.DataContracts Goals>Goal"`
	}{XMLNS, accountId, goals}
	reqBody, err := xml.MarshalIndent(SoapRequestEnvelope{ENVNS, a.authHeader("UpdateGoals"), SoapRequestBody{updateGoalsRequest}}, " ", " ")
	if err != nil {
		return goalResults, err
	}
	soapRespBody, err := a.Request("UpdateGoals", reqBody)
	if err != nil {
		return goalResults, err
	}
	return unmarshalGoalResults(soapRespBody)
}
