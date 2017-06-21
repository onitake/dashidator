package dashidator

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
	XmlNsXsi string `xmlns:xsi,attr`
	XmlNs string `xmlns,attr`
	XmlNsXlink string `xmlns:xlink,attr`
	XsiSchemaLocation string `xsi:schemaLocation,attr`
	Profiles string `profiles,attr`
	Type string `dynamic,attr`
	MinimumUpdatePeriod string `minimumUpdatePeriod,attr`
	PublishTime string `publishTime,attr`
	AvailabilityStartTime string `availabilityStartTime,attr`
	TimeShiftBufferDepth string `timeShiftBufferDepth,attr`
	SuggestedPresentationDelay string `suggestedPresentationDelay,attr`
	MinBufferTime string `minBufferTime,attr`
	ProgramInformation struct {
		Title string `Title`
	} `ProgramInformation`
	Location string `Location`
	Period struct {
		Id string `id,attr`
		Start string `start,attr`
		AdaptationSet []struct {
			Id string `id,attr`
			Group string `group,attr`
			MimeType string `mimeType,attr`
			Width string `width,attr`
			Height string `height,attr`
			Par string `par,attr`
			FrameRate string `frameRate,attr`
			SegmentAlignment string `segmentAlignment,attr`
			StartWithSAP string `startWithSAP,attr`
			SubsegmentAlignment string `subsegmentAlignment,attr`
			SubsegmentStartsWithSAP string `subsegmentStartsWithSAP,attr`
			SegmentTemplate struct {
				timescale string `timescale,attr`
				media string `media,attr`
				initialization string `initialization,attr`
				SegmentTimeline struct {
					S []struct {
						T string `t,attr`
						D string `d,attr`
					} `S`
				} `SegmentTimeline`
			} `SegmentTemplate`
			Representation struct {
				Id string `id`
				Codecs string `codecs,attr`
				Sar string `sar,attr`
				Bandwidth string `bandwidth,attr`
			} `Representation`
		} `AdaptationSet`
	} `Period`
	UTCTiming struct {
		SchemeIdUri string `schemeIdUri,attr`
		Value string `value,attr`
	} `UTCTiming`
}

func ReadMpd(reader io.Reader) (*Mpd, error) {
	decoder := xml.NewDecoder(reader)
	mpd := &Mpd{}
	err := decoder.Decode(mpd)
	return mpd, err
}
