

1.  pkg/util
2.  kubectl delete po nginx-7cdbd8cdc9-mkjzb --grace-period=0 --force 强制删除 ；不会延迟；以及有异常的情况
3.  kubectl api-resources  // 里面的所有资源都可以通过 type xxx struct 找到 即对于一个struct;

4.  kubectl api-versions
4.  kubectl get apiservices  etc...
5.  crd ---> operator ---> kubebuilder
5. client-go 
6. 认证 授权 准入控制





=============================================

1. kubectl get crd

2.  kubectl get --raw /apis/batch/v1

3. [root@master http-cache]# kubectl get --raw /apis/apiregistration.k8s.io/v1
{"kind":"APIResourceList","apiVersion":"v1","groupVersion":"apiregistration.k8s.io/v1","resources":[{"name":"apiservices","singularName":"","namespaced":false,"kind":"APIService","verbs":["create","delete","deletecollection","get","list","patch","update","watch"]},{"name":"apiservices/status","singularName":"","namespaced":false,"kind":"APIService","verbs":["get","patch","update"]}]}
[root@master http-cache]#


4. [root@master http-cache]# kubectl get --raw /apis/apiextensions.k8s.io/v1beta1
{"kind":"APIResourceList","apiVersion":"v1","groupVersion":"apiextensions.k8s.io/v1beta1","resources":[{"name":"customresourcedefinitions","singularName":"","namespaced":false,"kind":"CustomResourceDefinition","verbs":["create","delete","deletecollection","get","list","patch","update","watch"],"shortNames":["crd","crds"]},{"name":"customresourcedefinitions/status","singularName":"","namespaced":false,"kind":"CustomResourceDefinition","verbs":["get","patch","update"]}]}
[root@master http-cache]# 

5. [root@master http-cache]# kubectl get --raw /apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions
{"kind":"CustomResourceDefinitionList","apiVersion":"apiextensions.k8s.io/v1beta1","metadata":{"selfLink":"/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions","resourceVersion":"396039"},"items":[]}
[root@master http-cache]# 

6. (n-1)/2

7.  主从 ; svc / headless svc 

8. 
	"k8s.io/client-go/util/retry"
	retry.RetryOnConflict()
9. deploy.status.replicas = deploy.status.readyReplicas
10. Reconcile 调谐


11. # Build and push the app-operator image to a public registry such as docker.io

$ operator-sdk build wanminny/app-operator
$ docker push wanminny/app-operator

# Update the operator manifest to use the built image name (if you are performing these steps on OSX, see note below)
$ sed -i 's|REPLACE_IMAGE|docker.io/wanminy/app-operator|g' deploy/operator.yaml
# On OSX use:
$ sed -i "" 's|REPLACE_IMAGE|docker.io/wanminny/app-operator|g' deploy/operator.yaml

====== 》注意 这个 镜像是 完成功能的；类似于起来 operator-sdk up local (应用目录；否则有错误)；
=====》 如果不起来该镜像；cr等是拉不起来deploy;svc 等的；


12. kubectl delete node //kubeadm 重新join











