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