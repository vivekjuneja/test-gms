package domain

type MarathonApp struct {
    Id               string  `json:"id"`
    Cmd              string  `json:"cmd"`
    Instances        int     `json:"instances"`
    Cpus             float64 `json:"cpus"`
    Mem              int     `json:"mem"`
    Disk             int     `json:"disk"`
    ServiceInfo      []map[string]string `json:"service_info"`
    ContainerImage   string  `json:"container_image"`
    BuildId          string  `json:"build_id"`
    HaProxyGroup      string  `json:"haproxy_group"`
    Commit           string  `json:"commit"`
    JobName          string  `json:"job_name"`
    Project          string  `json:"project"`
    User             string  `json:"user"`
    Env              string  `json:"env"`
    DeployId         string  `json:"deploy_id"`
    Timestamp        string  `json:"timestamp"`
}

type MarathonGroup struct {
    Id      string          `json:"id"`
    Version string          `json:"version"`
    Groups  []MarathonGroup `json:"groups"`
    Apps    []MarathonApp   `json:"apps"`
}