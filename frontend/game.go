package main

import (
    "fmt"
    "os"
    "github.com/go-vgo/robotgo"
    "github.com/vcaesar/imgo"
)

var pad_x int
var pad_y int

func move(x, y string){
    //打开地图
    intX := toI(getCell("config", "C3"))
    intY := toI(getCell("config", "D3"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", true)
    robotgo.Sleep(1)

    //输入坐标x
    intX = toI(getCell("config", "C4"))
    intY = toI(getCell("config", "D4"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", true)
    for i := 0; i < 5; i++ {
        robotgo.KeyTap("backspace")
    }
    robotgo.TypeStr(x)
    robotgo.Sleep(1)

    //输入坐标y
    intX = toI(getCell("config", "C5"))
    intY = toI(getCell("config", "D5"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", true)
    for i := 0; i < 5; i++ {
        robotgo.KeyTap("backspace")
    }
    robotgo.TypeStr(y)
    robotgo.Sleep(1)

    //点击移动
    intX = toI(getCell("config", "C6"))
    intY = toI(getCell("config", "D6"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", true)
    robotgo.Sleep(1)
}

func take(command []string){
    //从地图移动到相应位置
    move(command[2], command[3])

    //点击占领
    intX := toI(getCell("config", "C7"))
    intY := toI(getCell("config", "D7"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", true)
    robotgo.Sleep(1)
    
    //选择部队
    intX = toI(getCell("config", "C8"))
    intY = toI(getCell("config", "D8"))
    id := toI(command[1])
    if id != 0 {
        intX = toI(getCell("config", "C8")) + (id - 3)*255
    }
    robotgo.Move(intX, intY)
    robotgo.Click("left", true)
    robotgo.Sleep(1)
    
    //点击确认
    intX = toI(getCell("config", "C9"))
    intY = toI(getCell("config", "D9"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", true)
    robotgo.Sleep(toI(command[5]))
}

type Point struct {
    X int `json:"x"`
    Y int `json:"y"`
}

func start() {
    //get game start xy
    img := robotgo.CaptureImg()
    imgo.Save("start.png", img)
    startxy := Point{}
    fileUpload("http://localhost:8080/start", "start.png", &startxy)
    robotgo.Move(startxy.X, startxy.Y)
    robotgo.Click("left", true)
    robotgo.Sleep(2)
}

func adjust() {
    //get game start xy
    img := robotgo.CaptureImg()
    imgo.Save("adjust.png", img)
    adjustxy := Point{}
    fileUpload("http://localhost:8080/adjust", "adjust.png", &adjustxy)
    fmt.Println(adjustxy)
    pad_x = adjustxy.X
    pad_y = adjustxy.Y
}


func run() {
    // 程序初始位置
    start()
    adjust()
    os.Exit(3)

    // 获取所有指令
    commands := getRows("command")
    for _, command := range commands {
        for _, colCell := range command {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
        switch command[0] {
            case "占领":
                take(command)
            case "移动":
                fmt.Println("not support")
            case "回城":
                fmt.Println("not support")
            default:
                fmt.Println("not support")
        }
    }
}