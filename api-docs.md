# Rancher Norman Stype API Docs

This is the API docs for [http://192.168.99.119:8080/v2-beta](http://192.168.99.119:8080/v2-beta).

## Schema List

Followings are the API links in the level.

- [accountLink](#schema-accountlink-api-list) [example](http://192.168.99.119:8080/v2-beta/accountlinks)
- [account](#schema-account-api-list) [example](http://192.168.99.119:8080/v2-beta/accounts)
- [agent](#schema-agent-api-list) [example](http://192.168.99.119:8080/v2-beta/agents)
- [apiKey](#schema-apikey-api-list) [example](http://192.168.99.119:8080/v2-beta/apikeys)
- [auditLog](#schema-auditlog-api-list) [example](http://192.168.99.119:8080/v2-beta/auditlogs)
- [azureadconfig](#schema-azureadconfig-api-list) [example](http://192.168.99.119:8080/v2-beta/azureadconfigs)
- [backupTarget](#schema-backuptarget-api-list) [example](http://192.168.99.119:8080/v2-beta/backuptargets)
- [backup](#schema-backup-api-list) [example](http://192.168.99.119:8080/v2-beta/backups)
- [certificate](#schema-certificate-api-list) [example](http://192.168.99.119:8080/v2-beta/certificates)
- [clusterMembership](#schema-clustermembership-api-list) [example](http://192.168.99.119:8080/v2-beta/clustermemberships)
- [composeProject](#schema-composeproject-api-list) [example](http://192.168.99.119:8080/v2-beta/composeprojects)
- [composeService](#schema-composeservice-api-list) [example](http://192.168.99.119:8080/v2-beta/composeservices)
- [configItemStatus](#schema-configitemstatus-api-list) [example](http://192.168.99.119:8080/v2-beta/configitemstatuses)
- [configItem](#schema-configitem-api-list) [example](http://192.168.99.119:8080/v2-beta/configitems)
- [containerEvent](#schema-containerevent-api-list) [example](http://192.168.99.119:8080/v2-beta/containerevents)
- [container](#schema-container-api-list) [example](http://192.168.99.119:8080/v2-beta/containers)
- [credential](#schema-credential-api-list) [example](http://192.168.99.119:8080/v2-beta/credentials)
- [databasechangelog](#schema-databasechangelog-api-list) [example](http://192.168.99.119:8080/v2-beta/databasechangelogs)
- [databasechangeloglock](#schema-databasechangeloglock-api-list) [example](http://192.168.99.119:8080/v2-beta/databasechangeloglocks)
- [dnsService](#schema-dnsservice-api-list) [example](http://192.168.99.119:8080/v2-beta/dnsservices)
- [extensionPoint](#schema-extensionpoint-api-list) [example](http://192.168.99.119:8080/v2-beta/extensionpoints)
- [externalCredential](#schema-externalcredential-api-list) [example](http://192.168.99.119:8080/v2-beta/externalcredentials)
- [externalDnsEvent](#schema-externaldnsevent-api-list) [example](http://192.168.99.119:8080/v2-beta/externaldnsevents)
- [externalEvent](#schema-externalevent-api-list) [example](http://192.168.99.119:8080/v2-beta/externalevents)
- [externalHandlerExternalHandlerProcessMap](#schema-externalhandlerexternalhandlerprocessmap-api-list) [example](http://192.168.99.119:8080/v2-beta/externalhandlerexternalhandlerprocessmaps)
- [externalHandlerProcess](#schema-externalhandlerprocess-api-list) [example](http://192.168.99.119:8080/v2-beta/externalhandlerprocesses)
- [externalHandler](#schema-externalhandler-api-list) [example](http://192.168.99.119:8080/v2-beta/externalhandlers)
- [externalHostEvent](#schema-externalhostevent-api-list) [example](http://192.168.99.119:8080/v2-beta/externalhostevents)
- [externalServiceEvent](#schema-externalserviceevent-api-list) [example](http://192.168.99.119:8080/v2-beta/externalserviceevents)
- [externalService](#schema-externalservice-api-list) [example](http://192.168.99.119:8080/v2-beta/externalservices)
- [externalStoragePoolEvent](#schema-externalstoragepoolevent-api-list) [example](http://192.168.99.119:8080/v2-beta/externalstoragepoolevents)
- [externalVolumeEvent](#schema-externalvolumeevent-api-list) [example](http://192.168.99.119:8080/v2-beta/externalvolumeevents)
- [genericObject](#schema-genericobject-api-list) [example](http://192.168.99.119:8080/v2-beta/genericobjects)
- [haConfigInput](#schema-haconfiginput-api-list) [example](http://192.168.99.119:8080/v2-beta/haconfiginputs)
- [haConfig](#schema-haconfig-api-list) [example](http://192.168.99.119:8080/v2-beta/haconfigs)
- [healthcheckInstanceHostMap](#schema-healthcheckinstancehostmap-api-list) [example](http://192.168.99.119:8080/v2-beta/healthcheckinstancehostmaps)
- [hostApiProxyToken](#schema-hostapiproxytoken-api-list) [example](http://192.168.99.119:8080/v2-beta/hostapiproxytokens)
- [hostTemplate](#schema-hosttemplate-api-list) [example](http://192.168.99.119:8080/v2-beta/hosttemplates)
- [host](#schema-host-api-list) [example](http://192.168.99.119:8080/v2-beta/hosts)
- [identity](#schema-identity-api-list) [example](http://192.168.99.119:8080/v2-beta/identities)
- [image](#schema-image-api-list) [example](http://192.168.99.119:8080/v2-beta/images)
- [instanceLink](#schema-instancelink-api-list) [example](http://192.168.99.119:8080/v2-beta/instancelinks)
- [instance](#schema-instance-api-list) [example](http://192.168.99.119:8080/v2-beta/instances)
- [ipAddress](#schema-ipaddress-api-list) [example](http://192.168.99.119:8080/v2-beta/ipaddresses)
- [kubernetesService](#schema-kubernetesservice-api-list) [example](http://192.168.99.119:8080/v2-beta/kubernetesservices)
- [kubernetesStack](#schema-kubernetesstack-api-list) [example](http://192.168.99.119:8080/v2-beta/kubernetesstacks)
- [label](#schema-label-api-list) [example](http://192.168.99.119:8080/v2-beta/labels)
- [loadBalancerService](#schema-loadbalancerservice-api-list) [example](http://192.168.99.119:8080/v2-beta/loadbalancerservices)
- [localAuthConfig](#schema-localauthconfig-api-list) [example](http://192.168.99.119:8080/v2-beta/localauthconfigs)
- [machineDriver](#schema-machinedriver-api-list) [example](http://192.168.99.119:8080/v2-beta/machinedrivers)
- [machine](#schema-machine-api-list) [example](http://192.168.99.119:8080/v2-beta/machines)
- [mount](#schema-mount-api-list) [example](http://192.168.99.119:8080/v2-beta/mounts)
- [networkDriverService](#schema-networkdriverservice-api-list) [example](http://192.168.99.119:8080/v2-beta/networkdriverservices)
- [networkDriver](#schema-networkdriver-api-list) [example](http://192.168.99.119:8080/v2-beta/networkdrivers)
- [networkPolicyRuleWithin](#schema-networkpolicyrulewithin-api-list) [example](http://192.168.99.119:8080/v2-beta/networkpolicyrulewithins)
- [network](#schema-network-api-list) [example](http://192.168.99.119:8080/v2-beta/networks)
- [openldapconfig](#schema-openldapconfig-api-list) [example](http://192.168.99.119:8080/v2-beta/openldapconfigs)
- [password](#schema-password-api-list) [example](http://192.168.99.119:8080/v2-beta/passwords)
- [physicalHost](#schema-physicalhost-api-list) [example](http://192.168.99.119:8080/v2-beta/physicalhosts)
- [port](#schema-port-api-list) [example](http://192.168.99.119:8080/v2-beta/ports)
- [processDefinition](#schema-processdefinition-api-list) [example](http://192.168.99.119:8080/v2-beta/processdefinitions)
- [processExecution](#schema-processexecution-api-list) [example](http://192.168.99.119:8080/v2-beta/processexecutions)
- [processInstance](#schema-processinstance-api-list) [example](http://192.168.99.119:8080/v2-beta/processinstances)
- [processPool](#schema-processpool-api-list) [example](http://192.168.99.119:8080/v2-beta/processpools)
- [processSummary](#schema-processsummary-api-list) [example](http://192.168.99.119:8080/v2-beta/processsummary)
- [projectMember](#schema-projectmember-api-list) [example](http://192.168.99.119:8080/v2-beta/projectmembers)
- [projectTemplate](#schema-projecttemplate-api-list) [example](http://192.168.99.119:8080/v2-beta/projecttemplates)
- [project](#schema-project-api-list) [example](http://192.168.99.119:8080/v2-beta/projects)
- [pullTask](#schema-pulltask-api-list) [example](http://192.168.99.119:8080/v2-beta/pulltasks)
- [region](#schema-region-api-list) [example](http://192.168.99.119:8080/v2-beta/regions)
- [register](#schema-register-api-list) [example](http://192.168.99.119:8080/v2-beta/register)
- [registrationToken](#schema-registrationtoken-api-list) [example](http://192.168.99.119:8080/v2-beta/registrationtokens)
- [registryCredential](#schema-registrycredential-api-list) [example](http://192.168.99.119:8080/v2-beta/registrycredentials)
- [registry](#schema-registry-api-list) [example](http://192.168.99.119:8080/v2-beta/registries)
- [resourceDefinition](#schema-resourcedefinition-api-list) [example](http://192.168.99.119:8080/v2-beta/resourcedefinitions)
- [scheduledUpgrade](#schema-scheduledupgrade-api-list) [example](http://192.168.99.119:8080/v2-beta/scheduledupgrades)
- [schema](#schema-schema-api-list) [example](http://192.168.99.119:8080/v2-beta/schemas)
- [secret](#schema-secret-api-list) [example](http://192.168.99.119:8080/v2-beta/secrets)
- [serviceConsumeMap](#schema-serviceconsumemap-api-list) [example](http://192.168.99.119:8080/v2-beta/serviceconsumemaps)
- [serviceEvent](#schema-serviceevent-api-list) [example](http://192.168.99.119:8080/v2-beta/serviceevents)
- [serviceExposeMap](#schema-serviceexposemap-api-list) [example](http://192.168.99.119:8080/v2-beta/serviceexposemaps)
- [serviceLog](#schema-servicelog-api-list) [example](http://192.168.99.119:8080/v2-beta/servicelogs)
- [serviceProxy](#schema-serviceproxy-api-list) [example](http://192.168.99.119:8080/v2-beta/serviceproxies)
- [service](#schema-service-api-list) [example](http://192.168.99.119:8080/v2-beta/services)
- [setting](#schema-setting-api-list) [example](http://192.168.99.119:8080/v2-beta/settings)
- [snapshot](#schema-snapshot-api-list) [example](http://192.168.99.119:8080/v2-beta/snapshots)
- [stack](#schema-stack-api-list) [example](http://192.168.99.119:8080/v2-beta/stacks)
- [storageDriverService](#schema-storagedriverservice-api-list) [example](http://192.168.99.119:8080/v2-beta/storagedriverservices)
- [storageDriver](#schema-storagedriver-api-list) [example](http://192.168.99.119:8080/v2-beta/storagedrivers)
- [storagePool](#schema-storagepool-api-list) [example](http://192.168.99.119:8080/v2-beta/storagepools)
- [subnet](#schema-subnet-api-list) [example](http://192.168.99.119:8080/v2-beta/subnets)
- [taskInstance](#schema-taskinstance-api-list) [example](http://192.168.99.119:8080/v2-beta/taskinstances)
- [task](#schema-task-api-list) [example](http://192.168.99.119:8080/v2-beta/tasks)
- [typeDocumentation](#schema-typedocumentation-api-list) [example](http://192.168.99.119:8080/v2-beta/typedocumentations)
- [userPreference](#schema-userpreference-api-list) [example](http://192.168.99.119:8080/v2-beta/userpreferences)
- [virtualMachine](#schema-virtualmachine-api-list) [example](http://192.168.99.119:8080/v2-beta/virtualmachines)
- [volumeTemplate](#schema-volumetemplate-api-list) [example](http://192.168.99.119:8080/v2-beta/volumetemplates)
- [volume](#schema-volume-api-list) [example](http://192.168.99.119:8080/v2-beta/volumes)

## API Type List

Following are the schema types in this endpoint.

### account
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/U | false | true | \* - 255 |  |
| externalIdType | string | C/R/U | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| identity | [identity](#identity) | -/R/- | false | false | - |  |
| kind | string | C/R/U | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive/upgrading;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | C/R/- | false | true | \* - 128 |  |
| version | string | -/R/- | false | true | \* - 128 |  |

### accountLink
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| external | boolean | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| linkedAccount | string | -/R/- | false | true | \* - 255 |  |
| linkedAccountUuid | string | -/R/- | false | false | - |  |
| linkedRegion | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### activeSetting
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| activeValue | json | -/R/- | false | false | - |  |
| id | string | -/R/- | false | false | - |  |
| inDb | boolean | -/R/- | false | false | - |  |
| name | string | -/R/- | false | false | - |  |
| source | string | -/R/- | false | false | - |  |
| value | string | -/R/U | false | false | - |  |

### addOutputsInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| outputs | map[string] | C/R/- | true | false | - |  |

### addRemoveServiceLinkInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| serviceLink | serviceLink | C/R/- | true | false | - |  |

### agent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| agentResourcesAccountId | [account](#account) | C/R/- | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalCredentials | array[externalCredential] | C/R/- | false | true | - |  |
| externalId | string | C/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| labels | map[string] | C/R/- | false | true | - |  |
| managedConfig | boolean | C/R/- | false | false | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| publicValue | string | C/R/- | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| secretValue | string | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/disconnected/disconnecting/finishing-reconnect/inactive/purged/purging/reconnected/reconnecting/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uri | string | C/R/- | false | true | \* - 255 |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### amazonec2Config
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accessKey | string | C/R/- | false | false | - | description: AWS Access Key;  |
| ami | string | C/R/- | false | false | - | description: AWS machine image;  |
| blockDurationMinutes | string | C/R/- | false | false | - | description: AWS spot instance duration in minutes (60, 120, 180, 240, 300, or 360);  |
| deviceName | string | C/R/- | false | false | - | description: AWS root device name;  |
| endpoint | string | C/R/- | false | false | - | description: Optional endpoint URL (hostname only or fully qualified URI);  |
| iamInstanceProfile | string | C/R/- | false | false | - | description: AWS IAM Instance Profile;  |
| insecureTransport | boolean | C/R/- | false | false | - | description: Disable SSL when sending requests;  |
| instanceType | string | C/R/- | false | false | - | description: AWS instance type;  |
| monitoring | boolean | C/R/- | false | false | - | description: Set this flag to enable CloudWatch monitoring;  |
| openPort | array[string] | C/R/- | false | false | - | description: Make the specified port number accessible from the Internet;  |
| privateAddressOnly | boolean | C/R/- | false | false | - | description: Only use a private IP address;  |
| region | string | C/R/- | false | false | - | description: AWS region;  |
| requestSpotInstance | boolean | C/R/- | false | false | - | description: Set this flag to request spot instance;  |
| retries | string | C/R/- | false | false | - | description: Set retry count for recoverable failures (use -1 to disable);  |
| rootSize | string | C/R/- | false | false | - | description: AWS root disk size (in GB);  |
| secretKey | string | C/R/- | false | false | - | description: AWS Secret Key;  |
| securityGroup | array[string] | C/R/- | false | false | - | description: AWS VPC security group;  |
| sessionToken | string | C/R/- | false | false | - | description: AWS Session Token;  |
| spotPrice | string | C/R/- | false | false | - | description: AWS spot instance bid price (in dollar);  |
| sshUser | string | C/R/- | false | false | - | description: Set the name of the ssh user;  |
| subnetId | string | C/R/- | false | false | - | description: AWS VPC subnet id;  |
| tags | string | C/R/- | false | false | - | description: AWS Tags (e.g. key1,value1,key2,value2);  |
| useEbsOptimizedInstance | boolean | C/R/- | false | false | - | description: Create an EBS optimized instance;  |
| usePrivateAddress | boolean | C/R/- | false | false | - | description: Force the usage of private IP address;  |
| volumeType | string | C/R/- | false | false | - | description: Amazon EBS volume type;  |
| vpcId | string | C/R/- | false | false | - | description: AWS VPC id;  |
| zone | string | C/R/- | false | false | - | description: AWS zone for instance (i.e. a,b,c,d,e);  |

### apiKey
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | C/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| publicValue | string | C/R/- | false | true | \* - 4096 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| secretValue | password | C/R/- | false | true | \* - 4096 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### auditLog
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| authType | enum | -/R/- | true | true | \* - 255 | enum options: AdminAuth/BasicAuth/RegistrationToken/TokenAccount/TokenAuth;  |
| authenticatedAsAccountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| authenticatedAsIdentityId | [identity](#identity) | -/R/- | false | true | \* - 255 |  |
| clientIp | string | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| description | string | -/R/- | false | true | \* - 1024 |  |
| eventType | string | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| requestObject | string | -/R/- | false | false | - |  |
| resourceId | int | -/R/- | false | true | \* - 255 |  |
| resourceType | string | -/R/- | false | true | \* - 255 |  |
| responseCode | int | -/R/- | false | false | - |  |
| responseObject | string | -/R/- | false | false | - |  |

### azureConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| availabilitySet | string | C/R/- | false | false | - | description: Azure Availability Set to place the virtual machine into;  |
| clientId | string | C/R/- | false | false | - | description: Azure Service Principal Account ID (optional, browser auth is used if not specified);  |
| clientSecret | string | C/R/- | false | false | - | description: Azure Service Principal Account password (optional, browser auth is used if not specified);  |
| dns | string | C/R/- | false | false | - | description: A unique DNS label for the public IP adddress;  |
| dockerPort | string | C/R/- | false | false | - | description: Port number for Docker engine;  |
| environment | string | C/R/- | false | false | - | description: Azure environment (e.g. AzurePublicCloud, AzureChinaCloud);  |
| image | string | C/R/- | false | false | - | description: Azure virtual machine OS image;  |
| location | string | C/R/- | false | false | - | description: Azure region to create the virtual machine;  |
| noPublicIp | boolean | C/R/- | false | false | - | description: Do not create a public IP address for the machine;  |
| openPort | array[string] | C/R/- | false | false | - | description: Make the specified port number accessible from the Internet;  |
| privateIpAddress | string | C/R/- | false | false | - | description: Specify a static private IP address for the machine;  |
| resourceGroup | string | C/R/- | false | false | - | description: Azure Resource Group name (will be created if missing);  |
| size | string | C/R/- | false | false | - | description: Size for Azure Virtual Machine;  |
| sshUser | string | C/R/- | false | false | - | description: Username for SSH login;  |
| staticPublicIp | boolean | C/R/- | false | false | - | description: Assign a static public IP address to the machine;  |
| storageType | string | C/R/- | false | false | - | description: Type of Storage Account to host the OS Disk for the machine;  |
| subnet | string | C/R/- | false | false | - | description: Azure Subnet Name to be used within the Virtual Network;  |
| subnetPrefix | string | C/R/- | false | false | - | description: Private CIDR block to be used for the new subnet, should comply RFC 1918;  |
| subscriptionId | string | C/R/- | false | false | - | description: Azure Subscription ID;  |
| usePrivateIp | boolean | C/R/- | false | false | - | description: Use private IP address of the machine to connect;  |
| vnet | string | C/R/- | false | false | - | description: Azure Virtual Network name to connect the virtual machine (in [resourcegroup:]name format);  |

### azureadconfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accessMode | string | C/R/- | true | false | - |  |
| adminAccountPassword | string | C/R/- | false | true | 1 - \* |  |
| adminAccountUsername | string | C/R/- | false | true | 1 - \* |  |
| clientId | string | C/R/- | false | true | - |  |
| domain | string | C/R/- | false | true | - |  |
| enabled | boolean | C/R/- | false | true | - |  |
| name | string | C/R/U | false | false | - |  |
| tenantId | string | C/R/- | false | true | - |  |

### backup
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| backupTargetId | [backupTarget](#backupTarget) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| snapshotId | [snapshot](#snapshot) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uri | string | -/R/- | false | true | \* - 4096 |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volumeId | [volume](#volume) | -/R/- | false | true | \* - 255 |  |

### backupTarget
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | true | false | 1 - 255 | valid chars: a-zA-Z0-9_.-;  |
| nfsConfig | nfsConfig | C/R/- | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### baseMachineConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |

### binding
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| services | map[serviceBinding] | C/R/- | false | false | - |  |

### blkioDeviceOption
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| readBps | int | C/R/- | false | true | - |  |
| readIops | int | C/R/- | false | true | - |  |
| weight | int | C/R/- | false | true | - |  |
| writeBps | int | C/R/- | false | true | - |  |
| writeIops | int | C/R/- | false | true | - |  |

### catalogTemplate
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| answers | map[json] | C/R/- | false | false | - |  |
| binding | binding | C/R/- | false | false | - |  |
| description | string | C/R/- | false | false | - |  |
| dockerCompose | string | C/R/- | false | false | - |  |
| name | string | C/R/- | false | false | - |  |
| rancherCompose | string | C/R/- | false | false | - |  |
| templateId | string | C/R/- | false | false | - |  |
| templateVersionId | string | C/R/- | false | false | - |  |

### certificate
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| CN | string | -/R/- | false | false | - |  |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| algorithm | string | -/R/- | false | false | - |  |
| cert | string | C/R/U | true | false | \* - 65535 |  |
| certChain | string | C/R/U | false | true | \* - 65535 |  |
| certFingerprint | string | -/R/- | false | false | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| expiresAt | string | -/R/- | false | false | - |  |
| id | int | -/R/- | false | true | - |  |
| issuedAt | string | -/R/- | false | false | - |  |
| issuer | string | -/R/- | false | false | - |  |
| key | string | C/R/U | true | false | \* - 65535 |  |
| keySize | int | -/R/- | false | false | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | true | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serialNumber | string | -/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested/updating-active;  |
| subjectAlternativeNames | array[string] | -/R/- | false | false | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| version | string | -/R/- | false | false | - |  |

### changeSecretInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| newSecret | string | C/R/- | true | false | - |  |
| oldSecret | string | C/R/- | true | false | - |  |

### clusterMembership
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| clustered | boolean | -/R/- | false | false | \* - 255 |  |
| config | string | -/R/- | false | true | \* - 16777215 |  |
| heartbeat | int | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| name | string | -/R/- | false | true | \* - 255 |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### composeConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| dockerComposeConfig | string | -/R/- | false | false | - |  |
| rancherComposeConfig | string | -/R/- | false | false | - |  |

### composeConfigInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| serviceIds | array[[service](#service)] | C/R/- | false | true | - |  |

### composeProject
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| answers | map[json] | C/R/- | false | false | - |  |
| binding | binding | C/R/U | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| environment | map[string] | C/R/- | false | false | - |  |
| externalId | string | C/R/U | false | true | \* - 128 |  |
| group | string | C/R/U | false | true | \* - 255 |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/- | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| previousEnvironment | map[string] | C/R/U | false | true | - |  |
| previousExternalId | string | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serviceIds | array[[service](#service)] | -/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/error/erroring/finishing-upgrade/removed/removing/requested/rolling-back/updating-active/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| templates | map[string] | C/R/- | true | true | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### composeService
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| currentScale | int | -/R/- | false | false | - |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| fqdn | string | -/R/- | false | true | - |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| launchConfig | launchConfig | C/R/- | false | true | - |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| publicEndpoints | array[publicEndpoint] | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| scale | int | C/R/U | false | false | - |  |
| scalePolicy | scalePolicy | C/R/U | false | true | - |  |
| selectorContainer | string | C/R/U | false | true | \* - 4096 |  |
| selectorLink | string | C/R/U | false | true | \* - 4096 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vip | string | C/R/- | false | true | \* - 255 |  |

### configItem
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| id | int | -/R/- | false | true | - |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| sourceVersion | string | -/R/- | false | true | \* - 1024 |  |

### configItemStatus
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| agentId | [agent](#agent) | -/R/- | false | true | \* - 255 |  |
| appliedUpdated | date | -/R/- | false | true | \* - 255 |  |
| appliedVersion | int | -/R/U | false | false | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| requestedUpdated | date | -/R/- | false | false | \* - 255 |  |
| requestedVersion | int | -/R/- | false | false | \* - 255 |  |
| sourceVersion | string | -/R/- | false | true | \* - 255 |  |

### container
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| agentId | [agent](#agent) | -/R/- | false | true | \* - 255 |  |
| allocationState | string | -/R/- | false | true | \* - 255 |  |
| blkioDeviceOptions | map[blkioDeviceOption] | C/R/- | false | true | - |  |
| blkioWeight | int | C/R/- | false | true | - |  |
| build | dockerBuild | C/R/- | false | true | - |  |
| capAdd | array[enum] | C/R/- | false | true | - | enum options: ALL/AUDIT_CONTROL/AUDIT_WRITE/BLOCK_SUSPEND/CHOWN/DAC_OVERRIDE/DAC_READ_SEARCH/FOWNER/FSETID/IPC_LOCK/IPC_OWNER/KILL/LEASE/LINUX_IMMUTABLE/MAC_ADMIN/MAC_OVERRIDE/MKNOD/NET_ADMIN/NET_BIND_SERVICE/NET_BROADCAST/NET_RAW/SETFCAP/SETGID/SETPCAP/SETUID/SYSLOG/SYS_ADMIN/SYS_BOOT/SYS_CHROOT/SYS_MODULE/SYS_NICE/SYS_PACCT/SYS_PTRACE/SYS_RAWIO/SYS_RESOURCE/SYS_TIME/SYS_TTY_CONFIG/WAKE_ALARM;  |
| capDrop | array[enum] | C/R/- | false | true | - | enum options: ALL/AUDIT_CONTROL/AUDIT_WRITE/BLOCK_SUSPEND/CHOWN/DAC_OVERRIDE/DAC_READ_SEARCH/FOWNER/FSETID/IPC_LOCK/IPC_OWNER/KILL/LEASE/LINUX_IMMUTABLE/MAC_ADMIN/MAC_OVERRIDE/MKNOD/NET_ADMIN/NET_BIND_SERVICE/NET_BROADCAST/NET_RAW/SETFCAP/SETGID/SETPCAP/SETUID/SYSLOG/SYS_ADMIN/SYS_BOOT/SYS_CHROOT/SYS_MODULE/SYS_NICE/SYS_PACCT/SYS_PTRACE/SYS_RAWIO/SYS_RESOURCE/SYS_TIME/SYS_TTY_CONFIG/WAKE_ALARM;  |
| cgroupParent | string | C/R/- | false | true | - |  |
| command | array[string] | C/R/- | false | true | - |  |
| count | int | C/R/- | false | true | - |  |
| cpuCount | int | C/R/- | false | true | - |  |
| cpuPercent | int | C/R/- | false | true | - |  |
| cpuPeriod | int | C/R/- | false | true | - |  |
| cpuQuota | int | C/R/- | false | true | - |  |
| cpuRealtimePeriod | int | C/R/- | false | true | - |  |
| cpuRealtimeRuntime | int | C/R/- | false | true | - |  |
| cpuSet | string | C/R/- | false | true | - |  |
| cpuSetMems | string | C/R/- | false | true | - |  |
| cpuShares | int | C/R/- | false | true | - |  |
| createIndex | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| dataVolumeMounts | map[[volume](#volume)] | C/R/- | false | true | - |  |
| dataVolumes | array[string] | C/R/- | false | true | - |  |
| dataVolumesFrom | array[[container](#container)] | C/R/- | false | true | - |  |
| deploymentUnitUuid | string | -/R/- | false | true | \* - 128 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| devices | array[string] | C/R/- | false | true | - |  |
| diskQuota | int | C/R/- | false | true | - |  |
| dns | array[string] | C/R/- | false | true | - |  |
| dnsOpt | array[string] | C/R/- | false | true | - |  |
| dnsSearch | array[string] | C/R/- | false | true | - |  |
| domainName | string | C/R/- | false | true | - |  |
| entryPoint | array[string] | C/R/- | false | true | - |  |
| environment | map[string] | C/R/- | false | false | - |  |
| expose | array[string] | C/R/- | false | true | - |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| extraHosts | array[string] | C/R/- | false | true | - |  |
| firstRunning | date | -/R/- | false | true | \* - 255 |  |
| groupAdd | array[string] | C/R/- | false | true | - |  |
| healthCheck | instanceHealthCheck | C/R/- | false | true | - |  |
| healthCmd | array[string] | C/R/- | false | true | - |  |
| healthInterval | int | C/R/- | false | true | - |  |
| healthRetries | int | C/R/- | false | true | - |  |
| healthState | enum | -/R/- | false | true | \* - 128 | enum options: healthy/unhealthy/updating-healthy/updating-unhealthy/initializing;  |
| healthTimeout | int | C/R/- | false | true | - |  |
| hostId | [host](#host) | -/R/- | false | true | - |  |
| hostname | string | C/R/- | false | true | 1 - 255 |  |
| id | int | -/R/- | false | true | - |  |
| imageUuid | string | C/R/- | false | true | - |  |
| instanceLinks | map[[instance](#instance)] | C/R/- | false | true | - |  |
| instanceTriggeredStop | enum | C/R/- | false | true | \* - 128 | enum options: stop/remove;  |
| ioMaximumBandwidth | int | C/R/- | false | true | - |  |
| ioMaximumIOps | int | C/R/- | false | true | - |  |
| ip | string | C/R/- | false | true | - |  |
| ip6 | string | C/R/- | false | true | - |  |
| ipcMode | string | C/R/- | false | true | - |  |
| isolation | string | C/R/- | false | true | - |  |
| kernelMemory | int | C/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| labels | map[string] | C/R/- | false | false | - |  |
| logConfig | logConfig | C/R/- | false | true | - |  |
| lxcConf | map[string] | C/R/- | false | true | - |  |
| memory | int | C/R/- | false | true | - |  |
| memoryReservation | int | C/R/- | false | true | \* - 255 |  |
| memorySwap | int | C/R/- | false | true | - |  |
| memorySwappiness | int | C/R/- | false | true | - |  |
| milliCpuReservation | int | C/R/- | false | true | \* - 255 |  |
| mounts | array[mountEntry] | -/R/- | false | false | - |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| nativeContainer | boolean | -/R/- | false | false | \* - 255 |  |
| netAlias | array[string] | C/R/- | false | true | - |  |
| networkContainerId | [container](#container) | C/R/- | false | true | \* - 255 |  |
| networkIds | array[[network](#network)] | C/R/- | false | true | - |  |
| networkMode | string | C/R/- | false | true | - |  |
| oomKillDisable | boolean | C/R/- | false | true | - |  |
| oomScoreAdj | int | C/R/- | false | true | - |  |
| pidMode | enum | C/R/- | false | true | - | enum options: host;  |
| pidsLimit | int | C/R/- | false | true | - |  |
| ports | array[string] | C/R/- | false | true | - |  |
| primaryIpAddress | string | -/R/- | false | false | - |  |
| primaryNetworkId | [network](#network) | -/R/- | false | false | - |  |
| privileged | boolean | C/R/- | false | false | - |  |
| publishAllPorts | boolean | C/R/- | false | false | - |  |
| readOnly | boolean | C/R/- | false | false | - |  |
| registryCredentialId | [registryCredential](#registryCredential) | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| requestedHostId | [host](#host) | C/R/- | false | true | - |  |
| restartPolicy | restartPolicy | C/R/- | false | true | - |  |
| runInit | boolean | C/R/- | false | false | - |  |
| secrets | array[secretReference] | C/R/- | false | false | - |  |
| securityOpt | array[string] | C/R/- | false | true | - |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| serviceIds | array[[service](#service)] | -/R/- | false | true | - |  |
| shmSize | int | C/R/- | false | true | - |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startCount | int | -/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: creating/error/erroring/migrating/purged/purging/removed/removing/requested/restarting/running/starting/stopped/stopping/updating-running/updating-stopped;  |
| stdinOpen | boolean | C/R/- | false | false | - |  |
| stopSignal | string | C/R/- | false | true | - |  |
| stopTimeout | int | C/R/- | false | true | - |  |
| storageOpt | map[string] | C/R/- | false | true | - |  |
| sysctls | map[string] | C/R/- | false | true | - |  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| tmpfs | map[string] | C/R/- | false | true | - |  |
| token | string | -/R/- | false | true | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| tty | boolean | C/R/- | false | false | - |  |
| ulimits | array[ulimit] | C/R/- | false | true | - |  |
| user | string | C/R/- | false | true | 1 - \* |  |
| userPorts | array[string] | -/R/- | false | true | - |  |
| usernsMode | string | C/R/- | false | true | - |  |
| uts | string | C/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| version | string | -/R/- | false | false | \* - 255 |  |
| volumeDriver | string | C/R/- | false | true | - |  |
| workingDir | string | C/R/- | false | true | 1 - \* |  |

### containerEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| dockerInspect | json | -/R/- | false | false | - |  |
| externalFrom | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 255 |  |
| externalStatus | string | -/R/- | false | true | \* - 255 |  |
| externalTimestamp | int | -/R/- | false | true | \* - 255 |  |
| hostId | [host](#host) | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| reportedHostUuid | string | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |

### containerExec
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| attachStdin | boolean | C/R/- | false | false | - |  |
| attachStdout | boolean | C/R/- | false | false | - |  |
| command | array[string] | C/R/- | true | false | - |  |
| tty | boolean | C/R/- | false | false | - |  |

### containerLogs
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| follow | boolean | C/R/- | false | false | - |  |
| lines | int | C/R/- | false | true | - |  |
| since | string | C/R/- | false | false | - |  |
| timestamps | boolean | C/R/- | false | true | - |  |

### containerProxy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| port | int | C/R/- | false | true | - |  |
| scheme | enum | C/R/- | false | false | - | enum options: http/https;  |

### credential
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| publicValue | string | C/R/- | false | true | \* - 4096 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| secretValue | password | C/R/- | false | true | \* - 4096 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### databasechangelog
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| author | string | -/R/- | false | true | \* - 255 |  |
| comments | string | -/R/- | false | true | \* - 255 |  |
| dateexecuted | date | -/R/- | false | false | \* - 255 |  |
| description | string | -/R/- | false | true | \* - 255 |  |
| exectype | string | -/R/- | false | true | \* - 10 |  |
| filename | string | -/R/- | false | true | \* - 255 |  |
| id | string | -/R/- | false | false | - |  |
| liquibase | string | -/R/- | false | true | \* - 20 |  |
| md5sum | string | -/R/- | false | true | \* - 35 |  |
| orderexecuted | int | -/R/- | false | false | \* - 255 |  |
| tag | string | -/R/- | false | true | \* - 255 |  |

### databasechangeloglock
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| id | int | -/R/- | false | true | - |  |
| locked | boolean | -/R/- | false | false | \* - 255 |  |
| lockedby | string | -/R/- | false | true | \* - 255 |  |
| lockgranted | date | -/R/- | false | true | \* - 255 |  |

### defaultNetwork
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| defaultPolicyAction | enum | C/R/U | false | false | - | enum options: allow/deny;  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| dns | array[string] | C/R/- | false | false | - |  |
| dnsSearch | array[string] | C/R/- | false | false | - |  |
| hostPorts | boolean | C/R/- | false | false | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | C/R/- | false | true | \* - 255 |  |
| metadata | map[json] | C/R/U | false | false | - |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| policy | array[networkPolicyRule] | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| subnets | array[subnet] | C/R/- | false | false | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### digitaloceanConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accessToken | string | C/R/- | false | false | - | description: Digital Ocean access token;  |
| backups | boolean | C/R/- | false | false | - | description: enable backups for droplet;  |
| image | string | C/R/- | false | false | - | description: Digital Ocean Image;  |
| ipv6 | boolean | C/R/- | false | false | - | description: enable ipv6 for droplet;  |
| monitoring | boolean | C/R/- | false | false | - | description: enable monitoring for droplet;  |
| privateNetworking | boolean | C/R/- | false | false | - | description: enable private networking for droplet;  |
| region | string | C/R/- | false | false | - | description: Digital Ocean region;  |
| size | string | C/R/- | false | false | - | description: Digital Ocean size;  |
| sshKeyFingerprint | string | C/R/- | false | false | - | description: SSH key fingerprint;  |
| sshPort | string | C/R/- | false | false | - | description: SSH port;  |
| sshUser | string | C/R/- | false | false | - | description: SSH username;  |
| tags | string | C/R/- | false | false | - | description: comma-separated list of tags to apply to the Droplet;  |

### dnsService
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| assignServiceIpAddress | boolean | C/R/- | false | false | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| fqdn | string | -/R/- | false | true | - |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| launchConfig | launchConfig | C/R/- | false | true | - |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| metadata | map[json] | C/R/U | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| retainIp | boolean | C/R/U | false | true | - |  |
| selectorLink | string | C/R/U | false | true | \* - 4096 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| upgrade | serviceUpgrade | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### dockerBuild
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| context | string | C/R/- | false | true | - |  |
| dockerfile | string | C/R/- | false | true | - |  |
| forcerm | boolean | C/R/- | false | false | - |  |
| nocache | boolean | C/R/- | false | false | - |  |
| remote | string | C/R/- | false | true | - |  |
| rm | boolean | C/R/- | false | false | - |  |

### extensionImplementation
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| className | string | -/R/- | false | false | - |  |
| name | string | -/R/- | false | false | - |  |
| properties | map[string] | -/R/- | false | false | - |  |

### extensionPoint
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| excludeSetting | string | -/R/- | false | false | - |  |
| implementations | array[extensionImplementation] | -/R/- | false | false | - |  |
| includeSetting | string | -/R/- | false | false | - |  |
| listSetting | string | -/R/- | false | false | - |  |
| name | string | C/R/U | false | false | - |  |

### externalCredential
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |

### externalDnsEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| eventType | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 255 |  |
| fqdn | string | -/R/- | false | false | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| reportedAccountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| serviceName | string | -/R/- | false | false | - |  |
| stackName | string | -/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| eventType | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| reportedAccountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalHandler
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| priority | int | C/R/U | false | true | \* - 255 |  |
| processConfigs | array[externalHandlerProcessConfig] | C/R/- | true | false | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| retries | int | C/R/U | false | true | 1 - \* |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| timeoutMillis | int | C/R/U | false | true | 1000 - \* |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalHandlerExternalHandlerProcessMap
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| eventName | string | -/R/- | false | true | \* - 255 |  |
| externalHandlerId | [externalHandler](#externalHandler) | -/R/- | false | true | \* - 255 |  |
| externalHandlerProcessId | [externalHandlerProcess](#externalHandlerProcess) | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| onError | string | -/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalHandlerProcess
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalHandlerProcessConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| name | string | C/R/- | false | false | - |  |
| onError | string | C/R/- | false | false | - |  |

### externalHostEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| deleteHost | boolean | C/R/- | false | false | - |  |
| eventType | string | C/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 255 |  |
| hostId | [host](#host) | C/R/- | false | true | - |  |
| hostLabel | string | C/R/- | false | true | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| reportedAccountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalService
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| externalIpAddresses | array[string] | C/R/U | false | true | - |  |
| fqdn | string | -/R/- | false | true | - |  |
| healthCheck | instanceHealthCheck | C/R/- | false | true | - |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| hostname | string | C/R/U | false | true | - |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| launchConfig | launchConfig | C/R/- | false | true | - |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| metadata | map[json] | C/R/U | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| upgrade | serviceUpgrade | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalServiceEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| environment | json | -/R/- | false | false | - |  |
| eventType | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| reportedAccountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| service | json | -/R/- | true | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalStoragePoolEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| eventType | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 255 |  |
| hostUuids | array[string] | -/R/- | false | false | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| reportedAccountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| storagePool | storagePool | -/R/- | true | false | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### externalVolumeEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| eventType | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| reportedAccountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volume | volume | -/R/- | true | false | - |  |

### fieldDocumentation
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| description | string | C/R/U | false | false | - |  |

### genericObject
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| key | string | C/R/- | false | true | \* - 255 |  |
| kind | string | C/R/- | false | true | \* - 255 |  |
| name | string | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| resourceData | map[json] | C/R/U | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### haConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| clusterSize | int | -/R/- | false | true | - |  |
| dbHost | string | -/R/- | false | false | - |  |
| dbSize | int | -/R/- | false | true | - |  |
| enabled | boolean | -/R/U | false | false | - |  |

### haConfigInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| cert | string | C/R/- | false | true | - |  |
| certChain | string | C/R/- | false | true | - |  |
| clusterSize | int | C/R/- | false | false | 1 - 65535 |  |
| hostRegistrationUrl | string | C/R/- | true | false | - |  |
| httpEnabled | boolean | C/R/- | false | false | - |  |
| httpPort | int | C/R/- | false | true | 1 - 65535 |  |
| httpsPort | int | C/R/- | false | false | 1 - 65535 |  |
| key | string | C/R/- | false | true | - |  |
| ppHttpPort | int | C/R/- | false | true | 1 - 65535 |  |
| ppHttpsPort | int | C/R/- | false | false | 1 - 65535 |  |
| redisPort | int | C/R/- | false | false | 1 - 65535 |  |
| swarmEnabled | boolean | C/R/- | false | false | - |  |
| swarmPort | int | C/R/- | false | true | 1 - 65535 |  |
| zookeeperClientPort | int | C/R/- | false | false | 1 - 65535 |  |
| zookeeperLeaderPort | int | C/R/- | false | false | 1 - 65535 |  |
| zookeeperQuorumPort | int | C/R/- | false | false | 1 - 65535 |  |

### healthcheckInstanceHostMap
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| hostId | [host](#host) | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### host
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| agentId | [agent](#agent) | -/R/- | false | true | \* - 255 |  |
| agentIpAddress | string | -/R/- | false | true | - |  |
| agentState | string | -/R/- | false | true | \* - 128 |  |
| amazonec2Config | amazonec2Config | -/R/- | false | true | - |  |
| apiProxy | string | -/R/U | false | true | - |  |
| authCertificateAuthority | string | -/R/- | false | true | - |  |
| authKey | string | -/R/- | false | true | - |  |
| azureConfig | azureConfig | -/R/- | false | true | - |  |
| computeTotal | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| digitaloceanConfig | digitaloceanConfig | -/R/- | false | true | - |  |
| dockerVersion | string | -/R/- | false | true | - |  |
| driver | string | -/R/- | false | true | - |  |
| engineEnv | map[string] | -/R/- | false | true | - |  |
| engineInsecureRegistry | array[string] | -/R/- | false | true | - |  |
| engineInstallUrl | string | -/R/- | false | true | - |  |
| engineLabel | map[string] | -/R/- | false | true | - |  |
| engineOpt | map[string] | -/R/- | false | true | - |  |
| engineRegistryMirror | array[string] | -/R/- | false | true | - |  |
| engineStorageDriver | string | -/R/- | false | true | - |  |
| hostTemplateId | [hostTemplate](#hostTemplate) | -/R/- | false | true | - |  |
| hostname | string | C/R/- | true | false | - |  |
| id | int | -/R/- | false | true | - |  |
| info | json | -/R/- | false | false | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | false | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| labels | map[string] | -/R/- | false | true | - |  |
| localStorageMb | int | C/R/U | false | true | \* - 255 |  |
| memory | int | C/R/U | false | true | \* - 255 |  |
| milliCpu | int | C/R/U | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | - |  |
| packetConfig | packetConfig | -/R/- | false | true | - |  |
| physicalHostId | [physicalHost](#physicalHost) | -/R/- | false | true | \* - 255 |  |
| publicEndpoints | array[publicEndpoint] | -/R/- | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/error/erroring/inactive/provisioned/provisioning/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### hostAccess
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| token | string | -/R/- | false | false | - |  |
| url | string | -/R/- | false | false | - |  |

### hostApiProxyToken
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| reportedUuid | string | -/R/- | false | false | - |  |
| token | string | -/R/- | false | false | - |  |
| url | string | -/R/- | false | false | - |  |

### hostTemplate
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| driver | string | C/R/- | false | true | \* - 255 |  |
| flavorPrefix | string | C/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| publicValues | map[json] | C/R/- | false | false | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| secretValues | map[json] | C/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### identity
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| all | string | C/R/- | false | true | - |  |
| externalId | string | C/R/- | true | false | - |  |
| externalIdType | string | C/R/- | true | false | - |  |
| id | string | -/R/- | false | true | - |  |
| login | string | C/R/- | false | true | - |  |
| name | string | C/R/- | false | true | - |  |
| profilePicture | string | C/R/- | false | true | - |  |
| profileUrl | string | C/R/- | false | true | - |  |
| projectId | [project](#project) | C/R/- | false | true | - |  |
| role | string | C/R/- | false | true | - |  |
| user | boolean | -/R/- | false | true | - |  |

### image
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### inServiceUpgradeStrategy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| batchSize | int | C/R/- | false | true | 1 - \* |  |
| intervalMillis | int | C/R/- | false | true | 100 - \* |  |
| launchConfig | launchConfig | C/R/- | false | true | - |  |
| previousLaunchConfig | launchConfig | -/R/- | false | true | - |  |
| previousSecondaryLaunchConfigs | array[secondaryLaunchConfig] | -/R/- | false | false | - |  |
| secondaryLaunchConfigs | array[secondaryLaunchConfig] | C/R/- | false | false | - |  |
| startFirst | boolean | C/R/- | false | false | - |  |

### instance
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| hostId | [host](#host) | -/R/- | false | true | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: creating/error/erroring/migrating/purged/purging/removed/removing/requested/restarting/running/starting/stopped/stopping/updating-running/updating-stopped;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### instanceConsole
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| kind | string | -/R/- | false | false | - |  |
| password | string | -/R/- | false | false | - |  |
| url | string | -/R/- | false | false | - |  |

### instanceConsoleInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |

### instanceHealthCheck
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| healthyThreshold | int | C/R/- | false | true | - |  |
| initializingTimeout | int | C/R/- | false | true | - |  |
| interval | int | C/R/- | false | true | - |  |
| name | string | C/R/U | false | true | - |  |
| port | int | C/R/- | true | true | 1 - 65535 |  |
| recreateOnQuorumStrategyConfig | recreateOnQuorumStrategyConfig | C/R/- | false | true | - |  |
| reinitializingTimeout | int | C/R/- | false | true | - |  |
| requestLine | string | C/R/- | false | true | - |  |
| responseTimeout | int | C/R/- | false | true | - |  |
| strategy | enum | C/R/- | false | true | - | enum options: none/recreate/recreateOnQuorum;  |
| unhealthyThreshold | int | C/R/- | false | true | - |  |

### instanceLink
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| linkName | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| ports | array[json] | -/R/- | false | false | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| targetInstanceId | [instance](#instance) | -/R/U | false | true | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### instanceStop
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| remove | boolean | C/R/- | false | false | - |  |
| timeout | int | C/R/- | false | false | - |  |

### ipAddress
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| address | string | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| networkId | [network](#network) | -/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/associated/associating/deactivating/disassociating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### kubernetesService
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| selectorContainer | string | -/R/- | false | true | \* - 4096 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| template | json | -/R/- | false | false | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vip | string | -/R/- | false | true | \* - 255 |  |

### kubernetesStack
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| answers | map[json] | C/R/- | false | false | - |  |
| binding | binding | C/R/U | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| environment | map[string] | C/R/- | false | false | - |  |
| externalId | string | C/R/U | false | true | \* - 128 |  |
| group | string | C/R/U | false | true | \* - 255 |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/- | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| namespace | string | C/R/- | true | false | - |  |
| previousEnvironment | map[string] | C/R/U | false | true | - |  |
| previousExternalId | string | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serviceIds | array[[service](#service)] | -/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/error/erroring/finishing-upgrade/removed/removing/requested/rolling-back/updating-active/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| templates | map[string] | C/R/- | false | true | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### kubernetesStackUpgrade
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| answers | map[json] | C/R/- | false | false | - |  |
| environment | map[string] | C/R/- | false | false | - |  |
| externalId | string | C/R/- | false | false | - |  |
| templates | map[string] | C/R/- | false | false | - |  |

### label
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| key | string | -/R/- | false | true | \* - 1024 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| value | string | -/R/- | false | true | \* - 4096 |  |

### launchConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| agentId | [agent](#agent) | -/R/- | false | true | \* - 255 |  |
| allocationState | string | -/R/- | false | true | \* - 255 |  |
| blkioDeviceOptions | map[blkioDeviceOption] | C/R/- | false | true | - |  |
| blkioWeight | int | C/R/- | false | true | - |  |
| build | dockerBuild | C/R/- | false | true | - |  |
| capAdd | array[enum] | C/R/- | false | true | - | enum options: ALL/AUDIT_CONTROL/AUDIT_WRITE/BLOCK_SUSPEND/CHOWN/DAC_OVERRIDE/DAC_READ_SEARCH/FOWNER/FSETID/IPC_LOCK/IPC_OWNER/KILL/LEASE/LINUX_IMMUTABLE/MAC_ADMIN/MAC_OVERRIDE/MKNOD/NET_ADMIN/NET_BIND_SERVICE/NET_BROADCAST/NET_RAW/SETFCAP/SETGID/SETPCAP/SETUID/SYSLOG/SYS_ADMIN/SYS_BOOT/SYS_CHROOT/SYS_MODULE/SYS_NICE/SYS_PACCT/SYS_PTRACE/SYS_RAWIO/SYS_RESOURCE/SYS_TIME/SYS_TTY_CONFIG/WAKE_ALARM;  |
| capDrop | array[enum] | C/R/- | false | true | - | enum options: ALL/AUDIT_CONTROL/AUDIT_WRITE/BLOCK_SUSPEND/CHOWN/DAC_OVERRIDE/DAC_READ_SEARCH/FOWNER/FSETID/IPC_LOCK/IPC_OWNER/KILL/LEASE/LINUX_IMMUTABLE/MAC_ADMIN/MAC_OVERRIDE/MKNOD/NET_ADMIN/NET_BIND_SERVICE/NET_BROADCAST/NET_RAW/SETFCAP/SETGID/SETPCAP/SETUID/SYSLOG/SYS_ADMIN/SYS_BOOT/SYS_CHROOT/SYS_MODULE/SYS_NICE/SYS_PACCT/SYS_PTRACE/SYS_RAWIO/SYS_RESOURCE/SYS_TIME/SYS_TTY_CONFIG/WAKE_ALARM;  |
| cgroupParent | string | C/R/- | false | true | - |  |
| command | array[string] | C/R/- | false | true | - |  |
| count | int | C/R/- | false | true | - |  |
| cpuCount | int | C/R/- | false | true | - |  |
| cpuPercent | int | C/R/- | false | true | - |  |
| cpuPeriod | int | C/R/- | false | true | - |  |
| cpuQuota | int | C/R/- | false | true | - |  |
| cpuRealtimePeriod | int | C/R/- | false | true | - |  |
| cpuRealtimeRuntime | int | C/R/- | false | true | - |  |
| cpuSet | string | C/R/- | false | true | - |  |
| cpuSetMems | string | C/R/- | false | true | - |  |
| cpuShares | int | C/R/- | false | true | - |  |
| createIndex | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| dataVolumeMounts | map[[volume](#volume)] | C/R/- | false | true | - |  |
| dataVolumes | array[string] | C/R/- | false | true | - |  |
| dataVolumesFrom | array[[container](#container)] | C/R/- | false | true | - |  |
| dataVolumesFromLaunchConfigs | array[string] | C/R/- | false | true | - |  |
| deploymentUnitUuid | string | -/R/- | false | true | \* - 128 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| devices | array[string] | C/R/- | false | true | - |  |
| diskQuota | int | C/R/- | false | true | - |  |
| disks | array[virtualMachineDisk] | C/R/- | false | true | - |  |
| dns | array[string] | C/R/- | false | true | - |  |
| dnsOpt | array[string] | C/R/- | false | true | - |  |
| dnsSearch | array[string] | C/R/- | false | true | - |  |
| domainName | string | C/R/- | false | true | - |  |
| drainTimeoutMs | int | C/R/U | false | false | - |  |
| entryPoint | array[string] | C/R/- | false | true | - |  |
| environment | map[string] | C/R/- | false | false | - |  |
| expose | array[string] | C/R/- | false | true | - |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| extraHosts | array[string] | C/R/- | false | true | - |  |
| firstRunning | date | -/R/- | false | true | \* - 255 |  |
| groupAdd | array[string] | C/R/- | false | true | - |  |
| healthCheck | instanceHealthCheck | C/R/- | false | true | - |  |
| healthCmd | array[string] | C/R/- | false | true | - |  |
| healthInterval | int | C/R/- | false | true | - |  |
| healthRetries | int | C/R/- | false | true | - |  |
| healthState | enum | -/R/- | false | true | \* - 128 | enum options: healthy/unhealthy/updating-healthy/updating-unhealthy/initializing;  |
| healthTimeout | int | C/R/- | false | true | - |  |
| hostId | [host](#host) | -/R/- | false | true | - |  |
| hostname | string | C/R/- | false | true | 1 - 255 |  |
| id | int | -/R/- | false | true | - |  |
| imageUuid | string | C/R/- | false | true | - |  |
| instanceLinks | map[[instance](#instance)] | C/R/- | false | true | - |  |
| instanceTriggeredStop | enum | C/R/- | false | true | \* - 128 | enum options: stop/remove;  |
| ioMaximumBandwidth | int | C/R/- | false | true | - |  |
| ioMaximumIOps | int | C/R/- | false | true | - |  |
| ip | string | C/R/- | false | true | - |  |
| ip6 | string | C/R/- | false | true | - |  |
| ipcMode | string | C/R/- | false | true | - |  |
| isolation | string | C/R/- | false | true | - |  |
| kernelMemory | int | C/R/- | false | true | - |  |
| kind | enum | C/R/- | false | true | \* - 255 | enum options: container/virtualMachine;  |
| labels | map[string] | C/R/- | false | false | - |  |
| logConfig | logConfig | C/R/- | false | true | - |  |
| lxcConf | map[string] | C/R/- | false | true | - |  |
| memory | int | C/R/- | false | true | - |  |
| memoryMb | int | C/R/- | false | true | \* - 255 |  |
| memoryReservation | int | C/R/- | false | true | \* - 255 |  |
| memorySwap | int | C/R/- | false | true | - |  |
| memorySwappiness | int | C/R/- | false | true | - |  |
| milliCpuReservation | int | C/R/- | false | true | \* - 255 |  |
| mounts | array[mountEntry] | -/R/- | false | false | - |  |
| nativeContainer | boolean | -/R/- | false | false | \* - 255 |  |
| netAlias | array[string] | C/R/- | false | true | - |  |
| networkContainerId | [container](#container) | C/R/- | false | true | \* - 255 |  |
| networkIds | array[[network](#network)] | C/R/- | false | true | - |  |
| networkLaunchConfig | string | C/R/- | false | true | - |  |
| networkMode | string | C/R/- | false | true | - |  |
| oomKillDisable | boolean | C/R/- | false | true | - |  |
| oomScoreAdj | int | C/R/- | false | true | - |  |
| pidMode | enum | C/R/- | false | true | - | enum options: host;  |
| pidsLimit | int | C/R/- | false | true | - |  |
| ports | array[string] | C/R/U | false | true | - |  |
| primaryIpAddress | string | -/R/- | false | false | - |  |
| primaryNetworkId | [network](#network) | -/R/- | false | false | - |  |
| privileged | boolean | C/R/- | false | false | - |  |
| publishAllPorts | boolean | C/R/- | false | false | - |  |
| readOnly | boolean | C/R/- | false | false | - |  |
| registryCredentialId | [registryCredential](#registryCredential) | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| requestedHostId | [host](#host) | C/R/- | false | true | - |  |
| requestedIpAddress | string | C/R/- | false | true | - |  |
| runInit | boolean | C/R/- | false | false | - |  |
| secrets | array[secretReference] | C/R/- | false | false | - |  |
| securityOpt | array[string] | C/R/- | false | true | - |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| serviceIds | array[[service](#service)] | -/R/- | false | true | - |  |
| shmSize | int | C/R/- | false | true | - |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startCount | int | -/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: creating/error/erroring/migrating/purged/purging/removed/removing/requested/restarting/running/starting/stopped/stopping/updating-running/updating-stopped;  |
| stdinOpen | boolean | C/R/- | false | false | - |  |
| stopSignal | string | C/R/- | false | true | - |  |
| stopTimeout | int | C/R/- | false | true | - |  |
| storageOpt | map[string] | C/R/- | false | true | - |  |
| sysctls | map[string] | C/R/- | false | true | - |  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| tmpfs | map[string] | C/R/- | false | true | - |  |
| token | string | -/R/- | false | true | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| tty | boolean | C/R/- | false | false | - |  |
| ulimits | array[ulimit] | C/R/- | false | true | - |  |
| user | string | C/R/- | false | true | 1 - \* |  |
| userPorts | array[string] | -/R/- | false | true | - |  |
| userdata | string | C/R/- | false | true | \* - 65535 |  |
| usernsMode | string | C/R/- | false | true | - |  |
| uts | string | C/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vcpu | int | C/R/- | false | true | - |  |
| version | string | -/R/- | false | false | \* - 255 |  |
| volumeDriver | string | C/R/- | false | true | - |  |
| workingDir | string | C/R/- | false | true | 1 - \* |  |

### lbConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| certificateIds | array[[certificate](#certificate)] | C/R/- | false | true | - |  |
| config | string | C/R/- | false | true | - |  |
| defaultCertificateId | [certificate](#certificate) | C/R/- | false | true | - |  |
| portRules | array[portRule] | C/R/- | false | false | - |  |
| stickinessPolicy | loadBalancerCookieStickinessPolicy | C/R/- | false | true | - |  |

### lbTargetConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| portRules | array[targetPortRule] | C/R/U | false | false | - |  |

### loadBalancerCookieStickinessPolicy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| cookie | string | C/R/U | false | false | - |  |
| domain | string | C/R/U | false | false | - |  |
| indirect | boolean | C/R/U | false | true | - |  |
| mode | enum | C/R/U | false | true | - | enum options: rewrite/insert/prefix;  |
| name | string | C/R/U | false | true | - |  |
| nocache | boolean | C/R/U | false | true | - |  |
| postonly | boolean | C/R/U | false | true | - |  |

### loadBalancerService
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| assignServiceIpAddress | boolean | C/R/- | false | false | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| currentScale | int | -/R/- | false | false | - |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| fqdn | string | -/R/- | false | true | - |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| launchConfig | launchConfig | C/R/U | true | true | - |  |
| lbConfig | lbConfig | C/R/U | true | true | - |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| metadata | map[json] | C/R/U | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| publicEndpoints | array[publicEndpoint] | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| retainIp | boolean | C/R/U | false | true | - |  |
| scale | int | C/R/U | false | false | - |  |
| scalePolicy | scalePolicy | C/R/U | false | true | - |  |
| selectorLink | string | C/R/U | false | true | \* - 4096 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| upgrade | serviceUpgrade | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vip | string | C/R/- | false | true | \* - 255 |  |

### localAuthConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accessMode | string | C/R/- | false | true | - |  |
| enabled | boolean | C/R/- | false | true | - |  |
| name | string | C/R/- | false | true | - |  |
| password | password | C/R/- | true | false | - |  |
| username | string | C/R/- | true | false | - |  |

### logConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| config | map[string] | C/R/- | false | true | - |  |
| driver | string | C/R/- | false | true | - |  |

### machine
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| amazonec2Config | amazonec2Config | -/R/- | false | true | - |  |
| authCertificateAuthority | string | -/R/- | false | true | - |  |
| authKey | string | -/R/- | false | true | - |  |
| azureConfig | azureConfig | -/R/- | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| digitaloceanConfig | digitaloceanConfig | -/R/- | false | true | - |  |
| dockerVersion | string | -/R/- | false | true | - |  |
| driver | string | -/R/- | false | true | - |  |
| engineEnv | map[string] | -/R/- | false | true | - |  |
| engineInsecureRegistry | array[string] | -/R/- | false | true | - |  |
| engineInstallUrl | string | -/R/- | false | true | - |  |
| engineLabel | map[string] | -/R/- | false | true | - |  |
| engineOpt | map[string] | -/R/- | false | true | - |  |
| engineRegistryMirror | array[string] | -/R/- | false | true | - |  |
| engineStorageDriver | string | -/R/- | false | true | - |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| hostTemplateId | [hostTemplate](#hostTemplate) | -/R/- | false | true | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| labels | map[string] | -/R/- | false | true | - |  |
| name | string | C/R/U | false | true | - |  |
| packetConfig | packetConfig | -/R/- | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: active/bootstrapping/created/creating/error/erroring/removed/removing/requested/updating;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### machineDriver
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| activateOnCreate | boolean | C/R/- | false | false | - |  |
| builtin | boolean | C/R/- | false | false | - |  |
| checksum | string | C/R/U | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| defaultActive | boolean | -/R/- | false | false | - |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/U | false | true | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | -/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/error/erroring/inactive/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uiUrl | string | C/R/U | false | true | - |  |
| url | string | C/R/U | true | false | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### mount
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| path | string | -/R/- | false | true | \* - 512 |  |
| permissions | string | -/R/- | false | true | \* - 128 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volumeId | [volume](#volume) | -/R/- | false | true | \* - 255 |  |

### mountEntry
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| instanceId | [instance](#instance) | -/R/- | false | false | - |  |
| instanceName | string | -/R/- | false | false | - |  |
| path | string | -/R/- | false | false | - |  |
| volumeId | [volume](#volume) | -/R/- | false | false | - |  |
| volumeName | string | -/R/- | false | false | - |  |

### network
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| defaultPolicyAction | enum | C/R/U | false | false | - | enum options: allow/deny;  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| dns | array[string] | C/R/- | false | false | - |  |
| dnsSearch | array[string] | C/R/- | false | false | - |  |
| hostPorts | boolean | C/R/- | false | false | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| metadata | map[json] | C/R/U | false | false | - |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| networkDriverId | [networkDriver](#networkDriver) | C/R/- | true | false | \* - 255 |  |
| policy | array[networkPolicyRule] | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| subnets | array[subnet] | C/R/- | false | false | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### networkDriver
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| cniConfig | map[json] | C/R/- | false | false | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| defaultNetwork | defaultNetwork | C/R/- | false | false | - |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| networkMetadata | map[json] | C/R/- | false | false | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/creating/deactivating/inactive/removed/removing/requested/updating;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### networkDriverService
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| assignServiceIpAddress | boolean | C/R/- | false | false | - |  |
| createIndex | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| currentScale | int | -/R/- | false | false | - |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| fqdn | string | -/R/- | false | true | - |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| launchConfig | launchConfig | C/R/- | false | true | - |  |
| lbConfig | lbTargetConfig | C/R/U | false | true | - |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| metadata | map[json] | C/R/U | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| networkDriver | networkDriver | C/R/U | true | false | - |  |
| publicEndpoints | array[publicEndpoint] | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| retainIp | boolean | C/R/U | false | true | - |  |
| scale | int | C/R/U | false | false | - |  |
| scalePolicy | scalePolicy | C/R/U | false | true | - |  |
| secondaryLaunchConfigs | array[secondaryLaunchConfig] | C/R/- | false | false | - |  |
| selectorContainer | string | C/R/U | false | true | \* - 4096 |  |
| selectorLink | string | C/R/U | false | true | \* - 4096 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| upgrade | serviceUpgrade | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vip | string | C/R/- | false | true | \* - 255 |  |

### networkPolicyRule
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| action | enum | C/R/- | false | false | - | enum options: allow/deny;  |
| between | networkPolicyRuleBetween | C/R/- | false | false | - |  |
| from | networkPolicyRuleMember | C/R/- | false | false | - |  |
| ports | array[string] | C/R/- | false | false | - |  |
| to | networkPolicyRuleMember | C/R/- | false | false | - |  |
| within | enum | C/R/- | false | false | - | enum options: stack/service/linked;  |

### networkPolicyRuleBetween
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| groupBy | string | C/R/- | false | false | - |  |
| selector | string | C/R/- | false | false | - |  |

### networkPolicyRuleMember
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| selector | string | C/R/- | false | false | - |  |

### networkPolicyRuleWithin
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |

### nfsConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| mountOptions | string | C/R/- | false | true | - |  |
| server | string | C/R/- | true | false | - |  |
| share | string | C/R/- | true | false | - |  |

### openldapconfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accessMode | string | C/R/- | true | false | - |  |
| allowedIdentities | array[identity] | C/R/- | false | false | - |  |
| connectionTimeout | int | C/R/- | true | false | - |  |
| domain | string | C/R/- | true | false | 1 - \* |  |
| enabled | boolean | C/R/- | false | true | - |  |
| groupDNField | string | C/R/- | false | false | - |  |
| groupMemberMappingAttribute | string | C/R/- | true | false | - |  |
| groupMemberUserAttribute | string | C/R/- | false | false | - |  |
| groupNameField | string | C/R/- | true | false | - |  |
| groupObjectClass | string | C/R/- | true | false | - |  |
| groupSearchDomain | string | C/R/- | false | true | - |  |
| groupSearchField | string | C/R/- | true | false | - |  |
| loginDomain | string | C/R/- | false | true | - |  |
| name | string | -/R/- | false | false | - |  |
| port | int | C/R/- | true | true | - |  |
| server | string | C/R/- | true | false | 1 - \* |  |
| serviceAccountPassword | string | C/R/- | true | true | 1 - \* |  |
| serviceAccountUsername | string | C/R/- | true | true | 1 - \* |  |
| tls | boolean | C/R/- | true | true | - |  |
| userDisabledBitMask | int | C/R/- | false | true | - |  |
| userEnabledAttribute | string | C/R/- | false | true | - |  |
| userLoginField | string | C/R/- | true | false | - |  |
| userMemberAttribute | string | C/R/- | true | false | - |  |
| userNameField | string | C/R/- | true | false | - |  |
| userObjectClass | string | C/R/- | true | false | - |  |
| userSearchField | string | C/R/- | true | false | - |  |

### password
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | C/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| publicValue | string | C/R/- | false | true | \* - 4096 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| secretValue | password | C/R/- | false | true | \* - 4096 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### physicalHost
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| driver | string | -/R/- | false | true | \* - 128 |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: active/bootstrapping/created/creating/error/erroring/removed/removing/requested/updating;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### port
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| bindAddress | string | -/R/- | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| privateIpAddressId | [ipAddress](#ipAddress) | -/R/- | false | true | \* - 255 |  |
| privatePort | int | -/R/- | false | true | \* - 255 |  |
| protocol | string | -/R/- | false | true | \* - 128 |  |
| publicIpAddressId | [ipAddress](#ipAddress) | -/R/- | false | true | \* - 255 |  |
| publicPort | int | -/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### portRule
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| backendName | string | C/R/- | false | true | - |  |
| environment | string | C/R/- | false | true | - |  |
| hostname | string | C/R/- | false | true | - |  |
| path | string | C/R/- | false | true | - |  |
| priority | int | C/R/- | false | true | 1 - \* |  |
| protocol | enum | C/R/- | false | false | - | enum options: http/tcp/https/tls/sni/udp;  |
| region | string | C/R/- | false | true | - |  |
| selector | string | C/R/- | false | true | - |  |
| serviceId | [service](#service) | C/R/- | false | true | - |  |
| sourcePort | int | C/R/- | true | true | 1 - 65535 |  |
| targetPort | int | C/R/- | false | true | 1 - 65535 |  |
| weight | int | C/R/- | false | true | \* - 256 |  |

### processDefinition
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| extensionBased | boolean | -/R/- | false | false | - |  |
| id | string | -/R/- | false | false | - |  |
| name | string | C/R/U | false | false | - |  |
| postProcessListeners | extensionPoint | -/R/- | false | false | - |  |
| preProcessListeners | extensionPoint | -/R/- | false | false | - |  |
| processHandlers | extensionPoint | -/R/- | false | false | - |  |
| resourceType | string | -/R/- | false | false | - |  |
| stateTransitions | array[stateTransition] | -/R/- | false | false | - |  |

### processExecution
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| created | date | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| log | map[json] | -/R/- | false | true | \* - 16777215 |  |
| processInstanceId | [processInstance](#processInstance) | -/R/- | false | false | \* - 255 |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### processInstance
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| endTime | date | -/R/- | false | true | \* - 255 |  |
| executionCount | int | -/R/- | false | false | \* - 255 |  |
| exitReason | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| phase | string | -/R/- | false | true | \* - 128 |  |
| priority | int | -/R/- | false | true | \* - 255 |  |
| processName | string | -/R/- | false | true | \* - 128 |  |
| resourceId | string | -/R/- | false | true | \* - 128 |  |
| resourceType | string | -/R/- | false | true | \* - 128 |  |
| result | string | -/R/- | false | true | \* - 128 |  |
| runAfter | date | -/R/- | false | true | \* - 255 |  |
| runningProcessServerId | string | -/R/- | false | true | \* - 128 |  |
| startProcessServerId | string | -/R/- | false | true | \* - 128 |  |
| startTime | date | -/R/- | false | true | \* - 255 |  |

### processPool
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| activeTasks | int | -/R/- | false | false | - |  |
| completedTasks | int | -/R/- | false | false | - |  |
| maxPoolSize | int | -/R/- | false | false | - |  |
| minPoolSize | int | -/R/- | false | false | - |  |
| name | string | -/R/- | false | false | - |  |
| poolSize | int | -/R/- | false | false | - |  |
| queueRemainingCapacity | int | -/R/- | false | false | - |  |
| queueSize | int | -/R/- | false | false | - |  |
| rejectedTasks | int | -/R/- | false | false | - |  |

### processSummary
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| delay | int | -/R/- | false | false | - |  |
| processName | string | -/R/- | false | false | - |  |
| ready | int | -/R/- | false | false | - |  |
| running | int | -/R/- | false | false | - |  |

### project
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| allowSystemRole | boolean | C/R/U | false | false | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| defaultNetworkId | [network](#network) | -/R/- | false | true | \* - 255 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| hostRemoveDelaySeconds | int | C/R/U | false | true | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| members | array[projectMember] | C/R/- | false | true | - |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| orchestration | string | -/R/- | false | false | - |  |
| projectLinks | array[[project](#project)] | C/R/U | false | true | - |  |
| projectTemplateId | [projectTemplate](#projectTemplate) | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| servicesPortRange | servicesPortRange | C/R/U | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive/upgrading;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | C/R/- | false | true | \* - 128 |  |
| version | string | -/R/- | false | true | \* - 128 |  |
| virtualMachine | boolean | C/R/U | false | false | - |  |

### projectMember
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| externalIdType | enum | C/R/- | false | true | \* - 255 | enum options: rancher_id/openldap_user/openldap_group/azuread_user/azuread_group/github_user/github_org/github_team/shibboleth_user/shibboleth_group/ldap_user/ldap_group;  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | -/R/- | false | true | \* - 255 |  |
| projectId | [project](#project) | -/R/- | false | false | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| role | enum | C/R/- | false | true | \* - 255 | enum options: member/owner/readonly/restricted;  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### projectTemplate
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/U | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| isPublic | boolean | C/R/U | false | false | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| stacks | array[catalogTemplate] | C/R/U | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### publicEndpoint
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| hostId | [host](#host) | C/R/- | false | false | - |  |
| instanceId | [instance](#instance) | C/R/- | false | false | - |  |
| ipAddress | string | C/R/- | false | false | - |  |
| port | int | C/R/- | false | true | - |  |
| serviceId | [service](#service) | C/R/- | false | false | - |  |

### publish
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| data | map[json] | C/R/- | false | true | - |  |
| id | string | C/R/- | false | false | - |  |
| name | string | C/R/- | true | false | - |  |
| previousIds | array[string] | C/R/- | false | false | - |  |
| publisher | string | C/R/- | false | false | - |  |
| resourceId | string | C/R/- | false | true | - |  |
| resourceType | string | C/R/- | false | true | - |  |
| time | int | C/R/- | false | true | - |  |
| transitioning | string | C/R/- | false | false | - |  |
| transitioningInternalMessage | string | C/R/- | false | false | - |  |
| transitioningMessage | string | C/R/- | false | false | - |  |
| transitioningProgress | int | C/R/- | false | true | - |  |

### pullTask
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| image | string | C/R/- | true | false | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| labels | map[string] | C/R/- | false | false | - |  |
| mode | enum | C/R/- | true | false | - | enum options: all/cached;  |
| name | string | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested;  |
| status | map[string] | -/R/- | false | false | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### recreateOnQuorumStrategyConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| quorum | int | C/R/- | true | true | - |  |

### region
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| local | boolean | C/R/U | false | true | \* - 255 |  |
| name | string | C/R/- | true | false | \* - 255 |  |
| publicValue | string | C/R/U | true | false | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| secretValue | string | C/R/U | true | false | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| url | string | C/R/U | false | true | \* - 255 |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### register
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accessKey | string | -/R/- | false | false | - |  |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| key | string | C/R/- | true | true | \* - 255 |  |
| kind | string | C/R/- | false | true | \* - 255 |  |
| name | string | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| resourceData | map[json] | C/R/U | false | false | - |  |
| secretKey | string | -/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### registrationToken
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| command | string | -/R/- | false | false | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| image | string | -/R/- | false | false | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| registrationUrl | string | -/R/- | false | false | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| token | string | -/R/- | false | false | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### registry
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| blockDevicePath | string | -/R/- | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| driverName | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serverAddress | string | C/R/- | true | false | 1 - \* |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volumeAccessMode | string | -/R/- | false | true | \* - 255 |  |
| volumeCapabilities | array[string] | -/R/- | false | true | - |  |

### registryCredential
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| publicValue | string | C/R/U | false | true | \* - 4096 |  |
| registryId | [registry](#registry) | C/R/- | true | false | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| secretValue | password | C/R/U | false | true | \* - 4096 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### resourceDefinition
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| id | string | -/R/- | false | false | - |  |
| name | string | -/R/- | false | false | - |  |

### restartPolicy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| maximumRetryCount | int | C/R/- | false | false | - |  |
| name | string | C/R/U | false | false | - |  |

### restoreFromBackupInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| backupId | [backup](#backup) | C/R/- | true | false | - |  |

### revertToSnapshotInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| snapshotId | [snapshot](#snapshot) | C/R/- | true | false | - |  |

### rollingRestartStrategy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| batchSize | int | C/R/- | false | true | 1 - \* |  |
| intervalMillis | int | C/R/- | false | true | 100 - \* |  |

### scalePolicy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| increment | int | C/R/- | false | true | 1 - \* |  |
| max | int | C/R/- | false | true | 1 - \* |  |
| min | int | C/R/- | true | true | 1 - \* |  |

### scheduledUpgrade
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| stackId | [stack](#stack) | -/R/- | false | true | \* - 255 |  |
| started | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: done/removed/removing/requested/running/scheduled/scheduling;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### schema
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| collectionActions | map[json] | C/R/- | false | false | - |  |
| collectionFields | map[json] | C/R/- | false | false | - |  |
| collectionFilters | map[json] | C/R/- | false | false | - |  |
| collectionMethods | array[string] | C/R/- | false | false | - |  |
| includeableLinks | array[string] | C/R/- | false | false | - |  |
| pluralName | string | C/R/- | false | false | - |  |
| resourceActions | map[json] | C/R/- | false | false | - |  |
| resourceFields | map[json] | C/R/- | false | false | - |  |
| resourceMethods | array[string] | C/R/- | false | false | - |  |

### secondaryLaunchConfig
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| agentId | [agent](#agent) | -/R/- | false | true | \* - 255 |  |
| allocationState | string | -/R/- | false | true | \* - 255 |  |
| blkioDeviceOptions | map[blkioDeviceOption] | C/R/- | false | true | - |  |
| blkioWeight | int | C/R/- | false | true | - |  |
| build | dockerBuild | C/R/- | false | true | - |  |
| capAdd | array[enum] | C/R/- | false | true | - | enum options: ALL/AUDIT_CONTROL/AUDIT_WRITE/BLOCK_SUSPEND/CHOWN/DAC_OVERRIDE/DAC_READ_SEARCH/FOWNER/FSETID/IPC_LOCK/IPC_OWNER/KILL/LEASE/LINUX_IMMUTABLE/MAC_ADMIN/MAC_OVERRIDE/MKNOD/NET_ADMIN/NET_BIND_SERVICE/NET_BROADCAST/NET_RAW/SETFCAP/SETGID/SETPCAP/SETUID/SYSLOG/SYS_ADMIN/SYS_BOOT/SYS_CHROOT/SYS_MODULE/SYS_NICE/SYS_PACCT/SYS_PTRACE/SYS_RAWIO/SYS_RESOURCE/SYS_TIME/SYS_TTY_CONFIG/WAKE_ALARM;  |
| capDrop | array[enum] | C/R/- | false | true | - | enum options: ALL/AUDIT_CONTROL/AUDIT_WRITE/BLOCK_SUSPEND/CHOWN/DAC_OVERRIDE/DAC_READ_SEARCH/FOWNER/FSETID/IPC_LOCK/IPC_OWNER/KILL/LEASE/LINUX_IMMUTABLE/MAC_ADMIN/MAC_OVERRIDE/MKNOD/NET_ADMIN/NET_BIND_SERVICE/NET_BROADCAST/NET_RAW/SETFCAP/SETGID/SETPCAP/SETUID/SYSLOG/SYS_ADMIN/SYS_BOOT/SYS_CHROOT/SYS_MODULE/SYS_NICE/SYS_PACCT/SYS_PTRACE/SYS_RAWIO/SYS_RESOURCE/SYS_TIME/SYS_TTY_CONFIG/WAKE_ALARM;  |
| cgroupParent | string | C/R/- | false | true | - |  |
| command | array[string] | C/R/- | false | true | - |  |
| count | int | C/R/- | false | true | - |  |
| cpuCount | int | C/R/- | false | true | - |  |
| cpuPercent | int | C/R/- | false | true | - |  |
| cpuPeriod | int | C/R/- | false | true | - |  |
| cpuQuota | int | C/R/- | false | true | - |  |
| cpuRealtimePeriod | int | C/R/- | false | true | - |  |
| cpuRealtimeRuntime | int | C/R/- | false | true | - |  |
| cpuSet | string | C/R/- | false | true | - |  |
| cpuSetMems | string | C/R/- | false | true | - |  |
| cpuShares | int | C/R/- | false | true | - |  |
| createIndex | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| dataVolumeMounts | map[[volume](#volume)] | C/R/- | false | true | - |  |
| dataVolumes | array[string] | C/R/- | false | true | - |  |
| dataVolumesFrom | array[[container](#container)] | C/R/- | false | true | - |  |
| dataVolumesFromLaunchConfigs | array[string] | C/R/- | false | true | - |  |
| deploymentUnitUuid | string | -/R/- | false | true | \* - 128 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| devices | array[string] | C/R/- | false | true | - |  |
| diskQuota | int | C/R/- | false | true | - |  |
| disks | array[virtualMachineDisk] | C/R/- | false | true | - |  |
| dns | array[string] | C/R/- | false | true | - |  |
| dnsOpt | array[string] | C/R/- | false | true | - |  |
| dnsSearch | array[string] | C/R/- | false | true | - |  |
| domainName | string | C/R/- | false | true | - |  |
| entryPoint | array[string] | C/R/- | false | true | - |  |
| environment | map[string] | C/R/- | false | false | - |  |
| expose | array[string] | C/R/- | false | true | - |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| extraHosts | array[string] | C/R/- | false | true | - |  |
| firstRunning | date | -/R/- | false | true | \* - 255 |  |
| groupAdd | array[string] | C/R/- | false | true | - |  |
| healthCheck | instanceHealthCheck | C/R/- | false | true | - |  |
| healthCmd | array[string] | C/R/- | false | true | - |  |
| healthInterval | int | C/R/- | false | true | - |  |
| healthRetries | int | C/R/- | false | true | - |  |
| healthState | enum | -/R/- | false | true | \* - 128 | enum options: healthy/unhealthy/updating-healthy/updating-unhealthy/initializing;  |
| healthTimeout | int | C/R/- | false | true | - |  |
| hostId | [host](#host) | -/R/- | false | true | - |  |
| hostname | string | C/R/- | false | true | 1 - 255 |  |
| id | int | -/R/- | false | true | - |  |
| imageUuid | string | C/R/- | false | true | - |  |
| instanceLinks | map[[instance](#instance)] | C/R/- | false | true | - |  |
| instanceTriggeredStop | enum | C/R/- | false | true | \* - 128 | enum options: stop/remove;  |
| ioMaximumBandwidth | int | C/R/- | false | true | - |  |
| ioMaximumIOps | int | C/R/- | false | true | - |  |
| ip | string | C/R/- | false | true | - |  |
| ip6 | string | C/R/- | false | true | - |  |
| ipcMode | string | C/R/- | false | true | - |  |
| isolation | string | C/R/- | false | true | - |  |
| kernelMemory | int | C/R/- | false | true | - |  |
| kind | enum | C/R/- | false | true | \* - 255 | enum options: container/virtualMachine;  |
| labels | map[string] | C/R/- | false | false | - |  |
| logConfig | logConfig | C/R/- | false | true | - |  |
| lxcConf | map[string] | C/R/- | false | true | - |  |
| memory | int | C/R/- | false | true | - |  |
| memoryMb | int | C/R/- | false | true | \* - 255 |  |
| memoryReservation | int | C/R/- | false | true | \* - 255 |  |
| memorySwap | int | C/R/- | false | true | - |  |
| memorySwappiness | int | C/R/- | false | true | - |  |
| milliCpuReservation | int | C/R/- | false | true | \* - 255 |  |
| mounts | array[mountEntry] | -/R/- | false | false | - |  |
| name | string | C/R/- | true | false | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| nativeContainer | boolean | -/R/- | false | false | \* - 255 |  |
| netAlias | array[string] | C/R/- | false | true | - |  |
| networkContainerId | [container](#container) | C/R/- | false | true | \* - 255 |  |
| networkIds | array[[network](#network)] | C/R/- | false | true | - |  |
| networkLaunchConfig | string | C/R/- | false | true | - |  |
| networkMode | string | C/R/- | false | true | - |  |
| oomKillDisable | boolean | C/R/- | false | true | - |  |
| oomScoreAdj | int | C/R/- | false | true | - |  |
| pidMode | enum | C/R/- | false | true | - | enum options: host;  |
| pidsLimit | int | C/R/- | false | true | - |  |
| ports | array[string] | C/R/U | false | true | - |  |
| primaryIpAddress | string | -/R/- | false | false | - |  |
| primaryNetworkId | [network](#network) | -/R/- | false | false | - |  |
| privileged | boolean | C/R/- | false | false | - |  |
| publishAllPorts | boolean | C/R/- | false | false | - |  |
| readOnly | boolean | C/R/- | false | false | - |  |
| registryCredentialId | [registryCredential](#registryCredential) | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| requestedHostId | [host](#host) | C/R/- | false | true | - |  |
| requestedIpAddress | string | C/R/- | false | true | - |  |
| runInit | boolean | C/R/- | false | false | - |  |
| secrets | array[secretReference] | C/R/- | false | false | - |  |
| securityOpt | array[string] | C/R/- | false | true | - |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| serviceIds | array[[service](#service)] | -/R/- | false | true | - |  |
| shmSize | int | C/R/- | false | true | - |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startCount | int | -/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: creating/error/erroring/migrating/purged/purging/removed/removing/requested/restarting/running/starting/stopped/stopping/updating-running/updating-stopped;  |
| stdinOpen | boolean | C/R/- | false | false | - |  |
| stopSignal | string | C/R/- | false | true | - |  |
| stopTimeout | int | C/R/- | false | true | - |  |
| storageOpt | map[string] | C/R/- | false | true | - |  |
| sysctls | map[string] | C/R/- | false | true | - |  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| tmpfs | map[string] | C/R/- | false | true | - |  |
| token | string | -/R/- | false | true | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| tty | boolean | C/R/- | false | false | - |  |
| ulimits | array[ulimit] | C/R/- | false | true | - |  |
| user | string | C/R/- | false | true | 1 - \* |  |
| userPorts | array[string] | -/R/- | false | true | - |  |
| userdata | string | C/R/- | false | true | \* - 65535 |  |
| usernsMode | string | C/R/- | false | true | - |  |
| uts | string | C/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vcpu | int | C/R/- | false | true | - |  |
| version | string | -/R/- | false | false | \* - 255 |  |
| volumeDriver | string | C/R/- | false | true | - |  |
| workingDir | string | C/R/- | false | true | 1 - \* |  |

### secret
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/- | true | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: active/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| value | string | C/R/- | false | true | \* - 16777215 |  |

### secretReference
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| gid | string | C/R/- | false | false | - |  |
| mode | string | C/R/- | false | false | - |  |
| name | string | C/R/- | false | false | - |  |
| secretId | [secret](#secret) | C/R/- | false | false | - |  |
| uid | string | C/R/- | false | false | - |  |

### service
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| assignServiceIpAddress | boolean | C/R/- | false | false | - |  |
| createIndex | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| currentScale | int | -/R/- | false | false | - |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| fqdn | string | -/R/- | false | true | - |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| launchConfig | launchConfig | C/R/- | false | true | - |  |
| lbConfig | lbTargetConfig | C/R/U | false | true | - |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| metadata | map[json] | C/R/U | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| publicEndpoints | array[publicEndpoint] | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| retainIp | boolean | C/R/U | false | true | - |  |
| scale | int | C/R/U | false | false | - |  |
| scalePolicy | scalePolicy | C/R/U | false | true | - |  |
| secondaryLaunchConfigs | array[secondaryLaunchConfig] | C/R/- | false | false | - |  |
| selectorContainer | string | C/R/U | false | true | \* - 4096 |  |
| selectorLink | string | C/R/U | false | true | \* - 4096 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| upgrade | serviceUpgrade | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vip | string | C/R/- | false | true | \* - 255 |  |

### serviceBinding
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| labels | map[string] | C/R/- | false | false | - |  |
| ports | array[string] | C/R/- | false | false | - |  |

### serviceConsumeMap
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| consumedService | string | -/R/- | false | true | \* - 255 |  |
| consumedServiceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | -/R/- | false | true | \* - 255 |  |
| ports | array[string] | -/R/- | true | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested/updating-active;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### serviceEvent
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalTimestamp | int | -/R/- | false | true | \* - 255 |  |
| healthcheckUuid | string | -/R/- | false | true | \* - 255 |  |
| hostId | [host](#host) | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| reportedHealth | string | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### serviceExposeMap
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| ipAddress | string | -/R/- | false | true | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| managed | boolean | -/R/- | false | false | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### serviceLink
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| name | string | C/R/- | false | true | - | valid chars: a-zA-Z0-9-._;  |
| service | string | C/R/- | false | false | - |  |
| serviceId | [service](#service) | C/R/- | false | false | - |  |
| uuid | string | -/R/- | false | false | - |  |

### serviceLog
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| endTime | date | -/R/- | false | true | \* - 255 |  |
| eventType | string | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| level | string | -/R/- | false | true | \* - 255 |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| subLog | boolean | -/R/- | false | false | \* - 255 |  |
| transactionId | string | -/R/- | false | true | \* - 255 |  |

### serviceProxy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| port | int | C/R/- | false | true | - |  |
| scheme | enum | C/R/- | false | false | - | enum options: http/https;  |
| service | string | C/R/- | true | false | 1 - \* |  |
| token | string | -/R/- | false | false | - |  |
| url | string | -/R/- | false | false | - |  |

### serviceRestart
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| rollingRestartStrategy | rollingRestartStrategy | C/R/- | true | false | - |  |

### serviceUpgrade
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| inServiceStrategy | inServiceUpgradeStrategy | C/R/- | false | true | - |  |
| toServiceStrategy | toServiceUpgradeStrategy | C/R/- | false | true | - |  |

### serviceUpgradeStrategy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| batchSize | int | C/R/- | false | true | 1 - \* |  |
| intervalMillis | int | C/R/- | false | true | 100 - \* |  |

### servicesPortRange
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| endPort | int | C/R/U | true | true | 1 - 65535 |  |
| startPort | int | C/R/U | true | true | 1 - 65535 |  |

### setProjectMembersInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| members | array[projectMember] | C/R/- | true | false | 1 - \* |  |

### setServiceLinksInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| serviceLinks | array[serviceLink] | C/R/- | false | true | - |  |

### setting
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| activeValue | string | -/R/- | false | false | - |  |
| id | int | -/R/- | false | true | - |  |
| inDb | boolean | -/R/- | false | false | - |  |
| name | string | C/R/- | false | true | \* - 255 |  |
| source | string | -/R/- | false | false | - |  |
| value | string | C/R/U | false | true | \* - 16777215 |  |

### snapshot
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: created/creating/removed/removing/requested;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volumeId | [volume](#volume) | -/R/- | true | false | \* - 255 |  |

### snapshotBackupInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| backupTargetId | [backupTarget](#backupTarget) | C/R/- | true | false | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| name | string | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |

### stack
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| answers | map[json] | C/R/- | false | false | - |  |
| binding | binding | C/R/U | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| dockerCompose | string | C/R/- | false | false | - |  |
| environment | map[string] | C/R/- | false | false | - |  |
| externalId | string | C/R/U | false | true | \* - 128 |  |
| group | string | C/R/U | false | true | \* - 255 |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| outputs | map[string] | C/R/U | false | true | - |  |
| previousEnvironment | map[string] | C/R/U | false | true | - |  |
| previousExternalId | string | C/R/U | false | true | - |  |
| rancherCompose | string | C/R/- | false | false | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| serviceIds | array[[service](#service)] | -/R/- | false | true | - |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/error/erroring/finishing-upgrade/removed/removing/requested/rolling-back/updating-active/upgraded/upgrading;  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| templates | map[string] | C/R/- | false | true | - |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### stackUpgrade
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| answers | map[json] | C/R/- | false | false | - |  |
| dockerCompose | string | C/R/- | false | false | - |  |
| environment | map[string] | C/R/- | false | false | - |  |
| externalId | string | C/R/- | false | false | - |  |
| rancherCompose | string | C/R/- | false | false | - |  |
| templates | map[string] | C/R/- | false | false | - |  |

### stateTransition
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |

### statsAccess
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| token | string | -/R/- | false | false | - |  |
| url | string | -/R/- | false | false | - |  |

### storageDriver
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| blockDevicePath | string | C/R/- | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| scope | enum | C/R/- | false | false | - | enum options: local/environment/custom;  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/creating/deactivating/inactive/removed/removing/requested/updating;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volumeAccessMode | enum | C/R/- | false | false | - | enum options: singleHostRW/singleInstanceRW/multiHostRW;  |
| volumeCapabilities | array[string] | C/R/- | false | true | - |  |

### storageDriverService
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| assignServiceIpAddress | boolean | C/R/- | false | false | - |  |
| createIndex | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| currentScale | int | -/R/- | false | false | - |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| externalId | string | C/R/- | false | true | \* - 255 |  |
| fqdn | string | -/R/- | false | true | - |  |
| healthState | string | -/R/- | false | true | \* - 128 |  |
| id | int | -/R/- | false | true | - |  |
| instanceIds | array[[instance](#instance)] | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| launchConfig | launchConfig | C/R/- | false | true | - |  |
| lbConfig | lbTargetConfig | C/R/U | false | true | - |  |
| linkedServices | map[[service](#service)] | -/R/- | false | true | - |  |
| metadata | map[json] | C/R/U | false | true | - |  |
| name | string | C/R/U | true | true | 1 - 63 | valid chars: a-zA-Z0-9-;  |
| publicEndpoints | array[publicEndpoint] | C/R/U | false | true | - |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| retainIp | boolean | C/R/U | false | true | - |  |
| scale | int | C/R/U | false | false | - |  |
| scalePolicy | scalePolicy | C/R/U | false | true | - |  |
| secondaryLaunchConfigs | array[secondaryLaunchConfig] | C/R/- | false | false | - |  |
| selectorContainer | string | C/R/U | false | true | \* - 4096 |  |
| selectorLink | string | C/R/U | false | true | \* - 4096 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | true | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/canceled-upgrade/canceling-upgrade/deactivating/finishing-upgrade/inactive/registering/removed/removing/requested/restarting/rolling-back/updating-active/updating-inactive/upgraded/upgrading;  |
| storageDriver | storageDriver | C/R/U | true | false | - |  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| upgrade | serviceUpgrade | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vip | string | C/R/- | false | true | \* - 255 |  |

### storagePool
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| blockDevicePath | string | -/R/- | false | true | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| driverName | string | -/R/- | false | true | \* - 255 |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| hostIds | array[[host](#host)] | -/R/- | false | true | - |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| storageDriverId | [storageDriver](#storageDriver) | -/R/- | false | true | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volumeAccessMode | string | -/R/- | false | true | \* - 255 |  |
| volumeCapabilities | array[string] | -/R/- | false | true | - |  |
| volumeIds | array[[volume](#volume)] | -/R/- | false | true | - |  |

### subnet
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| cidrSize | int | C/R/- | true | true | 8 - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| endAddress | string | C/R/- | false | true | \* - 255 |  |
| gateway | string | C/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| networkAddress | string | C/R/- | true | true | \* - 255 |  |
| networkId | [network](#network) | -/R/- | true | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| startAddress | string | C/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

### targetPortRule
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| backendName | string | C/R/- | false | true | - |  |
| hostname | string | C/R/- | false | true | - |  |
| path | string | C/R/- | false | true | - |  |
| targetPort | int | C/R/- | true | true | 1 - 65535 |  |

### task
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| id | int | -/R/- | false | true | - |  |
| name | string | C/R/U | false | true | \* - 128 |  |

### taskInstance
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| endTime | date | -/R/- | false | true | \* - 255 |  |
| exception | string | -/R/- | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| name | string | C/R/U | false | true | \* - 128 |  |
| serverId | string | -/R/- | false | true | \* - 128 |  |
| startTime | date | -/R/- | false | false | \* - 255 |  |
| taskId | [task](#task) | -/R/- | false | false | \* - 255 |  |

### toServiceUpgradeStrategy
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| batchSize | int | C/R/- | false | true | 1 - \* |  |
| finalScale | int | C/R/- | false | true | 1 - \* |  |
| intervalMillis | int | C/R/- | false | true | 100 - \* |  |
| toServiceId | [service](#service) | C/R/- | false | false | - |  |
| updateLinks | boolean | C/R/- | false | false | - |  |

### typeDocumentation
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| description | string | C/R/U | false | false | - |  |
| id | string | -/R/- | false | false | - |  |
| resourceFields | map[fieldDocumentation] | -/R/- | false | false | - |  |

### ulimit
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| hard | int | C/R/- | false | true | - |  |
| name | string | C/R/- | false | false | - |  |
| soft | int | C/R/- | false | true | - |  |

### userPreference
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| all | boolean | -/R/- | false | false | - |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| value | string | C/R/U | false | true | \* - 16777215 |  |

### virtualMachine
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| agentId | [agent](#agent) | -/R/- | false | true | \* - 255 |  |
| allocationState | string | -/R/- | false | true | \* - 255 |  |
| blkioDeviceOptions | map[blkioDeviceOption] | C/R/- | false | true | - |  |
| blkioWeight | int | C/R/- | false | true | - |  |
| cgroupParent | string | C/R/- | false | true | - |  |
| command | array[string] | C/R/- | false | true | - |  |
| count | int | C/R/- | false | true | - |  |
| cpuCount | int | C/R/- | false | true | - |  |
| cpuPercent | int | C/R/- | false | true | - |  |
| cpuPeriod | int | C/R/- | false | true | - |  |
| cpuQuota | int | C/R/- | false | true | - |  |
| cpuRealtimePeriod | int | C/R/- | false | true | - |  |
| cpuRealtimeRuntime | int | C/R/- | false | true | - |  |
| cpuSet | string | C/R/- | false | true | - |  |
| cpuSetMems | string | C/R/- | false | true | - |  |
| cpuShares | int | C/R/- | false | true | - |  |
| createIndex | int | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| deploymentUnitUuid | string | -/R/- | false | true | \* - 128 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| diskQuota | int | C/R/- | false | true | - |  |
| disks | array[virtualMachineDisk] | C/R/- | false | true | - |  |
| dns | array[string] | C/R/- | false | true | - |  |
| dnsOpt | array[string] | C/R/- | false | true | - |  |
| dnsSearch | array[string] | C/R/- | false | true | - |  |
| domainName | string | C/R/- | false | true | - |  |
| expose | array[string] | C/R/- | false | true | - |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| extraHosts | array[string] | C/R/- | false | true | - |  |
| firstRunning | date | -/R/- | false | true | \* - 255 |  |
| groupAdd | array[string] | C/R/- | false | true | - |  |
| healthCheck | instanceHealthCheck | C/R/- | false | true | - |  |
| healthCmd | array[string] | C/R/- | false | true | - |  |
| healthInterval | int | C/R/- | false | true | - |  |
| healthRetries | int | C/R/- | false | true | - |  |
| healthState | enum | -/R/- | false | true | \* - 128 | enum options: healthy/unhealthy/updating-healthy/updating-unhealthy/initializing;  |
| healthTimeout | int | C/R/- | false | true | - |  |
| hostId | [host](#host) | -/R/- | false | true | - |  |
| hostname | string | C/R/- | false | true | 1 - 255 |  |
| id | int | -/R/- | false | true | - |  |
| imageUuid | string | C/R/- | false | true | - |  |
| instanceLinks | map[[instance](#instance)] | C/R/- | false | true | - |  |
| instanceTriggeredStop | enum | C/R/- | false | true | \* - 128 | enum options: stop/remove;  |
| ioMaximumBandwidth | int | C/R/- | false | true | - |  |
| ioMaximumIOps | int | C/R/- | false | true | - |  |
| ip | string | C/R/- | false | true | - |  |
| ip6 | string | C/R/- | false | true | - |  |
| ipcMode | string | C/R/- | false | true | - |  |
| isolation | string | C/R/- | false | true | - |  |
| kernelMemory | int | C/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| labels | map[string] | C/R/- | false | false | - |  |
| logConfig | logConfig | C/R/- | false | true | - |  |
| memory | int | C/R/- | false | true | - |  |
| memoryMb | int | C/R/- | false | true | \* - 255 |  |
| memoryReservation | int | C/R/- | false | true | \* - 255 |  |
| memorySwap | int | C/R/- | false | true | - |  |
| memorySwappiness | int | C/R/- | false | true | - |  |
| milliCpuReservation | int | C/R/- | false | true | \* - 255 |  |
| mounts | array[mountEntry] | -/R/- | false | false | - |  |
| name | string | C/R/U | false | true | \* - 255 |  |
| nativeContainer | boolean | -/R/- | false | false | \* - 255 |  |
| netAlias | array[string] | C/R/- | false | true | - |  |
| networkIds | array[[network](#network)] | C/R/- | false | true | - |  |
| networkMode | string | C/R/- | false | true | - |  |
| oomKillDisable | boolean | C/R/- | false | true | - |  |
| oomScoreAdj | int | C/R/- | false | true | - |  |
| pidsLimit | int | C/R/- | false | true | - |  |
| ports | array[string] | C/R/- | false | true | - |  |
| primaryIpAddress | string | -/R/- | false | false | - |  |
| primaryNetworkId | [network](#network) | -/R/- | false | false | - |  |
| registryCredentialId | [registryCredential](#registryCredential) | C/R/- | false | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| requestedHostId | [host](#host) | C/R/- | false | true | - |  |
| restartPolicy | restartPolicy | C/R/- | false | true | - |  |
| runInit | boolean | C/R/- | false | false | - |  |
| securityOpt | array[string] | C/R/- | false | true | - |  |
| serviceId | [service](#service) | -/R/- | false | true | \* - 255 |  |
| serviceIds | array[[service](#service)] | -/R/- | false | true | - |  |
| shmSize | int | C/R/- | false | true | - |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| startCount | int | -/R/- | false | true | \* - 255 |  |
| startOnCreate | boolean | C/R/- | false | false | - |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: creating/error/erroring/migrating/purged/purging/removed/removing/requested/restarting/running/starting/stopped/stopping/updating-running/updating-stopped;  |
| stopSignal | string | C/R/- | false | true | - |  |
| stopTimeout | int | C/R/- | false | true | - |  |
| storageOpt | map[string] | C/R/- | false | true | - |  |
| sysctls | map[string] | C/R/- | false | true | - |  |
| system | boolean | -/R/- | false | false | \* - 255 |  |
| tmpfs | map[string] | C/R/- | false | true | - |  |
| token | string | -/R/- | false | true | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| ulimits | array[ulimit] | C/R/- | false | true | - |  |
| userPorts | array[string] | -/R/- | false | true | - |  |
| userdata | string | C/R/- | false | true | \* - 65535 |  |
| usernsMode | string | C/R/- | false | true | - |  |
| uts | string | C/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| vcpu | int | C/R/- | false | true | - |  |
| version | string | -/R/- | false | false | \* - 255 |  |
| volumeDriver | string | C/R/- | false | true | - |  |

### virtualMachineDisk
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| driver | string | C/R/- | false | false | - |  |
| name | string | C/R/- | false | true | 2 - \* | valid chars: a-z0-9_.-;  |
| opts | map[string] | C/R/- | false | false | - |  |
| readIops | int | C/R/- | false | true | - |  |
| root | boolean | C/R/- | false | false | - |  |
| size | string | C/R/- | false | false | - |  |
| writeIops | int | C/R/- | false | true | - |  |

### volume
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accessMode | string | -/R/- | false | true | \* - 255 |  |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 16777215 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| driver | string | C/R/- | false | false | - |  |
| driverOpts | map[string] | C/R/U | false | true | - |  |
| externalId | string | -/R/- | false | true | \* - 128 |  |
| hostId | [host](#host) | C/R/U | false | true | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| imageId | [image](#image) | -/R/- | false | true | \* - 255 |  |
| instanceId | [instance](#instance) | -/R/- | false | true | \* - 255 |  |
| isHostPath | boolean | -/R/- | false | false | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| mounts | array[mountEntry] | -/R/- | false | false | - |  |
| name | string | C/R/- | true | true | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| sizeMb | int | C/R/- | false | true | \* - 255 |  |
| stackId | [stack](#stack) | C/R/- | false | true | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/creating/deactivating/detached/inactive/purged/purging/registering/removed/removing/requested/restoring-active/restoring-inactive/reverting-active/reverting-inactive/updating-active/updating-inactive;  |
| storageDriverId | [storageDriver](#storageDriver) | C/R/- | false | false | \* - 255 |  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uri | string | -/R/- | false | true | \* - 255 |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |
| volumeTemplateId | [volumeTemplate](#volumeTemplate) | C/R/- | false | true | \* - 255 |  |

### volumeActivateInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| hostId | reference | C/R/- | false | true | - |  |

### volumeSnapshotInput
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| name | string | C/R/- | false | false | \* - 255 |  |

### volumeTemplate
  
| attribute name | type | C/R/U | required | nullable | length/range | comments |
| --- | --- | --- | --- | --- | --- | --- |
| accountId | [account](#account) | -/R/- | false | true | \* - 255 |  |
| created | date | -/R/- | false | true | \* - 255 |  |
| data | map[json] | -/R/- | false | true | \* - 65535 |  |
| description | string | C/R/U | false | true | \* - 1024 |  |
| driver | string | C/R/- | false | true | \* - 255 |  |
| driverOpts | map[string] | C/R/- | false | true | - |  |
| external | boolean | C/R/- | false | false | \* - 255 |  |
| id | int | -/R/- | false | true | - |  |
| kind | string | -/R/- | false | true | \* - 255 |  |
| name | string | C/R/U | true | true | \* - 255 |  |
| perContainer | boolean | C/R/- | false | false | \* - 255 |  |
| removeTime | date | -/R/- | false | true | \* - 255 |  |
| removed | date | -/R/- | false | true | \* - 255 |  |
| stackId | [stack](#stack) | C/R/- | true | false | \* - 255 |  |
| state | enum | -/R/- | false | false | \* - 128 | enum options: activating/active/deactivating/inactive/purged/purging/registering/removed/removing/requested/updating-active/updating-inactive;  |
| transitioning | enum | -/R/- | false | false | - | enum options: yes/no/error;  |
| transitioningMessage | string | -/R/- | false | true | - |  |
| transitioningProgress | int | -/R/- | false | true | - |  |
| uuid | string | -/R/- | false | true | \* - 128 |  |

## API List

This is the API List of this endppint.

### Schema account API List

#### GET /v2-beta/accounts

- input: null
- output: collection of [account](#account)

#### POST /v2-beta/accounts

- input: [account](#account)
- output: [account](#account)

#### GET /v2-beta/accounts/*resource-id*

- input: null
- output: [account](#account)

#### PUT /v2-beta/accounts/*resource-id*

- input: [account](#account)
- output: [account](#account)

#### DELETE /v2-beta/accounts/*resource-id*

- input: null
- output: [account](#account)

#### Action upgrade of schema object account, POST /v2-beta/accounts/*resource-id*?action=upgrade

- input: null
- output: [account](#account)

#### Action activate of schema object account, POST /v2-beta/accounts/*resource-id*?action=activate

- input: null
- output: [account](#account)

#### Action create of schema object account, POST /v2-beta/accounts/*resource-id*?action=create

- input: null
- output: [account](#account)

#### Action update of schema object account, POST /v2-beta/accounts/*resource-id*?action=update

- input: null
- output: [account](#account)

#### Action purge of schema object account, POST /v2-beta/accounts/*resource-id*?action=purge

- input: null
- output: [account](#account)

#### Action remove of schema object account, POST /v2-beta/accounts/*resource-id*?action=remove

- input: null
- output: [account](#account)

#### Action deactivate of schema object account, POST /v2-beta/accounts/*resource-id*?action=deactivate

- input: null
- output: [account](#account)

### Schema accountLink API List

#### GET /v2-beta/accountlinks

- input: null
- output: collection of [accountLink](#accountlink)

#### GET /v2-beta/accountlinks/*resource-id*

- input: null
- output: [accountLink](#accountlink)

#### Action activate of schema object accountLink, POST /v2-beta/accountlinks/*resource-id*?action=activate

- input: null
- output: [accountLink](#accountlink)

#### Action create of schema object accountLink, POST /v2-beta/accountlinks/*resource-id*?action=create

- input: null
- output: [accountLink](#accountlink)

#### Action update of schema object accountLink, POST /v2-beta/accountlinks/*resource-id*?action=update

- input: null
- output: [accountLink](#accountlink)

#### Action purge of schema object accountLink, POST /v2-beta/accountlinks/*resource-id*?action=purge

- input: null
- output: [accountLink](#accountlink)

#### Action remove of schema object accountLink, POST /v2-beta/accountlinks/*resource-id*?action=remove

- input: null
- output: [accountLink](#accountlink)

#### Action deactivate of schema object accountLink, POST /v2-beta/accountlinks/*resource-id*?action=deactivate

- input: null
- output: [accountLink](#accountlink)

### Schema agent API List

#### GET /v2-beta/agents

- input: null
- output: collection of [agent](#agent)

#### POST /v2-beta/agents

- input: [agent](#agent)
- output: [agent](#agent)

#### GET /v2-beta/agents/*resource-id*

- input: null
- output: [agent](#agent)

#### PUT /v2-beta/agents/*resource-id*

- input: [agent](#agent)
- output: [agent](#agent)

#### DELETE /v2-beta/agents/*resource-id*

- input: null
- output: [agent](#agent)

#### Action purge of schema object agent, POST /v2-beta/agents/*resource-id*?action=purge

- input: null
- output: [agent](#agent)

#### Action finishreconnect of schema object agent, POST /v2-beta/agents/*resource-id*?action=finishreconnect

- input: null
- output: [agent](#agent)

#### Action create of schema object agent, POST /v2-beta/agents/*resource-id*?action=create

- input: null
- output: [agent](#agent)

#### Action update of schema object agent, POST /v2-beta/agents/*resource-id*?action=update

- input: null
- output: [agent](#agent)

#### Action remove of schema object agent, POST /v2-beta/agents/*resource-id*?action=remove

- input: null
- output: [agent](#agent)

#### Action deactivate of schema object agent, POST /v2-beta/agents/*resource-id*?action=deactivate

- input: null
- output: [agent](#agent)

#### Action disconnect of schema object agent, POST /v2-beta/agents/*resource-id*?action=disconnect

- input: null
- output: [agent](#agent)

#### Action reconnect of schema object agent, POST /v2-beta/agents/*resource-id*?action=reconnect

- input: null
- output: [agent](#agent)

#### Action activate of schema object agent, POST /v2-beta/agents/*resource-id*?action=activate

- input: null
- output: [agent](#agent)

### Schema apiKey API List

#### GET /v2-beta/apikeys

- input: null
- output: collection of [apiKey](#apikey)

#### POST /v2-beta/apikeys

- input: [apiKey](#apikey)
- output: [apiKey](#apikey)

#### GET /v2-beta/apikeys/*resource-id*

- input: null
- output: [apiKey](#apikey)

#### PUT /v2-beta/apikeys/*resource-id*

- input: [apiKey](#apikey)
- output: [apiKey](#apikey)

#### DELETE /v2-beta/apikeys/*resource-id*

- input: null
- output: [apiKey](#apikey)

#### Action activate of schema object apiKey, POST /v2-beta/apikeys/*resource-id*?action=activate

- input: null
- output: [credential](#credential)

#### Action create of schema object apiKey, POST /v2-beta/apikeys/*resource-id*?action=create

- input: null
- output: [credential](#credential)

#### Action update of schema object apiKey, POST /v2-beta/apikeys/*resource-id*?action=update

- input: null
- output: [credential](#credential)

#### Action purge of schema object apiKey, POST /v2-beta/apikeys/*resource-id*?action=purge

- input: null
- output: [credential](#credential)

#### Action remove of schema object apiKey, POST /v2-beta/apikeys/*resource-id*?action=remove

- input: null
- output: [credential](#credential)

#### Action deactivate of schema object apiKey, POST /v2-beta/apikeys/*resource-id*?action=deactivate

- input: null
- output: [credential](#credential)

### Schema auditLog API List

#### GET /v2-beta/auditlogs

- input: null
- output: collection of [auditLog](#auditlog)

#### GET /v2-beta/auditlogs/*resource-id*

- input: null
- output: [auditLog](#auditlog)

### Schema azureadconfig API List

#### GET /v2-beta/azureadconfigs

- input: null
- output: collection of [azureadconfig](#azureadconfig)

#### POST /v2-beta/azureadconfigs

- input: [azureadconfig](#azureadconfig)
- output: [azureadconfig](#azureadconfig)

#### GET /v2-beta/azureadconfigs/*resource-id*

- input: null
- output: [azureadconfig](#azureadconfig)

#### PUT /v2-beta/azureadconfigs/*resource-id*

- input: [azureadconfig](#azureadconfig)
- output: [azureadconfig](#azureadconfig)

### Schema backup API List

#### GET /v2-beta/backups

- input: null
- output: collection of [backup](#backup)

#### GET /v2-beta/backups/*resource-id*

- input: null
- output: [backup](#backup)

#### Action create of schema object backup, POST /v2-beta/backups/*resource-id*?action=create

- input: null
- output: [backup](#backup)

#### Action remove of schema object backup, POST /v2-beta/backups/*resource-id*?action=remove

- input: null
- output: [backup](#backup)

### Schema backupTarget API List

#### GET /v2-beta/backuptargets

- input: null
- output: collection of [backupTarget](#backuptarget)

#### GET /v2-beta/backuptargets/*resource-id*

- input: null
- output: [backupTarget](#backuptarget)

#### Action create of schema object backupTarget, POST /v2-beta/backuptargets/*resource-id*?action=create

- input: null
- output: [backupTarget](#backuptarget)

#### Action remove of schema object backupTarget, POST /v2-beta/backuptargets/*resource-id*?action=remove

- input: null
- output: [backupTarget](#backuptarget)

### Schema certificate API List

#### GET /v2-beta/certificates

- input: null
- output: collection of [certificate](#certificate)

#### GET /v2-beta/certificates/*resource-id*

- input: null
- output: [certificate](#certificate)

#### Action remove of schema object certificate, POST /v2-beta/certificates/*resource-id*?action=remove

- input: null
- output: [certificate](#certificate)

#### Action create of schema object certificate, POST /v2-beta/certificates/*resource-id*?action=create

- input: null
- output: [certificate](#certificate)

#### Action update of schema object certificate, POST /v2-beta/certificates/*resource-id*?action=update

- input: null
- output: [certificate](#certificate)

### Schema clusterMembership API List

#### GET /v2-beta/clustermemberships

- input: null
- output: collection of [clusterMembership](#clustermembership)

#### GET /v2-beta/clustermemberships/*resource-id*

- input: null
- output: [clusterMembership](#clustermembership)

### Schema composeProject API List

#### GET /v2-beta/composeprojects

- input: null
- output: collection of [composeProject](#composeproject)

#### GET /v2-beta/composeprojects/*resource-id*

- input: null
- output: [composeProject](#composeproject)

#### Action create of schema object composeProject, POST /v2-beta/composeprojects/*resource-id*?action=create

- input: null
- output: [stack](#stack)

#### Action error of schema object composeProject, POST /v2-beta/composeprojects/*resource-id*?action=error

- input: null
- output: [stack](#stack)

#### Action cancelupgrade of schema object composeProject, POST /v2-beta/composeprojects/*resource-id*?action=cancelupgrade

- input: null
- output: [stack](#stack)

#### Action remove of schema object composeProject, POST /v2-beta/composeprojects/*resource-id*?action=remove

- input: null
- output: [stack](#stack)

#### Action finishupgrade of schema object composeProject, POST /v2-beta/composeprojects/*resource-id*?action=finishupgrade

- input: null
- output: [stack](#stack)

#### Action rollback of schema object composeProject, POST /v2-beta/composeprojects/*resource-id*?action=rollback

- input: null
- output: [stack](#stack)

### Schema composeService API List

#### GET /v2-beta/composeservices

- input: null
- output: collection of [composeService](#composeservice)

#### GET /v2-beta/composeservices/*resource-id*

- input: null
- output: [composeService](#composeservice)

#### Action continueupgrade of schema object composeService, POST /v2-beta/composeservices/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

#### Action activate of schema object composeService, POST /v2-beta/composeservices/*resource-id*?action=activate

- input: null
- output: [service](#service)

#### Action create of schema object composeService, POST /v2-beta/composeservices/*resource-id*?action=create

- input: null
- output: [service](#service)

#### Action cancelupgrade of schema object composeService, POST /v2-beta/composeservices/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action rollback of schema object composeService, POST /v2-beta/composeservices/*resource-id*?action=rollback

- input: null
- output: [service](#service)

#### Action remove of schema object composeService, POST /v2-beta/composeservices/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action finishupgrade of schema object composeService, POST /v2-beta/composeservices/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

### Schema configItem API List

#### GET /v2-beta/configitems

- input: null
- output: collection of [configItem](#configitem)

#### GET /v2-beta/configitems/*resource-id*

- input: null
- output: [configItem](#configitem)

### Schema configItemStatus API List

#### GET /v2-beta/configitemstatuses

- input: null
- output: collection of [configItemStatus](#configitemstatus)

#### GET /v2-beta/configitemstatuses/*resource-id*

- input: null
- output: [configItemStatus](#configitemstatus)

#### PUT /v2-beta/configitemstatuses/*resource-id*

- input: [configItemStatus](#configitemstatus)
- output: [configItemStatus](#configitemstatus)

### Schema container API List

#### GET /v2-beta/containers

- input: null
- output: collection of [container](#container)

#### GET /v2-beta/containers/*resource-id*

- input: null
- output: [container](#container)

#### Action update of schema object container, POST /v2-beta/containers/*resource-id*?action=update

- input: null
- output: [instance](#instance)

#### Action error of schema object container, POST /v2-beta/containers/*resource-id*?action=error

- input: null
- output: [instance](#instance)

#### Action allocate of schema object container, POST /v2-beta/containers/*resource-id*?action=allocate

- input: null
- output: [instance](#instance)

#### Action updatehealthy of schema object container, POST /v2-beta/containers/*resource-id*?action=updatehealthy

- input: null
- output: [instance](#instance)

#### Action logs of schema object container, POST /v2-beta/containers/*resource-id*?action=logs

- input: [containerLogs](#containerlogs)
- output: [hostAccess](#hostaccess)

#### Action execute of schema object container, POST /v2-beta/containers/*resource-id*?action=execute

- input: [containerExec](#containerexec)
- output: [hostAccess](#hostaccess)

#### Action proxy of schema object container, POST /v2-beta/containers/*resource-id*?action=proxy

- input: [containerProxy](#containerproxy)
- output: [hostAccess](#hostaccess)

#### Action console of schema object container, POST /v2-beta/containers/*resource-id*?action=console

- input: [instanceConsoleInput](#instanceconsoleinput)
- output: [instanceConsole](#instanceconsole)

#### Action deallocate of schema object container, POST /v2-beta/containers/*resource-id*?action=deallocate

- input: null
- output: [instance](#instance)

#### Action start of schema object container, POST /v2-beta/containers/*resource-id*?action=start

- input: null
- output: [instance](#instance)

#### Action purge of schema object container, POST /v2-beta/containers/*resource-id*?action=purge

- input: null
- output: [instance](#instance)

#### Action remove of schema object container, POST /v2-beta/containers/*resource-id*?action=remove

- input: null
- output: [instance](#instance)

#### Action create of schema object container, POST /v2-beta/containers/*resource-id*?action=create

- input: null
- output: [instance](#instance)

#### Action migrate of schema object container, POST /v2-beta/containers/*resource-id*?action=migrate

- input: null
- output: [instance](#instance)

#### Action updateunhealthy of schema object container, POST /v2-beta/containers/*resource-id*?action=updateunhealthy

- input: null
- output: [instance](#instance)

#### Action restart of schema object container, POST /v2-beta/containers/*resource-id*?action=restart

- input: null
- output: [instance](#instance)

#### Action updatereinitializing of schema object container, POST /v2-beta/containers/*resource-id*?action=updatereinitializing

- input: null
- output: [instance](#instance)

#### Action stop of schema object container, POST /v2-beta/containers/*resource-id*?action=stop

- input: [instanceStop](#instancestop)
- output: [instance](#instance)

### Schema containerEvent API List

#### GET /v2-beta/containerevents

- input: null
- output: collection of [containerEvent](#containerevent)

#### GET /v2-beta/containerevents/*resource-id*

- input: null
- output: [containerEvent](#containerevent)

#### Action create of schema object containerEvent, POST /v2-beta/containerevents/*resource-id*?action=create

- input: null
- output: [containerEvent](#containerevent)

#### Action remove of schema object containerEvent, POST /v2-beta/containerevents/*resource-id*?action=remove

- input: null
- output: [containerEvent](#containerevent)

### Schema credential API List

#### GET /v2-beta/credentials

- input: null
- output: collection of [credential](#credential)

#### POST /v2-beta/credentials

- input: [credential](#credential)
- output: [credential](#credential)

#### GET /v2-beta/credentials/*resource-id*

- input: null
- output: [credential](#credential)

#### Action activate of schema object credential, POST /v2-beta/credentials/*resource-id*?action=activate

- input: null
- output: [credential](#credential)

#### Action create of schema object credential, POST /v2-beta/credentials/*resource-id*?action=create

- input: null
- output: [credential](#credential)

#### Action update of schema object credential, POST /v2-beta/credentials/*resource-id*?action=update

- input: null
- output: [credential](#credential)

#### Action purge of schema object credential, POST /v2-beta/credentials/*resource-id*?action=purge

- input: null
- output: [credential](#credential)

#### Action remove of schema object credential, POST /v2-beta/credentials/*resource-id*?action=remove

- input: null
- output: [credential](#credential)

#### Action deactivate of schema object credential, POST /v2-beta/credentials/*resource-id*?action=deactivate

- input: null
- output: [credential](#credential)

### Schema databasechangelog API List

#### GET /v2-beta/databasechangelogs

- input: null
- output: collection of [databasechangelog](#databasechangelog)

#### GET /v2-beta/databasechangelogs/*resource-id*

- input: null
- output: [databasechangelog](#databasechangelog)

#### DELETE /v2-beta/databasechangelogs/*resource-id*

- input: null
- output: [databasechangelog](#databasechangelog)

### Schema databasechangeloglock API List

#### GET /v2-beta/databasechangeloglocks

- input: null
- output: collection of [databasechangeloglock](#databasechangeloglock)

#### GET /v2-beta/databasechangeloglocks/*resource-id*

- input: null
- output: [databasechangeloglock](#databasechangeloglock)

#### DELETE /v2-beta/databasechangeloglocks/*resource-id*

- input: null
- output: [databasechangeloglock](#databasechangeloglock)

### Schema dnsService API List

#### GET /v2-beta/dnsservices

- input: null
- output: collection of [dnsService](#dnsservice)

#### GET /v2-beta/dnsservices/*resource-id*

- input: null
- output: [dnsService](#dnsservice)

#### Action finishupgrade of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

#### Action activate of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=activate

- input: null
- output: [service](#service)

#### Action create of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=create

- input: null
- output: [service](#service)

#### Action deactivate of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=deactivate

- input: null
- output: [service](#service)

#### Action restart of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=restart

- input: [serviceRestart](#servicerestart)
- output: [service](#service)

#### Action update of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=update

- input: null
- output: [service](#service)

#### Action remove of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action setservicelinks of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=setservicelinks

- input: [setServiceLinksInput](#setservicelinksinput)
- output: [service](#service)

#### Action rollback of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=rollback

- input: null
- output: [service](#service)

#### Action continueupgrade of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

#### Action removeservicelink of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=removeservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action addservicelink of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=addservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action cancelupgrade of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action upgrade of schema object dnsService, POST /v2-beta/dnsservices/*resource-id*?action=upgrade

- input: [serviceUpgrade](#serviceupgrade)
- output: [service](#service)

### Schema extensionPoint API List

#### GET /v2-beta/extensionpoints

- input: null
- output: collection of [extensionPoint](#extensionpoint)

#### GET /v2-beta/extensionpoints/*resource-id*

- input: null
- output: [extensionPoint](#extensionpoint)

### Schema externalCredential API List

#### GET /v2-beta/externalcredentials

- input: null
- output: collection of [externalCredential](#externalcredential)

#### GET /v2-beta/externalcredentials/*resource-id*

- input: null
- output: [externalCredential](#externalcredential)

### Schema externalDnsEvent API List

#### GET /v2-beta/externaldnsevents

- input: null
- output: collection of [externalDnsEvent](#externaldnsevent)

#### GET /v2-beta/externaldnsevents/*resource-id*

- input: null
- output: [externalDnsEvent](#externaldnsevent)

#### Action create of schema object externalDnsEvent, POST /v2-beta/externaldnsevents/*resource-id*?action=create

- input: null
- output: [externalEvent](#externalevent)

#### Action remove of schema object externalDnsEvent, POST /v2-beta/externaldnsevents/*resource-id*?action=remove

- input: null
- output: [externalEvent](#externalevent)

### Schema externalEvent API List

#### GET /v2-beta/externalevents

- input: null
- output: collection of [externalEvent](#externalevent)

#### GET /v2-beta/externalevents/*resource-id*

- input: null
- output: [externalEvent](#externalevent)

#### Action create of schema object externalEvent, POST /v2-beta/externalevents/*resource-id*?action=create

- input: null
- output: [externalEvent](#externalevent)

#### Action remove of schema object externalEvent, POST /v2-beta/externalevents/*resource-id*?action=remove

- input: null
- output: [externalEvent](#externalevent)

### Schema externalHandler API List

#### GET /v2-beta/externalhandlers

- input: null
- output: collection of [externalHandler](#externalhandler)

#### POST /v2-beta/externalhandlers

- input: [externalHandler](#externalhandler)
- output: [externalHandler](#externalhandler)

#### GET /v2-beta/externalhandlers/*resource-id*

- input: null
- output: [externalHandler](#externalhandler)

#### PUT /v2-beta/externalhandlers/*resource-id*

- input: [externalHandler](#externalhandler)
- output: [externalHandler](#externalhandler)

#### Action deactivate of schema object externalHandler, POST /v2-beta/externalhandlers/*resource-id*?action=deactivate

- input: null
- output: [externalHandler](#externalhandler)

#### Action activate of schema object externalHandler, POST /v2-beta/externalhandlers/*resource-id*?action=activate

- input: null
- output: [externalHandler](#externalhandler)

#### Action create of schema object externalHandler, POST /v2-beta/externalhandlers/*resource-id*?action=create

- input: null
- output: [externalHandler](#externalhandler)

#### Action update of schema object externalHandler, POST /v2-beta/externalhandlers/*resource-id*?action=update

- input: null
- output: [externalHandler](#externalhandler)

#### Action purge of schema object externalHandler, POST /v2-beta/externalhandlers/*resource-id*?action=purge

- input: null
- output: [externalHandler](#externalhandler)

#### Action remove of schema object externalHandler, POST /v2-beta/externalhandlers/*resource-id*?action=remove

- input: null
- output: [externalHandler](#externalhandler)

### Schema externalHandlerExternalHandlerProcessMap API List

#### GET /v2-beta/externalhandlerexternalhandlerprocessmaps

- input: null
- output: collection of [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

#### GET /v2-beta/externalhandlerexternalhandlerprocessmaps/*resource-id*

- input: null
- output: [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

#### Action activate of schema object externalHandlerExternalHandlerProcessMap, POST /v2-beta/externalhandlerexternalhandlerprocessmaps/*resource-id*?action=activate

- input: null
- output: [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

#### Action create of schema object externalHandlerExternalHandlerProcessMap, POST /v2-beta/externalhandlerexternalhandlerprocessmaps/*resource-id*?action=create

- input: null
- output: [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

#### Action update of schema object externalHandlerExternalHandlerProcessMap, POST /v2-beta/externalhandlerexternalhandlerprocessmaps/*resource-id*?action=update

- input: null
- output: [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

#### Action purge of schema object externalHandlerExternalHandlerProcessMap, POST /v2-beta/externalhandlerexternalhandlerprocessmaps/*resource-id*?action=purge

- input: null
- output: [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

#### Action remove of schema object externalHandlerExternalHandlerProcessMap, POST /v2-beta/externalhandlerexternalhandlerprocessmaps/*resource-id*?action=remove

- input: null
- output: [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

#### Action deactivate of schema object externalHandlerExternalHandlerProcessMap, POST /v2-beta/externalhandlerexternalhandlerprocessmaps/*resource-id*?action=deactivate

- input: null
- output: [externalHandlerExternalHandlerProcessMap](#externalhandlerexternalhandlerprocessmap)

### Schema externalHandlerProcess API List

#### GET /v2-beta/externalhandlerprocesses

- input: null
- output: collection of [externalHandlerProcess](#externalhandlerprocess)

#### GET /v2-beta/externalhandlerprocesses/*resource-id*

- input: null
- output: [externalHandlerProcess](#externalhandlerprocess)

#### Action purge of schema object externalHandlerProcess, POST /v2-beta/externalhandlerprocesses/*resource-id*?action=purge

- input: null
- output: [externalHandlerProcess](#externalhandlerprocess)

#### Action remove of schema object externalHandlerProcess, POST /v2-beta/externalhandlerprocesses/*resource-id*?action=remove

- input: null
- output: [externalHandlerProcess](#externalhandlerprocess)

#### Action deactivate of schema object externalHandlerProcess, POST /v2-beta/externalhandlerprocesses/*resource-id*?action=deactivate

- input: null
- output: [externalHandlerProcess](#externalhandlerprocess)

#### Action activate of schema object externalHandlerProcess, POST /v2-beta/externalhandlerprocesses/*resource-id*?action=activate

- input: null
- output: [externalHandlerProcess](#externalhandlerprocess)

#### Action create of schema object externalHandlerProcess, POST /v2-beta/externalhandlerprocesses/*resource-id*?action=create

- input: null
- output: [externalHandlerProcess](#externalhandlerprocess)

#### Action update of schema object externalHandlerProcess, POST /v2-beta/externalhandlerprocesses/*resource-id*?action=update

- input: null
- output: [externalHandlerProcess](#externalhandlerprocess)

### Schema externalHostEvent API List

#### GET /v2-beta/externalhostevents

- input: null
- output: collection of [externalHostEvent](#externalhostevent)

#### GET /v2-beta/externalhostevents/*resource-id*

- input: null
- output: [externalHostEvent](#externalhostevent)

#### Action create of schema object externalHostEvent, POST /v2-beta/externalhostevents/*resource-id*?action=create

- input: null
- output: [externalEvent](#externalevent)

#### Action remove of schema object externalHostEvent, POST /v2-beta/externalhostevents/*resource-id*?action=remove

- input: null
- output: [externalEvent](#externalevent)

### Schema externalService API List

#### GET /v2-beta/externalservices

- input: null
- output: collection of [externalService](#externalservice)

#### GET /v2-beta/externalservices/*resource-id*

- input: null
- output: [externalService](#externalservice)

#### Action update of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=update

- input: null
- output: [service](#service)

#### Action deactivate of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=deactivate

- input: null
- output: [service](#service)

#### Action activate of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=activate

- input: null
- output: [service](#service)

#### Action create of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=create

- input: null
- output: [service](#service)

#### Action cancelupgrade of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action rollback of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=rollback

- input: null
- output: [service](#service)

#### Action upgrade of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=upgrade

- input: [serviceUpgrade](#serviceupgrade)
- output: [service](#service)

#### Action restart of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=restart

- input: [serviceRestart](#servicerestart)
- output: [service](#service)

#### Action remove of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action finishupgrade of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

#### Action continueupgrade of schema object externalService, POST /v2-beta/externalservices/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

### Schema externalServiceEvent API List

#### GET /v2-beta/externalserviceevents

- input: null
- output: collection of [externalServiceEvent](#externalserviceevent)

#### GET /v2-beta/externalserviceevents/*resource-id*

- input: null
- output: [externalServiceEvent](#externalserviceevent)

#### Action create of schema object externalServiceEvent, POST /v2-beta/externalserviceevents/*resource-id*?action=create

- input: null
- output: [externalEvent](#externalevent)

#### Action remove of schema object externalServiceEvent, POST /v2-beta/externalserviceevents/*resource-id*?action=remove

- input: null
- output: [externalEvent](#externalevent)

### Schema externalStoragePoolEvent API List

#### GET /v2-beta/externalstoragepoolevents

- input: null
- output: collection of [externalStoragePoolEvent](#externalstoragepoolevent)

#### GET /v2-beta/externalstoragepoolevents/*resource-id*

- input: null
- output: [externalStoragePoolEvent](#externalstoragepoolevent)

#### Action remove of schema object externalStoragePoolEvent, POST /v2-beta/externalstoragepoolevents/*resource-id*?action=remove

- input: null
- output: [externalEvent](#externalevent)

#### Action create of schema object externalStoragePoolEvent, POST /v2-beta/externalstoragepoolevents/*resource-id*?action=create

- input: null
- output: [externalEvent](#externalevent)

### Schema externalVolumeEvent API List

#### GET /v2-beta/externalvolumeevents

- input: null
- output: collection of [externalVolumeEvent](#externalvolumeevent)

#### GET /v2-beta/externalvolumeevents/*resource-id*

- input: null
- output: [externalVolumeEvent](#externalvolumeevent)

#### Action remove of schema object externalVolumeEvent, POST /v2-beta/externalvolumeevents/*resource-id*?action=remove

- input: null
- output: [externalEvent](#externalevent)

#### Action create of schema object externalVolumeEvent, POST /v2-beta/externalvolumeevents/*resource-id*?action=create

- input: null
- output: [externalEvent](#externalevent)

### Schema genericObject API List

#### GET /v2-beta/genericobjects

- input: null
- output: collection of [genericObject](#genericobject)

#### POST /v2-beta/genericobjects

- input: [genericObject](#genericobject)
- output: [genericObject](#genericobject)

#### GET /v2-beta/genericobjects/*resource-id*

- input: null
- output: [genericObject](#genericobject)

#### PUT /v2-beta/genericobjects/*resource-id*

- input: [genericObject](#genericobject)
- output: [genericObject](#genericobject)

#### DELETE /v2-beta/genericobjects/*resource-id*

- input: null
- output: [genericObject](#genericobject)

#### Action create of schema object genericObject, POST /v2-beta/genericobjects/*resource-id*?action=create

- input: null
- output: [genericObject](#genericobject)

#### Action remove of schema object genericObject, POST /v2-beta/genericobjects/*resource-id*?action=remove

- input: null
- output: [genericObject](#genericobject)

### Schema haConfig API List

#### GET /v2-beta/haconfigs

- input: null
- output: collection of [haConfig](#haconfig)

#### GET /v2-beta/haconfigs/*resource-id*

- input: null
- output: [haConfig](#haconfig)

#### PUT /v2-beta/haconfigs/*resource-id*

- input: [haConfig](#haconfig)
- output: [haConfig](#haconfig)

#### Action createscript of schema object haConfig, POST /v2-beta/haconfigs/*resource-id*?action=createscript

- input: [haConfigInput](#haconfiginput)
- output: null

### Schema haConfigInput API List

#### GET /v2-beta/haconfiginputs

- input: null
- output: collection of [haConfigInput](#haconfiginput)

#### POST /v2-beta/haconfiginputs

- input: [haConfigInput](#haconfiginput)
- output: [haConfigInput](#haconfiginput)

#### GET /v2-beta/haconfiginputs/*resource-id*

- input: null
- output: [haConfigInput](#haconfiginput)

### Schema healthcheckInstanceHostMap API List

#### GET /v2-beta/healthcheckinstancehostmaps

- input: null
- output: collection of [healthcheckInstanceHostMap](#healthcheckinstancehostmap)

#### GET /v2-beta/healthcheckinstancehostmaps/*resource-id*

- input: null
- output: [healthcheckInstanceHostMap](#healthcheckinstancehostmap)

#### Action create of schema object healthcheckInstanceHostMap, POST /v2-beta/healthcheckinstancehostmaps/*resource-id*?action=create

- input: null
- output: [healthcheckInstanceHostMap](#healthcheckinstancehostmap)

#### Action remove of schema object healthcheckInstanceHostMap, POST /v2-beta/healthcheckinstancehostmaps/*resource-id*?action=remove

- input: null
- output: [healthcheckInstanceHostMap](#healthcheckinstancehostmap)

### Schema host API List

#### GET /v2-beta/hosts

- input: null
- output: collection of [host](#host)

#### GET /v2-beta/hosts/*resource-id*

- input: null
- output: [host](#host)

#### Action evacuate of schema object host, POST /v2-beta/hosts/*resource-id*?action=evacuate

- input: null
- output: [host](#host)

#### Action provision of schema object host, POST /v2-beta/hosts/*resource-id*?action=provision

- input: null
- output: [host](#host)

#### Action dockersocket of schema object host, POST /v2-beta/hosts/*resource-id*?action=dockersocket

- input: null
- output: [hostAccess](#hostaccess)

#### Action activate of schema object host, POST /v2-beta/hosts/*resource-id*?action=activate

- input: null
- output: [host](#host)

#### Action update of schema object host, POST /v2-beta/hosts/*resource-id*?action=update

- input: null
- output: [host](#host)

#### Action error of schema object host, POST /v2-beta/hosts/*resource-id*?action=error

- input: null
- output: [host](#host)

#### Action remove of schema object host, POST /v2-beta/hosts/*resource-id*?action=remove

- input: null
- output: [host](#host)

#### Action upgrade of schema object host, POST /v2-beta/hosts/*resource-id*?action=upgrade

- input: null
- output: null

#### Action create of schema object host, POST /v2-beta/hosts/*resource-id*?action=create

- input: null
- output: [host](#host)

#### Action purge of schema object host, POST /v2-beta/hosts/*resource-id*?action=purge

- input: null
- output: [host](#host)

#### Action delete of schema object host, POST /v2-beta/hosts/*resource-id*?action=delete

- input: null
- output: null

#### Action deactivate of schema object host, POST /v2-beta/hosts/*resource-id*?action=deactivate

- input: null
- output: [host](#host)

### Schema hostApiProxyToken API List

#### GET /v2-beta/hostapiproxytokens

- input: null
- output: collection of [hostApiProxyToken](#hostapiproxytoken)

#### GET /v2-beta/hostapiproxytokens/*resource-id*

- input: null
- output: [hostApiProxyToken](#hostapiproxytoken)

### Schema hostTemplate API List

#### GET /v2-beta/hosttemplates

- input: null
- output: collection of [hostTemplate](#hosttemplate)

#### GET /v2-beta/hosttemplates/*resource-id*

- input: null
- output: [hostTemplate](#hosttemplate)

#### Action create of schema object hostTemplate, POST /v2-beta/hosttemplates/*resource-id*?action=create

- input: null
- output: [hostTemplate](#hosttemplate)

#### Action remove of schema object hostTemplate, POST /v2-beta/hosttemplates/*resource-id*?action=remove

- input: null
- output: [hostTemplate](#hosttemplate)

### Schema identity API List

#### GET /v2-beta/identities

- input: null
- output: collection of [identity](#identity)

#### GET /v2-beta/identities/*resource-id*

- input: null
- output: [identity](#identity)

### Schema image API List

#### GET /v2-beta/images

- input: null
- output: collection of [image](#image)

#### GET /v2-beta/images/*resource-id*

- input: null
- output: [image](#image)

#### Action activate of schema object image, POST /v2-beta/images/*resource-id*?action=activate

- input: null
- output: [image](#image)

#### Action create of schema object image, POST /v2-beta/images/*resource-id*?action=create

- input: null
- output: [image](#image)

#### Action update of schema object image, POST /v2-beta/images/*resource-id*?action=update

- input: null
- output: [image](#image)

#### Action purge of schema object image, POST /v2-beta/images/*resource-id*?action=purge

- input: null
- output: [image](#image)

#### Action remove of schema object image, POST /v2-beta/images/*resource-id*?action=remove

- input: null
- output: [image](#image)

#### Action deactivate of schema object image, POST /v2-beta/images/*resource-id*?action=deactivate

- input: null
- output: [image](#image)

### Schema instance API List

#### GET /v2-beta/instances

- input: null
- output: collection of [instance](#instance)

#### GET /v2-beta/instances/*resource-id*

- input: null
- output: [instance](#instance)

#### Action purge of schema object instance, POST /v2-beta/instances/*resource-id*?action=purge

- input: null
- output: [instance](#instance)

#### Action updatehealthy of schema object instance, POST /v2-beta/instances/*resource-id*?action=updatehealthy

- input: null
- output: [instance](#instance)

#### Action migrate of schema object instance, POST /v2-beta/instances/*resource-id*?action=migrate

- input: null
- output: [instance](#instance)

#### Action updateunhealthy of schema object instance, POST /v2-beta/instances/*resource-id*?action=updateunhealthy

- input: null
- output: [instance](#instance)

#### Action restart of schema object instance, POST /v2-beta/instances/*resource-id*?action=restart

- input: null
- output: [instance](#instance)

#### Action updatereinitializing of schema object instance, POST /v2-beta/instances/*resource-id*?action=updatereinitializing

- input: null
- output: [instance](#instance)

#### Action stop of schema object instance, POST /v2-beta/instances/*resource-id*?action=stop

- input: [instanceStop](#instancestop)
- output: [instance](#instance)

#### Action allocate of schema object instance, POST /v2-beta/instances/*resource-id*?action=allocate

- input: null
- output: [instance](#instance)

#### Action console of schema object instance, POST /v2-beta/instances/*resource-id*?action=console

- input: [instanceConsoleInput](#instanceconsoleinput)
- output: [instanceConsole](#instanceconsole)

#### Action deallocate of schema object instance, POST /v2-beta/instances/*resource-id*?action=deallocate

- input: null
- output: [instance](#instance)

#### Action start of schema object instance, POST /v2-beta/instances/*resource-id*?action=start

- input: null
- output: [instance](#instance)

#### Action update of schema object instance, POST /v2-beta/instances/*resource-id*?action=update

- input: null
- output: [instance](#instance)

#### Action error of schema object instance, POST /v2-beta/instances/*resource-id*?action=error

- input: null
- output: [instance](#instance)

#### Action remove of schema object instance, POST /v2-beta/instances/*resource-id*?action=remove

- input: null
- output: [instance](#instance)

#### Action create of schema object instance, POST /v2-beta/instances/*resource-id*?action=create

- input: null
- output: [instance](#instance)

### Schema instanceLink API List

#### GET /v2-beta/instancelinks

- input: null
- output: collection of [instanceLink](#instancelink)

#### GET /v2-beta/instancelinks/*resource-id*

- input: null
- output: [instanceLink](#instancelink)

#### Action activate of schema object instanceLink, POST /v2-beta/instancelinks/*resource-id*?action=activate

- input: null
- output: [instanceLink](#instancelink)

#### Action create of schema object instanceLink, POST /v2-beta/instancelinks/*resource-id*?action=create

- input: null
- output: [instanceLink](#instancelink)

#### Action update of schema object instanceLink, POST /v2-beta/instancelinks/*resource-id*?action=update

- input: null
- output: [instanceLink](#instancelink)

#### Action purge of schema object instanceLink, POST /v2-beta/instancelinks/*resource-id*?action=purge

- input: null
- output: [instanceLink](#instancelink)

#### Action remove of schema object instanceLink, POST /v2-beta/instancelinks/*resource-id*?action=remove

- input: null
- output: [instanceLink](#instancelink)

#### Action deactivate of schema object instanceLink, POST /v2-beta/instancelinks/*resource-id*?action=deactivate

- input: null
- output: [instanceLink](#instancelink)

### Schema ipAddress API List

#### GET /v2-beta/ipaddresses

- input: null
- output: collection of [ipAddress](#ipaddress)

#### GET /v2-beta/ipaddresses/*resource-id*

- input: null
- output: [ipAddress](#ipaddress)

#### Action update of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=update

- input: null
- output: [ipAddress](#ipaddress)

#### Action purge of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=purge

- input: null
- output: [ipAddress](#ipaddress)

#### Action remove of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=remove

- input: null
- output: [ipAddress](#ipaddress)

#### Action associate of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=associate

- input: null
- output: [ipAddress](#ipaddress)

#### Action deactivate of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=deactivate

- input: null
- output: [ipAddress](#ipaddress)

#### Action disassociate of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=disassociate

- input: null
- output: [ipAddress](#ipaddress)

#### Action activate of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=activate

- input: null
- output: [ipAddress](#ipaddress)

#### Action create of schema object ipAddress, POST /v2-beta/ipaddresses/*resource-id*?action=create

- input: null
- output: [ipAddress](#ipaddress)

### Schema kubernetesService API List

#### GET /v2-beta/kubernetesservices

- input: null
- output: collection of [kubernetesService](#kubernetesservice)

#### GET /v2-beta/kubernetesservices/*resource-id*

- input: null
- output: [kubernetesService](#kubernetesservice)

#### Action rollback of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=rollback

- input: null
- output: [service](#service)

#### Action remove of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action setservicelinks of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=setservicelinks

- input: [setServiceLinksInput](#setservicelinksinput)
- output: [service](#service)

#### Action restart of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=restart

- input: [serviceRestart](#servicerestart)
- output: [service](#service)

#### Action update of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=update

- input: null
- output: [service](#service)

#### Action activate of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=activate

- input: null
- output: [service](#service)

#### Action create of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=create

- input: null
- output: [service](#service)

#### Action cancelupgrade of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action upgrade of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=upgrade

- input: [serviceUpgrade](#serviceupgrade)
- output: [service](#service)

#### Action removeservicelink of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=removeservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action continueupgrade of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

#### Action deactivate of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=deactivate

- input: null
- output: [service](#service)

#### Action finishupgrade of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

#### Action addservicelink of schema object kubernetesService, POST /v2-beta/kubernetesservices/*resource-id*?action=addservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

### Schema kubernetesStack API List

#### GET /v2-beta/kubernetesstacks

- input: null
- output: collection of [kubernetesStack](#kubernetesstack)

#### GET /v2-beta/kubernetesstacks/*resource-id*

- input: null
- output: [kubernetesStack](#kubernetesstack)

#### Action upgrade of schema object kubernetesStack, POST /v2-beta/kubernetesstacks/*resource-id*?action=upgrade

- input: [kubernetesStackUpgrade](#kubernetesstackupgrade)
- output: [kubernetesStack](#kubernetesstack)

#### Action create of schema object kubernetesStack, POST /v2-beta/kubernetesstacks/*resource-id*?action=create

- input: null
- output: [stack](#stack)

#### Action error of schema object kubernetesStack, POST /v2-beta/kubernetesstacks/*resource-id*?action=error

- input: null
- output: [stack](#stack)

#### Action cancelupgrade of schema object kubernetesStack, POST /v2-beta/kubernetesstacks/*resource-id*?action=cancelupgrade

- input: null
- output: [stack](#stack)

#### Action remove of schema object kubernetesStack, POST /v2-beta/kubernetesstacks/*resource-id*?action=remove

- input: null
- output: [stack](#stack)

#### Action finishupgrade of schema object kubernetesStack, POST /v2-beta/kubernetesstacks/*resource-id*?action=finishupgrade

- input: null
- output: [stack](#stack)

#### Action rollback of schema object kubernetesStack, POST /v2-beta/kubernetesstacks/*resource-id*?action=rollback

- input: null
- output: [stack](#stack)

### Schema label API List

#### GET /v2-beta/labels

- input: null
- output: collection of [label](#label)

#### GET /v2-beta/labels/*resource-id*

- input: null
- output: [label](#label)

#### Action create of schema object label, POST /v2-beta/labels/*resource-id*?action=create

- input: null
- output: [label](#label)

#### Action remove of schema object label, POST /v2-beta/labels/*resource-id*?action=remove

- input: null
- output: [label](#label)

### Schema loadBalancerService API List

#### GET /v2-beta/loadbalancerservices

- input: null
- output: collection of [loadBalancerService](#loadbalancerservice)

#### GET /v2-beta/loadbalancerservices/*resource-id*

- input: null
- output: [loadBalancerService](#loadbalancerservice)

#### Action remove of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action continueupgrade of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

#### Action cancelupgrade of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action upgrade of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=upgrade

- input: [serviceUpgrade](#serviceupgrade)
- output: [service](#service)

#### Action restart of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=restart

- input: [serviceRestart](#servicerestart)
- output: [service](#service)

#### Action deactivate of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=deactivate

- input: null
- output: [service](#service)

#### Action addservicelink of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=addservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action setservicelinks of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=setservicelinks

- input: [setServiceLinksInput](#setservicelinksinput)
- output: [service](#service)

#### Action removeservicelink of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=removeservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action create of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=create

- input: null
- output: [service](#service)

#### Action rollback of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=rollback

- input: null
- output: [service](#service)

#### Action update of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=update

- input: null
- output: [service](#service)

#### Action finishupgrade of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

#### Action activate of schema object loadBalancerService, POST /v2-beta/loadbalancerservices/*resource-id*?action=activate

- input: null
- output: [service](#service)

### Schema localAuthConfig API List

#### GET /v2-beta/localauthconfigs

- input: null
- output: collection of [localAuthConfig](#localauthconfig)

#### POST /v2-beta/localauthconfigs

- input: [localAuthConfig](#localauthconfig)
- output: [localAuthConfig](#localauthconfig)

#### GET /v2-beta/localauthconfigs/*resource-id*

- input: null
- output: [localAuthConfig](#localauthconfig)

### Schema machine API List

#### GET /v2-beta/machines

- input: null
- output: collection of [machine](#machine)

#### GET /v2-beta/machines/*resource-id*

- input: null
- output: [machine](#machine)

#### Action update of schema object machine, POST /v2-beta/machines/*resource-id*?action=update

- input: null
- output: [physicalHost](#physicalhost)

#### Action bootstrap of schema object machine, POST /v2-beta/machines/*resource-id*?action=bootstrap

- input: null
- output: [physicalHost](#physicalhost)

#### Action error of schema object machine, POST /v2-beta/machines/*resource-id*?action=error

- input: null
- output: [physicalHost](#physicalhost)

#### Action remove of schema object machine, POST /v2-beta/machines/*resource-id*?action=remove

- input: null
- output: [physicalHost](#physicalhost)

#### Action create of schema object machine, POST /v2-beta/machines/*resource-id*?action=create

- input: null
- output: [physicalHost](#physicalhost)

### Schema machineDriver API List

#### GET /v2-beta/machinedrivers

- input: null
- output: collection of [machineDriver](#machinedriver)

#### POST /v2-beta/machinedrivers

- input: [machineDriver](#machinedriver)
- output: [machineDriver](#machinedriver)

#### GET /v2-beta/machinedrivers/*resource-id*

- input: null
- output: [machineDriver](#machinedriver)

#### PUT /v2-beta/machinedrivers/*resource-id*

- input: [machineDriver](#machinedriver)
- output: [machineDriver](#machinedriver)

#### DELETE /v2-beta/machinedrivers/*resource-id*

- input: null
- output: [machineDriver](#machinedriver)

#### Action create of schema object machineDriver, POST /v2-beta/machinedrivers/*resource-id*?action=create

- input: null
- output: [machineDriver](#machinedriver)

#### Action update of schema object machineDriver, POST /v2-beta/machinedrivers/*resource-id*?action=update

- input: null
- output: [machineDriver](#machinedriver)

#### Action error of schema object machineDriver, POST /v2-beta/machinedrivers/*resource-id*?action=error

- input: null
- output: [machineDriver](#machinedriver)

#### Action remove of schema object machineDriver, POST /v2-beta/machinedrivers/*resource-id*?action=remove

- input: null
- output: [machineDriver](#machinedriver)

#### Action deactivate of schema object machineDriver, POST /v2-beta/machinedrivers/*resource-id*?action=deactivate

- input: null
- output: [machineDriver](#machinedriver)

#### Action reactivate of schema object machineDriver, POST /v2-beta/machinedrivers/*resource-id*?action=reactivate

- input: null
- output: [machineDriver](#machinedriver)

#### Action activate of schema object machineDriver, POST /v2-beta/machinedrivers/*resource-id*?action=activate

- input: null
- output: [machineDriver](#machinedriver)

### Schema mount API List

#### GET /v2-beta/mounts

- input: null
- output: collection of [mount](#mount)

#### GET /v2-beta/mounts/*resource-id*

- input: null
- output: [mount](#mount)

#### Action create of schema object mount, POST /v2-beta/mounts/*resource-id*?action=create

- input: null
- output: [mount](#mount)

#### Action remove of schema object mount, POST /v2-beta/mounts/*resource-id*?action=remove

- input: null
- output: [mount](#mount)

#### Action deactivate of schema object mount, POST /v2-beta/mounts/*resource-id*?action=deactivate

- input: null
- output: [mount](#mount)

### Schema network API List

#### GET /v2-beta/networks

- input: null
- output: collection of [network](#network)

#### GET /v2-beta/networks/*resource-id*

- input: null
- output: [network](#network)

#### Action purge of schema object network, POST /v2-beta/networks/*resource-id*?action=purge

- input: null
- output: [network](#network)

#### Action remove of schema object network, POST /v2-beta/networks/*resource-id*?action=remove

- input: null
- output: [network](#network)

#### Action deactivate of schema object network, POST /v2-beta/networks/*resource-id*?action=deactivate

- input: null
- output: [network](#network)

#### Action activate of schema object network, POST /v2-beta/networks/*resource-id*?action=activate

- input: null
- output: [network](#network)

#### Action create of schema object network, POST /v2-beta/networks/*resource-id*?action=create

- input: null
- output: [network](#network)

#### Action update of schema object network, POST /v2-beta/networks/*resource-id*?action=update

- input: null
- output: [network](#network)

### Schema networkDriver API List

#### GET /v2-beta/networkdrivers

- input: null
- output: collection of [networkDriver](#networkdriver)

#### GET /v2-beta/networkdrivers/*resource-id*

- input: null
- output: [networkDriver](#networkdriver)

#### Action activate of schema object networkDriver, POST /v2-beta/networkdrivers/*resource-id*?action=activate

- input: null
- output: [networkDriver](#networkdriver)

#### Action create of schema object networkDriver, POST /v2-beta/networkdrivers/*resource-id*?action=create

- input: null
- output: [networkDriver](#networkdriver)

#### Action update of schema object networkDriver, POST /v2-beta/networkdrivers/*resource-id*?action=update

- input: null
- output: [networkDriver](#networkdriver)

#### Action remove of schema object networkDriver, POST /v2-beta/networkdrivers/*resource-id*?action=remove

- input: null
- output: [networkDriver](#networkdriver)

#### Action deactivate of schema object networkDriver, POST /v2-beta/networkdrivers/*resource-id*?action=deactivate

- input: null
- output: [networkDriver](#networkdriver)

### Schema networkDriverService API List

#### GET /v2-beta/networkdriverservices

- input: null
- output: collection of [networkDriverService](#networkdriverservice)

#### GET /v2-beta/networkdriverservices/*resource-id*

- input: null
- output: [networkDriverService](#networkdriverservice)

#### Action activate of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=activate

- input: null
- output: [service](#service)

#### Action addservicelink of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=addservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action create of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=create

- input: null
- output: [service](#service)

#### Action upgrade of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=upgrade

- input: [serviceUpgrade](#serviceupgrade)
- output: [service](#service)

#### Action restart of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=restart

- input: [serviceRestart](#servicerestart)
- output: [service](#service)

#### Action update of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=update

- input: null
- output: [service](#service)

#### Action cancelupgrade of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action remove of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action finishupgrade of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

#### Action continueupgrade of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

#### Action setservicelinks of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=setservicelinks

- input: [setServiceLinksInput](#setservicelinksinput)
- output: [service](#service)

#### Action rollback of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=rollback

- input: null
- output: [service](#service)

#### Action deactivate of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=deactivate

- input: null
- output: [service](#service)

#### Action removeservicelink of schema object networkDriverService, POST /v2-beta/networkdriverservices/*resource-id*?action=removeservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

### Schema networkPolicyRuleWithin API List

#### GET /v2-beta/networkpolicyrulewithins

- input: null
- output: collection of [networkPolicyRuleWithin](#networkpolicyrulewithin)

#### POST /v2-beta/networkpolicyrulewithins

- input: [networkPolicyRuleWithin](#networkpolicyrulewithin)
- output: [networkPolicyRuleWithin](#networkpolicyrulewithin)

#### GET /v2-beta/networkpolicyrulewithins/*resource-id*

- input: null
- output: [networkPolicyRuleWithin](#networkpolicyrulewithin)

### Schema openldapconfig API List

#### GET /v2-beta/openldapconfigs

- input: null
- output: collection of [openldapconfig](#openldapconfig)

#### POST /v2-beta/openldapconfigs

- input: [openldapconfig](#openldapconfig)
- output: [openldapconfig](#openldapconfig)

#### GET /v2-beta/openldapconfigs/*resource-id*

- input: null
- output: [openldapconfig](#openldapconfig)

#### PUT /v2-beta/openldapconfigs/*resource-id*

- input: [openldapconfig](#openldapconfig)
- output: [openldapconfig](#openldapconfig)

### Schema password API List

#### GET /v2-beta/passwords

- input: null
- output: collection of [password](#password)

#### POST /v2-beta/passwords

- input: [password](#password)
- output: [password](#password)

#### GET /v2-beta/passwords/*resource-id*

- input: null
- output: [password](#password)

#### PUT /v2-beta/passwords/*resource-id*

- input: [password](#password)
- output: [password](#password)

#### DELETE /v2-beta/passwords/*resource-id*

- input: null
- output: [password](#password)

#### Action create of schema object password, POST /v2-beta/passwords/*resource-id*?action=create

- input: null
- output: [credential](#credential)

#### Action update of schema object password, POST /v2-beta/passwords/*resource-id*?action=update

- input: null
- output: [credential](#credential)

#### Action purge of schema object password, POST /v2-beta/passwords/*resource-id*?action=purge

- input: null
- output: [credential](#credential)

#### Action remove of schema object password, POST /v2-beta/passwords/*resource-id*?action=remove

- input: null
- output: [credential](#credential)

#### Action deactivate of schema object password, POST /v2-beta/passwords/*resource-id*?action=deactivate

- input: null
- output: [credential](#credential)

#### Action changesecret of schema object password, POST /v2-beta/passwords/*resource-id*?action=changesecret

- input: [changeSecretInput](#changesecretinput)
- output: [changeSecretInput](#changesecretinput)

#### Action activate of schema object password, POST /v2-beta/passwords/*resource-id*?action=activate

- input: null
- output: [credential](#credential)

### Schema physicalHost API List

#### GET /v2-beta/physicalhosts

- input: null
- output: collection of [physicalHost](#physicalhost)

#### GET /v2-beta/physicalhosts/*resource-id*

- input: null
- output: [physicalHost](#physicalhost)

#### Action create of schema object physicalHost, POST /v2-beta/physicalhosts/*resource-id*?action=create

- input: null
- output: [physicalHost](#physicalhost)

#### Action update of schema object physicalHost, POST /v2-beta/physicalhosts/*resource-id*?action=update

- input: null
- output: [physicalHost](#physicalhost)

#### Action bootstrap of schema object physicalHost, POST /v2-beta/physicalhosts/*resource-id*?action=bootstrap

- input: null
- output: [physicalHost](#physicalhost)

#### Action error of schema object physicalHost, POST /v2-beta/physicalhosts/*resource-id*?action=error

- input: null
- output: [physicalHost](#physicalhost)

#### Action remove of schema object physicalHost, POST /v2-beta/physicalhosts/*resource-id*?action=remove

- input: null
- output: [physicalHost](#physicalhost)

### Schema port API List

#### GET /v2-beta/ports

- input: null
- output: collection of [port](#port)

#### GET /v2-beta/ports/*resource-id*

- input: null
- output: [port](#port)

#### Action deactivate of schema object port, POST /v2-beta/ports/*resource-id*?action=deactivate

- input: null
- output: [port](#port)

#### Action activate of schema object port, POST /v2-beta/ports/*resource-id*?action=activate

- input: null
- output: [port](#port)

#### Action create of schema object port, POST /v2-beta/ports/*resource-id*?action=create

- input: null
- output: [port](#port)

#### Action update of schema object port, POST /v2-beta/ports/*resource-id*?action=update

- input: null
- output: [port](#port)

#### Action purge of schema object port, POST /v2-beta/ports/*resource-id*?action=purge

- input: null
- output: [port](#port)

#### Action remove of schema object port, POST /v2-beta/ports/*resource-id*?action=remove

- input: null
- output: [port](#port)

### Schema processDefinition API List

#### GET /v2-beta/processdefinitions

- input: null
- output: collection of [processDefinition](#processdefinition)

#### GET /v2-beta/processdefinitions/*resource-id*

- input: null
- output: [processDefinition](#processdefinition)

### Schema processExecution API List

#### GET /v2-beta/processexecutions

- input: null
- output: collection of [processExecution](#processexecution)

#### GET /v2-beta/processexecutions/*resource-id*

- input: null
- output: [processExecution](#processexecution)

### Schema processInstance API List

#### GET /v2-beta/processinstances

- input: null
- output: collection of [processInstance](#processinstance)

#### GET /v2-beta/processinstances/*resource-id*

- input: null
- output: [processInstance](#processinstance)

#### Action replay of schema object processInstance, POST /v2-beta/processinstances/*resource-id*?action=replay

- input: null
- output: [processInstance](#processinstance)

### Schema processPool API List

#### GET /v2-beta/processpools

- input: null
- output: collection of [processPool](#processpool)

#### GET /v2-beta/processpools/*resource-id*

- input: null
- output: [processPool](#processpool)

### Schema processSummary API List

#### GET /v2-beta/processsummary

- input: null
- output: collection of [processSummary](#processsummary)

#### GET /v2-beta/processsummary/*resource-id*

- input: null
- output: [processSummary](#processsummary)

### Schema project API List

#### GET /v2-beta/projects

- input: null
- output: collection of [project](#project)

#### POST /v2-beta/projects

- input: [project](#project)
- output: [project](#project)

#### GET /v2-beta/projects/*resource-id*

- input: null
- output: [project](#project)

#### PUT /v2-beta/projects/*resource-id*

- input: [project](#project)
- output: [project](#project)

#### DELETE /v2-beta/projects/*resource-id*

- input: null
- output: [project](#project)

#### Action activate of schema object project, POST /v2-beta/projects/*resource-id*?action=activate

- input: null
- output: [account](#account)

#### Action create of schema object project, POST /v2-beta/projects/*resource-id*?action=create

- input: null
- output: [account](#account)

#### Action purge of schema object project, POST /v2-beta/projects/*resource-id*?action=purge

- input: null
- output: [account](#account)

#### Action remove of schema object project, POST /v2-beta/projects/*resource-id*?action=remove

- input: null
- output: [account](#account)

#### Action deactivate of schema object project, POST /v2-beta/projects/*resource-id*?action=deactivate

- input: null
- output: [account](#account)

#### Action defaultNetworkId of schema object project, POST /v2-beta/projects/*resource-id*?action=defaultNetworkId

- input: null
- output: null

#### Action upgrade of schema object project, POST /v2-beta/projects/*resource-id*?action=upgrade

- input: null
- output: [account](#account)

#### Action update of schema object project, POST /v2-beta/projects/*resource-id*?action=update

- input: null
- output: [account](#account)

#### Action setmembers of schema object project, POST /v2-beta/projects/*resource-id*?action=setmembers

- input: [setProjectMembersInput](#setprojectmembersinput)
- output: [setProjectMembersInput](#setprojectmembersinput)

#### Action delete of schema object project, POST /v2-beta/projects/*resource-id*?action=delete

- input: null
- output: null

### Schema projectMember API List

#### GET /v2-beta/projectmembers

- input: null
- output: collection of [projectMember](#projectmember)

#### POST /v2-beta/projectmembers

- input: [projectMember](#projectmember)
- output: [projectMember](#projectmember)

#### GET /v2-beta/projectmembers/*resource-id*

- input: null
- output: [projectMember](#projectmember)

#### Action activate of schema object projectMember, POST /v2-beta/projectmembers/*resource-id*?action=activate

- input: null
- output: [projectMember](#projectmember)

#### Action create of schema object projectMember, POST /v2-beta/projectmembers/*resource-id*?action=create

- input: null
- output: [projectMember](#projectmember)

#### Action update of schema object projectMember, POST /v2-beta/projectmembers/*resource-id*?action=update

- input: null
- output: [projectMember](#projectmember)

#### Action purge of schema object projectMember, POST /v2-beta/projectmembers/*resource-id*?action=purge

- input: null
- output: [projectMember](#projectmember)

#### Action remove of schema object projectMember, POST /v2-beta/projectmembers/*resource-id*?action=remove

- input: null
- output: [projectMember](#projectmember)

#### Action deactivate of schema object projectMember, POST /v2-beta/projectmembers/*resource-id*?action=deactivate

- input: null
- output: [projectMember](#projectmember)

### Schema projectTemplate API List

#### GET /v2-beta/projecttemplates

- input: null
- output: collection of [projectTemplate](#projecttemplate)

#### POST /v2-beta/projecttemplates

- input: [projectTemplate](#projecttemplate)
- output: [projectTemplate](#projecttemplate)

#### GET /v2-beta/projecttemplates/*resource-id*

- input: null
- output: [projectTemplate](#projecttemplate)

#### PUT /v2-beta/projecttemplates/*resource-id*

- input: [projectTemplate](#projecttemplate)
- output: [projectTemplate](#projecttemplate)

#### DELETE /v2-beta/projecttemplates/*resource-id*

- input: null
- output: [projectTemplate](#projecttemplate)

#### Action create of schema object projectTemplate, POST /v2-beta/projecttemplates/*resource-id*?action=create

- input: null
- output: [projectTemplate](#projecttemplate)

#### Action remove of schema object projectTemplate, POST /v2-beta/projecttemplates/*resource-id*?action=remove

- input: null
- output: [projectTemplate](#projecttemplate)

### Schema pullTask API List

#### GET /v2-beta/pulltasks

- input: null
- output: collection of [pullTask](#pulltask)

#### GET /v2-beta/pulltasks/*resource-id*

- input: null
- output: [pullTask](#pulltask)

#### Action create of schema object pullTask, POST /v2-beta/pulltasks/*resource-id*?action=create

- input: null
- output: [genericObject](#genericobject)

#### Action remove of schema object pullTask, POST /v2-beta/pulltasks/*resource-id*?action=remove

- input: null
- output: [genericObject](#genericobject)

### Schema region API List

#### GET /v2-beta/regions

- input: null
- output: collection of [region](#region)

#### POST /v2-beta/regions

- input: [region](#region)
- output: [region](#region)

#### GET /v2-beta/regions/*resource-id*

- input: null
- output: [region](#region)

#### PUT /v2-beta/regions/*resource-id*

- input: [region](#region)
- output: [region](#region)

#### DELETE /v2-beta/regions/*resource-id*

- input: null
- output: [region](#region)

#### Action create of schema object region, POST /v2-beta/regions/*resource-id*?action=create

- input: null
- output: [region](#region)

#### Action update of schema object region, POST /v2-beta/regions/*resource-id*?action=update

- input: null
- output: [region](#region)

#### Action purge of schema object region, POST /v2-beta/regions/*resource-id*?action=purge

- input: null
- output: [region](#region)

#### Action remove of schema object region, POST /v2-beta/regions/*resource-id*?action=remove

- input: null
- output: [region](#region)

#### Action deactivate of schema object region, POST /v2-beta/regions/*resource-id*?action=deactivate

- input: null
- output: [region](#region)

#### Action activate of schema object region, POST /v2-beta/regions/*resource-id*?action=activate

- input: null
- output: [region](#region)

### Schema register API List

#### GET /v2-beta/register

- input: null
- output: collection of [register](#register)

#### GET /v2-beta/register/*resource-id*

- input: null
- output: [register](#register)

#### Action create of schema object register, POST /v2-beta/register/*resource-id*?action=create

- input: null
- output: [genericObject](#genericobject)

#### Action remove of schema object register, POST /v2-beta/register/*resource-id*?action=remove

- input: null
- output: [genericObject](#genericobject)

#### Action stop of schema object register, POST /v2-beta/register/*resource-id*?action=stop

- input: [instanceStop](#instancestop)
- output: [instance](#instance)

### Schema registrationToken API List

#### GET /v2-beta/registrationtokens

- input: null
- output: collection of [registrationToken](#registrationtoken)

#### GET /v2-beta/registrationtokens/*resource-id*

- input: null
- output: [registrationToken](#registrationtoken)

#### Action activate of schema object registrationToken, POST /v2-beta/registrationtokens/*resource-id*?action=activate

- input: null
- output: [credential](#credential)

#### Action create of schema object registrationToken, POST /v2-beta/registrationtokens/*resource-id*?action=create

- input: null
- output: [credential](#credential)

#### Action update of schema object registrationToken, POST /v2-beta/registrationtokens/*resource-id*?action=update

- input: null
- output: [credential](#credential)

#### Action purge of schema object registrationToken, POST /v2-beta/registrationtokens/*resource-id*?action=purge

- input: null
- output: [credential](#credential)

#### Action remove of schema object registrationToken, POST /v2-beta/registrationtokens/*resource-id*?action=remove

- input: null
- output: [credential](#credential)

#### Action deactivate of schema object registrationToken, POST /v2-beta/registrationtokens/*resource-id*?action=deactivate

- input: null
- output: [credential](#credential)

### Schema registry API List

#### GET /v2-beta/registries

- input: null
- output: collection of [registry](#registry)

#### GET /v2-beta/registries/*resource-id*

- input: null
- output: [registry](#registry)

#### Action activate of schema object registry, POST /v2-beta/registries/*resource-id*?action=activate

- input: null
- output: [storagePool](#storagepool)

#### Action create of schema object registry, POST /v2-beta/registries/*resource-id*?action=create

- input: null
- output: [storagePool](#storagepool)

#### Action update of schema object registry, POST /v2-beta/registries/*resource-id*?action=update

- input: null
- output: [storagePool](#storagepool)

#### Action purge of schema object registry, POST /v2-beta/registries/*resource-id*?action=purge

- input: null
- output: [storagePool](#storagepool)

#### Action remove of schema object registry, POST /v2-beta/registries/*resource-id*?action=remove

- input: null
- output: [storagePool](#storagepool)

#### Action deactivate of schema object registry, POST /v2-beta/registries/*resource-id*?action=deactivate

- input: null
- output: [storagePool](#storagepool)

### Schema registryCredential API List

#### GET /v2-beta/registrycredentials

- input: null
- output: collection of [registryCredential](#registrycredential)

#### GET /v2-beta/registrycredentials/*resource-id*

- input: null
- output: [registryCredential](#registrycredential)

#### Action deactivate of schema object registryCredential, POST /v2-beta/registrycredentials/*resource-id*?action=deactivate

- input: null
- output: [credential](#credential)

#### Action activate of schema object registryCredential, POST /v2-beta/registrycredentials/*resource-id*?action=activate

- input: null
- output: [credential](#credential)

#### Action create of schema object registryCredential, POST /v2-beta/registrycredentials/*resource-id*?action=create

- input: null
- output: [credential](#credential)

#### Action update of schema object registryCredential, POST /v2-beta/registrycredentials/*resource-id*?action=update

- input: null
- output: [credential](#credential)

#### Action purge of schema object registryCredential, POST /v2-beta/registrycredentials/*resource-id*?action=purge

- input: null
- output: [credential](#credential)

#### Action remove of schema object registryCredential, POST /v2-beta/registrycredentials/*resource-id*?action=remove

- input: null
- output: [credential](#credential)

### Schema resourceDefinition API List

#### GET /v2-beta/resourcedefinitions

- input: null
- output: collection of [resourceDefinition](#resourcedefinition)

#### GET /v2-beta/resourcedefinitions/*resource-id*

- input: null
- output: [resourceDefinition](#resourcedefinition)

### Schema scheduledUpgrade API List

#### GET /v2-beta/scheduledupgrades

- input: null
- output: collection of [scheduledUpgrade](#scheduledupgrade)

#### GET /v2-beta/scheduledupgrades/*resource-id*

- input: null
- output: [scheduledUpgrade](#scheduledupgrade)

#### Action create of schema object scheduledUpgrade, POST /v2-beta/scheduledupgrades/*resource-id*?action=create

- input: null
- output: [scheduledUpgrade](#scheduledupgrade)

#### Action remove of schema object scheduledUpgrade, POST /v2-beta/scheduledupgrades/*resource-id*?action=remove

- input: null
- output: [scheduledUpgrade](#scheduledupgrade)

#### Action start of schema object scheduledUpgrade, POST /v2-beta/scheduledupgrades/*resource-id*?action=start

- input: null
- output: [scheduledUpgrade](#scheduledupgrade)

### Schema schema API List

#### GET /v2-beta/schemas

- input: null
- output: collection of [schema](#schema)

#### GET /v2-beta/schemas/*resource-id*

- input: null
- output: [schema](#schema)

### Schema secret API List

#### GET /v2-beta/secrets

- input: null
- output: collection of [secret](#secret)

#### GET /v2-beta/secrets/*resource-id*

- input: null
- output: [secret](#secret)

#### Action create of schema object secret, POST /v2-beta/secrets/*resource-id*?action=create

- input: null
- output: [secret](#secret)

#### Action remove of schema object secret, POST /v2-beta/secrets/*resource-id*?action=remove

- input: null
- output: [secret](#secret)

### Schema service API List

#### GET /v2-beta/services

- input: null
- output: collection of [service](#service)

#### GET /v2-beta/services/*resource-id*

- input: null
- output: [service](#service)

#### Action remove of schema object service, POST /v2-beta/services/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action removeservicelink of schema object service, POST /v2-beta/services/*resource-id*?action=removeservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action activate of schema object service, POST /v2-beta/services/*resource-id*?action=activate

- input: null
- output: [service](#service)

#### Action addservicelink of schema object service, POST /v2-beta/services/*resource-id*?action=addservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action upgrade of schema object service, POST /v2-beta/services/*resource-id*?action=upgrade

- input: [serviceUpgrade](#serviceupgrade)
- output: [service](#service)

#### Action create of schema object service, POST /v2-beta/services/*resource-id*?action=create

- input: null
- output: [service](#service)

#### Action cancelupgrade of schema object service, POST /v2-beta/services/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action deactivate of schema object service, POST /v2-beta/services/*resource-id*?action=deactivate

- input: null
- output: [service](#service)

#### Action finishupgrade of schema object service, POST /v2-beta/services/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

#### Action restart of schema object service, POST /v2-beta/services/*resource-id*?action=restart

- input: [serviceRestart](#servicerestart)
- output: [service](#service)

#### Action update of schema object service, POST /v2-beta/services/*resource-id*?action=update

- input: null
- output: [service](#service)

#### Action continueupgrade of schema object service, POST /v2-beta/services/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

#### Action setservicelinks of schema object service, POST /v2-beta/services/*resource-id*?action=setservicelinks

- input: [setServiceLinksInput](#setservicelinksinput)
- output: [service](#service)

#### Action rollback of schema object service, POST /v2-beta/services/*resource-id*?action=rollback

- input: null
- output: [service](#service)

### Schema serviceConsumeMap API List

#### GET /v2-beta/serviceconsumemaps

- input: null
- output: collection of [serviceConsumeMap](#serviceconsumemap)

#### GET /v2-beta/serviceconsumemaps/*resource-id*

- input: null
- output: [serviceConsumeMap](#serviceconsumemap)

#### Action create of schema object serviceConsumeMap, POST /v2-beta/serviceconsumemaps/*resource-id*?action=create

- input: null
- output: [serviceConsumeMap](#serviceconsumemap)

#### Action update of schema object serviceConsumeMap, POST /v2-beta/serviceconsumemaps/*resource-id*?action=update

- input: null
- output: [serviceConsumeMap](#serviceconsumemap)

#### Action remove of schema object serviceConsumeMap, POST /v2-beta/serviceconsumemaps/*resource-id*?action=remove

- input: null
- output: [serviceConsumeMap](#serviceconsumemap)

### Schema serviceEvent API List

#### GET /v2-beta/serviceevents

- input: null
- output: collection of [serviceEvent](#serviceevent)

#### GET /v2-beta/serviceevents/*resource-id*

- input: null
- output: [serviceEvent](#serviceevent)

#### Action create of schema object serviceEvent, POST /v2-beta/serviceevents/*resource-id*?action=create

- input: null
- output: [serviceEvent](#serviceevent)

#### Action remove of schema object serviceEvent, POST /v2-beta/serviceevents/*resource-id*?action=remove

- input: null
- output: [serviceEvent](#serviceevent)

### Schema serviceExposeMap API List

#### GET /v2-beta/serviceexposemaps

- input: null
- output: collection of [serviceExposeMap](#serviceexposemap)

#### GET /v2-beta/serviceexposemaps/*resource-id*

- input: null
- output: [serviceExposeMap](#serviceexposemap)

#### Action create of schema object serviceExposeMap, POST /v2-beta/serviceexposemaps/*resource-id*?action=create

- input: null
- output: [serviceExposeMap](#serviceexposemap)

#### Action remove of schema object serviceExposeMap, POST /v2-beta/serviceexposemaps/*resource-id*?action=remove

- input: null
- output: [serviceExposeMap](#serviceexposemap)

### Schema serviceLog API List

#### GET /v2-beta/servicelogs

- input: null
- output: collection of [serviceLog](#servicelog)

#### GET /v2-beta/servicelogs/*resource-id*

- input: null
- output: [serviceLog](#servicelog)

### Schema serviceProxy API List

#### GET /v2-beta/serviceproxies

- input: null
- output: collection of [serviceProxy](#serviceproxy)

#### GET /v2-beta/serviceproxies/*resource-id*

- input: null
- output: [serviceProxy](#serviceproxy)

### Schema setting API List

#### GET /v2-beta/settings

- input: null
- output: collection of [setting](#setting)

#### POST /v2-beta/settings

- input: [setting](#setting)
- output: [setting](#setting)

#### GET /v2-beta/settings/*resource-id*

- input: null
- output: [setting](#setting)

#### PUT /v2-beta/settings/*resource-id*

- input: [setting](#setting)
- output: [setting](#setting)

#### DELETE /v2-beta/settings/*resource-id*

- input: null
- output: [setting](#setting)

### Schema snapshot API List

#### GET /v2-beta/snapshots

- input: null
- output: collection of [snapshot](#snapshot)

#### GET /v2-beta/snapshots/*resource-id*

- input: null
- output: [snapshot](#snapshot)

#### Action backup of schema object snapshot, POST /v2-beta/snapshots/*resource-id*?action=backup

- input: [snapshotBackupInput](#snapshotbackupinput)
- output: [backup](#backup)

#### Action create of schema object snapshot, POST /v2-beta/snapshots/*resource-id*?action=create

- input: null
- output: [snapshot](#snapshot)

#### Action remove of schema object snapshot, POST /v2-beta/snapshots/*resource-id*?action=remove

- input: null
- output: [snapshot](#snapshot)

### Schema stack API List

#### GET /v2-beta/stacks

- input: null
- output: collection of [stack](#stack)

#### GET /v2-beta/stacks/*resource-id*

- input: null
- output: [stack](#stack)

#### Action activateservices of schema object stack, POST /v2-beta/stacks/*resource-id*?action=activateservices

- input: null
- output: [stack](#stack)

#### Action upgrade of schema object stack, POST /v2-beta/stacks/*resource-id*?action=upgrade

- input: [stackUpgrade](#stackupgrade)
- output: [stack](#stack)

#### Action error of schema object stack, POST /v2-beta/stacks/*resource-id*?action=error

- input: null
- output: [stack](#stack)

#### Action remove of schema object stack, POST /v2-beta/stacks/*resource-id*?action=remove

- input: null
- output: [stack](#stack)

#### Action cancelupgrade of schema object stack, POST /v2-beta/stacks/*resource-id*?action=cancelupgrade

- input: null
- output: [stack](#stack)

#### Action finishupgrade of schema object stack, POST /v2-beta/stacks/*resource-id*?action=finishupgrade

- input: null
- output: [stack](#stack)

#### Action rollback of schema object stack, POST /v2-beta/stacks/*resource-id*?action=rollback

- input: null
- output: [stack](#stack)

#### Action deactivateservices of schema object stack, POST /v2-beta/stacks/*resource-id*?action=deactivateservices

- input: null
- output: [stack](#stack)

#### Action addoutputs of schema object stack, POST /v2-beta/stacks/*resource-id*?action=addoutputs

- input: [addOutputsInput](#addoutputsinput)
- output: [stack](#stack)

#### Action create of schema object stack, POST /v2-beta/stacks/*resource-id*?action=create

- input: null
- output: [stack](#stack)

#### Action update of schema object stack, POST /v2-beta/stacks/*resource-id*?action=update

- input: null
- output: [stack](#stack)

#### Action exportconfig of schema object stack, POST /v2-beta/stacks/*resource-id*?action=exportconfig

- input: [composeConfigInput](#composeconfiginput)
- output: [composeConfig](#composeconfig)

### Schema storageDriver API List

#### GET /v2-beta/storagedrivers

- input: null
- output: collection of [storageDriver](#storagedriver)

#### GET /v2-beta/storagedrivers/*resource-id*

- input: null
- output: [storageDriver](#storagedriver)

#### Action deactivate of schema object storageDriver, POST /v2-beta/storagedrivers/*resource-id*?action=deactivate

- input: null
- output: [storageDriver](#storagedriver)

#### Action activate of schema object storageDriver, POST /v2-beta/storagedrivers/*resource-id*?action=activate

- input: null
- output: [storageDriver](#storagedriver)

#### Action create of schema object storageDriver, POST /v2-beta/storagedrivers/*resource-id*?action=create

- input: null
- output: [storageDriver](#storagedriver)

#### Action update of schema object storageDriver, POST /v2-beta/storagedrivers/*resource-id*?action=update

- input: null
- output: [storageDriver](#storagedriver)

#### Action remove of schema object storageDriver, POST /v2-beta/storagedrivers/*resource-id*?action=remove

- input: null
- output: [storageDriver](#storagedriver)

### Schema storageDriverService API List

#### GET /v2-beta/storagedriverservices

- input: null
- output: collection of [storageDriverService](#storagedriverservice)

#### GET /v2-beta/storagedriverservices/*resource-id*

- input: null
- output: [storageDriverService](#storagedriverservice)

#### Action continueupgrade of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=continueupgrade

- input: null
- output: [service](#service)

#### Action setservicelinks of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=setservicelinks

- input: [setServiceLinksInput](#setservicelinksinput)
- output: [service](#service)

#### Action addservicelink of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=addservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action remove of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=remove

- input: null
- output: [service](#service)

#### Action deactivate of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=deactivate

- input: null
- output: [service](#service)

#### Action finishupgrade of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=finishupgrade

- input: null
- output: [service](#service)

#### Action removeservicelink of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=removeservicelink

- input: [addRemoveServiceLinkInput](#addremoveservicelinkinput)
- output: [service](#service)

#### Action rollback of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=rollback

- input: null
- output: [service](#service)

#### Action restart of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=restart

- input: [serviceRestart](#servicerestart)
- output: [service](#service)

#### Action update of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=update

- input: null
- output: [service](#service)

#### Action cancelupgrade of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=cancelupgrade

- input: null
- output: [service](#service)

#### Action upgrade of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=upgrade

- input: [serviceUpgrade](#serviceupgrade)
- output: [service](#service)

#### Action activate of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=activate

- input: null
- output: [service](#service)

#### Action create of schema object storageDriverService, POST /v2-beta/storagedriverservices/*resource-id*?action=create

- input: null
- output: [service](#service)

### Schema storagePool API List

#### GET /v2-beta/storagepools

- input: null
- output: collection of [storagePool](#storagepool)

#### GET /v2-beta/storagepools/*resource-id*

- input: null
- output: [storagePool](#storagepool)

#### Action activate of schema object storagePool, POST /v2-beta/storagepools/*resource-id*?action=activate

- input: null
- output: [storagePool](#storagepool)

#### Action create of schema object storagePool, POST /v2-beta/storagepools/*resource-id*?action=create

- input: null
- output: [storagePool](#storagepool)

#### Action update of schema object storagePool, POST /v2-beta/storagepools/*resource-id*?action=update

- input: null
- output: [storagePool](#storagepool)

#### Action purge of schema object storagePool, POST /v2-beta/storagepools/*resource-id*?action=purge

- input: null
- output: [storagePool](#storagepool)

#### Action remove of schema object storagePool, POST /v2-beta/storagepools/*resource-id*?action=remove

- input: null
- output: [storagePool](#storagepool)

#### Action deactivate of schema object storagePool, POST /v2-beta/storagepools/*resource-id*?action=deactivate

- input: null
- output: [storagePool](#storagepool)

### Schema subnet API List

#### GET /v2-beta/subnets

- input: null
- output: collection of [subnet](#subnet)

#### GET /v2-beta/subnets/*resource-id*

- input: null
- output: [subnet](#subnet)

#### Action purge of schema object subnet, POST /v2-beta/subnets/*resource-id*?action=purge

- input: null
- output: [subnet](#subnet)

#### Action remove of schema object subnet, POST /v2-beta/subnets/*resource-id*?action=remove

- input: null
- output: [subnet](#subnet)

#### Action deactivate of schema object subnet, POST /v2-beta/subnets/*resource-id*?action=deactivate

- input: null
- output: [subnet](#subnet)

#### Action activate of schema object subnet, POST /v2-beta/subnets/*resource-id*?action=activate

- input: null
- output: [subnet](#subnet)

#### Action create of schema object subnet, POST /v2-beta/subnets/*resource-id*?action=create

- input: null
- output: [subnet](#subnet)

#### Action update of schema object subnet, POST /v2-beta/subnets/*resource-id*?action=update

- input: null
- output: [subnet](#subnet)

### Schema task API List

#### GET /v2-beta/tasks

- input: null
- output: collection of [task](#task)

#### GET /v2-beta/tasks/*resource-id*

- input: null
- output: [task](#task)

#### Action execute of schema object task, POST /v2-beta/tasks/*resource-id*?action=execute

- input: null
- output: [task](#task)

### Schema taskInstance API List

#### GET /v2-beta/taskinstances

- input: null
- output: collection of [taskInstance](#taskinstance)

#### GET /v2-beta/taskinstances/*resource-id*

- input: null
- output: [taskInstance](#taskinstance)

### Schema typeDocumentation API List

#### GET /v2-beta/typedocumentations

- input: null
- output: collection of [typeDocumentation](#typedocumentation)

#### GET /v2-beta/typedocumentations/*resource-id*

- input: null
- output: [typeDocumentation](#typedocumentation)

### Schema userPreference API List

#### GET /v2-beta/userpreferences

- input: null
- output: collection of [userPreference](#userpreference)

#### POST /v2-beta/userpreferences

- input: [userPreference](#userpreference)
- output: [userPreference](#userpreference)

#### GET /v2-beta/userpreferences/*resource-id*

- input: null
- output: [userPreference](#userpreference)

#### PUT /v2-beta/userpreferences/*resource-id*

- input: [userPreference](#userpreference)
- output: [userPreference](#userpreference)

#### DELETE /v2-beta/userpreferences/*resource-id*

- input: null
- output: [userPreference](#userpreference)

#### Action update of schema object userPreference, POST /v2-beta/userpreferences/*resource-id*?action=update

- input: null
- output: [userPreference](#userpreference)

#### Action purge of schema object userPreference, POST /v2-beta/userpreferences/*resource-id*?action=purge

- input: null
- output: [userPreference](#userpreference)

#### Action remove of schema object userPreference, POST /v2-beta/userpreferences/*resource-id*?action=remove

- input: null
- output: [userPreference](#userpreference)

#### Action deactivate of schema object userPreference, POST /v2-beta/userpreferences/*resource-id*?action=deactivate

- input: null
- output: [userPreference](#userpreference)

#### Action activate of schema object userPreference, POST /v2-beta/userpreferences/*resource-id*?action=activate

- input: null
- output: [userPreference](#userpreference)

#### Action create of schema object userPreference, POST /v2-beta/userpreferences/*resource-id*?action=create

- input: null
- output: [userPreference](#userpreference)

### Schema virtualMachine API List

#### GET /v2-beta/virtualmachines

- input: null
- output: collection of [virtualMachine](#virtualmachine)

#### GET /v2-beta/virtualmachines/*resource-id*

- input: null
- output: [virtualMachine](#virtualmachine)

#### Action restart of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=restart

- input: null
- output: [instance](#instance)

#### Action purge of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=purge

- input: null
- output: [instance](#instance)

#### Action create of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=create

- input: null
- output: [instance](#instance)

#### Action migrate of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=migrate

- input: null
- output: [instance](#instance)

#### Action proxy of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=proxy

- input: [containerProxy](#containerproxy)
- output: [hostAccess](#hostaccess)

#### Action console of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=console

- input: [instanceConsoleInput](#instanceconsoleinput)
- output: [instanceConsole](#instanceconsole)

#### Action deallocate of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=deallocate

- input: null
- output: [instance](#instance)

#### Action start of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=start

- input: null
- output: [instance](#instance)

#### Action updatehealthy of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=updatehealthy

- input: null
- output: [instance](#instance)

#### Action logs of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=logs

- input: [containerLogs](#containerlogs)
- output: [hostAccess](#hostaccess)

#### Action execute of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=execute

- input: [containerExec](#containerexec)
- output: [hostAccess](#hostaccess)

#### Action updatereinitializing of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=updatereinitializing

- input: null
- output: [instance](#instance)

#### Action update of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=update

- input: null
- output: [instance](#instance)

#### Action remove of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=remove

- input: null
- output: [instance](#instance)

#### Action allocate of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=allocate

- input: null
- output: [instance](#instance)

#### Action updateunhealthy of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=updateunhealthy

- input: null
- output: [instance](#instance)

#### Action error of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=error

- input: null
- output: [instance](#instance)

#### Action stop of schema object virtualMachine, POST /v2-beta/virtualmachines/*resource-id*?action=stop

- input: [instanceStop](#instancestop)
- output: [instance](#instance)

### Schema volume API List

#### GET /v2-beta/volumes

- input: null
- output: collection of [volume](#volume)

#### GET /v2-beta/volumes/*resource-id*

- input: null
- output: [volume](#volume)

#### Action allocate of schema object volume, POST /v2-beta/volumes/*resource-id*?action=allocate

- input: null
- output: [volume](#volume)

#### Action deallocate of schema object volume, POST /v2-beta/volumes/*resource-id*?action=deallocate

- input: null
- output: [volume](#volume)

#### Action create of schema object volume, POST /v2-beta/volumes/*resource-id*?action=create

- input: null
- output: [volume](#volume)

#### Action update of schema object volume, POST /v2-beta/volumes/*resource-id*?action=update

- input: null
- output: [volume](#volume)

#### Action snapshot of schema object volume, POST /v2-beta/volumes/*resource-id*?action=snapshot

- input: [volumeSnapshotInput](#volumesnapshotinput)
- output: [snapshot](#snapshot)

#### Action restorefrombackup of schema object volume, POST /v2-beta/volumes/*resource-id*?action=restorefrombackup

- input: [restoreFromBackupInput](#restorefrombackupinput)
- output: [volume](#volume)

#### Action purge of schema object volume, POST /v2-beta/volumes/*resource-id*?action=purge

- input: null
- output: [volume](#volume)

#### Action remove of schema object volume, POST /v2-beta/volumes/*resource-id*?action=remove

- input: null
- output: [volume](#volume)

#### Action reverttosnapshot of schema object volume, POST /v2-beta/volumes/*resource-id*?action=reverttosnapshot

- input: [revertToSnapshotInput](#reverttosnapshotinput)
- output: [volume](#volume)

### Schema volumeTemplate API List

#### GET /v2-beta/volumetemplates

- input: null
- output: collection of [volumeTemplate](#volumetemplate)

#### GET /v2-beta/volumetemplates/*resource-id*

- input: null
- output: [volumeTemplate](#volumetemplate)

#### Action activate of schema object volumeTemplate, POST /v2-beta/volumetemplates/*resource-id*?action=activate

- input: null
- output: [volumeTemplate](#volumetemplate)

#### Action create of schema object volumeTemplate, POST /v2-beta/volumetemplates/*resource-id*?action=create

- input: null
- output: [volumeTemplate](#volumetemplate)

#### Action update of schema object volumeTemplate, POST /v2-beta/volumetemplates/*resource-id*?action=update

- input: null
- output: [volumeTemplate](#volumetemplate)

#### Action purge of schema object volumeTemplate, POST /v2-beta/volumetemplates/*resource-id*?action=purge

- input: null
- output: [volumeTemplate](#volumetemplate)

#### Action remove of schema object volumeTemplate, POST /v2-beta/volumetemplates/*resource-id*?action=remove

- input: null
- output: [volumeTemplate](#volumetemplate)

#### Action deactivate of schema object volumeTemplate, POST /v2-beta/volumetemplates/*resource-id*?action=deactivate

- input: null
- output: [volumeTemplate](#volumetemplate)

