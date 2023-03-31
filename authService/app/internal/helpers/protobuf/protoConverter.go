package protobuf

import (
	"authService/app/protobuf/pb"
	"time"
)

func DateToProtobuf(time time.Time) *pb.DateTime {

	return &pb.DateTime{
		Day:   int32(time.Day()),
		Month: int32(time.Month()),
		Year:  int32(time.Year()),
	}
}
