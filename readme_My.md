## How to run this CLI app ? 
There is two ways to run this app:
1. Use `go run . [flags]`; 
2. Use `./uniq [flags]`; .


## How to run my tests ?
Use `go test -v ./tests/cli`

**Notes**: if you use stdin for input, use `ctrl+Z+Enter` to break the input. 
(You dont should do this if you out you data in stdin from file like this: `cat some_txt_file.txt|go run ./cmd/cli`).
