package issue

import (
	"time"

	"github.com/google/uuid"
)

type Issue struct {
	SysId                    uuid.UUID `json:"sysid" maelstrom:"required"`
	Identifier               string    `json:"identifier" maelstrom:"required"`
	SummaryBrief             string    `json:"summary_brief" maelstrom:"required"`
	SummaryLong              string    `json:"summary_long" maelstrom:"required"`
	ProblemDescription       string    `json:"problem_description" maelstrom:"required"`
	WorkAround               string    `json:"work_around" maelstrom:"required"`
	StepsToReproduce         string    `json:"steps_to_reproduce" maelstrom:"required"`
	Kind                     uuid.UUID `json:"kind" maelstrom:"required" maelstrom:"required"`
	DateFound                time.Time `json:"date_found" maelstrom:"required"`
	DateReported             time.Time `json:"date_reported" maelstrom:"required"`
	DateInput                time.Time `json:"date_input" maelstrom:"required"`
	FoundByPrimary           uuid.UUID `json:"found_by_primary" maelstrom:"required"`
	FoundByTeamPrimary       uuid.UUID `json:"found_by_team_primary" maelstrom:"required"`
	ReportedByPrimary        uuid.UUID `json:"reported_by_primary" maelstrom:"required"`
	ReportedByTeamPrimary    uuid.UUID `json:"reported_by_team_primary" maelstrom:"required"`
	InputByPrimary           uuid.UUID `json:"input_by_primary" maelstrom:"required"`
	InputByTeamPrimary       uuid.UUID `json:"input_by_team_primary" maelstrom:"required"`
	Severity                 uuid.UUID `json:"severity" maelstrom:"required"`
	Priority                 uuid.UUID `json:"priority" maelstrom:"required"`
	OrganizationValue        uuid.UUID `json:"organization_value" maelstrom:"required"`
	CurrentStatus            uuid.UUID `json:"current_status" maelstrom:"required"`
	CurrentState             uuid.UUID `json:"current_state" maelstrom:"required"`
	IsResolved               bool      `json:"is_resolved" maelstrom:"required"`
	DateResolved             time.Time `json:"date_resolved" maelstrom:"required"`
	ResolvedByPrimary        uuid.UUID `json:"resolved_by_primary" maelstrom:"required"`
	ResolvedByTeamPrimary    uuid.UUID `json:"resolved_by_team_primary" maelstrom:"required"`
	ResolutionDueDate        time.Time `json:"resolution_due_date" maelstrom:"required"`
	ResolutionEffortUnit     uuid.UUID `json:"resolution_effort_unit" maelstrom:"required"`
	ResolutionEffort         string    `json:"resolution_effort" maelstrom:"required"`
	EstimatedResolutionDate  time.Time `json:"estimated_resolution_date" maelstrom:"required"`
	TargetResolutionDate     time.Time `json:"target_resolution_date" maelstrom:"required"`
	RootCauseAnalysis        string    `json:"root_cause_analysis" maelstrom:"required"`
	FixDescription           string    `json:"fix_description" maelstrom:"required"`
	AssignedToPrimary        uuid.UUID `json:"assigned_to_primary" maelstrom:"required"`
	AssignedToTeamPrimary    uuid.UUID `json:"assigned_to_team_primary" maelstrom:"required"`
	TargetOriginalBuild      uuid.UUID `json:"target_original_build" maelstrom:"required"`
	EstimatedOriginalBuild   uuid.UUID `json:"estimated_original_build" maelstrom:"required"`
	ActualOriginalBuild      uuid.UUID `json:"actual_original_build" maelstrom:"required"`
	TargetOriginalRelease    uuid.UUID `json:"target_original_release" maelstrom:"required"`
	EstimatedOriginalRelease uuid.UUID `json:"estimated_original_release" maelstrom:"required"`
	ActualOriginalRelease    uuid.UUID `json:"actual_original_release" maelstrom:"required"`
}
