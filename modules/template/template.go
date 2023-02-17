package template

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/Mrpye/evaluate-js-templates-api/modules/body_types"
	"github.com/Mrpye/golib/lib"
)

func ParseTemplate(template_data body_types.Template) (string, error) {
	//*********************
	//Create a function map
	//*********************
	funcMap := template.FuncMap{
		"base64enc":   lib.Base64EncString,
		"base64dec":   lib.Base64DecString,
		"gzip_base64": lib.GzipBase64String,
		"lc":          strings.ToLower,
		"uc":          strings.ToUpper,
		"domain":      lib.GetDomainOrIP,
		"port":        lib.GetPortString,
		"port_int":    lib.GetPortInt,
		"clean":       lib.Clean,
		"concat":      lib.Concat,
		"replace":     strings.ReplaceAll,
		"contains":    lib.CommaListContainsString,
		"not":         lib.NOT,
		"or":          lib.OR,
		"and":         lib.AND,
		"plus":        lib.Plus,
		"minus":       lib.Minus,
		"multiply":    lib.Multiply,
		"divide":      lib.Divide,
	}

	//*****************
	//Pase the template
	//*****************
	tpl := strings.ReplaceAll(template_data.Template, "<%", "{{")
	tpl = strings.ReplaceAll(tpl, "%>", "}}")
	tmpl, err := template.New("CodeRun").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return "", err
	}

	//**************************************
	//Run the template to verify the output.
	//**************************************
	var tpl_buffer bytes.Buffer
	err = tmpl.Execute(&tpl_buffer, template_data.Model)
	if err != nil {
		return "", err
	}

	return tpl_buffer.String(), nil
}
