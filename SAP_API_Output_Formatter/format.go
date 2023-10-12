package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-inspection-lot-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			InspectionLot:                  data.InspectionLot,
			Material:                       data.Material,
			Batch:                          data.Batch,
			Plant:                          data.Plant,
			InspectionLotOrigin:            data.InspectionLotOrigin,
			OrderInternalBillOfOperations:  data.OrderInternalBillOfOperations,
			ManufacturingOrder:             data.ManufacturingOrder,
			InspectionLotText:              data.InspectionLotText,
			InspectionLotType:              data.InspectionLotType,
			InspectionLotQuantity:          data.InspectionLotQuantity,
			InspectionLotActualQuantity:    data.InspectionLotActualQuantity,
			InspectionLotDefectiveQuantity: data.InspectionLotDefectiveQuantity,
			InspectionLotQuantityUnit:      data.InspectionLotQuantityUnit,
			InspLotCreatedOnLocalDate:      data.InspLotCreatedOnLocalDate,
			InspSubsetFieldCombination:     data.InspSubsetFieldCombination,
			InspLotNmbrOpenLongTermCharc:   data.InspLotNmbrOpenLongTermCharc,
			StatusObject:                   data.StatusObject,
			StatusObjectCategory:           data.StatusObjectCategory,
			InspectionLotObjectText:        data.InspectionLotObjectText,
			StatusProfile:                  data.StatusProfile,
			MatlQualityAuthorizationGroup:  data.MatlQualityAuthorizationGroup,
			InspectionLotHasQuantity:       data.InspectionLotHasQuantity,
			InspLotIsCreatedAutomatically:  data.InspLotIsCreatedAutomatically,
			InspectionLotHasPartialLots:    data.InspectionLotHasPartialLots,
			InspectionLotHasAppraisalCosts: data.InspectionLotHasAppraisalCosts,
			InspLotHasSubsets:              data.InspLotHasSubsets,
			InspLotIsAutomUsgeDcsnPossible: data.InspLotIsAutomUsgeDcsnPossible,
			PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
			InspLotHasConfignSpecification: data.InspLotHasConfignSpecification,
			GoodsReceiptIsMovedToBlkdStock: data.GoodsReceiptIsMovedToBlkdStock,
			InspLotIsDocumentationRequired: data.InspLotIsDocumentationRequired,
			InspLotIsTaskListRequired:      data.InspLotIsTaskListRequired,
			InspLotHasManualSampleSize:     data.InspLotHasManualSampleSize,
			InspLotHasMaterialSpec:         data.InspLotHasMaterialSpec,
			InspectionLotApproval:          data.InspectionLotApproval,
			InspLotDigitalSgntrResultsRecg: data.InspLotDigitalSgntrResultsRecg,
			InspLotDigitalSgntrInUsgeDcsn:  data.InspLotDigitalSgntrInUsgeDcsn,
			InspLotIsBatchRequired:         data.InspLotIsBatchRequired,
			InspLotUsageInStabilityStudy:   data.InspLotUsageInStabilityStudy,
			InspLotIsStockPostingCompleted: data.InspLotIsStockPostingCompleted,
			InspectionLotHasUsageDecision:  data.InspectionLotHasUsageDecision,
			NumberOfSerialNumbers:          data.NumberOfSerialNumbers,
			InspLotIsSerialNmbrPossible:    data.InspLotIsSerialNmbrPossible,
			InspectionLotIsSkipped:         data.InspectionLotIsSkipped,
			InspectionLotIsFullInspection:  data.InspectionLotIsFullInspection,
			InspectionLotDynamicLevel:      data.InspectionLotDynamicLevel,
			SamplingProcedure:              data.SamplingProcedure,
			InspLotCreatedOnLocalTime:      data.InspLotCreatedOnLocalTime,
			InspectionLotCreatedBy:         data.InspectionLotCreatedBy,
			InspectionLotCreatedOn:         data.InspectionLotCreatedOn,
			InspectionLotCreatedOnTime:     data.InspectionLotCreatedOnTime,
			InspectionLotChangedBy:         data.InspectionLotChangedBy,
			InspectionLotChangeDate:        data.InspectionLotChangeDate,
			InspectionLotChangeTime:        data.InspectionLotChangeTime,
			ChangedDateTime:                data.ChangedDateTime,
			InspectionLotStartDate:         data.InspectionLotStartDate,
			InspectionLotStartTime:         data.InspectionLotStartTime,
			InspectionLotEndDate:           data.InspectionLotEndDate,
			InspectionLotEndTime:           data.InspectionLotEndTime,
			BillOfOperationsType:           data.BillOfOperationsType,
			BillOfOperationsGroup:          data.BillOfOperationsGroup,
			BillOfOperationsUsage:          data.BillOfOperationsUsage,
			BillOfOperationsVariant:        data.BillOfOperationsVariant,
			BillOfOperationsChangeStateID:  data.BillOfOperationsChangeStateID,
			InspectionSubsetType:           data.InspectionSubsetType,
			SmplDrawingProcedure:           data.SmplDrawingProcedure,
			SmplDrawingProcedureVersion:    data.SmplDrawingProcedureVersion,
			SmplDrwgProcedIsConfRequired:   data.SmplDrwgProcedIsConfRequired,
			InspLotSelectionMaterial:       data.InspLotSelectionMaterial,
			InspLotSelMatlRevisionLvl:      data.InspLotSelMatlRevisionLvl,
			InspLotSelectionPlant:          data.InspLotSelectionPlant,
			InspLotSelectionSupplier:       data.InspLotSelectionSupplier,
			InspLotSelectionManufacturer:   data.InspLotSelectionManufacturer,
			InspLotSelectionCustomer:       data.InspLotSelectionCustomer,
			InspLotSelBillOfOperationsUsge: data.InspLotSelBillOfOperationsUsge,
			InspLotSelectionValidFromDate:  data.InspLotSelectionValidFromDate,
			ProductionVersion:              data.ProductionVersion,
			SalesOperationsPlanningOrder:   data.SalesOperationsPlanningOrder,
			IsBusinessPurposeCompleted:     data.IsBusinessPurposeCompleted,
			Customer:                       data.Customer,
			Supplier:                       data.Supplier,
			BatchBySupplier:                data.BatchBySupplier,
			Manufacturer:                   data.Manufacturer,
			ManufacturerPartNmbr:           data.ManufacturerPartNmbr,
			MaterialRevisionLevel:          data.MaterialRevisionLevel,
			MaterialIsBatchManaged:         data.MaterialIsBatchManaged,
			BatchStorageLocation:           data.BatchStorageLocation,
			MaterialCompIsSpecialStock:     data.MaterialCompIsSpecialStock,
			PurchasingOrganization:         data.PurchasingOrganization,
			PurchasingDocument:             data.PurchasingDocument,
			PurchasingDocumentItem:         data.PurchasingDocumentItem,
			ScheduleLine:                   data.ScheduleLine,
			AccountingDocumentType:         data.AccountingDocumentType,
			MaterialDocumentYear:           data.MaterialDocumentYear,
			MaterialDocument:               data.MaterialDocument,
			MaterialDocumentItem:           data.MaterialDocumentItem,
			MatlDocLatestPostgDate:         data.MatlDocLatestPostgDate,
			GoodsMovementType:              data.GoodsMovementType,
			InspectionLotPlant:             data.InspectionLotPlant,
			InspectionLotStorageLocation:   data.InspectionLotStorageLocation,
			Warehouse:                      data.Warehouse,
			StorageType:                    data.StorageType,
			StorageBin:                     data.StorageBin,
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			DeliveryDocument:               data.DeliveryDocument,
			DeliveryDocumentItem:           data.DeliveryDocumentItem,
			DeliveryCategory:               data.DeliveryCategory,
			InspectionDeliveryCategory:     data.InspectionDeliveryCategory,
			Route:                          data.Route,
			BillToPartyCountry:             data.BillToPartyCountry,
			SoldToParty:                    data.SoldToParty,
			SalesOrganization:              data.SalesOrganization,
			MaterialByCustomer:             data.MaterialByCustomer,
			Language:                       data.Language,
			InspLotNmbrAddlRecordedCharc:   data.InspLotNmbrAddlRecordedCharc,
			InspLotNmbrOpenShortTermCharc:  data.InspLotNmbrOpenShortTermCharc,
			InspectionLotContainer:         data.InspectionLotContainer,
			InspectionLotContainerUnit:     data.InspectionLotContainerUnit,
			InspectionLotSampleQuantity:    data.InspectionLotSampleQuantity,
			InspectionLotSampleUnit:        data.InspectionLotSampleUnit,
			InspLotDynamicRule:             data.InspLotDynamicRule,
			InspLotDynamicTrggrPoint:       data.InspLotDynamicTrggrPoint,
			InspectionDynamicStage:         data.InspectionDynamicStage,
			InspectionSeverity:             data.InspectionSeverity,
			InspLotQtyToFree:               data.InspLotQtyToFree,
			InspLotQtyToScrap:              data.InspLotQtyToScrap,
			InspLotQtyToSample:             data.InspLotQtyToSample,
			InspLotQtyToBlocked:            data.InspLotQtyToBlocked,
			InspLotQtyToReserves:           data.InspLotQtyToReserves,
			InspLotQtyToAnotherMaterial:    data.InspLotQtyToAnotherMaterial,
			InspLotMaterialPostedTo:        data.InspLotMaterialPostedTo,
			InspLotBatchTransferredTo:      data.InspLotBatchTransferredTo,
			InspLotQtyReturnedToSupplier:   data.InspLotQtyReturnedToSupplier,
			InspLotQtyToSpecialStock:       data.InspLotQtyToSpecialStock,
			InspLotQtyToOtherStock:         data.InspLotQtyToOtherStock,
			InspLotQtyToBePosted:           data.InspLotQtyToBePosted,
			InspLotSmplQtyForLongTermChar:  data.InspLotSmplQtyForLongTermChar,
			InspLotQtyInspected:            data.InspLotQtyInspected,
			InspLotQtyDestroyed:            data.InspLotQtyDestroyed,
			InspectionLotScrapRatio:        data.InspectionLotScrapRatio,
			InspLotUsageDecisionTool:       data.InspLotUsageDecisionTool,
			InspectionLotAllowedScrapRatio: data.InspectionLotAllowedScrapRatio,
			QualityCostCollector:           data.QualityCostCollector,
			ConsumptionPosting:             data.ConsumptionPosting,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
			InspLotAcctAssgmtKey:           data.InspLotAcctAssgmtKey,
			CostCenter:                     data.CostCenter,
			ReservationItem:                data.ReservationItem,
			MasterFixedAsset:               data.MasterFixedAsset,
			FixedAsset:                     data.FixedAsset,
			SalesOrdStockWBSElement:        data.SalesOrdStockWBSElement,
			ProjectNetwork:                 data.ProjectNetwork,
			NetworkActivityInternalID:      data.NetworkActivityInternalID,
			InventorySpclStkSalesDocument:  data.InventorySpclStkSalesDocument,
			InventorySpclStkSalesDocItm:    data.InventorySpclStkSalesDocItm,
			ProfitabilitySegment:           data.ProfitabilitySegment,
			ProfitCenter:                   data.ProfitCenter,
			BusinessArea:                   data.BusinessArea,
			GLAccount:                      data.GLAccount,
			ControllingArea:                data.ControllingArea,
			CompanyCode:                    data.CompanyCode,
			SerialNumberProfile:            data.SerialNumberProfile,
			InspLotCostCollectorSalesOrder: data.InspLotCostCollectorSalesOrder,
			InspLotCostCollectorSlsOrdItem: data.InspLotCostCollectorSlsOrdItem,
			InspLotCostCollectorWBSElement: data.InspLotCostCollectorWBSElement,
			InspLotExternalNumber:          data.InspLotExternalNumber,
			InspectionLotPriorityPoints:    data.InspectionLotPriorityPoints,
			MaintenancePlan:                data.MaintenancePlan,
			MaintenancePlanItemIntID:       data.MaintenancePlanItemIntID,
			MaintenanceStrategy:            data.MaintenanceStrategy,
			to_InspectionLotWithStatus:     data.to_InspectionLotWithStatus,
			InspectionLotStatusCreated:     data.InspectionLotStatusCreated,
			InspectionLotStatusReleased:    data.InspectionLotStatusReleased,
			InspectionLotStatusSkip:        data.InspectionLotStatusSkip,
			InspLotStatusRsltsConfirmed:    data.InspLotStatusRsltsConfirmed,
			InspLotStsDefectsRecorded:      data.InspLotStsDefectsRecorded,
			InspLotStsShrtTrmInspCmpltd:    data.InspLotStsShrtTrmInspCmpltd,
			InspLotStatusInspCompleted:     data.InspLotStatusInspCompleted,
			InspLotStatusCanceled:          data.InspLotStatusCanceled,
			InspLotStatusRepair:            data.InspLotStatusRepair,
		})
	}

	return header, nil
}

func ConvertToOperation(raw []byte, l *logger.Logger) ([]Operation, error) {
	pm := &responses.Operation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Operation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	operation := make([]Operation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		operation = append(operation, Operation{
			InspectionLot:                 data.InspectionLot,
			InspPlanOperationInternalID:   data.InspPlanOperationInternalID,
			OrderInternalBillOfOperations: data.OrderInternalBillOfOperations,
			InspectionOperation:           data.InspectionOperation,
			InspectionOperationPlant:      data.InspectionOperationPlant,
			BillOfOperationsType:          data.BillOfOperationsType,
			BOOOperationInternalID:        data.BOOOperationInternalID,
			WorkCenterInternalID:          data.WorkCenterInternalID,
			StatusObject:                  data.StatusObject,
			OperationControlProfile:       data.OperationControlProfile,
			OperationConfirmation:         data.OperationConfirmation,
			InspectionSubSystem:           data.InspectionSubSystem,
			OperationText:                 data.OperationText,
		})
	}

	return operation, nil
}
