package main

import "encoding/xml"

type DesignMapXML struct {
	XMLName             xml.Name `xml:"Document"`
	Text                string   `xml:",chardata"`
	IdPkg               string   `xml:"idPkg,attr"`
	DOMVersion          string   `xml:"DOMVersion,attr"`
	Self                string   `xml:"Self,attr"`
	StoryList           string   `xml:"StoryList,attr"`
	Name                string   `xml:"Name,attr"`
	ZeroPoint           string   `xml:"ZeroPoint,attr"`
	ActiveLayer         string   `xml:"ActiveLayer,attr"`
	CMYKProfile         string   `xml:"CMYKProfile,attr"`
	RGBProfile          string   `xml:"RGBProfile,attr"`
	SolidColorIntent    string   `xml:"SolidColorIntent,attr"`
	AfterBlendingIntent string   `xml:"AfterBlendingIntent,attr"`
	DefaultImageIntent  string   `xml:"DefaultImageIntent,attr"`
	RGBPolicy           string   `xml:"RGBPolicy,attr"`
	CMYKPolicy          string   `xml:"CMYKPolicy,attr"`
	AccurateLABSpots    string   `xml:"AccurateLABSpots,attr"`
	Properties          struct {
		Text  string `xml:",chardata"`
		Label struct {
			Text         string `xml:",chardata"`
			KeyValuePair struct {
				Text  string `xml:",chardata"`
				Key   string `xml:"Key,attr"`
				Value string `xml:"Value,attr"`
			} `xml:"KeyValuePair"`
		} `xml:"Label"`
	} `xml:"Properties"`
	Language struct {
		Text                string `xml:",chardata"`
		Self                string `xml:"Self,attr"`
		Name                string `xml:"Name,attr"`
		SingleQuotes        string `xml:"SingleQuotes,attr"`
		DoubleQuotes        string `xml:"DoubleQuotes,attr"`
		PrimaryLanguageName string `xml:"PrimaryLanguageName,attr"`
		SublanguageName     string `xml:"SublanguageName,attr"`
		ID                  string `xml:"Id,attr"`
		HyphenationVendor   string `xml:"HyphenationVendor,attr"`
		SpellingVendor      string `xml:"SpellingVendor,attr"`
	} `xml:"Language"`
	Graphic struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"Graphic"`
	Fonts struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"Fonts"`
	Styles struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"Styles"`
	NumberingList struct {
		Text                           string `xml:",chardata"`
		Self                           string `xml:"Self,attr"`
		Name                           string `xml:"Name,attr"`
		ContinueNumbersAcrossStories   string `xml:"ContinueNumbersAcrossStories,attr"`
		ContinueNumbersAcrossDocuments string `xml:"ContinueNumbersAcrossDocuments,attr"`
	} `xml:"NumberingList"`
	NamedGrid struct {
		Text                string `xml:",chardata"`
		Self                string `xml:"Self,attr"`
		Name                string `xml:"Name,attr"`
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
	} `xml:"NamedGrid"`
	ConditionalTextPreference struct {
		Text                    string `xml:",chardata"`
		ShowConditionIndicators string `xml:"ShowConditionIndicators,attr"`
		ActiveConditionSet      string `xml:"ActiveConditionSet,attr"`
	} `xml:"ConditionalTextPreference"`
	Preferences struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"Preferences"`
	EndnoteOption struct {
		Text                 string `xml:",chardata"`
		EndnoteTitle         string `xml:"EndnoteTitle,attr"`
		EndnoteTitleStyle    string `xml:"EndnoteTitleStyle,attr"`
		StartEndnoteNumberAt string `xml:"StartEndnoteNumberAt,attr"`
		EndnoteMarkerStyle   string `xml:"EndnoteMarkerStyle,attr"`
		EndnoteTextStyle     string `xml:"EndnoteTextStyle,attr"`
		EndnoteSeparatorText string `xml:"EndnoteSeparatorText,attr"`
		EndnotePrefix        string `xml:"EndnotePrefix,attr"`
		EndnoteSuffix        string `xml:"EndnoteSuffix,attr"`
		Properties           struct {
			Text                  string `xml:",chardata"`
			EndnoteNumberingStyle struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"EndnoteNumberingStyle"`
			RestartEndnoteNumbering struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"RestartEndnoteNumbering"`
			EndnoteMarkerPositioning struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"EndnoteMarkerPositioning"`
			ScopeValue struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"ScopeValue"`
			FrameCreateOption struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"FrameCreateOption"`
			ShowEndnotePrefixSuffix struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"ShowEndnotePrefixSuffix"`
		} `xml:"Properties"`
	} `xml:"EndnoteOption"`
	TextFrameFootnoteOptionsObject struct {
		Text                  string `xml:",chardata"`
		EnableOverrides       string `xml:"EnableOverrides,attr"`
		SpanFootnotesAcross   string `xml:"SpanFootnotesAcross,attr"`
		MinimumSpacingOption  string `xml:"MinimumSpacingOption,attr"`
		SpaceBetweenFootnotes string `xml:"SpaceBetweenFootnotes,attr"`
	} `xml:"TextFrameFootnoteOptionsObject"`
	LinkedStoryOption struct {
		Text                      string `xml:",chardata"`
		UpdateWhileSaving         string `xml:"UpdateWhileSaving,attr"`
		WarnOnUpdateOfEditedStory string `xml:"WarnOnUpdateOfEditedStory,attr"`
		RemoveForcedLineBreaks    string `xml:"RemoveForcedLineBreaks,attr"`
		ApplyStyleMappings        string `xml:"ApplyStyleMappings,attr"`
	} `xml:"LinkedStoryOption"`
	LinkedPageItemOption struct {
		Text                         string `xml:",chardata"`
		UpdateLinkWhileSaving        string `xml:"UpdateLinkWhileSaving,attr"`
		WarnOnUpdateOfEditedPageItem string `xml:"WarnOnUpdateOfEditedPageItem,attr"`
		PreserveSizeAndShape         string `xml:"PreserveSizeAndShape,attr"`
		PreserveAppearance           string `xml:"PreserveAppearance,attr"`
		PreserveInteractivity        string `xml:"PreserveInteractivity,attr"`
		PreserveFrameContent         string `xml:"PreserveFrameContent,attr"`
		PreserveOthers               string `xml:"PreserveOthers,attr"`
	} `xml:"LinkedPageItemOption"`
	WatermarkPreference struct {
		Text                        string `xml:",chardata"`
		WatermarkVisibility         string `xml:"WatermarkVisibility,attr"`
		WatermarkDoPrint            string `xml:"WatermarkDoPrint,attr"`
		WatermarkDrawInBack         string `xml:"WatermarkDrawInBack,attr"`
		WatermarkText               string `xml:"WatermarkText,attr"`
		WatermarkFontFamily         string `xml:"WatermarkFontFamily,attr"`
		WatermarkFontStyle          string `xml:"WatermarkFontStyle,attr"`
		WatermarkFontPointSize      string `xml:"WatermarkFontPointSize,attr"`
		WatermarkOpacity            string `xml:"WatermarkOpacity,attr"`
		WatermarkRotation           string `xml:"WatermarkRotation,attr"`
		WatermarkHorizontalPosition string `xml:"WatermarkHorizontalPosition,attr"`
		WatermarkHorizontalOffset   string `xml:"WatermarkHorizontalOffset,attr"`
		WatermarkVerticalPosition   string `xml:"WatermarkVerticalPosition,attr"`
		WatermarkVerticalOffset     string `xml:"WatermarkVerticalOffset,attr"`
		Properties                  struct {
			Text               string `xml:",chardata"`
			WatermarkFontColor struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"WatermarkFontColor"`
		} `xml:"Properties"`
	} `xml:"WatermarkPreference"`
	TaggedPDFPreference struct {
		Text           string `xml:",chardata"`
		StructureOrder string `xml:"StructureOrder,attr"`
	} `xml:"TaggedPDFPreference"`
	AdjustLayoutPreference struct {
		Text                              string `xml:",chardata"`
		EnableAdjustLayout                string `xml:"EnableAdjustLayout,attr"`
		AllowLockedObjectsToAdjust        string `xml:"AllowLockedObjectsToAdjust,attr"`
		AllowFontSizeAndLeadingAdjustment string `xml:"AllowFontSizeAndLeadingAdjustment,attr"`
		ImposeFontSizeRestriction         string `xml:"ImposeFontSizeRestriction,attr"`
		MinimumFontSize                   string `xml:"MinimumFontSize,attr"`
		MaximumFontSize                   string `xml:"MaximumFontSize,attr"`
		EnableAutoAdjustMargins           string `xml:"EnableAutoAdjustMargins,attr"`
	} `xml:"AdjustLayoutPreference"`
	HTMLFXLExportPreference struct {
		Text                string `xml:",chardata"`
		EpubPageRange       string `xml:"EpubPageRange,attr"`
		EpubPageRangeFormat string `xml:"EpubPageRangeFormat,attr"`
	} `xml:"HTMLFXLExportPreference"`
	PublishExportPreference struct {
		Text                   string `xml:",chardata"`
		PublishCover           string `xml:"PublishCover,attr"`
		CoverImageFile         string `xml:"CoverImageFile,attr"`
		PublishPageRange       string `xml:"PublishPageRange,attr"`
		PublishPageRangeFormat string `xml:"PublishPageRangeFormat,attr"`
		ImageConversion        string `xml:"ImageConversion,attr"`
		ImageExportResolution  string `xml:"ImageExportResolution,attr"`
		PublishDescription     string `xml:"PublishDescription,attr"`
		PublishFileName        string `xml:"PublishFileName,attr"`
		PublishFormat          string `xml:"PublishFormat,attr"`
		CoverPage              string `xml:"CoverPage,attr"`
		GIFOptionsPalette      string `xml:"GIFOptionsPalette,attr"`
		JPEGOptionsQuality     string `xml:"JPEGOptionsQuality,attr"`
		PublishPdf             string `xml:"PublishPdf,attr"`
	} `xml:"PublishExportPreference"`
	TextVariable []struct {
		Text                            string `xml:",chardata"`
		Self                            string `xml:"Self,attr"`
		Name                            string `xml:"Name,attr"`
		VariableType                    string `xml:"VariableType,attr"`
		ChapterNumberVariablePreference struct {
			Text       string `xml:",chardata"`
			TextBefore string `xml:"TextBefore,attr"`
			Format     string `xml:"Format,attr"`
			TextAfter  string `xml:"TextAfter,attr"`
		} `xml:"ChapterNumberVariablePreference"`
		DateVariablePreference struct {
			Text       string `xml:",chardata"`
			TextBefore string `xml:"TextBefore,attr"`
			Format     string `xml:"Format,attr"`
			TextAfter  string `xml:"TextAfter,attr"`
		} `xml:"DateVariablePreference"`
		FileNameVariablePreference struct {
			Text             string `xml:",chardata"`
			TextBefore       string `xml:"TextBefore,attr"`
			IncludePath      string `xml:"IncludePath,attr"`
			IncludeExtension string `xml:"IncludeExtension,attr"`
			TextAfter        string `xml:"TextAfter,attr"`
		} `xml:"FileNameVariablePreference"`
		CaptionMetadataVariablePreference struct {
			Text                 string `xml:",chardata"`
			TextBefore           string `xml:"TextBefore,attr"`
			MetadataProviderName string `xml:"MetadataProviderName,attr"`
			TextAfter            string `xml:"TextAfter,attr"`
		} `xml:"CaptionMetadataVariablePreference"`
		PageNumberVariablePreference struct {
			Text       string `xml:",chardata"`
			TextBefore string `xml:"TextBefore,attr"`
			Format     string `xml:"Format,attr"`
			TextAfter  string `xml:"TextAfter,attr"`
			Scope      string `xml:"Scope,attr"`
		} `xml:"PageNumberVariablePreference"`
		MatchParagraphStylePreference struct {
			Text                  string `xml:",chardata"`
			TextBefore            string `xml:"TextBefore,attr"`
			TextAfter             string `xml:"TextAfter,attr"`
			AppliedParagraphStyle string `xml:"AppliedParagraphStyle,attr"`
			SearchStrategy        string `xml:"SearchStrategy,attr"`
			ChangeCase            string `xml:"ChangeCase,attr"`
			DeleteEndPunctuation  string `xml:"DeleteEndPunctuation,attr"`
		} `xml:"MatchParagraphStylePreference"`
	} `xml:"TextVariable"`
	Tags struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"Tags"`
	Layer struct {
		Text       string `xml:",chardata"`
		Self       string `xml:"Self,attr"`
		Name       string `xml:"Name,attr"`
		Visible    string `xml:"Visible,attr"`
		Locked     string `xml:"Locked,attr"`
		IgnoreWrap string `xml:"IgnoreWrap,attr"`
		ShowGuides string `xml:"ShowGuides,attr"`
		LockGuides string `xml:"LockGuides,attr"`
		UI         string `xml:"UI,attr"`
		Expendable string `xml:"Expendable,attr"`
		Printable  string `xml:"Printable,attr"`
		Properties struct {
			Text       string `xml:",chardata"`
			LayerColor struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"type,attr"`
				ListItem []struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ListItem"`
			} `xml:"LayerColor"`
		} `xml:"Properties"`
	} `xml:"Layer"`
	MasterSpread struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"MasterSpread"`
	Spread []struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"Spread"`
	Section []struct {
		Text                  string `xml:",chardata"`
		Self                  string `xml:"Self,attr"`
		Length                string `xml:"Length,attr"`
		Name                  string `xml:"Name,attr"`
		ContinueNumbering     string `xml:"ContinueNumbering,attr"`
		IncludeSectionPrefix  string `xml:"IncludeSectionPrefix,attr"`
		Marker                string `xml:"Marker,attr"`
		PageStart             string `xml:"PageStart,attr"`
		SectionPrefix         string `xml:"SectionPrefix,attr"`
		AlternateLayoutLength string `xml:"AlternateLayoutLength,attr"`
		AlternateLayout       string `xml:"AlternateLayout,attr"`
		Properties            struct {
			Text            string `xml:",chardata"`
			PageNumberStyle struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"PageNumberStyle"`
		} `xml:"Properties"`
	} `xml:"Section"`
	DocumentUser struct {
		Text       string `xml:",chardata"`
		Self       string `xml:"Self,attr"`
		UserName   string `xml:"UserName,attr"`
		Properties struct {
			Text      string `xml:",chardata"`
			UserColor struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"UserColor"`
		} `xml:"Properties"`
	} `xml:"DocumentUser"`
	CrossReferenceFormat []struct {
		Text                  string `xml:",chardata"`
		Self                  string `xml:"Self,attr"`
		Name                  string `xml:"Name,attr"`
		AppliedCharacterStyle string `xml:"AppliedCharacterStyle,attr"`
		BuildingBlock         []struct {
			Text                  string `xml:",chardata"`
			Self                  string `xml:"Self,attr"`
			BlockType             string `xml:"BlockType,attr"`
			AppliedCharacterStyle string `xml:"AppliedCharacterStyle,attr"`
			CustomText            string `xml:"CustomText,attr"`
			AppliedDelimiter      string `xml:"AppliedDelimiter,attr"`
			IncludeDelimiter      string `xml:"IncludeDelimiter,attr"`
		} `xml:"BuildingBlock"`
	} `xml:"CrossReferenceFormat"`
	BackingStory struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"BackingStory"`
	Story []struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"Story"`
	HyperlinkURLDestination []struct {
		Text                 string `xml:",chardata"`
		Self                 string `xml:"Self,attr"`
		Name                 string `xml:"Name,attr"`
		DestinationURL       string `xml:"DestinationURL,attr"`
		Hidden               string `xml:"Hidden,attr"`
		DestinationUniqueKey string `xml:"DestinationUniqueKey,attr"`
	} `xml:"HyperlinkURLDestination"`
	HyperlinkPageItemSource []struct {
		Text           string `xml:",chardata"`
		Self           string `xml:"Self,attr"`
		Name           string `xml:"Name,attr"`
		SourcePageItem string `xml:"SourcePageItem,attr"`
		Hidden         string `xml:"Hidden,attr"`
	} `xml:"HyperlinkPageItemSource"`
	Hyperlink []struct {
		Text                 string `xml:",chardata"`
		Self                 string `xml:"Self,attr"`
		Name                 string `xml:"Name,attr"`
		Source               string `xml:"Source,attr"`
		Visible              string `xml:"Visible,attr"`
		Highlight            string `xml:"Highlight,attr"`
		Width                string `xml:"Width,attr"`
		BorderStyle          string `xml:"BorderStyle,attr"`
		Hidden               string `xml:"Hidden,attr"`
		DestinationUniqueKey string `xml:"DestinationUniqueKey,attr"`
		Properties           struct {
			Text        string `xml:",chardata"`
			BorderColor struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"BorderColor"`
			Destination struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"type,attr"`
				ListItem []struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ListItem"`
			} `xml:"Destination"`
		} `xml:"Properties"`
	} `xml:"Hyperlink"`
	ColorGroup struct {
		Text             string `xml:",chardata"`
		Self             string `xml:"Self,attr"`
		Name             string `xml:"Name,attr"`
		IsRootColorGroup string `xml:"IsRootColorGroup,attr"`
		ColorGroupSwatch []struct {
			Text          string `xml:",chardata"`
			Self          string `xml:"Self,attr"`
			SwatchItemRef string `xml:"SwatchItemRef,attr"`
		} `xml:"ColorGroupSwatch"`
	} `xml:"ColorGroup"`
	IndexingSortOption []struct {
		Text       string `xml:",chardata"`
		Self       string `xml:"Self,attr"`
		Name       string `xml:"Name,attr"`
		Include    string `xml:"Include,attr"`
		Priority   string `xml:"Priority,attr"`
		HeaderType string `xml:"HeaderType,attr"`
	} `xml:"IndexingSortOption"`
	ABullet []struct {
		Text           string `xml:",chardata"`
		Self           string `xml:"Self,attr"`
		CharacterType  string `xml:"CharacterType,attr"`
		CharacterValue string `xml:"CharacterValue,attr"`
		Properties     struct {
			Text        string `xml:",chardata"`
			BulletsFont struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"BulletsFont"`
			BulletsFontStyle struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"BulletsFontStyle"`
		} `xml:"Properties"`
	} `xml:"ABullet"`
	Assignment struct {
		Text                    string `xml:",chardata"`
		Self                    string `xml:"Self,attr"`
		Name                    string `xml:"Name,attr"`
		UserName                string `xml:"UserName,attr"`
		ExportOptions           string `xml:"ExportOptions,attr"`
		IncludeLinksWhenPackage string `xml:"IncludeLinksWhenPackage,attr"`
		FilePath                string `xml:"FilePath,attr"`
		Properties              struct {
			Text       string `xml:",chardata"`
			FrameColor struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"FrameColor"`
		} `xml:"Properties"`
	} `xml:"Assignment"`
}
