
BIN := wac

all: $(BIN)

$(BIN):
	go build

clean:
	rm -f $(BIN)
