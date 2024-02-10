package index

import (
	"github.com/dunhamsteve/iwork/proto/KN"
	"github.com/dunhamsteve/iwork/proto/TSWP"
	"github.com/golang/protobuf/proto"
)

func decodeKeynote(typ uint32, payload []byte) (interface{}, error) {
	switch typ {
	case 1:
		var value = &KN.DocumentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10:
		var value = &KN.ThemeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 100:
		var value = &KN.CommandBuildSetValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 10011:
		var value = &TSWP.SectionPlaceholderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 101:
		var value = &KN.CommandShowInsertSlideArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 102:
		var value = &KN.CommandShowMoveSlideArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 103:
		var value = &KN.CommandShowRemoveSlideArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 104:
		var value = &KN.CommandSlideInsertDrawablesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 105:
		var value = &KN.CommandSlideRemoveDrawableArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 106:
		var value = &KN.CommandSlideNodeSetPropertyArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 107:
		var value = &KN.CommandSlideInsertBuildArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 108:
		var value = &KN.CommandSlideMoveBuildWithoutMovingChunksArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 109:
		var value = &KN.CommandSlideRemoveBuildArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 11:
		var value = &KN.PasteboardNativeStorageArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 110:
		var value = &KN.CommandSlideInsertBuildChunkArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 111:
		var value = &KN.CommandSlideMoveBuildChunkArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 112:
		var value = &KN.CommandSlideRemoveBuildChunkArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 113:
		var value = &KN.CommandSlideSetValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 114:
		var value = &KN.CommandTransitionSetValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 115:
		var value = &KN.UIStateCommandGroupArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 116:
		var value = &KN.CommandSlidePasteDrawablesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 117:
		var value = &KN.CommandSlideApplyThemeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 118:
		var value = &KN.CommandSlideMoveDrawableZOrderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 119:
		var value = &KN.CommandChangeMasterSlideArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 12:
		var value = &KN.PlaceholderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 123:
		var value = &KN.CommandShowSetSlideNumberVisibilityArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 124:
		var value = &KN.CommandShowSetValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 128:
		var value = &KN.CommandShowMarkOutOfSyncRecordingArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 129:
		var value = &KN.CommandShowRemoveRecordingArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 130:
		var value = &KN.CommandShowReplaceRecordingArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 131:
		var value = &KN.CommandShowSetSoundtrack{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 132:
		var value = &KN.CommandSoundtrackSetValue{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 133:
		var value = &KN.CommandMasterRescaleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 134:
		var value = &KN.CommandMoveMastersArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 135:
		var value = &KN.CommandInsertMasterArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 136:
		var value = &KN.CommandSlideSetStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 137:
		var value = &KN.CommandSlideSetPlaceholdersForTagsArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 138:
		var value = &KN.CommandBuildChunkSetValueArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 139:
		var value = &KN.CommandSlideMoveBuildChunksArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 14:
		var value = &TSWP.TextualAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 140:
		var value = &KN.CommandRemoveMasterArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 141:
		var value = &KN.CommandRenameMasterArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 142:
		var value = &KN.CommandMasterSetThumbnailTextArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 143:
		var value = &KN.CommandShowChangeThemeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 144:
		var value = &KN.CommandSlidePrimitiveSetMasterArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 145:
		var value = &KN.CommandMasterSetBodyStylesArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 146:
		var value = &KN.CommandSlideReapplyMasterArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 147:
		var value = &KN.SlideCollectionCommandSelectionBehaviorArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 148:
		var value = &KN.ChartInfoGeometryCommandArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 15:
		var value = &KN.NoteArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 16:
		var value = &KN.RecordingArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 17:
		var value = &KN.RecordingEventTrackArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 18:
		var value = &KN.RecordingMovieTrackArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 19:
		var value = &KN.ClassicStylesheetRecordArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 2:
		var value = &KN.ShowArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 20:
		var value = &KN.ClassicThemeRecordArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 21:
		var value = &KN.Soundtrack{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 22:
		var value = &KN.SlideNumberAttachmentArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 23:
		var value = &KN.DesktopUILayoutArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 24:
		var value = &KN.CanvasSelectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 25:
		var value = &KN.SlideCollectionSelectionArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 3:
		var value = &KN.UIStateArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 4:
		var value = &KN.SlideNodeArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 5:
		var value = &KN.SlideArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 6:
		var value = &KN.SlideArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 7:
		var value = &KN.PlaceholderArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 8:
		var value = &KN.BuildArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	case 9:
		var value = &KN.SlideStyleArchive{}
		err := proto.Unmarshal(payload, value)
		return value, err

	default:
		return decodeCommon(typ, payload)
	}
}
