package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App
var WorkTime = 30
var RestTime = 5

func main() {
	a = app.New()
	icon, _ := fyne.LoadResourceFromURLString("http://127.0.0.1/standing.png")
	a.SetIcon(icon)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("For The Sex",
			fyne.NewMenuItem("Show", func() {
				fmt.Println("show setting")
				openSetting()
			}))
		desk.SetSystemTrayMenu(m)
	}
	time.AfterFunc(time.Duration(time.Minute*time.Duration(WorkTime)), showAlert)

	fmt.Println("running")
	a.Run()
}

func showAlert() {
	laterTime := RestTime
	w := a.NewWindow("For The Sex!")

	btnStand := widget.NewButton("Stand!", func() {
		fmt.Println("Stand")
		w.Close()
		countReturnWork()
	})
	separtor := widget.NewSeparator()

	btnLater := widget.NewButton(fmt.Sprintf("%d Minutes Later!", laterTime), func() {
		fmt.Println("Later")
		w.Close()
		time.AfterFunc(time.Duration(time.Minute*time.Duration(laterTime)), showAlert)
	})

	laterSlider := widget.NewSlider(1, 120)
	laterSlider.Value = float64(laterTime)
	laterSlider.OnChanged = func(i float64) {
		btnLater.SetText(fmt.Sprintf("%.0f Minutes Later!", i))
		laterTime = int(i)
	}

	title := widget.NewLabel("Stand Up!!!")
	title.Alignment = fyne.TextAlignCenter

	view := container.NewVBox(title, btnStand, separtor, btnLater, laterSlider)
	w.SetCloseIntercept(func() { fmt.Print("close") })
	w.SetContent(view)
	w.Resize(fyne.NewSize(333, 100))
	w.CenterOnScreen()
	w.Show()
}

func countReturnWork() {
	w := a.NewWindow("Walk Around...")
	counting := widget.NewLabel("Counting")
	counting.Alignment = fyne.TextAlignCenter
	view := container.NewVBox(counting)

	timeRemain := RestTime * 60
	go func() {
		for range time.Tick(time.Second) {
			if timeRemain > 0 {
				timeRemain--
				counting.SetText(int2Duration(timeRemain))
			} else {
				counting.SetText("Work! Work!")
				view.Add(widget.NewButton("Sit!", func() {
					fmt.Println("sit")
					time.AfterFunc(time.Duration(time.Minute*time.Duration(WorkTime)), showAlert)
					w.Close()
				}))
				break
			}
		}
	}()

	w.SetContent(view)
	w.Resize(fyne.NewSize(300, 100))
	w.CenterOnScreen()
	w.SetCloseIntercept(func() { fmt.Print("close") })
	w.Show()

}

func openSetting() {
	w := a.NewWindow("Settings")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.Resize(fyne.NewSize(300, 200))
	w.CenterOnScreen()
	w.Show()
}

// func checkSingleton() (windows.Handle, error) {
//     path, err := os.Executable()
//     if err != nil {
//         return 0, err
//     }
//     hashName := md5.Sum([]byte(path))
//     name, err := syscall.UTF16PtrFromString("Local\\" + hex.EncodeToString(hashName[:]))
//     if err != nil {
//         return 0, err
//     }
//     return windows.CreateMutex(nil, false, name)
// }

func int2Duration(i int) string {
	return fmt.Sprintf("Time Remain: %02d:%02d", i/60, i%60)
}
