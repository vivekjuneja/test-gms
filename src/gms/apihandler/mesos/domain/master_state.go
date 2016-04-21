package domain

type MasterState struct {
	ActivatedSlaves     int    `json:"activated_slaves"`
	BuildDate           string `json:"build_date"`
	BuildTime           int    `json:"build_time"`
	BuildUser           string `json:"build_user"`
	CompletedFrameworks []struct {
		Active         bool          `json:"active"`
		Capabilities   []interface{} `json:"capabilities"`
		Checkpoint     bool          `json:"checkpoint"`
		CompletedTasks []struct {
			ExecutorID  string `json:"executor_id"`
			FrameworkID string `json:"framework_id"`
			ID          string `json:"id"`
			Labels      []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"labels"`
			Name      string `json:"name"`
			Resources struct {
				Cpus  float64 `json:"cpus"`
				Disk  int     `json:"disk"`
				Mem   int     `json:"mem"`
				Ports string  `json:"ports"`
			} `json:"resources"`
			SlaveID  string `json:"slave_id"`
			State    string `json:"state"`
			Statuses []struct {
				ContainerStatus struct {
					NetworkInfos []struct {
						IPAddress string `json:"ip_address"`
					} `json:"network_infos"`
				} `json:"container_status"`
				State     string  `json:"state"`
				Timestamp float64 `json:"timestamp"`
			} `json:"statuses"`
		} `json:"completed_tasks"`
		Executors        []interface{} `json:"executors"`
		FailoverTimeout  int           `json:"failover_timeout"`
		Hostname         string        `json:"hostname"`
		ID               string        `json:"id"`
		Name             string        `json:"name"`
		OfferedResources struct {
			Cpus int `json:"cpus"`
			Disk int `json:"disk"`
			Mem  int `json:"mem"`
		} `json:"offered_resources"`
		Offers           []interface{} `json:"offers"`
		Pid              string        `json:"pid"`
		RegisteredTime   float64       `json:"registered_time"`
		ReregisteredTime float64       `json:"reregistered_time"`
		Resources        struct {
			Cpus int `json:"cpus"`
			Disk int `json:"disk"`
			Mem  int `json:"mem"`
		} `json:"resources"`
		Role             string        `json:"role"`
		Tasks            []interface{} `json:"tasks"`
		UnregisteredTime float64       `json:"unregistered_time"`
		UsedResources    struct {
			Cpus int `json:"cpus"`
			Disk int `json:"disk"`
			Mem  int `json:"mem"`
		} `json:"used_resources"`
		User     string `json:"user"`
		WebuiURL string `json:"webui_url"`
	} `json:"completed_frameworks"`
	DeactivatedSlaves int     `json:"deactivated_slaves"`
	ElectedTime       float64 `json:"elected_time"`
	Flags             struct {
		AdvertiseIP               string `json:"advertise_ip"`
		AdvertisePort             string `json:"advertise_port"`
		AllocationInterval        string `json:"allocation_interval"`
		Allocator                 string `json:"allocator"`
		Authenticate              string `json:"authenticate"`
		AuthenticateSlaves        string `json:"authenticate_slaves"`
		Authenticators            string `json:"authenticators"`
		Authorizers               string `json:"authorizers"`
		FrameworkSorter           string `json:"framework_sorter"`
		Help                      string `json:"help"`
		Hostname                  string `json:"hostname"`
		HostnameLookup            string `json:"hostname_lookup"`
		InitializeDriverLogging   string `json:"initialize_driver_logging"`
		LogAutoInitialize         string `json:"log_auto_initialize"`
		LogDir                    string `json:"log_dir"`
		Logbufsecs                string `json:"logbufsecs"`
		LoggingLevel              string `json:"logging_level"`
		MaxSlavePingTimeouts      string `json:"max_slave_ping_timeouts"`
		Port                      string `json:"port"`
		Quiet                     string `json:"quiet"`
		Quorum                    string `json:"quorum"`
		RecoverySlaveRemovalLimit string `json:"recovery_slave_removal_limit"`
		Registry                  string `json:"registry"`
		RegistryFetchTimeout      string `json:"registry_fetch_timeout"`
		RegistryStoreTimeout      string `json:"registry_store_timeout"`
		RegistryStrict            string `json:"registry_strict"`
		RootSubmissions           string `json:"root_submissions"`
		SlavePingTimeout          string `json:"slave_ping_timeout"`
		SlaveReregisterTimeout    string `json:"slave_reregister_timeout"`
		UserSorter                string `json:"user_sorter"`
		Version                   string `json:"version"`
		WebuiDir                  string `json:"webui_dir"`
		WorkDir                   string `json:"work_dir"`
		Zk                        string `json:"zk"`
		ZkSessionTimeout          string `json:"zk_session_timeout"`
	} `json:"flags"`
	Frameworks []struct {
		Active         bool          `json:"active"`
		Capabilities   []interface{} `json:"capabilities"`
		Checkpoint     bool          `json:"checkpoint"`
		CompletedTasks []struct {
			ExecutorID  string `json:"executor_id"`
			FrameworkID string `json:"framework_id"`
			ID          string `json:"id"`
			Labels      []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"labels"`
			Name      string `json:"name"`
			Resources struct {
				Cpus  float64 `json:"cpus"`
				Disk  int     `json:"disk"`
				Mem   int     `json:"mem"`
				Ports string  `json:"ports"`
			} `json:"resources"`
			SlaveID  string `json:"slave_id"`
			State    string `json:"state"`
			Statuses []struct {
				ContainerStatus struct {
					NetworkInfos []struct {
						IPAddress string `json:"ip_address"`
					} `json:"network_infos"`
				} `json:"container_status"`
				State     string  `json:"state"`
				Timestamp float64 `json:"timestamp"`
			} `json:"statuses"`
		} `json:"completed_tasks"`
		Executors        []interface{} `json:"executors"`
		FailoverTimeout  int           `json:"failover_timeout"`
		Hostname         string        `json:"hostname"`
		ID               string        `json:"id"`
		Name             string        `json:"name"`
		OfferedResources struct {
			Cpus int `json:"cpus"`
			Disk int `json:"disk"`
			Mem  int `json:"mem"`
		} `json:"offered_resources"`
		Offers           []interface{} `json:"offers"`
		Pid              string        `json:"pid"`
		RegisteredTime   float64       `json:"registered_time"`
		ReregisteredTime float64       `json:"reregistered_time"`
		Resources        struct {
			Cpus  float64 `json:"cpus"`
			Disk  int     `json:"disk"`
			Mem   int     `json:"mem"`
			Ports string  `json:"ports"`
		} `json:"resources"`
		Role  string `json:"role"`
		Tasks []struct {
			ExecutorID  string `json:"executor_id"`
			FrameworkID string `json:"framework_id"`
			ID          string `json:"id"`
			Labels      []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"labels"`
			Name      string `json:"name"`
			Resources struct {
				Cpus  float64 `json:"cpus"`
				Disk  int     `json:"disk"`
				Mem   int     `json:"mem"`
				Ports string  `json:"ports"`
			} `json:"resources"`
			SlaveID  string `json:"slave_id"`
			State    string `json:"state"`
			Statuses []struct {
				ContainerStatus struct {
					NetworkInfos []struct {
						IPAddress string `json:"ip_address"`
					} `json:"network_infos"`
				} `json:"container_status"`
				Labels []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"labels"`
				State     string  `json:"state"`
				Timestamp float64 `json:"timestamp"`
			} `json:"statuses"`
		} `json:"tasks"`
		UnregisteredTime int `json:"unregistered_time"`
		UsedResources    struct {
			Cpus  float64 `json:"cpus"`
			Disk  int     `json:"disk"`
			Mem   int     `json:"mem"`
			Ports string  `json:"ports"`
		} `json:"used_resources"`
		User     string `json:"user"`
		WebuiURL string `json:"webui_url"`
	} `json:"frameworks"`
	Hostname    string        `json:"hostname"`
	ID          string        `json:"id"`
	Leader      string        `json:"leader"`
	LogDir      string        `json:"log_dir"`
	OrphanTasks []interface{} `json:"orphan_tasks"`
	Pid         string        `json:"pid"`
	Slaves      []struct {
		Active           bool     `json:"active"`
		Attributes       struct{} `json:"attributes"`
		Hostname         string   `json:"hostname"`
		ID               string   `json:"id"`
		OfferedResources struct {
			Cpus int `json:"cpus"`
			Disk int `json:"disk"`
			Mem  int `json:"mem"`
		} `json:"offered_resources"`
		Pid               string   `json:"pid"`
		RegisteredTime    float64  `json:"registered_time"`
		ReregisteredTime  float64  `json:"reregistered_time"`
		ReservedResources struct{} `json:"reserved_resources"`
		Resources         struct {
			Cpus  int    `json:"cpus"`
			Disk  int    `json:"disk"`
			Mem   int    `json:"mem"`
			Ports string `json:"ports"`
		} `json:"resources"`
		UnreservedResources struct {
			Cpus  int    `json:"cpus"`
			Disk  int    `json:"disk"`
			Mem   int    `json:"mem"`
			Ports string `json:"ports"`
		} `json:"unreserved_resources"`
		UsedResources struct {
			Cpus  float64 `json:"cpus"`
			Disk  int     `json:"disk"`
			Mem   int     `json:"mem"`
			Ports string  `json:"ports"`
		} `json:"used_resources"`
	} `json:"slaves"`
	StartTime              float64       `json:"start_time"`
	UnregisteredFrameworks []interface{} `json:"unregistered_frameworks"`
	Version                string        `json:"version"`
}
