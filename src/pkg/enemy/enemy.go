package enemy

type EnemyRank int

const (
	EnemyRank_INVALID EnemyRank = -1
	EnemyRank_COMMON  EnemyRank = 0
	EnemyRank_ELITE   EnemyRank = 1
	EnemyRank_BOSS    EnemyRank = 2
)

type EnemyConfig struct {
	Rank EnemyRank
}
