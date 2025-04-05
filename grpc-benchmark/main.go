package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

type TestConfig struct {
	ServiceName   string
	Method        string
	DataDir       string
	TotalRequests uint
	Concurrency   uint
	Timeout       time.Duration
}

func main() {
	services := []TestConfig{
		{
			ServiceName:   "ExportLogs",
			Method:        "plcruntime.PLCRuntimeService.ExportLogs",
			DataDir:       "testdata/export-logs",
			TotalRequests: 1000,
			Concurrency:   50,
			Timeout:       10 * time.Second,
		},
		{
			ServiceName:   "CreateProject",
			Method:        "plcruntime.PLCRuntimeService.CreateProject",
			DataDir:       "testdata/create-project",
			TotalRequests: 800,
			Concurrency:   30,
			Timeout:       15 * time.Second,
		},
		{
			ServiceName:   "ImportProgram",
			Method:        "plcruntime.PLCRuntimeService.ImportProgram",
			DataDir:       "testdata/import-program",
			TotalRequests: 1200,
			Concurrency:   40,
			Timeout:       20 * time.Second,
		},
		{
			ServiceName:   "ExportProject",
			Method:        "plcruntime.PLCRuntimeService.ExportProject",
			DataDir:       "testdata/export-project",
			TotalRequests: 600,
			Concurrency:   20,
			Timeout:       10 * time.Second,
		},
	}

	for _, config := range services {
		testService(config)
	}
}

func testService(config TestConfig) {
	files, err := filepath.Glob(filepath.Join(config.DataDir, "*.json"))
	if err != nil {
		fmt.Printf("[%s] Error finding test files: %v\n", config.ServiceName, err)
		return
	}

	for _, dataFile := range files {
		startTime := time.Now()
		report, err := runner.Run(
			config.Method,
			"localhost:50051",
			runner.WithProtoFile("proto/plcruntime.proto", []string{}),
			runner.WithDataFromFile(dataFile),
			runner.WithInsecure(true),
			runner.WithTotalRequests(config.TotalRequests),
			runner.WithConcurrency(config.Concurrency),
			runner.WithTimeout(config.Timeout),
		)

		if err != nil {
			fmt.Printf("[%s] Test failed: %v\n", config.ServiceName, err)
			continue
		}

		// 生成报告
		reportFile := fmt.Sprintf("reports/%s_%s.json", 
			config.ServiceName,
			time.Now().Format("20060102-150405"))
		
		if err := saveReport(report, reportFile); err != nil {
			fmt.Printf("[%s] Failed to save report: %v\n", config.ServiceName, err)
		}

		// 打印摘要
		fmt.Printf("\n=== %s 测试结果 (%s) ===\n", config.ServiceName, filepath.Base(dataFile))
		fmt.Printf("总请求数: %d\n", report.Count)
		fmt.Printf("平均耗时: %v\n", report.Average)
		fmt.Printf("最长耗时: %v\n", report.Slowest)
		fmt.Printf("最短耗时: %v\n", report.Fastest)
		fmt.Printf("QPS: %.2f\n", report.Rps)
		fmt.Printf("总耗时: %v\n", time.Since(startTime))
	}
}

func saveReport(report *runner.Report, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	p := printer.ReportPrinter{
		Out:    file,
		Report: report,
	}

	return p.Print("pretty")
}