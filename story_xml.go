package main

import "encoding/xml"

type StoryXML struct {
	XMLName    xml.Name `xml:"Story"`
	Text       string   `xml:",chardata"`
	IdPkg      string   `xml:"idPkg,attr"`
	DOMVersion string   `xml:"DOMVersion,attr"`
	Story      struct {
		Text             string `xml:",chardata"`
		Self             string `xml:"Self,attr"`
		UserText         string `xml:"UserText,attr"`
		IsEndnoteStory   string `xml:"IsEndnoteStory,attr"`
		AppliedTOCStyle  string `xml:"AppliedTOCStyle,attr"`
		TrackChanges     string `xml:"TrackChanges,attr"`
		StoryTitle       string `xml:"StoryTitle,attr"`
		AppliedNamedGrid string `xml:"AppliedNamedGrid,attr"`
		StoryPreference  struct {
			Text                   string `xml:",chardata"`
			OpticalMarginAlignment string `xml:"OpticalMarginAlignment,attr"`
			OpticalMarginSize      string `xml:"OpticalMarginSize,attr"`
			FrameType              string `xml:"FrameType,attr"`
			StoryOrientation       string `xml:"StoryOrientation,attr"`
			StoryDirection         string `xml:"StoryDirection,attr"`
		} `xml:"StoryPreference"`
		InCopyExportOption struct {
			Text                  string `xml:",chardata"`
			IncludeGraphicProxies string `xml:"IncludeGraphicProxies,attr"`
			IncludeAllResources   string `xml:"IncludeAllResources,attr"`
		} `xml:"InCopyExportOption"`
		ParagraphStyleRange []struct {
			Text                    string `xml:",chardata"`
			AppliedParagraphStyle   string `xml:"AppliedParagraphStyle,attr"`
			Hyphenation             string `xml:"Hyphenation,attr"`
			HyphenationZone         string `xml:"HyphenationZone,attr"`
			RuleAboveLineWeight     string `xml:"RuleAboveLineWeight,attr"`
			RuleBelowLineWeight     string `xml:"RuleBelowLineWeight,attr"`
			SplitColumnInsideGutter string `xml:"SplitColumnInsideGutter,attr"`
			Properties              struct {
				Text    string `xml:",chardata"`
				TabList struct {
					Text     string `xml:",chardata"`
					Type     string `xml:"type,attr"`
					ListItem struct {
						Text      string `xml:",chardata"`
						Type      string `xml:"type,attr"`
						Alignment struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"Alignment"`
						AlignmentCharacter struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"AlignmentCharacter"`
						Leader struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"Leader"`
						Position struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"Position"`
					} `xml:"ListItem"`
				} `xml:"TabList"`
			} `xml:"Properties"`
			CharacterStyleRange []struct {
				Text                  string `xml:",chardata"`
				AppliedCharacterStyle string `xml:"AppliedCharacterStyle,attr"`
				FillColor             string `xml:"FillColor,attr"`
				FontStyle             string `xml:"FontStyle,attr"`
				PointSize             string `xml:"PointSize,attr"`
				StrokeWeight          string `xml:"StrokeWeight,attr"`
				MiterLimit            string `xml:"MiterLimit,attr"`
				RubyFontSize          string `xml:"RubyFontSize,attr"`
				KentenFontSize        string `xml:"KentenFontSize,attr"`
				Properties            struct {
					Text    string `xml:",chardata"`
					Leading struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"Leading"`
					AppliedFont struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"AppliedFont"`
				} `xml:"Properties"`
				Content string `xml:"Content"`
			} `xml:"CharacterStyleRange"`
		} `xml:"ParagraphStyleRange"`
	} `xml:"Story"`
}
