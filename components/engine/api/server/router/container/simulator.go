package container

import (
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
)

// Port An open port on a container
// swagger:model Port
type Port struct {

	// IP
	IP string `json:"IP,omitempty"`

	// Port on the container
	// Required: true
	PrivatePort uint16 `json:"PrivatePort"`

	// Port exposed on the host
	PublicPort uint16 `json:"PublicPort,omitempty"`

	// type
	// Required: true
	Type string `json:"Type"`
}

// MountPoint represents a mount point configuration inside the container.
// This is used for reporting the mountpoints in use by a container.
// type MountPoint struct {
// 	Type        mount.Type `json:",omitempty"`
// 	Name        string     `json:",omitempty"`
// 	Source      string
// 	Destination string
// 	Driver      string `json:",omitempty"`
// 	Mode        string
// 	RW          bool
// 	Propagation mount.Propagation
// }

// Container contains response of Engine API:
// GET "/containers/json"
type Container struct {
	ID         string `json:"Id"`
	Names      []string
	Image      string
	ImageID    string
	Command    string
	Created    int64
	Ports      []Port
	SizeRw     int64 `json:",omitempty"`
	SizeRootFs int64 `json:",omitempty"`
	Labels     map[string]string
	State      string
	Status     string
	// HostConfig struct {
	// 	NetworkMode string `json:",omitempty"`
	// }
	//NetworkSettings *SummaryNetworkSettings
	// Mounts          []MountPoint
}

func simContainers() ([]*types.Container, error) {
	var cont types.Container
	cont.ID = "12345"
	cont.Names = []string{"confuse-cont"}
	cont.Image = "myimage"
	cont.ImageID = "myimage-id"
	cont.Command = "command"
	cont.Created = 1513629086
	cont.State = "Running"
	cont.Status = "Up 47 hours"
	var contList = make([]*types.Container, 1)
	contList[0] = &cont
	return contList, nil
}

func simContainerInpsect() (string, error) {
	inspectStr, err := ioutil.ReadFile("/tmp/cont.inspect") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return string(inspectStr), nil
}
