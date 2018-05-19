# comprtest

Sample output:

hplinux@hplinux:~/go/src/comprtst$ 
hplinux@hplinux:~/go/src/comprtst$ go build -tags debug
hplinux@hplinux:~/go/src/comprtst$ ./comprtst /etc/passwd
pageCount: 1
8192            1          888         2169       2.44
hplinux@hplinux:~/go/src/comprtst$ go build
hplinux@hplinux:~/go/src/comprtst$ ./comprtst /etc/passwd
8192            1          888         2169       2.44
hplinux@hplinux:~/go/src/comprtst$ 

