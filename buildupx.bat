@echo off
set proj_name=beego-learn
set ts=%date:~2,2%%date:~5,2%%date:~8,2%.%time:~0,2%%time:~3,2%%time:~6,2%
set ts=%ts: =0%
set exe_name=%proj_name%-%ts%.exe
set upx_name=%proj_name%-%ts%-upx.exe
go build -o %exe_name% && upx -9 -o %upx_name% %exe_name% && rm %exe_name%
