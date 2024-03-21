package pkg

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"os"
)

// GetHelmChartReleases - Example of listing helm chart releases for a given chart name
func GetHelmChartReleases(chartName *string) ([]*release.Release, error) {
	var resultList []*release.Release

	// Create the Helm client
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), nil); err != nil {
		return nil, err
	}

	// Create the list action
	list := action.NewList(actionConfig)
	list.AllNamespaces = true

	// Get the list of releases
	releases, err := list.Run()
	if err != nil {
		return nil, err
	}

	// Filter the releases by chart name
	for _, rel := range releases {
		if rel.Chart != nil && rel.Chart.Metadata != nil && rel.Chart.Metadata.Name == *chartName {
			resultList = append(resultList, rel)
		}
	}

	// Return the list of releases
	return resultList, nil
}
