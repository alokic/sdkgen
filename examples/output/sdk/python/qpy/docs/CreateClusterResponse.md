# CreateClusterResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**datadog_settings** | **dict(str, object)** |  | [optional] 
**disallow_cluster_termination** | **bool** |  | [optional] 
**ec2_settings** | [**CreateClusterResponseEc2Settings**](CreateClusterResponseEc2Settings.md) |  | [optional] 
**enable_ganglia_monitoring** | **bool** |  | [optional] 
**engine_config** | [**CreateClusterResponseEngineConfig**](CreateClusterResponseEngineConfig.md) |  | [optional] 
**errors** | **list[str]** |  | [optional] 
**hadoop_settings** | [**CreateClusterResponseHadoopSettings**](CreateClusterResponseHadoopSettings.md) |  | [optional] 
**id** | **float** |  | [optional] 
**label** | **list[str]** |  | [optional] 
**node_bootstrap_file** | **str** |  | [optional] 
**node_configuration** | [**CreateClusterResponseNodeConfiguration**](CreateClusterResponseNodeConfiguration.md) |  | [optional] 
**presto_settings** | [**CreateClusterResponsePrestoSettings**](CreateClusterResponsePrestoSettings.md) |  | [optional] 
**security_settings** | [**CreateClusterResponseSecuritySettings**](CreateClusterResponseSecuritySettings.md) |  | [optional] 
**spark_s3_package_name** | **str** |  | [optional] 
**spark_settings** | **dict(str, object)** |  | [optional] 
**state** | **str** |  | [optional] 
**zeppelin_s3_package_name** | **str** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


