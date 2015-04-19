## Simple GO application for processing text from tcp connect and providing statistical data through http. ##

Note. Currently application processes only latin symbols and digital.

### Installing and Quick Start ###

// clone
```bash
git clone https://github.com/iostrovok/test_go_server.git
cd test_go_server
```

// load go packages
```bash
make install
```

// make exe files
```bash
make build
```

// start
```bash
./run
```

// or read help
```bash
./run -h
```

### Simple flow ###

Next steps:

First console:
```bash
./run
```

Second console:
```bash
./add_word_file -f=./my_large_text_file.txt 
```

Browser:

http://localhost:8080/stats

http://localhost:8080/stats?n=10

http://localhost:8080/stats?n=10&d=true


