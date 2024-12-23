:: This file is the Windows version of the Maven wrapper script. 
:: It allows you to run Maven commands without requiring Maven to be installed on your machine.

@echo off
setlocal

set MAVEN_HOME=%~dp0\.mvn\wrapper\maven-wrapper.jar
set MAVEN_OPTS=%MAVEN_OPTS% -Dfile.encoding=UTF-8

if not exist "%MAVEN_HOME%" (
    echo "Maven wrapper jar not found. Please run mvnw to download it."
    exit /b 1
)

java -cp "%MAVEN_HOME%" org.apache.maven.wrapper.MavenWrapperMain %*