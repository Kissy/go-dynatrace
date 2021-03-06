// Code generated by go-swagger; DO NOT EDIT.

package dynatrace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScopeConditions scope conditions
// swagger:model ScopeConditions
type ScopeConditions struct {

	// Only applies to this host group.
	HostGroup string `json:"hostGroup,omitempty"`

	// Only applies to this process group. Note that this can't be transferred between different clusters or environments.
	ProcessGroup string `json:"processGroup,omitempty"`

	// Only applies to this service technology.
	// Enum: [ACTIVE_MQ ACTIVE_MQ_ARTEMIS ADO_NET AIX AKKA AMAZON_REDSHIFT AMQP APACHE_CAMEL APACHE_CASSANDRA APACHE_COUCH_DB APACHE_DERBY APACHE_HTTP_SERVER APACHE_KAFKA APACHE_SOLR APACHE_STORM APACHE_SYNAPSE APACHE_TOMCAT APPARMOR APPLICATION_INSIGHTS_SDK ASP_DOTNET ASP_DOTNET_CORE ASP_DOTNET_CORE_SIGNALR ASP_DOTNET_SIGNALR AWS_LAMBDA AWS_RDS AWS_SERVICE AXIS AZURE_FUNCTIONS AZURE_SERVICE_BUS AZURE_SERVICE_FABRIC AZURE_STORAGE BOSHBPM CITRIX CLOUDFOUNDRY CLOUDFOUNDRY_AUCTIONEER CLOUDFOUNDRY_BOSH CLOUDFOUNDRY_GOROUTER COLDFUSION CONTAINERD COUCHBASE CRIO CXF DATASTAX DB2 DIEGO_CELL DOCKER DOTNET DOTNET_REMOTING ELASTIC_SEARCH ENVOY ERLANG ETCD F5_LTM GARDEN GLASSFISH GO GRPC GRSECURITY HADOOP HADOOP_HDFS HADOOP_YARN HAPROXY HEAT HESSIAN HORNET_Q IBM_CICS_REGION IBM_CICS_TRANSACTION_GATEWAY IBM_IMS_CONNECT_REGION IBM_IMS_CONTROL_REGION IBM_IMS_MESSAGE_PROCESSING_REGION IBM_IMS_SOAP_GATEWAY IBM_INTEGRATION_BUS IBM_MQ IBM_MQ_CLIENT IBM_WEBSHPRERE_APPLICATION_SERVER IBM_WEBSHPRERE_LIBERTY IIS IIS_APP_POOL ISTIO JAVA JAX_WS JBOSS JBOSS_EAP JDK_HTTP_SERVER JERSEY JETTY JRUBY JYTHON KUBERNETES LIBVIRT LINKERD MARIADB MEMCACHED MICROSOFT_SQL_SERVER MONGODB MSSQL_CLIENT MULE_ESB MYSQL MYSQL_CONNECTOR NETFLIX_SERVO NETTY NGINX NODE_JS ONEAGENT_SDK OPENCENSUS OPENSHIFT OPENSTACK_COMPUTE OPENSTACK_CONTROLLER OPENTELEMETRY OPENTRACING OPEN_LIBERTY ORACLE_DATABASE ORACLE_WEBLOGIC OWIN PERL PHP PHP_FPM PLAY POSTGRE_SQL POSTGRE_SQL_DOTNET_DATA_PROVIDER PROGRESS PYTHON RABBIT_MQ REDIS RESTEASY RESTLET RIAK RUBY SAG_WEBMETHODS_IS SAP SAP_HANADB SAP_HYBRIS SAP_MAXDB SAP_SYBASE SCALA SELINUX SHAREPOINT SPARK SPRING SQLITE THRIFT TIBCO TIBCO_BUSINESS_WORKS TIBCO_EMS VARNISH_CACHE VIM2 VIRTUAL_MACHINE_KVM VIRTUAL_MACHINE_QEMU WILDFLY WINDOWS_CONTAINERS WINK ZERO_MQ]
	ServiceTechnology string `json:"serviceTechnology,omitempty"`

	// Only apply to process groups matching this tag.
	TagOfProcessGroup string `json:"tagOfProcessGroup,omitempty"`
}

