// Copyright 2022 The Kube-burner Authors.
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

package main

import (
	"embed"
	_ "embed"
	"log"
	"os"
	"time"

	"github.com/cloud-bulldozer/go-commons/indexers"
	"github.com/kube-burner/kube-burner/pkg/util"
	"github.com/kube-burner/kube-burner/pkg/workloads"
	uid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"kube-burner.io/ocp"
)

//go:embed config/*
var ocpConfig embed.FS

func openShiftCmd() *cobra.Command {
	var workloadConfig workloads.Config
	var wh workloads.WorkloadHelper
	var metricsProfileType string
	ocpCmd := &cobra.Command{
		Use:  "kube-burner-ocp",
		Long: `kube-burner plugin designed to be used with OpenShift clusters as a quick way to run well-known workloads`,
	}
	ocpCmd.PersistentFlags().StringVar(&workloadConfig.EsServer, "es-server", "", "Elastic Search endpoint")
	ocpCmd.PersistentFlags().StringVar(&workloadConfig.EsIndex, "es-index", "", "Elastic Search index")
	localIndexing := ocpCmd.PersistentFlags().Bool("local-indexing", false, "Enable local indexing")
	ocpCmd.PersistentFlags().StringVar(&workloadConfig.MetricsEndpoint, "metrics-endpoint", "", "YAML file with a list of metric endpoints")
	ocpCmd.PersistentFlags().BoolVar(&workloadConfig.Alerting, "alerting", true, "Enable alerting")
	ocpCmd.PersistentFlags().StringVar(&workloadConfig.UUID, "uuid", uid.NewV4().String(), "Benchmark UUID")
	ocpCmd.PersistentFlags().DurationVar(&workloadConfig.Timeout, "timeout", 4*time.Hour, "Benchmark timeout")
	ocpCmd.PersistentFlags().IntVar(&workloadConfig.QPS, "qps", 20, "QPS")
	ocpCmd.PersistentFlags().IntVar(&workloadConfig.Burst, "burst", 20, "Burst")
	ocpCmd.PersistentFlags().BoolVar(&workloadConfig.Gc, "gc", true, "Garbage collect created namespaces")
	ocpCmd.PersistentFlags().BoolVar(&workloadConfig.GcMetrics, "gc-metrics", false, "Collect metrics during garbage collection")
	ocpCmd.PersistentFlags().StringVar(&workloadConfig.UserMetadata, "user-metadata", "", "User provided metadata file, in YAML format")
	extract := ocpCmd.PersistentFlags().Bool("extract", false, "Extract workload in the current directory")
	ocpCmd.PersistentFlags().StringVar(&metricsProfileType, "profile-type", "both", "Metrics profile to use, supported options are: regular, reporting or both")
	ocpCmd.MarkFlagsRequiredTogether("es-server", "es-index")
	ocpCmd.MarkFlagsMutuallyExclusive("es-server", "local-indexing")
	ocpCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		util.ConfigureLogging(cmd)
		if *extract {
			if err := workloads.ExtractWorkload(ocpConfig, cmd.Name(), "alerts.yml", "metrics.yml", "metrics-aggregated.yml", "metrics-report.yml"); err != nil {
				log.Fatal(err.Error())
			}
			os.Exit(0)
		}
		if workloadConfig.EsServer != "" || *localIndexing {
			workloadConfig.Indexing = true
			if workloadConfig.EsServer != "" {
				workloadConfig.Indexer = indexers.ElasticIndexer
			} else {
				workloadConfig.Indexer = indexers.LocalIndexer
			}
			wh = workloads.NewWorkloadHelper(workloadConfig, workloads.OCP, ocpConfig)
		}
	}
	ocpCmd.AddCommand(
		ocp.NewClusterDensity(&wh, "cluster-density-v2"),
		ocp.NewClusterDensity(&wh, "cluster-density-ms"),
		ocp.NewCrdScale(&wh),
		ocp.NewNetworkPolicy(&wh, "networkpolicy-multitenant"),
		ocp.NewNetworkPolicy(&wh, "networkpolicy-matchlabels"),
		ocp.NewNetworkPolicy(&wh, "networkpolicy-matchexpressions"),
		ocp.NewNodeDensity(&wh),
		ocp.NewNodeDensityHeavy(&wh),
		ocp.NewNodeDensityCNI(&wh),
		ocp.NewIndex(&wh.MetricsEndpoint, &wh.MetadataAgent),
		ocp.NewPVCDensity(&wh),
		ocp.NewWebBurner(&wh, "web-burner-init"),
		ocp.NewWebBurner(&wh, "web-burner-node-density"),
		ocp.NewWebBurner(&wh, "web-burner-cluster-density"),
	)
	util.SetupCmd(ocpCmd)
	return ocpCmd
}

func main() {
	if openShiftCmd().Execute() != nil {
		os.Exit(1)
	}
}
