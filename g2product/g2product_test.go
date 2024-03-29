package g2product

import (
	"context"
	"fmt"
	"os"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/g2-sdk-go/g2api"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
	printResults      = false
)

var (
	g2productSingleton *G2product
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) *G2product {
	return getG2Product(ctx)
}

func getG2Product(ctx context.Context) *G2product {
	if g2productSingleton == nil {
		g2productSingleton = &G2product{
			LicenseResult: `{"customer":"Senzing Public Test License","contract":"EVALUATION - support@senzing.com","issueDate":"2022-11-29","licenseType":"EVAL (Solely for non-productive use)","licenseLevel":"STANDARD","billing":"MONTHLY","expireDate":"2023-11-29","recordLimit":50000}`,
			VersionResult: `{"PRODUCT_NAME":"Senzing API","VERSION":"3.5.0","BUILD_VERSION":"3.5.0.23041","BUILD_DATE":"2023-02-09","BUILD_NUMBER":"2023_02_09__23_01","COMPATIBILITY_VERSION":{"CONFIG_VERSION":"10"},"SCHEMA_VERSION":{"ENGINE_SCHEMA_VERSION":"3.5","MINIMUM_REQUIRED_SCHEMA_VERSION":"3.0","MAXIMUM_REQUIRED_SCHEMA_VERSION":"3.99"}}`,
		}
	}
	return g2productSingleton
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if printResults {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, g2product g2api.G2product, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

func testErrorNoFail(test *testing.T, ctx context.Context, g2product g2api.G2product, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestG2product_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	g2product.SetObserverOrigin(ctx, origin)
}

func TestG2product_GetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	g2product.SetObserverOrigin(ctx, origin)
	actual := g2product.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

func TestG2product_Init(test *testing.T) {
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	moduleName := "Test module name"
	iniParams := "{}"
	verboseLogging := int64(0)
	err := g2product.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2product, err)
}

func TestG2product_License(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	actual, err := g2product.License(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestG2product_Version(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	actual, err := g2product.Version(ctx)
	testError(test, ctx, g2product, err)
	printActual(test, actual)
}

func TestG2product_Destroy(test *testing.T) {
	ctx := context.TODO()
	g2product := getTestObject(ctx, test)
	err := g2product.Destroy(ctx)
	testError(test, ctx, g2product, err)
}
