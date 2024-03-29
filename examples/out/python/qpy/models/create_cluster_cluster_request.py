# coding: utf-8

"""
    Qubole API Sepcification

    Operations, requests and responses  # noqa: E501

    OpenAPI spec version: v1.11.0
    Contact: support@qubole.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six


class CreateClusterClusterRequest(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'datadog_settings': 'CreateClusterClusterRequestDatadogSettings',
        'disallow_cluster_termination': 'bool',
        'ec2_settings': 'CreateClusterClusterRequestEc2Settings',
        'enable_ganglia_monitoring': 'bool',
        'engine_config': 'CreateClusterClusterRequestEngineConfig',
        'hadoop_settings': 'CreateClusterClusterRequestHadoopSettings',
        'label': 'list[str]',
        'node_bootstrap_file': 'str',
        'node_configuration': 'CreateClusterClusterRequestNodeConfiguration',
        'presto_settings': 'CreateClusterClusterRequestPrestoSettings',
        'presto_version': 'str',
        'security_settings': 'CreateClusterClusterRequestSecuritySettings',
        'spark_settings': 'CreateClusterClusterRequestSparkSettings',
        'spark_version': 'str',
        'state': 'str',
        'use_hadoop2': 'bool',
        'zeppelin_interpreter_mode': 'str'
    }

    attribute_map = {
        'datadog_settings': 'datadog_settings',
        'disallow_cluster_termination': 'disallow_cluster_termination',
        'ec2_settings': 'ec2_settings',
        'enable_ganglia_monitoring': 'enable_ganglia_monitoring',
        'engine_config': 'engine_config',
        'hadoop_settings': 'hadoop_settings',
        'label': 'label',
        'node_bootstrap_file': 'node_bootstrap_file',
        'node_configuration': 'node_configuration',
        'presto_settings': 'presto_settings',
        'presto_version': 'presto_version',
        'security_settings': 'security_settings',
        'spark_settings': 'spark_settings',
        'spark_version': 'spark_version',
        'state': 'state',
        'use_hadoop2': 'use_hadoop2',
        'zeppelin_interpreter_mode': 'zeppelin_interpreter_mode'
    }

    def __init__(self, datadog_settings=None, disallow_cluster_termination=None, ec2_settings=None, enable_ganglia_monitoring=None, engine_config=None, hadoop_settings=None, label=None, node_bootstrap_file=None, node_configuration=None, presto_settings=None, presto_version=None, security_settings=None, spark_settings=None, spark_version=None, state=None, use_hadoop2=None, zeppelin_interpreter_mode=None):  # noqa: E501
        """CreateClusterClusterRequest - a model defined in OpenAPI"""  # noqa: E501

        self._datadog_settings = None
        self._disallow_cluster_termination = None
        self._ec2_settings = None
        self._enable_ganglia_monitoring = None
        self._engine_config = None
        self._hadoop_settings = None
        self._label = None
        self._node_bootstrap_file = None
        self._node_configuration = None
        self._presto_settings = None
        self._presto_version = None
        self._security_settings = None
        self._spark_settings = None
        self._spark_version = None
        self._state = None
        self._use_hadoop2 = None
        self._zeppelin_interpreter_mode = None
        self.discriminator = None

        if datadog_settings is not None:
            self.datadog_settings = datadog_settings
        if disallow_cluster_termination is not None:
            self.disallow_cluster_termination = disallow_cluster_termination
        if ec2_settings is not None:
            self.ec2_settings = ec2_settings
        if enable_ganglia_monitoring is not None:
            self.enable_ganglia_monitoring = enable_ganglia_monitoring
        if engine_config is not None:
            self.engine_config = engine_config
        if hadoop_settings is not None:
            self.hadoop_settings = hadoop_settings
        if label is not None:
            self.label = label
        if node_bootstrap_file is not None:
            self.node_bootstrap_file = node_bootstrap_file
        if node_configuration is not None:
            self.node_configuration = node_configuration
        if presto_settings is not None:
            self.presto_settings = presto_settings
        if presto_version is not None:
            self.presto_version = presto_version
        if security_settings is not None:
            self.security_settings = security_settings
        if spark_settings is not None:
            self.spark_settings = spark_settings
        if spark_version is not None:
            self.spark_version = spark_version
        if state is not None:
            self.state = state
        if use_hadoop2 is not None:
            self.use_hadoop2 = use_hadoop2
        if zeppelin_interpreter_mode is not None:
            self.zeppelin_interpreter_mode = zeppelin_interpreter_mode

    @property
    def datadog_settings(self):
        """Gets the datadog_settings of this CreateClusterClusterRequest.  # noqa: E501


        :return: The datadog_settings of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestDatadogSettings
        """
        return self._datadog_settings

    @datadog_settings.setter
    def datadog_settings(self, datadog_settings):
        """Sets the datadog_settings of this CreateClusterClusterRequest.


        :param datadog_settings: The datadog_settings of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestDatadogSettings
        """

        self._datadog_settings = datadog_settings

    @property
    def disallow_cluster_termination(self):
        """Gets the disallow_cluster_termination of this CreateClusterClusterRequest.  # noqa: E501


        :return: The disallow_cluster_termination of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: bool
        """
        return self._disallow_cluster_termination

    @disallow_cluster_termination.setter
    def disallow_cluster_termination(self, disallow_cluster_termination):
        """Sets the disallow_cluster_termination of this CreateClusterClusterRequest.


        :param disallow_cluster_termination: The disallow_cluster_termination of this CreateClusterClusterRequest.  # noqa: E501
        :type: bool
        """

        self._disallow_cluster_termination = disallow_cluster_termination

    @property
    def ec2_settings(self):
        """Gets the ec2_settings of this CreateClusterClusterRequest.  # noqa: E501


        :return: The ec2_settings of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestEc2Settings
        """
        return self._ec2_settings

    @ec2_settings.setter
    def ec2_settings(self, ec2_settings):
        """Sets the ec2_settings of this CreateClusterClusterRequest.


        :param ec2_settings: The ec2_settings of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestEc2Settings
        """

        self._ec2_settings = ec2_settings

    @property
    def enable_ganglia_monitoring(self):
        """Gets the enable_ganglia_monitoring of this CreateClusterClusterRequest.  # noqa: E501


        :return: The enable_ganglia_monitoring of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: bool
        """
        return self._enable_ganglia_monitoring

    @enable_ganglia_monitoring.setter
    def enable_ganglia_monitoring(self, enable_ganglia_monitoring):
        """Sets the enable_ganglia_monitoring of this CreateClusterClusterRequest.


        :param enable_ganglia_monitoring: The enable_ganglia_monitoring of this CreateClusterClusterRequest.  # noqa: E501
        :type: bool
        """

        self._enable_ganglia_monitoring = enable_ganglia_monitoring

    @property
    def engine_config(self):
        """Gets the engine_config of this CreateClusterClusterRequest.  # noqa: E501


        :return: The engine_config of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestEngineConfig
        """
        return self._engine_config

    @engine_config.setter
    def engine_config(self, engine_config):
        """Sets the engine_config of this CreateClusterClusterRequest.


        :param engine_config: The engine_config of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestEngineConfig
        """

        self._engine_config = engine_config

    @property
    def hadoop_settings(self):
        """Gets the hadoop_settings of this CreateClusterClusterRequest.  # noqa: E501


        :return: The hadoop_settings of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestHadoopSettings
        """
        return self._hadoop_settings

    @hadoop_settings.setter
    def hadoop_settings(self, hadoop_settings):
        """Sets the hadoop_settings of this CreateClusterClusterRequest.


        :param hadoop_settings: The hadoop_settings of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestHadoopSettings
        """

        self._hadoop_settings = hadoop_settings

    @property
    def label(self):
        """Gets the label of this CreateClusterClusterRequest.  # noqa: E501


        :return: The label of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: list[str]
        """
        return self._label

    @label.setter
    def label(self, label):
        """Sets the label of this CreateClusterClusterRequest.


        :param label: The label of this CreateClusterClusterRequest.  # noqa: E501
        :type: list[str]
        """

        self._label = label

    @property
    def node_bootstrap_file(self):
        """Gets the node_bootstrap_file of this CreateClusterClusterRequest.  # noqa: E501


        :return: The node_bootstrap_file of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: str
        """
        return self._node_bootstrap_file

    @node_bootstrap_file.setter
    def node_bootstrap_file(self, node_bootstrap_file):
        """Sets the node_bootstrap_file of this CreateClusterClusterRequest.


        :param node_bootstrap_file: The node_bootstrap_file of this CreateClusterClusterRequest.  # noqa: E501
        :type: str
        """

        self._node_bootstrap_file = node_bootstrap_file

    @property
    def node_configuration(self):
        """Gets the node_configuration of this CreateClusterClusterRequest.  # noqa: E501


        :return: The node_configuration of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestNodeConfiguration
        """
        return self._node_configuration

    @node_configuration.setter
    def node_configuration(self, node_configuration):
        """Sets the node_configuration of this CreateClusterClusterRequest.


        :param node_configuration: The node_configuration of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestNodeConfiguration
        """

        self._node_configuration = node_configuration

    @property
    def presto_settings(self):
        """Gets the presto_settings of this CreateClusterClusterRequest.  # noqa: E501


        :return: The presto_settings of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestPrestoSettings
        """
        return self._presto_settings

    @presto_settings.setter
    def presto_settings(self, presto_settings):
        """Sets the presto_settings of this CreateClusterClusterRequest.


        :param presto_settings: The presto_settings of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestPrestoSettings
        """

        self._presto_settings = presto_settings

    @property
    def presto_version(self):
        """Gets the presto_version of this CreateClusterClusterRequest.  # noqa: E501


        :return: The presto_version of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: str
        """
        return self._presto_version

    @presto_version.setter
    def presto_version(self, presto_version):
        """Sets the presto_version of this CreateClusterClusterRequest.


        :param presto_version: The presto_version of this CreateClusterClusterRequest.  # noqa: E501
        :type: str
        """

        self._presto_version = presto_version

    @property
    def security_settings(self):
        """Gets the security_settings of this CreateClusterClusterRequest.  # noqa: E501


        :return: The security_settings of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestSecuritySettings
        """
        return self._security_settings

    @security_settings.setter
    def security_settings(self, security_settings):
        """Sets the security_settings of this CreateClusterClusterRequest.


        :param security_settings: The security_settings of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestSecuritySettings
        """

        self._security_settings = security_settings

    @property
    def spark_settings(self):
        """Gets the spark_settings of this CreateClusterClusterRequest.  # noqa: E501


        :return: The spark_settings of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: CreateClusterClusterRequestSparkSettings
        """
        return self._spark_settings

    @spark_settings.setter
    def spark_settings(self, spark_settings):
        """Sets the spark_settings of this CreateClusterClusterRequest.


        :param spark_settings: The spark_settings of this CreateClusterClusterRequest.  # noqa: E501
        :type: CreateClusterClusterRequestSparkSettings
        """

        self._spark_settings = spark_settings

    @property
    def spark_version(self):
        """Gets the spark_version of this CreateClusterClusterRequest.  # noqa: E501


        :return: The spark_version of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: str
        """
        return self._spark_version

    @spark_version.setter
    def spark_version(self, spark_version):
        """Sets the spark_version of this CreateClusterClusterRequest.


        :param spark_version: The spark_version of this CreateClusterClusterRequest.  # noqa: E501
        :type: str
        """

        self._spark_version = spark_version

    @property
    def state(self):
        """Gets the state of this CreateClusterClusterRequest.  # noqa: E501


        :return: The state of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: str
        """
        return self._state

    @state.setter
    def state(self, state):
        """Sets the state of this CreateClusterClusterRequest.


        :param state: The state of this CreateClusterClusterRequest.  # noqa: E501
        :type: str
        """

        self._state = state

    @property
    def use_hadoop2(self):
        """Gets the use_hadoop2 of this CreateClusterClusterRequest.  # noqa: E501


        :return: The use_hadoop2 of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: bool
        """
        return self._use_hadoop2

    @use_hadoop2.setter
    def use_hadoop2(self, use_hadoop2):
        """Sets the use_hadoop2 of this CreateClusterClusterRequest.


        :param use_hadoop2: The use_hadoop2 of this CreateClusterClusterRequest.  # noqa: E501
        :type: bool
        """

        self._use_hadoop2 = use_hadoop2

    @property
    def zeppelin_interpreter_mode(self):
        """Gets the zeppelin_interpreter_mode of this CreateClusterClusterRequest.  # noqa: E501


        :return: The zeppelin_interpreter_mode of this CreateClusterClusterRequest.  # noqa: E501
        :rtype: str
        """
        return self._zeppelin_interpreter_mode

    @zeppelin_interpreter_mode.setter
    def zeppelin_interpreter_mode(self, zeppelin_interpreter_mode):
        """Sets the zeppelin_interpreter_mode of this CreateClusterClusterRequest.


        :param zeppelin_interpreter_mode: The zeppelin_interpreter_mode of this CreateClusterClusterRequest.  # noqa: E501
        :type: str
        """

        self._zeppelin_interpreter_mode = zeppelin_interpreter_mode

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CreateClusterClusterRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
