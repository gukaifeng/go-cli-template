GO = go
PROGRAM = myprogram

all: $(PROGRAM)

$(PROGRAM): main.go cmd/*.go
	$(GO) build -o $(PROGRAM) main.go

clean:
	rm -f $(PROGRAM)

PREFIX = /usr/local

install: $(PROGRAM)
	install -D $(PROGRAM) $(DESTDIR)$(PREFIX)/bin/$(PROGRAM)