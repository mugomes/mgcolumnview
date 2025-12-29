// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

package mgcolumnview

import "fyne.io/fyne/v2"

// layout personalizado com colunas fixas
type fixedColumnsLayout struct {
	colWidths []float32
}

func (l *fixedColumnsLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	x := float32(0)
	for i, o := range objects {
		if i >= len(l.colWidths) {
			break
		}
		w := l.colWidths[i]
		o.Resize(fyne.NewSize(w, size.Height))
		o.Move(fyne.NewPos(x, 0))
		x += w
	}
}

func (l *fixedColumnsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var totalWidth float32
	var maxHeight float32
	for i, o := range objects {
		if i >= len(l.colWidths) {
			break
		}
		totalWidth += l.colWidths[i]
		h := o.MinSize().Height
		if h > maxHeight {
			maxHeight = h
		}
	}
	return fyne.NewSize(totalWidth, maxHeight)
}