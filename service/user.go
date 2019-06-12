package service

import (
	"context"
	"fmt"

	"github.com/yukpiz/go-protobuf-example/pb"
)

type UserService struct{}

func (s *UserService) GetUser(c context.Context, msg *pb.UserMessage) (*pb.UserResponse, error) {
	// TODO: ここにビジネスロジックを...
	// serviceはapplicationに相当しそう？
	// serviceにも抽象を持って、handlerから抽象に依存するようにしたい

	// TODO: ここでrepositoryを挟めば良さそう
	// u, err := s.UserRepository.FindByID(c, msg.Id)
	// if err != nil {
	//     return err
	// }
	// return u, nil
	var res *pb.UserResponse
	switch msg.Id {
	case 1:
		res = &pb.UserResponse{
			Id:        msg.Id,
			LastName:  "yuk1",
			FirstName: "piz1",
			Email:     "yukpiz1@gmail.com",
		}
	case 2:
		res = &pb.UserResponse{
			Id:        msg.Id,
			LastName:  "yuk2",
			FirstName: "piz2",
			Email:     "yukpiz2@gmail.com",
		}

	case 3:
		res = &pb.UserResponse{
			Id:        msg.Id,
			LastName:  "yuk3",
			FirstName: "piz3",
			Email:     "yukpiz3@gmail.com",
		}
	default:
		return nil, fmt.Errorf("not found record: user_id=%d", msg.Id)
	}
	return res, nil
}