// Validate validates this scope conditions
func (m *ScopeConditions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceTechnology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scopeConditionsTypeServiceTechnologyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE_MQ","ACTIVE_MQ_ARTEMIS","ADO_NET","AIX","AKKA","AMAZON_REDSHIFT","AMQP","APACHE_CAMEL","APACHE_CASSANDRA","APACHE_COUCH_DB","APACHE_DERBY","APACHE_HTTP_SERVER","APACHE_KAFKA","APACHE_SOLR","APACHE_STORM","APACHE_SYNAPSE","APACHE_TOMCAT","APPARMOR","APPLICATION_INSIGHTS_SDK","ASP_DOTNET","ASP_DOTNET_CORE","ASP_DOTNET_CORE_SIGNALR","ASP_DOTNET_SIGNALR","AWS_LAMBDA","AWS_RDS","AWS_SERVICE","AXIS","AZURE_FUNCTIONS","AZURE_SERVICE_BUS","AZURE_SERVICE_FABRIC","AZURE_STORAGE","BOSHBPM","CITRIX","CLOUDFOUNDRY","CLOUDFOUNDRY_AUCTIONEER","CLOUDFOUNDRY_BOSH","CLOUDFOUNDRY_GOROUTER","COLDFUSION","CONTAINERD","COUCHBASE","CRIO","CXF","DATASTAX","DB2","DIEGO_CELL","DOCKER","DOTNET","DOTNET_REMOTING","ELASTIC_SEARCH","ENVOY","ERLANG","ETCD","F5_LTM","GARDEN","GLASSFISH","GO","GRPC","GRSECURITY","HADOOP","HADOOP_HDFS","HADOOP_YARN","HAPROXY","HEAT","HESSIAN","HORNET_Q","IBM_CICS_REGION","IBM_CICS_TRANSACTION_GATEWAY","IBM_IMS_CONNECT_REGION","IBM_IMS_CONTROL_REGION","IBM_IMS_MESSAGE_PROCESSING_REGION","IBM_IMS_SOAP_GATEWAY","IBM_INTEGRATION_BUS","IBM_MQ","IBM_MQ_CLIENT","IBM_WEBSHPRERE_APPLICATION_SERVER","IBM_WEBSHPRERE_LIBERTY","IIS","IIS_APP_POOL","ISTIO","JAVA","JAX_WS","JBOSS","JBOSS_EAP","JDK_HTTP_SERVER","JERSEY","JETTY","JRUBY","JYTHON","KUBERNETES","LIBVIRT","LINKERD","MARIADB","MEMCACHED","MICROSOFT_SQL_SERVER","MONGODB","MSSQL_CLIENT","MULE_ESB","MYSQL","MYSQL_CONNECTOR","NETFLIX_SERVO","NETTY","NGINX","NODE_JS","ONEAGENT_SDK","OPENCENSUS","OPENSHIFT","OPENSTACK_COMPUTE","OPENSTACK_CONTROLLER","OPENTELEMETRY","OPENTRACING","OPEN_LIBERTY","ORACLE_DATABASE","ORACLE_WEBLOGIC","OWIN","PERL","PHP","PHP_FPM","PLAY","POSTGRE_SQL","POSTGRE_SQL_DOTNET_DATA_PROVIDER","PROGRESS","PYTHON","RABBIT_MQ","REDIS","RESTEASY","RESTLET","RIAK","RUBY","SAG_WEBMETHODS_IS","SAP","SAP_HANADB","SAP_HYBRIS","SAP_MAXDB","SAP_SYBASE","SCALA","SELINUX","SHAREPOINT","SPARK","SPRING","SQLITE","THRIFT","TIBCO","TIBCO_BUSINESS_WORKS","TIBCO_EMS","VARNISH_CACHE","VIM2","VIRTUAL_MACHINE_KVM","VIRTUAL_MACHINE_QEMU","WILDFLY","WINDOWS_CONTAINERS","WINK","ZERO_MQ"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scopeConditionsTypeServiceTechnologyPropEnum = append(scopeConditionsTypeServiceTechnologyPropEnum, v)
	}
}

const (

	// ScopeConditionsServiceTechnologyACTIVEMQ captures enum value "ACTIVE_MQ"
	ScopeConditionsServiceTechnologyACTIVEMQ string = "ACTIVE_MQ"

	// ScopeConditionsServiceTechnologyACTIVEMQARTEMIS captures enum value "ACTIVE_MQ_ARTEMIS"
	ScopeConditionsServiceTechnologyACTIVEMQARTEMIS string = "ACTIVE_MQ_ARTEMIS"

	// ScopeConditionsServiceTechnologyADONET captures enum value "ADO_NET"
	ScopeConditionsServiceTechnologyADONET string = "ADO_NET"

	// ScopeConditionsServiceTechnologyAIX captures enum value "AIX"
	ScopeConditionsServiceTechnologyAIX string = "AIX"

	// ScopeConditionsServiceTechnologyAKKA captures enum value "AKKA"
	ScopeConditionsServiceTechnologyAKKA string = "AKKA"

	// ScopeConditionsServiceTechnologyAMAZONREDSHIFT captures enum value "AMAZON_REDSHIFT"
	ScopeConditionsServiceTechnologyAMAZONREDSHIFT string = "AMAZON_REDSHIFT"

	// ScopeConditionsServiceTechnologyAMQP captures enum value "AMQP"
	ScopeConditionsServiceTechnologyAMQP string = "AMQP"

	// ScopeConditionsServiceTechnologyAPACHECAMEL captures enum value "APACHE_CAMEL"
	ScopeConditionsServiceTechnologyAPACHECAMEL string = "APACHE_CAMEL"

	// ScopeConditionsServiceTechnologyAPACHECASSANDRA captures enum value "APACHE_CASSANDRA"
	ScopeConditionsServiceTechnologyAPACHECASSANDRA string = "APACHE_CASSANDRA"

	// ScopeConditionsServiceTechnologyAPACHECOUCHDB captures enum value "APACHE_COUCH_DB"
	ScopeConditionsServiceTechnologyAPACHECOUCHDB string = "APACHE_COUCH_DB"

	// ScopeConditionsServiceTechnologyAPACHEDERBY captures enum value "APACHE_DERBY"
	ScopeConditionsServiceTechnologyAPACHEDERBY string = "APACHE_DERBY"

	// ScopeConditionsServiceTechnologyAPACHEHTTPSERVER captures enum value "APACHE_HTTP_SERVER"
	ScopeConditionsServiceTechnologyAPACHEHTTPSERVER string = "APACHE_HTTP_SERVER"

	// ScopeConditionsServiceTechnologyAPACHEKAFKA captures enum value "APACHE_KAFKA"
	ScopeConditionsServiceTechnologyAPACHEKAFKA string = "APACHE_KAFKA"

	// ScopeConditionsServiceTechnologyAPACHESOLR captures enum value "APACHE_SOLR"
	ScopeConditionsServiceTechnologyAPACHESOLR string = "APACHE_SOLR"

	// ScopeConditionsServiceTechnologyAPACHESTORM captures enum value "APACHE_STORM"
	ScopeConditionsServiceTechnologyAPACHESTORM string = "APACHE_STORM"

	// ScopeConditionsServiceTechnologyAPACHESYNAPSE captures enum value "APACHE_SYNAPSE"
	ScopeConditionsServiceTechnologyAPACHESYNAPSE string = "APACHE_SYNAPSE"

	// ScopeConditionsServiceTechnologyAPACHETOMCAT captures enum value "APACHE_TOMCAT"
	ScopeConditionsServiceTechnologyAPACHETOMCAT string = "APACHE_TOMCAT"

	// ScopeConditionsServiceTechnologyAPPARMOR captures enum value "APPARMOR"
	ScopeConditionsServiceTechnologyAPPARMOR string = "APPARMOR"

	// ScopeConditionsServiceTechnologyAPPLICATIONINSIGHTSSDK captures enum value "APPLICATION_INSIGHTS_SDK"
	ScopeConditionsServiceTechnologyAPPLICATIONINSIGHTSSDK string = "APPLICATION_INSIGHTS_SDK"

	// ScopeConditionsServiceTechnologyASPDOTNET captures enum value "ASP_DOTNET"
	ScopeConditionsServiceTechnologyASPDOTNET string = "ASP_DOTNET"

	// ScopeConditionsServiceTechnologyASPDOTNETCORE captures enum value "ASP_DOTNET_CORE"
	ScopeConditionsServiceTechnologyASPDOTNETCORE string = "ASP_DOTNET_CORE"

	// ScopeConditionsServiceTechnologyASPDOTNETCORESIGNALR captures enum value "ASP_DOTNET_CORE_SIGNALR"
	ScopeConditionsServiceTechnologyASPDOTNETCORESIGNALR string = "ASP_DOTNET_CORE_SIGNALR"

	// ScopeConditionsServiceTechnologyASPDOTNETSIGNALR captures enum value "ASP_DOTNET_SIGNALR"
	ScopeConditionsServiceTechnologyASPDOTNETSIGNALR string = "ASP_DOTNET_SIGNALR"

	// ScopeConditionsServiceTechnologyAWSLAMBDA captures enum value "AWS_LAMBDA"
	ScopeConditionsServiceTechnologyAWSLAMBDA string = "AWS_LAMBDA"

	// ScopeConditionsServiceTechnologyAWSRDS captures enum value "AWS_RDS"
	ScopeConditionsServiceTechnologyAWSRDS string = "AWS_RDS"

	// ScopeConditionsServiceTechnologyAWSSERVICE captures enum value "AWS_SERVICE"
	ScopeConditionsServiceTechnologyAWSSERVICE string = "AWS_SERVICE"

	// ScopeConditionsServiceTechnologyAXIS captures enum value "AXIS"
	ScopeConditionsServiceTechnologyAXIS string = "AXIS"

	// ScopeConditionsServiceTechnologyAZUREFUNCTIONS captures enum value "AZURE_FUNCTIONS"
	ScopeConditionsServiceTechnologyAZUREFUNCTIONS string = "AZURE_FUNCTIONS"

	// ScopeConditionsServiceTechnologyAZURESERVICEBUS captures enum value "AZURE_SERVICE_BUS"
	ScopeConditionsServiceTechnologyAZURESERVICEBUS string = "AZURE_SERVICE_BUS"

	// ScopeConditionsServiceTechnologyAZURESERVICEFABRIC captures enum value "AZURE_SERVICE_FABRIC"
	ScopeConditionsServiceTechnologyAZURESERVICEFABRIC string = "AZURE_SERVICE_FABRIC"

	// ScopeConditionsServiceTechnologyAZURESTORAGE captures enum value "AZURE_STORAGE"
	ScopeConditionsServiceTechnologyAZURESTORAGE string = "AZURE_STORAGE"

	// ScopeConditionsServiceTechnologyBOSHBPM captures enum value "BOSHBPM"
	ScopeConditionsServiceTechnologyBOSHBPM string = "BOSHBPM"

	// ScopeConditionsServiceTechnologyCITRIX captures enum value "CITRIX"
	ScopeConditionsServiceTechnologyCITRIX string = "CITRIX"

	// ScopeConditionsServiceTechnologyCLOUDFOUNDRY captures enum value "CLOUDFOUNDRY"
	ScopeConditionsServiceTechnologyCLOUDFOUNDRY string = "CLOUDFOUNDRY"

	// ScopeConditionsServiceTechnologyCLOUDFOUNDRYAUCTIONEER captures enum value "CLOUDFOUNDRY_AUCTIONEER"
	ScopeConditionsServiceTechnologyCLOUDFOUNDRYAUCTIONEER string = "CLOUDFOUNDRY_AUCTIONEER"

	// ScopeConditionsServiceTechnologyCLOUDFOUNDRYBOSH captures enum value "CLOUDFOUNDRY_BOSH"
	ScopeConditionsServiceTechnologyCLOUDFOUNDRYBOSH string = "CLOUDFOUNDRY_BOSH"

	// ScopeConditionsServiceTechnologyCLOUDFOUNDRYGOROUTER captures enum value "CLOUDFOUNDRY_GOROUTER"
	ScopeConditionsServiceTechnologyCLOUDFOUNDRYGOROUTER string = "CLOUDFOUNDRY_GOROUTER"

	// ScopeConditionsServiceTechnologyCOLDFUSION captures enum value "COLDFUSION"
	ScopeConditionsServiceTechnologyCOLDFUSION string = "COLDFUSION"

	// ScopeConditionsServiceTechnologyCONTAINERD captures enum value "CONTAINERD"
	ScopeConditionsServiceTechnologyCONTAINERD string = "CONTAINERD"

	// ScopeConditionsServiceTechnologyCOUCHBASE captures enum value "COUCHBASE"
	ScopeConditionsServiceTechnologyCOUCHBASE string = "COUCHBASE"

	// ScopeConditionsServiceTechnologyCRIO captures enum value "CRIO"
	ScopeConditionsServiceTechnologyCRIO string = "CRIO"

	// ScopeConditionsServiceTechnologyCXF captures enum value "CXF"
	ScopeConditionsServiceTechnologyCXF string = "CXF"

	// ScopeConditionsServiceTechnologyDATASTAX captures enum value "DATASTAX"
	ScopeConditionsServiceTechnologyDATASTAX string = "DATASTAX"

	// ScopeConditionsServiceTechnologyDB2 captures enum value "DB2"
	ScopeConditionsServiceTechnologyDB2 string = "DB2"

	// ScopeConditionsServiceTechnologyDIEGOCELL captures enum value "DIEGO_CELL"
	ScopeConditionsServiceTechnologyDIEGOCELL string = "DIEGO_CELL"

	// ScopeConditionsServiceTechnologyDOCKER captures enum value "DOCKER"
	ScopeConditionsServiceTechnologyDOCKER string = "DOCKER"

	// ScopeConditionsServiceTechnologyDOTNET captures enum value "DOTNET"
	ScopeConditionsServiceTechnologyDOTNET string = "DOTNET"

	// ScopeConditionsServiceTechnologyDOTNETREMOTING captures enum value "DOTNET_REMOTING"
	ScopeConditionsServiceTechnologyDOTNETREMOTING string = "DOTNET_REMOTING"

	// ScopeConditionsServiceTechnologyELASTICSEARCH captures enum value "ELASTIC_SEARCH"
	ScopeConditionsServiceTechnologyELASTICSEARCH string = "ELASTIC_SEARCH"

	// ScopeConditionsServiceTechnologyENVOY captures enum value "ENVOY"
	ScopeConditionsServiceTechnologyENVOY string = "ENVOY"

	// ScopeConditionsServiceTechnologyERLANG captures enum value "ERLANG"
	ScopeConditionsServiceTechnologyERLANG string = "ERLANG"

	// ScopeConditionsServiceTechnologyETCD captures enum value "ETCD"
	ScopeConditionsServiceTechnologyETCD string = "ETCD"

	// ScopeConditionsServiceTechnologyF5LTM captures enum value "F5_LTM"
	ScopeConditionsServiceTechnologyF5LTM string = "F5_LTM"

	// ScopeConditionsServiceTechnologyGARDEN captures enum value "GARDEN"
	ScopeConditionsServiceTechnologyGARDEN string = "GARDEN"

	// ScopeConditionsServiceTechnologyGLASSFISH captures enum value "GLASSFISH"
	ScopeConditionsServiceTechnologyGLASSFISH string = "GLASSFISH"

	// ScopeConditionsServiceTechnologyGO captures enum value "GO"
	ScopeConditionsServiceTechnologyGO string = "GO"

	// ScopeConditionsServiceTechnologyGRPC captures enum value "GRPC"
	ScopeConditionsServiceTechnologyGRPC string = "GRPC"

	// ScopeConditionsServiceTechnologyGRSECURITY captures enum value "GRSECURITY"
	ScopeConditionsServiceTechnologyGRSECURITY string = "GRSECURITY"

	// ScopeConditionsServiceTechnologyHADOOP captures enum value "HADOOP"
	ScopeConditionsServiceTechnologyHADOOP string = "HADOOP"

	// ScopeConditionsServiceTechnologyHADOOPHDFS captures enum value "HADOOP_HDFS"
	ScopeConditionsServiceTechnologyHADOOPHDFS string = "HADOOP_HDFS"

	// ScopeConditionsServiceTechnologyHADOOPYARN captures enum value "HADOOP_YARN"
	ScopeConditionsServiceTechnologyHADOOPYARN string = "HADOOP_YARN"

	// ScopeConditionsServiceTechnologyHAPROXY captures enum value "HAPROXY"
	ScopeConditionsServiceTechnologyHAPROXY string = "HAPROXY"

	// ScopeConditionsServiceTechnologyHEAT captures enum value "HEAT"
	ScopeConditionsServiceTechnologyHEAT string = "HEAT"

	// ScopeConditionsServiceTechnologyHESSIAN captures enum value "HESSIAN"
	ScopeConditionsServiceTechnologyHESSIAN string = "HESSIAN"

	// ScopeConditionsServiceTechnologyHORNETQ captures enum value "HORNET_Q"
	ScopeConditionsServiceTechnologyHORNETQ string = "HORNET_Q"

	// ScopeConditionsServiceTechnologyIBMCICSREGION captures enum value "IBM_CICS_REGION"
	ScopeConditionsServiceTechnologyIBMCICSREGION string = "IBM_CICS_REGION"

	// ScopeConditionsServiceTechnologyIBMCICSTRANSACTIONGATEWAY captures enum value "IBM_CICS_TRANSACTION_GATEWAY"
	ScopeConditionsServiceTechnologyIBMCICSTRANSACTIONGATEWAY string = "IBM_CICS_TRANSACTION_GATEWAY"

	// ScopeConditionsServiceTechnologyIBMIMSCONNECTREGION captures enum value "IBM_IMS_CONNECT_REGION"
	ScopeConditionsServiceTechnologyIBMIMSCONNECTREGION string = "IBM_IMS_CONNECT_REGION"

	// ScopeConditionsServiceTechnologyIBMIMSCONTROLREGION captures enum value "IBM_IMS_CONTROL_REGION"
	ScopeConditionsServiceTechnologyIBMIMSCONTROLREGION string = "IBM_IMS_CONTROL_REGION"

	// ScopeConditionsServiceTechnologyIBMIMSMESSAGEPROCESSINGREGION captures enum value "IBM_IMS_MESSAGE_PROCESSING_REGION"
	ScopeConditionsServiceTechnologyIBMIMSMESSAGEPROCESSINGREGION string = "IBM_IMS_MESSAGE_PROCESSING_REGION"

	// ScopeConditionsServiceTechnologyIBMIMSSOAPGATEWAY captures enum value "IBM_IMS_SOAP_GATEWAY"
	ScopeConditionsServiceTechnologyIBMIMSSOAPGATEWAY string = "IBM_IMS_SOAP_GATEWAY"

	// ScopeConditionsServiceTechnologyIBMINTEGRATIONBUS captures enum value "IBM_INTEGRATION_BUS"
	ScopeConditionsServiceTechnologyIBMINTEGRATIONBUS string = "IBM_INTEGRATION_BUS"

	// ScopeConditionsServiceTechnologyIBMMQ captures enum value "IBM_MQ"
	ScopeConditionsServiceTechnologyIBMMQ string = "IBM_MQ"

	// ScopeConditionsServiceTechnologyIBMMQCLIENT captures enum value "IBM_MQ_CLIENT"
	ScopeConditionsServiceTechnologyIBMMQCLIENT string = "IBM_MQ_CLIENT"

	// ScopeConditionsServiceTechnologyIBMWEBSHPREREAPPLICATIONSERVER captures enum value "IBM_WEBSHPRERE_APPLICATION_SERVER"
	ScopeConditionsServiceTechnologyIBMWEBSHPREREAPPLICATIONSERVER string = "IBM_WEBSHPRERE_APPLICATION_SERVER"

	// ScopeConditionsServiceTechnologyIBMWEBSHPRERELIBERTY captures enum value "IBM_WEBSHPRERE_LIBERTY"
	ScopeConditionsServiceTechnologyIBMWEBSHPRERELIBERTY string = "IBM_WEBSHPRERE_LIBERTY"

	// ScopeConditionsServiceTechnologyIIS captures enum value "IIS"
	ScopeConditionsServiceTechnologyIIS string = "IIS"

	// ScopeConditionsServiceTechnologyIISAPPPOOL captures enum value "IIS_APP_POOL"
	ScopeConditionsServiceTechnologyIISAPPPOOL string = "IIS_APP_POOL"

	// ScopeConditionsServiceTechnologyISTIO captures enum value "ISTIO"
	ScopeConditionsServiceTechnologyISTIO string = "ISTIO"

	// ScopeConditionsServiceTechnologyJAVA captures enum value "JAVA"
	ScopeConditionsServiceTechnologyJAVA string = "JAVA"

	// ScopeConditionsServiceTechnologyJAXWS captures enum value "JAX_WS"
	ScopeConditionsServiceTechnologyJAXWS string = "JAX_WS"

	// ScopeConditionsServiceTechnologyJBOSS captures enum value "JBOSS"
	ScopeConditionsServiceTechnologyJBOSS string = "JBOSS"

	// ScopeConditionsServiceTechnologyJBOSSEAP captures enum value "JBOSS_EAP"
	ScopeConditionsServiceTechnologyJBOSSEAP string = "JBOSS_EAP"

	// ScopeConditionsServiceTechnologyJDKHTTPSERVER captures enum value "JDK_HTTP_SERVER"
	ScopeConditionsServiceTechnologyJDKHTTPSERVER string = "JDK_HTTP_SERVER"

	// ScopeConditionsServiceTechnologyJERSEY captures enum value "JERSEY"
	ScopeConditionsServiceTechnologyJERSEY string = "JERSEY"

	// ScopeConditionsServiceTechnologyJETTY captures enum value "JETTY"
	ScopeConditionsServiceTechnologyJETTY string = "JETTY"

	// ScopeConditionsServiceTechnologyJRUBY captures enum value "JRUBY"
	ScopeConditionsServiceTechnologyJRUBY string = "JRUBY"

	// ScopeConditionsServiceTechnologyJYTHON captures enum value "JYTHON"
	ScopeConditionsServiceTechnologyJYTHON string = "JYTHON"

	// ScopeConditionsServiceTechnologyKUBERNETES captures enum value "KUBERNETES"
	ScopeConditionsServiceTechnologyKUBERNETES string = "KUBERNETES"

	// ScopeConditionsServiceTechnologyLIBVIRT captures enum value "LIBVIRT"
	ScopeConditionsServiceTechnologyLIBVIRT string = "LIBVIRT"

	// ScopeConditionsServiceTechnologyLINKERD captures enum value "LINKERD"
	ScopeConditionsServiceTechnologyLINKERD string = "LINKERD"

	// ScopeConditionsServiceTechnologyMARIADB captures enum value "MARIADB"
	ScopeConditionsServiceTechnologyMARIADB string = "MARIADB"

	// ScopeConditionsServiceTechnologyMEMCACHED captures enum value "MEMCACHED"
	ScopeConditionsServiceTechnologyMEMCACHED string = "MEMCACHED"

	// ScopeConditionsServiceTechnologyMICROSOFTSQLSERVER captures enum value "MICROSOFT_SQL_SERVER"
	ScopeConditionsServiceTechnologyMICROSOFTSQLSERVER string = "MICROSOFT_SQL_SERVER"

	// ScopeConditionsServiceTechnologyMONGODB captures enum value "MONGODB"
	ScopeConditionsServiceTechnologyMONGODB string = "MONGODB"

	// ScopeConditionsServiceTechnologyMSSQLCLIENT captures enum value "MSSQL_CLIENT"
	ScopeConditionsServiceTechnologyMSSQLCLIENT string = "MSSQL_CLIENT"

	// ScopeConditionsServiceTechnologyMULEESB captures enum value "MULE_ESB"
	ScopeConditionsServiceTechnologyMULEESB string = "MULE_ESB"

	// ScopeConditionsServiceTechnologyMYSQL captures enum value "MYSQL"
	ScopeConditionsServiceTechnologyMYSQL string = "MYSQL"

	// ScopeConditionsServiceTechnologyMYSQLCONNECTOR captures enum value "MYSQL_CONNECTOR"
	ScopeConditionsServiceTechnologyMYSQLCONNECTOR string = "MYSQL_CONNECTOR"

	// ScopeConditionsServiceTechnologyNETFLIXSERVO captures enum value "NETFLIX_SERVO"
	ScopeConditionsServiceTechnologyNETFLIXSERVO string = "NETFLIX_SERVO"

	// ScopeConditionsServiceTechnologyNETTY captures enum value "NETTY"
	ScopeConditionsServiceTechnologyNETTY string = "NETTY"

	// ScopeConditionsServiceTechnologyNGINX captures enum value "NGINX"
	ScopeConditionsServiceTechnologyNGINX string = "NGINX"

	// ScopeConditionsServiceTechnologyNODEJS captures enum value "NODE_JS"
	ScopeConditionsServiceTechnologyNODEJS string = "NODE_JS"

	// ScopeConditionsServiceTechnologyONEAGENTSDK captures enum value "ONEAGENT_SDK"
	ScopeConditionsServiceTechnologyONEAGENTSDK string = "ONEAGENT_SDK"

	// ScopeConditionsServiceTechnologyOPENCENSUS captures enum value "OPENCENSUS"
	ScopeConditionsServiceTechnologyOPENCENSUS string = "OPENCENSUS"

	// ScopeConditionsServiceTechnologyOPENSHIFT captures enum value "OPENSHIFT"
	ScopeConditionsServiceTechnologyOPENSHIFT string = "OPENSHIFT"

	// ScopeConditionsServiceTechnologyOPENSTACKCOMPUTE captures enum value "OPENSTACK_COMPUTE"
	ScopeConditionsServiceTechnologyOPENSTACKCOMPUTE string = "OPENSTACK_COMPUTE"

	// ScopeConditionsServiceTechnologyOPENSTACKCONTROLLER captures enum value "OPENSTACK_CONTROLLER"
	ScopeConditionsServiceTechnologyOPENSTACKCONTROLLER string = "OPENSTACK_CONTROLLER"

	// ScopeConditionsServiceTechnologyOPENTELEMETRY captures enum value "OPENTELEMETRY"
	ScopeConditionsServiceTechnologyOPENTELEMETRY string = "OPENTELEMETRY"

	// ScopeConditionsServiceTechnologyOPENTRACING captures enum value "OPENTRACING"
	ScopeConditionsServiceTechnologyOPENTRACING string = "OPENTRACING"

	// ScopeConditionsServiceTechnologyOPENLIBERTY captures enum value "OPEN_LIBERTY"
	ScopeConditionsServiceTechnologyOPENLIBERTY string = "OPEN_LIBERTY"

	// ScopeConditionsServiceTechnologyORACLEDATABASE captures enum value "ORACLE_DATABASE"
	ScopeConditionsServiceTechnologyORACLEDATABASE string = "ORACLE_DATABASE"

	// ScopeConditionsServiceTechnologyORACLEWEBLOGIC captures enum value "ORACLE_WEBLOGIC"
	ScopeConditionsServiceTechnologyORACLEWEBLOGIC string = "ORACLE_WEBLOGIC"

	// ScopeConditionsServiceTechnologyOWIN captures enum value "OWIN"
	ScopeConditionsServiceTechnologyOWIN string = "OWIN"

	// ScopeConditionsServiceTechnologyPERL captures enum value "PERL"
	ScopeConditionsServiceTechnologyPERL string = "PERL"

	// ScopeConditionsServiceTechnologyPHP captures enum value "PHP"
	ScopeConditionsServiceTechnologyPHP string = "PHP"

	// ScopeConditionsServiceTechnologyPHPFPM captures enum value "PHP_FPM"
	ScopeConditionsServiceTechnologyPHPFPM string = "PHP_FPM"

	// ScopeConditionsServiceTechnologyPLAY captures enum value "PLAY"
	ScopeConditionsServiceTechnologyPLAY string = "PLAY"

	// ScopeConditionsServiceTechnologyPOSTGRESQL captures enum value "POSTGRE_SQL"
	ScopeConditionsServiceTechnologyPOSTGRESQL string = "POSTGRE_SQL"

	// ScopeConditionsServiceTechnologyPOSTGRESQLDOTNETDATAPROVIDER captures enum value "POSTGRE_SQL_DOTNET_DATA_PROVIDER"
	ScopeConditionsServiceTechnologyPOSTGRESQLDOTNETDATAPROVIDER string = "POSTGRE_SQL_DOTNET_DATA_PROVIDER"

	// ScopeConditionsServiceTechnologyPROGRESS captures enum value "PROGRESS"
	ScopeConditionsServiceTechnologyPROGRESS string = "PROGRESS"

	// ScopeConditionsServiceTechnologyPYTHON captures enum value "PYTHON"
	ScopeConditionsServiceTechnologyPYTHON string = "PYTHON"

	// ScopeConditionsServiceTechnologyRABBITMQ captures enum value "RABBIT_MQ"
	ScopeConditionsServiceTechnologyRABBITMQ string = "RABBIT_MQ"

	// ScopeConditionsServiceTechnologyREDIS captures enum value "REDIS"
	ScopeConditionsServiceTechnologyREDIS string = "REDIS"

	// ScopeConditionsServiceTechnologyRESTEASY captures enum value "RESTEASY"
	ScopeConditionsServiceTechnologyRESTEASY string = "RESTEASY"

	// ScopeConditionsServiceTechnologyRESTLET captures enum value "RESTLET"
	ScopeConditionsServiceTechnologyRESTLET string = "RESTLET"

	// ScopeConditionsServiceTechnologyRIAK captures enum value "RIAK"
	ScopeConditionsServiceTechnologyRIAK string = "RIAK"

	// ScopeConditionsServiceTechnologyRUBY captures enum value "RUBY"
	ScopeConditionsServiceTechnologyRUBY string = "RUBY"

	// ScopeConditionsServiceTechnologySAGWEBMETHODSIS captures enum value "SAG_WEBMETHODS_IS"
	ScopeConditionsServiceTechnologySAGWEBMETHODSIS string = "SAG_WEBMETHODS_IS"

	// ScopeConditionsServiceTechnologySAP captures enum value "SAP"
	ScopeConditionsServiceTechnologySAP string = "SAP"

	// ScopeConditionsServiceTechnologySAPHANADB captures enum value "SAP_HANADB"
	ScopeConditionsServiceTechnologySAPHANADB string = "SAP_HANADB"

	// ScopeConditionsServiceTechnologySAPHYBRIS captures enum value "SAP_HYBRIS"
	ScopeConditionsServiceTechnologySAPHYBRIS string = "SAP_HYBRIS"

	// ScopeConditionsServiceTechnologySAPMAXDB captures enum value "SAP_MAXDB"
	ScopeConditionsServiceTechnologySAPMAXDB string = "SAP_MAXDB"

	// ScopeConditionsServiceTechnologySAPSYBASE captures enum value "SAP_SYBASE"
	ScopeConditionsServiceTechnologySAPSYBASE string = "SAP_SYBASE"

	// ScopeConditionsServiceTechnologySCALA captures enum value "SCALA"
	ScopeConditionsServiceTechnologySCALA string = "SCALA"

	// ScopeConditionsServiceTechnologySELINUX captures enum value "SELINUX"
	ScopeConditionsServiceTechnologySELINUX string = "SELINUX"

	// ScopeConditionsServiceTechnologySHAREPOINT captures enum value "SHAREPOINT"
	ScopeConditionsServiceTechnologySHAREPOINT string = "SHAREPOINT"

	// ScopeConditionsServiceTechnologySPARK captures enum value "SPARK"
	ScopeConditionsServiceTechnologySPARK string = "SPARK"

	// ScopeConditionsServiceTechnologySPRING captures enum value "SPRING"
	ScopeConditionsServiceTechnologySPRING string = "SPRING"

	// ScopeConditionsServiceTechnologySQLITE captures enum value "SQLITE"
	ScopeConditionsServiceTechnologySQLITE string = "SQLITE"

	// ScopeConditionsServiceTechnologyTHRIFT captures enum value "THRIFT"
	ScopeConditionsServiceTechnologyTHRIFT string = "THRIFT"

	// ScopeConditionsServiceTechnologyTIBCO captures enum value "TIBCO"
	ScopeConditionsServiceTechnologyTIBCO string = "TIBCO"

	// ScopeConditionsServiceTechnologyTIBCOBUSINESSWORKS captures enum value "TIBCO_BUSINESS_WORKS"
	ScopeConditionsServiceTechnologyTIBCOBUSINESSWORKS string = "TIBCO_BUSINESS_WORKS"

	// ScopeConditionsServiceTechnologyTIBCOEMS captures enum value "TIBCO_EMS"
	ScopeConditionsServiceTechnologyTIBCOEMS string = "TIBCO_EMS"

	// ScopeConditionsServiceTechnologyVARNISHCACHE captures enum value "VARNISH_CACHE"
	ScopeConditionsServiceTechnologyVARNISHCACHE string = "VARNISH_CACHE"

	// ScopeConditionsServiceTechnologyVIM2 captures enum value "VIM2"
	ScopeConditionsServiceTechnologyVIM2 string = "VIM2"

	// ScopeConditionsServiceTechnologyVIRTUALMACHINEKVM captures enum value "VIRTUAL_MACHINE_KVM"
	ScopeConditionsServiceTechnologyVIRTUALMACHINEKVM string = "VIRTUAL_MACHINE_KVM"

	// ScopeConditionsServiceTechnologyVIRTUALMACHINEQEMU captures enum value "VIRTUAL_MACHINE_QEMU"
	ScopeConditionsServiceTechnologyVIRTUALMACHINEQEMU string = "VIRTUAL_MACHINE_QEMU"

	// ScopeConditionsServiceTechnologyWILDFLY captures enum value "WILDFLY"
	ScopeConditionsServiceTechnologyWILDFLY string = "WILDFLY"

	// ScopeConditionsServiceTechnologyWINDOWSCONTAINERS captures enum value "WINDOWS_CONTAINERS"
	ScopeConditionsServiceTechnologyWINDOWSCONTAINERS string = "WINDOWS_CONTAINERS"

	// ScopeConditionsServiceTechnologyWINK captures enum value "WINK"
	ScopeConditionsServiceTechnologyWINK string = "WINK"

	// ScopeConditionsServiceTechnologyZEROMQ captures enum value "ZERO_MQ"
	ScopeConditionsServiceTechnologyZEROMQ string = "ZERO_MQ"
)

// prop value enum
func (m *ScopeConditions) validateServiceTechnologyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scopeConditionsTypeServiceTechnologyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScopeConditions) validateServiceTechnology(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceTechnology) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceTechnologyEnum("serviceTechnology", "body", m.ServiceTechnology); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScopeConditions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScopeConditions) UnmarshalBinary(b []byte) error {
	var res ScopeConditions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
