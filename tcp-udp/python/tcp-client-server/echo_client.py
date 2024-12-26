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