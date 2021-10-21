package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gen-k8s-user/auth"
	"gen-k8s-user/utils"

	"encoding/json"

	"github.com/spf13/cobra"
)

type CommandOptions struct {
	UserName   string
	NameSpaces string
	Clusters   string
	Bindings   string
}

func (o *CommandOptions) Validate() (Binder, error) {
	b := Binder{}
	b.User = strings.Split(o.UserName, ",")
	b.Namespaces = strings.Split(o.NameSpaces, ",")
	b.Clusters = strings.Split(o.Clusters, ",")
	b.Roles = strings.Split(o.Bindings, ",")
	return b, nil
}

type Binder struct {
	User       []string
	Namespaces []string
	Clusters   []string
	Roles      []string
}

func (b *Binder) GenUser() {
	for _, u := range b.User {
		for _, n := range b.Namespaces {
			for _, c := range b.Clusters {
				b.gen(u, c, n)
			}
		}
	}
}

func (b *Binder) gen(name, cluster, ns string) {
	k, err := utils.NewKubeClient()
	if err != nil {
		log.Fatalln(err)
	}
	kclient, err := k.GetClientSet(cluster)
	if err != nil {
		log.Fatalln(err)
	}
	client := auth.NewK8sClient(kclient)
	sa, err := client.CreateAccount(name, ns)
	if err != nil {
		log.Fatalln(err)
	}
	for _, r := range b.Roles {
		err = client.BindRole(r, ns, sa)
		if err != nil {
			log.Fatalln(err)
		}
	}
	time.Sleep(time.Duration(2) * time.Second)
	kubeconfig, err := client.GenKubeConfig(sa)
	if err != nil {
		log.Fatalln(err)
	}
	d, err := json.Marshal(&kubeconfig)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(d))
}

// 命令行
func main() {
	c := CommandOptions{}
	cmd := &cobra.Command{
		Use:   "gen-user NAME [flags]",
		Short: "generator k8s user and output his kubeconfig",
		Long: `To generate the account of k8s and output kubeconfig. Command like:
		
		gen-user <User-Name> --namespace <ns1>,<ns2> --binding <role1>,<role2> --cluster <cluster-context1>,<cluster-context2>

		for example:
		
		gen-user shikanon --namespace default,rcmd --binding rcmd-dev --cluster new-test
		`,
		Run: func(cmd *cobra.Command, args []string) {
			c.UserName = args[0]
			b, err := c.Validate()
			if err != nil {
				log.Fatal(err)
			}
			b.GenUser()
		},
	}
	cmd.Flags().StringVarP(&c.NameSpaces, "namespace", "n", c.NameSpaces, "The user created by which namespaces")
	cmd.Flags().StringVarP(&c.Clusters, "cluster", "c", c.Clusters, "The user created by cluter which is current kubeconfig context name")
	cmd.Flags().StringVarP(&c.Bindings, "binding", "b", c.Bindings, "The user binding by cluter role")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
