import socket
import threading
import json

class ChatServer:
    def __init__(self, host='127.0.0.1', port=8888):
        self.host = host
        self.port = port
        self.server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.clients = {}  # {client_socket: username}
        
    def broadcast(self, message, sender=None):
        for client in self.clients:
            if client != sender:
                try:
                    client.send(json.dumps(message).encode())
                except:
                    self.remove_client(client)

    def handle_client(self, client_socket, addr):
        try:
            username = json.loads(client_socket.recv(1024).decode())['username']
            self.clients[client_socket] = username
            self.broadcast({"type": "join", "username": username})
            
            while True:
                message = json.loads(client_socket.recv(1024).decode())
                if message["type"] == "message":
                    self.broadcast({
                        "type": "message",
                        "username": username,
                        "content": message["content"]
                    }, client_socket)
        except:
            self.remove_client(client_socket)

    def remove_client(self, client_socket):
        if client_socket in self.clients:
            username = self.clients[client_socket]
            del self.clients[client_socket]
            client_socket.close()
            self.broadcast({"type": "leave", "username": username})

    def start(self):
        self.server_socket.bind((self.host, self.port))
        self.server_socket.listen(5)
        print(f"Chat server running on {self.host}:{self.port}")
        
        try:
            while True:
                client_sock, addr = self.server_socket.accept()
                thread = threading.Thread(target=self.handle_client, args=(client_sock, addr))
                thread.start()
        finally:
            self.server_socket.close()

if __name__ == "__main__":
    server = ChatServer()
    server.start()