package externalModels

type Date struct {
	bMonth int64
	bDay   int64
	bYear  int64
}

func New(day int64, month int64, year int64) Date {

	return Date{
		bMonth: month,
		bDay:   day,
		bYear:  year,
	}
}
func (d *Date) Day() int64 {
	return d.bDay
}
func (d *Date) Month() int64 {
	return d.bMonth
}
func (d *Date) Year() int64 {
	return d.bYear
}
