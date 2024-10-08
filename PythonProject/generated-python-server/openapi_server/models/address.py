from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from openapi_server.models.base_model import Model
from openapi_server import util


class Address(Model):
    """NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

    Do not edit the class manually.
    """

    def __init__(self, line1=None, line2=None, line3=None, city=None, state=None, postal_code=None, country=None):  # noqa: E501
        """Address - a model defined in OpenAPI

        :param line1: The line1 of this Address.  # noqa: E501
        :type line1: str
        :param line2: The line2 of this Address.  # noqa: E501
        :type line2: str
        :param line3: The line3 of this Address.  # noqa: E501
        :type line3: str
        :param city: The city of this Address.  # noqa: E501
        :type city: str
        :param state: The state of this Address.  # noqa: E501
        :type state: str
        :param postal_code: The postal_code of this Address.  # noqa: E501
        :type postal_code: str
        :param country: The country of this Address.  # noqa: E501
        :type country: str
        """
        self.openapi_types = {
            'line1': str,
            'line2': str,
            'line3': str,
            'city': str,
            'state': str,
            'postal_code': str,
            'country': str
        }

        self.attribute_map = {
            'line1': 'line1',
            'line2': 'line2',
            'line3': 'line3',
            'city': 'city',
            'state': 'state',
            'postal_code': 'postalCode',
            'country': 'country'
        }

        self._line1 = line1
        self._line2 = line2
        self._line3 = line3
        self._city = city
        self._state = state
        self._postal_code = postal_code
        self._country = country

    @classmethod
    def from_dict(cls, dikt) -> 'Address':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The Address of this Address.  # noqa: E501
        :rtype: Address
        """
        return util.deserialize_model(dikt, cls)

    @property
    def line1(self) -> str:
        """Gets the line1 of this Address.

        First line of the street address  # noqa: E501

        :return: The line1 of this Address.
        :rtype: str
        """
        return self._line1

    @line1.setter
    def line1(self, line1: str):
        """Sets the line1 of this Address.

        First line of the street address  # noqa: E501

        :param line1: The line1 of this Address.
        :type line1: str
        """
        if line1 is not None and len(line1) > 180:
            raise ValueError("Invalid value for `line1`, length must be less than or equal to `180`")  # noqa: E501

        self._line1 = line1

    @property
    def line2(self) -> str:
        """Gets the line2 of this Address.

        Second line of the street address  # noqa: E501

        :return: The line2 of this Address.
        :rtype: str
        """
        return self._line2

    @line2.setter
    def line2(self, line2: str):
        """Sets the line2 of this Address.

        Second line of the street address  # noqa: E501

        :param line2: The line2 of this Address.
        :type line2: str
        """
        if line2 is not None and len(line2) > 150:
            raise ValueError("Invalid value for `line2`, length must be less than or equal to `150`")  # noqa: E501

        self._line2 = line2

    @property
    def line3(self) -> str:
        """Gets the line3 of this Address.

        Third line of the street address  # noqa: E501

        :return: The line3 of this Address.
        :rtype: str
        """
        return self._line3

    @line3.setter
    def line3(self, line3: str):
        """Sets the line3 of this Address.

        Third line of the street address  # noqa: E501

        :param line3: The line3 of this Address.
        :type line3: str
        """
        if line3 is not None and len(line3) > 150:
            raise ValueError("Invalid value for `line3`, length must be less than or equal to `150`")  # noqa: E501

        self._line3 = line3

    @property
    def city(self) -> str:
        """Gets the city of this Address.

        The city of the address  # noqa: E501

        :return: The city of this Address.
        :rtype: str
        """
        return self._city

    @city.setter
    def city(self, city: str):
        """Sets the city of this Address.

        The city of the address  # noqa: E501

        :param city: The city of this Address.
        :type city: str
        """
        if city is not None and len(city) > 50:
            raise ValueError("Invalid value for `city`, length must be less than or equal to `50`")  # noqa: E501

        self._city = city

    @property
    def state(self) -> str:
        """Gets the state of this Address.

        The state of the address following the ISO 3166-2 state/province code without the country prefix.  # noqa: E501

        :return: The state of this Address.
        :rtype: str
        """
        return self._state

    @state.setter
    def state(self, state: str):
        """Sets the state of this Address.

        The state of the address following the ISO 3166-2 state/province code without the country prefix.  # noqa: E501

        :param state: The state of this Address.
        :type state: str
        """
        if state is not None and len(state) > 50:
            raise ValueError("Invalid value for `state`, length must be less than or equal to `50`")  # noqa: E501

        self._state = state

    @property
    def postal_code(self) -> str:
        """Gets the postal_code of this Address.

        Postal Code / Zip Code of the address.  # noqa: E501

        :return: The postal_code of this Address.
        :rtype: str
        """
        return self._postal_code

    @postal_code.setter
    def postal_code(self, postal_code: str):
        """Sets the postal_code of this Address.

        Postal Code / Zip Code of the address.  # noqa: E501

        :param postal_code: The postal_code of this Address.
        :type postal_code: str
        """
        if postal_code is not None and len(postal_code) > 20:
            raise ValueError("Invalid value for `postal_code`, length must be less than or equal to `20`")  # noqa: E501

        self._postal_code = postal_code

    @property
    def country(self) -> str:
        """Gets the country of this Address.

        The country of the address following the ISO 3166-1 alpha-2 standard.  # noqa: E501

        :return: The country of this Address.
        :rtype: str
        """
        return self._country

    @country.setter
    def country(self, country: str):
        """Sets the country of this Address.

        The country of the address following the ISO 3166-1 alpha-2 standard.  # noqa: E501

        :param country: The country of this Address.
        :type country: str
        """
        if country is not None and len(country) > 2:
            raise ValueError("Invalid value for `country`, length must be less than or equal to `2`")  # noqa: E501

        self._country = country
