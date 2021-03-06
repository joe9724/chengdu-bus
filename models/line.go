package models

type LineStation struct {
	LineStations []LineStationModel `json:"lineStations"`
}
type LineModel struct {
	LineID   int64  `json:"line_id"`
	LineName string `json:"line_name"`
	EndEarlytime string `json:"end_earlytime,omitempty"`
	EndLatetime string `json:"end_latetime,omitempty"`
	EndStop string `json:"end_stop,omitempty"`
	StartEarlytime string `json:"start_earlytime,omitempty"`
	StartLatetime string `json:"start_latetime,omitempty"`
	StartStop string `json:"start_stop,omitempty"`
}
type StationModel struct {
	LineID    int64       `json:"line_id"`
	StName    string      `json:"st_name"`
	StLat     float64     `json:"st_lat"`
	StLon     float64     `json:"st_lon"`
	PassLines []LineModel `json:"passlines"`
}
type LineStationModel struct {
	IsNewData   string    `json:"is_new_data"`
	RouteID     int64     `json:"routeId"`
	RouteName   string    `json:"routeName"`
	RouteType   string    `json:"routeType"`
	SegmentList []Segment `json:"segmentList"`
}
type Segment struct {
	SegmentID           int64     `json:"segmentId"`
	SegmentName         string    `json:"segmentName"`
	FirstTime           string    `json:"firstTime"`
	LastTime            string    `json:"lasttime"`
	RoutePrice          string    `json:"routePrice"`
	NormalTimeSpan      int64     `json:"normalTimeSpan"`
	PeakTimeSpan        int64     `json:"peakTimeSpan"`
	StationList         []Station `json:"stationList"`
	FirstLastShiftInfo  string    `json:"firstLastShiftInfo"`
	FirstLastShiftInfo2 string    `json:"firstLastShiftInfo2"`
	Memos               string    `json:"memos"`
	RunDirection        string    `json:"runDirection"`
	Baidumapid          string    `json:"baidumapid"`
	Amapid              string    `json:"amapid"`
	DrawType            string    `json:"drawType"`
}
type Station struct {
	StationID          string     `json:"stationId"`
	StationName        string     `json:"stationName"`
	StationNo          string     `json:"stationNo"`
	StationPosition    StationPos `json:"stationPostion"`
	Stationmemo        string     `json:"stationmemo"`
	Speed              string     `json:"speed"`
	DualSerial         int64      `json:"dualSerial"`
	StationSort        int64      `json:"stationSort"`
	StationSectionList string     `json:"stationSectionList"`
}
type StationPos struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}
type BtkLineStation struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Dir   string `json:"dir"`
	Info  string `json:"info"`
	SEG   []Seg  `json:"seg"`
	Price string `json:"price"`
}
type Seg struct {
	SegID    int64        `json:"sid"`
	SegName  string       `json:"sname"`
	Stations []BtkStation `json:"stations"`
	Info     string       `json:"info"`
	Price    string       `json:"price"`
}
type BtkStation struct {
	StID   string  `json:"st_id"`
	StName string  `json:"st_name"`
	StNo   string  `json:"stno"`
	Lat    float32 `json:"st_lat"`
	Lon    float32 `json:"st_lon"`
}
type Realtime struct {
	RouterID          int64    `json:"routerId"`
	RunBusNum         int64    `json:"runBusNum"`
	RStaRealTInfoList []RsInfo `json:"rstaRealTInfoList"`
	IsEnd             string   `json:"IsEnd"`
}
type RsInfo struct {
	StationID          string `json:"stationId"`
	RStanum            int64  `json:"rstanum"`
	ExpArriveBusStaNum int64  `json:"expArriveBusStaNum"`
	StopBusStaNum      int64  `json:"stopBusStaNum"`
	BusType            int64  `json:"busType"`
	MediaRouteName     string `json:"mediaRouteName"`
}

type SearchRet struct {
	Lines    []LineModel    `json:"lines"`
	Stations []StationModel `json:"stations"`
}
