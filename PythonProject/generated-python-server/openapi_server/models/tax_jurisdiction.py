from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from openapi_server.models.base_model import Model
from openapi_server.models.tax_jurisdiction_type import TaxJurisdictionType
from openapi_server import util

from openapi_server.models.tax_jurisdiction_type import TaxJurisdictionType  # noqa: E501

class TaxJurisdiction(Model):
    """NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

    Do not edit the class manually.
    """

    def __init__(self, code=None, type=None, name=None):  # noqa: E501
        """TaxJurisdiction - a model defined in OpenAPI

        :param code: The code of this TaxJurisdiction.  # noqa: E501
        :type code: str
        :param type: The type of this TaxJurisdiction.  # noqa: E501
        :type type: TaxJurisdictionType
        :param name: The name of this TaxJurisdiction.  # noqa: E501
        :type name: str
        """
        self.openapi_types = {
            'code': str,
            'type': TaxJurisdictionType,
            'name': str
        }

        self.attribute_map = {
            'code': 'code',
            'type': 'type',
            'name': 'name'
        }

        self._code = code
        self._type = type
        self._name = name

    @classmethod
    def from_dict(cls, dikt) -> 'TaxJurisdiction':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The TaxJurisdiction of this TaxJurisdiction.  # noqa: E501
        :rtype: TaxJurisdiction
        """
        return util.deserialize_model(dikt, cls)

    @property
    def code(self) -> str:
        """Gets the code of this TaxJurisdiction.

        The jurisdiction code.  # noqa: E501

        :return: The code of this TaxJurisdiction.
        :rtype: str
        """
        return self._code

    @code.setter
    def code(self, code: str):
        """Sets the code of this TaxJurisdiction.

        The jurisdiction code.  # noqa: E501

        :param code: The code of this TaxJurisdiction.
        :type code: str
        """
        if code is None:
            raise ValueError("Invalid value for `code`, must not be `None`")  # noqa: E501
        if code is not None and len(code) > 50:
            raise ValueError("Invalid value for `code`, length must be less than or equal to `50`")  # noqa: E501

        self._code = code

    @property
    def type(self) -> TaxJurisdictionType:
        """Gets the type of this TaxJurisdiction.


        :return: The type of this TaxJurisdiction.
        :rtype: TaxJurisdictionType
        """
        return self._type

    @type.setter
    def type(self, type: TaxJurisdictionType):
        """Sets the type of this TaxJurisdiction.


        :param type: The type of this TaxJurisdiction.
        :type type: TaxJurisdictionType
        """
        if type is None:
            raise ValueError("Invalid value for `type`, must not be `None`")  # noqa: E501

        self._type = type

    @property
    def name(self) -> str:
        """Gets the name of this TaxJurisdiction.

        The jurisdiction name.  # noqa: E501

        :return: The name of this TaxJurisdiction.
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name: str):
        """Sets the name of this TaxJurisdiction.

        The jurisdiction name.  # noqa: E501

        :param name: The name of this TaxJurisdiction.
        :type name: str
        """
        if name is None:
            raise ValueError("Invalid value for `name`, must not be `None`")  # noqa: E501
        if name is not None and len(name) > 50:
            raise ValueError("Invalid value for `name`, length must be less than or equal to `50`")  # noqa: E501

        self._name = name
