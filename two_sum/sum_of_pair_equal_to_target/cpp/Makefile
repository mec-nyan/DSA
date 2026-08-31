IN = two_sum.cpp
OUT = two_sum.o

build: $(IN)
	@g++-16 --std=c++23 -Wall -o $(OUT) $(IN)

clean:
	@if [ -e $(OUT) ]; then rm $(OUT); else echo "Nothing to clean."; fi

run: build
	@./$(OUT)
