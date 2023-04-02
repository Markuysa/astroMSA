package protobuf

import (
	"github.com/Markuysa/astroMSA/authService/app/internal/model"
	"github.com/Markuysa/astroMSA/authService/app/pkg/externalModels"
	"github.com/Markuysa/astroMSA/authService/app/protobuf/pb"
	"google.golang.org/genproto/googleapis/type/date"
)

func DateToProtobuf(date2 externalModels.Date) *date.Date {

	return &date.Date{
		Day:   int32(date2.Day()),
		Month: int32(date2.Month()),
		Year:  int32(date2.Year()),
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

	return &model.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
		BirthInfo: externalModels.New(int64(request.BirthInfo.Day),
			int64(request.BirthInfo.Month),
			int64(request.BirthInfo.Year),
		),
		Notifications: request.Notifications,
	}
}
