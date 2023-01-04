def readFile(file_name, sheet_name)
    config = {
        'start' => [],
        'commands' => [],
    }
    book = Spreadsheet.open(file_name)
    sheet = book.worksheet(sheet_name)
    sheet.each do |row|
      break if row[0].nil?
      if row[0] == '程序初始位置'
        config['start'] = [row[1].to_i, row[2].to_i]
        next
      end
      next if row[0] == '指令'
      config['commands'].push({
        'type' => row[0],
        'id' => row[1].to_i,
        'x' => row[2].to_i,
        'y' => row[3].to_i,
        'wait' => row[5].to_i,
      })
    end
    config
end

def start(w, h)
    mouse_move(w, h)
    sleep 2
    left_click
end

def goto(x,y)
    # 地图
    mouse_move(260,750)
    sleep 2
    left_click

    #输入坐标
    mouse_move(1170,760)
    sleep 2
    left_click

    #输入坐标
    4.times do
        key_stroke(0x08)
    end 
    sleep 2
    left_click

    #输入坐标
    type(x)
    sleep 2
    left_click

    #输入坐标
    mouse_move(1300,760)
    sleep 2
    left_click

    #输入坐标
    4.times do
        key_stroke(0x08)
    end 
    sleep 2
    left_click

    #输入坐标
    type(y)
    sleep 2
    left_click

    #输入坐标
    mouse_move(1430,760)
    sleep 2
    left_click
end

def take_place(x, y)
    goto(x, y)

    #点击占领
    mouse_move(1040,395)
    sleep 2
    left_click
    
    #选择部队
    mouse_move(806,840)
    sleep 2
    left_click
    
    #点击确认
    mouse_move(1190,700)
    sleep 2
    left_click
end

def del_troops(x, y)
    #选最后一支部队
    mouse_move(x,y)
    sleep 2
    left_click

    #退阵
    mouse_move(667,323)
    sleep 2
    left_click

    #确认
    mouse_move(739,699)
    sleep 2
    left_click

    #退阵
    mouse_move(887,323)
    sleep 2
    left_click

    #确认
    mouse_move(739,699)
    sleep 2
    left_click

    #退阵
    mouse_move(1100,323)
    sleep 2
    left_click

    #确认
    mouse_move(739,699)
    sleep 2
    left_click

    #退出
    mouse_move(1383,252)
    sleep 2
    left_click
end

def add_troops(x, y)
    #选最后一支部队
    mouse_move(x,y)
    sleep 2
    left_click

    #上阵
    mouse_move(607,481)
    sleep 2
    left_click

    #按等级倒叙
    2.times do
        mouse_move(710,375)
        sleep 2
        left_click
    end
    
    #出阵
    mouse_move(1406,441)
    sleep 2
    left_click

    #点5次增加500部队
    5.times do
        mouse_move(1119,722)
        sleep 2
        left_click
    end

    #确认
    mouse_move(1295,815)
    sleep 2
    left_click

    #退出部队
    mouse_move(1383,252)
    sleep 2
    left_click

    #退出城市
    mouse_move(1410,243)
    sleep 2
    left_click
end

def back_city(x,y, m)
   #部队所在位置
   goto(x, y)

   #点击部队
   mouse_move(817,533)
   sleep 2
   left_click

   #回城
   mouse_move(1380,700)
   sleep 2
   left_click

   sleep m
end

def change_troops(x, y, m)
   #点坐标
   mouse_move(1393,330)
   sleep 2
   left_click

   #点主城
   mouse_move(1281,380)
   sleep 2
   left_click

   #点入城
   mouse_move(815,539)
   sleep 2
   left_click

   #部队退阵
   del_troops(1322,839)

   #新建部队
   add_troops(1322,839)

   goto(x, y)

   #点击占领
   mouse_move(1040,395)
   sleep 2
   left_click
   
   #选择部队
   mouse_move(1312,834)
   sleep 2
   left_click
   
   #点击确认
   mouse_move(1190,700)
   sleep 2
   left_click

   sleep m
end

def select_troop(id)
   #中间点一下
   mouse_move(700, 460)
   sleep 2
   left_click
   #点击部队
   mouse_move(1244, 272)
   sleep 2
   left_click

   if id <= 3 then
    left_drag(1244, 300, 1244, 400)
    sleep 2
   else
    left_drag(1244, 400, 1244, 300)
    sleep 2
   end

   case id
   when 1
    mouse_move(1273, 316)
   when 2
    mouse_move(1273, 366)
   when 3
    mouse_move(1273, 416)
   when 4
    mouse_move(1273, 436)
   when 5
    mouse_move(1273, 486)
   else
    print('not support')
   end
   left_click
   sleep 2
end
