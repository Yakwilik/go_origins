.PHONY: build run

run:
	go run uniq/main.go -c -f 1 -s 1 $(filter-out $@,$(MAKECMDGOALS))


%:
	@: