package pkg

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
	v1 "k8s.io/api/core/v1"
	metav "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodes(l string) []v1.Node {
	client := getClient()
	nodeClient := client.CoreV1().Nodes()
	var options metav.ListOptions
	if l != "" {
		options = metav.ListOptions{LabelSelector: l}
	}
	nodes, err := nodeClient.List(context.TODO(), options)
	if err != nil {
		log.Fatalln("Could not get nodes!")
		os.Exit(2)
	}
	return nodes.Items
}

func ListNodes(nodes []v1.Node) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Available", "Capacity", "Used%"})

	rows := make([]table.Row, len(nodes))
	for i, node := range nodes {
		name := node.ObjectMeta.Name
		capacity := node.Status.Capacity.StorageEphemeral().AsApproximateFloat64() / 1024 / 1024 / 1024
		available := node.Status.Allocatable.StorageEphemeral().AsApproximateFloat64() / 1024 / 1024 / 1024
		used := (capacity - available) / capacity * 100
		rows[i] = table.Row{
			name,
			fmt.Sprintf("%.2fGi", available),
			fmt.Sprintf("%.2fGi", capacity),
			fmt.Sprintf("%.2f", used),
		}
	}
	t.AppendRows(rows)
	style := table.StyleBold
	style.Options = table.OptionsNoBordersAndSeparators
	t.SetStyle(style)
	t.Render()
}
