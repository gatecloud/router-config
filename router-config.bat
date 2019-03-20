
c:
cd C:\%userprofile%\go\src\router-config
rm router-config.exe
go build
start router-config.exe
start chrome http://localhost:7000/home
echo "hostname is " 
hostname
pause
taskkill /IM router-config.exe