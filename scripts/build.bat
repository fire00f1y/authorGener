REM This is designed to be able to be run from within the system GOPATH. If you clone this outside that gopath, you'll have to set a custom gopath (It's commented out now, you just need to modify)

REM SET GOPATH=%GOPATH%;C:\Your\Fav\Location
SET GOARCH=amd64

call pushd %GOPATH%
call go get -d .\...
call popd

SET GOOS=windows
call go install github.com/fire00f1y/authorGener

SET GOOS=linux
call go install github.com/fire00f1y/authorGener

SET GOOS=darwin
call go install github.com/fire00f1y/authorGener