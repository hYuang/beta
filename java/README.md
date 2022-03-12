# readme

    jvmti with go 
    jdk17 

## build command  

    go build -o javatool.dll -buildmode=c-shared javatool.go

## check the export function  

    dumpbin /exports java.dll
