package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_APIServiceCABundleInjectorConfig = map[string]string{
	"":             "APIServiceCABundleInjectorConfig provides information to configure an APIService CA Bundle Injector controller",
	"caBundleFile": "caBundleFile holds the ca bundle to apply to APIServices.",
}

func (APIServiceCABundleInjectorConfig) SwaggerDoc() map[string]string {
	return map_APIServiceCABundleInjectorConfig
}

var map_ConfigMapCABundleInjectorConfig = map[string]string{
	"":             "ConfigMapCABundleInjectorConfig provides information to configure a ConfigMap CA Bundle Injector controller",
	"caBundleFile": "caBundleFile holds the ca bundle to apply to ConfigMaps.",
}

func (ConfigMapCABundleInjectorConfig) SwaggerDoc() map[string]string {
	return map_ConfigMapCABundleInjectorConfig
}

var map_ServiceCertSignerOperatorConfig = map[string]string{
	"": "ServiceCertSignerOperatorConfig provides information to configure an operator to manage the service cert signing controllers",
}

func (ServiceCertSignerOperatorConfig) SwaggerDoc() map[string]string {
	return map_ServiceCertSignerOperatorConfig
}

var map_ServiceCertSignerOperatorConfigList = map[string]string{
	"":      "ServiceCertSignerOperatorConfigList is a collection of items",
	"items": "Items contains the items",
}

func (ServiceCertSignerOperatorConfigList) SwaggerDoc() map[string]string {
	return map_ServiceCertSignerOperatorConfigList
}

var map_ServiceCertSignerOperatorConfigSpec = map[string]string{
	"serviceServingCertSignerConfig":   "serviceServingCertSignerConfig holds a sparse config that the user wants for this component.  It only needs to be the overrides from the defaults it will end up overlaying in the following order: 1. hardcoded default 2. this config",
	"apiServiceCABundleInjectorConfig": "apiServiceCABundleInjectorConfig holds a sparse config that the user wants for this component.  It only needs to be the overrides from the defaults it will end up overlaying in the following order: 1. hardcoded default 2. this config",
	"configMapCABundleInjectorConfig":  "configMapCABundleInjectorConfig holds a sparse config that the user wants for this component.  It only needs to be the overrides from the defaults it will end up overlaying in the following order: 1. hardcoded default 2. this config",
}

func (ServiceCertSignerOperatorConfigSpec) SwaggerDoc() map[string]string {
	return map_ServiceCertSignerOperatorConfigSpec
}

var map_ServiceServingCertSignerConfig = map[string]string{
	"":                     "ServiceServingCertSignerConfig provides information to configure a serving serving cert signing controller",
	"signer":               "signer holds the signing information used to automatically sign serving certificates.",
	"intermediateCertFile": "IntermediateCertFile is the name of a file containing a PEM-encoded certificate. Only required if the initial CA has been rotated. The certificate should consist of the public key of the current CA signed by the private key of the previous CA. When included with a serving cert generated by the current CA, this certificate should allow clients with a stale CA bundle to trust the serving cert.",
}

func (ServiceServingCertSignerConfig) SwaggerDoc() map[string]string {
	return map_ServiceServingCertSignerConfig
}

// AUTO-GENERATED FUNCTIONS END HERE
