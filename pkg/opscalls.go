package pkg

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func SearchForObjectsByName(ObjectNameToSearch string, SrchMaxRslt int, ObjSrchType []string, gblparams GlobalParams) {

	// THIS NEEDS TO BE ULTRA DYNAMIC
	// this is where tons of filters can be added

	srchfilters := &[]SearchFilterSubType{
		{
			ObjectName: ObjectNameToSearch,
			FilterType: "object_name",
		},
		{
			ObjectTypes: ObjSrchType,
			FilterType:  "object_type",
		},
	}

	srchfiltermain := &SearchFilter{
		FilterDefinitions: *srchfilters,
		MaxResults:        SrchMaxRslt,
	}

	jsonFilter, err := json.Marshal(srchfiltermain)
	if err != nil {
		log.Fatalln(err)
	}

	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/search"

	respcode, respbyte := HTTPCall(uriBuild, "POST", jsonFilter, "application/json", "application/json")

	if respcode == 200 {
		fmt.Printf("Successfully searched : %v.\n", ObjectNameToSearch)
		var jsonResultData SearchResult
		json.Unmarshal(respbyte, &jsonResultData)
		SearchResultAsTable(jsonResultData)
		// log.Println(jsonResultData)
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		fmt.Printf("                 : %v\n", string(respbyte))
		os.Exit(1)
	}

}

func GetSystemStatus(gblparams GlobalParams) {

	_, respbyte := HTTPCall("http://"+gblparams.AEHostname+":"+gblparams.AEPort+"/ae/api/v1/0/system/health?details=true", "GET", nil, "application/json", "application/json")

	var jsonData SystemServiceStatus
	json.Unmarshal(respbyte, &jsonData)

	// fmt.Println(jsonData)
	SysStatusAsTable(jsonData)

}

func GetActiveExecutions(LimitJobsOnly bool, gblparams GlobalParams) {

	limitScope := ""
	if LimitJobsOnly {
		limitScope = "?type=JOBS&type=JOBP&type=SCRI"
	}

	// _, respbyte := HTTPCall("http://automicAEserver:8088/ae/api/v1/"+ClientID+"/executions?type=JOBS&type=JOBP&type=SCRI", "GET", nil)
	_, respbyte := HTTPCall("http://"+gblparams.AEHostname+":"+gblparams.AEPort+"/ae/api/v1/"+gblparams.AEClientID+"/executions"+limitScope, "GET", nil, "application/json", "application/json")

	var jsonData Activities
	json.Unmarshal(respbyte, &jsonData)

	ActivitiesAsTable(jsonData)

}

func GetActiveRuntimeMetadata(TaskRunID int, gblparams GlobalParams) ExecuteObjectMetadata {

	uriToCall := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID)

	respcode, respByte := HTTPCall(uriToCall, "GET", nil, "application/json", "application/json")

	if respcode != 200 {

		fmt.Printf("REST call failed : %v\n", respcode)
		fmt.Printf("                 : %v\n", string(respByte))
		os.Exit(1)
	}

	var responseData ExecuteObjectMetadata
	json.Unmarshal(respByte, &responseData)
	return responseData

}

func GetSystemAgents(gblparams GlobalParams) {

	_, respbyte := HTTPCall("http://"+gblparams.AEHostname+":"+gblparams.AEPort+"/ae/api/v1/"+gblparams.AEClientID+"/system/agents?active=true", "GET", nil, "application/json", "application/json")

	//http://automicAEserver:8088/ae/api/v1/120/system/agents?active=true

	var jsonData AgentsData
	json.Unmarshal(respbyte, &jsonData)

	fmt.Println(jsonData.ActiveObjects)

}

func GetTaskComments(TaskRunID int, gblparams GlobalParams) {

	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID) + "/comments"
	_, respbyte := HTTPCall(uriBuild, "GET", nil, "application/json", "application/json")

	var jsonData []Comments
	json.Unmarshal(respbyte, &jsonData)

	// fmt.Println(string(respbyte))

	CommentsAsTable(jsonData)

}

