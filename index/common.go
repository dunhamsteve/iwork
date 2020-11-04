package index

import (
	"fmt"

	"github.com/dunhamsteve/iwork/proto/TSA"
	"github.com/dunhamsteve/iwork/proto/TSCE"
	"github.com/dunhamsteve/iwork/proto/TSCH"
	"github.com/dunhamsteve/iwork/proto/TSCH/PreUFF"
	"github.com/dunhamsteve/iwork/proto/TSD"
	"github.com/dunhamsteve/iwork/proto/TSK"
	"github.com/dunhamsteve/iwork/proto/TSP"
	"github.com/dunhamsteve/iwork/proto/TSS"
	"github.com/dunhamsteve/iwork/proto/TST"
	"github.com/dunhamsteve/iwork/proto/TSWP"

	"github.com/golang/protobuf/proto"
)

func decodeCommon(typ uint32, payload []byte) (interface{}, error) {
	switch typ {

	case 11000:
		var value = &TSP.PasteboardObject{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 11006:
		var value = &TSP.PackageMetadata{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 11007:
		var value = &TSP.PasteboardMetadata{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 11008:
		var value = &TSP.ObjectContainer{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 200:
		var value = &TSK.DocumentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2001:
		var value = &TSWP.StorageArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2002:
		var value = &TSWP.SelectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2003:
		var value = &TSWP.DrawableAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2004:
		var value = &TSWP.TextualAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2005:
		var value = &TSWP.StorageArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2006:
		var value = &TSWP.UIGraphicalAttachment{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2007:
		var value = &TSWP.TextualAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2008:
		var value = &TSWP.FootnoteReferenceAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2009:
		var value = &TSWP.TextualAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 201:
		var value = &TSK.CommandHistory{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2010:
		var value = &TSWP.TSWPTOCPageNumberAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2011:
		var value = &TSWP.ShapeInfoArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2013:
		var value = &TSWP.HighlightArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2014:
		var value = &TSWP.CommentInfoArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 202:
		var value = &TSK.CommandGroupArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2021:
		var value = &TSWP.CharacterStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2022:
		var value = &TSWP.ParagraphStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2023:
		var value = &TSWP.ListStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2024:
		var value = &TSWP.ColumnStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2025:
		var value = &TSWP.ShapeStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2026:
		var value = &TSWP.TOCEntryStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 203:
		var value = &TSK.CommandContainerArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2031:
		var value = &TSWP.PlaceholderSmartFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2032:
		var value = &TSWP.HyperlinkFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2033:
		var value = &TSWP.FilenameSmartFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2034:
		var value = &TSWP.DateTimeSmartFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2035:
		var value = &TSWP.BookmarkFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2036:
		var value = &TSWP.MergeSmartFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2037:
		var value = &TSWP.CitationRecordArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2038:
		var value = &TSWP.CitationSmartFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2039:
		var value = &TSWP.UnsupportedHyperlinkFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 204:
		var value = &TSK.ReplaceAllCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2040:
		var value = &TSWP.BibliographySmartFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2041:
		var value = &TSWP.TOCSmartFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2042:
		var value = &TSWP.RubyFieldArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2043:
		var value = &TSWP.NumberAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 205:
		var value = &TSK.TreeNode{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2050:
		var value = &TSWP.TextStylePresetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2051:
		var value = &TSWP.TOCSettingsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2052:
		var value = &TSWP.TOCEntryInstanceArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 206:
		var value = &TSK.ProgressiveCommandGroupArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2060:
		var value = &TSWP.ChangeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2061:
		var value = &TSK.DeprecatedChangeAuthorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2062:
		var value = &TSWP.ChangeSessionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 208:
		var value = &TSK.CommandSelectionBehaviorHistoryArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 209:
		var value = &TSK.UndoRedoStateCommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 210:
		var value = &TSK.ViewStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2101:
		var value = &TSWP.TextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2102:
		var value = &TSWP.InsertAttachmentCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2104:
		var value = &TSWP.ReplaceAllTextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2105:
		var value = &TSWP.FormatTextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2107:
		var value = &TSWP.ApplyPlaceholderTextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2108:
		var value = &TSWP.ApplyHighlightTextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 211:
		var value = &TSK.DocumentSupportArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2113:
		var value = &TSWP.CreateHyperlinkCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2114:
		var value = &TSWP.RemoveHyperlinkCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2115:
		var value = &TSWP.ModifyHyperlinkCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2116:
		var value = &TSWP.ApplyRubyTextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2117:
		var value = &TSWP.RemoveRubyTextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2118:
		var value = &TSWP.ModifyRubyTextCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2119:
		var value = &TSWP.UpdateDateTimeFieldCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 212:
		var value = &TSK.AnnotationAuthorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2120:
		var value = &TSWP.ModifyTOCSettingsBaseCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2121:
		var value = &TSWP.ModifyTOCSettingsForTOCInfoCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2122:
		var value = &TSWP.ModifyTOCSettingsPresetForThemeCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 213:
		var value = &TSK.AnnotationAuthorStorageArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 214:
		var value = &TSK.AddAnnotationAuthorCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 215:
		var value = &TSK.SetAnnotationAuthorColorCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2206:
		var value = &TSWP.AnchorAttachmentCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2207:
		var value = &TSWP.TextApplyThemeCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2231:
		var value = &TSWP.ShapeApplyPresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2232:
		var value = &TSWP.ShapePasteStyleCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2240:
		var value = &TSWP.TOCInfoArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2241:
		var value = &TSWP.TOCAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2242:
		var value = &TSWP.TOCLayoutHintArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2400:
		var value = &TSWP.StyleBaseCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2401:
		var value = &TSWP.StyleCreateCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2402:
		var value = &TSWP.StyleRenameCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2403:
		var value = &TSWP.StyleUpdateCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2404:
		var value = &TSWP.StyleDeleteCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2405:
		var value = &TSWP.StyleReorderCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2406:
		var value = &TSWP.StyleUpdatePropertyMapCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3002:
		var value = &TSD.DrawableArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3003:
		var value = &TSD.ContainerArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3004:
		var value = &TSD.ShapeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3005:
		var value = &TSD.ImageArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3006:
		var value = &TSD.MaskArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3007:
		var value = &TSD.MovieArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3008:
		var value = &TSD.GroupArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3009:
		var value = &TSD.ConnectionLineArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3015:
		var value = &TSD.ShapeStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3016:
		var value = &TSD.MediaStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3020:
		var value = &TSD.DrawablesCommandGroupArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3021:
		var value = &TSD.InfoGeometryCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3022:
		var value = &TSD.DrawablePathSourceCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3023:
		var value = &TSD.ShapePathSourceFlipCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3024:
		var value = &TSD.ImageMaskCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3025:
		var value = &TSD.ImageMediaCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3026:
		var value = &TSD.ImageReplaceCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3027:
		var value = &TSD.MediaOriginalSizeCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3028:
		var value = &TSD.ShapeStyleSetValueCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3030:
		var value = &TSD.MediaStyleSetValueCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3031:
		var value = &TSD.ShapeApplyPresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3032:
		var value = &TSD.MediaApplyPresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3033:
		var value = &TSD.DrawableApplyThemeCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3034:
		var value = &TSD.MovieSetValueCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3035:
		var value = &TSD.ShapeSetLineEndCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3036:
		var value = &TSD.ExteriorTextWrapCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3037:
		var value = &TSD.MediaFlagsCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3038:
		var value = &TSD.GroupDrawablesCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3039:
		var value = &TSD.UngroupGroupCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3040:
		var value = &TSD.DrawableHyperlinkCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3041:
		var value = &TSD.ConnectionLineConnectCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3042:
		var value = &TSD.InstantAlphaCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3043:
		var value = &TSD.DrawableLockCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3045:
		var value = &TSD.CanvasSelectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3046:
		var value = &TSD.CommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3047:
		var value = &TSD.GuideStorageArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3048:
		var value = &TSD.StyledInfoSetStyleCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3049:
		var value = &TSD.DrawableInfoCommentCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3050:
		var value = &TSD.GuideCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3051:
		var value = &TSD.DrawableAspectRatioLockedCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3052:
		var value = &TSD.ContainerRemoveChildrenCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3053:
		var value = &TSD.ContainerInsertChildrenCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3054:
		var value = &TSD.ContainerReorderChildrenCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3055:
		var value = &TSD.ImageAdjustmentsCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3056:
		var value = &TSD.CommentStorageArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3057:
		var value = &TSD.ThemeReplaceFillPresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3058:
		var value = &TSD.DrawableAccessibilityDescriptionCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3059:
		var value = &TSD.PasteStyleCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3060:
		var value = &TSD.CommentStorageApplyCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 400:
		var value = &TSS.StyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 4000:
		var value = &TSCE.CalculationEngineArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 4001:
		var value = &TSCE.FormulaRewriteCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 4002:
		var value = &TSCE.TrackedReferencesRewriteCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 4003:
		var value = &TSCE.NamedReferenceManagerArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 4004:
		var value = &TSCE.ReferenceTrackerArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 4005:
		var value = &TSCE.TrackedReferenceArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 401:
		var value = &TSS.StylesheetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 402:
		var value = &TSS.ThemeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 410:
		var value = &TSS.ApplyThemeCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 411:
		var value = &TSS.ApplyThemeChildCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 412:
		var value = &TSS.StyleUpdatePropertyMapCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 413:
		var value = &TSS.ThemeReplacePresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 414:
		var value = &TSS.ThemeAddStylePresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 415:
		var value = &TSS.ThemeRemoveStylePresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 416:
		var value = &TSS.ThemeReplaceColorPresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 417:
		var value = &TSS.ThemeMovePresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 418:
		var value = &TSS.ThemeReplaceStylePresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5000:
		var value = &PreUFF.ChartInfoArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5002:
		var value = &PreUFF.ChartGridArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5004:
		var value = &TSCH.ChartMediatorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5010:
		var value = &PreUFF.ChartStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5011:
		var value = &PreUFF.ChartSeriesStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5012:
		var value = &PreUFF.ChartAxisStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5013:
		var value = &PreUFF.LegendStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5014:
		var value = &PreUFF.ChartNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5015:
		var value = &PreUFF.ChartSeriesNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5016:
		var value = &PreUFF.ChartAxisNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5017:
		var value = &PreUFF.LegendNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5020:
		var value = &TSCH.ChartStylePreset{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5021:
		var value = &TSCH.ChartDrawableArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5022:
		var value = &TSCH.ChartStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5023:
		var value = &TSCH.ChartNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5024:
		var value = &TSCH.LegendStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5025:
		var value = &TSCH.LegendNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5026:
		var value = &TSCH.ChartAxisStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5027:
		var value = &TSCH.ChartAxisNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5028:
		var value = &TSCH.ChartSeriesStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5029:
		var value = &TSCH.ChartSeriesNonStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5103:
		var value = &TSCH.CommandSetChartTypeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5104:
		var value = &TSCH.CommandSetSeriesNameArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5105:
		var value = &TSCH.CommandSetCategoryNameArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5107:
		var value = &TSCH.CommandSetScatterFormatArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5108:
		var value = &TSCH.CommandSetLegendFrameArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5109:
		var value = &TSCH.CommandSetGridValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5110:
		var value = &TSCH.CommandSetGridDirectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5113:
		var value = &TSCH.SynchronousCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5114:
		var value = &TSCH.CommandReplaceAllArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5115:
		var value = &TSCH.CommandAddGridRowsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5116:
		var value = &TSCH.CommandAddGridColumnsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5117:
		var value = &TSCH.CommandSetPreviewLocArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5118:
		var value = &TSCH.CommandMoveGridRowsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5119:
		var value = &TSCH.CommandMoveGridColumnsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5120:
		var value = &TSCH.CommandDeleteGridRowsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5121:
		var value = &TSCH.CommandDeleteGridColumnsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5122:
		var value = &TSCH.CommandSetPieWedgeExplosion{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5123:
		var value = &TSCH.CommandStyleSwapArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5124:
		var value = &TSCH.CommandChartApplyTheme{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5125:
		var value = &TSCH.CommandChartApplyPreset{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5126:
		var value = &TSCH.ChartCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5127:
		var value = &TSCH.CommandReplaceGridValuesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5129:
		var value = &TSCH.StylePasteboardDataArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5130:
		var value = &TSCH.CommandSetMultiDataSetIndexArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5131:
		var value = &TSCH.CommandReplaceThemePresetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5132:
		var value = &TSCH.CommandInvalidateWPCaches{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 600:
		var value = &TSA.DocumentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6000:
		var value = &TST.TableInfoArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6001:
		var value = &TST.TableModelArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6002:
		var value = &TST.Tile{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6003:
		var value = &TST.TableStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6004:
		var value = &TST.CellStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6005:
		var value = &TST.TableDataList{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6006:
		var value = &TST.HeaderStorageBucket{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6007:
		var value = &TST.WPTableInfoArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6008:
		var value = &TST.TableStylePresetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6009:
		var value = &TST.TableStrokePresetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 601:
		var value = &TSA.FunctionBrowserStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6010:
		var value = &TST.ConditionalStyleSetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 602:
		var value = &TSA.PropagatePresetCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6100:
		var value = &TST.TableCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6101:
		var value = &TST.CommandDeleteCellsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6102:
		var value = &TST.CommandInsertColumnsOrRowsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6103:
		var value = &TST.CommandRemoveColumnsOrRowsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6104:
		var value = &TST.CommandResizeColumnOrRowArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6105:
		var value = &TST.CommandSetCellArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6106:
		var value = &TST.CommandSetNumberOfHeadersOrFootersArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6107:
		var value = &TST.CommandSetTableNameArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6108:
		var value = &TST.CommandStyleCellsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6109:
		var value = &TST.CommandFillCellsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6110:
		var value = &TST.CommandReplaceAllTextArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6111:
		var value = &TST.CommandChangeFreezeHeaderStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6112:
		var value = &TST.CommandReplaceTextArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6113:
		var value = &TST.CommandPasteArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6114:
		var value = &TST.CommandSetTableNameEnabledArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6115:
		var value = &TST.CommandMoveRowsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6116:
		var value = &TST.CommandMoveColumnsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6117:
		var value = &TST.CommandApplyTableStylePresetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6118:
		var value = &TST.CommandApplyStrokePresetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6119:
		var value = &TST.CommandSetExplicitFormatArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6120:
		var value = &TST.CommandSetRepeatingHeaderEnabledArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6121:
		var value = &TST.CommandApplyThemeToTableArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6122:
		var value = &TST.CommandApplyThemeChildForTableArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6123:
		var value = &TST.CommandSortArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6124:
		var value = &TST.CommandToggleTextPropertyArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6125:
		var value = &TST.CommandStyleTableArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6126:
		var value = &TST.CommandSetNumberOfDecimalPlacesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6127:
		var value = &TST.CommandSetShowThousandsSeparatorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6128:
		var value = &TST.CommandSetNegativeNumberStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6129:
		var value = &TST.CommandSetFractionAccuracyArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6130:
		var value = &TST.CommandSetSingleNumberFormatParameterArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6131:
		var value = &TST.CommandSetCurrencyCodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6132:
		var value = &TST.CommandSetUseAccountingStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6134:
		var value = &TST.CommandRewriteFormulasForSortArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6135:
		var value = &TST.CommandRewriteFormulasForTectonicShiftArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6136:
		var value = &TST.CommandSetTableFontNameArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6137:
		var value = &TST.CommandSetTableFontSizeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6138:
		var value = &TST.CommandRewriteFormulasForMoveArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6139:
		var value = &TST.CommandFixStylesInHeadersOrFootersArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6141:
		var value = &TST.CommandResetFillPropertyToDefault{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6142:
		var value = &TST.CommandSetTableNameHeightArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6143:
		var value = &TST.CommandMergeUnmergeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6144:
		var value = &TST.MergeRegionMapArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6145:
		var value = &TST.CommandHideShowArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6146:
		var value = &TST.CommandSetBaseArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6147:
		var value = &TST.CommandSetBasePlacesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6148:
		var value = &TST.CommandSetBaseUseMinusSignArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6179:
		var value = &TST.FormulaEqualsTokenAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6181:
		var value = &TST.TokenAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6182:
		var value = &TST.ExpressionNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6183:
		var value = &TST.BooleanNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6184:
		var value = &TST.NumberNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6185:
		var value = &TST.StringNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6186:
		var value = &TST.ArrayNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6187:
		var value = &TST.ListNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6188:
		var value = &TST.OperatorNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6189:
		var value = &TST.FunctionNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6190:
		var value = &TST.DateNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6191:
		var value = &TST.ReferenceNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6192:
		var value = &TST.DurationNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6193:
		var value = &TST.ArgumentPlaceholderNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6194:
		var value = &TST.PostfixOperatorNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6195:
		var value = &TST.PrefixOperatorNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6196:
		var value = &TST.FunctionEndNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6197:
		var value = &TST.EmptyExpressionNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6198:
		var value = &TST.LayoutHintArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6199:
		var value = &TST.CompletionTokenAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6200:
		var value = &TST.FormulaEditingCommandGroupArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6201:
		var value = &TST.TableDataList{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6202:
		var value = &TST.CommandCoerceMultipleCellsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6203:
		var value = &TST.CommandSetMultipleCellsCustomArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6204:
		var value = &TST.HiddenStateFormulaOwnerArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6205:
		var value = &TST.CommandSetAutomaticDurationUnitsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6206:
		var value = &TST.PopUpMenuModel{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6207:
		var value = &TST.CommandSetControlMinimumArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6208:
		var value = &TST.CommandSetControlMaximumArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6209:
		var value = &TST.CommandSetControlIncrementArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6210:
		var value = &TST.CommandSetControlCellsDisplayNumberFormatArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6211:
		var value = &TST.CommandSetMultipleCellsMultipleChoiceListArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6212:
		var value = &TST.CommandSetMultipleChoiceListFormatForEditedItemArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6213:
		var value = &TST.CommandSetMultipleChoiceListFormatForDeleteItemArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6214:
		var value = &TST.CommandSetMultipleChoiceListFormatForReorderItemArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6215:
		var value = &TST.CommandSetMultipleChoiceListFormatForInitialValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6216:
		var value = &TST.CommandRewriteFormulasForCellMergeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6217:
		var value = &TST.TableInfoGeometryCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6218:
		var value = &TST.RichTextPayloadArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6219:
		var value = &TST.EditingStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6220:
		var value = &TST.FilterSetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6221:
		var value = &TST.CommandSetFiltersEnabledArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6222:
		var value = &TST.CommandRewriteFilterFormulasForTectonicShiftArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6223:
		var value = &TST.CommandRewriteFilterFormulasForSortArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6224:
		var value = &TST.CommandRewriteFilterFormulasForTableResizeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6225:
		var value = &TST.CommandSetAutomaticFormatArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6226:
		var value = &TST.CommandTextPreflightInsertCellArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6227:
		var value = &TST.FormulaEditingCommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6228:
		var value = &TST.CommandDeleteCellContentsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6229:
		var value = &TST.CommandPostflightSetCellArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6231:
		var value = &TST.CommandRewriteConditionalStylesForTectonicShiftArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6232:
		var value = &TST.CommandRewriteConditionalStylesForSortArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6233:
		var value = &TST.CommandRewriteConditionalStylesForRangeMoveArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6234:
		var value = &TST.CommandRewriteConditionalStylesForCellMergeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6235:
		var value = &TST.IdentifierNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6236:
		var value = &TST.UndoRedoStateCommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6237:
		var value = &TST.CommandSetStyleApplyClearsAllFlagArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6238:
		var value = &TST.CommandSetDateTimeFormatArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6239:
		var value = &TST.TableCommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6240:
		var value = &TST.CommandAddQuickFilterRulesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6241:
		var value = &TST.CommandModifyFilterRuleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6242:
		var value = &TST.CommandDeleteFilterRulesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6244:
		var value = &TST.CommandApplyCellCommentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6245:
		var value = &TST.CommandApplyConditionalStyleSetArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6246:
		var value = &TST.CommandSetFormulaTokenizationArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6247:
		var value = &TST.TableStyleNetworkArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6248:
		var value = &TST.CommandSetFilterEnabledArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6249:
		var value = &TST.CommandSetFilterRuleEnabledArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6250:
		var value = &TST.CommandSetFilterSetTypeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6251:
		var value = &TST.CommandSetStyleNetworkArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6252:
		var value = &TST.CommandMutateCellsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6253:
		var value = &TST.DisableTableNameSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6254:
		var value = &TST.CommandDisableFilterRulesForColumnArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6255:
		var value = &TST.CommandSetTextStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6256:
		var value = &TST.CommandNotifyForTransformingArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	default:
		return nil, fmt.Errorf("Unknown type %d", typ)
	}
}
