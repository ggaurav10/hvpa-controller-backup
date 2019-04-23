# HVPA controller

## Goal
A pod can be scaled horizontally by deploying more replicas of the pod with the same values of “request” and “limit” parameters, or can be scaled vertically by increasing the “request” parameter.
Currently, separate components – Horizontal Pod Autoscaler and Vertical Pod Autoscaler  - take care of scaling an application running inside a pod. These two components run independent of each other. Scaling by one of the components may disrupt the calculations for the other, which may result in destablising the application.

The proposal is to develop an HVPA controller. It takes only the scaling recommendations from both the pod autoscalers, however, applies those recommendations on its own so that horizontal and vertical scaling can be done in tandem.

A minor goal is to make it also possible to use HVPA only with VPA (without HPA). In this mode, HVPA can be used as an alternative updater for VPA where the upstream resources are updated directly instead of updating the `pods` directly.

## Non-goal

It is not a goal of HVPA to replace Horizontal Pod Autoscaler and Vertical Pod Autoscaler. It proposes to make use of HPA and VPA components (with possible small changes) for recommendations for scaling while taking more control of actually applying the recommendations.

## Modes in which recommendation can be applied:

Most of these modes will require at least one if not both of the pod autoscalers (HPA and VPA) to be provisioned with only recommendations switched on.

