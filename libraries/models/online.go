package models

import "timingniao_wlx_api/libraries/types"

type Online struct {
	Model
	Type   types.OnlineType
	Online bool
}
