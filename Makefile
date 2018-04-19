.PHONY : build
build :
	go build -o ./bin/pgn2text ./cmd/pgn2text/*.go

.PHONY : run
run : 
	./bin/pgn2text

.PHONY : test
test :
	cd cmd/pgn2text && go test
	cd pkg/pgn2text && go test

