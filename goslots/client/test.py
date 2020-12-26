import socket
import time
import json

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(('127.0.0.1', 1234))    

print(sock.recv(1024))
for i in range(100000):
	sock.send("500")
	print(sock.recv(1024))
	time.sleep(1)


sock.close()

print("done!")