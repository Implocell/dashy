.PHONY: build
build:
	cd frontend && npm run build
	cd server && go build -o ../dashy

watch-server:
	cd server && nodemon --watch './main.go' --watch './frontend/**/*' --signal SIGTERM --exec 'go' run .

watch-frontend:
	cd frontend && npm run dev

.PHONY: dev
dev:
	${MAKE} -j4 watch-frontend watch-server

.PHONY: install-dev-requirements
install-dev-requirements:
	npm install -g nodemon