package udocker

import (
	"strings"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestEndpointVersion(t *testing.T) {
	urls := func(version VersionedEndpoints) string {
		return strings.Join([]string{
			version.ListContainersEndpoint(),
			version.ContainerInspectEndpoint("CONTAINER_ID"),
			version.TasksEndpoint(),
			version.ServicesEndpoint(),
			version.NodesEndpoint(),
			version.NetworkInspectEndpoint("NETWORK_ID"),
		}, "\n")
	}

	assert.Equal(t, "\n"+urls(DefaultVersion), `
/v1.48/containers/json
/v1.48/containers/CONTAINER_ID/json
/v1.48/tasks?filters=%7B%22desired-state%22%3A%5B%22running%22%5D%7D
/v1.48/services
/v1.48/nodes
/v1.48/networks/NETWORK_ID`)

	assert.Equal(t, "\n"+urls(EndpointVersion("1.69")), `
/v1.69/containers/json
/v1.69/containers/CONTAINER_ID/json
/v1.69/tasks?filters=%7B%22desired-state%22%3A%5B%22running%22%5D%7D
/v1.69/services
/v1.69/nodes
/v1.69/networks/NETWORK_ID`)
}
