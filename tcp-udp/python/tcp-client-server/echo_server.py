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