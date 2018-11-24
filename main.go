package main

import (
	"os"
	"github.com/moonlightnvkz/CarrierDensity/formula"
	//"/formula"

	// "github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func init() {

}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetTitle("bridge Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()
   
    var sigma float64 = formula.FindConductivity()
	fmt.Println(sigma)

	// series := charts.NewQLineSeriesFromPointer(view.RootObject().FindChild("series", core.Qt__FindChildrenRecursively).Pointer())

	app.Exec()
}
