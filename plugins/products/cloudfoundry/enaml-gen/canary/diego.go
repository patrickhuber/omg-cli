package canary 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Diego struct {

	/*Canary - Descr: Datadog API key for the canary app Default: <nil>
*/
	Canary *Canary `yaml:"canary,omitempty"`

	/*Ssl - Descr: Toggles cli verification of the Elastic Runtime API SSL certificate Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

}