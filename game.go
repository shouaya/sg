package main

import (
    "fmt"
    "github.com/go-vgo/robotgo"
)

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

func start(x, y int) {
    robotgo.Move(x, y)
    robotgo.Click("left", true)
}

func run() {
    // 程序初始位置
    intStartX := toI(getCell("config", "C2"))
    intStartY := toI(getCell("config", "D2"))
    start(intStartX, intStartY)
    robotgo.Sleep(1)

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