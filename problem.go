package zabbix

// Trigger represent Zabbix trigger object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/trigger/object
type ProblemObject struct {
	EventId         string   `json:"eventid"`
	Source          string   `json:"source"`
	Object          string   `json:"object"`
	ObjectId        string   `json:"objectid"`
	Clock           string   `json:"clock"`
	Ns              string   `json:"ns"`
	RecoveryEventId string   `json:"r_eventid"`
	RecoveryClock   string   `json:"r_clock"`
	RecoveryNs      string   `json:"r_ns,string"`
	Correlationid   string   `json:"correlationid,string"`
	UserId          string   `json:"userid"`
	Name            string   `json:"name"`
	Acknowledged    string   `json:"acknowledged"`
	Severity        string   `json:"severity"`
	Suppressed      string   `json:"suppressed"`
	OpData          string   `json:"opdata"`
	URLS            []string `json:"urls"`
}

// Triggers is an array of Trigger
type ProblemsObject []Trigger

// TriggersGet Wrapper for trigger.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/trigger/get
func (api *API) ProblemsGet(params Params) (res ProblemsObject, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("problem.get", params, &res)
	return
}

// TriggerGetByID Gets trigger by Id only if there is exactly 1 matching host.
/*func (api *API) TriggerGetByID(id string) (res *Trigger, err error) {
	triggers, err := api.TriggersGet(Params{"triggerids": id})
	if err != nil {
		return
	}

	if len(triggers) != 1 {
		e := ExpectedOneResult(len(triggers))
		err = &e
		return
	}
	res = &triggers[0]
	return
}

// TriggersCreate Wrapper for trigger.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/trigger/create
func (api *API) TriggersCreate(triggers Triggers) (err error) {
	response, err := api.CallWithError("trigger.create", triggers)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	triggerids := result["triggerids"].([]interface{})
	for i, id := range triggerids {
		triggers[i].TriggerID = id.(string)
	}
	return
}

// TriggersUpdate Wrapper for trigger.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/trigger/update
func (api *API) TriggersUpdate(triggers Triggers) (err error) {
	// Clear up unwanted paramters (UUID) that are used for update commands
	for idx := range triggers {
		triggers[idx].UUID = ""
	}
	_, err = api.CallWithError("trigger.update", triggers)
	return
}

// TriggersDelete Wrapper for trigger.delete
// Cleans ItemId in all triggers elements if call succeed.
// https://www.zabbix.com/documentation/3.2/manual/api/reference/trigger/delete
func (api *API) TriggersDelete(triggers Triggers) (err error) {
	ids := make([]string, len(triggers))
	for i, trigger := range triggers {
		ids[i] = trigger.TriggerID
	}

	err = api.TriggersDeleteByIds(ids)
	if err == nil {
		for i := range triggers {
			triggers[i].TriggerID = ""
		}
	}
	return
}

// TriggersDeleteByIds Wrapper for trigger.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/trigger/delete
func (api *API) TriggersDeleteByIds(ids []string) (err error) {
	deleteIds, err := api.TriggersDeleteIDs(ids)
	if err != nil {
		return
	}
	l := len(deleteIds)
	if len(ids) != l {
		err = &ExpectedMore{len(ids), l}
	}
	return
}

// TriggersDeleteIDs Wrapper for trigger.delete
// return the id of the deleted trigger
func (api *API) TriggersDeleteIDs(ids []string) (triggerids []interface{}, err error) {
	response, err := api.CallWithError("trigger.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	triggerids1, ok := result["triggerids"].([]interface{})
	if !ok {
		triggerids2 := result["triggerids"].(map[string]interface{})
		for _, id := range triggerids2 {
			triggerids = append(triggerids, id)
		}
	} else {
		triggerids = triggerids1
	}
	return

}
*/
