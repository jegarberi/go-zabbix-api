package zabbix

type (
	// AvailableType (readonly) Availability of Zabbix agent
	// see "available" in: https://www.zabbix.com/documentation/3.2/manual/api/reference/host/object
	AvailableType int

	// StatusType Status and function of the host.
	// see "status" in:	https://www.zabbix.com/documentation/3.2/manual/api/reference/host/object
	StatusType int
)

const (
	// Unknown (default)
	Unknown AvailableType = 0
	// Available host is available
	Available AvailableType = 1
	// Unavailable host is unavailable
	Unavailable AvailableType = 2
)

const (
	// Monitored monitored host(default)
	Monitored StatusType = 0
	// Unmonitored unmonitored host
	Unmonitored StatusType = 1
)

// Host represent Zabbix host object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/object
type Host struct {
	HostID        string         `json:"hostid,omitempty" zabbix:"id"`
	Host          string         `json:"host"`
	Name          string         `json:"name"`
	Status        StatusType     `json:"status,string"`
	Description   string         `json:"description"`
	InventoryMode int            `json:"inventory_mode,string"`
	IPMIAuthType  int            `json:"ipmi_authtype,string"`
	IPMIPassword  string         `json:"ipmi_password"`
	IPMIUsername  string         `json:"ipmi_username"`
	IPMIPrivilege int            `json:"ipmi_privilege,string"`
	ProxyHostID   string         `json:"proxy_hostid,omitempty"`
	Tags          []HostTag      `json:"tags"`
	GroupIds      HostGroupIDs   `json:"groups"`
	Interfaces    HostInterfaces `json:"interfaces"`
	TemplateIDs   TemplateIDs    `json:"templates"`
	Macros        []Macro        `json:"macros"`
}

// Hosts is an array of Host
type Hosts []Host

type HostTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

// HostsGet Wrapper for host.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/get
func (api *API) HostsGet(params Params) (res Hosts, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("host.get", params, &res)
	return
}

// HostsGetByHostGroupIds Gets hosts by host group Ids.
func (api *API) HostsGetByHostGroupIds(ids []string) (res Hosts, err error) {
	return api.HostsGet(Params{"groupids": ids, "selectMacros": "extend", "selectInterfaces": "extend"})
}

// HostsGetByHostGroups Gets hosts by host groups.
func (api *API) HostsGetByHostGroups(hostGroups HostGroups) (res Hosts, err error) {
	ids := make([]string, len(hostGroups))
	for i, id := range hostGroups {
		ids[i] = id.GroupID
	}
	return api.HostsGetByHostGroupIds(ids)
}

// HostGetByID Gets host by Id only if there is exactly 1 matching host.
func (api *API) HostGetByID(id string) (res *Host, err error) {
	params := Params{
		"hostids":          id,
		"selectMacros":     "extend",
		"selectInterfaces": "extend",
	}
	hosts, err := api.HostsGet(params)
	if err != nil {
		return
	}

	if len(hosts) == 1 {
		res = &hosts[0]
	} else {
		e := ExpectedOneResult(len(hosts))
		err = &e
	}
	return
}

// HostGetByHost Gets host by Host only if there is exactly 1 matching host.
func (api *API) HostGetByHost(host string) (res *Host, err error) {
	hosts, err := api.HostsGet(Params{"filter": map[string]string{"host": host}, "selectMacros": "extend", "selectInterfaces": "extend"})
	if err != nil {
		return
	}

	if len(hosts) == 1 {
		res = &hosts[0]
	} else {
		e := ExpectedOneResult(len(hosts))
		err = &e
	}
	return
}

// HostsCreate Wrapper for host.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/create
func (api *API) HostsCreate(hosts Hosts) (err error) {
	response, err := api.CallWithError("host.create", hosts)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	hostids := result["hostids"].([]interface{})
	for i, id := range hostids {
		hosts[i].HostID = id.(string)
	}
	return
}

// HostsUpdate Wrapper for host.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/update
func (api *API) HostsUpdate(hosts Hosts) (err error) {
	_, err = api.CallWithError("host.update", hosts)
	return
}

// HostsDelete Wrapper for host.delete
// Cleans HostId in all hosts elements if call succeed.
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/delete
func (api *API) HostsDelete(hosts Hosts) (err error) {
	ids := make([]string, len(hosts))
	for i, host := range hosts {
		ids[i] = host.HostID
	}

	err = api.HostsDeleteByIds(ids)
	if err == nil {
		for i := range hosts {
			hosts[i].HostID = ""
		}
	}
	return
}

// HostsDeleteByIds Wrapper for host.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/delete
func (api *API) HostsDeleteByIds(ids []string) (err error) {
	hostIds := ids
	//for i, id := range ids {
	//	hostIds[i] = map[string]string{"hostid": id}
	//}

	response, err := api.CallWithError("host.delete", hostIds)
	if err != nil {
		// Zabbix 2.4 uses new syntax only
		if e, ok := err.(*Error); ok && e.Code == -32500 {
			response, err = api.CallWithError("host.delete", ids)
		}
	}
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	hostids := result["hostids"].([]interface{})
	if len(ids) != len(hostids) {
		err = &ExpectedMore{len(ids), len(hostids)}
	}
	return
}

func (host *Host) GetID() string {
	return host.HostID
}

func (host *Host) SetID(id string) {
	host.HostID = id
}

func (host *Host) GetAPIModule() string {
	return "host"
}

func (host *Host) GetExtraParams() Params {
	return Params{
		"selectMacros":     "extend",
		"selectInterfaces": "extend",
		"selectTags":       "extend",
		"selectGroups":     []string{"groupid"},
		"templated_hosts":  nil,
		//"selectParentTemplates": "extend",
	}
}
