package ui

import "fyne.io/fyne"

type AbLayout struct {

	Width int
	Height int
}


func (d *AbLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {

	return fyne.NewSize(d.Width, d.Height)
}
func (d *AbLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {

	for _, o := range objects {

		obSize := o.Size()

		if obSize.Width>0 && obSize.Height>0 {
			o.Resize(o.Size())

		}else {
			o.Resize(o.MinSize())
		}
	}

}
