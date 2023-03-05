package model

type Prediction struct {
	Color         string `json:"color"`
	Compatibility string `json:"compatibility"`
	Description   string `json:"description"`
	LuckyNumber   string `json:"luckyNumber"`
	LuckyTime     string `json:"luckyTime"`
	Mood          string `json:"mood"`
}
