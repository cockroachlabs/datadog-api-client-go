// Get the list of available monthly custom reports returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("GetMonthlyCustomReports", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageMeteringApi.GetMonthlyCustomReports(ctx, *datadog.NewGetMonthlyCustomReportsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetMonthlyCustomReports`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetMonthlyCustomReports`:\n%s\n", responseContent)
}