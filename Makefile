.PHONY: build run

run:
	go run uniq/main.go -u input.txt $(filter-out $@,$(MAKECMDGOALS))


%:
	@: