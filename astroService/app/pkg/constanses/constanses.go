package constanses

import "time"

var (
	ARIES       = "aries"
	TAURUS      = "taurus"
	GEMINI      = "gemini"
	CANCER      = "cancer"
	LEO         = "leo"
	VIRGO       = "virgo"
	LIBRA       = "libra"
	SCORPIO     = "scorpio"
	SAGITTARIUS = "sagittarius"
	CAPRICORN   = "capricorn"
	AQUARiUS    = "aquarius"
	PISCES      = "pisces"

	AriesStartDate = time.Time{}.AddDate(0, 2, 20)
	AriesEndDate   = time.Time{}.AddDate(0, 3, 19)

	TaurusStartDate = time.Time{}.AddDate(0, 3, 20)
	TaurusEndDate   = time.Time{}.AddDate(0, 4, 20)

	GeminiStartDate = time.Time{}.AddDate(0, 4, 20)
	GeminiEndDate   = time.Time{}.AddDate(0, 5, 20)

	CancerStartDate = time.Time{}.AddDate(0, 5, 21)
	CancerEndDate   = time.Time{}.AddDate(0, 6, 21)

	LeoStartDate = time.Time{}.AddDate(0, 6, 22)
	LeoEndDate   = time.Time{}.AddDate(0, 7, 22)

	VirgoStartDate = time.Time{}.AddDate(0, 7, 23)
	VirgoEndDate   = time.Time{}.AddDate(0, 8, 21)

	LibraStartDate = time.Time{}.AddDate(0, 8, 22)
	LibraEndDate   = time.Time{}.AddDate(0, 9, 22)

	ScorpioStartDate = time.Time{}.AddDate(0, 9, 23)
	ScorpioEndDate   = time.Time{}.AddDate(0, 10, 21)

	SagittariusStartDate = time.Time{}.AddDate(0, 10, 22)
	SagittariusEndDate   = time.Time{}.AddDate(0, 11, 20)

	CapricornStartDate = time.Time{}.AddDate(0, 11, 21)
	CapricornEndDate   = time.Time{}.AddDate(0, 12, 19)

	AquariusStartDate = time.Time{}.AddDate(0, 12, 20)
	AquariusEndDate   = time.Time{}.AddDate(0, 1, 17)

	PiscesStartDate = time.Time{}.AddDate(0, 1, 18)
	PiscesEndDate   = time.Time{}.AddDate(0, 2, 19)
)
