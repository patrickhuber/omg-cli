package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type FailedJobs struct {

	/*CutoffAgeInDays - Descr: How old a failed job should stay in cloud controller database before being cleaned up Default: 31
*/
	CutoffAgeInDays interface{} `yaml:"cutoff_age_in_days,omitempty"`

}