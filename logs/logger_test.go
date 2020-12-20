package logs

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/johnfg10/authit/logs/mock_logs"
)

type TestStruct struct {
	ConfidentialString string `confidential:""`
	NormalString       string
}

//mockgen -destination="./mock_logs/mock_logger.go" . ObjectPrinter
func TestObjectConfidentialityPrinter(t *testing.T) {
	controller := gomock.NewController(t)
	mock := mock_logs.NewMockObjectPrinter(controller)

	sucessStr := fmt.Sprintf("ConfidentialString: %s, NormalString: %s", ExcludedConfidential, "This should be printed")

	mockTestStruct := TestStruct{
		ConfidentialString: "This shouldnt be printer",
		NormalString:       "This should be printed",
	}

	mock.EXPECT().PrintField(gomock.Eq("ConfidentialString"), gomock.Eq(ExcludedConfidential))
	mock.EXPECT().PrintField(gomock.Eq("NormalString"), gomock.Eq("This should be printed"))
	mock.EXPECT().Build().Return(sucessStr)

	logger := Logger{
		ObjectPrint: mock,
	}

	if sucessStr != logger.Print(mockTestStruct) {
		t.Fail()
	}
}

func TestObjectConfidentialityPrinterNonStruct(t *testing.T) {
	controller := gomock.NewController(t)
	mock := mock_logs.NewMockObjectPrinter(controller)

	//sucessStr := fmt.Sprintf("ConfidentialString: %s, NormalString: %s", ExcludedConfidential, "This should be printed")

	// mockTestStruct := TestStruct{
	// 	ConfidentialString: "This shouldnt be printer",
	// 	NormalString:       "This should be printed",
	// }

	// mock.EXPECT().PrintField(gomock.Eq("ConfidentialString"), gomock.Eq(ExcludedConfidential))
	// mock.EXPECT().PrintField(gomock.Eq("NormalString"), gomock.Eq("This should be printed"))
	// mock.EXPECT().Build().Return(sucessStr)

	logger := Logger{
		ObjectPrint: mock,
	}

	if "Test" != logger.Print("Test") {
		t.Fail()
	}
}

func TestStringObjectPrinter(t *testing.T) {

	sop := NewStringObjectPrinter()

	sop.PrintField("Test Name", "Test Value")

	value := sop.Build()
	t.Log("Test", value)
	if value != "Test Name: Test Value," {
		t.Fail()
	}

	if sop.builder.String() != "" {
		t.Fail()
	}
}
