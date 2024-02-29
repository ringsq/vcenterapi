package vcenter

type LoginResult struct {
	Value string `json:"value"`
}

type DatacenterResult []struct {
	ID   string `json:"datacenter"`
	Name string `json:"name"`
}

type ClusterListResult []struct {
	ID         string `json:"cluster"`
	DrsEnabled bool   `json:"drs_enabled"`
	HaEnabled  bool   `json:"ha_enabled"`
	Name       string `json:"name"`
}

type VMlistResult []struct {
	CPUCount      int    `json:"cpu_count"`
	MemorySizeMiB int    `json:"memory_size_MiB"`
	Name          string `json:"name"`
	PowerState    string `json:"power_state"`
	ID            string `json:"vm"`
}

type VMDetail struct {
	Boot struct {
		Delay          int    `json:"delay"`
		EnterSetupMode bool   `json:"enter_setup_mode"`
		Retry          bool   `json:"retry"`
		RetryDelay     int    `json:"retry_delay"`
		Type           string `json:"type"`
	} `json:"boot"`
	BootDevices []interface{} `json:"boot_devices"`
	CPU         struct {
		CoresPerSocket   int  `json:"cores_per_socket"`
		Count            int  `json:"count"`
		HotAddEnabled    bool `json:"hot_add_enabled"`
		HotRemoveEnabled bool `json:"hot_remove_enabled"`
	} `json:"cpu"`
	Cdroms   interface{}     `json:"cdroms"`
	Disks    map[string]Disk `json:"disks"`
	Floppies struct {
	} `json:"floppies"`
	GuestOS  *string `json:"guest_OS"`
	Hardware struct {
		UpgradePolicy string `json:"upgrade_policy"`
		UpgradeStatus string `json:"upgrade_status"`
		Version       string `json:"version"`
	} `json:"hardware"`
	Identity struct {
		BiosUUID     string `json:"bios_uuid"`
		InstanceUUID string `json:"instance_uuid"`
		Name         string `json:"name"`
	} `json:"identity"`
	InstantCloneFrozen bool `json:"instant_clone_frozen"`
	Memory             struct {
		HotAddEnabled          bool `json:"hot_add_enabled"`
		HotAddIncrementSizeMiB int  `json:"hot_add_increment_size_MiB"`
		HotAddLimitMiB         int  `json:"hot_add_limit_MiB"`
		SizeMiB                int  `json:"size_MiB"`
	} `json:"memory"`
	Name         string         `json:"name"`
	Nics         map[string]NIC `json:"nics"`
	NvmeAdapters struct {
	} `json:"nvme_adapters"`
	ParallelPorts struct {
	} `json:"parallel_ports"`
	PowerState   string `json:"power_state"`
	SataAdapters struct {
	} `json:"sata_adapters"`
	ScsiAdapters map[string]struct {
		Label         string `json:"label"`
		PciSlotNumber int    `json:"pci_slot_number"`
		Scsi          Scsi   `json:"scsi"`
		Sharing       string `json:"sharing"`
		Type          string `json:"type"`
	} `json:"scsi_adapters"`
	SerialPorts struct {
	} `json:"serial_ports"`
}
type Backing struct {
	Type     string `json:"type"`
	VmdkFile string `json:"vmdk_file"`
}
type Scsi struct {
	Bus  int `json:"bus"`
	Unit int `json:"unit"`
}
type Disk struct {
	Backing  Backing `json:"backing"`
	Capacity int     `json:"capacity"`
	Label    string  `json:"label"`
	Scsi     Scsi    `json:"scsi"`
	Type     string  `json:"type"`
}

type NIC struct {
	AllowGuestControl bool `json:"allow_guest_control"`
	Backing           struct {
		ConnectionCookie      int    `json:"connection_cookie"`
		DistributedPort       string `json:"distributed_port"`
		DistributedSwitchUUID string `json:"distributed_switch_uuid"`
		Network               string `json:"network"`
		Type                  string `json:"type"`
	} `json:"backing"`
	Label                   string `json:"label"`
	MacAddress              string `json:"mac_address"`
	MacType                 string `json:"mac_type"`
	StartConnected          bool   `json:"start_connected"`
	State                   string `json:"state"`
	Type                    string `json:"type"`
	UptCompatibilityEnabled bool   `json:"upt_compatibility_enabled"`
	WakeOnLanEnabled        bool   `json:"wake_on_lan_enabled"`
}
