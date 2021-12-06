package main
 
import (
    "fmt"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 
    "k8s.io/client-go/kubernetes"
 
    "k8s.io/client-go/tools/clientcmd"
)
 
func main() {
 
    config, err := clientcmd.buildconfigfromflags("", "kube/config")
    if err != nil {
        panic(err)
    }
    client, _ := kubernetes.newforconfig(config)
    pods ,err := client.corev1().pods("ferry").list(metav1.listoptions{})
    if err != nil {
        fmt.println(err)
        return
    }
 
    for _,v := range  pods.items {
        fmt.printf("Namespace is:% v \ n pod Name:% v \ n IP:% v \ n \ n",v.namespace,v.name,v.status.podip)
    }
}
