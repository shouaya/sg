package main

import (
    "fmt"
    "os"
    "github.com/go-vgo/robotgo"
)

var pad_x int
var pad_y int
var api_url string

func move(x, y string){
    //打开地图
    intX := toI(getCell("config", "C3"))
    intY := toI(getCell("config", "D3"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", false)
    robotgo.Sleep(1)

    //输入坐标x
    intX = toI(getCell("config", "C4"))
    intY = toI(getCell("config", "D4"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", false)
    for i := 0; i < 5; i++ {
        robotgo.KeyTap("backspace")
    }
    robotgo.TypeStr(x)
    robotgo.Sleep(1)

    //输入坐标y
    intX = toI(getCell("config", "C5"))
    intY = toI(getCell("config", "D5"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", false)
    for i := 0; i < 5; i++ {
        robotgo.KeyTap("backspace")
    }
    robotgo.TypeStr(y)
    robotgo.Sleep(1)

    //点击移动
    intX = toI(getCell("config", "C6"))
    intY = toI(getCell("config", "D6"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", false)
    robotgo.Sleep(1)
}

func take(command []string){
    //从地图移动到相应位置
    move(command[2], command[3])

    //点击占领
    intX := toI(getCell("config", "C7"))
    intY := toI(getCell("config", "D7"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", false)
    robotgo.Sleep(1)
    
    //选择部队
    intX = toI(getCell("config", "C8"))
    intY = toI(getCell("config", "D8"))
    id := toI(command[1])
    if id != 0 {
        intX = toI(getCell("config", "C8")) + (id - 3)*255
    }
    robotgo.Move(intX, intY)
    robotgo.Click("left", false)
    robotgo.Sleep(1)
    
    //点击确认
    intX = toI(getCell("config", "C9"))
    intY = toI(getCell("config", "D9"))
    robotgo.Move(intX, intY)
    robotgo.Click("left", false)
    robotgo.Sleep(toI(command[5]))
}

type Point struct {
    X int `json:"x"`
    Y int `json:"y"`
}

func findGameIcon() {
    //get game start xy
    img := robotgo.CaptureImg()
    //imgo.Save("start.png", img)
    robotgo.SaveJpeg(img, "start.jpeg", 10)
    xy := Point{}
    fileUpload(api_url + "start", "start.jpeg", &xy)

    sx, sy := robotgo.GetScreenSize()
    fmt.Println("屏幕分辨利率: ", sx, sy)
    if(xy.Y < sy - 30){
        robotgo.Alert("error", "please open the game")
        os.Exit(3)
    }
    
    robotgo.Move(xy.X, xy.Y)
    robotgo.Click("left", false)
    robotgo.Sleep(2)
}

func adjustPad() {
    //get game start xy
    img := robotgo.CaptureImg()
    //imgo.Save("adjust.png", img)
    robotgo.SaveJpeg(img, "adjust.jpeg", 10)

    xy := Point{}
    fileUpload(api_url + "adjust", "adjust.jpeg", &xy)
    pad_x = xy.X
    pad_y = xy.Y
}

func initWindow() {
    //find back button
    img := robotgo.CaptureImg(pad_x, pad_y, 1280, 750)
    //imgo.Save("game.png", img)
    robotgo.SaveJpeg(img, "game.jpeg", 10)
    xy := Point{}
    fileUpload(api_url + "back", "game.jpeg", &xy)

    if((xy.X > 1100 && xy.X < 1200)|| (xy.Y > 30 && xy.Y < 60)){
        robotgo.Move(xy.X + pad_x + 10, xy.Y + pad_y + 10)
        robotgo.Click("left", false)
        robotgo.Sleep(2)
    }
}

func run() {

    //api_url = "https://sg.weget.jp/"
    api_url = "http://localhost:8080/"

    fmt.Println("获取game图标位置并点开")
    findGameIcon()

    fmt.Println("获取game窗口与屏幕间距")
    adjustPad()

    fmt.Println("初期话游戏界面")
    initWindow()
    os.Exit(3)
    //1280 750
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