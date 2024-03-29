package mocks

import (
	"github.com/h2non/gock"
	"github.com/phaus/go-proxmox/tests/mocks/config"
	"github.com/phaus/go-proxmox/tests/mocks/pve6x"
	"github.com/phaus/go-proxmox/tests/mocks/pve7x"
	"github.com/phaus/go-proxmox/tests/mocks/pve8x"
)

func On(c config.Config) {
	ProxmoxVE7x(c) // default pve7
}

func Off() {
	gock.Off()
}

func ProxmoxVE8x(c config.Config) {
	config.C = c
	pve8x.Load()
}

func ProxmoxVE7x(c config.Config) {
	config.C = c
	pve7x.Load()
}

func ProxmoxVE6x(c config.Config) {
	config.C = c
	pve6x.Load()
}
