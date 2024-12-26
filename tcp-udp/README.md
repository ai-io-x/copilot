# TCP Server/Client:

- prompt: How to write a simple backend service over TCP/udp protocol

- GitHub Copilot: Let's build a simple TCP echo server that accepts connections and echoes back messages.

Steps:
1. Create TCP server that listens on a port
2. Accept client connections
3. Handle each client in a separate thread
4. Echo received messages back to client
5. Add proper error handling and cleanup

Here's the implementation in Python:

```python
import socket
import threading

class TCPServer:
    def __init__(self, host='127.0.0.1', port=8888):
        self.host = host
        self.port = port
        self.server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

    def handle_client(self, client_socket, addr):
        print(f"New connection from {addr}")
        try:
            while True:
                data = client_socket.recv(1024)
                if not data:
                    break
                print(f"Received from {addr}: {data.decode()}")
                client_socket.send(data)  # Echo back
        except Exception as e:
            print(f"Error handling client {addr}: {e}")
        finally:
            client_socket.close()
            print(f"Connection closed from {addr}")

    def start(self):
        self.server_socket.bind((self.host, self.port))
        self.server_socket.listen(5)
        print(f"Server listening on {self.host}:{self.port}")

        try:
            while True:
                client_sock, addr = self.server_socket.accept()
                client_thread = threading.Thread(
                    target=self.handle_client,
                    args=(client_sock, addr)
                )
                client_thread.start()
        except KeyboardInterrupt:
            print("Server shutting down...")
        finally:
            self.server_socket.close()

if __name__ == "__main__":
    server = TCPServer()
    server.start()
```

And a simple test client:

```python
import socket

def test_client():
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(('127.0.0.1', 8888))
    
    try:
        message = "Hello, Server!"
        client.send(message.encode())
        response = client.recv(1024)
        print(f"Server response: {response.decode()}")
    finally:
        client.close()

if __name__ == "__main__":
    test_client()
```

To run:
1. Start server: `python echo_server.py`
2. In another terminal: `python echo_client.py`

# UDP Server/Client:

- prompt: How to implement UDP instead?

- GitHub Copilot: Building UDP Server/Client

## Plan:
1. Create UDP server class
2. Implement datagram receiving loop
3. Add message handling logic  
4. Create UDP client class
5. Add error handling and cleanup
6. Test communication

## Implementation:

```python
import socket

class UDPServer:
    def __init__(self, host='127.0.0.1', port=8888):
        self.host = host
        self.port = port
        self.server_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        
    def start(self):
        self.server_socket.bind((self.host, self.port))
        print(f"UDP Server listening on {self.host}:{self.port}")
        
        try:
            while True:
                data, addr = self.server_socket.recvfrom(1024)
                print(f"Received from {addr}: {data.decode()}")
                # Echo back
                self.server_socket.sendto(data, addr)
        except KeyboardInterrupt:
            print("Server shutting down...")
        finally:
            self.server_socket.close()

if __name__ == "__main__":
    server = UDPServer()
    server.start()
```

```python
import socket

class UDPClient:
    def __init__(self, host='127.0.0.1', port=8888):
        self.server_addr = (host, port)
        self.client_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        
    def send_message(self, message):
        try:
            # Send data
            self.client_socket.sendto(message.encode(), self.server_addr)
            
            # Receive response
            data, _ = self.client_socket.recvfrom(1024)
            print(f"Server response: {data.decode()}")
        finally:
            self.client_socket.close()

if __name__ == "__main__":
    client = UDPClient()
    client.send_message("Hello UDP Server!")
```

## Usage:
1. Start server: `python udp_server.py`
2. In new terminal: `python udp_client.py`

# HTTP Server:

- prompt: How to implement HTTP instead?

- Building a Simple HTTP Server - Step by Step Plan

1. **Basic Structure**
   - Create HTTPServer class to handle incoming requests
   - Implement request parser for HTTP headers and body
   - Setup socket listener on port 8080

2. **Request Handling**
   - Parse incoming HTTP requests
   - Extract method, path, headers
   - Handle different HTTP methods (GET, POST)

3. **Response Generation**
   - Create HTTP response builder
   - Add status codes and headers
   - Support basic content types

4. **Implementation**

