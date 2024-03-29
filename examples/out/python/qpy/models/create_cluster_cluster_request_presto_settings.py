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


class CreateClusterClusterRequestPrestoSettings(object):
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
        'custom_config': 'dict(str, object)',
        'enable_presto': 'bool'
    }

    attribute_map = {
        'custom_config': 'custom_config',
        'enable_presto': 'enable_presto'
    }

    def __init__(self, custom_config=None, enable_presto=None):  # noqa: E501
        """CreateClusterClusterRequestPrestoSettings - a model defined in OpenAPI"""  # noqa: E501

        self._custom_config = None
        self._enable_presto = None
        self.discriminator = None

        if custom_config is not None:
            self.custom_config = custom_config
        if enable_presto is not None:
            self.enable_presto = enable_presto

    @property
    def custom_config(self):
        """Gets the custom_config of this CreateClusterClusterRequestPrestoSettings.  # noqa: E501


        :return: The custom_config of this CreateClusterClusterRequestPrestoSettings.  # noqa: E501
        :rtype: dict(str, object)
        """
        return self._custom_config

    @custom_config.setter
    def custom_config(self, custom_config):
        """Sets the custom_config of this CreateClusterClusterRequestPrestoSettings.


        :param custom_config: The custom_config of this CreateClusterClusterRequestPrestoSettings.  # noqa: E501
        :type: dict(str, object)
        """

        self._custom_config = custom_config

    @property
    def enable_presto(self):
        """Gets the enable_presto of this CreateClusterClusterRequestPrestoSettings.  # noqa: E501


        :return: The enable_presto of this CreateClusterClusterRequestPrestoSettings.  # noqa: E501
        :rtype: bool
        """
        return self._enable_presto

    @enable_presto.setter
    def enable_presto(self, enable_presto):
        """Sets the enable_presto of this CreateClusterClusterRequestPrestoSettings.


        :param enable_presto: The enable_presto of this CreateClusterClusterRequestPrestoSettings.  # noqa: E501
        :type: bool
        """

        self._enable_presto = enable_presto

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
        if not isinstance(other, CreateClusterClusterRequestPrestoSettings):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
