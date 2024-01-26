package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/senzing-garage/g2-sdk-go-mock/g2config"
	"github.com/senzing-garage/g2-sdk-go-mock/g2configmgr"
	"github.com/senzing-garage/g2-sdk-go-mock/g2diagnostic"
	"github.com/senzing-garage/g2-sdk-go-mock/g2engine"
	"github.com/senzing-garage/g2-sdk-go-mock/g2product"
	"github.com/senzing-garage/g2-sdk-go/g2api"
	"github.com/senzing-garage/go-common/truthset"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdTemplate = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var Messages = map[int]string{
	1:    "%s",
	2:    "WithInfo: %s",
	2001: "Testing %s.",
	2002: "Physical cores: %d.",
	2003: "withInfo",
	2004: "License",
	2999: "Cannot retrieve last error message.",
}

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"
var logger logging.LoggingInterface

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getG2config(ctx context.Context) (g2api.G2config, error) {
	result := g2config.G2config{}
	moduleName := "Test module name"
	verboseLogging := int64(0) // 0 for no Senzing logging; 1 for logging
	iniParams := `{}`
	err := result.Init(ctx, moduleName, iniParams, verboseLogging)
	return &result, err
}

func getG2configmgr(ctx context.Context) (g2api.G2configmgr, error) {
	result := g2configmgr.G2configmgr{}
	moduleName := "Test module name"
	verboseLogging := int64(0)
	iniParams := `{}`
	err := result.Init(ctx, moduleName, iniParams, verboseLogging)
	return &result, err
}

func getG2diagnostic(ctx context.Context) (g2api.G2diagnostic, error) {
	result := g2diagnostic.G2diagnostic{}
	moduleName := "Test module name"
	verboseLogging := int64(0)
	iniParams := `{}`
	err := result.Init(ctx, moduleName, iniParams, verboseLogging)
	return &result, err
}

func getG2engine(ctx context.Context) (g2api.G2engine, error) {
	result := g2engine.G2engine{}
	moduleName := "Test module name"
	verboseLogging := int64(0)
	iniParams := `{}`
	err := result.Init(ctx, moduleName, iniParams, verboseLogging)
	return &result, err
}

func getG2product(ctx context.Context) (g2api.G2product, error) {
	result := g2product.G2product{}
	moduleName := "Test module name"
	verboseLogging := int64(0)
	iniParams := `{}`
	err := result.Init(ctx, moduleName, iniParams, verboseLogging)
	return &result, err
}

func getLogger(ctx context.Context) (logging.LoggingInterface, error) {

	logger, err := logging.NewSenzingLogger("my-unique-%04d", Messages)
	if err != nil {
		fmt.Println(err)
	}

	return logger, err
}

func demonstrateConfigFunctions(ctx context.Context, g2Config g2api.G2config, g2Configmgr g2api.G2configmgr) error {
	now := time.Now()

	// Using G2Config: Create a default configuration in memory.

	configHandle, err := g2Config.Create(ctx)
	if err != nil {
		return logger.NewError(5100, err)
	}

	// Using G2Config: Add data source to in-memory configuration.

	for _, testDataSource := range truthset.TruthsetDataSources {
		_, err := g2Config.AddDataSource(ctx, configHandle, testDataSource.Json)
		if err != nil {
			return logger.NewError(5101, err)
		}
	}

	// Using G2Config: Persist configuration to a string.

	configStr, err := g2Config.Save(ctx, configHandle)
	if err != nil {
		return logger.NewError(5102, err)
	}

	// Using G2Configmgr: Persist configuration string to database.

	configComments := fmt.Sprintf("Created by g2diagnostic_test at %s", now.UTC())
	configID, err := g2Configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		return logger.NewError(5103, err)
	}

	// Using G2Configmgr: Set new configuration as the default.

	err = g2Configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		return logger.NewError(5104, err)
	}

	return err
}

func demonstrateAddRecord(ctx context.Context, g2Engine g2api.G2engine) (string, error) {
	dataSourceCode := "TEST"
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000000))
	if err != nil {
		panic(err)
	}
	recordID := randomNumber.String()
	jsonData := fmt.Sprintf(
		"%s%s%s",
		`{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "`,
		recordID,
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`)
	loadID := dataSourceCode
	var flags int64 = 0

	// Using G2Engine: Add record and return "withInfo".

	return g2Engine.AddRecordWithInfo(ctx, dataSourceCode, recordID, jsonData, loadID, flags)
}

func demonstrateAdditionalFunctions(ctx context.Context, g2Diagnostic g2api.G2diagnostic, g2Engine g2api.G2engine, g2Product g2api.G2product) error {
	// Using G2Diagnostic: Check physical cores.

	actual, err := g2Diagnostic.GetPhysicalCores(ctx)
	if err != nil {
		failOnError(5300, err)
	}
	logger.Log(2002, actual)

	// Using G2Engine: Purge repository.

	err = g2Engine.PurgeRepository(ctx)
	if err != nil {
		failOnError(5301, err)
	}

	// Using G2Engine: Add records with information returned.

	withInfo, err := demonstrateAddRecord(ctx, g2Engine)
	if err != nil {
		failOnError(5302, err)
	}
	logger.Log(2003, withInfo)

	// Using G2Product: Show license metadata.

	license, err := g2Product.License(ctx)
	if err != nil {
		failOnError(5303, err)
	}
	logger.Log(2004, license)

	// Using G2Engine: Purge repository again.

	err = g2Engine.PurgeRepository(ctx)
	if err != nil {
		failOnError(5304, err)
	}

	return err
}

func destroyObjects(ctx context.Context, g2Config g2api.G2config, g2Configmgr g2api.G2configmgr, g2Diagnostic g2api.G2diagnostic, g2Engine g2api.G2engine, g2Product g2api.G2product) error {

	err := g2Config.Destroy(ctx)
	if err != nil {
		failOnError(5401, err)
	}

	err = g2Configmgr.Destroy(ctx)
	if err != nil {
		failOnError(5402, err)
	}

	err = g2Diagnostic.Destroy(ctx)
	if err != nil {
		failOnError(5403, err)
	}

	err = g2Engine.Destroy(ctx)
	if err != nil {
		failOnError(5404, err)
	}

	err = g2Product.Destroy(ctx)
	if err != nil {
		failOnError(5405, err)
	}

	return err
}

func failOnError(msgId int, err error) {
	logger.Log(msgId, err)
	panic(err.Error())
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	var err error = nil
	ctx := context.TODO()

	// Configure the "log" standard library.

	log.SetFlags(0)
	logger, err = getLogger(ctx)
	if err != nil {
		failOnError(5000, err)
	}

	// Test logger.

	programmMetadataMap := map[string]interface{}{
		"ProgramName":    programName,
		"BuildVersion":   buildVersion,
		"BuildIteration": buildIteration,
	}

	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
	logger.Log(2001, "Just a test of logging", programmMetadataMap)

	// Create observers.

	observer1 := &observer.ObserverNull{
		Id: "Observer 1",
	}

	// Get Senzing objects for installing a Senzing Engine configuration.

	g2Config, err := getG2config(ctx)
	if err != nil {
		failOnError(5001, err)
	}
	err = g2Config.RegisterObserver(ctx, observer1)
	if err != nil {
		panic(err)
	}

	g2Configmgr, err := getG2configmgr(ctx)
	if err != nil {
		failOnError(5005, err)
	}
	err = g2Configmgr.RegisterObserver(ctx, observer1)
	if err != nil {
		panic(err)
	}

	// Persist the Senzing configuration to the Senzing repository.

	err = demonstrateConfigFunctions(ctx, g2Config, g2Configmgr)
	if err != nil {
		failOnError(5008, err)
	}

	// Now that a Senzing configuration is installed, get the remainder of the Senzing objects.

	g2Diagnostic, err := getG2diagnostic(ctx)
	if err != nil {
		failOnError(5009, err)
	}
	err = g2Diagnostic.RegisterObserver(ctx, observer1)
	if err != nil {
		panic(err)
	}

	g2Engine, err := getG2engine(ctx)
	if err != nil {
		failOnError(5010, err)
	}
	err = g2Engine.RegisterObserver(ctx, observer1)
	if err != nil {
		panic(err)
	}

	g2Product, err := getG2product(ctx)
	if err != nil {
		failOnError(5011, err)
	}
	err = g2Product.RegisterObserver(ctx, observer1)
	if err != nil {
		panic(err)
	}

	// Demonstrate tests.

	err = demonstrateAdditionalFunctions(ctx, g2Diagnostic, g2Engine, g2Product)
	if err != nil {
		failOnError(5015, err)
	}

	// Destroy Senzing objects.

	err = destroyObjects(ctx, g2Config, g2Configmgr, g2Diagnostic, g2Engine, g2Product)
	if err != nil {
		failOnError(5016, err)
	}

	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
}
