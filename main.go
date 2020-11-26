package main

import (
	"log"

	"github.com/keti-openfx/openfx/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

/* Cache 객체 선언 */ 
var Cache CacheMap


func main() {

	/* kubernetes가 pod에 제공하는 config 객체를 반환
	 * kubernetes에서 실행중인 Pod 내부에서 실행하지 않으면, 에러 발생 */
	conf, err := rest.InClusterConfig()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Kubernetes Host: %s\n", conf.Host)

	/* 주어진 config을 위한 clienteset 반환 */
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		log.Panic(err)
	}

	/* Cache 객체 초기화 및 캐시 관리 스레드 실행*/
	Cache = NewCache()
	go Cache.CacheManager()

	/* FxGateway config 생성 */
	c := config.NewFxGatewayConfig(Version)

	/* FxGateway 생성 */
	s := NewFxGateway(c, clientset)

	/* FxGateway 실행 */
	log.Printf("OpenFx Gateway Service Starting...\n")
	s.Start()
}
