# learning-go-with-ido
All the go projects in one place            

Learing with the [course](https://www.udemy.com/course/go-the-complete-developers-guide/)               
This repo defines the learning path of go               

# Helpful commands
compile:
```bash
go build FILENAME
```
compile&run:
```bash
go run ALL_FILES_NAMES
```
```bash
# run all go files in the directory
go run $(find . -type f -name "*.go" | tr "\n" " ")

```
What is the difference between those:          
* `go build` -> compile          
* `go run` -> compile & execute             
* `go fmt` -> format all code in directory              
* `go install` -> compile and install a package          
* `go get` -> download raw source code from projects        
* `go test` -> run any tests associated with current project        

# packages
Where is the packages? [here](https://pkg.go.dev/std)         
`main` is a key word for packages that will complie and build binary files, there are helper packages athat won't build an artifact.

# Projects Overview
* Hello World! [hello-world-project](/first-project/)              
* Go Basics! [fun-with-types](/second-project-cards/)              

