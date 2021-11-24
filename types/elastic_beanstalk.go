package types

type EbDockerrunAwsJson struct {
	AWSEBDockerrunVersion interface{}             `json:"AWSEBDockerrunVersion"`
	Volumes               []EbVolume              `json:"volumes,omitempty"`
	ContainerDefinitions  []EbContainerDefinition `json:"containerDefinitions,omitempty"`
}

type EbContainerDefinition struct {
	Name              string          `json:"name"`
	Image             string          `json:"image"`
	Environment       []EbEnvironment `json:"environment,omitempty"`
	Essential         bool            `json:"essential"`
	Cpu               int64           `json:"cpu"`
	Memory            int64           `json:"memory"`
	MemoryReservation int64           `json:"memoryReservation"`
	MountPoints       []EbMountPoint  `json:"mountPoints,omitempty"`
	User              string          `json:"user,omitempty"`
	Privileged        bool            `json:"privileged,omitempty"`
	PortMappings      []EbPortMapping `json:"portMappings,omitempty"`
	Ulimits           []EbUlimit      `json:"ulimits,omitempty"`
	Links             []string        `json:"links,omitempty"`
	Commands          []string        `json:"command,omitempty"`
}

type EbEnvironment struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type EbVolume struct {
	Name string `json:"name"`
	Host EbHost `json:"host"`
}

type EbHost struct {
	SourcePath string `json:"sourcePath"`
}

type EbMountPoint struct {
	SourceVolume  string `json:"sourceVolume"`
	ContainerPath string `json:"containerPath"`
	ReadOnly      bool   `json:"readOnly"`
}

type EbUlimit struct {
	Name      string `json:"name"`
	SoftLimit int    `json:"softLimit"`
	HardLimit int    `json:"hardLimit"`
}

type EbPortMapping struct {
	HostPort      int `json:"hostPort"`
	ContainerPort int `json:"containerPort"`
}
