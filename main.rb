require './lib.rb'
require 'auto_click'
require 'spreadsheet'
include AutoClickMethods

$config = readFile('steps.xls', 'config')
# open window
start($config['程序初始位置'][0], $config['程序初始位置'][1])
book = Spreadsheet.open('steps.xls')
sheet = book.worksheet('command')
sheet.each do |command|
    puts command.join(',')
    next if command[0] == '指令'
    #select_troop(command['id'])
    case command[0]
    when '占领'
      take_place(command[2].to_i, command[3].to_i, command[1].to_i)
    when '行军'
      #mouse_move(1273, 366)
    when '回城'
      #back_city(1273, 416)
    when '退阵'
      #del_troops(1273, 436)
    when '上阵'
      #add_troops(1273, 486)
    else
      print('not support')
    end
    sleep command[5].to_i
end