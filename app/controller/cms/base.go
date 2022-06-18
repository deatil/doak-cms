package cms

import (
    "github.com/deatil/doak-cms/app/response"
)

// Base
type Base struct{}

func (this *Base) Theme(tpl string) string {
    return response.CmsTheme(tpl)
}
