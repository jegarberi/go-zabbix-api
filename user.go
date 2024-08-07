package zabbix

import (
	"encoding/json"
	"fmt"
)

// User represent Zabbix user group object
// https://www.zabbix.com/documentation/6.2/manual/api/reference/user/object
type User struct {
	UserID   string        `json:"userid,omitempty" zabbix:"id"`
	Username string        `json:"username"`
	Name     string        `json:"name,omitempty"`
	Surname  string        `json:"surname,omitempty"`
	RoleID   string        `json:"roleid"`
	Groups   []UserGroupID `json:"usrgrps"`
}

func (u *User) GetExtraParams() Params {
	return Params{"selectUsrgrps": []string{"usrgrpid"}}
}

type UserGroupID string

func (u *UserGroupID) UnmarshalJSON(bytes []byte) error {
	data := make(map[string]string)
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}
	id, ok := data["usrgrpid"]
	if !ok {
		return fmt.Errorf("couldn't find key usergrpid")
	}
	*u = UserGroupID(id)
	return nil
}

func (u *UserGroupID) MarshalJSON() (bytes []byte, err error) {
	data := map[string]string{"usrgrpid": string(*u)}
	return json.Marshal(data)
}

// Users is an array of User
type Users []User

// UsersGet Wrapper for user.get
// https://www.zabbix.com/documentation/4.0/manual/api/reference/user/get
func (api *API) UsersGet(params Params) (res Users, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("user.get", params, &res)
	return
}

func (u *User) GetID() string {
	return u.UserID
}

func (u *User) SetID(id string) {
	u.UserID = id
}

func (u *User) GetAPIModule() string {
	return "user"
}