VPA already supports this with [`UpdateMode: Off`](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1beta2/types.go#L105).

HPA will have to be enhanced to support simlar recommendation-only mode.

### HPA preferred
When scaling, pods are scaled vertically up only after they are scaled to maximum number of replicas possible. Also, they are scaled vertically down only after they are scaled to minimum number of replicas possible.

#### Pros:
* Easy to implement for cases where HPA and VPA act on same metrices
* Vertical scaling results in rolling update, which may not be ideal for some applications. In this mode, vertical scaling is done only when horizontal scaling is not possible anymore
* Only VPA needs to be deployed with `UpdateMode: Off`. HPA can be deployed noremally.

#### Cons:
* If HPA and VPA act on different metrices `m1` and `m2` respectively, then there may be a case when according to `m1`, horizontal scaling is not required, however, according to `m2` vertical scaling is required. In such a case, there will not be any kind of scaling.

### HPA biased
In case both HPA and VPA have recommendations to scale, HVPA controller will always apply HPA's recommendations. VPA's recommendation will be applied only when HPA doesn't have any recommendation.

#### Pros:
* Works even if HPA and VPA act on different metrices
* Vertical scaling is done only when required
#### Cons:
* Might need both VPA and HPA to be deployed in a recommendation-only mode.

### VPA biased
In case both HPA and VPA have recommendations to scale, HVPA controller will always apply VPA's recommendations. HPA's recommendation will be applied only when VPA doesn't have any recommendation.

#### Pros:
* Works even if HPA and VPA act on different metrices
#### Cons:
* May have more frequent rolling updates. Can it be mitigated by setting user provided thresholds for vertical autoscaling?
* Does it make sense to give preference to VPA instead of HPA, when vertical scaling may result in rolling updates?
* Will need both VPA and HPA to be deployed in a recommendation-only mode.

### Mixed
In this mode, along with the preference for autoscaling (HPA or VPA), user will also provide thresholds for horizontal and vertical autoscaling, say, `x` and `y` respectively. If only one autoscaler has a new recommendation, it will be applied. If both the components have recommendations:
* HPA's recommendation will be applied if the difference between the recommended replicas and the current replicas is greater than `x`.
* VPA's recommendation will be applied if the difference between the recommended value and the current value is greater than `y`, or the recommended value is more than the `limit`.
* In case both above conditions are satisifed, user provided preference will be taken into account

#### Pros:
* Works even if HPA and VPA act on different metrices
* Disruptions because of rolling updates can be controlled as rolling updates will be applied less frequently because of higher threshold
* Provides more control to users for autoscaling apps
* Relies on users' understanding of application, and its scaling
#### Cons:
* Relies on users' understanding of application, and its scaling
* Will need both VPA and HPA to be deployed in a recommendation-only mode.

### Weight based scaling
In this mode, deployment's scaling is divided into 3 stages depending on number of replicas, and a user can provide an interval `x1` to `x2` and 3 `vpaWeight`s between and including `0` and `1`. Stage 1 will be from HPA's `minReplicas` to `x1`. Stage 2 will be from `x1` to `x2`. Stage 3 will be from `x2` to HPA's `maxReplicas`. HVPA controller will consider VPA's recommendations according to the weights provided, and scale the deployment accordingly. The weights given to the HPA's recommendation will be `1 - vpaWeight`.

According to VPA, new `requests` will be `currentRequest + (targetRequest - currentRequest) * vpaWeight`

According to HPA, new number of replicas will be ceil of `currentReplicas + (desiredReplicas - currentReplicas) * (1 - vpaWeight)`

`vpaWeight` = 1 will result in pure vertical scaling
`vpaWeight` = 0 will result in pure horizontal scaling

```
Resource request vs number of replicas curve could typically look like:
Here FirstStageVpaWeight = 0, SecondStageVpaWeight is a fraction, ThirdStageVpaWeight = 0

resource
request
^
|
|                  ____________
|                 /|
|                /
|               /  |
|              /
|  ___________/    |
|
|-------------|----|-----------|----->
   min        x1    x2        max     #Replicas
```

```
Example of vpaWeight vs number of replicas.
Here FirstStageVpaWeight = 0, SecondStageVpaWeight = 0.4, ThirdStageVpaWeight = 0

vpaWeight
   ^
   |
0.4|           |-------|
   |           |       |
  0|   ________|       |____________
   |------------------------------------->
               x1     x2               #Replicas
```

#### Pros
* Works even if HPA and VPA act on different metrices
* Gives better control to user on scaling of apps
#### Cons
* Need to define behavior when:
    * current number of replicas is `x2`, the deployment has not yet scaled up to `maxAllowed`, and user has chosen to scale only horizontally after `x2`. In this case, the deployment cannot scale up vertically anymore if the load increases.
    * current number of replicas is `x1`, the deployment has not yet scaled down to `minAllowed`, and user has chosen to scale only horizontally for lower values than `x1`. In this case, the deployment cannot scale down vertically anymore if the load decreases.
* Will need both VPA and HPA to be deployed in a recommendation-only mode.

#### Mitigation
* `x1` and `x2` will be optional, If either or both of them are not provided, vertical scaling will be done until `minAllowed` and/or `maxAllowed` correspondingly.
* Once replica count reaches `maxReplicas`, and VPA still recommends higher resource requirements, then vertical scaling will be done unconditionally

#### Spec

```golang
// HvpaSpec defines the desired state of Hvpa
type HvpaSpec struct {
    // Scaling of a deployment is divided into 3 stages.
	// Stage 1 starts at HPA's minReplicas
	// Stage 3 ends at HPA's maxReplicas
	// When in stage 2, weight given to VPA's recommendation will be SecondStageVpaWeigh
	// StageTwoStartReplicaCount defines the number of replicas when stage 2 starts
	// +optional
	StageTwoStartReplicaCount int `json:"StageTwoStartReplicaCount,omitempty"`

	// StageTwoStopReplicaCount defines the number of replicas when stage 2 stops
	// +optional
	StageTwoStopReplicaCount int `json:"StageTwoStopReplicaCount,omitempty"`

	// FirstStageVpaWeight defines the weight to be given VPA recommendations in stage 1
	FirstStageVpaWeight VpaWeight `json:"FirstStageVpaWeight,omitempty"`

	// SecondStageVpaWeight defines the weight to be given VPA recommendations in stage 2
	SecondStageVpaWeight VpaWeight `json:"SecondStageVpaWeight,omitempty"`

	// ThirdStageVpaWeight defines the weight to be given VPA recommendations in stage 3
	ThirdStageVpaWeight VpaWeight `json:"ThirdStageVpaWeight,omitempty"`

	// HpaSpec defines the spec of HPA
	HpaSpec scaling_v1.HorizontalPodAutoscaler `json:"hpaSpec,omitempty"`

	// VpaSpec defines the spec of VPA
	VpaSpec vpa_api.VerticalPodAutoscaler `json:"vpaSpec,omitempty"`
}

// VpaWeight - weight to provide to VPA scaling
type VpaWeight float32

const (
	// VpaOnly - only vertical scaling
	VpaOnly VpaWeight = 1.0
	// HpaOnly - only horizontal scaling
	HpaOnly VpaWeight = 0
)
```

Example 1:

```yaml
apiVersion: autoscaling.k8s.io/v1alpha1
kind: Hvpa
metadata:
  name: hvpa-sample
spec:
  stageTwoStartReplicaCount: 3
  stageTwoStopReplicaCount: 7
  firstStageVpaWeight: 0
  secondStageVpaWeight: 0.4
  thirdStageVpaWeight: 0
  hpaSpec:
    maxReplicas: 10
    minReplicas: 1
    scaleTargetRef:
      apiVersion: apps/v1
      kind: Deployment
      name: resource-consumer
    targetCPUUtilizationPercentage: 60
  vpaSpec:
    resourcePolicy:
      containerPolicies:
      - MinAllowed:
          memory: 400Mi
        containerName: resource-consumer
        maxAllowed:
          memory: 3000Mi
    updatePolicy:
      updateMode: "Off"
```
```
resource
request
^
|
|                  _________
|                 /|
|                /
|               /  |
|              /
|  ___________/    |
|
|-------------|----|--------|----->
   1          3    7        10     #Replicas
```


Example 2:
```yaml
apiVersion: autoscaling.k8s.io/v1alpha1
kind: Hvpa
metadata:
  name: hvpa-sample
spec:
  stageTwoStartReplicaCount: 1
  stageTwoStopReplicaCount: 4
  firstStageVpaWeight: 1
  secondStageVpaWeight: 0.4
  thirdStageVpaWeight: 0
  hpaSpec:
    .
    .
    .
```
```
resource
request
^
| .     ______
|      /
|     /  
|    /
|   /  
|  |
|  |  
|  |
|
|--|----|-----|------->
   1   4       10     #Replicas
```


Example 3: `x2` is not provided. After maxReplicas, there is only vertical scaling
```yaml
apiVersion: autoscaling.k8s.io/v1alpha1
kind: Hvpa
metadata:
  name: hvpa-sample
spec:
  stageTwoStartReplicaCount: 3
  firstStageVpaWeight: 0
  secondStageVpaWeight: 0.4
  thirdStageVpaWeight: 0
  hpaSpec:
    minReplicas: 1
    maxReplicas: 10
    .
    .
    .
```
```
resource
request
^
|
|           |
|           |
|           .     
|          /
|         /
|        / 
|       /
|  ____/ 
|
|--|---|----|------>
   1   3    10     #Replicas
```

### Weight based scaling for multiple intervals of replica counts
Instead of just one interval of `x1` to `x2` as in above approach, user will be able to provide multiple such intervals and corresponding `vpaWeight`s in an array `WeightBasedScaling`.

User may also provide values for `scaleUpDelay` and/or `scaleDownDelay`. These values will be used to keep check on frequent scaling recommendations. If HPA or VPA recommends higher resources, then these recommendations will be applied only after duration provided in `scaleUpDelay` has passed. Same will be the case for scaling down.

User can also provide thresholds for vertical scaling of CPU and/or memory as either absolute values or percentage of current `requests` value or both. In case, both types of thresholds are provided, minimum of the 2 will be taken. In case none of the thresholds are provided, HVPA will use default values of `200M` for memory and `200m` for CPU.

Providing 2 values helps in cases where `minAllowed` and `maxAllowed` for VPA has a large range. For example, let's say `minAllowed` is `100M` and `maxAllowed` is `4000M` for memory. Here user can mention threshold values - `500M` and `80%`. For smaller range of `requests`, eg. `200M`, a user might want to scale up if the new recommendation is higher than `360M` (+80%). For larger values, eg. `3500M`, VPA threshold value of `500M` might make more sense than 80% of `3500M`, ie, `2800M`.

If only one type of threshold is provided, that is used as is.

#### Pros
* Works even if HPA and VPA act on different metrices
* Gives better control than previous approaches to user on scaling of apps 
#### Cons
* Need to define behavior when:
    * current number of replicas is `x2`, the deployment has not yet scaled up to `maxAllowed`, and user has chosen to scale only horizontally after `x2`. In this case, the deployment cannot scale up vertically anymore if the load increases.
    * current number of replicas is `x1`, the deployment has not yet scaled down to `minAllowed`, and user has chosen to scale only horizontally for lower values than `x1`. In this case, the deployment cannot scale down vertically anymore if the load decreases.
* Will need both VPA and HPA to be deployed in a recommendation-only mode. 

#### Spec
[Here](https://github.com/ggaurav10/hvpa-controller/blob/master/pkg/apis/autoscaling/v1alpha1/hvpa_types.go) is the spec for HVPA controller

Example 1:

```yaml
apiVersion: autoscaling.k8s.io/v1alpha1
kind: Hvpa
metadata:
  name: hvpa-sample
spec:
  scaleUpDelay: "2m"
  scaleDownDelay: "3m"
  minCpuChange:
    value: "200m"
    percentage: 70
  minMemChange:
    value: "200M"
    percentage: 80
  hpaTemplate:
    maxReplicas: 10
    minReplicas: 1
    metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 50
    - type: Resource
      resource:
        name: memory
        target:
          type: Utilization
          averageUtilization: 50
  vpaTemplate:
    resourcePolicy:
      containerPolicies:
      - MinAllowed:
          memory: 400Mi
        containerName: resource-consumer
        maxAllowed:
          memory: 3000Mi
  weightBasedScalingIntervals:
    - vpaWeight: 0
      startReplicaCount: 1
      lastReplicaCount: 2
    - vpaWeight: 0.6
      startReplicaCount: 3
      lastReplicaCount: 7
    - vpaWeight: 0
      startReplicaCount: 7
      lastReplicaCount: 10
  targetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: resource-consumer
```
```
resource
request
^
|
|                  _________
|                 /|
|                /
|               /  |
|              /
|  ___________/    |
|
|-------------|----|--------|----->
   1          3    7        10     #Replicas
```


Example 2:
```yaml
apiVersion: autoscaling.k8s.io/v1alpha1
kind: Hvpa
metadata:
  name: hvpa-sample
spec:
  weightBasedScalingIntervals:
    - vpaWeight: 1
      startReplicaCount: 1
      lastReplicaCount: 1
    - vpaWeight: 0.6
      startReplicaCount: 2
      lastReplicaCount: 4
    - vpaWeight: 0
      startReplicaCount: 5
      lastReplicaCount: 10
  hpaTemplate:
    .
    .
    .
```
```
resource
request
^
|       ______
|      /
|     /  
|    /
|   /  
|  |
|  |  
|  |
|
|--|----|-----|------->
   1   4       10     #Replicas
```


Example 3: After maxReplicas, even if weight is not mentioned, there is only vertical scaling because HPA is limited by maxReplicas
```yaml
apiVersion: autoscaling.k8s.io/v1alpha1
kind: Hvpa
metadata:
  name: hvpa-sample
spec:
  weightBasedScalingIntervals:
    - vpaWeight: 0
      startReplicaCount: 1
      lastReplicaCount: 3
    - vpaWeight: 0.6
      startReplicaCount: 4
      lastReplicaCount: 10
  hpaTemplate:
    minReplicas: 1
    maxReplicas: 10
    .
    .
    .
```
```
resource
request
^
|
|           |
|           |
|           .     
|          /
|         /
|        / 
|       /
|  ____/ 
|
|--|---|----|------>
   1   3    10     #Replicas
```

## [Currently preferred approach](#Weight-based-scaling-for-multiple-intervals-of-replica-counts)