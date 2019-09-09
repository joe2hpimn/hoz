### Connection Protocol
Currently implemented HTTP protocol, local and server are directly connected

### Data Communication Protocol
The package [head, body], head is 4 bytes containing the current encrypted packet length, and the body is the encrypted bytes. After the read is completed, the decryption is performed.

### About encryption
The default is OORR, which is implemented by myself. It is mainly used to perform or perform operations on bytes. Modify all bytes to achieve the purpose of encryption. It is fast and consumes less resources.

### How to use it
./client_side -addr ":1080" -remote "127.0.0.1:10800" -password "oor://!@adDxS$&(dl/*?QKc$mJ?PdTkajGzSNMILH{t4_hvFR>" <br>
./server_side -addr ":10800" -password "oor://!@adDxS$&(dl/*?QKc$mJ?PdTkajGzSNMILH{t4_hvFR>"
### TODO
SOCKS5 protocol
