package entity

type User struct {
	ID         string
	Name       string
	Surname    string
	Patronymic string

	Email        string
	PasswordHash string
	PhoneNumber  string

	Roles []Role
}

func (u *User) AddRole(role Role) error {
	u.Roles = append(u.Roles, role)
	return nil
}

func (u *User) HasPermission(permission string) bool {
	for _, role := range u.Roles {
		if role.CanAccess(permission) {
			return true
		}
	}
	return false
}
