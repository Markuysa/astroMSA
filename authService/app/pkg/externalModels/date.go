package externalModels

type Date struct {
	BMonth int64
	BDay   int64
	BYear  int64
}

func New(day int64, month int64, year int64) *Date {
	return &Date{
		BMonth: month,
		BDay:   day,
		BYear:  year,
	}
}
func (d *Date) Day() int64 {
	return d.BDay
}
func (d *Date) Month() int64 {
	return d.BMonth
}
func (d *Date) Year() int64 {
	return d.BYear
}
