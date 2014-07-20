protoc --go_out=. *.proto
protoc --cpp_out=cpp *.proto

go install

cp cpp/* /Users/aaaa/git/cetris/Cetris/pb/
