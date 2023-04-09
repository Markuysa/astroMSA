package protobuf

import (
	pb "github.com/Markuysa/astroMSA/apigw/app/protobuf/gen"
	"github.com/Markuysa/astroMSA/authService/app/internal/model"
	"github.com/Markuysa/astroMSA/authService/app/pkg/externalModels"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

// DateToProtobuf function converts
// internal structure of Date to protobuf date
func DateToProtobuf(date2 externalModels.Date) *date.Date {

	return &date.Date{
		Day:   int32(date2.Day()),
		Month: int32(date2.Month()),
		Year:  int32(date2.Year()),
	}
}

// DateToTime converts internal structure of Date
// to default package struct of Time
func DateToTime(date2 externalModels.Date) *time.Time {
	datetime := time.Time{}.AddDate(int(date2.BYear), int(date2.BMonth), int(date2.BDay))
	return &datetime
}

// TimeToInternalDate is the opposite to DateToTime
// converts the default package Time to internal struct Date
func TimeToInternalDate(date time.Time) externalModels.Date {
	return externalModels.Date{
		BDay:   int64(date.Day()),
		BMonth: int64(date.Month()),
		BYear:  int64(date.Year()),
	}

}

// ConvertUserToPbResponse converts the internal model User
// to protobuf response struct
func ConvertUserToPbResponse(user *model.User) *pb.GetUserResponse {
	return &pb.GetUserResponse{User: &pb.User{
		Email:     user.Email,
		Sign:      user.Sign,
		Name:      user.Name,
		ID:        uint64(int64(user.ID)),
		BirthInfo: DateToProtobuf(user.BirthInfo),
	}}
}

// ConvertUserRequestToUserStruct converts the add user request
// with user data to internal model of user
func ConvertUserRequestToUserStruct(request *pb.AddUserRequest) *model.User {

	return &model.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
		BirthInfo: *externalModels.New(
			int64(request.BirthInfo.Day),
			int64(request.BirthInfo.Month),
			int64(request.BirthInfo.Year),
		),
		Notifications: request.Notifications,
	}
}
