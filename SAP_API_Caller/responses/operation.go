package responses

type Operation struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			InspectionLot                 string `json:"InspectionLot"`
			InspPlanOperationInternalID   string `json:"InspPlanOperationInternalID"`
			OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
			InspectionOperation           string `json:"InspectionOperation"`
			InspectionOperationPlant      string `json:"InspectionOperationPlant"`
			BillOfOperationsType          string `json:"BillOfOperationsType"`
			BOOOperationInternalID        string `json:"BOOOperationInternalID"`
			WorkCenterInternalID          string `json:"WorkCenterInternalID"`
			StatusObject                  string `json:"StatusObject"`
			OperationControlProfile       string `json:"OperationControlProfile"`
			OperationConfirmation         bool   `json:"OperationConfirmation"`
			InspectionSubSystem           string `json:"InspectionSubSystem"`
			OperationText                 string `json:"OperationText"`
		} `json:"results"`
	} `json:"d"`
}
