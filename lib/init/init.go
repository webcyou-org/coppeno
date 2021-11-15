package initialization

import (
	"github.com/webcyou-org/coppeno/utils"
)

func Start() error {
	return InitCoppenoJson()
}

func InitCoppenoJson() error {
	return utils.CreateCoppenoJson()
}
