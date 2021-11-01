package model

type Permission struct {
	Id             int    `json:"id"`
	PermissionName string `json:"permission_name"`
	Url            string `json:"url"`
	Method         string `json:"method"`
}

func (p Permission) TableName() string {
	return "permission"
}
