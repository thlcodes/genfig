package plugins

import (
	"io"
	"strings"
	"text/template"

	"github.com/theliebeskind/genfig/types"
)

type envUpdatePlugin struct {
	s   types.SchemaMap
	tpl *template.Template
}

var (
	envUpdater = envUpdatePlugin{
		s: types.SchemaMap{},
		tpl: template.Must(template.
			New("envUpdater").
			Funcs(template.FuncMap{
				"upper": strings.ToUpper,
				"title": strings.Title,
				"cleanPrefixEnv": func(s string) string {
					return strings.Join(strings.Split(s, "_")[1:], "_")
				},
				"makePath": func(s string) string {
					return strings.Join(strings.Split(s, "_")[1:], ".")
				},
				"renameSlice": func(s string) string {
					if strings.HasPrefix(s, "[]") {
						return s[2:] + "Slice"
					}
					return s
				},
			}).
			Parse(`import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func (c *Config) UpdateFromEnv() {
	var val string
	var exists bool
{{range $_, $v := .}}{{if not $v.IsStruct}}
	if val, exists = os.LookupEnv("{{cleanPrefixEnv (upper $v.Path)}}"); exists { {{if eq $v.Content "string"}}
		c.{{makePath $v.Path}} = val {{else}}
		if v, err := parse{{title (renameSlice $v.Content)}}(val); err == nil {
			c.{{makePath $v.Path}} = v
		} else {
			fmt.Printf("Genfig: could not parse {{$v.Content}} from {{upper $v.Path}} ('%s')\n", val)
		} {{end}}
	}
{{end}}{{end}}
}

// these are wrappers, so that they can
// a) be referenced easily be the code generator and
// b) be replaces easily by you (or me)
func parseInt64(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 0)
	return
}

func parseFloat64(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 0)
	return
}

func parseBool(s string) (b bool, err error) {
	b, err = strconv.ParseBool(s)
	return
}

func parseStringSlice(s string) (a []string, err error) {
	err = json.Unmarshal([]byte(s), &a)
	return
}

func parseInt64Slice(s string) (a []int64, err error) {
	err = json.Unmarshal([]byte(s), &a)
	return
}

func parseFloat64Slice(s string) (a []float64, err error) {
	err = json.Unmarshal([]byte(s), &a)
	return
}

func parseInterfaceSlice(s string) (a []interface{}, err error) {
	err = json.Unmarshal([]byte(s), &a)
	return
}
`))}
)

func init() {
	Plugins["env_updater"] = &envUpdater
}

func (p *envUpdatePlugin) SetSchemaMap(s types.SchemaMap) {
	p.s = s
}

func (p *envUpdatePlugin) WriteTo(w io.Writer) (l int64, err error) {
	err = p.tpl.Execute(w, p.s)
	return
}