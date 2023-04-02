package protobuf

import (
	"github.com/Markuysa/astroMSA/authService/app/internal/model"
	"github.com/Markuysa/astroMSA/authService/app/pkg/externalModels"
	"github.com/Markuysa/astroMSA/authService/app/protobuf/pb"
	"google.golang.org/genproto/googleapis/type/date"
	"log"
	"time"
)

func DateToProtobuf(date2 externalModels.Date) *date.Date {

	return &date.Date{
		Day:   int32(date2.Day()),
		Month: int32(date2.Month()),
		Year:  int32(date2.Year()),
	}
}
func DateToTime(date2 externalModels.Date) *time.Time {
	datetime := time.Time{}.AddDate(int(date2.BYear), int(date2.BMonth), int(date2.BDay))
	return &datetime
}
func TimeToInternalDate(date time.Time) externalModels.Date {
	return externalModels.Date{
		BDay:   int64(date.Day()),
		BMonth: int64(date.Month()),
		BYear:  int64(date.Year()),
	}

}
func ConvertUserToPbResponse(user *model.User) *pb.GetUserResponse {
	return &pb.GetUserResponse{User: &pb.User{
		Email:     user.Email,
		Sign:      user.Sign,
		Name:      user.Name,
		ID:        uint64(int64(user.ID)),
		BirthInfo: DateToProtobuf(user.BirthInfo),
	}}
}

func ConvertUserRequestToUserStruct(request *pb.AddUserRequest) *model.User {
	datee := externalModels.New(
		int64(request.BirthInfo.Day),
		int64(request.BirthInfo.Month),
		int64(request.BirthInfo.Year),
	)
	log.Println(datee)
	return &model.User{
		Email:         request.Email,
		Name:          request.Name,
		Password:      request.Password,
		BirthInfo:     *datee,
		Notifications: request.Notifications,
	}
}
