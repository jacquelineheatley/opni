package controllers

import (
	"context"
	"fmt"

	. "github.com/kralicky/kmatch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1beta1 "github.com/rancher/opni/apis/core/v1beta1"
	"github.com/rancher/opni/pkg/auth/openid"
	cfgv1beta1 "github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/rancher/opni/pkg/noauth"
	opnimeta "github.com/rancher/opni/pkg/util/meta"
	"github.com/samber/lo"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Core Gateway Controller", Ordered, Label("controller", "slow"), func() {
	When("creating a Gateway resource", func() {
		var gw *corev1beta1.Gateway
		It("should succeed", func() {
			gw = &corev1beta1.Gateway{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: makeTestNamespace(),
				},
				Spec: corev1beta1.GatewaySpec{
					Image: &opnimeta.ImageSpec{
						Image: lo.ToPtr("rancher/opni:latest"),
					},
					Auth: corev1beta1.AuthSpec{
						Provider: cfgv1beta1.AuthProviderNoAuth,
						Noauth:   &noauth.ServerConfig{},
					},
					Alerting: corev1beta1.AlertingSpec{
						Enabled:     true,
						ServiceType: corev1.ServiceTypeLoadBalancer,
						WebPort:     9093,
						ClusterPort: 9094,
						ConfigName:  "alertmanager-config",
						GatewayVolumeMounts: []opnimeta.ExtraVolumeMount{
							{
								Name:      "alerting-storage",
								MountPath: "/var/logs/alerting",
								ReadOnly:  false,
								VolumeSource: corev1.VolumeSource{
									NFS: &corev1.NFSVolumeSource{
										Server: "localhost",
										Path:   "/var/logs/alerting",
									},
								},
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), gw)).To(Succeed())
			Eventually(Object(gw)).Should(Exist())
		})

		It("should create the gateway deployment", func() {
			Eventually(Object(&appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-gateway",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HaveMatchingContainer(And(
					HaveImage("rancher/opni:latest"),
					HavePorts(
						"http",
						"metrics",
						"management-grpc",
						"management-http",
						"management-web",
						"noauth",
					),
					HaveVolumeMounts(
						"config",
						"certs",
						"cortex-client-certs",
						"cortex-server-cacert",
						"alerting-storage",
					),
				)),
				HaveMatchingVolume(And(
					HaveName("config"),
					HaveVolumeSource("ConfigMap"),
				)),
				HaveMatchingVolume(And(
					HaveName("certs"),
					HaveVolumeSource("Secret"),
				)),
				HaveMatchingVolume(And(
					HaveName("cortex-client-certs"),
					HaveVolumeSource("Secret"),
				)),
				HaveMatchingVolume(And(
					HaveName("cortex-server-cacert"),
					HaveVolumeSource("Secret"),
				)),
			))
		})

		It("should create the gateway services", func() {
			Eventually(Object(&corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HavePorts(
					"grpc",
				),
				HaveType(corev1.ServiceTypeLoadBalancer),
			))
			Eventually(Object(&corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-internal",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HavePorts(
					"http",
					"management-grpc",
					"management-http",
				),
				Not(HavePorts(
					"management-web",
				)),
				HaveType(corev1.ServiceTypeClusterIP),
			))
			Eventually(Object(&corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-admin-dashboard",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HavePorts(
					"web",
				),
				HaveType(corev1.ServiceTypeClusterIP),
			))
		})
		It("should create the gateway configmap", func() {
			Eventually(Object(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-gateway",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HaveData("config.yaml", nil),
			))
		})

		It("should create gateway rbac", func() {
			Eventually(Object(&corev1.ServiceAccount{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
			Eventually(Object(&rbacv1.Role{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-crd",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
			Eventually(Object(&rbacv1.RoleBinding{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-crd",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
		})

		It("should create the gateway servicemonitor", func() {
			Eventually(Object(&monitoringv1.ServiceMonitor{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-gateway",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
		})
		It("should configure openid auth", func() {
			gw = &corev1beta1.Gateway{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-openid",
					Namespace: makeTestNamespace(),
				},
				Spec: corev1beta1.GatewaySpec{
					Image: &opnimeta.ImageSpec{
						Image: lo.ToPtr("rancher/opni:latest"),
					},
					Auth: corev1beta1.AuthSpec{
						Provider: cfgv1beta1.AuthProviderOpenID,
						Openid: &corev1beta1.OpenIDConfigSpec{
							ClientID:          "test-client-id",
							ClientSecret:      "test-client-secret",
							Scopes:            []string{"openid", "profile", "email"},
							RoleAttributePath: "test-role-attribute-path",
							OpenidConfig: openid.OpenidConfig{
								Discovery: &openid.DiscoverySpec{
									Issuer: "https://test-issuer/",
								},
							},
						},
					},
				},
			}
			Expect(k8sClient.Create(context.Background(), gw)).To(Succeed())
			Eventually(Object(gw)).Should(Exist())

			cm := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-gateway",
					Namespace: gw.Namespace,
				},
			}
			Eventually(Object(cm)).Should(ExistAnd(
				HaveOwner(gw),
				HaveData("config.yaml", func(data string) bool {
					fmt.Println(data)
					return true
				}),
			))
		})
		//FIXME: opni alertmanager command needs to be in rancher:main for this test to be possible
		XIt("should create the alerting Objects", func() {
			Eventually(Object(&appsv1.StatefulSet{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-alerting-internal",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HaveMatchingContainer(And(
					HaveImage("bitnami/alertmanager:latest"),
					HavePorts(
						"web-port",
						"cluster-port",
					),
					HaveVolumeMounts(
						"opni-alertmanager-data",
						"opni-alertmanager-config",
					),
				)),
				HaveMatchingVolume(
					And(
						HaveName("opni-alertmanager-data"),
						HaveVolumeSource("PersistentVolumeClaim"),
					),
				),
				HaveMatchingVolume(
					And(
						HaveName("opni-alertmanager-config"),
						HaveVolumeSource("ConfigMap"),
					),
				),
			))

			Eventually(Object(&corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "opni-alerting",
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
				HavePorts(
					"web-port",
					"cluster-port",
				),
				HaveType(corev1.ServiceTypeLoadBalancer),
			))
			Eventually(Object(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      gw.Spec.Alerting.ConfigName,
					Namespace: gw.Namespace,
				},
			})).Should(ExistAnd(
				HaveOwner(gw),
			))
		})
		//FIXME: opni alertmanager command needs to be in rancher:main for this test to be possible
		XWhen("disabling alerting", func() {
			It("should remove the alerting objects", func() {
				updateObject(gw, func(gw *corev1beta1.Gateway) *corev1beta1.Gateway {
					gw.Spec.Alerting.Enabled = false
					return gw
				})
				Eventually(Object(&appsv1.StatefulSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "opni-alerting-internal",
						Namespace: gw.Namespace,
					},
				})).ShouldNot(Exist())

				Eventually(Object(&corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "opni-alerting",
						Namespace: gw.Namespace,
					},
				})).ShouldNot(Exist())
			})
		})
	})
})
