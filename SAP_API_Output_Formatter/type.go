package sap_api_output_formatter

type Header struct {
	InspectionLot                  string `json:"InspectionLot"`
	Material                       string `json:"Material"`
	Batch                          string `json:"Batch"`
	Plant                          bool   `json:"Plant"`
	InspectionLotOrigin            string `json:"InspectionLotOrigin"`
	OrderInternalBillOfOperations  string `json:"OrderInternalBillOfOperations"`
	ManufacturingOrder             string `json:"ManufacturingOrder"`
	InspectionLotText              string `json:"InspectionLotText"`
	InspectionLotType              string `json:"InspectionLotType"`
	InspectionLotQuantity          string `json:"InspectionLotQuantity"`
	InspectionLotActualQuantity    string `json:"InspectionLotActualQuantity"`
	InspectionLotDefectiveQuantity string `json:"InspectionLotDefectiveQuantity"`
	InspectionLotQuantityUnit      string `json:"InspectionLotQuantityUnit"`
	InspLotCreatedOnLocalDate      bool   `json:"InspLotCreatedOnLocalDate"`
	InspSubsetFieldCombination     string `json:"InspSubsetFieldCombination"`
	InspLotNmbrOpenLongTermCharc   string `json:"InspLotNmbrOpenLongTermCharc"`
	StatusObject                   string `json:"StatusObject"`
	StatusObjectCategory           string `json:"StatusObjectCategory"`
	InspectionLotObjectText        string `json:"InspectionLotObjectText"`
	StatusProfile                  string `json:"StatusProfile"`
	MatlQualityAuthorizationGroup  string `json:"MatlQualityAuthorizationGroup"`
	InspectionLotHasQuantity       string `json:"InspectionLotHasQuantity"`
	InspLotIsCreatedAutomatically  string `json:"InspLotIsCreatedAutomatically"`
	InspectionLotHasPartialLots    string `json:"InspectionLotHasPartialLots"`
	InspectionLotHasAppraisalCosts string `json:"InspectionLotHasAppraisalCosts"`
	InspLotHasSubsets              string `json:"InspLotHasSubsets"`
	InspLotIsAutomUsgeDcsnPossible string `json:"InspLotIsAutomUsgeDcsnPossible"`
	PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
	InspLotHasConfignSpecification string `json:"InspLotHasConfignSpecification"`
	GoodsReceiptIsMovedToBlkdStock string `json:"GoodsReceiptIsMovedToBlkdStock"`
	InspLotIsDocumentationRequired string `json:"InspLotIsDocumentationRequired"`
	InspLotIsTaskListRequired      bool   `json:"InspLotIsTaskListRequired"`
	InspLotHasManualSampleSize     string `json:"InspLotHasManualSampleSize"`
	InspLotHasMaterialSpec         string `json:"InspLotHasMaterialSpec"`
	InspectionLotApproval          string `json:"InspectionLotApproval"`
	InspLotDigitalSgntrResultsRecg string `json:"InspLotDigitalSgntrResultsRecg"`
	InspLotDigitalSgntrInUsgeDcsn  string `json:"InspLotDigitalSgntrInUsgeDcsn"`
	InspLotIsBatchRequired         string `json:"InspLotIsBatchRequired"`
	InspLotUsageInStabilityStudy   string `json:"InspLotUsageInStabilityStudy"`
	InspLotIsStockPostingCompleted string `json:"InspLotIsStockPostingCompleted"`
	InspectionLotHasUsageDecision  string `json:"InspectionLotHasUsageDecision"`
	NumberOfSerialNumbers          bool   `json:"NumberOfSerialNumbers"`
	InspLotIsSerialNmbrPossible    string `json:"InspLotIsSerialNmbrPossible"`
	InspectionLotIsSkipped         string `json:"InspectionLotIsSkipped"`
	InspectionLotIsFullInspection  string `json:"InspectionLotIsFullInspection"`
	InspectionLotDynamicLevel      string `json:"InspectionLotDynamicLevel"`
	SamplingProcedure              string `json:"SamplingProcedure"`
	InspLotCreatedOnLocalTime      string `json:"InspLotCreatedOnLocalTime"`
	InspectionLotCreatedBy         string `json:"InspectionLotCreatedBy"`
	InspectionLotCreatedOn         string `json:"InspectionLotCreatedOn"`
	InspectionLotCreatedOnTime     string `json:"InspectionLotCreatedOnTime"`
	InspectionLotChangedBy         string `json:"InspectionLotChangedBy"`
	InspectionLotChangeDate        string `json:"InspectionLotChangeDate"`
	InspectionLotChangeTime        string `json:"InspectionLotChangeTime"`
	ChangedDateTime                string `json:"ChangedDateTime"`
	InspectionLotStartDate         string `json:"InspectionLotStartDate"`
	InspectionLotStartTime         string `json:"InspectionLotStartTime"`
	InspectionLotEndDate           string `json:"InspectionLotEndDate"`
	InspectionLotEndTime           string `json:"InspectionLotEndTime"`
	BillOfOperationsType           bool   `json:"BillOfOperationsType"`
	BillOfOperationsGroup          string `json:"BillOfOperationsGroup"`
	BillOfOperationsUsage          string `json:"BillOfOperationsUsage"`
	BillOfOperationsVariant        string `json:"BillOfOperationsVariant"`
	BillOfOperationsChangeStateID  string `json:"BillOfOperationsChangeStateID"`
	InspectionSubsetType           string `json:"InspectionSubsetType"`
	SmplDrawingProcedure           string `json:"SmplDrawingProcedure"`
	SmplDrawingProcedureVersion    string `json:"SmplDrawingProcedureVersion"`
	SmplDrwgProcedIsConfRequired   string `json:"SmplDrwgProcedIsConfRequired"`
	InspLotSelectionMaterial       string `json:"InspLotSelectionMaterial"`
	InspLotSelMatlRevisionLvl      bool   `json:"InspLotSelMatlRevisionLvl"`
	InspLotSelectionPlant          string `json:"InspLotSelectionPlant"`
	InspLotSelectionSupplier       string `json:"InspLotSelectionSupplier"`
	InspLotSelectionManufacturer   string `json:"InspLotSelectionManufacturer"`
	InspLotSelectionCustomer       string `json:"InspLotSelectionCustomer"`
	InspLotSelBillOfOperationsUsge string `json:"InspLotSelBillOfOperationsUsge"`
	InspLotSelectionValidFromDate  string `json:"InspLotSelectionValidFromDate"`
	ProductionVersion              string `json:"ProductionVersion"`
	SalesOperationsPlanningOrder   string `json:"SalesOperationsPlanningOrder"`
	IsBusinessPurposeCompleted     string `json:"IsBusinessPurposeCompleted"`
	Customer                       string `json:"Customer"`
	Supplier                       string `json:"Supplier"`
	BatchBySupplier                string `json:"BatchBySupplier"`
	Manufacturer                   string `json:"Manufacturer"`
	ManufacturerPartNmbr           string `json:"ManufacturerPartNmbr"`
	MaterialRevisionLevel          string `json:"MaterialRevisionLevel"`
	MaterialIsBatchManaged         string `json:"MaterialIsBatchManaged"`
	BatchStorageLocation           string `json:"BatchStorageLocation"`
	MaterialCompIsSpecialStock     bool   `json:"MaterialCompIsSpecialStock"`
	PurchasingOrganization         string `json:"PurchasingOrganization"`
	PurchasingDocument             string `json:"PurchasingDocument"`
	PurchasingDocumentItem         string `json:"PurchasingDocumentItem"`
	ScheduleLine                   string `json:"ScheduleLine"`
	AccountingDocumentType         string `json:"AccountingDocumentType"`
	MaterialDocumentYear           string `json:"MaterialDocumentYear"`
	MaterialDocument               string `json:"MaterialDocument"`
	MaterialDocumentItem           string `json:"MaterialDocumentItem"`
	MatlDocLatestPostgDate         string `json:"MatlDocLatestPostgDate"`
	GoodsMovementType              bool   `json:"GoodsMovementType"`
	InspectionLotPlant             string `json:"InspectionLotPlant"`
	InspectionLotStorageLocation   string `json:"InspectionLotStorageLocation"`
	Warehouse                      string `json:"Warehouse"`
	StorageType                    string `json:"StorageType"`
	StorageBin                     string `json:"StorageBin"`
	SalesOrder                     string `json:"SalesOrder"`
	SalesOrderItem                 string `json:"SalesOrderItem"`
	DeliveryDocument               string `json:"DeliveryDocument"`
	DeliveryDocumentItem           string `json:"DeliveryDocumentItem"`
	DeliveryCategory               string `json:"DeliveryCategory"`
	InspectionDeliveryCategory     string `json:"InspectionDeliveryCategory"`
	Route                          string `json:"Route"`
	BillToPartyCountry             string `json:"BillToPartyCountry"`
	SoldToParty                    string `json:"SoldToParty"`
	SalesOrganization              string `json:"SalesOrganization"`
	MaterialByCustomer             string `json:"MaterialByCustomer"`
	Language                       string `json:"Language"`
	InspLotNmbrAddlRecordedCharc   string `json:"InspLotNmbrAddlRecordedCharc"`
	InspLotNmbrOpenShortTermCharc  string `json:"InspLotNmbrOpenShortTermCharc"`
	InspectionLotContainer         string `json:"InspectionLotContainer"`
	InspectionLotContainerUnit     string `json:"InspectionLotContainerUnit"`
	InspectionLotSampleQuantity    string `json:"InspectionLotSampleQuantity"`
	InspectionLotSampleUnit        string `json:"InspectionLotSampleUnit"`
	InspLotDynamicRule             string `json:"InspLotDynamicRule"`
	InspLotDynamicTrggrPoint       string `json:"InspLotDynamicTrggrPoint"`
	InspectionDynamicStage         string `json:"InspectionDynamicStage"`
	InspectionSeverity             string `json:"InspectionSeverity"`
	InspLotQtyToFree               string `json:"InspLotQtyToFree"`
	InspLotQtyToScrap              string `json:"InspLotQtyToScrap"`
	InspLotQtyToSample             string `json:"InspLotQtyToSample"`
	InspLotQtyToBlocked            string `json:"InspLotQtyToBlocked"`
	InspLotQtyToReserves           string `json:"InspLotQtyToReserves"`
	InspLotQtyToAnotherMaterial    string `json:"InspLotQtyToAnotherMaterial"`
	InspLotMaterialPostedTo        string `json:"InspLotMaterialPostedTo"`
	InspLotBatchTransferredTo      string `json:"InspLotBatchTransferredTo"`
	InspLotQtyReturnedToSupplier   string `json:"InspLotQtyReturnedToSupplier"`
	InspLotQtyToSpecialStock       string `json:"InspLotQtyToSpecialStock"`
	InspLotQtyToOtherStock         string `json:"InspLotQtyToOtherStock"`
	InspLotQtyToBePosted           bool   `json:"InspLotQtyToBePosted"`
	InspLotSmplQtyForLongTermChar  string `json:"InspLotSmplQtyForLongTermChar"`
	InspLotQtyInspected            string `json:"InspLotQtyInspected"`
	InspLotQtyDestroyed            string `json:"InspLotQtyDestroyed"`
	InspectionLotScrapRatio        string `json:"InspectionLotScrapRatio"`
	InspLotUsageDecisionTool       string `json:"InspLotUsageDecisionTool"`
	InspectionLotAllowedScrapRatio string `json:"InspectionLotAllowedScrapRatio"`
	QualityCostCollector           string `json:"QualityCostCollector"`
	ConsumptionPosting             string `json:"ConsumptionPosting"`
	AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
	PurchasingDocumentItemCategory bool   `json:"PurchasingDocumentItemCategory"`
	InspLotAcctAssgmtKey           string `json:"InspLotAcctAssgmtKey"`
	CostCenter                     string `json:"CostCenter"`
	ReservationItem                string `json:"ReservationItem"`
	MasterFixedAsset               string `json:"MasterFixedAsset"`
	FixedAsset                     string `json:"FixedAsset"`
	SalesOrdStockWBSElement        string `json:"SalesOrdStockWBSElement"`
	ProjectNetwork                 string `json:"ProjectNetwork"`
	NetworkActivityInternalID      string `json:"NetworkActivityInternalID"`
	InventorySpclStkSalesDocument  string `json:"InventorySpclStkSalesDocument"`
	InventorySpclStkSalesDocItm    string `json:"InventorySpclStkSalesDocItm"`
	ProfitabilitySegment           string `json:"ProfitabilitySegment"`
	ProfitCenter                   string `json:"ProfitCenter"`
	BusinessArea                   string `json:"BusinessArea"`
	GLAccount                      string `json:"GLAccount"`
	ControllingArea                string `json:"ControllingArea"`
	CompanyCode                    string `json:"CompanyCode"`
	SerialNumberProfile            string `json:"SerialNumberProfile"`
	InspLotCostCollectorSalesOrder bool   `json:"InspLotCostCollectorSalesOrder"`
	InspLotCostCollectorSlsOrdItem string `json:"InspLotCostCollectorSlsOrdItem"`
	InspLotCostCollectorWBSElement string `json:"InspLotCostCollectorWBSElement"`
	InspLotExternalNumber          string `json:"InspLotExternalNumber"`
	InspectionLotPriorityPoints    string `json:"InspectionLotPriorityPoints"`
	MaintenancePlan                string `json:"MaintenancePlan"`
	MaintenancePlanItemIntID       string `json:"MaintenancePlanItemIntID"`
	MaintenanceStrategy            string `json:"MaintenanceStrategy"`
	to_InspectionLotWithStatus     string `json:"to_InspectionLotWithStatus"`
	InspectionLotStatusCreated     bool   `json:"InspectionLotStatusCreated"`
	InspectionLotStatusReleased    string `json:"InspectionLotStatusReleased"`
	InspectionLotStatusSkip        string `json:"InspectionLotStatusSkip"`
	InspLotStatusRsltsConfirmed    string `json:"InspLotStatusRsltsConfirmed"`
	InspLotStsDefectsRecorded      string `json:"InspLotStsDefectsRecorded"`
	InspLotStsShrtTrmInspCmpltd    string `json:"InspLotStsShrtTrmInspCmpltd"`
	InspLotStatusInspCompleted     string `json:"InspLotStatusInspCompleted"`
	InspLotStatusCanceled          string `json:"InspLotStatusCanceled"`
	InspLotStatusRepair            string `json:"InspLotStatusRepair"`
}

type MaterialAssignment struct {
	Material                      string `json:"Material"`
	Plant                         string `json:"Plant"`
	InspectionPlanGroup           string `json:"InspectionPlanGroup"`
	InspectionPlan                string `json:"InspectionPlan"`
	InspPlanMatlAssignment        string `json:"InspPlanMatlAssignment"`
	InspPlanMatlAssgmtIntVersion  string `json:"InspPlanMatlAssgmtIntVersion"`
	InspectionPlanInternalVersion string `json:"InspectionPlanInternalVersion"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	ChangeNumber                  string `json:"ChangeNumber"`
	CreationDate                  string `json:"CreationDate"`
	LastChangeDate                string `json:"LastChangeDate"`
	IsDeleted                     bool   `json:"IsDeleted"`
	Supplier                      string `json:"Supplier"`
	Customer                      string `json:"Customer"`
	MultipleSpecificationObject   string `json:"MultipleSpecificationObject"`
	MultipleSpecificationObjType  string `json:"MultipleSpecificationObjType"`
	BOOSearchText                 string `json:"BOOSearchText"`
	ChangedDateTime               string `json:"ChangedDateTime"`
}

type Operation struct {
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
}
