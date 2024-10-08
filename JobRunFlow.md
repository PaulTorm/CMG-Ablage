# Different Flows for Job Executions
Our diagrams do not include authentication, this will be handled separately.
Most likely the workers will authenticate via JWT, BasicAuth or something else that can be provisioned by Environments.
The internal connections can be handled via similar or different authentications.

## Full execution with job
```
@startuml
participant WorkerRegistry as wr
participant Worker as w
participant WorkerGateway as wg
participant Job as j
w -> wg: getJob
wg -> j: getJobForWorkerX
j --> wg: returnJobForX
wg --> w: returnJob
w -> wr: status: Working
wr --> w: status received
note over w : Execute Job
w -> wg: jobFinished
note left: Contains  Artifacts, Logs and Result
wg -> j: jobXFinished
j --> wg: jobFinishedReceived
wg --> w: jobFinishedReceived
w -> wr: status: readyForWork
wr --> w: status received
@enduml
```

## Full execution no job
```
@startuml
participant WorkerRegistry as wr
participant Worker as w
participant WorkerGateway as wg
participant Job as j
w -> wg: getJob
wg -> j: getJobForWorkerX
j --> wg: noJobForX
wg --> w: noJob
@enduml
```

## Job has an error
```
@startuml
participant WorkerRegistry as wr
participant Worker as w
participant WorkerGateway as wg
participant Job as j
w -> wg: getJob
wg -> j: getJobForWorkerX
j -[#red]-> wg: Error/No connection
wg --> w: noJob
@enduml
```

## WorkerGateway has an error
```
@startuml
participant WorkerRegistry as wr
participant Worker as w
participant WorkerGateway as wg
participant Job as j
w -> wg: getJob
wg -[#red]-> w: Error/No connection
@enduml
```

## Job has Non-Zero Exitcode
```
@startuml
participant WorkerRegistry as wr
participant Worker as w
participant WorkerGateway as wg
participant Job as j
w -> wg: getJob
wg -> j: getJobForWorkerX
j --> wg: returnJobForX
wg --> w: returnJob
w -> wr: status: Working
wr --> w: status received
note over w #red: Execute Job - Error
w -[#red]> wg: jobFinished
note left: Contains  Artifacts, Logs and Result
wg -[#red]> j: jobXFinished
j --> wg: jobFinishedReceived
wg --> w: jobFinishedReceived
w -> wr: status: readyForWork
wr --> w: status received
@enduml
```

Idea: When job has finished with Non-zero exitcode, the job will contain this information inside the returnvalues but not handle it any different. The Job-Scheduler or something else should decide how to continue with this job.


## WorkerRegistry has an Error
```
@startuml
participant WorkerRegistry as wr
participant Worker as w
participant WorkerGateway as wg
participant Job as j
w -> wg: getJob
wg -> j: getJobForWorkerX
j --> wg: returnJobForX
wg --> w: returnJob
w -> wr: status: Working
wr -[#red]-> w: Error/No connection
w -[#red]> wg: refuse job
wg -[#red]> j: refuse job
j --> wg: refuse accepted
wg --> w: refuse accepted
@enduml
```
Idea: There should no job be executed if the Worker Registry is not available / faulty. This will prevent false scheduling