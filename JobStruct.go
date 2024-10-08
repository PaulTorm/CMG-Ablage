package main

type Job struct {
	id               int
	status           int // e.g 0 = new job, 1 = scheduled job, 2 = currently in work, 3 = finished
	image            string
	execution_args   []string
	environment_vars []string
	logs_path        string
	artifacts_path   string
	scheduled_worker string
	exitcode         int
}
