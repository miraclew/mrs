protoc --go_out=. *.proto
protoc --cpp_out=cpp *.proto

go install

GOOGLE_NAMESPACE=google_public
(
    sed -i '' -e "s/namespace\ google /namespace\ ${GOOGLE_NAMESPACE} /g" $(find . -name \*.h -type f)
    sed -i '' -e "s/namespace\ google /namespace\ ${GOOGLE_NAMESPACE} /g" $(find . -name \*.cc -type f)
    sed -i '' -e "s/namespace\ google /namespace\ ${GOOGLE_NAMESPACE} /g" $(find . -name \*.proto -type f)
    sed -i '' -e "s/google::protobuf/${GOOGLE_NAMESPACE}::protobuf/g" $(find . -name \*.h -type f)
    sed -i '' -e "s/google::protobuf/${GOOGLE_NAMESPACE}::protobuf/g" $(find . -name \*.cc -type f)
    sed -i '' -e "s/google::protobuf/${GOOGLE_NAMESPACE}::protobuf/g" $(find . -name \*.proto -type f)
)
cp cpp/* ~/git/cetris/Cetris/pb
