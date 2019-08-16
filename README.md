## Metallurgy Go
A Golang command-line application for bulk converting images between PNG and JPG. The app finds all files with the extension given to `--in` flag and converts them to format given to `--out` flag. __NOTE__: This app is single-threaded, so the larger the file set given, the longer it will take. It will be fast, but not as fast as possible. Concurrent execution will be implemented in a future release.

---

### Install Options
The following presents two methods of installation: the first to OS (Mac/Linux), the second using Go tools

---

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

---

#### Golang
Get the repo:
```bash
go get github.com/jwhittle933/metallurgygo
```

Install:
```bash
cd $GOPATH/src/github.com/metallurgygo && go install
```

(Optionally) change the binary name for easy use:
```bash
mv $GOPATH/bin/metallurgygo $GOPATH/bin/mgo
```

---

### Usage
```
mgo --dir <path/to/files> --in <input format> --out <output format> --save <path/to/save/location>
```
Each flag has its own defaults which can be seen with the `--help` flag:
```
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
  -v
        This option turns on verbose logging
```
