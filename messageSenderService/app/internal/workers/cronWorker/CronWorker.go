package cronWorker

import (
	"fmt"
	"time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR_TO_TICK int = 8
const MINUTE_TO_TICK int = 00
const SECOND_TO_TICK int = 00

type JobTicker struct {
	timer *time.Timer
}

func Init() *JobTicker {
	return &JobTicker{timer: time.NewTimer(INTERVAL_PERIOD)}
}

func (t *JobTicker) StartTicker() {
	jobTicker := &JobTicker{}
	jobTicker.updateTimer()
	for {
		<-jobTicker.timer.C
		// here some useful work
		fmt.Println(time.Now(), "- just ticked")
		jobTicker.updateTimer()
	}
}

func (t *JobTicker) updateTimer() {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick, "- next tick")
	diff := nextTick.Sub(time.Now())
	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}
