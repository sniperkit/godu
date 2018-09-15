/*
Sniperkit-Bot
- Status: analyzed
*/

package core

import (
	"fmt"
	"sync"
)

func PrepareTree(tree *File, limit int64) error {
	PruneTree(tree, limit)
	if len(tree.Files) == 0 {
		return fmt.Errorf("the folder '%s' doesn't contain any files bigger than %dMB", tree.Name, limit/MEGABYTE)
	}
	SortDesc(tree)
	return nil
}

func StartProcessing(
	folder *File,
	commands <-chan Executer,
	states chan<- State,
	lastStateChan chan<- *State,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	state := State{
		Folder:      folder,
		MarkedFiles: make(map[*File]struct{}),
	}
	states <- state
	for {
		command, more := <-commands
		if !more {
			close(states)
			break
		}
		if newState, err := command.Execute(state); err == nil {
			state = newState
			states <- state
		}
	}
	lastStateChan <- &state
}
