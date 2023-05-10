package constanses

import (
	"github.com/Markuysa/astroMSA/astroService/app/pkg/model"
	date "github.com/Markuysa/astroMSA/authService/app/pkg/externalModels"
	"time"
)

// constances to work with astro stuff
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

	AriesStartDate = date.New(21, 3, 0)
	AriesEndDate   = date.New(19, 4, 0)

	TaurusStartDate = date.New(20, 4, 0)
	TaurusEndDate   = date.New(20, 5, 0)

	GeminiStartDate = date.New(21, 5, 0)
	GeminiEndDate   = date.New(21, 6, 0)

	CancerStartDate = date.New(22, 6, 0)
	CancerEndDate   = date.New(22, 7, 0)

	LeoStartDate = date.New(23, 7, 0)
	LeoEndDate   = date.New(22, 8, 0)

	VirgoStartDate = date.New(23, 8, 0)
	VirgoEndDate   = date.New(22, 9, 0)

	LibraStartDate = date.New(23, 9, 0)
	LibraEndDate   = date.New(23, 10, 0)

	ScorpioStartDate = date.New(24, 10, 0)
	ScorpioEndDate   = date.New(22, 11, 0)

	SagittariusStartDate = date.New(23, 11, 0)
	SagittariusEndDate   = date.New(21, 12, 0)

	CapricornStartDate = date.New(22, 12, 0)
	CapricornEndDate   = date.New(20, 1, 0)

	AquariusStartDate = date.New(21, 1, 0)
	AquariusEndDate   = date.New(18, 2, 0)

	PiscesStartDate = date.New(19, 2, 0)
	PiscesEndDate   = date.New(20, 3, 0)
)

