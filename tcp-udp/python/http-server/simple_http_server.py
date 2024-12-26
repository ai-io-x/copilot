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