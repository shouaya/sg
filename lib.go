package main

import (
    "fmt"
    "strconv"
    "github.com/xuri/excelize"
    "github.com/go-vgo/robotgo"
    "github.com/vcaesar/imgo"
)

func run() {
    f, err := excelize.OpenFile("steps.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // Get value from cell by given worksheet name and cell reference.
    cellx, err := f.GetCellValue("config", "C2")
    intx, err := strconv.Atoi(cellx)
    celly, err := f.GetCellValue("config", "D2")
    inty, err := strconv.Atoi(celly)
    start(intx, inty)
    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("command")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}

func screen() {
    x, y := robotgo.Location()
    fmt.Println("pos: ", x, y)

    color := robotgo.GetPixelColor(100, 200)
    fmt.Println("color---- ", color)

    sx, sy := robotgo.GetScreenSize()
    fmt.Println("get screen size: ", sx, sy)

    bit := robotgo.CaptureScreen(10, 10, 30, 30)
    defer robotgo.FreeBitmap(bit)

    img := robotgo.ToImage(bit)
    imgo.Save("test.png", img)

    num := robotgo.DisplaysNum()
    for i := 0; i < num; i++ {
        robotgo.DisplayID = i
        img1 := robotgo.CaptureImg()
        path1 := "save_" + strconv.Itoa(i)
        robotgo.Save(img1, path1+".png")
        robotgo.SaveJpeg(img1, path1+".jpeg", 50)

        img2 := robotgo.CaptureImg(10, 10, 20, 20)
        robotgo.Save(img2, "test_"+strconv.Itoa(i)+".png")
    }
}

func start(x, y int) {
    robotgo.Move(x, y)
    robotgo.Click("left", true)
}