// HardCoded predictions for test
var (
	AriesPrediction = &model.Prediction{
		DateRange:     "21.03-19.04",
		CurrentDate:   time.Now().String(),
		Description:   "The day will be favorable for communication. If you need help and support, you will surely get them. In the morning, not everything can turn out as we would like. But you shouldn't worry too much about it. Soon the situation will change for the better, and you will see new opportunities.",
		Compatibility: LIBRA,
		Mood:          "Funny",
		Color:         "Black",
		LuckyNumber:   "1",
		LuckyTime:     "Night",
	}
	TaurusPrediction = &model.Prediction{
		DateRange:     "20.04-20.05",
		CurrentDate:   time.Now().String(),
		Description:   "The morning will hardly pass quietly. There will be more things to do at this time than usual, and you will have to hurry to cope with everything on time. It is not necessary to cope with everything yourself. For help, you can turn to both old and new acquaintances, they will not refuse.",
		Compatibility: ARIES,
		Mood:          "Funny",
		Color:         "Red",
		LuckyNumber:   "2",
		LuckyTime:     "Day",
	}
	GeminiPrediction = &model.Prediction{
		DateRange:     "21.05-21.06",
		CurrentDate:   time.Now().String(),
		Description:   "There will be more worries than usual. It is possible that someone close to you will need your help, and you will not want to refuse requests. Some Twins will need to change their plans to go on trips or focus on things that previously seemed not too important and urgent.",
		Compatibility: LIBRA,
		Mood:          "Sad",
		Color:         "White",
		LuckyNumber:   "51",
		LuckyTime:     "Day",
	}
	CancerPrediction = &model.Prediction{
		DateRange:     "22.06-22.07",
		CurrentDate:   time.Now().String(),
		Description:   "The day will be favorable for communication, meetings and conversations with people who are important to you. There will be no reasons for serious disagreements with loved ones. Some Cancers will make peace with those with whom they were in a quarrel. It is possible that old acquaintances will remind you of themselves. They will surely offer something interesting.",
		Compatibility: GEMINI,
		Mood:          "Strange",
		Color:         "Blue",
		LuckyNumber:   "12",
		LuckyTime:     "Night",
	}
	LeoPrediction = &model.Prediction{
		DateRange:     "23.07-22.08",
		CurrentDate:   time.Now().String(),
		Description:   "Today, the Lions will reap the fruits of their past decisions. And most likely they will be pleasant. Rely on your gut – it will tell you in which direction you should move, help you dodge problems. The horoscope for today for the sign of Leo advises talking to people close in spirit – it will not only bring interesting news, but also cheer up.",
		Compatibility: ARIES,
		Mood:          "Funny",
		Color:         "Gold",
		LuckyNumber:   "1",
		LuckyTime:     "Night",
	}
	VirgoPrediction = &model.Prediction{
		DateRange:     "23.08-22.09",
		CurrentDate:   time.Now().String(),
		Description:   "Today, Virgos should rely on their gut – it will tell you who you can trust and who you shouldn't deal with. Pay attention to microemotions – they are now more eloquent than words. The horoscope for today for the sign of Virgo advises to make plans, but not to dedicate anyone to them. They may not come true. It is better to spend the evening with your family – you just need to communicate with them.",
		Compatibility: GEMINI,
		Mood:          "Funny",
		Color:         "Pink",
		LuckyNumber:   "31",
		LuckyTime:     "Morning",
	}
	LibraPrediction = &model.Prediction{
		DateRange:     "23.09-23.10",
		CurrentDate:   time.Now().String(),
		Description:   "Libra should not lose heart today – the day will be difficult, but still interesting. It's better to keep your plans a secret. You can tell about them only to the closest ones. Give at least an hour to sports – you will feel cheerful and ready for new adventures. The horoscope for today for the sign of Libra advises to plunge into the sea of romance with your head – with a partner you will become even closer.",
		Compatibility: LIBRA,
		Mood:          "Funny",
		Color:         "Red",
		LuckyNumber:   "1",
		LuckyTime:     "Night",
	}
	ScorpioPrediction = &model.Prediction{
		DateRange:     "24.10-22.11",
		CurrentDate:   time.Now().String(),
		Description:   "Today, Scorpios should rely not on logic, but on intuition – an inner flair will tell you who you can communicate with on this day, and who you should not trust. Not everything will be said out loud now, so look closely at the microemotions of the interlocutor, they will be much more eloquent. The horoscope for today for the sign of Scorpio advises not to torment ideas, act here and now.",
		Compatibility: CANCER,
		Mood:          "Funny",
		Color:         "Red",
		LuckyNumber:   "1",
		LuckyTime:     "Night",
	}
	SagittariusPrediction = &model.Prediction{
		DateRange:     "23.11-21.12",
		CurrentDate:   time.Now().String(),
		Description:   "Today, Sagittarians need to rest – you should not even take up household chores. It is better to spend the day in peace and comfort. The only thing you can do now is a hobby – you will understand how to monetize it. The horoscope for today for the sign of Cancer advises you to listen to your gut – it will indicate who you can trust, and with which people you should not even talk.",
		Compatibility: CAPRICORN,
		Mood:          "Sad",
		Color:         "Blue",
		LuckyNumber:   "41",
		LuckyTime:     "Pink",
	}
	CapricornPrediction = &model.Prediction{
		DateRange:     "22.12-20.01",
		CurrentDate:   time.Now().String(),
		Description:   "Today it will not be easy for Capricorns: something or someone will make them nervous and angry. Keep cool so as not to destroy what you have been building for many months – you can ruin both relationships and business. The horoscope for today for the Pisces sign advises not to write a script for this day – your plans are not destined to come true. It's better to be open to everything new, it will color your life in bright colors.",
		Compatibility: LEO,
		Mood:          "Strange",
		Color:         "Blue",
		LuckyNumber:   "74",
		LuckyTime:     "Orange",
	}
	AquariusPrediction = &model.Prediction{
		DateRange:     "22.12-20.01",
		CurrentDate:   time.Now().String(),
		Description:   "Today, Aquarians invest money somewhere is the most terrible idea. After all, you can lose hard earned. It is better to keep them in a secret place and sometimes count them. And financial energy is right there. News on this day can turn your life upside down, of course, in a good way. The horoscope for today for the sign of Aquarius advises not to write a script for this day – your plans are not destined to come true.",
		Compatibility: SAGITTARIUS,
		Mood:          "Strange",
		Color:         "Blue",
		LuckyNumber:   "123",
		LuckyTime:     "BLack",
	}
	PiscesPrediction = &model.Prediction{
		DateRange:     "22.12-20.01",
		CurrentDate:   time.Now().String(),
		Description:   "Today it will not be easy for Capricorns: something or someone will make them nervous and angry. Keep cool so as not to destroy what you have been building for many months – you can ruin both relationships and business. The horoscope for today for the Pisces sign advises not to write a script for this day – your plans are not destined to come true. It's better to be open to everything new, it will color your life in bright colors.",
		Compatibility: TAURUS,
		Mood:          "Boring",
		Color:         "Silver",
		LuckyNumber:   "95",
		LuckyTime:     "Now",
	}
)

func ReturnConst(sign string) (*model.Prediction, error) {
	switch sign {
	case ARIES:
		{
			return AriesPrediction, nil
		}
	case TAURUS:
		{
			return TaurusPrediction, nil
		}
	case GEMINI:
		{
			return GeminiPrediction, nil
		}
	case CANCER:
		{
			return CancerPrediction, nil
		}
	case LEO:
		{
			return LeoPrediction, nil
		}
	case VIRGO:
		{
			return VirgoPrediction, nil
		}
	case LIBRA:
		{
			return LibraPrediction, nil
		}
	case SCORPIO:
		{
			return ScorpioPrediction, nil
		}
	case SAGITTARIUS:
		{
			return SagittariusPrediction, nil
		}
	case CAPRICORN:
		{
			return CapricornPrediction, nil
		}
	case AQUARiUS:
		{
			return AquariusPrediction, nil
		}
	case PISCES:
		{
			return PiscesPrediction, nil
		}
	}
	return nil, nil
}
