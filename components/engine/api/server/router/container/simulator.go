package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
)

var inspectRepo = "/var/lib/simdocker/"

func simContainerPs() ([]*types.Container, error) {
	containers, err := ioutil.ReadDir(inspectRepo)
	if err != nil {
		return nil, err
	}
	var contList = make([]*types.Container, len(containers))
	var idx = 0
	for _, cont := range containers {
		var pathBuffer bytes.Buffer
		pathBuffer.WriteString(inspectRepo)
		pathBuffer.WriteString(cont.Name())
		inspectData, _ := simReadInpsect(pathBuffer.String())
		var contJson = types.ContainerJSON{}
		json.Unmarshal(inspectData, &contJson)

		var contType types.Container
		contType.ID = contJson.ID
		contType.Names = []string{contJson.Name}
		contType.Image = contJson.Config.Image
		contType.Status = contJson.State.Status
		var command bytes.Buffer
		command.WriteString(contJson.Path)
		command.WriteString(" ")
		args := contJson.Args
		for _, arg := range args {
			command.WriteString(arg)
			command.WriteString(" ")
		}
		// var tmpPorts []types.Port
		contType.Command = command.String()
		cTime, _ := time.Parse(time.RFC3339Nano, contJson.Created)
		contType.Created = cTime.Unix()
		// contList = append(contList, &contType)
		contList[idx] = &contType
		idx++
	}
	return contList, nil
	// var cont types.Container
	// cont.ID = "12345"
	// cont.Names = []string{"confuse-cont"}
	// cont.Image = "myimage"
	// cont.ImageID = "myimage-id"
	// cont.Command = "command"
	// cont.Created = 1513629086
	// cont.State = "Running"
	// cont.Status = "Up 47 hours"

	// contList[0] = &cont

}

func simContainerInpsect(name string) ([]byte, error) {
	containers, err := ioutil.ReadDir(inspectRepo)
	if err != nil {
		return nil, err
	}
	for _, contFile := range containers {
		if strings.Compare(contFile.Name(), name) == 0 {
			var pathBuffer bytes.Buffer
			pathBuffer.WriteString(inspectRepo)
			pathBuffer.WriteString(contFile.Name())
			inspectData, _ := simReadInpsect(pathBuffer.String())
			return inspectData, nil
		}
	}
	return nil, nil
}

func simReadInpsect(filepath string) ([]byte, error) {
	inspectData, err := ioutil.ReadFile(filepath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return inspectData, nil
}