func SetTaskComment(TaskRunID int, AddComment string, gblparams GlobalParams) {

	jsonData := []byte("{\"comment\": \"" + AddComment + "\" }")
	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID) + "/comments"
	respcode, respByte := HTTPCall(uriBuild, "POST", jsonData, "application/json", "application/json")

	if respcode == 200 {
		fmt.Printf("Successfully ADDED COMMENT on RunID %v.\n", strconv.Itoa(TaskRunID))
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		fmt.Printf("                 : %v\n", string(respByte))
		os.Exit(1)
	}

}

func GetPreviousExecutions(TaskName, TaskType, StartTime, EndTime string, ShowDeactivated bool, gblparams GlobalParams) {

	// &time_frame_from=2022-12-28T11:55:00Z
	// &time_frame_to=2022-12-28T13:00:00Z

	deactState := "false"
	if ShowDeactivated {
		deactState = "true"
	}

	// uriBuild := "http://automicAEserver:8088/ae/api/v1/" + ClientID + "/executions?type=" + TaskType + "&name=" + TaskName + "&time_frame_from=" + StartTime + "&time_frame_to=" + EndTime + "&include_deactivated=" + deactState
	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions?name=" + TaskName + "&time_frame_from=" + StartTime + "&time_frame_to=" + EndTime + "&include_deactivated=" + deactState
	_, respbyte := HTTPCall(uriBuild, "GET", nil, "application/json", "application/json")

	var jsonData Activities
	json.Unmarshal(respbyte, &jsonData)

	ActivitiesAsTable(jsonData)

}

func GetChildExecutions(TaskRunID int, gblparams GlobalParams) {

	// &time_frame_from=2022-12-28T11:55:00Z
	// &time_frame_to=2022-12-28T13:00:00Z

	// uriBuild := "http://automicAEserver:8088/ae/api/v1/" + ClientID + "/executions?type=" + TaskType + "&name=" + TaskName + "&time_frame_from=" + StartTime + "&time_frame_to=" + EndTime + "&include_deactivated=" + deactState
	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID) + "/children"
	_, respbyte := HTTPCall(uriBuild, "GET", nil, "application/json", "application/json")

	var jsonData Activities
	json.Unmarshal(respbyte, &jsonData)

	ActivitiesAsTable(jsonData)

}

func GetExportObjectCode(ObjectName string, gblparams GlobalParams) ([]byte, error) {

	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/objects/" + ObjectName
	respcode, respbyte := HTTPCall(uriBuild, "GET", nil, "application/json", "application/json")

	if respcode == 200 {
		fmt.Printf("Successfully extracted object definition %v.\n", ObjectName)
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		fmt.Printf("                 : %v\n", string(respbyte))
		return nil, fmt.Errorf("REST call failed. See previous error.")
	}

	return respbyte, nil

}

func ImportObjectCode(ImportFilename string, OverwriteObject bool, gblparams GlobalParams) {

	overwriteobj := "false"
	if OverwriteObject {
		overwriteobj = "true"
	}

	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/objects?overwrite_existing_objects=" + overwriteobj

	importFileData, err := os.ReadFile(ImportFilename)
	if err != nil {
		log.Println(err)
	}

	respcode, respByte := HTTPCall(uriBuild, "POST", importFileData, "application/json", "application/json")

	switch respcode {
	case 200:
		fmt.Printf("[ %v ] : Successfully imported object.\n", respcode)
	case 400:
		fmt.Printf("[ %v ] : Import has failed.\n", respcode)
		fmt.Printf("%v\n", string(respByte))
		fmt.Printf(" ==> Try using the --overwrite flag")
		// os.Exit(1)
	default:
		fmt.Printf("REST call failed : %v\n", respcode)
		fmt.Printf("                 : %v\n", string(respByte))
	}

	// log.Println(respStatus)
	// log.Println(string(respbyte))

}

