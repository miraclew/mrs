package pb

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/json"
	"log"
	"testing"
)

func TestA(t *testing.T) {
	test := &Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newTest := &Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	// etc.

	log.Printf("%s\n", newTest.String())

	bytes, _ := json.Marshal(newTest)
	log.Println(string(bytes))
}
