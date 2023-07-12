package duemail

import "strings"

const (
	prefix = `<style type="text/css">
    table.tftable {font-size:12px;color: #333333;width:100%;border-width: 1px;border-color: #e5d254;border-collapse: collapse;}
    table.tftable th {font-size:12px;background-color: #b8851e;border-width: 1px;padding: 8px;border-style: solid;border-color: #ebe23a;text-align:left;}
    table.tftable tr {background-color:#ffffff;}
    table.tftable td {font-size:12px;border-width: 1px;padding: 8px;border-style: solid;border-color: #cfad30;}
</style>

<table id="tfhover" class="tftable" border="1" lang="zh-cn"><head><meta charset="utf-8"/>`
	trStart = `<tr>`
	trEnd   = `</tr>`
	thStart = `<th>`
	thEnd   = `</th>`
	tdStart = `<td>`
	tdEnd   = `</td>`
)

type Table struct {
	Header  []string
	Content [][]string
}

func MakeTableHtml(table Table) string {
	var content strings.Builder
	content.WriteString(prefix)
	content.WriteString(trStart)
	for i := 0; i < len(table.Header); i++ {
		content.WriteString(thStart)
		content.WriteString(table.Header[i])
		content.WriteString(thEnd)
	}
	content.WriteString(trEnd)
	for i := 0; i < len(table.Content); i++ {
		content.WriteString(trStart)
		for j := 0; j < len(table.Content[i]); j++ {
			content.WriteString(tdStart)
			content.WriteString(table.Content[i][j])
			content.WriteString(tdEnd)
		}
		content.WriteString(trEnd)
	}
	return content.String()
}
