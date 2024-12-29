package model

type TargetType int

const (
	// Placeholder/sentinel value, in case some fucky stuff happens
	TARGET_NULL    TargetType = 0
	TARGET_SELF    TargetType = 1
	TARGET_ALLIES  TargetType = 2
	TARGET_ENEMIES TargetType = 3
)

// Roles in this game: Used by wengines and characters

type Role int

const (
	Role_INVALID Role = -1
	Role_ATTACK  Role = 0
	Role_ANOMALY Role = 1
	Role_SUPPORT Role = 2
	Role_DEFENSE Role = 3
	Role_STUN    Role = 4
)
