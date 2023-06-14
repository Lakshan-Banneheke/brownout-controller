# Brownout Controller for Achieveing Energy Efficiency of Containerized Applications on the Edge

An open-source software application that contains the brownout controller to control and monitor the brownout-enabled edge cluster.

The software application includes,
1. Brownout Controller - Developed using Go
2. API for the Brownout Controller - Developed using Go client for Kubernetes

## Overview of the Brownout Approach
A novel strategy that utilizes Brownout, a container orchestration-based approach for optimizing energy efficiency on the edge while maintaining Quality of Service (QoS). Our proposed approach consists of an analytical model that predicts the edge cluster’s current power consumption and a brownout controller, which activates and deactivates containers based on container selection policies and QoS requirements. The experiment results of the developed Brownout Controller equipped with the Node Idling Selection Policy (NISP) demonstrates that our approach is able to achieve a power saving of 34% while maintaining QoS.

## System Architecture
<img src="https://github.com/Lakshan-Banneheke/brownout-controller/assets/62496951/5275b7b7-0b15-4a5b-b85e-c13adde24b90" width=50% height=50%>

## Brownout Controller

The Brownout controller is the most critical component in the system for reducing power consumption. It consists of 4 components.
1. Brownout Activator.
2. Brownout Algorithm.
3. Container selection policies.
4. Actuators.
  
Once activated by the Brownout Activator, the Brownout Algorithm employs one of the container selection policies to determine the containers that should be deactivated in order to reduce power consumption in the edge cluster while maintaining QoS. The deactivation is done by ensuring that the power consumption is not beyond the upper power threshold calculated (𝑃𝑢𝑡) based on 𝐴𝑆𝑅 and policy constant. Here, the accepted success rate (ASR) and the minimum accepted success rate (MASR) should be provided by the user. Finally, the actuators implement the output configuration of containers by activating/deactivating optional containers on the edge cluster. Overall, the Brownout Controller continuously monitors the cluster and activates/deactivates containers to reduce power consumption while maintaining the success rate between 𝐴𝑆𝑅 and 𝑀 𝐴𝑆𝑅.

## Brownout Algorithm
<img src="https://github.com/Lakshan-Banneheke/brownout-controller/assets/62496951/536b8798-e2ec-4474-aa94-a8b463dbd863" width=30% height=30%>

### Container Selection Policies
The Brownout Controller uses one of the below container selection policies. This can be given as an environment variable.
- LUCF - Least Utilized Container First Policy
- HUCF - Highest Utilized Container First Policy
- RCSP - Random Container Selection Policy
- NISP - Node Idling Selection Policy

## API
The software client exposes an API to access the Brownout Controller.
The [BonE Dashboard](https://github.com/Thushani-Jayasekera/BonE_Dashboard) accesses the controller through this API.

### API Endpoints
![Screenshot 2023-06-14 203058](https://github.com/Lakshan-Banneheke/brownout-controller/assets/62496951/48a4ab3e-44ce-47bb-86db-0d2aa391160b)

