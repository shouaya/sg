import cv2
import numpy as np

image= cv2.imread('desktop.png')
gray= cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

template= cv2.imread('game_icon.png',0)


result= cv2.matchTemplate(gray, template, cv2.TM_CCOEFF)
min_val, max_val, min_loc, max_loc= cv2.minMaxLoc(result)

height, width= template.shape[:2]

top_left= max_loc
bottom_right= (top_left[0] + width, top_left[1] + height)

print(top_left, bottom_right)

cv2.rectangle(image, top_left, bottom_right, (0,0,255),5)

cv2.imwrite('result.png', image)

