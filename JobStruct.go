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

// id -> A job has to be uniquely identified
// status -> displays which status a job currently is in, this helps to remove states from application and store those only inside of the data
// image -> image of executing job
// execution_args -> args that should be passed to docker run [args]
// environment_vars -> Environment vars that should be set for the container --> possibly redundand with execution args
// logs_path -> Path on a persistent volume that contains a logfile for the complete job
// artifacts_path -> Artifacts that are returned by the job, need to specify more like: how does the application know what is an artifact?, how is the artifact extracted? !!! This is a question that should be discussed !!!
// scheduled_worker -> worker that should execute / has executed the Job
// exitcode -> exitcode of the container

// Maybe add a 2nd type of artifacts path -> one is where on the container fs the artifacts will be found, the other where to save it
