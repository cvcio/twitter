test:
	godotenv -f ./.env go test -v
test-cover:
	godotenv -f ./.env go test -v -cover