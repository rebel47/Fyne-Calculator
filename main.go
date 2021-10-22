package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))
	output := ""
	input := widget.NewLabel("")
	isHistory := false
	historystr := ""
	history := widget.NewLabel(historystr)
	var historyArr []string
	historyBtn := widget.NewButton("History", func() {
		if isHistory {
			historystr = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historystr = historystr + historyArr[i]
				historystr += "\n"
			}
		}
		isHistory = !isHistory
		history.SetText(historystr)

	})
	backBtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	openBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	closeBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	divideBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	nineBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	multiplyBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	subBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	addBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	decimalBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + "=" + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				output = "Error!!"
			}
		} else {
			output = "Error!!"
		}
		input.SetText(output)
	})
	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				closeBtn,
				divideBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiplyBtn,
			),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				subBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				addBtn,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					decimalBtn),
				equalBtn,
			),
		),
	))

	w.ShowAndRun()
}
