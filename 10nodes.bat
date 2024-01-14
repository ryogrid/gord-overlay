go build -o gordolctl.exe main.go
start /d . .\gordolctl.exe -l 127.0.0.1:10000
start /d . .\gordolctl.exe -l 127.0.0.1:10002 -n 127.0.0.1:10000
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10004 -n 127.0.0.1:10002
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10006 -n 127.0.0.1:10004
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10008 -n 127.0.0.1:10006
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10010 -n 127.0.0.1:10008
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10012 -n 127.0.0.1:10010
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10014 -n 127.0.0.1:10012
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10016 -n 127.0.0.1:10016
timeout 20
start /d . .\gordolctl.exe -l 127.0.0.1:10018 -n 127.0.0.1:10018
timeout 20
