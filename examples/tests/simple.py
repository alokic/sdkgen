from __future__ import print_function
import time
import qpy
from qpy.rest import ApiException
from pprint import pprint

configuration=qpy.Configuration()
configuration.host="https://qa3.qubole.net/api"
user_agent = 'user_agent_example' # str | User-Agent
x_auth_token = 'x_auth_token' # str | X-AUTH-TOKEN


######### ACCOUNT CREATION START ############
# api_instance = qpy.V20Api(qpy.ApiClient(configuration))
# request_body = {"account":{"CacheQuotaSizeInGB":50,"aws_region":"us-east-1","compute_access_key":"$COMPUTE_ACCESS_KEY","compute_secret_key":"$COMPUTE_SECRET_KEY","compute_type":"CUSTOMER_MANAGED","defloc":"$DEFLOC","idle_cluster_timeout":3,"idle_session_timeout":2880,"level":"free","name":"Test","storage_access_key":"$STORAGE_ACCESS_KEY","storage_secret_key":"$STORAGE_SECRET_KEY","storage_type":"CUSTOMER_MANAGED","sub_account_creation":False}} # dict(str, object) |  (optional)

# try:
#     api_response = api_instance.create_account(x_auth_token, user_agent, request_body=request_body)
#     pprint(api_response)
# except ApiException as e:
#     print("Exception when calling V13Api->create_account: %s\n" % e)



######### CLUSTER CREATION START ############
request_body = {
      "label": [
      "spark-cluster-3"
  ],
  "spark_version": "2.3",
  "ec2_settings": {
      "use_account_compute_creds": True
  },
  "node_configuration": {
      "max_nodes": 2,
      "master_instance_type": "m1.large",
      "slave_instance_type": "m1.xlarge",
      "initial_nodes": 1
  },
  "hadoop_settings":{
      "use_spark": True
  },
  "spark_settings":{
      "custom_config": {}
  },
  "xyz": 2.3
}
api_instance = qpy.V13Api(qpy.ApiClient(configuration))
try:
    api_response = api_instance.create_cluster(x_auth_token, user_agent, request_body=request_body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling V13Api->create_account: %s\n" % e)