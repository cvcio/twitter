test:
	godotenv -f ./.env go test -v
test-cover:
	godotenv -f ./.env go test -v -cover

test-channels:
	godotenv -f ./.env go test -timeout 120s -run Test_GetUserFollowers_Error

test-stream:
	godotenv -f ./.env go test -timeout 120s -run Test_GetFilterStream