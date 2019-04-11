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


class CreateClusterClusterRequestSecuritySettings(object):
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
        'encrypted_ephemerals': 'bool',
        'persistent_security_groups': 'str',
        'ssh_public_key': 'str'
    }

    attribute_map = {
        'encrypted_ephemerals': 'encrypted_ephemerals',
        'persistent_security_groups': 'persistent_security_groups',
        'ssh_public_key': 'ssh_public_key'
    }

    def __init__(self, encrypted_ephemerals=None, persistent_security_groups=None, ssh_public_key=None):  # noqa: E501
        """CreateClusterClusterRequestSecuritySettings - a model defined in OpenAPI"""  # noqa: E501

        self._encrypted_ephemerals = None
        self._persistent_security_groups = None
        self._ssh_public_key = None
        self.discriminator = None

        if encrypted_ephemerals is not None:
            self.encrypted_ephemerals = encrypted_ephemerals
        if persistent_security_groups is not None:
            self.persistent_security_groups = persistent_security_groups
        if ssh_public_key is not None:
            self.ssh_public_key = ssh_public_key

    @property
    def encrypted_ephemerals(self):
        """Gets the encrypted_ephemerals of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501


        :return: The encrypted_ephemerals of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501
        :rtype: bool
        """
        return self._encrypted_ephemerals

    @encrypted_ephemerals.setter
    def encrypted_ephemerals(self, encrypted_ephemerals):
        """Sets the encrypted_ephemerals of this CreateClusterClusterRequestSecuritySettings.


        :param encrypted_ephemerals: The encrypted_ephemerals of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501
        :type: bool
        """

        self._encrypted_ephemerals = encrypted_ephemerals

    @property
    def persistent_security_groups(self):
        """Gets the persistent_security_groups of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501


        :return: The persistent_security_groups of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501
        :rtype: str
        """
        return self._persistent_security_groups

    @persistent_security_groups.setter
    def persistent_security_groups(self, persistent_security_groups):
        """Sets the persistent_security_groups of this CreateClusterClusterRequestSecuritySettings.


        :param persistent_security_groups: The persistent_security_groups of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501
        :type: str
        """

        self._persistent_security_groups = persistent_security_groups

    @property
    def ssh_public_key(self):
        """Gets the ssh_public_key of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501


        :return: The ssh_public_key of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501
        :rtype: str
        """
        return self._ssh_public_key

    @ssh_public_key.setter
    def ssh_public_key(self, ssh_public_key):
        """Sets the ssh_public_key of this CreateClusterClusterRequestSecuritySettings.


        :param ssh_public_key: The ssh_public_key of this CreateClusterClusterRequestSecuritySettings.  # noqa: E501
        :type: str
        """

        self._ssh_public_key = ssh_public_key

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
        if not isinstance(other, CreateClusterClusterRequestSecuritySettings):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
