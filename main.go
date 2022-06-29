// Demo code for the Grid primitive.
package main

import (
        "strings"
        "os/exec"
	"github.com/rivo/tview"
        "log"
        "os"
        // "fmt"
        "github.com/gdamore/tcell/v2"
)


func main() {
// init the app (usefull for closing it lol)

    app := tview.NewApplication()
    
    main := tview.NewList()   //the input field (to add or remove packages from the selected venv)
    menu := tview.NewList()
    //inputs : 
    inputField := tview.NewInputField().
                  SetLabel("add package: ")
  inputField.SetDoneFunc(func(key tcell.Key) {
          a,_ := menu.GetItemText(menu.GetCurrentItem())
          dirname , err := os.UserHomeDir()
          out, err := exec.Command(dirname+"/.config/lazyvenv/"+a+"/bin/pip","install" , inputField.GetText()).Output()
            if err != nil {
              log.Fatal(err)
          }
          print(out)
    })

    venvInput := tview.NewInputField().
                  SetLabel("add venv: ")
    venvInput.SetDoneFunc(func(key tcell.Key) {
          a,_ := menu.GetItemText(menu.GetCurrentItem())             
          dirname , err := os.UserHomeDir()
          out, err := exec.Command(dirname+"/.config/lazyvenv/"+a+"/bin/pip","freeze").Output()
            if err != nil {
              log.Fatal(err)
          }
          print(out)
    })


    // // title & helper
    // title  = tview.TextView("LazyVenv")
    // helper = tview.TextView("mouse -> select items / C-c -> exit / d -> delete venv/package ")


// get the venvs
    
    files,err := os.ReadDir("/home/mehdi/.config/lazyvenv/")
    if err != nil{
      log.Fatal(err)
    }

    //add the items to the list
    c := 'a'
    c2 := 'a' 
    for _,file := range files {
        menu.AddItem(file.Name(), "",c , func(){
        c2 = 'a' 
        a,_ := menu.GetItemText(menu.GetCurrentItem())
        // inputField.SetText(a)

        dirname , err := os.UserHomeDir()
        out, err := exec.Command(dirname+"/.config/lazyvenv/"+a+"/bin/pip","freeze").Output()
        if err != nil {
            log.Fatal(err)
        }
        main.Clear()
        for _,pack := range strings.Split(string(out),"\n") {

          main.AddItem(pack, "",c2 ,nil) 
          c2++
        }

        })       
     c++ 
   }

 
    main.SetBackgroundColor(tcell.ColorDefault).SetTitle("packages")
    menu.SetBackgroundColor(tcell.ColorDefault).SetTitle("venvs")
    inputField.SetBackgroundColor(tcell.ColorDefault)
    venvInput.SetBackgroundColor(tcell.ColorDefault)

    // title.SetBackgroundColor(tcell.ColorDefault)
    // help.SetBackgroundColor(tcell.ColorDefault)

    //add the inputField to the grid
    grid := tview.NewGrid().
            SetRows( 0, 1).
            SetColumns(17,-1).
            SetBorders(true). 
            AddItem(inputField, 1, 1, 1, 1, 0, 0, false).
            AddItem(venvInput, 1, 0, 1, 1, 0, 0, false)

    // add the venv list & the the packages list to the menu
    // Layout for screens wider than 100 cells.
    grid.AddItem(menu, 0, 0, 1, 1, 0, 100, false).
         AddItem(main, 0, 1, 1, 1, 3, 90, true)


    grid.SetBackgroundColor(tcell.ColorDefault)
    // Key presses : 

    main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
      switch event.Rune(){
      case 'd':
        print("deleted an item")
      }
      return event
    })

    menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
      switch event.Rune(){
      case 'd':
        print("deleted a venv")

      case 'c':
        print("cloning a venv")
      }
      return event
    })
      

         
    // start the app
    if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
  
            panic(err)
    }

  }
