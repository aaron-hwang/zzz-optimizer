package model

type TargetType int

const (
	// Placeholder/sentinel value, in case some fucky stuff happens
	TARGET_NULL    TargetType = 0
	TARGET_SELF    TargetType = 1
	TARGET_ALLIES  TargetType = 2
	TARGET_ENEMIES TargetType = 3
)
