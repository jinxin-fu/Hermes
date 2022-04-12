package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HermesdModel = (*customHermesdModel)(nil)

type (
	// HermesdModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHermesdModel.
	HermesdModel interface {
		hermesdModel
	}

	customHermesdModel struct {
		*defaultHermesdModel
	}
)

// NewHermesdModel returns a model for the database table.
func NewHermesdModel(conn sqlx.SqlConn, c cache.CacheConf) HermesdModel {
	return &customHermesdModel{
		defaultHermesdModel: newHermesdModel(conn, c),
	}
}
