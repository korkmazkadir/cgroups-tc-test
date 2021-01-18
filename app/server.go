package app

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Server implements a TCP file server.
type Server struct {
	pid int

	// protects following fields
	rwLock sync.RWMutex

	availableFileMap map[string]File
}

var (
	// ErrFileNotAvailable is raised if the file is not registered in the available file list
	ErrFileNotAvailable = fmt.Errorf("file is not available")
	// ErrInvalidFileName is raised if the file name is not valid one
	ErrInvalidFileName = fmt.Errorf("file name is invalid")
	// ErrFileAlreadyRegistered is raised when an already used name is used to register a file
	ErrFileAlreadyRegistered = fmt.Errorf("file already registered")
)

// NewServer creates a new server and returns it
func NewServer() *Server {
	server := new(Server)
	server.pid = os.Getpid()
	server.availableFileMap = make(map[string]File)
	return server
}

// GetFile call accepts file name as a parameter and returns the file
func (s *Server) GetFile(fileName string, file *File) error {

	if fileName == "" {
		return ErrInvalidFileName
	}

	s.rwLock.RLock()
	localFile, ok := s.availableFileMap[fileName]
	s.rwLock.RUnlock()

	if ok == false {
		return ErrFileNotAvailable
	}

	file.Data = localFile.Data
	log.Printf("[%d]\t[GET] Name %s Size %d Hash: %s\n", s.pid, fileName, len(file.Data), file.Hash())

	return nil
}

// PutFile puts a file with a specific name
func (s *Server) PutFile(file File, reply *int) error {

	if file.Name == "" {
		return ErrInvalidFileName
	}

	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	_, ok := s.availableFileMap[file.Name]

	if ok {
		return ErrFileAlreadyRegistered
	}

	s.availableFileMap[file.Name] = file

	log.Printf("[%d]\t[PUT] Name %s Size %d Hash: %s\n", s.pid, file.Name, len(file.Data), file.Hash())

	return nil
}
