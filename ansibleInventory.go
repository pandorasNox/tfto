package main

import (
	"strings"
)

type AnsibleInventory map[string]Group

type Group struct {
	Hosts map[string]Host `json:"hosts"`
}

type Host struct {
	Ansible_host string `json:"ansible_host"`
	Ansible_user string `json:"ansible_user"`
}

func createAnsibleInventoryHetzner(hetznerServerResources []TFResource) AnsibleInventory {
	AnsibleInventory := make(map[string]Group)

	for _, hr := range hetznerServerResources {
		for _, instance := range hr.Instances {
			//extract group names of hetzner
			aig := instance.Attributes.Labels.AnsibleInventoryGroups
			ansibleGroupNames := strings.Split(aig, ".")

			name := instance.Attributes.Name
			hostAddress := instance.Attributes.Ipv4Address

			for _, groupName := range ansibleGroupNames {
				_, groupExists := AnsibleInventory[groupName]
				if !groupExists {
					newHosts := make(map[string]Host)
					AnsibleInventory[groupName] = Group{Hosts: newHosts}
				}
				AnsibleInventory[groupName].Hosts[name] = Host{hostAddress, "root"}
			}
		}
	}

	return AnsibleInventory
}
