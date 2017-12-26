package container

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// getContainersByName inspects container's configuration and serializes it as json.
func (s *containerRouter) getContainersByName(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	// displaySize := httputils.BoolValue(r, "size")

	// version := httputils.VersionFromContext(ctx)
	// json, err := s.backend.ContainerInspect(vars["name"], displaySize, version)
	// if err != nil {
	// 	return err
	// }

	inspectData, _ := simContainerInpsect(vars["name"])
	var contType = &types.ContainerJSON{}
	json.Unmarshal(inspectData, &contType)
	// var contList = make([]types.ContainerJSON, 1)
	// contList = append(contList, contType)
	// out, _ := json.Marshal(contType)
	fmt.Println("Container-inspect = ", contType)
	return httputils.WriteJSON(w, http.StatusOK, contType)
}
