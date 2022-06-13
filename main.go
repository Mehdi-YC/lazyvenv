// Demo code for the Grid primitive.
package main

import (
	"github.com/rivo/tview"
        "github.com/gdamore/tcell/v2"
)

func main() {

  //to remove later
    newPrimitive := func(text string) tview.Primitive {
            return tview.NewTextView().
                    SetTextAlign(tview.AlignCenter).
                    SetText(text)

    }





    //add List of venv : 
    menu := tview.NewList().
            AddItem("List item 1", "", 'a', nil).
            AddItem("List item 2", "", 'b', nil).
            AddItem("List item 3", "", 'c', nil).
            AddItem("List item 4", "", 'd', nil)
    main := newPrimitive("Main content")
    



    //the input field (to add or remove packages from the selected venv)
    inputField := tview.NewInputField().
            SetFieldBackgroundColor(tcell.ColorBlack)



    //add the inputField to the grid
    grid := tview.NewGrid().
            SetRows(7, 0, 1).
            SetColumns(17,-1).
            SetBorders(true).
            AddItem(newPrimitive("Header"), 0, 0, 1, 2, 0, 0, false).
            AddItem(inputField, 2, 0, 1, 2, 0, 0, false)


    // add the venv list & the the packages list to the menu
    // Layout for screens wider than 100 cells.
    grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
         AddItem(main, 1, 1, 1, 1, 3, 90, false)




    // start the app
    if err := tview.NewApplication().SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
            panic(err)
    }
}
