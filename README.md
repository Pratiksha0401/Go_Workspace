#Go Commands
 
#command to init mod to root directory

go mod init <name of package>  eg. go mod init main

#to build 

go build <fileName>.go

#to run test 

go test   //only




#how to build file to install in linux

$go build -o /usr/local/go/bin/demo(binaryfile name created by build cmd) (your go path)

$demo(binaryfulename)



#how to build file to install in windows 

$go install filename.go

$filename

<br/>
#you cannot have main function in same directory more than one but if you have then test them by running separately

<br/>
<br/>
slice_intersection.go:6:2: no required module provides package golang.org/x/exp/constraints; to add it:
    run go cmd =>  $go get golang.org/x/exp/constraints