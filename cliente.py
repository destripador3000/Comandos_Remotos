import socket

PORT = 8080
FORMAT = 'utf-8'
DISCONNECT_MESSAGE = "!DESCONECTAR"
SERVER = "127.0.0.1"
ADDR = (SERVER, PORT)

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(ADDR)

def send(msg):
    message = (msg + '\n').encode(FORMAT)
    client.sendall(message)
    
    response = client.recv(4096).decode(FORMAT)
    print(response)

try:
    while True:
        cmd = input(">>> ")
        if cmd.lower() in ["exit", "quit"]:
            send(DISCONNECT_MESSAGE)
            break
        send(cmd)
except KeyboardInterrupt:
    send(DISCONNECT_MESSAGE)

client.close()
