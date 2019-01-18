# comprtest

debug.go defines nebutil.Nebprint as a function that does a real print by calling fmt.Printf.

release.go defines nebutil.Nebprint as an empty function.
comprtest.go just calls nebutil.Nebprint

Depending on how we compile it (with tags debug or not), we will see the debug print.

Sample output:

hplinux@hplinux: 

hplinux@hplinux: go build -tags debug

hplinux@hplinux: ./comprtst /etc/passwd

pageCount: 1

8192            1          888         2169       2.44

hplinux@hplinux: go build

hplinux@hplinux: ./comprtst /etc/passwd

8192            1          888         2169       2.44

hplinux@hplinux:

