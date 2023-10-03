//go:build linux

package g2product

import (
	"context"
	"fmt"

	"github.com/senzing/go-logging/logging"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleG2product_SetObserverOrigin() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	origin := "Machine: nn; Task: UnitTest"
	g2product.SetObserverOrigin(ctx, origin)
	// Output:
}

func ExampleG2product_GetObserverOrigin() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2config/g2product_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	origin := "Machine: nn; Task: UnitTest"
	g2product.SetObserverOrigin(ctx, origin)
	result := g2product.GetObserverOrigin(ctx)
	fmt.Println(result)
	// Output: Machine: nn; Task: UnitTest
}

func ExampleG2product_Init() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	moduleName := "Test module name"
	iniParams := "{}"
	verboseLogging := 0
	err := g2product.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		// This should produce a "senzing-60164002" error.
	}
	// Output:
}

func ExampleG2product_License() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	result, err := g2product.License(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"customer":"Senzing Public Test License","contract":"EVALUATION - support@senzing.com","issueDate":"2022-11-29","licenseType":"EVAL (Solely for non-productive use)","licenseLevel":"STANDARD","billing":"MONTHLY","expireDate":"2023-11-29","recordLimit":50000}
}

func ExampleG2product_SetLogLevel() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	err := g2product.SetLogLevel(ctx, logging.LevelInfoName)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2product_ValidateLicenseFile() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	licenseFilePath := "/etc/opt/senzing/g2.lic"
	result, err := g2product.ValidateLicenseFile(ctx, licenseFilePath)
	if err != nil {
		fmt.Println("Invalid license")
	} else {
		fmt.Println(result)
	}
	// Output: Success
}

func ExampleG2product_ValidateLicenseStringBase64() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	licenseString := "AQAAADgCAAAAAAAAU2VuemluZyBQdWJsaWMgVGVzdCBMaWNlbnNlAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAARVZBTFVBVElPTiAtIHN1cHBvcnRAc2VuemluZy5jb20AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADIwMjItMTEtMjkAAAAAAAAAAAAARVZBTAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFNUQU5EQVJEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFDDAAAAAAAAMjAyMy0xMS0yOQAAAAAAAAAAAABNT05USExZAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACQfw5e19QAHetkvd+vk0cYHtLaQCLmgx2WUfLorDfLQq15UXmOawNIXc1XguPd8zJtnOaeI6CB2smxVaj10mJE2ndGPZ1JjGk9likrdAj3rw+h6+C/Lyzx/52U8AuaN1kWgErDKdNE9qL6AnnN5LLi7Xs87opP7wbVMOdzsfXx2Xi3H7dSDIam7FitF6brSFoBFtIJac/V/Zc3b8jL/a1o5b1eImQldaYcT4jFrRZkdiVO/SiuLslEb8or3alzT0XsoUJnfQWmh0BjehBK9W74jGw859v/L1SGn1zBYKQ4m8JBiUOytmc9ekLbUKjIg/sCdmGMIYLywKqxb9mZo2TLZBNOpYWVwfaD/6O57jSixfJEHcLx30RPd9PKRO0Nm+4nPdOMMLmd4aAcGPtGMpI6ldTiK9hQyUfrvc9z4gYE3dWhz2Qu3mZFpaAEuZLlKtxaqEtVLWIfKGxwxPargPEfcLsv+30fdjSy8QaHeU638tj67I0uCEgnn5aB8pqZYxLxJx67hvVKOVsnbXQRTSZ00QGX1yTA+fNygqZ5W65wZShhICq5Fz8wPUeSbF7oCcE5VhFfDnSyi5v0YTNlYbF8LOAqXPTi+0KP11Wo24PjLsqYCBVvmOg9ohZ89iOoINwUB32G8VucRfgKKhpXhom47jObq4kSnihxRbTwJRx4o"
	result, err := g2product.ValidateLicenseStringBase64(ctx, licenseString)
	if err != nil {
		fmt.Println("Invalid license")
	} else {
		fmt.Println(result)
	}
	// Output: Success
}

func ExampleG2product_Version() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	result, err := g2product.Version(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(result, 43))
	// Output: {"PRODUCT_NAME":"Senzing API","VERSION":...
}

func ExampleG2product_Destroy() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getG2Product(ctx)
	err := g2product.Destroy(ctx)
	if err != nil {
		// This should produce a "senzing-60164001" error.
	}
	// Output:
}
