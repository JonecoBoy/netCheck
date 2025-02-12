package entity

import pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"

type Address struct {
	id           pkgEntity.ID
	street       string
	number       string
	complement   string
	neighborhood string
	city         string
	state        string
	zip          string
	country      string
}
