CC = /usr/bin/gcc
CFLAGS = -Wall -Wextra -Wpedantic

.PHONY: all clean

all:
	@echo "Usage: make <day number> (e.g., make 1)"

%:
	@if [ -d day$@ ]; then \
		cd day$@ && $(CC) $(CFLAGS) main.c ../utils.c && ./a.out; \
	else \
		echo "Directory day$@ does not exist."; \
	fi

clean:
	rm -f day*/a.out
