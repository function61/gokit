package udocker

const (
	SwarmServiceNameLabelKey = "com.docker.swarm.service.name"
)

// stupid Docker requires "inspect" to get actually interesting details
type ContainerListItem struct {
	Id              string            `json:"Id"`
	Names           []string          `json:"Names"`
	Image           string            `json:"Image"`
	Labels          map[string]string `json:"Labels"`
	NetworkSettings struct {
		Networks map[string]struct {
			IPAddress string `json:"IPAddress"`
		} `json:"Networks"`
	} `json:"NetworkSettings"`
}

type Container struct {
	Id     string          `json:"Id"`
	Name   string          `json:"Name"`
	Image  string          `json:"Image"`
	Config ContainerConfig `json:"Config"`
	Mounts []Mount         `json:"Mounts"`
}

type Mount struct {
	Type   string `json:"Type"`
	Name   string `json:"Name"`
	Source string `json:"Source"`
	Driver string `json:"Driver"`
}

type ContainerConfig struct {
	Labels map[string]string `json:"Labels"`
	Env    []string          `json:"Env"`
}

type Service struct {
	ID   string      `json:"ID"`
	Spec ServiceSpec `json:"Spec"`
}

type ServiceSpec struct {
	Name         string                  `json:"Name"`
	TaskTemplate ServiceSpecTaskTemplate `json:"TaskTemplate"`
	Labels       map[string]string       `json:"Labels"`
}

type ServiceSpecTaskTemplate struct {
	ContainerSpec ServiceSpecTaskTemplateContainerSpec `json:"ContainerSpec"`
}

type ServiceSpecTaskTemplateContainerSpec struct {
	Image string   `json:"Image"`
	Env   []string `json:"Env"`
}

type Task struct {
	ID                  string                  `json:"ID"`
	ServiceID           string                  `json:"ServiceID"`
	NodeID              string                  `json:"NodeID"`
	NetworksAttachments []TaskNetworkAttachment `json:"NetworksAttachments"`
}

type TaskNetworkAttachment struct {
	Network   TaskNetworkAttachmentNetwork `json:"Network"`
	Addresses []string                     `json:"Addresses"`
}

type TaskNetworkAttachmentNetwork struct {
	ID   string                           `json:"ID"`
	Spec TaskNetworkAttachmentNetworkSpec `json:"Spec"`
}

type TaskNetworkAttachmentNetworkSpec struct {
	Name string `json:"Name"`
}

type Node struct {
	ID          string          `json:"ID"`
	Description NodeDescription `json:"Description"`
}

type NodeDescription struct {
	Hostname string `json:"Hostname"`
}
