NAME=galaxy

.PHONY: build
build: .env
	@go build -v -tags x11 -buildvcs=false -o ./bin/${NAME} ./cmd/${NAME}

.PHONY: watch
watch:
	@inotifywait -e close_write,moved_to,create -rmq pkg cmd @bin | \
	while read -r directory events filename; do \
		echo $$'\n'; echo "> $$directory$$filename changed. Recompiling...";\
		make -s build; \
	done

.PHONY: run
run: .env build
	./bin/${NAME}

.env: .env.dist
	cp .env.dist .env
