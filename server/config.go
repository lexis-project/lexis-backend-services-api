package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/segmentio/encoding/json"

	"github.com/spf13/viper"
	l "gitlab.com/cyclops-utilities/logging"
)

// General service structs
type apiKey struct {
	Enabled bool `json:"enabled"`
	Key     string
	Place   string
	Token   string `json:"token"`
}

type policy struct {
	Endpoint string
	Method   string
	Allowed  []string
	Checks   []string
}

type generalConfig struct {
	CertificateFile    string `json:"certificate_file"`
	CertificateKey     string `json:"certificate_key"`
	CORSEnabled        bool
	CORSHeaders        []string
	CORSMethods        []string
	CORSOrigins        []string
	HTTPSEnabled       bool `json:"https_enabled"`
	InsecureSkipVerify bool
	LogFile            string `json:"log_file"`
	LogLevel           string `json:"log_level"`
	LogToConsole       bool   `json:"log_to_console"`
	ServerPort         int    `json:"server_port"`
}

type keycloakConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Enabled      bool   `json:"enabled"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Realm        string `json:"realm"`
	RedirectURL  string `json:"redirect_url"`
	UseHTTP      bool   `json:"use_http"`
}

// Interfaces connection structs
type approvalServiceConfig struct {
	BaseURL string `json:"base_url"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type csConfig struct {
	BaseURL string `json:"base_url"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type cdrConfig struct {
	BaseURL string `json:"base_url"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type dataCatalogServiceConfig struct {
	BaseURL string `json:"base_url"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type udrConfig struct {
	BaseURL string `json:"base_url"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type userOrgServiceConfig struct {
	BaseURL string `json:"base_url"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

type workflowServiceConfig struct {
	BaseURL string `json:"base_url"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// Main config struct
type configuration struct {
	APIKey             apiKey `json:"api_key"`
	PoliciesFile       string
	Policies           []policy
	AuthzEnabled       bool
	ApprovalService    approvalServiceConfig    `json:"approval_service"`
	CSService          csConfig                 `json:"cs_service"`
	CDRService         cdrConfig                `json:"cdr_service"`
	DataCatalogService dataCatalogServiceConfig `json:"data_catalog_service"`
	General            generalConfig            `json:"general"`
	Keycloak           keycloakConfig           `json:"keycloak"`
	UDRService         udrConfig                `json:"udr_service"`
	UserOrgService     userOrgServiceConfig     `json:"user_org_service"`
	WorkflowService    workflowServiceConfig    `json:"workflow_service"`
}

// parsePolicies handles the filling of the config struct that contains the authz
// policies that the service will enforce for any entry/call.
// Returns:
// - pols: and array of policies
func parsePolicies() (policies []policy) {

	policyFile, e := os.Open(cfg.PoliciesFile)

	if e != nil {

		panic("Authz policy file couldn't be found\n")

	}

	r := csv.NewReader(policyFile)
	r.Comma = '|'
	r.Comment = '-'
	r.TrimLeadingSpace = true

	policiesList, e := r.ReadAll()

	if e != nil {

		panic("Authz policy file couldn't be readed\n")

	}

	for _, pol := range policiesList {

		newPolicy := policy{
			Endpoint: strings.TrimSpace(pol[0]),
			Method:   strings.TrimSpace(pol[1]),
			Allowed:  strings.Fields(pol[2]),
			Checks:   strings.Fields(pol[3]),
		}

		policies = append(policies, newPolicy)

	}

	return

}

// parseConfig creates a valid configuration from the input file read with viper
func parseConfig() (c configuration) {

	c = configuration{

		APIKey: apiKey{
			Enabled: viper.GetBool("apikey.enabled"),
			Key:     viper.GetString("apikey.key"),
			Place:   viper.GetString("apikey.place"),
			Token:   viper.GetString("apikey.token"),
		},

		General: generalConfig{
			CertificateFile:    viper.GetString("general.certificatefile"),
			CertificateKey:     viper.GetString("general.certificatekey"),
			CORSEnabled:        viper.GetBool("general.corsenabled"),
			CORSHeaders:        viper.GetStringSlice("general.corsheaders"),
			CORSMethods:        viper.GetStringSlice("general.corsmethods"),
			CORSOrigins:        viper.GetStringSlice("general.corsorigins"),
			HTTPSEnabled:       viper.GetBool("general.httpsenabled"),
			InsecureSkipVerify: viper.GetBool("general.insecureskipverify"),
			LogFile:            viper.GetString("general.logfile"),
			LogToConsole:       viper.GetBool("general.logtoconsole"),
			LogLevel:           viper.GetString("general.loglevel"),
			ServerPort:         viper.GetInt("general.serverport"),
		},

		Keycloak: keycloakConfig{
			ClientID:     viper.GetString("keycloak.clientid"),
			ClientSecret: viper.GetString("keycloak.clientsecret"),
			Enabled:      viper.GetBool("keycloak.enabled"),
			Host:         viper.GetString("keycloak.host"),
			Port:         viper.GetInt("keycloak.port"),
			Realm:        viper.GetString("keycloak.realm"),
			RedirectURL:  viper.GetString("keycloak.redirecturl"),
			UseHTTP:      viper.GetBool("keycloak.usehttp"),
		},

		ApprovalService: approvalServiceConfig{
			Host:    viper.GetString("approvalservice.host"),
			Port:    viper.GetInt("approvalservice.port"),
			BaseURL: viper.GetString("approvalservice.baseurl"),
		},

		CSService: csConfig{
			Host:    viper.GetString("csservice.host"),
			Port:    viper.GetInt("csservice.port"),
			BaseURL: viper.GetString("csservice.baseurl"),
		},

		CDRService: cdrConfig{
			Host:    viper.GetString("cdrservice.host"),
			Port:    viper.GetInt("cdrservice.port"),
			BaseURL: viper.GetString("cdrservice.baseurl"),
		},

		DataCatalogService: dataCatalogServiceConfig{
			Host:    viper.GetString("datacatalogservice.host"),
			Port:    viper.GetInt("datacatalogservice.port"),
			BaseURL: viper.GetString("datacatalogservice.baseurl"),
		},

		UDRService: udrConfig{
			Host:    viper.GetString("udrservice.host"),
			Port:    viper.GetInt("udrservice.port"),
			BaseURL: viper.GetString("udrservice.baseurl"),
		},

		UserOrgService: userOrgServiceConfig{
			Host:    viper.GetString("userorgservice.host"),
			Port:    viper.GetInt("userorgservice.port"),
			BaseURL: viper.GetString("userorgservice.baseurl"),
		},

		WorkflowService: workflowServiceConfig{
			Host:    viper.GetString("workflowservice.host"),
			Port:    viper.GetInt("workflowservice.port"),
			BaseURL: viper.GetString("workflowservice.baseurl"),
		},

		PoliciesFile: viper.GetString("authz.policiesfile"),
		AuthzEnabled: viper.GetBool("authz.authzenabled"),
	}

	return

}

// dumpConfig dumps the configuration in json format to the log system
func dumpConfig(c configuration) {

	cfgCopy := c

	// deal with configuration params that should be masked
	cfgCopy.Keycloak.ClientSecret = masked(c.Keycloak.ClientSecret, 4)
	cfgCopy.APIKey.Token = masked(c.APIKey.Token, 4)

	// mmrshalindent creates a string containing newlines; each line starts with
	// two spaces and two spaces are added for each indent...
	configJson, _ := json.MarshalIndent(cfgCopy, "  ", "  ")

	l.Info.Printf("Configuration settings:\n")
	l.Info.Printf("%v\n", string(configJson))

}

// function which returns asterisks in place of string except for last um=nmakedChars chars
func masked(s string, unmaskedChars int) (returnString string) {

	if len(s) <= unmaskedChars {

		returnString = s

		return

	}

	asteriskString := strings.Repeat("*", (len(s) - unmaskedChars))
	returnString = asteriskString + string(s[len(s)-unmaskedChars:])

	return

}

// readConfigfile reads in a config file in toml format with the name specified by
// the input parameter filename
func readConfigFile(f string) {

	viper.SetConfigName(f) // name of config file (without extension)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".") // path to look for the config file in

	if e := viper.ReadInConfig(); e != nil {

		fmt.Printf("Fatal error reading config file: %v\n", e)

		os.Exit(1)

	}

}
