package dao

import (
	"github.com/go-mogu/mogu-picture/internal/app/picture/dao/internal"
)

type networkDiskDao struct {
	*internal.NetworkDiskDao
}

var (
	NetworkDisk = networkDiskDao{
		internal.NewNetworkDiskDao(),
	}
)
