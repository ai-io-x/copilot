import socket
import threading
import json
import tkinter as tk
from tkinter import scrolledtext

class ChatClient:
    def __init__(self, host='127.0.0.1', port=8888):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect((host, port))
        
        self.window = tk.Tk()
        self.window.title("Chat Client")
        
        self.chat_area = scrolledtext.ScrolledText(self.window)
        self.chat_area.pack(padx=10, pady=10)
        
        self.msg_entry = tk.Entry(self.window)
        self.msg_entry.pack(padx=10, pady=5, fill=tk.X)
        self.msg_entry.bind("<Return>", self.send_message)
        
        self.username = input("Enter username: ")
        self.socket.send(json.dumps({"username": self.username}).encode())
        
        receive_thread = threading.Thread(target=self.receive_messages)
        receive_thread.daemon = True
        receive_thread.start()
        
    def receive_messages(self):
        while True:
            try:
                message = json.loads(self.socket.recv(1024).decode())
                if message["type"] == "message":
                    self.chat_area.insert(tk.END, 
                        f"{message['username']}: {message['content']}\n")
                elif message["type"] in ["join", "leave"]:
                    self.chat_area.insert(tk.END,
                        f"*** {message['username']} has {message['type']}ed ***\n")
                self.chat_area.see(tk.END)
            except:
                break
                
    def send_message(self, event=None):
        content = self.msg_entry.get()
        if content:
            self.socket.send(json.dumps({
                "type": "message",
                "content": content
            }).encode())
            self.msg_entry.delete(0, tk.END)
    
    def start(self):
        self.window.mainloop()

if __name__ == "__main__":
    client = ChatClient()
    client.start()