func ExecuteTaskByName(TaskName string, gblparams GlobalParams) {

	var jsonData = []byte(fmt.Sprintf("{ \"object_name\": \"%v\"}", TaskName))

	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions"
	respcode, respbyte := HTTPCall(uriBuild, "POST", jsonData, "application/json", "application/json")
	if respcode == 200 {
		// extract the runid for the task
		var rtm RunTimeMeta
		json.Unmarshal(respbyte, &rtm)

		fmt.Printf("EXECUTE task : [ %v ]\n", TaskName)
		fmt.Printf("RUN ID       : [ %v ]\n", rtm.RunID)
		fmt.Printf("Get REPORT   : [ %v %v ]\n", "util.exe report --reporttype REP -r ", rtm.RunID)
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		os.Exit(1)
	}
}

func KillActiveExecution(TaskRunID int, gblparams GlobalParams) {

	responseMeta := GetActiveRuntimeMetadata(TaskRunID, gblparams)

	var jsonData []byte

	switch responseMeta.Type {
	case "JSCH|JOBP|JOBG":
		jsonData = []byte(`{
			"action": "cancel",
			"cancel": {
				"recursive": "true"
			}
		}`)
	default:
		jsonData = []byte(`{
			"action": "cancel",
			"cancel": {}
		}`)
	}

	uriBuild := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID) + "/status"

	respcode, respByte := HTTPCall(uriBuild, "POST", jsonData, "application/json", "application/json")

	if respcode == 200 {
		fmt.Printf("Successfully issued CANCEL on RunID %v.\n", strconv.Itoa(TaskRunID))
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		fmt.Printf("                 : %v\n", string(respByte))
		os.Exit(1)
	}

}

func RestartExecution(TaskRunID int, gblparams GlobalParams) {

	var jsonData []byte

	jsonData = []byte(`{
		"action": "restart"
	}`)

	uriToCall := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID) + "/status"
	respcode, respByte := HTTPCall(uriToCall, "POST", jsonData, "application/json", "application/json")

	if respcode == 200 {
		fmt.Printf("Successfully issued RESTART on RunID %v.\n", strconv.Itoa(TaskRunID))
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		fmt.Printf("                 : %v\n", string(respByte))
		os.Exit(1)
	}

}

func RunScriptCode(JCLScriptText string, gblparams GlobalParams) {

	var jsonData = []byte(`{
		"script": "` + JCLScriptText + `",
		"queue": "CLIENT_QUEUE"
	  }`)

	uriToCall := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/scripts"

	respcode, respByte := HTTPCall(uriToCall, "POST", jsonData, "application/json", "application/json")

	if respcode == 200 {
		fmt.Println("Successfully issued SCRI.")
		fmt.Printf(" : %v\n", string(respByte))
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		os.Exit(1)
	}

}

func GetTaskReportsAvailable(TaskRunID int, gblparams GlobalParams) {

	uriToCall := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID) + "/reports"
	respcode, respdata := HTTPCall(uriToCall, "GET", nil, "application/json", "application/json")

	if respcode == 200 {
		fmt.Printf("Successfully issued REPORT LIST on RunID %v.\n", strconv.Itoa(TaskRunID))
	} else {
		fmt.Printf("REST call failed : %v\n", respcode)
		os.Exit(1)
	}

	var jsonData []ReportTypesAvailable
	json.Unmarshal(respdata, &jsonData)

	fmt.Println("\n\n======================================= REPORT TYPES =======================================")
	// fmt.Println("If IN_DB is false and no timestamp, the process is still active and report will not be fetched.")
	if len(jsonData) > 0 {
		for _, rpttype := range jsonData {
			if len(rpttype.EndTimestamp) == 0 && !rpttype.ReportInDB {
				fmt.Printf("Process still active, %v cannot be fetched yet!\n", rpttype.ReportType)
			} else {
				fmt.Printf("[%v][%v] - Type: %v\n", rpttype.EndTimestamp, rpttype.ReportInDB, rpttype.ReportType)
			}
		}
	} else {
		fmt.Println("No report types available.")
	}
	fmt.Println("======================================= REPORT TYPES =======================================\n\n")

}

