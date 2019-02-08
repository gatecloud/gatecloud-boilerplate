c:
cd %userprofile%\go\src\gatecloud-boilerplate\api
if exist api.exe rm api.exe
go build
start api.exe

cd %userprofile%\go\src\gatecloud-boilerplate\proxy
if exist proxy.exe rm proxy.exe
go build
start proxy.exe


cd %userprofile%\go\src\gatecloud-boilerplate
echo "hostname is " 
hostname
pause

taskkill /IM api.exe
taskkill /IM proxy.exe