package structs

type Policy struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	Rules []Rules `json:"rules"`
}

type Criteria struct {
	MinSeverity string `json:"min_severity"`
}
type BlockDownload struct {
	Unscanned bool `json:"unscanned"`
	Active    bool `json:"active"`
}
type Actions struct {
	Mails         []string      `json:"mails"`
	FailBuild     bool          `json:"fail_build"`
	BlockDownload BlockDownload `json:"block_download"`
}
type Rules struct {
	Name     string   `json:"name"`
	Priority int      `json:"priority"`
	Criteria Criteria `json:"criteria"`
	Actions  Actions  `json:"actions"`
}