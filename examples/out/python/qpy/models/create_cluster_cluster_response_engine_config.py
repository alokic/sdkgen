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


class CreateClusterClusterResponseEngineConfig(object):
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
        'dbtap_id': 'float',
        'fernet_key': 'str',
        'overrides': 'str',
        'type': 'str'
    }

    attribute_map = {
        'dbtap_id': 'dbtap_id',
        'fernet_key': 'fernet_key',
        'overrides': 'overrides',
        'type': 'type'
    }

    def __init__(self, dbtap_id=None, fernet_key=None, overrides=None, type=None):  # noqa: E501
        """CreateClusterClusterResponseEngineConfig - a model defined in OpenAPI"""  # noqa: E501

        self._dbtap_id = None
        self._fernet_key = None
        self._overrides = None
        self._type = None
        self.discriminator = None

        if dbtap_id is not None:
            self.dbtap_id = dbtap_id
        if fernet_key is not None:
            self.fernet_key = fernet_key
        if overrides is not None:
            self.overrides = overrides
        if type is not None:
            self.type = type

    @property
    def dbtap_id(self):
        """Gets the dbtap_id of this CreateClusterClusterResponseEngineConfig.  # noqa: E501


        :return: The dbtap_id of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :rtype: float
        """
        return self._dbtap_id

    @dbtap_id.setter
    def dbtap_id(self, dbtap_id):
        """Sets the dbtap_id of this CreateClusterClusterResponseEngineConfig.


        :param dbtap_id: The dbtap_id of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :type: float
        """

        self._dbtap_id = dbtap_id

    @property
    def fernet_key(self):
        """Gets the fernet_key of this CreateClusterClusterResponseEngineConfig.  # noqa: E501


        :return: The fernet_key of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :rtype: str
        """
        return self._fernet_key

    @fernet_key.setter
    def fernet_key(self, fernet_key):
        """Sets the fernet_key of this CreateClusterClusterResponseEngineConfig.


        :param fernet_key: The fernet_key of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :type: str
        """

        self._fernet_key = fernet_key

    @property
    def overrides(self):
        """Gets the overrides of this CreateClusterClusterResponseEngineConfig.  # noqa: E501


        :return: The overrides of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :rtype: str
        """
        return self._overrides

    @overrides.setter
    def overrides(self, overrides):
        """Sets the overrides of this CreateClusterClusterResponseEngineConfig.


        :param overrides: The overrides of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :type: str
        """

        self._overrides = overrides

    @property
    def type(self):
        """Gets the type of this CreateClusterClusterResponseEngineConfig.  # noqa: E501


        :return: The type of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :rtype: str
        """
        return self._type

    @type.setter
    def type(self, type):
        """Sets the type of this CreateClusterClusterResponseEngineConfig.


        :param type: The type of this CreateClusterClusterResponseEngineConfig.  # noqa: E501
        :type: str
        """

        self._type = type

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
        if not isinstance(other, CreateClusterClusterResponseEngineConfig):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
