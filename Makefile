all:
	@echo "make"
run:
	@make push
push:
	-git add . && git commit -m 'auto commit' && git push origin master
test:
	@go test ./