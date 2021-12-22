package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	remoteworkflowapi "code.it4i.cz/lexis/wp4/alien4cloud-interface/client"
	kclib "code.it4i.cz/lexis/wp4/keycloak-lib"
	remoteapprovalapi "github.com/lexis-project/lexis-backend-services-interface-approval-system.git/client"
	remotedataapi "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi"
	"github.com/lexis-project/lexis-backend-services-api.git/server/approvalapi"
	"github.com/lexis-project/lexis-backend-services-api.git/server/datasetapi"
	"github.com/lexis-project/lexis-backend-services-api.git/server/heappeapi"
	"github.com/lexis-project/lexis-backend-services-api.git/server/usageManager"
	"github.com/lexis-project/lexis-backend-services-api.git/server/userorgapi"
	"github.com/lexis-project/lexis-backend-services-api.git/server/workflowapi"
	remoteuserorgapi "github.com/lexis-project/lexis-backend-services-userorg-service.git/client"
	"github.com/rs/cors"
	"github.com/segmentio/encoding/json"
	cdrapi "gitlab.com/cyclops-community/cdr-client-interface/client"
	csapi "gitlab.com/cyclops-community/cs-client-interface/client"
	udrapi "gitlab.com/cyclops-community/udr-client-interface/client"
	l "gitlab.com/cyclops-utilities/logging"
)

var (
	cfg     configuration
	version string
)

// getBasePath function is to get the base URL of the server.
// Returns:
// - String with the value of the base URL of the server.
func getBasePath() string {

	type jsonBasePath struct {
		BasePath string
	}

	bp := jsonBasePath{}

	e := json.Unmarshal(restapi.SwaggerJSON, &bp)

	if e != nil {

		l.Warning.Printf("Unmarshalling of the basepath failed: %v\n", e)

	}

	return bp.BasePath

}

// init function - reads in configuration file and creates logger
func init() {

	confFile := flag.String("conf", "./config", "configuration file path (without toml extension)")

	flag.Parse()

	//placeholder code as the default value will ensure this situation will never arise
	if len(*confFile) == 0 {

		fmt.Println("Usage: lexis-portal-api --conf=/path/to/configuration/file")

		os.Exit(0)

	}

	readConfigFile(*confFile)

	cfg = parseConfig()

	cfg.Policies = parsePolicies()

	dumpConfig(cfg)

	// when communicating with other services, they may not be secured with valid Https
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: cfg.General.InsecureSkipVerify}

	l.InitLogger(cfg.General.LogFile, cfg.General.LogLevel, cfg.General.LogToConsole)

	l.Info.Printf("Initilizing Keycloak Lib...")

	kclib.InitLib("config_keycloak.toml")

	l.Info.Printf("LEXIS Portal API Service version %v initialized", version)

}

// main function creates the database connection and launches the endpoint handlers
func main() {

	as := remoteapprovalapi.Config{
		URL: &url.URL{
			Scheme: "http",
			Host:   cfg.ApprovalService.Host + ":" + strconv.Itoa(cfg.ApprovalService.Port),
			Path:   cfg.ApprovalService.BaseURL,
		},
	}

	a := approvalapi.New(as)

	dc := remotedataapi.Config{
		URL: &url.URL{
			Scheme: "http",
			Host:   cfg.DataCatalogService.Host + ":" + strconv.Itoa(cfg.DataCatalogService.Port),
			Path:   cfg.DataCatalogService.BaseURL,
		},
	}

	d := datasetapi.New(dc, cfg.Keycloak.RedirectURL)

	uoc := remoteuserorgapi.Config{
		URL: &url.URL{
			Scheme: "http",
			Host:   cfg.UserOrgService.Host + ":" + strconv.Itoa(cfg.UserOrgService.Port),
			Path:   cfg.UserOrgService.BaseURL,
		},
	}

	uo := userorgapi.New(uoc)

	cs := csapi.Config{
		URL: &url.URL{
			Scheme: "http",
			Host:   cfg.CSService.Host + ":" + strconv.Itoa(cfg.CSService.Port),
			Path:   cfg.CSService.BaseURL,
		},
	}

	cc := cdrapi.Config{
		URL: &url.URL{
			Scheme: "http",
			Host:   cfg.CDRService.Host + ":" + strconv.Itoa(cfg.CDRService.Port),
			Path:   cfg.CDRService.BaseURL,
		},
	}

	uc := udrapi.Config{
		URL: &url.URL{
			Scheme: "http",
			Host:   cfg.UDRService.Host + ":" + strconv.Itoa(cfg.UDRService.Port),
			Path:   cfg.UDRService.BaseURL,
		},
	}

	u := usageManager.New(uoc, uc, cc, cs)

	wc := remoteworkflowapi.Config{
		URL: &url.URL{
			Scheme: "http",
			Host:   cfg.WorkflowService.Host + ":" + strconv.Itoa(cfg.WorkflowService.Port),
			Path:   cfg.WorkflowService.BaseURL,
		},
	}

	w := workflowapi.New(wc)

	hp := heappeapi.New()

	// Initiate the http handler, with the objects that are implementing the business logic.
	rc := restapi.Config{
		ApprovalSystemManagementAPI: a,
		DataSetManagementAPI:        d,
		StagingAPI:                  d,
		UsageManagementAPI:          u,
		UserOrgManagementAPI:        uo,
		WorkflowManagementAPI:       w,
		ClusterInformationAPI:       hp,
		Logger:                      l.Info.Printf,
		AuthAPIKeyHeader:            AuthAPIKey,
		AuthAPIKeyParam:             AuthAPIKey,
		AuthKeycloak:                AuthKeycloak,
		Authorizer:                  Authorizer,
	}

	h, e := restapi.Handler(rc)

	if e != nil {

		l.Error.Printf("Error creating REST handler: %s ...Exiting\n", e)

		os.Exit(2)

	}

	hcors := cors.New(
		cors.Options{
			Debug:          (cfg.General.LogLevel == "DEBUG"),
			AllowedOrigins: cfg.General.CORSOrigins,
			AllowedHeaders: cfg.General.CORSHeaders,
			AllowedMethods: cfg.General.CORSMethods,
		}).Handler(h)

	// note that this runs on all interfaces right now
	serviceLocation := ":" + strconv.Itoa(cfg.General.ServerPort)

	l.Info.Printf("Starting to serve the LEXIS Portal API Service, access server on https://localhost:%v\n", serviceLocation)

	// Run the standard http server
	if cfg.General.HTTPSEnabled {

		l.Error.Printf((http.ListenAndServeTLS(serviceLocation, cfg.General.CertificateFile, cfg.General.CertificateKey, hcors)).Error())

	} else {

		l.Warning.Printf("Running without TLS security - do not use in production scenario...")

		l.Error.Printf((http.ListenAndServe(serviceLocation, hcors)).Error())

	}

}
