package domain

type MarathonApp struct {
    Id               string  `json:"id"`
    Cmd              string  `json:"cmd"`
    Instances        int     `json:"instances"`
    Cpus             float64 `json:"cpus"`
    Mem              int     `json:"mem"`
    Disk             int     `json:"disk"`
    ServicePorts     []int   `json:"service_port"`
    ContainerPorts   []int   `json:"container_port"`
    ContainerType    string  `json:"container_type"`
    ContainerNetwork string  `json:"container_network"`
    ContainerImage   string  `json:"container_image"`
    BuildId          string  `json:"build_id"`
    TriggeredBy      string  `json:"triggered_by"`
    Commit           string  `json:"commit"`
    JobName          string  `json:"job_name"`
    Project          string  `json:"project"`
    User             string  `json:"user"`
    Env              string  `json:"env"`
    DeployId         string  `json:"deploy_id"`
    Version          string  `json:"version"`
}

type MarathonGroup struct {
    Id      string          `json:"id"`
    Version string          `json:"version"`
    Groups  []MarathonGroup `json:"groups"`
    Apps    []MarathonApp   `json:"apps"`
}