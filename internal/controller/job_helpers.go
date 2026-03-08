/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"github.com/osac-project/osac-operator/api/v1alpha1"
)

// findJobByID finds a job by its ID in the jobs array.
// Returns a pointer to the job if found, nil otherwise.
func findJobByID(jobs []v1alpha1.JobStatus, jobID string) *v1alpha1.JobStatus {
	for i := range jobs {
		if jobs[i].JobID == jobID {
			return &jobs[i]
		}
	}
	return nil
}

// updateJob updates an existing job by ID with new values.
// Returns true if the job was found and updated, false otherwise.
func updateJob(jobs []v1alpha1.JobStatus, updatedJob v1alpha1.JobStatus) bool {
	job := findJobByID(jobs, updatedJob.JobID)
	if job == nil {
		return false
	}
	*job = updatedJob
	return true
}
