package georeplication

import (
	"github.com/pborman/uuid"
)

const (
	// GeorepStatusCreated represents Created State
	GeorepStatusCreated = "Created"

	// GeorepStatusStarted represents Started State
	GeorepStatusStarted = "Started"

	// GeorepStatusStopped represents Stopped State
	GeorepStatusStopped = "Stopped"

	// GeorepStatusPaused represents Paused State
	GeorepStatusPaused = "Paused"
)

// GeorepSlaveHost represents Slave host UUID and Hostname
type GeorepSlaveHost struct {
	NodeID   uuid.UUID `json:"nodeid"`
	Hostname string    `json:"host"`
}

// GeorepWorker represents Geo-replication Worker
type GeorepWorker struct {
	MasterNode                 string `json:"master_node"`
	MasterNodeID               string `json:"node_id"`
	MasterBrickPath            string `json:"master_brick_path"`
	MasterBrick                string `json:"master_brick"`
	Status                     string `json:"worker_status"`
	LastSyncedTime             string `json:"last_synced"`
	LastSyncedTimeUTC          string `json:"last_synced_utc"`
	LastEntrySyncedTime        string `json:"last_synced_entry"`
	SlaveNode                  string `json:"slave_node"`
	CheckpointTime             string `json:"checkpoint_time"`
	CheckpointTimeUTC          string `json:"checkpoint_time_utc"`
	CheckpointCompleted        string `json:"checkpoint_completed"`
	CheckpointCompletedTime    string `json:"checkpoint_completion_time"`
	CheckpointCompletedTimeUTC string `json:"checkpoint_completion_time_utc"`
	MetaOps                    string `json:"meta"`
	EntryOps                   string `json:"entry"`
	DataOps                    string `json:"data"`
	FailedOps                  string `json:"failures"`
	CrawlStatus                string `json:"crawl_status"`
}

// GeorepSession represents Geo-replication session
type GeorepSession struct {
	MasterID   uuid.UUID         `json:"master_volume_id"`
	SlaveID    uuid.UUID         `json:"slave_volume_id"`
	MasterVol  string            `json:"master_volume"`
	SlaveUser  string            `json:"slave_user"`
	SlaveHosts []GeorepSlaveHost `json:"slave_hosts"`
	SlaveVol   string            `json:"slave_volume"`
	Status     string            `json:"monitor_status"`
	Workers    []GeorepWorker    `json:"workers"`
	Options    map[string]string `json:"options"`
}

// GeorepOption represents Config details
type GeorepOption struct {
	Name         string `json:"name"`
	Value        string `json:"value"`
	DefaultValue string `json:"default_value"`
	Configurable bool   `json:"configurable"`
	Modified     bool   `json:"modified"`
}
