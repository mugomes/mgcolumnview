// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package mgcolumnview

import (
	"image/color"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type rowData struct {
	id   int
	data []string
}

type ColumnView struct {
	widget.BaseWidget
	headers        []string
	widths         []float32
	data           []rowData
	selected       map[int]bool
	nextID         int
	enableCheck    bool
	selectAllCheck *widget.Check

	header *fyne.Container
	body   *fyne.Container
	scroll *container.Scroll
}

// NewColumnView cria o componente com cabeçalhos e dados
func NewColumnView(headers []string, widths []float32, enableCheck bool) *ColumnView {
	cv := &ColumnView{
		headers:  headers,
		selected: make(map[int]bool),
		widths:   widths,
	}
	cv.enableCheck = enableCheck
	cv.ExtendBaseWidget(cv)
	cv.build()
	return cv
}

// build constrói os elementos principais
func (cv *ColumnView) build() {
	cv.header = cv.makeHeader()
	cv.body = cv.makeBody()
	cv.scroll = container.NewVScroll(cv.body)
}

// CreateRenderer implementa o renderer
func (cv *ColumnView) CreateRenderer() fyne.WidgetRenderer {
	content := container.NewBorder(cv.header, nil, nil, nil, cv.scroll)
	return widget.NewSimpleRenderer(content)
}

// SetData atualiza os dados
// func (cv *ColumnView) SetData(data [][]string) {
// 	cv.data = data
// 	cv.RefreshBody()
// }

// OnToggle define o callback ao marcar/desmarcar
// func (cv *ColumnView) OnToggle(f func(row int, checked bool, data []string)) {
// 	cv.onToggle = f
// }

// AddRow
func (cv *ColumnView) AddRow(row []string) {
	// adiciona strings vazias se faltar colunas
	for len(row) < len(cv.headers) {
		row = append(row, "")
	}
	r := rowData{
		id:   cv.nextID,
		data: row,
	}
	cv.nextID++
	cv.data = append(cv.data, r)
	cv.selected[r.id] = false
	cv.RefreshBody()
}

// RemoveSelected remove todas as linhas que estão marcadas
func (cv *ColumnView) RemoveSelected() {
	newData := []rowData{}
	newSelected := make(map[int]bool)
	for _, r := range cv.data {
		if !cv.selected[r.id] {
			newSelected[r.id] = false
			newData = append(newData, r)
		}
	}
	cv.data = newData
	cv.selected = newSelected
	cv.RefreshBody()
}

// Selected retorna todas as linhas que estão marcadas (checkbox = true)
func (cv *ColumnView) ListSelected() [][]string {
	var selectedData [][]string
	for _, row := range cv.data {
		if cv.selected[row.id] {
			// copia para não expor referência direta
			copied := make([]string, len(row.data))
			copy(copied, row.data)
			selectedData = append(selectedData, copied)
		}
	}
	return selectedData
}

// Selected retorna todas as linhas que estão marcadas (checkbox = true)
func (cv *ColumnView) ListAll() [][]string {
	result := make([][]string, len(cv.data))
	for i, r := range cv.data {
		// cria uma cópia para evitar alterações externas
		copied := make([]string, len(r.data))
		copy(copied, r.data)
		result[i] = copied
	}
	return result
}

func (cv *ColumnView) RemoveAll() {
	cv.data = []rowData{}            // limpa todas as linhas
	cv.selected = make(map[int]bool) // limpa seleção
	cv.nextID = 0                    // reseta o contador de IDs
	if cv.selectAllCheck != nil {
		cv.selectAllCheck.SetChecked(false) // desmarca checkbox "Selecionar Todos"
	}
	cv.RefreshBody()
}

// Novo campo no ColumnView
func (cv *ColumnView) makeHeader() *fyne.Container {
	cells := []fyne.CanvasObject{}

	// checkbox "Selecionar Todos"
	if cv.enableCheck {
		cv.selectAllCheck = widget.NewCheck("", func(checked bool) {
			for _, row := range cv.data {
				cv.selected[row.id] = checked
			}
			cv.RefreshBody()
		})
		cells = append(cells, cv.selectAllCheck)
	}

	for colIndex, h := range cv.headers {
		col := colIndex
		btn := widget.NewButton(h, func() {
			cv.sortByColumn(col)
		})
		btn.Importance = widget.LowImportance // remove borda/efeito de botão
		btn.Alignment = widget.ButtonAlignLeading
		cells = append(cells, container.NewStack(btn))
	}

	var bgColor color.Color
	rect := canvas.NewRectangle(bgColor)
	rect.SetMinSize(fyne.NewSize(400, 28))

	grid1 := container.New(&fixedColumnsLayout{colWidths: cv.widths}, cells...)
	//line := container.NewStack(rect, container.NewHBox(cells...))
	return container.NewHBox(grid1)
}

// Sort
func (cv *ColumnView) sortByColumn(col int) {
	sort.SliceStable(cv.data, func(i, j int) bool {
		a, b := "", ""
		if col < len(cv.data[i].data) {
			a = cv.data[i].data[col]
		}
		if col < len(cv.data[j].data) {
			b = cv.data[j].data[col]
		}
		return a < b
	})
	cv.RefreshBody()
}

// RefreshBody reconstrói apenas o corpo (linhas)
func (cv *ColumnView) RefreshBody() {
	cv.body.Objects = cv.makeBody().Objects
	cv.body.Refresh()
}

// makeHeader cria a linha de cabeçalhos
// func (cv *ColumnView) makeHeader() *fyne.Container {
// 	cells := []fyne.CanvasObject{widget.NewLabel("")} // espaço do checkbox
// 	for _, h := range cv.headers {
// 		lbl := widget.NewLabelWithStyle(h, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
// 		// lbl.TextStyle = fyne.TextStyle{Bold: true}
// 		cells = append(cells, container.NewStack(lbl))
// 	}
// 	return container.NewHBox(cells...)
// }

// makeBody cria as linhas com checkbox e conteúdo

func (cv *ColumnView) makeBody() *fyne.Container {
	rows := []fyne.CanvasObject{}

	for _, r := range cv.data {
		row := r        // captura por cópia
		rowID := row.id // captura o ID

		// background
		var bgColor color.Color
		rect := canvas.NewRectangle(bgColor)
		rect.SetMinSize(fyne.NewSize(400, 28))

		// checkbox da linha
		var cells = []fyne.CanvasObject{}
		if cv.enableCheck {
			check := widget.NewCheck("", func(checked bool) {
				cv.selected[rowID] = checked
				// if cv.onToggle != nil {
				// 	// cria cópia da row.data para evitar problemas
				// 	copied := make([]string, len(row.data))
				// 	copy(copied, row.data)
				// 	// cv.onToggle(rowID, checked, copied)
				// }
			})
			check.Checked = cv.selected[rowID]

			cells = []fyne.CanvasObject{check}
		}

		// preenche células da linha
		for colIndex := 0; colIndex < len(cv.headers); colIndex++ {
			var cellText string
			if colIndex < len(row.data) {
				cellText = row.data[colIndex]
			}
			lbl := widget.NewLabel(cellText)
			lbl.Alignment = fyne.TextAlignLeading
			lbl.Truncation = fyne.TextTruncateEllipsis
			cells = append(cells, lbl)
		}

		grid1 := container.New(&fixedColumnsLayout{colWidths: cv.widths}, cells...)
		line := container.NewStack(rect, container.NewVBox(grid1))
		rows = append(rows, line)
	}

	return container.NewVBox(rows...)
}
