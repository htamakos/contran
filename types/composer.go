package types

type ComposerConfig struct {
	Version  string                     `yaml:"version"`
	Services map[string]ComposerService `yaml:"services"`
}

type ComposerService struct {
	Name           string                 `yaml:"-" json:"-"`
	Image          string                 `yaml:"image,omitempty" json:"image,omitempty"`
	Environment    map[string]interface{} `yaml:"environment,omitempty" json:"environment,omitempty"`
	Commands       []string               `yaml:"command,omitempty" json:"command,omitempty"`
	CpuShares      int64                  `yaml:"cpu_shares,omitempty"`
	Cpus           float32                `yaml:"cpus,omitempty" json:"cpus,omitempty"`
	MemLimit       string                 `yaml:"mem_limit,omitempty" json:"mem_limit,omitempty"`
	MemReservation string                 `yaml:"mem_reservation,omitempty" json:"mem_reservation,omitempty"`
	User           string                 `yaml:"user,omitempty" json:"user,omitempty"`
	Privileged     bool                   `yaml:"privileged,omitempty" json:"privileged,omitempty"`
	Links          []string               `yaml:"links,omitempty" json:"links,omitempty"`
	Ports          []string               `yaml:"ports,omitempty" json:"ports,omitempty"`
	Volumes        []string               `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	Ulimits        map[string]interface{} `yaml:"ulimits,omitempty" json:"ulimits,omitempty"`
}
