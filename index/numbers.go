package index

import (
	"github.com/dunhamsteve/iwork/proto/TN"
	"github.com/dunhamsteve/iwork/proto/TSWP"
	"github.com/golang/protobuf/proto"
)

func decodeNumbers(typ uint32, payload []byte) (interface{}, error) {
	switch typ {

	case 1:
		var value = &TN.DocumentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10011:
		var value = &TSWP.SectionPlaceholderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12002:
		var value = &TN.CommandSheetInsertDrawablesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12003:
		var value = &TN.CommandDocumentInsertSheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12004:
		var value = &TN.CommandDocumentRemoveSheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12005:
		var value = &TN.CommandSetSheetNameArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12006:
		var value = &TN.ChartMediatorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12007:
		var value = &TN.CommandPasteDrawablesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12008:
		var value = &TN.CommandDocumentReorderSheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12009:
		var value = &TN.ThemeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12010:
		var value = &TN.CommandPasteSheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12011:
		var value = &TN.CommandReorderSidebarItemChildrenAchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12012:
		var value = &TN.CommandSheetRemoveDrawablesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12013:
		var value = &TN.CommandSheetMoveDrawableZOrderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12014:
		var value = &TN.CommandChartMediatorSetEditingState{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12015:
		var value = &TN.CommandFormChooseTargetTableArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12016:
		var value = &TN.CommandChartMediatorUpdateForEntityDelete{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12017:
		var value = &TN.CommandSetPageOrientationArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12018:
		var value = &TN.CommandSetContentScaleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12019:
		var value = &TN.CommandSetShowPageNumbersValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12021:
		var value = &TN.CommandSetAutofitValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12024:
		var value = &TN.UndoRedoStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12025:
		var value = &TN.CommandDocumentReplaceLastSheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12026:
		var value = &TN.UIStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12027:
		var value = &TN.ChartCommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12028:
		var value = &TN.SheetSelectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12029:
		var value = &TN.SheetCommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12030:
		var value = &TN.CommandSetDocumentPrinterOptions{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2:
		var value = &TN.SheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3:
		var value = &TN.FormBasedSheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 7:
		var value = &TN.PlaceholderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	default:
		return decodeCommon(typ, payload)
	}
}
