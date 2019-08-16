## Metallurgy Go
A Golang command-line application for bulk converting images

### Install Options
The following presents two methods of installation: the first to the OS (Mac/Linux), the second using Go topls

#### OS
Clone the repo:
```bash
git clone https://github.com/jwhittle933/metallurgygo
```

Build the application:
```bash
go build -o ./mgo
```

Move the binary:
```bash
mv mgo /usr/local/bin/
```

(Optionally) remove the source code:
```bash
rm -rf ./metallurgygo
```

#### Golang
Get the repo:
```bash
go get github.com/jwhittle933/metallurgygo
```

Install:
```bash
cd $GOPATH/src/github.com/metallurgygo && go install $GOPATH/src/github.com/metallurgygo
```

(Optionally) change the binary name for easy use:
```bash
mv $GOPATH/bin/metallurgygo $GOPATH/bin/mgo
```

### Usage
```bash
mgo --dir <path/to/files> --in <input format> --out <output format> --save <path/to/save/location>
```
Each flag has its own defaults:
```bash
$ mgo --help
[Info] 2019/08/15 21:49:04 logger.go:47: USAGE
  -dir string
    	The directory location of your files (default ".")
  -help:
    	Print flags
  -in string
    	The filetype to start with (default ".png")
  -out string
    	The filetype to convert to (default ".jpg")
  -save string
    	The save location for files (default ".")
```
