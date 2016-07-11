package consul 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Consul struct {

	/*JoinHosts - Descr: Hostnames/IPs representing all initial consul servers. Use this or consul.join_host / consul.size Default: <nil>
*/
	JoinHosts interface{} `yaml:"join_hosts,omitempty"`

	/*SslKey - Descr: The content of the key file Default: <nil>
*/
	SslKey interface{} `yaml:"ssl_key,omitempty"`

	/*SslCert - Descr: The content of the cert file Default: <nil>
*/
	SslCert interface{} `yaml:"ssl_cert,omitempty"`

	/*Encrypt - Descr: A key to encrypt the traffic between the consul agents (use consul keygen) Default: <nil>
*/
	Encrypt interface{} `yaml:"encrypt,omitempty"`

	/*JoinHost - Descr: Hostname/IP for initial cluster node for other consul servers to join. Default: <nil>
*/
	JoinHost interface{} `yaml:"join_host,omitempty"`

	/*AgentConfig - Descr: override hash for the consul agent.json configuration Default: <nil>
*/
	AgentConfig interface{} `yaml:"agent_config,omitempty"`

	/*Domain - Descr: All DNS queries for this domain will be handled by consul. Default: consul
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*User - Descr: User that consul is ran under Default: vcap
*/
	User interface{} `yaml:"user,omitempty"`

	/*ClientAddr - Descr: The IP to use for client communication Default: <nil>
*/
	ClientAddr interface{} `yaml:"client_addr,omitempty"`

	/*Server - Descr: Should the agent run in server or client mode Default: true
*/
	Server interface{} `yaml:"server,omitempty"`

	/*LeaveOnTerminate - Descr: If enabled, gracefully leave the cluster when the process shuts down. Default: false
*/
	LeaveOnTerminate interface{} `yaml:"leave_on_terminate,omitempty"`

	/*DefaultRecursor - Descr: DNS recursor to use if BOSH not provising DNS Default: 8.8.8.8
*/
	DefaultRecursor interface{} `yaml:"default_recursor,omitempty"`

	/*Services - Descr: a map of service configurations, keyed by service name Default: <nil>
*/
	Services interface{} `yaml:"services,omitempty"`

	/*SslCa - Descr: The content of the ca file Default: <nil>
*/
	SslCa interface{} `yaml:"ssl_ca,omitempty"`

}