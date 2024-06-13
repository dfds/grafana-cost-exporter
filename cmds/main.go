package main

import (
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.dfds.cloud/grafana-cost-exporter/conf"
	"go.dfds.cloud/grafana-cost-exporter/internal/loki_client"
	"go.dfds.cloud/grafana-cost-exporter/internal/metrics"
	"go.dfds.cloud/grafana-cost-exporter/internal/prom_client"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(pprof.New())

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	go worker()
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

func worker() {
	config, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	sleepInterval, err := time.ParseDuration(fmt.Sprintf("%ds", config.WorkerInterval))
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Getting Grafana Cloud cost data")

		lokiClient := loki_client.NewClient(config.Loki.Endpoint, &loki_client.Auth{
			Username: config.Loki.Auth.Username,
			Token:    config.Loki.Auth.Token,
		})
		promClient := prom_client.NewClient(config.Prometheus.Endpoint)
		gatherer := metrics.NewGatherer(lokiClient, promClient)
		//data := gatherer.GetAllMetrics()
		gatherer.GetAllMetrics()

		//metricsByCaps := metrics.ByCapability(data)
		//
		//pricingData := conf.LoadData()
		//
		//pricingProd := metrics.Pricing{
		//	NetworkTransfer: pricingData.Pricing.Prod.NetworkTransfer,
		//	Storage:         pricingData.Pricing.Prod.Storage,
		//}
		//
		//pricingDev := metrics.Pricing{
		//	NetworkTransfer: pricingData.Pricing.Dev.NetworkTransfer,
		//	Storage:         pricingData.Pricing.Dev.Storage,
		//}
		//
		//csvData := metrics.CapabilityResponseToCostCsv(metricsByCaps, pricingProd, pricingDev)

		//serialised, err := json.Marshal(csvData)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//fmt.Println(string(serialised))

		fmt.Println("New metrics published")
		time.Sleep(sleepInterval)
	}
}
