echo "Testing..."
go test -v -coverprofile cover.out ./ 
echo "Testing success!"

echo "Building..."
go build .

echo "Move build to bin/golang_test..."

mv golang_test bin/golang_test
echo "Building success!"