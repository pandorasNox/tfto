package main

type kubelifePatch struct {
	Server kubelifeServers `json:"server"`
}

type kubelifeServers []kubelifeServer

type kubelifeServer struct {
	Name              string `json:"name"`
	Role              string `json:"role"`
	Ipv4AddressPublic string `json:"ipv4_address_public"`
	Ipv6Address       string `json:"ipv6_address"`
	Ipv6Network       string `json:"ipv6_network"`
}

func createKubelifePatch(hetznerServerResources []TFResource) kubelifePatch {
	var kp = kubelifePatch{}

	for _, hr := range hetznerServerResources {
		for _, instance := range hr.Instances {
			name := instance.Attributes.Name
			nodeRole := instance.Attributes.Labels.K8sNodeRole
			ipv4AddressPublic := instance.Attributes.Ipv4Address
			ipv6Address := instance.Attributes.Ipv6Address
			ipv6Network := instance.Attributes.Ipv6Network

			kp.Server = append(kp.Server, kubelifeServer{name, nodeRole, ipv4AddressPublic, ipv6Address, ipv6Network})
		}
	}

	return kp
}
