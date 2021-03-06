package gogen

import (
	"testing"
	"os"
	def "github.com/bulletRush/qcloud-sdk/generator"
	"fmt"
	engine "github.com/bulletRush/qcloud-sdk/qcloud"
)

var (
	diskTypeParamDef = def.ParamDefinition{
		Name: "diskType",
		Type: "string",
		Rule: "root|data",
		Optional: true,
		Describe: "disk usage type",
	}
	payModeParamDef = def.ParamDefinition{
		Name: "payMode",
		Type: "string",
		Rule: "prePay|postPay",
		Optional: true,
		Describe: "disk pay mode",
	}
	storageTypeParamDef = def.ParamDefinition{
		Name: "storageType",
		Type: "string",
		Rule: "cloudBasic|cloudSSD",
		Optional: false,
		Describe: "storage type",
	}
	interfaceDef = def.InterfaceDefinition{
		Name: "DescribeCbsStorages",
		Brief: "list cbs storages",
		Describe: "see xx for more",
		InputParamList: []def.ParamDefinition{
			diskTypeParamDef, payModeParamDef, storageTypeParamDef,
		},
	}
)

func TestNewGoGenerator(t *testing.T) {

}

func newGoGenerator() *GoGenerator {
	return NewGoGenerator(os.Stdout, "svc", "CbsService")
}

func TestGoGenerator_GenRequestDefinition(t *testing.T) {
	gg := newGoGenerator()
	gg.GenRequestDefinition(interfaceDef)
	gg.Output()
}

func xTestGoGenerator_GenInputParamCheck(t *testing.T) {
	gg := newGoGenerator()
	gg.GenInputParamCheck(diskTypeParamDef)
}

func xTestGoGenerator_GenCheckCall(t *testing.T) {
	gg := newGoGenerator()
	fmt.Println(gg.GenCheckCall(payModeParamDef))
}

func TestGoGenerator_GenFunc(t *testing.T) {
	gg := newGoGenerator()
	gg.GenFunc( interfaceDef)
	gg.Output()
}

func TestGoGenerator_GenResponseDefinition(t *testing.T) {
	gg := newGoGenerator()
	gg.GenResponseDefinition(interfaceDef)
	gg.Output()
}

func newQloudEngine() engine.QcloudEngine {
	return engine.NewQcloudEngine()
}

func TestNewCbsService(t *testing.T) {
	engine := newQloudEngine()
	svc := NewCbsService(engine)
	svc.DescribeCbsStorages(DescribeCbsS)
}
