package global

import (
	"github.com/provider-go/pkg/cache"
	"github.com/provider-go/pkg/smcc"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Cache cache.Cache
	SMCC  smcc.SMCC
)
