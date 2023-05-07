package pkg

import "encoding/json"

type GlobalParams struct {
	AEHostname string
	AEPort     string
	AEClientID string
}

type SystemServiceStatus struct {
	SystemStatus           string                `json:"status"`
	SystemActiveExecutiosn int                   `json:"count_active_executions"`
	PWPService             ServiceStatusMetadata `json:"pwp"`
	WPService              ServiceStatusMetadata `json:"wp"`
	JWPService             ServiceStatusMetadata `json:"jwp"`
	JCPService             ServiceStatusMetadata `json:"jcp"`
	RESTService            ServiceStatusMetadata `json:"rest"`
	CPService              ServiceStatusMetadata `json:"cp"`
}

type ServiceStatusMetadata struct {
	ServiceStatus       string                      `json:"status"`
	InstancesRunning    int                         `json:"instancesRunning"`
	ServiceAvailability []ServiceStatusAvailability `json:"available"`
}
type ServiceStatusAvailability struct {
	ServiceName         string `json:"name"`
	NumberOfConnections int    `json:"count_of_connections"`
	LastSignOfLife      string `json:"last_life_sign"`
}

type Activities struct {
	TotalEntries  int                     `json:"total"`
	ActiveObjects []ExecuteObjectMetadata `json:"data"`
	HasMore       bool                    `json:"hasmore"`
}

type ExecuteObjectMetadata struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	Queue            string `json:"queue"`
	RunID            int    `json:"run_id"`
	Status           int    `json:"status"`
	StatusText       string `json:"status_text"`
	ActivationTime   string `json:"activation_time"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	Agent            string `json:"agent"`
	AgentPlatform    string `json:"platform"`
	ParentID         int    `json:"parent"`
	RefRunID         int    `json:"reference_run_id"`
	LineNum          int    `json:"line_number"`
	UserID           string `json:"user"`
	EstimatedRuntime int    `json:"estimated_runtime"`
	Title            string `json:"title"`
	Alias            string `json:"alias"`
	Activator        int    `json:"activator"`
	ActivatorObjType string `json:"activator_object_type"`
}

type ReportData struct {
	TotalPages        int              `json:"total"`
	ReportDataContent []ReportDataPage `json:"data"`
}
type ReportDataPage struct {
	PageNumber    int    `json:"page"`
	ReportContent string `json:"content"`
}

type ReportTypesAvailable struct {
	EndTimestamp string `json:"end_timestamp"`
	ReportType   string `json:"type"`
	ReportInDB   bool   `json:"is_db"`
}

type Comments struct {
	Timestamp string `json:"timestamp"`
	Comment   string `json:"comment"`
	Username  string `json:"user"`
}

type SingleObjectDump struct {
	Total       int             `json:"total"`
	RawJSONData json.RawMessage `json:"data"`
	FolderPath  string          `json:"path"`
	ClientID    int             `json:"client"`
	HasMore     bool            `json:"hasmore"`
}

type AgentsData struct {
	TotalEntries  int             `json:"total"`
	ActiveObjects []AgentMetaData `json:"data"`
	HasMore       bool            `json:"hasmore"`
}
type AgentMetaData struct {
	AgentName          string `json:"name"`
	AgentPlatform      string `json:"platform"`
	AgentAuthenticated bool   `json:"authenticated"`
	AgentVersion       string `json:"version"`
	AgentHardware      string `json:"hardware"`
	AgentIPAddress     string `json:"ip_address"`
	AgentPort          int    `json:"port"`
	AgentSoftware      string `json:"software"`
}

type RunTimeMeta struct {
	RunID int `json:"run_id"`
}

type SearchResultData struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	SubType      string `json:"sub_type"`
	Platform     string `json:"platform"`
	ID           string `json:"id"`
	Title        string `json:"title"`
	ArchiveKey1  string `json:"archive_key_1"`
	ArchiveKey2  string `json:"archive_key_2"`
	Agent        string `json:"agent"`
	Login        string `json:"login"`
	FolderPath   string `json:"folder_path"`
	Client       string `json:"client"`
	IsInactive   string `json:"is_inactive"`
	CreatedBy    string `json:"created_by"`
	ModifiedBy   string `json:"modified_by"`
	CreationDate int64  `json:"creation_date"`
	ModifiedDate int64  `json:"modified_date"`
	LastUsedDate int64  `json:"last_used_date"`
}

type SearchResult struct {
	SrchRsltData []SearchResultData `json:"data"`
	TotalRecords int32              `json:"total"`
	HasMore      bool               `json:"hasmore"`
}

// 1.
// "object_types": [
// 	"JOBP",
// 	"JOBS"
//   ],
//   "filter_identifier": "object_type"

// 2.
//   "query": "*queue*",
//   "filter_identifier": "title"

// 3.
// "object_name": "JOBS.NEW.1",
// "filter_identifier": "object_name"

type SearchFilterSubType struct {
	ObjectTypes []string `json:"object_types,omitempty"`
	ObjectTitle string   `json:"query,omitempty"`
	ObjectName  string   `json:"object_name,omitempty"`
	FilterType  string   `json:"filter_identifier"`
}

type SearchFilter struct {
	FilterDefinitions []SearchFilterSubType `json:"filters"`
	MaxResults        int                   `json:"max_results"`
}
