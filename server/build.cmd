@echo off
setlocal enabledelayedexpansion

:: 配置参数
set OUTPUT_DIR=dist
set APP_NAME=myapp
set TARGET_OS=linux
set TARGET_ARCH=amd64

:: 设置交叉编译环境变量
set CGO_ENABLED=0
set GOOS=!TARGET_OS!
set GOARCH=!TARGET_ARCH!

:: 调试输出
echo 当前 GOOS=!GOOS!
echo 当前 GOARCH=!GOARCH!

:: 创建输出目录并编译
mkdir "%OUTPUT_DIR%" >nul 2>&1
@REM go build -trimpath -ldflags="-s -w" -o "%OUTPUT_DIR%\!APP_NAME%-!TARGET_OS%-!TARGET_ARCH!"
go build -o "%OUTPUT_DIR%\%APP_NAME%-%TARGET_OS%-%TARGET_ARCH%"

:: 结果提示
if %errorlevel% equ 0 (
    echo 编译成功！文件保存在: %OUTPUT_DIR%\
) else (
    echo 编译失败，请检查 Go 版本和环境变量
)
endlocal