```python
import socket
import threading

class HTTPServer:
    def __init__(self, host='127.0.0.1', port=8080):
        self.host = host
        self.port = port
        self.server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

    def parse_request(self, request):
        request_lines = request.split('\r\n')
        request_line = request_lines[0]
        method, path, _ = request_line.split()
        headers = {}
        
        # Parse headers
        for line in request_lines[1:]:
            if line == '':
                break
            key, value = line.split(': ')
            headers[key] = value
            
        return method, path, headers

    def create_response(self, status_code, content):
        response = f'HTTP/1.1 {status_code}\r\n'
        response += 'Content-Type: text/plain\r\n'
        response += f'Content-Length: {len(content)}\r\n'
        response += '\r\n'
        response += content
        return response

    def handle_client(self, client_socket, addr):
        try:
            request_data = client_socket.recv(1024).decode()
            if not request_data:
                return

            method, path, headers = self.parse_request(request_data)
            print(f"Received {method} request for {path}")

            # Handle request
            if method == 'GET':
                response = self.create_response('200 OK', f'Requested path: {path}')
            else:
                response = self.create_response('405 Method Not Allowed', 'Only GET is supported')

            client_socket.send(response.encode())
        except Exception as e:
            print(f"Error handling client {addr}: {e}")
        finally:
            client_socket.close()

    def start(self):
        self.server_socket.bind((self.host, self.port))
        self.server_socket.listen(5)
        print(f"HTTP Server running on http://{self.host}:{self.port}")

        try:
            while True:
                client_sock, addr = self.server_socket.accept()
                client_thread = threading.Thread(
                    target=self.handle_client,
                    args=(client_sock, addr)
                )
                client_thread.start()
        except KeyboardInterrupt:
            print("\nServer shutting down...")
        finally:
            self.server_socket.close()

if __name__ == "__main__":
    server = HTTPServer()
    server.start()
```

5. **Testing**
- Run server: `python simple_http_server.py`
- Test using curl: `curl http://localhost:8080/test`
- Or open in browser: `http://localhost:8080/test`

This implements a basic HTTP server handling GET requests with proper HTTP headers and response formatting. It supports threading for multiple connections and includes basic error handling.

# TCP vs UDP Server Comparison

## Key Differences

1. **Connection Handling**
   - TCP: Connection-oriented, requires handshake
   - UDP: Connectionless, no handshake needed

2. **Data Delivery**
   - TCP: Guaranteed delivery, ordered packets
   - UDP: No guarantee, packets may arrive out of order

3. **Speed**
   - TCP: Slower due to overhead
   - UDP: Faster, minimal overhead

4. **Use Cases**
   - TCP: Web browsers, email, file transfer
   - UDP: Video streaming, gaming, VoIP

## Code Comparison

```python
# TCP Server - Reliable, ordered
import socket

class TCPServer:
    def __init__(self, port=8888):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.bind(('localhost', port))
        self.socket.listen(1)
        
    def start(self):
        while True:
            # Wait for connection
            client, address = self.socket.accept()
            # Receive data
            data = client.recv(1024)
            # Send response
            client.send(data)
            client.close()
```

```python
# UDP Server - Fast, no guarantees
import socket

class UDPServer:
    def __init__(self, port=8888):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        self.socket.bind(('localhost', port))
        
    def start(self):
        while True:
            # Receive data directly (no connection needed)
            data, addr = self.socket.recvfrom(1024)
            # Send response back to address
            self.socket.sendto(data, addr)
```

## When to Use Which?

- **Use TCP when:**
  - Data must arrive completely and in order
  - Reliability is critical
  - Network errors must be handled automatically

- **Use UDP when:**
  - Speed is priority
  - Some data loss is acceptable
  - Real-time applications
  - Broadcasting/Multicasting needed


# Building a Chat Application - Plan

1. **Analysis for Chat Requirements**
   - Real-time messaging needed
   - Multiple clients support
   - Message ordering important
   - Reliable delivery required

2. **Decision: Use TCP**
   - Guarantees message delivery
   - Maintains message order
   - Handles connection state
   - Better for chat integrity

## Implementation

```python
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
```

```python
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
```

## Usage
1. Start server: `python chat_server.py`
2. Start multiple clients: `python chat_client.py`
3. Enter username for each client
4. Start chatting!

TCP is better for chat because:
- Guaranteed message delivery
- Messages arrive in order
- Connection state tracking
- Better for text-based communication
- Built-in error handling