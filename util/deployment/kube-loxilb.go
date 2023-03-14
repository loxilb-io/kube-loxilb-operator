/*
 * Copyright (c) 2023 NetLOX Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package deployment

import (
	"fmt"
	"strings"

	kubeloxilb "loxilb-io/kube-loxilb-operator/api/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetKubeLoxilbDeployment(k *kubeloxilb.Kubeloxilbapp) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k.Name,
			Namespace: k.Namespace,
			Labels:    k.Labels,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: k.ObjectMeta.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: k.Labels,
				},
				Spec: corev1.PodSpec{
					HostNetwork:        true,
					ServiceAccountName: "kube-loxilb-operator-kube-loxilb",
					Containers: []corev1.Container{{
						Image:           k.Spec.ContainerImage,
						Name:            "kube-loxilb",
						ImagePullPolicy: corev1.PullPolicy(k.Spec.ImagePullPolicy),
						Command:         []string{"/bin/kube-loxilb"},
						Args:            MakeKubeloxilbArgs(k),
						//Resources: corev1.ResourceRequirements{
						//	Requests: corev1.ResourceList{
						//		corev1.ResourceCPU:    resource.MustParse("100m"),
						//		corev1.ResourceMemory: resource.MustParse("50Mi"),
						//	},
						//	Limits: corev1.ResourceList{
						//		corev1.ResourceCPU:    resource.MustParse("100m"),
						//		corev1.ResourceMemory: resource.MustParse("50Mi"),
						//	},
						//},
					}},
				},
			},
		},
	}
}

func MakeKubeloxilbArgs(k *kubeloxilb.Kubeloxilbapp) []string {
	// TODO: check validation
	args := []string{}

	if len(k.Spec.LoxiURL) > 0 {
		loxiurlStr := strings.Join(k.Spec.LoxiURL, ",")
		args = append(args, fmt.Sprintf("--loxiURL=%s", loxiurlStr))
	}

	if k.Spec.ExternalCIDR != "" {
		args = append(args, fmt.Sprintf("--externalCIDR=%s", k.Spec.ExternalCIDR))
	}

	if k.Spec.SetBGP {
		args = append(args, "--setBGP=true")
	}

	if k.Spec.SetLBMode >= 0 || k.Spec.SetLBMode <= 2 {
		args = append(args, fmt.Sprintf("--setLBMode=%d", k.Spec.SetLBMode))
	}

	return args
}
