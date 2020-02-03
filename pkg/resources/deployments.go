//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package resources

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ControllerDeployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name:      CertManagerControllerName,
		Namespace: DeployNamespace,
		Labels:    ControllerLabelMap,
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: &replicaCount,
		Selector: &metav1.LabelSelector{
			MatchLabels: ControllerLabelMap,
		},
		Template: corev1.PodTemplateSpec{ // missing: annotations
			ObjectMeta: metav1.ObjectMeta{
				Labels: ControllerLabelMap,
			},
			Spec: certManagerControllerPod,
		},
	},
}

var WebhookDeployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name:      CertManagerWebhookName,
		Namespace: DeployNamespace,
		Labels:    WebhookLabelMap,
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: &replicaCount,
		Selector: &metav1.LabelSelector{
			MatchLabels: WebhookLabelMap,
		},
		Template: corev1.PodTemplateSpec{ // missing: annotations
			ObjectMeta: metav1.ObjectMeta{
				Labels: WebhookLabelMap,
			},
			Spec: certManagerWebhookPod,
		},
	},
}

var CainjectorDeployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name:      CertManagerCainjectorName,
		Namespace: DeployNamespace,
		Labels:    CainjectorLabelMap,
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: &replicaCount,
		Selector: &metav1.LabelSelector{
			MatchLabels: CainjectorLabelMap,
		},
		Template: corev1.PodTemplateSpec{ // missing: annotations
			ObjectMeta: metav1.ObjectMeta{
				Labels: CainjectorLabelMap,
			},
			Spec: certManagerCainjectorPod,
		},
	},
}
