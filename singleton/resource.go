package singleton

import rice "github.com/GeertJohan/go.rice"

var resource *Resource

// Resource static file in go
type Resource struct {
	tpl *rice.Box
}

// GetResource Resource
func GetResource() *Resource {
	if resource == nil {
		resource = &Resource{
			tpl: rice.MustFindBox("templates"),
		}
	}
	return resource
}

// GetTpl tpl string
func (r *Resource) GetTpl(name string) string {
	return r.tpl.MustString(name)
}
