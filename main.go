// Demo code for the Grid primitive.
package main

import (
        "os/exec"
	"github.com/rivo/tview"
        "log"
        "os"
        "fmt"
        "github.com/gdamore/tcell/v2"
)

func main() {
// init the app (usefull for closing it lol)

    app := tview.NewApplication()
    
    //to remove later
    newPrimitive := func(text string) tview.Primitive {
            return tview.NewTextView().
                    SetTextAlign(tview.AlignCenter).
                    SetText(text)

    }


    main := newPrimitive("Main content")
    


    //the input field (to add or remove packages from the selected venv)
    inputField := tview.NewInputField().
                  SetLabel("add or remove : ").
            SetFieldBackgroundColor(tcell.ColorBlack)

    title := `
  _
        | | __ _ _____   ___   _____ _ ____   __
        | |/ _' |_  / | | \ \ / / _ \ '_ \ \ / /
       | | (_| |/ /| |_| |\ V /  __/ | | \ V /
      |_|\__,_/___|\__, | \_/ \___|_| |_|\_/
|___/

    ` 

// get the venvs
    menu := tview.NewList()
    
    files,err := os.ReadDir("/home/mehdi/.config/lazyvenv/")
    if err != nil{
      log.Fatal(err)
    }

    //add the items to the list
    c := 'a'
    for _,file := range files {
        menu.AddItem(file.Name(), "",c , func(){

        a,_ := menu.GetItemText(menu.GetCurrentItem())
        inputField.SetText(a)


        out, err := exec.Command("date").Output()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("The date is %s\n", out)
        c++

        })       
      }

  

    //add the inputField to the grid
    grid := tview.NewGrid().
            SetRows(10, 0, 1).
            SetColumns(17,-1).
            SetBorders(true).
            AddItem(newPrimitive(title), 0, 0, 1, 2, 0, 0, false).
            AddItem(inputField, 2, 0, 1, 2, 0, 0, false)


    // add the venv list & the the packages list to the menu
    // Layout for screens wider than 100 cells.
    grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
         AddItem(main, 1, 1, 1, 1, 3, 90, false)



    // start the app
    if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
            panic(err)
    }
}
