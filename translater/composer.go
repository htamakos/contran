package translater

import (
	"errors"
	"strconv"

	"github.com/htamakos/contran/types"
	"gopkg.in/yaml.v2"
)

const VERSION string = "2.4"

type Composer struct{}

func (c *Composer) Input(values []byte) (*types.CommonSpecs, error) {
	// TODO: implement this method
	return nil, errors.New("Composer.Input: Not Implement method error")
}

func (c *Composer) Output(specs *types.CommonSpecs) ([]byte, error) {
	services := map[string]types.ComposerService{}

	for _, v := range specs.ContainerSpecs {
		service := types.ComposerService{
			Name:  v.Name,
			Image: v.Image,
		}

		if v.User != "" {
			service.User = v.User
		}

		if v.Cpu.Shares != 0 {
			service.CpuShares = v.Cpu.Shares
		}

		if v.MemoryLimit.Value != 0 {
			service.MemLimit = v.MemoryLimit.String()
		}

		if v.MemoryReservation.Value != 0 {
			service.MemReservation = v.MemoryReservation.String()
		}

		if len(v.Commands) > 0 {
			service.Commands = v.Commands
		}

		if len(v.Environments) > 0 {
			environments := map[string]string{}
			for _, e := range v.Environments {
				environments[e.Name] = e.Value
			}
			service.Environment = environments
		}

		if len(v.Ports) > 0 {
			ports := make([]string, len(v.Ports))

			for i, p := range v.Ports {
				port := ""
				if p.HostIp != "" {
					port = p.HostIp + ":"
				}
				port = port + strconv.Itoa(p.HostPort) + ":" + strconv.Itoa(p.ContainerPort)
				ports[i] = port
			}
			service.Ports = ports
		}

		if len(v.Links) > 0 {
			service.Links = v.Links
		}

		if len(v.Volumes) > 0 {
			volumes := make([]string, 0)
			for _, vol := range v.Volumes {
				sourcePath, ok := specs.VolumeSpecs[vol.Source]
				if !ok {
					continue
				}
				volStr := sourcePath.Path + ":" + vol.Target
				if vol.ReadOnly {
					volStr = volStr + ":ro"
				}
				volumes = append(volumes, volStr)
			}
			service.Volumes = volumes
		}

		if len(v.Ulimits) > 0 {
			limits := make(map[string]interface{})
			for _, l := range v.Ulimits {
				limits[l.Name] = struct {
					Soft int
					Hard int
				}{
					Soft: l.SoftLimit,
					Hard: l.HardLimit,
				}
			}

			service.Ulimits = limits
		}

		if v.Privileged {
			service.Privileged = v.Privileged
		}

		services[v.Name] = service
	}

	outputs, err := yaml.Marshal(types.ComposerConfig{
		Version:  VERSION,
		Services: services,
	})

	if err != nil {
		return nil, err
	}

	return outputs, nil
}
