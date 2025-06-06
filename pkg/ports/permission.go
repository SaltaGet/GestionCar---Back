package ports


type PermissionService interface {
	PermissionByRoleID(roleID string) (permissions *[]string, err error)
}

type PermissionRepository interface {
	PermissionByRoleID(roleID string) (permissions *[]string, err error)
}