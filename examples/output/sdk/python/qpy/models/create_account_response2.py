# coding: utf-8

"""
    Quubole API Sepcification

    Operations, requests and responses  # noqa: E501

    OpenAPI spec version: v2.0
    Contact: support@qubole.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six


class CreateAccountResponse2(object):
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
        'account_id': 'str',
        'authentication_token': 'str',
        'groups': 'list[str]',
        'is_aws_keys_enabled': 'bool',
        'name': 'str',
        'state': 'str',
        'status': 'str'
    }

    attribute_map = {
        'account_id': 'account_id',
        'authentication_token': 'authentication_token',
        'groups': 'groups',
        'is_aws_keys_enabled': 'is_aws_keys_enabled',
        'name': 'name',
        'state': 'state',
        'status': 'status'
    }

    def __init__(self, account_id=None, authentication_token=None, groups=None, is_aws_keys_enabled=None, name=None, state=None, status=None):  # noqa: E501
        """CreateAccountResponse2 - a model defined in OpenAPI"""  # noqa: E501

        self._account_id = None
        self._authentication_token = None
        self._groups = None
        self._is_aws_keys_enabled = None
        self._name = None
        self._state = None
        self._status = None
        self.discriminator = None

        if account_id is not None:
            self.account_id = account_id
        if authentication_token is not None:
            self.authentication_token = authentication_token
        if groups is not None:
            self.groups = groups
        if is_aws_keys_enabled is not None:
            self.is_aws_keys_enabled = is_aws_keys_enabled
        if name is not None:
            self.name = name
        if state is not None:
            self.state = state
        if status is not None:
            self.status = status

    @property
    def account_id(self):
        """Gets the account_id of this CreateAccountResponse2.  # noqa: E501


        :return: The account_id of this CreateAccountResponse2.  # noqa: E501
        :rtype: str
        """
        return self._account_id

    @account_id.setter
    def account_id(self, account_id):
        """Sets the account_id of this CreateAccountResponse2.


        :param account_id: The account_id of this CreateAccountResponse2.  # noqa: E501
        :type: str
        """

        self._account_id = account_id

    @property
    def authentication_token(self):
        """Gets the authentication_token of this CreateAccountResponse2.  # noqa: E501


        :return: The authentication_token of this CreateAccountResponse2.  # noqa: E501
        :rtype: str
        """
        return self._authentication_token

    @authentication_token.setter
    def authentication_token(self, authentication_token):
        """Sets the authentication_token of this CreateAccountResponse2.


        :param authentication_token: The authentication_token of this CreateAccountResponse2.  # noqa: E501
        :type: str
        """

        self._authentication_token = authentication_token

    @property
    def groups(self):
        """Gets the groups of this CreateAccountResponse2.  # noqa: E501


        :return: The groups of this CreateAccountResponse2.  # noqa: E501
        :rtype: list[str]
        """
        return self._groups

    @groups.setter
    def groups(self, groups):
        """Sets the groups of this CreateAccountResponse2.


        :param groups: The groups of this CreateAccountResponse2.  # noqa: E501
        :type: list[str]
        """

        self._groups = groups

    @property
    def is_aws_keys_enabled(self):
        """Gets the is_aws_keys_enabled of this CreateAccountResponse2.  # noqa: E501


        :return: The is_aws_keys_enabled of this CreateAccountResponse2.  # noqa: E501
        :rtype: bool
        """
        return self._is_aws_keys_enabled

    @is_aws_keys_enabled.setter
    def is_aws_keys_enabled(self, is_aws_keys_enabled):
        """Sets the is_aws_keys_enabled of this CreateAccountResponse2.


        :param is_aws_keys_enabled: The is_aws_keys_enabled of this CreateAccountResponse2.  # noqa: E501
        :type: bool
        """

        self._is_aws_keys_enabled = is_aws_keys_enabled

    @property
    def name(self):
        """Gets the name of this CreateAccountResponse2.  # noqa: E501


        :return: The name of this CreateAccountResponse2.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this CreateAccountResponse2.


        :param name: The name of this CreateAccountResponse2.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def state(self):
        """Gets the state of this CreateAccountResponse2.  # noqa: E501


        :return: The state of this CreateAccountResponse2.  # noqa: E501
        :rtype: str
        """
        return self._state

    @state.setter
    def state(self, state):
        """Sets the state of this CreateAccountResponse2.


        :param state: The state of this CreateAccountResponse2.  # noqa: E501
        :type: str
        """

        self._state = state

    @property
    def status(self):
        """Gets the status of this CreateAccountResponse2.  # noqa: E501


        :return: The status of this CreateAccountResponse2.  # noqa: E501
        :rtype: str
        """
        return self._status

    @status.setter
    def status(self, status):
        """Sets the status of this CreateAccountResponse2.


        :param status: The status of this CreateAccountResponse2.  # noqa: E501
        :type: str
        """

        self._status = status

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
        if not isinstance(other, CreateAccountResponse2):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
