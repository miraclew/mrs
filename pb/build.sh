protoc --go_out=. *.proto
protoc --cpp_out=cpp *.proto

go install