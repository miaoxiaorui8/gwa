package render

import (
	"gwa_server/render/admin"
)

type RenderGroup struct {
	AdminRenderGroup admin.RenderGroup
}

var RenderGroupHtml = new(RenderGroup)
