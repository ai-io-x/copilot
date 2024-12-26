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