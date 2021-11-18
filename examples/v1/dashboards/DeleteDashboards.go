// Delete dashboards returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "dashboard" in the system
	DASHBOARD_ID := os.Getenv("DASHBOARD_ID")

	body := datadog.DashboardBulkDeleteRequest{
		Data: []datadog.DashboardBulkActionData{
			datadog.DashboardBulkActionData{
				Id:   DASHBOARD_ID,
				Type: datadog.DashboardResourceType("dashboard"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	r, err := apiClient.DashboardsApi.DeleteDashboards(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboards`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}