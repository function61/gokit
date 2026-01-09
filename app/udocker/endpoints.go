package udocker

var (
	// API version 1.48 = Docker 28.0 = 2025-02
	DefaultVersion = EndpointVersion("1.48")
)

type VersionedEndpoints struct {
	baseURL string
}

// creates "factory" with which you can make URLs like `/v1.44/containers/json`
func EndpointVersion(version string) VersionedEndpoints {
	return VersionedEndpoints{"/v" + version}
}

func (v VersionedEndpoints) ListContainersEndpoint() string { return v.baseURL + "/containers/json" }

func (v VersionedEndpoints) TasksEndpoint() string {
	return v.baseURL + "/tasks?filters=%7B%22desired-state%22%3A%5B%22running%22%5D%7D"
}

func (v VersionedEndpoints) ServicesEndpoint() string { return v.baseURL + "/services" }

func (v VersionedEndpoints) NodesEndpoint() string { return v.baseURL + "/nodes" }

func (v VersionedEndpoints) ContainerInspectEndpoint(containerId string) string {
	return v.baseURL + "/containers/" + containerId + "/json"
}

func (v VersionedEndpoints) NetworkInspectEndpoint(networkId string) string {
	return v.baseURL + "/networks/" + networkId
}
