package collector

import (
	"QoS/config"
	"QoS/tool"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func Refresh(name string, cpu string, mem string) error {
	config := config.Yaml{}
	f, err := ioutil.ReadFile("/home/crrr/k8s/" + name + ".yaml")
	if err != nil {
		fmt.Println("read fail", err)
		return err
	}
	if yaml.Unmarshal(f, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
		return err
	}
	config.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu = cpu
	config.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu = cpu
	config.Spec.Template.Spec.Containers[0].Resources.Limits.Memory = mem
	config.Spec.Template.Spec.Containers[0].Resources.Requests.Memory = mem
	data, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("生成config.yaml出错: %v", err)
		return err
	}
	err = ioutil.WriteFile("/home/crrr/k8s/"+name+".yaml", data, 0777)
	if err != nil {
		log.Fatalf("write config.yaml err: %v", err)
		return err
	}
	err = tool.Exec("kubectl delete rc -n db " + name + "-rc")
	err = tool.Exec("kubectl apply -f /home/crrr/k8s/" + name + ".yaml")
	return nil
}
