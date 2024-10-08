from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from openapi_server.models.base_model import Model
from openapi_server.models.credential_status import CredentialStatus
from openapi_server import util

from openapi_server.models.credential_status import CredentialStatus  # noqa: E501

class CredentialValidationResponse(Model):
    """NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

    Do not edit the class manually.
    """

    def __init__(self, status=None):  # noqa: E501
        """CredentialValidationResponse - a model defined in OpenAPI

        :param status: The status of this CredentialValidationResponse.  # noqa: E501
        :type status: CredentialStatus
        """
        self.openapi_types = {
            'status': CredentialStatus
        }

        self.attribute_map = {
            'status': 'status'
        }

        self._status = status

    @classmethod
    def from_dict(cls, dikt) -> 'CredentialValidationResponse':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The CredentialValidationResponse of this CredentialValidationResponse.  # noqa: E501
        :rtype: CredentialValidationResponse
        """
        return util.deserialize_model(dikt, cls)

    @property
    def status(self) -> CredentialStatus:
        """Gets the status of this CredentialValidationResponse.


        :return: The status of this CredentialValidationResponse.
        :rtype: CredentialStatus
        """
        return self._status

    @status.setter
    def status(self, status: CredentialStatus):
        """Sets the status of this CredentialValidationResponse.


        :param status: The status of this CredentialValidationResponse.
        :type status: CredentialStatus
        """
        if status is None:
            raise ValueError("Invalid value for `status`, must not be `None`")  # noqa: E501

        self._status = status
