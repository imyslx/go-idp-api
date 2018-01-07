package controller

import (
	"encoding/json"

	"github.com/couchbase/gocb"
	"github.com/goadesign/goa"
	"github.com/imyslx/go-idp-api/app"

	zlog "github.com/rs/zerolog/log"
)

// HostsController implements the hosts resource.
type HostsController struct {
	*goa.Controller
}

// NewHostsController creates a hosts controller.
func NewHostsController(service *goa.Service) *HostsController {
	return &HostsController{Controller: service.NewController("HostsController")}
}

// List runs the list action.
func (c *HostsController) List(ctx *app.ListHostsContext) error {

	// Execute N1QL Query.
	baseQuery :=
		"SELECT Hostname, Status, `Role`, Type, OperatingSystem, Tag, Kernel, MonitoringStatus" +
			" FROM `idp_database`"
	rows := ExecuteQuery(ctx.Payload, baseQuery)

	// Create responses.
	resp := app.BiResult{}
	count := 0
	status := "success"

	var row interface{}
	var err error
	for rows.Next(&row) {
		bit := new(app.BasicInfoType)
		jsonByte, err := json.Marshal(row)
		if err != nil {
			zlog.Error().Err(err).Msg("Could not marshal to json result rows.")
			status = "warning: Some Errors in Response."
			continue
		}
		json.Unmarshal(jsonByte, bit)
		resp.BasicInfo = append(resp.BasicInfo, bit)
		count++
	}
	if err = rows.Close(); err != nil {
		zlog.Error().Err(err).Msg("Couldn't get all the rows.")
	}

	resp.Count = &count
	resp.ResponseStatus = &status

	return ctx.OK(&resp)
}

// Simplelist runs the simplelist action.
func (c *HostsController) Simplelist(ctx *app.SimplelistHostsContext) error {

	baseQuery := "SELECT Hostname FROM `idp_database`"
	rows := ExecuteQuery(ctx.Payload, baseQuery)

	// Create responses.
	resp := app.SlResult{}
	var row interface{}
	var err error
	for rows.Next(&row) {
		str := new(app.SimpleListType)
		jsonByte, err := json.Marshal(row)
		if err != nil {
			zlog.Error().Err(err).Msg("Could not marshal to json result rows.")
			continue
		}
		json.Unmarshal(jsonByte, str)
		resp.Hostname = append(resp.Hostname, str.Hostname[0])
	}
	if err = rows.Close(); err != nil {
		zlog.Error().Err(err).Msg("Couldn't get all the rows.")
	}

	return ctx.OK(&resp)
}

// ExecuteQuery : Execute the query.
func ExecuteQuery(params *app.HostsPayload, baseQuery string) gocb.QueryResults {

	// Get connect to couchbase bucket.
	bucket := GetCbBucket("")

	// Create N1QL Query.
	query := CreateQuery(params, baseQuery)
	zlog.Debug().Msg("N1QL query: " + query)

	// Execute
	rows, err := bucket.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), nil)
	if err != nil {
		zlog.Error().Err(err).Msg("Could not query in N1QL.")
	}

	return rows
}

// CreateQuery : Create query for n1ql.
func CreateQuery(params *app.HostsPayload, baseQuery string) string {

	flag := false
	where := " WHERE "
	if params.Hostname != nil {
		AddWhere(&flag, &where, " Hostname LIKE '%"+*params.Hostname+"%' ")
	}
	if params.Os != nil {
		AddWhere(&flag, &where, " OperatingSystem LIKE '%"+*params.Os+"%' ")
	}
	if params.Status != nil {
		status := ""
		if *params.Status {
			status = "true"
		} else {
			status = "false"
		}
		AddWhere(&flag, &where, " Status = "+status+" ")
	}

	if flag {
		return baseQuery + where
	} else {
		return baseQuery
	}
}

// AddWhere : Add strings for where.
func AddWhere(flag *bool, where *string, str string) {
	if *flag {
		*where += " AND "
	}
	*where += str
	*flag = true
}
