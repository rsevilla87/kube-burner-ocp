// Copyright 2024 The Kube-burner Authors.
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
	"fmt"
	"os"
	"time"

	"github.com/kube-burner/kube-burner/pkg/workloads"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// Returns virt-density workload
func NewVirtDensity(wh *workloads.WorkloadHelper) *cobra.Command {
	var vmsPerNode int
	var vmReadyThreshold time.Duration
	var metricsProfiles []string
	var rc int
	cmd := &cobra.Command{
		Use:          "virt-density",
		Short:        "Runs virt-density workload",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			totalPods := clusterMetadata.WorkerNodesCount * vmsPerNode
			vmCount, err := wh.MetadataAgent.GetCurrentPodCount()
			if err != nil {
				log.Fatal(err.Error())
			}
			os.Setenv("JOB_ITERATIONS", fmt.Sprint(totalPods-vmCount))
			os.Setenv("VM_READY_THRESHOLD", fmt.Sprintf("%v", vmReadyThreshold))
		},
		Run: func(cmd *cobra.Command, args []string) {
			setMetrics(cmd, metricsProfiles)
			rc = wh.Run(cmd.Name())
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			os.Exit(rc)
		},
	}
	cmd.Flags().IntVar(&vmsPerNode, "pods-per-node", 245, "Pods per node")
	cmd.Flags().DurationVar(&vmReadyThreshold, "pod-ready-threshold", 15*time.Second, "Pod ready timeout threshold")
	cmd.Flags().StringSliceVar(&metricsProfiles, "metrics-profile", []string{"metrics.yml"}, "Comma separated list of metrics profiles to use")
	return cmd
}
