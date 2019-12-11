run:
	go run main.go

build:
	go build -gcflags "-m -m" -o binary_lab_1

bench:
	go test . --bench=. --memprofile mem.out

memp: bench
	go tool pprof mem.out