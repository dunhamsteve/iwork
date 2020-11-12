package index

import (
	"github.com/dunhamsteve/iwork/proto/TP"

	"github.com/golang/protobuf/proto"
)

func decodePages(typ uint32, payload []byte) (interface{}, error) {
	value, err := decodeCommon(typ, payload)
	if err == nil {
		return value, err
	}

	switch typ {

	case 10000:
		var value = &TP.DocumentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10001:
		var value = &TP.ThemeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10010:
		var value = &TP.FloatingDrawablesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10011:
		var value = &TP.SectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10012:
		var value = &TP.SettingsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10015:
		var value = &TP.DrawablesZOrderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10101:
		var value = &TP.InsertDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10102:
		var value = &TP.RemoveDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10108:
		var value = &TP.PasteAnchoredDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10109:
		var value = &TP.PasteDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10110:
		var value = &TP.MoveDrawablesAttachedCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10111:
		var value = &TP.MoveDrawablesFloatingCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10112:
		var value = &TP.MoveInlineDrawableAnchoredCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10113:
		var value = &TP.InsertFootnoteCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10114:
		var value = &TP.ChangeFootnoteFormatCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10115:
		var value = &TP.ChangeFootnoteKindCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10116:
		var value = &TP.ChangeFootnoteNumberingCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10117:
		var value = &TP.ToggleBodyLayoutDirectionCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10118:
		var value = &TP.ChangeFootnoteSpacingCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10119:
		var value = &TP.MoveAnchoredDrawableInlineCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10120:
		var value = &TP.ChangeSectionMarginsCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10121:
		var value = &TP.ChangeDocumentPrinterOptionsCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10125:
		var value = &TP.InsertMasterDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10126:
		var value = &TP.RemoveMasterDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10127:
		var value = &TP.PasteMasterDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10128:
		var value = &TP.NudgeDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10130:
		var value = &TP.MoveDrawablesPageIndexCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10131:
		var value = &TP.LayoutStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10132:
		var value = &TP.CanvasSelectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10133:
		var value = &TP.ViewStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10134:
		var value = &TP.ChangeHeaderFooterVisibilityCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10140:
		var value = &TP.MoveMasterDrawableZOrderCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10141:
		var value = &TP.SwapDrawableZOrderCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10142:
		var value = &TP.RemoveAnchoredDrawableCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10143:
		var value = &TP.PageMasterArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10147:
		var value = &TP.UIStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10148:
		var value = &TP.ChangeCTVisibilityCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10149:
		var value = &TP.TrackChangesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10150:
		var value = &TP.DocumentHyphenationCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10151:
		var value = &TP.DocumentLigaturesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10152:
		var value = &TP.InsertSectionBreakCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10153:
		var value = &TP.DeleteSectionCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10154:
		var value = &TP.ReplaceSectionCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10155:
		var value = &TP.ChangeSectionPropertyCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10156:
		var value = &TP.DocumentHasBodyCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10157:
		var value = &TP.PauseChangeTrackingCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 7:
		var value = &TP.PlaceholderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	default:
		return decodeCommon(typ, payload)
	}
}
