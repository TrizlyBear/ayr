package database

type Server struct {
	Id				string
	Config			ServerConfig
	LevelRewards	map[int]string
	Roles			RoleMenu
}

type ServerConfig struct {
	LevelingEnabled	bool
	RoleByMessage	bool
	LevelUpMessage	bool
}

type RoleMenu struct {
	
}