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


class CreateClusterClusterResponsePrestoSettings(object):
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
        'enable_presto': 'bool'
    }

    attribute_map = {
        'enable_presto': 'enable_presto'
    }

    def __init__(self, enable_presto=None):  # noqa: E501
        """CreateClusterClusterResponsePrestoSettings - a model defined in OpenAPI"""  # noqa: E501

        self._enable_presto = None
        self.discriminator = None

        if enable_presto is not None:
            self.enable_presto = enable_presto

    @property
    def enable_presto(self):
        """Gets the enable_presto of this CreateClusterClusterResponsePrestoSettings.  # noqa: E501


        :return: The enable_presto of this CreateClusterClusterResponsePrestoSettings.  # noqa: E501
        :rtype: bool
        """
        return self._enable_presto

    @enable_presto.setter
    def enable_presto(self, enable_presto):
        """Sets the enable_presto of this CreateClusterClusterResponsePrestoSettings.


        :param enable_presto: The enable_presto of this CreateClusterClusterResponsePrestoSettings.  # noqa: E501
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
        if not isinstance(other, CreateClusterClusterResponsePrestoSettings):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
