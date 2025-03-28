package main

import (
	"flag"
	"fmt"
	"github.com/mbausson/handshake/pkg/api"
	"github.com/mbausson/handshake/pkg/graph"
	"sync"
)

var cache sync.Map

func main() {
	var startId, targetId string
	var maxDepth int

	flag.StringVar(&api.Key, "key", "", "Your steam API key")
	flag.StringVar(&startId, "from", "", "The steam user's ID from which the handshakes start")
	flag.StringVar(&targetId, "to", "", "The steam user's ID to find")
	flag.IntVar(&maxDepth, "depth", 6, "The maximum depth of relations until the algorithm stops")

	flag.Parse()

	root := &graph.Node{
		startId,
		nil,
		nil,
	}

	currentNodes := []*graph.Node{root}

	for depth := 0; depth < maxDepth; depth++ {
		fmt.Printf("%v - %v total friends\n", depth+1, len(currentNodes))

		if nodeFound := graph.FindNode(currentNodes, targetId); nodeFound != nil {
			displayResult(startId, nodeFound)
			return
		}

		currentNodes = fetchFriends(currentNodes)
	}

	fmt.Printf("No friends found ! Reached maximum depth (%v)\n", maxDepth)
}

func fetchFriends(nodes []*graph.Node) []*graph.Node {
	wg := sync.WaitGroup{}
	result := []*graph.Node{}

	for _, node := range nodes {
		wg.Add(1)

		go func(steamId string) {
			defer wg.Done()

			if cachedNode, ok := cache.Load(steamId); ok {
				result = append(result, cachedNode.(*graph.Node))
				return
			}

			var friends, err = api.GetFriends(node.Data)

			if err != nil {
				fmt.Println(fmt.Errorf("error fetching friends: %v", err))
				return
			}

			for _, friend := range friends {
				newNode := &graph.Node{
					Data: friend.SteamId,
					Root: node,
				}

				cache.Store(steamId, newNode)

				result = append(result, newNode)
			}
		}(node.Data)
	}

	wg.Wait()
	return result
}

func displayResult(startingId string, endNode *graph.Node) {
	displayUser := func(steamId string) {
		user, err := api.GetUser(steamId)

		if err != nil {
			panic(err)
		}

		fmt.Printf("--> User '%s' %s\n", user.Name, user.Url)
	}

	currentNode := endNode

	for currentNode.Root != nil {
		displayUser(currentNode.Data)

		currentNode = currentNode.Root
	}

	displayUser(startingId)
}
