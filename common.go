// Copyright 2023 The Kube-burner Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ocp

import (
	"os"

	"github.com/kube-burner/kube-burner/pkg/measurements/types"
	"github.com/spf13/cobra"
)

func getMetrics(cmd *cobra.Command, metricsProfile string) []string {
	var metricsProfiles []string
	profileType, _ := cmd.Root().PersistentFlags().GetString("profile-type")
	switch ProfileType(profileType) {
	case Reporting:
		metricsProfiles = []string{"metrics-report.yml"}
		os.Setenv("POD_LATENCY_METRICS", string(types.Quantiles))
	case Regular:
		metricsProfiles = []string{metricsProfile}
	case Both:
		metricsProfiles = []string{"metrics-report.yml", metricsProfile}
	}
	return metricsProfiles
}
