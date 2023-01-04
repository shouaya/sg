require './lib.rb'
require 'auto_click'
require 'spreadsheet'
include AutoClickMethods

config = readFile('steps.xls', 'Sheet1')
# open window
start(config['start'][0], config['start'][1])
# config['commands'].each do |command|
#     select_troop(command['id'])
#     case command['type']
#     when '占领'
#       take_place(1273, 316)
#     when '行军'
#       mouse_move(1273, 366)
#     when '回城'
#       back_city(1273, 416)
#     when '退阵'
#       del_troops(1273, 436)
#     when '上阵'
#       add_troops(1273, 486)
#     when '上阵'
#       add_troops(1273, 486)
#     when '上阵'
#       add_troops(1273, 486)
#     else
#       print('not support')
#     end
# end