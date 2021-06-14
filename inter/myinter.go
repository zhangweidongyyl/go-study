package inter

type UserInfoInterface interface {
	getusername(uid uint32) (string, error)

	setusername(uid uint32) (bool, error)
}
