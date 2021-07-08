package main

import "encoding/xml"

type SpreadXML struct {
	XMLName    xml.Name `xml:"Spread"`
	Text       string   `xml:",chardata"`
	IdPkg      string   `xml:"idPkg,attr"`
	DOMVersion string   `xml:"DOMVersion,attr"`
	Spread     struct {
		Text                    string `xml:",chardata"`
		Self                    string `xml:"Self,attr"`
		PageTransitionType      string `xml:"PageTransitionType,attr"`
		PageTransitionDirection string `xml:"PageTransitionDirection,attr"`
		PageTransitionDuration  string `xml:"PageTransitionDuration,attr"`
		ShowMasterItems         string `xml:"ShowMasterItems,attr"`
		PageCount               string `xml:"PageCount,attr"`
		BindingLocation         string `xml:"BindingLocation,attr"`
		AllowPageShuffle        string `xml:"AllowPageShuffle,attr"`
		ItemTransform           string `xml:"ItemTransform,attr"`
		FlattenerOverride       string `xml:"FlattenerOverride,attr"`
		FlattenerPreference     struct {
			Text                        string `xml:",chardata"`
			LineArtAndTextResolution    string `xml:"LineArtAndTextResolution,attr"`
			GradientAndMeshResolution   string `xml:"GradientAndMeshResolution,attr"`
			ClipComplexRegions          string `xml:"ClipComplexRegions,attr"`
			ConvertAllStrokesToOutlines string `xml:"ConvertAllStrokesToOutlines,attr"`
			ConvertAllTextToOutlines    string `xml:"ConvertAllTextToOutlines,attr"`
			Properties                  struct {
				Text                string `xml:",chardata"`
				RasterVectorBalance struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"RasterVectorBalance"`
			} `xml:"Properties"`
		} `xml:"FlattenerPreference"`
		Page struct {
			Text                   string `xml:",chardata"`
			Self                   string `xml:"Self,attr"`
			TabOrder               string `xml:"TabOrder,attr"`
			AppliedMaster          string `xml:"AppliedMaster,attr"`
			OverrideList           string `xml:"OverrideList,attr"`
			MasterPageTransform    string `xml:"MasterPageTransform,attr"`
			Name                   string `xml:"Name,attr"`
			AppliedTrapPreset      string `xml:"AppliedTrapPreset,attr"`
			GeometricBounds        string `xml:"GeometricBounds,attr"`
			ItemTransform          string `xml:"ItemTransform,attr"`
			AppliedAlternateLayout string `xml:"AppliedAlternateLayout,attr"`
			LayoutRule             string `xml:"LayoutRule,attr"`
			SnapshotBlendingMode   string `xml:"SnapshotBlendingMode,attr"`
			OptionalPage           string `xml:"OptionalPage,attr"`
			GridStartingPoint      string `xml:"GridStartingPoint,attr"`
			UseMasterGrid          string `xml:"UseMasterGrid,attr"`
			Properties             struct {
				Text      string `xml:",chardata"`
				PageColor struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"PageColor"`
				Descriptor struct {
					Text     string `xml:",chardata"`
					Type     string `xml:"type,attr"`
					ListItem []struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"ListItem"`
				} `xml:"Descriptor"`
			} `xml:"Properties"`
			Guide []struct {
				Text                    string `xml:",chardata"`
				Self                    string `xml:"Self,attr"`
				OverriddenPageItemProps string `xml:"OverriddenPageItemProps,attr"`
				Orientation             string `xml:"Orientation,attr"`
				Location                string `xml:"Location,attr"`
				FitToPage               string `xml:"FitToPage,attr"`
				ViewThreshold           string `xml:"ViewThreshold,attr"`
				Locked                  string `xml:"Locked,attr"`
				ItemLayer               string `xml:"ItemLayer,attr"`
				PageIndex               string `xml:"PageIndex,attr"`
				GuideType               string `xml:"GuideType,attr"`
				GuideZone               string `xml:"GuideZone,attr"`
				Properties              struct {
					Text       string `xml:",chardata"`
					GuideColor struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"GuideColor"`
				} `xml:"Properties"`
			} `xml:"Guide"`
			MarginPreference struct {
				Text             string `xml:",chardata"`
				ColumnCount      string `xml:"ColumnCount,attr"`
				ColumnGutter     string `xml:"ColumnGutter,attr"`
				Top              string `xml:"Top,attr"`
				Bottom           string `xml:"Bottom,attr"`
				Left             string `xml:"Left,attr"`
				Right            string `xml:"Right,attr"`
				ColumnDirection  string `xml:"ColumnDirection,attr"`
				ColumnsPositions string `xml:"ColumnsPositions,attr"`
			} `xml:"MarginPreference"`
			GridDataInformation struct {
				Text               string `xml:",chardata"`
				FontStyle          string `xml:"FontStyle,attr"`
				PointSize          string `xml:"PointSize,attr"`
				CharacterAki       string `xml:"CharacterAki,attr"`
				LineAki            string `xml:"LineAki,attr"`
				HorizontalScale    string `xml:"HorizontalScale,attr"`
				VerticalScale      string `xml:"VerticalScale,attr"`
				LineAlignment      string `xml:"LineAlignment,attr"`
				GridAlignment      string `xml:"GridAlignment,attr"`
				CharacterAlignment string `xml:"CharacterAlignment,attr"`
				Properties         struct {
					Text        string `xml:",chardata"`
					AppliedFont struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"AppliedFont"`
				} `xml:"Properties"`
			} `xml:"GridDataInformation"`
		} `xml:"Page"`
		Polygon []struct {
			Text                            string `xml:",chardata"`
			Self                            string `xml:"Self,attr"`
			ContentType                     string `xml:"ContentType,attr"`
			StoryTitle                      string `xml:"StoryTitle,attr"`
			OverriddenPageItemProps         string `xml:"OverriddenPageItemProps,attr"`
			Visible                         string `xml:"Visible,attr"`
			Name                            string `xml:"Name,attr"`
			HorizontalLayoutConstraints     string `xml:"HorizontalLayoutConstraints,attr"`
			VerticalLayoutConstraints       string `xml:"VerticalLayoutConstraints,attr"`
			FillColor                       string `xml:"FillColor,attr"`
			StrokeWeight                    string `xml:"StrokeWeight,attr"`
			MiterLimit                      string `xml:"MiterLimit,attr"`
			GradientFillStart               string `xml:"GradientFillStart,attr"`
			GradientFillLength              string `xml:"GradientFillLength,attr"`
			GradientFillAngle               string `xml:"GradientFillAngle,attr"`
			GradientStrokeStart             string `xml:"GradientStrokeStart,attr"`
			GradientStrokeLength            string `xml:"GradientStrokeLength,attr"`
			GradientStrokeAngle             string `xml:"GradientStrokeAngle,attr"`
			ItemLayer                       string `xml:"ItemLayer,attr"`
			Locked                          string `xml:"Locked,attr"`
			LocalDisplaySetting             string `xml:"LocalDisplaySetting,attr"`
			GradientFillHiliteLength        string `xml:"GradientFillHiliteLength,attr"`
			GradientFillHiliteAngle         string `xml:"GradientFillHiliteAngle,attr"`
			GradientStrokeHiliteLength      string `xml:"GradientStrokeHiliteLength,attr"`
			GradientStrokeHiliteAngle       string `xml:"GradientStrokeHiliteAngle,attr"`
			AppliedObjectStyle              string `xml:"AppliedObjectStyle,attr"`
			ItemTransform                   string `xml:"ItemTransform,attr"`
			ParentInterfaceChangeCount      string `xml:"ParentInterfaceChangeCount,attr"`
			TargetInterfaceChangeCount      string `xml:"TargetInterfaceChangeCount,attr"`
			LastUpdatedInterfaceChangeCount string `xml:"LastUpdatedInterfaceChangeCount,attr"`
			Properties                      struct {
				Text         string `xml:",chardata"`
				PathGeometry struct {
					Text             string `xml:",chardata"`
					GeometryPathType []struct {
						Text           string `xml:",chardata"`
						PathOpen       string `xml:"PathOpen,attr"`
						PathPointArray struct {
							Text          string `xml:",chardata"`
							PathPointType []struct {
								Text           string `xml:",chardata"`
								Anchor         string `xml:"Anchor,attr"`
								LeftDirection  string `xml:"LeftDirection,attr"`
								RightDirection string `xml:"RightDirection,attr"`
							} `xml:"PathPointType"`
						} `xml:"PathPointArray"`
					} `xml:"GeometryPathType"`
				} `xml:"PathGeometry"`
			} `xml:"Properties"`
			ObjectExportOption struct {
				Text                         string `xml:",chardata"`
				AltTextSourceType            string `xml:"AltTextSourceType,attr"`
				ActualTextSourceType         string `xml:"ActualTextSourceType,attr"`
				CustomAltText                string `xml:"CustomAltText,attr"`
				CustomActualText             string `xml:"CustomActualText,attr"`
				ApplyTagType                 string `xml:"ApplyTagType,attr"`
				ImageConversionType          string `xml:"ImageConversionType,attr"`
				ImageExportResolution        string `xml:"ImageExportResolution,attr"`
				GIFOptionsPalette            string `xml:"GIFOptionsPalette,attr"`
				GIFOptionsInterlaced         string `xml:"GIFOptionsInterlaced,attr"`
				JPEGOptionsQuality           string `xml:"JPEGOptionsQuality,attr"`
				JPEGOptionsFormat            string `xml:"JPEGOptionsFormat,attr"`
				ImageAlignment               string `xml:"ImageAlignment,attr"`
				ImageSpaceBefore             string `xml:"ImageSpaceBefore,attr"`
				ImageSpaceAfter              string `xml:"ImageSpaceAfter,attr"`
				UseImagePageBreak            string `xml:"UseImagePageBreak,attr"`
				ImagePageBreak               string `xml:"ImagePageBreak,attr"`
				CustomImageAlignment         string `xml:"CustomImageAlignment,attr"`
				SpaceUnit                    string `xml:"SpaceUnit,attr"`
				CustomLayout                 string `xml:"CustomLayout,attr"`
				CustomLayoutType             string `xml:"CustomLayoutType,attr"`
				EpubType                     string `xml:"EpubType,attr"`
				SizeType                     string `xml:"SizeType,attr"`
				CustomSize                   string `xml:"CustomSize,attr"`
				PreserveAppearanceFromLayout string `xml:"PreserveAppearanceFromLayout,attr"`
				Properties                   struct {
					Text                string `xml:",chardata"`
					AltMetadataProperty struct {
						Text            string `xml:",chardata"`
						NamespacePrefix string `xml:"NamespacePrefix,attr"`
						PropertyPath    string `xml:"PropertyPath,attr"`
					} `xml:"AltMetadataProperty"`
					ActualMetadataProperty struct {
						Text            string `xml:",chardata"`
						NamespacePrefix string `xml:"NamespacePrefix,attr"`
						PropertyPath    string `xml:"PropertyPath,attr"`
					} `xml:"ActualMetadataProperty"`
				} `xml:"Properties"`
			} `xml:"ObjectExportOption"`
			TextWrapPreference struct {
				Text                  string `xml:",chardata"`
				Inverse               string `xml:"Inverse,attr"`
				ApplyToMasterPageOnly string `xml:"ApplyToMasterPageOnly,attr"`
				TextWrapSide          string `xml:"TextWrapSide,attr"`
				TextWrapMode          string `xml:"TextWrapMode,attr"`
				Properties            struct {
					Text           string `xml:",chardata"`
					TextWrapOffset struct {
						Text   string `xml:",chardata"`
						Top    string `xml:"Top,attr"`
						Left   string `xml:"Left,attr"`
						Bottom string `xml:"Bottom,attr"`
						Right  string `xml:"Right,attr"`
					} `xml:"TextWrapOffset"`
				} `xml:"Properties"`
			} `xml:"TextWrapPreference"`
			InCopyExportOption struct {
				Text                  string `xml:",chardata"`
				IncludeGraphicProxies string `xml:"IncludeGraphicProxies,attr"`
				IncludeAllResources   string `xml:"IncludeAllResources,attr"`
			} `xml:"InCopyExportOption"`
			FrameFittingOption struct {
				Text       string `xml:",chardata"`
				LeftCrop   string `xml:"LeftCrop,attr"`
				TopCrop    string `xml:"TopCrop,attr"`
				RightCrop  string `xml:"RightCrop,attr"`
				BottomCrop string `xml:"BottomCrop,attr"`
			} `xml:"FrameFittingOption"`
		} `xml:"Polygon"`
		TextFrame []struct {
			Text                            string `xml:",chardata"`
			Self                            string `xml:"Self,attr"`
			ParentStory                     string `xml:"ParentStory,attr"`
			PreviousTextFrame               string `xml:"PreviousTextFrame,attr"`
			NextTextFrame                   string `xml:"NextTextFrame,attr"`
			ContentType                     string `xml:"ContentType,attr"`
			OverriddenPageItemProps         string `xml:"OverriddenPageItemProps,attr"`
			Visible                         string `xml:"Visible,attr"`
			Name                            string `xml:"Name,attr"`
			HorizontalLayoutConstraints     string `xml:"HorizontalLayoutConstraints,attr"`
			VerticalLayoutConstraints       string `xml:"VerticalLayoutConstraints,attr"`
			GradientFillStart               string `xml:"GradientFillStart,attr"`
			GradientFillLength              string `xml:"GradientFillLength,attr"`
			GradientFillAngle               string `xml:"GradientFillAngle,attr"`
			GradientStrokeStart             string `xml:"GradientStrokeStart,attr"`
			GradientStrokeLength            string `xml:"GradientStrokeLength,attr"`
			GradientStrokeAngle             string `xml:"GradientStrokeAngle,attr"`
			ItemLayer                       string `xml:"ItemLayer,attr"`
			Locked                          string `xml:"Locked,attr"`
			LocalDisplaySetting             string `xml:"LocalDisplaySetting,attr"`
			GradientFillHiliteLength        string `xml:"GradientFillHiliteLength,attr"`
			GradientFillHiliteAngle         string `xml:"GradientFillHiliteAngle,attr"`
			GradientStrokeHiliteLength      string `xml:"GradientStrokeHiliteLength,attr"`
			GradientStrokeHiliteAngle       string `xml:"GradientStrokeHiliteAngle,attr"`
			AppliedObjectStyle              string `xml:"AppliedObjectStyle,attr"`
			ItemTransform                   string `xml:"ItemTransform,attr"`
			ParentInterfaceChangeCount      string `xml:"ParentInterfaceChangeCount,attr"`
			TargetInterfaceChangeCount      string `xml:"TargetInterfaceChangeCount,attr"`
			LastUpdatedInterfaceChangeCount string `xml:"LastUpdatedInterfaceChangeCount,attr"`
			Properties                      struct {
				Text         string `xml:",chardata"`
				PathGeometry struct {
					Text             string `xml:",chardata"`
					GeometryPathType struct {
						Text           string `xml:",chardata"`
						PathOpen       string `xml:"PathOpen,attr"`
						PathPointArray struct {
							Text          string `xml:",chardata"`
							PathPointType []struct {
								Text           string `xml:",chardata"`
								Anchor         string `xml:"Anchor,attr"`
								LeftDirection  string `xml:"LeftDirection,attr"`
								RightDirection string `xml:"RightDirection,attr"`
							} `xml:"PathPointType"`
						} `xml:"PathPointArray"`
					} `xml:"GeometryPathType"`
				} `xml:"PathGeometry"`
			} `xml:"Properties"`
			ObjectExportOption struct {
				Text                         string `xml:",chardata"`
				AltTextSourceType            string `xml:"AltTextSourceType,attr"`
				ActualTextSourceType         string `xml:"ActualTextSourceType,attr"`
				CustomAltText                string `xml:"CustomAltText,attr"`
				CustomActualText             string `xml:"CustomActualText,attr"`
				ApplyTagType                 string `xml:"ApplyTagType,attr"`
				ImageConversionType          string `xml:"ImageConversionType,attr"`
				ImageExportResolution        string `xml:"ImageExportResolution,attr"`
				GIFOptionsPalette            string `xml:"GIFOptionsPalette,attr"`
				GIFOptionsInterlaced         string `xml:"GIFOptionsInterlaced,attr"`
				JPEGOptionsQuality           string `xml:"JPEGOptionsQuality,attr"`
				JPEGOptionsFormat            string `xml:"JPEGOptionsFormat,attr"`
				ImageAlignment               string `xml:"ImageAlignment,attr"`
				ImageSpaceBefore             string `xml:"ImageSpaceBefore,attr"`
				ImageSpaceAfter              string `xml:"ImageSpaceAfter,attr"`
				UseImagePageBreak            string `xml:"UseImagePageBreak,attr"`
				ImagePageBreak               string `xml:"ImagePageBreak,attr"`
				CustomImageAlignment         string `xml:"CustomImageAlignment,attr"`
				SpaceUnit                    string `xml:"SpaceUnit,attr"`
				CustomLayout                 string `xml:"CustomLayout,attr"`
				CustomLayoutType             string `xml:"CustomLayoutType,attr"`
				EpubType                     string `xml:"EpubType,attr"`
				SizeType                     string `xml:"SizeType,attr"`
				CustomSize                   string `xml:"CustomSize,attr"`
				PreserveAppearanceFromLayout string `xml:"PreserveAppearanceFromLayout,attr"`
				Properties                   struct {
					Text                string `xml:",chardata"`
					AltMetadataProperty struct {
						Text            string `xml:",chardata"`
						NamespacePrefix string `xml:"NamespacePrefix,attr"`
						PropertyPath    string `xml:"PropertyPath,attr"`
					} `xml:"AltMetadataProperty"`
					ActualMetadataProperty struct {
						Text            string `xml:",chardata"`
						NamespacePrefix string `xml:"NamespacePrefix,attr"`
						PropertyPath    string `xml:"PropertyPath,attr"`
					} `xml:"ActualMetadataProperty"`
				} `xml:"Properties"`
			} `xml:"ObjectExportOption"`
			TextFramePreference struct {
				Text                 string `xml:",chardata"`
				TextColumnCount      string `xml:"TextColumnCount,attr"`
				TextColumnFixedWidth string `xml:"TextColumnFixedWidth,attr"`
				TextColumnMaxWidth   string `xml:"TextColumnMaxWidth,attr"`
				Properties           struct {
					Text         string `xml:",chardata"`
					InsetSpacing struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"type,attr"`
						ListItem []struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"ListItem"`
					} `xml:"InsetSpacing"`
				} `xml:"Properties"`
			} `xml:"TextFramePreference"`
			TextWrapPreference struct {
				Text                  string `xml:",chardata"`
				Inverse               string `xml:"Inverse,attr"`
				ApplyToMasterPageOnly string `xml:"ApplyToMasterPageOnly,attr"`
				TextWrapSide          string `xml:"TextWrapSide,attr"`
				TextWrapMode          string `xml:"TextWrapMode,attr"`
				Properties            struct {
					Text           string `xml:",chardata"`
					TextWrapOffset struct {
						Text   string `xml:",chardata"`
						Top    string `xml:"Top,attr"`
						Left   string `xml:"Left,attr"`
						Bottom string `xml:"Bottom,attr"`
						Right  string `xml:"Right,attr"`
					} `xml:"TextWrapOffset"`
				} `xml:"Properties"`
			} `xml:"TextWrapPreference"`
		} `xml:"TextFrame"`
	} `xml:"Spread"`
}
