package neweb_pay

import (
	"bytes"
	"github.com/sony/sonyflake"
	"html/template"
	"strconv"
)

func PtrNilString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func GenSonyflake() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return ""
	}
	return strconv.FormatUint(id, 16)
}

var OrderTemplateText = `<form id="order_form" action="{{.Action}}" enctype="application/x-www-form-urlencoded" method="POST">
{{range $key,$element := .Values}}    <input type="hidden" name="{{$key}}" id="{{$key}}" value="{{$element}}" />
{{end -}}
</form>
<script>document.querySelector("#order_form").submit();</script>`

type OrderTmplArgs struct {
	Values map[string]string
	Action string
}

var OrderTmpl = template.Must(template.New("AutoPostOrder").Parse(OrderTemplateText))

func GenerateAutoSubmitHtmlForm(params map[string]string, targetUrl string) string {

	var result bytes.Buffer
	err := OrderTmpl.Execute(&result, OrderTmplArgs{
		Values: params,
		Action: targetUrl,
	})
	if err != nil {
		panic(err)
	}
	return result.String()
}
