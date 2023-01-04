@echo off
setlocal enabledelayedexpansion
cd %~dp0
ruby update.rb && bundle install && ruby main.rb
pause