func GetTaskReportOutput(TaskRunID int, ReportType string, gblparams GlobalParams) {

	uriToCall := "http://" + gblparams.AEHostname + ":" + gblparams.AEPort + "/ae/api/v1/" + gblparams.AEClientID + "/executions/" + strconv.Itoa(TaskRunID) + "/reports/" + ReportType

	respcode, respdata := HTTPCall(uriToCall, "GET", nil, "application/json", "application/json")

	if respcode == 200 {
		fmt.Printf("Successfully issued REPORT on RunID %v.\n", strconv.Itoa(TaskRunID))
	} else {
		GetTaskReportsAvailable(TaskRunID, gblparams)
		os.Exit(0)
	}

	var jsonData ReportData
	json.Unmarshal(respdata, &jsonData)

	fmt.Println("\n\n======================================= REPORT CONTENT =======================================")
	fmt.Println(jsonData.ReportDataContent[0].ReportContent)
	fmt.Println("======================================= REPORT CONTENT =======================================\n\n")

}

// ============================================================

func basicAuthConv(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func HTTPCall(URI string, Method string, POSTData []byte, AcceptType, ContentType string) (int, []byte) {

	fmt.Printf("Calling URL : %s\n", URI)

	if AcceptType == "" {
		AcceptType = "application/json"
	}
	if ContentType == "" {
		ContentType = "application/json"
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// request, err := http.NewRequest("http://automicAEserver:8088/ae/api/v1/ping")
	// request, err := http.NewRequest("GET", "http://automicAEserver:8088/ae/api/v1/0/system/health", nil)

	request, err := http.NewRequest(Method, URI, bytes.NewBuffer(POSTData))
	if err != nil {
		log.Println(err)
		return 500, nil
	}
	request.Header.Add("Authorization", "Basic "+basicAuthConv("myUsername", "myPassword"))
	request.Header.Set("Accept", AcceptType)
	request.Header.Set("Accept", AcceptType)

	request.Header.Set("Content-Type", ContentType)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return 500, nil
	}

	// defer resp.Body.Close()

	respbyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode, respbyte
}

func SysStatusAsTable(jsonData SystemServiceStatus) {
	var tmpdata []string
	var table *tablewriter.Table

	// used if you want the table data held as a
	//   string you can maniplulate
	// tableString := &strings.Builder{}
	//table = tablewriter.NewWriter(tableString)

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"service name", "status", "instances", "process name", "process connections", "process last accessed"})
	// table.SetAutoMergeCells(true)
	// table.SetAutoMergeCellsByColumnIndex([]int{0})

	tmpdata = []string{"PWP", jsonData.PWPService.ServiceStatus, strconv.Itoa(jsonData.PWPService.InstancesRunning), "", "", ""}
	table.Append(tmpdata)
	for _, svcinfo := range jsonData.PWPService.ServiceAvailability {
		tmpdata = []string{"", "", "", svcinfo.ServiceName, strconv.Itoa(svcinfo.NumberOfConnections), svcinfo.LastSignOfLife}
		table.Append(tmpdata)
	}

	tmpdata = []string{"WP", jsonData.WPService.ServiceStatus, strconv.Itoa(jsonData.WPService.InstancesRunning), "", "", ""}
	table.Append(tmpdata)
	for _, svcinfo := range jsonData.WPService.ServiceAvailability {
		tmpdata = []string{"", "", "", svcinfo.ServiceName, strconv.Itoa(svcinfo.NumberOfConnections), svcinfo.LastSignOfLife}
		table.Append(tmpdata)
	}

	tmpdata = []string{"CP", jsonData.CPService.ServiceStatus, strconv.Itoa(jsonData.CPService.InstancesRunning), "", "", ""}
	table.Append(tmpdata)
	for _, svcinfo := range jsonData.CPService.ServiceAvailability {
		tmpdata = []string{"", "", "", svcinfo.ServiceName, strconv.Itoa(svcinfo.NumberOfConnections), svcinfo.LastSignOfLife}
		table.Append(tmpdata)
	}

	tmpdata = []string{"JWP", jsonData.JWPService.ServiceStatus, strconv.Itoa(jsonData.JWPService.InstancesRunning), "", "", ""}
	table.Append(tmpdata)
	for _, svcinfo := range jsonData.JWPService.ServiceAvailability {
		tmpdata = []string{"", "", "", svcinfo.ServiceName, strconv.Itoa(svcinfo.NumberOfConnections), svcinfo.LastSignOfLife}
		table.Append(tmpdata)
	}

	tmpdata = []string{"JCP", jsonData.JCPService.ServiceStatus, strconv.Itoa(jsonData.JCPService.InstancesRunning), "", "", ""}
	table.Append(tmpdata)
	for _, svcinfo := range jsonData.JCPService.ServiceAvailability {
		tmpdata = []string{"", "", "", svcinfo.ServiceName, strconv.Itoa(svcinfo.NumberOfConnections), svcinfo.LastSignOfLife}
		table.Append(tmpdata)
	}

	tmpdata = []string{"REST", jsonData.RESTService.ServiceStatus, strconv.Itoa(jsonData.RESTService.InstancesRunning), "", "", ""}
	table.Append(tmpdata)
	for _, svcinfo := range jsonData.RESTService.ServiceAvailability {
		tmpdata = []string{"", "", "", svcinfo.ServiceName, strconv.Itoa(svcinfo.NumberOfConnections), svcinfo.LastSignOfLife}
		table.Append(tmpdata)
	}

	table.Render()
}

