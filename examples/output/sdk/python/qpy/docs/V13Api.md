# qpy.V13Api

All URIs are relative to *https://api.qubole.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account**](V13Api.md#create_account) | **POST** /v1.3/accounts/ | 
[**create_account_0**](V13Api.md#create_account_0) | **PATCH** /v1.3/accounts/users | 
[**create_account_1**](V13Api.md#create_account_1) | **PUT** /v1.3/accounts/users/123/customers/12 | 
[**create_cluster**](V13Api.md#create_cluster) | **POST** /v1.3/clusters | 


# **create_account**
> CreateAccountResponse create_account(x_auth_token, user_agent, request_body=request_body)



Create Account

### Example

```python
from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = qpy.V13Api()
x_auth_token = 'x_auth_token_example' # str | X-AUTH-TOKEN
user_agent = 'user_agent_example' # str | User-Agent
request_body = {"account":{"CacheQuotaSizeInGB":50,"aws_region":"us-east-1","compute_access_key":"$COMPUTE_ACCESS_KEY","compute_secret_key":"$COMPUTE_SECRET_KEY","compute_type":"CUSTOMER_MANAGED","defloc":"$DEFLOC","idle_cluster_timeout":3,"idle_session_timeout":2880,"level":"free","name":"Test","storage_access_key":"$STORAGE_ACCESS_KEY","storage_secret_key":"$STORAGE_SECRET_KEY","storage_type":"CUSTOMER_MANAGED","sub_account_creation":false}} # dict(str, object) |  (optional)

try:
    api_response = api_instance.create_account(x_auth_token, user_agent, request_body=request_body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V13Api->create_account: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x_auth_token** | **str**| X-AUTH-TOKEN | 
 **user_agent** | **str**| User-Agent | 
 **request_body** | [**dict(str, object)**](object.md)|  | [optional] 

### Return type

[**CreateAccountResponse**](CreateAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_account_0**
> CreateAccountResponse2 create_account_0(x_auth_token, user_agent, request_body=request_body)



Create Account

### Example

```python
from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = qpy.V13Api()
x_auth_token = 'x_auth_token_example' # str | X-AUTH-TOKEN
user_agent = 'user_agent_example' # str | User-Agent
request_body = {"account":{"CacheQuotaSizeInGB":50,"aws_region":"us-east-1","compute_access_key":"$COMPUTE_ACCESS_KEY","compute_secret_key":"$COMPUTE_SECRET_KEY","compute_type":"CUSTOMER_MANAGED","defloc":"$DEFLOC","idle_cluster_timeout":3,"idle_session_timeout":2880,"level":"free","name":"Test","storage_access_key":"$STORAGE_ACCESS_KEY","storage_secret_key":"$STORAGE_SECRET_KEY","storage_type":"CUSTOMER_MANAGED","sub_account_creation":false}} # dict(str, object) |  (optional)

try:
    api_response = api_instance.create_account_0(x_auth_token, user_agent, request_body=request_body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V13Api->create_account_0: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x_auth_token** | **str**| X-AUTH-TOKEN | 
 **user_agent** | **str**| User-Agent | 
 **request_body** | [**dict(str, object)**](object.md)|  | [optional] 

### Return type

[**CreateAccountResponse2**](CreateAccountResponse2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_account_1**
> CreateAccountResponse2 create_account_1(x_auth_token, user_agent, request_body=request_body)



Create Account

### Example

```python
from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = qpy.V13Api()
x_auth_token = 'x_auth_token_example' # str | X-AUTH-TOKEN
user_agent = 'user_agent_example' # str | User-Agent
request_body = {"account":{"CacheQuotaSizeInGB":50,"aws_region":"us-east-1","compute_access_key":"$COMPUTE_ACCESS_KEY","compute_secret_key":"$COMPUTE_SECRET_KEY","compute_type":"CUSTOMER_MANAGED","defloc":"$DEFLOC","idle_cluster_timeout":3,"idle_session_timeout":2880,"level":"free","name":"Test","storage_access_key":"$STORAGE_ACCESS_KEY","storage_secret_key":"$STORAGE_SECRET_KEY","storage_type":"CUSTOMER_MANAGED","sub_account_creation":false}} # dict(str, object) |  (optional)

try:
    api_response = api_instance.create_account_1(x_auth_token, user_agent, request_body=request_body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V13Api->create_account_1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x_auth_token** | **str**| X-AUTH-TOKEN | 
 **user_agent** | **str**| User-Agent | 
 **request_body** | [**dict(str, object)**](object.md)|  | [optional] 

### Return type

[**CreateAccountResponse2**](CreateAccountResponse2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_cluster**
> CreateClusterResponse create_cluster(x_auth_token, user_agent, request_body=request_body)



Create Cluster

### Example

```python
from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = qpy.V13Api()
x_auth_token = 'x_auth_token_example' # str | X-AUTH-TOKEN
user_agent = 'user_agent_example' # str | User-Agent
request_body = {"datadog_settings":{"datadog_api_token":"","datadog_app_token":""},"disallow_cluster_termination":false,"ec2_settings":{"aws_preferred_availability_zone":"Any","aws_region":"us-west-2","bastion_node_port":"","bastion_node_public_dns":"","bastion_node_user":"ec2-user","compute_access_key":"\u003cyour_ec2_compute_access_key\u003e","compute_secret_key":"\u003cyour_ec2_compute_secret_key\u003e","compute_validated":true,"master_elastic_ip":"","role_instance_profile":"","subnet_id":"","use_account_compute_creds":true,"vpc_id":""},"enable_ganglia_monitoring":false,"engine_config":{"dbtap_id":"9670","fernet_key":"\u003cyour-fernet-key\u003e","hive_settings":{"hive.qubole.metadata.cache":"true","hive_version":"2.1.1","hs2_thrift_port":"10007","is_hs2":true,"overrides":"hive.execution.engine=tez"},"overrides":"core.dag_concurrency=32","type":"airflow"},"hadoop_settings":{"custom_config":{},"fairscheduler_settings":{"default_pool":"","fairscheduler_config_xml":""},"max_nodes":10,"use_hadoop2":false,"use_qubole_placement_policy":"","use_spark":true},"label":["my_cluster"],"node_bootstrap_file":"node_bootstrap.sh","node_configuration":{"custom_ec2_tags":"{\"project\": \"webportal\", \"owner\": \"john@example.com\"}","ebs_upscaling_config":{"absolute_free_space_threshold":100,"max_ebs_volume_count":5,"percent_free_space_threshold":20,"sampling_interval":40,"sampling_window":8},"ebs_volume_count":0,"ebs_volume_size":"100GB","ebs_volume_type":"standard","fallback_to_ondemand":"spot","heterogeneous_instance_config":{"memory":[{"instance_type":"m4.4xlarge","weight":1},{"instance_type":"m4.2xlarge","weight":0.5},{"instance_type":"m4.xlarge","weight":0.25}]},"idle_cluster_timeout":0,"idle_cluster_timeout_in_secs":120,"initial_nodes":1,"master_instance_type":"m1.large","max_nodes":10,"node_base_cooldown_period":5,"node_spot_cooldown_period":10,"slave_instance_type":"m1.xlarge","slave_request_type":"spot","spot_block_settings":{"duration":120},"spot_instance_settings":{"maximum_bid_price_percentage":100,"maximum_spot_instance_percentage":60,"timeout_for_request":10},"stable_spot_instance_settings":{"maximum_bid_price_percentage":100,"timeout_for_request":10}},"presto_settings":{"custom_config":{},"enable_presto":false},"presto_version":"0.193","security_settings":{"encrypted_ephemerals":false,"persistent_security_groups":"","ssh_public_key":""},"spark_settings":{"custom_config":{}},"spark_version":"2.3","state":"DOWN","use_hadoop2":false,"zeppelin_interpreter_mode":"legacy"} # dict(str, object) |  (optional)

try:
    api_response = api_instance.create_cluster(x_auth_token, user_agent, request_body=request_body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V13Api->create_cluster: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x_auth_token** | **str**| X-AUTH-TOKEN | 
 **user_agent** | **str**| User-Agent | 
 **request_body** | [**dict(str, object)**](object.md)|  | [optional] 

### Return type

[**CreateClusterResponse**](CreateClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

