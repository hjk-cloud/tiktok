@echo off

@REM 设置要执行的任务cd ../cmd/tiktok-schedule/
@REM set task="..\cmd\douyin\douyin.exe"
@REM set task="D:\GitProjects\tiktok\cmd\douyin\douyin.exe"
set task=" ..\cmd\tiktok-schedule\tiktok-schedule.exe"

rem 设置任务要在每天的几时几分运行
set hour=18
set minute=40

rem 将任务设置为每天在指定时间运行
@REM at %hour%:%minute% %task%
schtasks /create /tn "tiktok-schedule" /tr %task% /sc DAILY /st %hour%:%minute%
@REM SCHTASKS /Delete /TN tiktok-schedule

rem 显示设置成功的提示
echo.
echo Task set successfully!
echo.
:: return to the directory of this script
cd /d %~dp0

pause
exit