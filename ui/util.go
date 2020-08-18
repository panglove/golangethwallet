package ui

import "fyne.io/fyne"

func SetWidgetHCenter(ob fyne.CanvasObject,container fyne.Size){


	newDelPos := fyne.NewPos( (container.Width-ob.Size().Width) / 2,0)

	ob.Move(ob.Position().Add(newDelPos))


}

func SetWidgetY(ob fyne.CanvasObject,y int){

	x :=ob.Position().X



	ob.Move(fyne.NewPos(x,y))
}

func SetWidgetX(ob fyne.CanvasObject,x int){

	y :=ob.Position().Y



	ob.Move(fyne.NewPos(x,y))
}

func SetWidgetPosition(ob fyne.CanvasObject,x int, y int){


	ob.Move(fyne.NewPos(x,y))
}

