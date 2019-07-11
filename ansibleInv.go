package main

type ansibleInv struct {
	all struct {
		hosts []host
	}
	master struct {
		hosts []host
	}
	worker struct {
		hosts []host
	}
}

type host struct {
	Ansible_host string `json:"ansible_host"`
	Ansible_user string `json:"ansible_user"`
}
