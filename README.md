# hoz
<b>通信协议</b>
包[head,body],head为4字节包含当前加密包长度，body为加密的byte字节，读取完成后进行解密

<b>连接协议</b>
目前实现了HTTP协议，local和server直接连接

<b>加密</b>
默认为OORR，主要是对byte进行或与运算，打乱所有的byte达到加密的目的，速度很快，耗费资源少

<b>Todo</b>
socks5 协议
