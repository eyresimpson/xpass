@REM 生产环境用这个，不会出现黑色窗口
go build -ldflags="-H windowsgui" -o ./dist/xpass.exe ./main.go
@REM 代码调试时用这个，可以看到程序输出的错误
@REM go build -o ./dist/xpass.exe ./main.go
