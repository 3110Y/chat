package user

type userService struct {
	userRepository *userRepository
}

func NewUserService(userRepository *userRepository) *userService {
	return &userService{userRepository: userRepository}
}
