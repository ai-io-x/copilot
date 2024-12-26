# Plan for Terminal-Based Chat Client

1. Remove GUI components
2. Create command-line interface
3. Use threading for message receiving
4. Handle user input in main thread
5. Add quit command support

```python


import socket
import threading
import json
import sys

class ChatClient:
    def __init__(self, host='127.0.0.1', port=8888):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect((host, port))
        self.running = True
        
    def receive_messages(self):
        while self.running:
            try:
                message = json.loads(self.socket.recv(1024).decode())
                if message["type"] == "message":
                    print(f"\r{message['username']}: {message['content']}")
                elif message["type"] in ["join", "leave"]:
                    print(f"\r*** {message['username']} has {message['type']}ed ***")
                print("\rYou: ", end='', flush=True)
            except:
                self.running = False
                break

    def start(self):
        username = input("Enter username: ")
        self.socket.send(json.dumps({"username": username}).encode())
        
        # Start receive thread
        receive_thread = threading.Thread(target=self.receive_messages)
        receive_thread.daemon = True
        receive_thread.start()
        
        # Main loop for sending messages
        print("Type 'quit' to exit")
        while self.running:
            try:
                message = input("You: ")
                if message.lower() == 'quit':
                    self.running = False
                    break
                self.socket.send(json.dumps({
                    "type": "message",
                    "content": message
                }).encode())
            except KeyboardInterrupt:
                self.running = False
                break
        
        self.socket.close()
        print("\nDisconnected from server")

if __name__ == "__main__":
    client = ChatClient()
    client.start()
```

Usage:
1. Start server first
2. Run client: `python chat_client.py`
3. Enter username
4. Type messages and press Enter to send
5. Type 'quit' to exit