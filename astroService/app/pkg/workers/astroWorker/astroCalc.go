package astroWorker

import (
	"github.com/Markuysa/astroMSA/astroService/app/pkg/constanses"
	"time"
)

func CalculateSign(birthTime time.Time) string {
	birthDay := birthTime.Day()
	birthMonth := birthTime.Month()
	switch {
	case birthDay >= constanses.AriesStartDate.Day() && birthMonth == constanses.AriesStartDate.Month() ||
		birthMonth == constanses.AriesEndDate.Month() && birthDay <= constanses.AriesEndDate.Day():
		return constanses.ARIES
	case birthDay >= constanses.TaurusStartDate.Day() && birthMonth == constanses.TaurusEndDate.Month() ||
		birthMonth == constanses.TaurusEndDate.Month() && birthDay <= constanses.TaurusEndDate.Day():
		return constanses.TAURUS
	case birthDay >= constanses.GeminiStartDate.Day() && birthMonth == constanses.GeminiStartDate.Month() ||
		birthMonth == constanses.GeminiEndDate.Month() && birthDay <= constanses.GeminiEndDate.Day():
		return constanses.GEMINI
	case birthDay >= constanses.CancerStartDate.Day() && birthMonth == constanses.CancerStartDate.Month() ||
		birthMonth == constanses.CancerEndDate.Month() && birthDay <= constanses.CancerEndDate.Day():
		return constanses.CANCER
	case birthDay >= constanses.LeoStartDate.Day() && birthMonth == constanses.LeoStartDate.Month() ||
		birthMonth == constanses.LeoEndDate.Month() && birthDay <= constanses.LeoEndDate.Day():
		return constanses.LEO
	case birthDay >= constanses.VirgoStartDate.Day() && birthMonth == constanses.VirgoStartDate.Month() ||
		birthMonth == constanses.VirgoEndDate.Month() && birthDay <= constanses.VirgoEndDate.Day():
		return constanses.VIRGO
	case birthDay >= constanses.LibraStartDate.Day() && birthMonth == constanses.LibraStartDate.Month() ||
		birthMonth == constanses.LibraEndDate.Month() && birthDay <= constanses.LibraEndDate.Day():
		return constanses.LIBRA
	case birthDay >= constanses.ScorpioStartDate.Day() && birthMonth == constanses.ScorpioStartDate.Month() ||
		birthMonth == constanses.ScorpioEndDate.Month() && birthDay <= constanses.ScorpioEndDate.Day():
		return constanses.SCORPIO
	case birthDay >= constanses.SagittariusStartDate.Day() && birthMonth == constanses.SagittariusStartDate.Month() ||
		birthMonth == constanses.SagittariusEndDate.Month() && birthDay <= constanses.SagittariusEndDate.Day():
		return constanses.SAGITTARIUS
	case birthDay >= constanses.CapricornStartDate.Day() && birthMonth == constanses.CapricornStartDate.Month() ||
		birthMonth == constanses.CapricornEndDate.Month() && birthDay <= constanses.CapricornEndDate.Day():
		return constanses.CAPRICORN
	case birthDay >= constanses.AquariusStartDate.Day() && birthMonth == constanses.AquariusStartDate.Month() ||
		birthMonth == constanses.AquariusEndDate.Month() && birthDay <= constanses.AquariusEndDate.Day():
		return constanses.AQUARiUS
	case birthDay >= constanses.PiscesStartDate.Day() && birthMonth == constanses.PiscesStartDate.Month() ||
		birthMonth == constanses.PiscesEndDate.Month() && birthDay <= constanses.PiscesEndDate.Day():
		return constanses.PISCES
	}
	return "error"
}
