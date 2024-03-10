.PHONY: build run

run:
	go run uniq/main.go input.txt output.txt $(filter-out $@,$(MAKECMDGOALS))


%:
	@: