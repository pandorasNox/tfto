package main

type AnsibleInventory map[string]Group

type Group struct {
	Hosts map[string]Host `json:"hosts"`
}

type Host struct {
	Ansible_host string `json:"ansible_host"`
	Ansible_user string `json:"ansible_user"`
}
