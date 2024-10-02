package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct{
	gorm.Model
	Id int `gorm:"primaryKey"`
	Url string `gorm:"not null"`
	Used int 
	Alive bool `gorm:"default:true"`
}
func (s *Server) IncreaseUsed() int {
	s.Used = s.Used + 1
	return s.Used
	
}
func (s *Server) Find(servers *[]Server) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Find(servers).Error
}
func (s *Server) Dead() bool {
	s.Alive = false
	return s.Alive
}
func (s *Server) get_the_next() *Server {
	var servers []Server
	result := s.Find(&servers)
	if result != nil || len(servers) == 0 {
		return nil
	}

	var leastUsedServer *Server
	least := servers[0].Used
	leastUsedServer = &servers[0]

	for _, server := range servers[1:] { 
		if server.Used < least {
			least = server.Used
			leastUsedServer = &server 
		}
		
		
	

	return leastUsedServer
}
