import os
import time
from bottle import route, run, request, response, template
import cv2
import numpy as np
import json
import math

@route('/start' , method='POST')
def start():
    image = request.files.get('file')
    save_path = os.path.join('/tmp/', time.strftime("%Y%m%d-%H%M%S-") + image.filename)
    image.save(save_path)
    point = findImage(save_path, 'start_icon.png', True)
    os.remove(save_path)
    response.content_type = 'application/json'
    return json.dumps({ "x": point[0], "y": point[1] })

@route('/adjust' , method='POST')
def adjust():
    image = request.files.get('file')
    save_path = os.path.join('/tmp/', time.strftime("%Y%m%d-%H%M%S-") + image.filename)
    image.save(save_path)
    point = findImage(save_path, 'top_icon.png', False)
    os.remove(save_path)
    response.content_type = 'application/json'
    return json.dumps({ "x": point[0], "y": point[1] })

@route('/back' , method='POST')
def back():
    image = request.files.get('file')
    save_path = os.path.join('/tmp/', time.strftime("%Y%m%d-%H%M%S-") + image.filename)
    image.save(save_path)
    point = findImage(save_path, 'back_icon.png', False)
    os.remove(save_path)
    response.content_type = 'application/json'
    return json.dumps({ "x": point[0], "y": point[1] })

def findImage(bigImage, smallImage, padding):
    image= cv2.imread(bigImage)
    gray= cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    template= cv2.imread(smallImage, 0)
    result= cv2.matchTemplate(gray, template, cv2.TM_CCOEFF)
    min_val, max_val, min_loc, max_loc= cv2.minMaxLoc(result)
    height, width= template.shape[:2]
    top_left = max_loc
    if padding:
        return (top_left[0] + math.ceil(width/2), top_left[1] + math.ceil(height/2))
    return top_left

run(host='0.0.0.0', port=8080)
