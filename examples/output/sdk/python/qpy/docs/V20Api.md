# qpy.V20Api

All URIs are relative to *https://api.qubole.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account**](V20Api.md#create_account) | **POST** /v2/accounts/ | 
[**get_account**](V20Api.md#get_account) | **GET** /v2/{user_id}/1 | 
[**get_account_0**](V20Api.md#get_account_0) | **GET** /v2/accounts/{user_id}/1 | 


# **create_account**
> CreateAccountResponse2 create_account(user_agent, x_auth_token, request_body=request_body)



Create Account

### Example

```python
from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = qpy.V20Api()
user_agent = 'user_agent_example' # str | User-Agent
x_auth_token = 'x_auth_token_example' # str | X-AUTH-TOKEN
request_body = {"account":{"CacheQuotaSizeInGB":50,"aws_region":"us-east-1","compute_access_key":"$COMPUTE_ACCESS_KEY","compute_secret_key":"$COMPUTE_SECRET_KEY","compute_type":"CUSTOMER_MANAGED","defloc":"$DEFLOC","idle_cluster_timeout":3,"idle_session_timeout":2880,"level":"free","name":"Test","storage_access_key":"$STORAGE_ACCESS_KEY","storage_secret_key":"$STORAGE_SECRET_KEY","storage_type":"CUSTOMER_MANAGED","sub_account_creation":false}} # dict(str, object) |  (optional)

try:
    api_response = api_instance.create_account(user_agent, x_auth_token, request_body=request_body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V20Api->create_account: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_agent** | **str**| User-Agent | 
 **x_auth_token** | **str**| X-AUTH-TOKEN | 
 **request_body** | [**dict(str, object)**](object.md)|  | [optional] 

### Return type

[**CreateAccountResponse2**](CreateAccountResponse2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account**
> GetAccountResponse get_account(x_auth_token, user_agent, user_id, id=id)



Get an Account

### Example

```python
from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = qpy.V20Api()
x_auth_token = 'x_auth_token_example' # str | X-AUTH-TOKEN
user_agent = 'user_agent_example' # str | User-Agent
user_id = 'user_id_example' # str | user_id
id = 3.4 # float | id (optional)

try:
    api_response = api_instance.get_account(x_auth_token, user_agent, user_id, id=id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V20Api->get_account: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x_auth_token** | **str**| X-AUTH-TOKEN | 
 **user_agent** | **str**| User-Agent | 
 **user_id** | **str**| user_id | 
 **id** | **float**| id | [optional] 

### Return type

[**GetAccountResponse**](GetAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_0**
> GetAccountResponse2 get_account_0(x_auth_token, user_agent, user_id, id=id)



Get an Account

### Example

```python
from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = qpy.V20Api()
x_auth_token = 'x_auth_token_example' # str | X-AUTH-TOKEN
user_agent = 'user_agent_example' # str | User-Agent
user_id = 'user_id_example' # str | user_id
id = 3.4 # float | id (optional)

try:
    api_response = api_instance.get_account_0(x_auth_token, user_agent, user_id, id=id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V20Api->get_account_0: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x_auth_token** | **str**| X-AUTH-TOKEN | 
 **user_agent** | **str**| User-Agent | 
 **user_id** | **str**| user_id | 
 **id** | **float**| id | [optional] 

### Return type

[**GetAccountResponse2**](GetAccountResponse2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

