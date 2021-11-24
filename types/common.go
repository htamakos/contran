package types

import "strconv"

type CommonSpecs struct {
	VolumeSpecs    map[string]VolumeSpec
	ContainerSpecs map[string]ContainerSpec
}

type VolumeSpec struct {
	Name string
	Path string
}

type ContainerSpec struct {
	Name              string
	Image             string
	Commands          []string
	Environments      []ContainerEnvironment
	MemoryLimit       ContainerMemoryLimit
	MemoryReservation ContainerMemoryReservation
	Cpu               ContainerCpu
	User              string
	Privileged        bool
	Restart           string
	Runtime           string
	Ports             []ContainerPort
	Links             []string
	Ulimits           []ContainerUlimit
	Volumes           []ContainerVolume
}

type ContainerEnvironment struct {
	Name  string
	Value interface{}
}

type ContainerMemoryLimit struct {
	Unit  string
	Value int64
}

type ContainerMemoryReservation struct {
	Unit  string
	Value int64
}

func (m *ContainerMemoryLimit) String() string {
	return strconv.FormatInt(m.Value, 10) + m.Unit
}

func (m *ContainerMemoryReservation) String() string {
	return strconv.FormatInt(m.Value, 10) + m.Unit
}

type ContainerPort struct {
	HostIp        string
	HostPort      int
	ContainerPort int
}

type ContainerVolume struct {
	Source   string
	Target   string
	ReadOnly bool
}

type ContainerUlimit struct {
	Name      string
	HardLimit int
	SoftLimit int
}

type ContainerCpu struct {
	Count     int64
	Percent   float32
	Set       []int16
	Quota     int64
	RtPeriod  int64
	RtRuntime int64
	Shares    int64
}
