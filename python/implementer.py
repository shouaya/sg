import asyncio
import socketio
import pyautogui

sio = socketio.AsyncClient()

@sio.event
async def connect():
    print('connection established')
    await sio.emit('chat message', {'response': 'hello'})

@sio.on('chat message')
async def on_message(data):
    print('message received with ', data)
    if "e" in data:
        if data['e'] == 'c':
            print("click")
            pyautogui.click(data['x'], data['y'])
        else:
            print("drag")
            pyautogui.moveTo(data['fx'], data['fy'], duration = 0)
            pyautogui.dragTo(data['tx'], data['ty'], duration = 1)

@sio.event
async def disconnect():
    print('disconnected from server')

async def main():
    await sio.connect('http://localhost:3000')
    await sio.wait()

if __name__ == '__main__':
    asyncio.run(main())