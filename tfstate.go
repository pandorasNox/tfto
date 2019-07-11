package main

type tfstate struct {
	TerraformVersion string `json:"terraform_version"`
	Version          int64  `json:"version"`
	Resources        []struct {
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
				Labels      struct {
					K8sNodeType string `json:"k8s_node_type"`
				} `json:"labels"`
				Location   string      `json:"location"`
				Name       string      `json:"name"`
				ServerType string      `json:"server_type"`
				SSHKeys    interface{} `json:"ssh_keys"`
				Status     string      `json:"status"`
				UserData   interface{} `json:"user_data"`
			} `json:"attributes"`
		} `json:"instances"`
	} `json:"resources"`
}