func ActivitiesAsTable(jsonData Activities) {
	var tmpdata []string
	var table *tablewriter.Table

	// used if you want the table data held as a
	//   string you can maniplulate
	// tableString := &strings.Builder{}
	//table = tablewriter.NewWriter(tableString)

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"runid", "task type", "task name", "status", "agent", "activation", "start time", "end time", "estimated end"})
	// table.SetAutoMergeCells(true)
	// table.SetAutoMergeCellsByColumnIndex([]int{0})

	for _, info := range jsonData.ActiveObjects {

		// convert from RFC33339/ISO8601 format string to time type
		EstEndTime, _ := time.Parse(time.RFC3339, info.StartTime)
		// add on the estimate secs runtime
		EstEndTime = EstEndTime.Add(time.Second * time.Duration(info.EstimatedRuntime))

		tmpdata = []string{strconv.Itoa(info.RunID), info.Type, info.Alias, fmt.Sprintf("(%v) %v", strconv.Itoa(info.Status), info.StatusText), info.Agent, info.ActivationTime, info.StartTime, info.EndTime, EstEndTime.Format(time.RFC3339)}
		table.Append(tmpdata)
	}

	table.Render()
}

func SearchResultAsTable(jsonData SearchResult) {
	var tmpdata []string
	var table *tablewriter.Table

	// used if you want the table data held as a
	//   string you can maniplulate
	// tableString := &strings.Builder{}
	//table = tablewriter.NewWriter(tableString)

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ohidnr", "name", "type", "subtype", "title", "folderpath"})
	// table.SetAutoMergeCells(true)
	// table.SetAutoMergeCellsByColumnIndex([]int{0})

	for _, info := range jsonData.SrchRsltData {

		// convert from RFC33339/ISO8601 format string to time type
		// EstEndTime, _ := time.Parse(time.RFC3339, info.StartTime)
		// // add on the estimate secs runtime
		// EstEndTime = EstEndTime.Add(time.Second * time.Duration(info.EstimatedRuntime))

		tmpdata = []string{info.ID, info.Name, info.Type, info.SubType, info.Title, info.FolderPath}
		table.Append(tmpdata)
	}

	table.Render()
}

func CommentsAsTable(jsonData []Comments) {
	var tmpdata []string
	var table *tablewriter.Table

	// used if you want the table data held as a
	//   string you can maniplulate
	// tableString := &strings.Builder{}
	//table = tablewriter.NewWriter(tableString)

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"timestamp", "user", "comment"})
	// table.SetAutoMergeCells(true)
	// table.SetAutoMergeCellsByColumnIndex([]int{0})

	for _, info := range jsonData {

		// convert from RFC33339/ISO8601 format string to time type
		Timestamp, _ := time.Parse(time.RFC3339, info.Timestamp)

		tmpdata = []string{Timestamp.Format(time.RFC3339), info.Username, info.Comment}
		table.Append(tmpdata)
	}

	table.Render()
}
