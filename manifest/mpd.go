package manifest

import (
	"io"
	"encoding/xml"
)

const (
	MpdXmlNsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	MpdXmlNs2011 = "urn:mpeg:dash:schema:mpd:2011"
	MpdXmlNsXlink = "http://www.w3.org/1999/xlink"
	// Prepend with the desired xmlns
	MpdXsiSchemaLocation = "http://standards.iso.org/ittf/PubliclyAvailableStandards/MPEG-DASH_schema_files/DASH-MPD.xsd"
)

type Mpd struct {
	XmlNsXsi string `xml:"xmlns:xsi,attr"`
	XmlNs string `xml:"xmlns,attr""`
	XmlNsXlink string `xml:"xmlns:xlink,attr"`
	XsiSchemaLocation string `xml:"xsi:schemaLocation,attr"`
	Profiles string `xml:"profiles,attr"`
	Type string `xml:"type,attr"`
	MinimumUpdatePeriod string `xml:"minimumUpdatePeriod,attr"`
	PublishTime string `xml:"publishTime,attr"`
	AvailabilityStartTime string `xml:"availabilityStartTime,attr"`
	TimeShiftBufferDepth string `xml:"timeShiftBufferDepth,attr"`
	SuggestedPresentationDelay string `xml:"suggestedPresentationDelay,attr"`
	MinBufferTime string `xml:"minBufferTime,attr"`
	ProgramInformation struct {
		Title string `xml:"Title"`
	} `xml:"ProgramInformation"`
	Location string `xml:"Location"`
	Period struct {
		Id int `xml:"id,attr"`
		Start string `xml:"start,attr"`
		AdaptationSet []struct {
			Id int `xml:"id,attr"`
			Group int `xml:"group,attr"`
			MimeType string `xml:"mimeType,attr"`
			Width int `xml:"width,attr"`
			Height int `xml:"height,attr"`
			Par string `xml:"par,attr"`
			FrameRate int `xml:"frameRate,attr"`
			SegmentAlignment bool `xml:"segmentAlignment,attr"`
			StartWithSAP int `xml:"startWithSAP,attr"`
			SubsegmentAlignment bool `xml:"subsegmentAlignment,attr"`
			SubsegmentStartsWithSAP int `xml:"subsegmentStartsWithSAP,attr"`
			SegmentTemplate struct {
				timescale string `xml:"timescale,attr"`
				media string `xml:"media,attr"`
				initialization string `xml:"initialization,attr"`
				SegmentTimeline struct {
					S []struct {
						T int64 `xml:"t,attr"`
						D int64 `xml:"d,attr"`
					} `xml:"S"`
				} `xml:"SegmentTimeline"`
			} `xml:"SegmentTemplate"`
			Representation struct {
				Id int `xml:"id"`
				Codecs string `xml:"codecs,attr"`
				Sar string `xml:"sar,attr"`
				Bandwidth int `xml:"bandwidth,attr"`
			} `xml:"Representation"`
		} `xml:"AdaptationSet"`
	} `xml:"Period"`
	UTCTiming struct {
		SchemeIdUri string `xml:"schemeIdUri,attr"`
		Value string `xml:"value,attr"`
	} `xml:"UTCTiming"`
}

func ReadMpd(reader io.Reader) (*Mpd, error) {
	decoder := xml.NewDecoder(reader)
	mpd := &Mpd{}
	err := decoder.Decode(mpd)
	return mpd, err
}
