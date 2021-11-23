package translater

import (
	"encoding/json"
	"errors"

	"github.com/htamakos/contran/types"
)

type Eb struct{}

func (eb *Eb) Input(values []byte) (*types.CommonSpecs, error) {
	var ebJson types.EbDockerrunAwsJson

	err := json.Unmarshal(values, &ebJson)
	if err != nil {
		return nil, err
	}

	volumeSpecs := map[string]types.VolumeSpec{}

	for _, v := range ebJson.Volumes {
		volumeSpecs[v.Name] = types.VolumeSpec{
			Name: v.Name,
			Path: v.Host.SourcePath,
		}
	}

	containerSpecs := map[string]types.ContainerSpec{}
	for _, c := range ebJson.ContainerDefinitions {
		containerSpec := types.ContainerSpec{
			Name:  c.Name,
			Image: c.Image,
		}

		if c.Memory != 0 {
			containerSpec.MemoryLimit = types.ContainerMemoryLimit{
				Unit:  "m",
				Value: c.Memory,
			}
		}

		if c.MemoryReservation != 0 {
			containerSpec.MemoryReservation = types.ContainerMemoryReservation{
				Unit:  "m",
				Value: c.MemoryReservation,
			}
		}

		if len(c.Environment) > 0 {
			environments := make([]types.ContainerEnvironment, len(c.Environment))
			for i, e := range c.Environment {
				environments[i] = types.ContainerEnvironment{
					Name:  e.Name,
					Value: e.Value,
				}
			}
			containerSpec.Environments = environments
		}

		if len(c.PortMappings) > 0 {
			containerPorts := make([]types.ContainerPort, len(c.PortMappings))
			for i, p := range c.PortMappings {
				containerPorts[i] = types.ContainerPort{
					HostPort:      p.HostPort,
					ContainerPort: p.ContainerPort,
				}
			}
			containerSpec.Ports = containerPorts
		}

		if len(c.Links) > 0 {
			containerSpec.Links = c.Links
		}

		if len(c.MountPoints) > 0 {
			volumes := make([]types.ContainerVolume, len(c.MountPoints))
			for i, mp := range c.MountPoints {
				volumes[i] = types.ContainerVolume{
					Source:   mp.SourceVolume,
					Target:   mp.ContainerPath,
					ReadOnly: mp.ReadOnly,
				}
			}
			containerSpec.Volumes = volumes
		}

		if len(c.Ulimits) > 0 {
			limits := make([]types.ContainerUlimit, len(c.Ulimits))

			for i, l := range c.Ulimits {
				limits[i] = types.ContainerUlimit{
					Name:      l.Name,
					SoftLimit: l.SoftLimit,
					HardLimit: l.HardLimit,
				}
			}
			containerSpec.Ulimits = limits
		}

		if len(c.Commands) > 0 {
			containerSpec.Commands = c.Commands
		}

		if c.Privileged {
			containerSpec.Privileged = c.Privileged
		}

		containerSpecs[c.Name] = containerSpec
	}

	return &types.CommonSpecs{
		VolumeSpecs:    volumeSpecs,
		ContainerSpecs: containerSpecs,
	}, nil
}

func (eb *Eb) Output(specs *types.CommonSpecs) ([]byte, error) {
	// TODO: implement this method
	return nil, errors.New("Eb.Input: Not Implement method error")
}
