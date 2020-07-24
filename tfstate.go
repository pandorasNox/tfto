package main

type tfstate struct {
	TerraformVersion string       `json:"terraform_version"`
	Version          int64        `json:"version"`
	Resources        []TFResource `json:"resources"`
}

type TFResource struct {
	Name      string `json:"name"`
	Provider  string `json:"provider"`
	Type      string `json:"type"`
	Instances []struct {
		IndexKey      int64 `json:"index_key"`
		SchemaVersion int64 `json:"schema_version"`
		Attributes    struct {
			Datacenter  string `json:"datacenter"`
			ID          string `json:"id"`
			Image       string `json:"image"`
			Ipv4Address string `json:"ipv4_address"`
			Ipv6Address string `json:"ipv6_address"`
			Ipv6Network string `json:"ipv6_network"`
			Labels      struct {
				AnsibleInventoryGroups string `json:"ansible_inventory_groups"`
				K8sNodeRole            string `json:"k8s_node_role"`
			} `json:"labels"`
			Location   string      `json:"location"`
			Name       string      `json:"name"`
			ServerType string      `json:"server_type"`
			SSHKeys    interface{} `json:"ssh_keys"`
			Status     string      `json:"status"`
			UserData   interface{} `json:"user_data"`
		} `json:"attributes"`
	} `json:"instances"`
}
