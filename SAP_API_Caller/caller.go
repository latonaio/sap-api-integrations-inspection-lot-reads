package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-inspection-lot-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetInspectionLot(inspectionLotGroup, inspectionLot, plant, material, billOfOperationsDesc, inspectionSpecification string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(inspectionLotGroup, inspectionLot)
				wg.Done()
			}()
		case "Operation":
			func() {
				c.Operation(inspectionLotGroup, inspectionLot)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(inspectionLotGroup, inspectionLot string) {
	data, err := c.callInspectionLotSrvAPIRequirementHeader("A_InspectionLot", inspectionLotGroup, inspectionLot)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callInspectionLotSrvAPIRequirementHeader(api, inspectionLotGroup, inspectionLot string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_INSPECTIONLOT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, inspectionLotGroup, inspectionLot)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Operation(inspectionLotGroup, inspectionLot string) {
	data, err := c.callInspectionLotSrvAPIRequirementOperation("A_InspLotOpCharacteristic", inspectionLotGroup, inspectionLot)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callInspectionLotSrvAPIRequirementOperation(api, inspectionLotGroup, inspectionLot string) ([]sap_api_output_formatter.Operation, error) {
	url := strings.Join([]string{c.baseURL, "API_INSPECTIONLOT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithOperation(req, inspectionLotGroup, inspectionLot)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, inspectionLotGroup, inspectionLot string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InspectionLotGroup eq '%s' and InspectionLot eq '%s'", inspectionLotGroup, inspectionLot))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithOperation(req *http.Request, inspectionLotGroup, inspectionLot string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("InspectionLotGroup eq '%s' and InspectionLot eq '%s'", inspectionLotGroup, inspectionLot))
	req.URL.RawQuery = params.Encode()
}
