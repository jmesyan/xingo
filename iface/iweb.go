package iface

import (
	"github.com/celrenheit/lion"
)

type IWeb interface {
	Ready(*lion.Router)
}
